package what

import (
	"crypto/tls"
	"io"
	"net"
	"strconv"

	"github.com/starter-go/pktline"
)

// PktlineConnection ...
type PktlineConnection interface {
	Connection

	pktline.Reader

	pktline.Writer
}

////////////////////////////////////////////////////////////////////////////////

// PktlineConnector ...
type PktlineConnector struct {
}

func (inst *PktlineConnector) _impl() Connector {
	return inst
}

// Registration ...
func (inst *PktlineConnector) Registration() *ConnectorRegistration {
	return &ConnectorRegistration{
		Priority:  1,
		Enabled:   true,
		Name:      "PktlineConnector",
		Mode:      ModePktline,
		Connector: inst,
		Support:   nil,
	}
}

// Connect ...
func (inst *PktlineConnector) Connect(c *Context, p *ConnectionParams) (Connection, error) {
	conn := &defaultPktlineConnection{context: c}
	err := conn.open()
	if err != nil {
		return nil, err
	}
	return conn, nil
}

////////////////////////////////////////////////////////////////////////////////

type defaultPktlineConnection struct {
	context     *Context
	conn1       net.Conn
	conn2       net.Conn
	innerReader pktline.Reader
	innerWriter pktline.Writer
}

func (inst *defaultPktlineConnection) _impl() PktlineConnection {
	return inst
}

func (inst *defaultPktlineConnection) open() error {

	config1 := &inst.context.Config
	host := config1.Host
	port := config1.Port

	if host == "" {
		host = "127.0.0.1"
	}

	if port < 1 {
		port = 10217
	}

	addr := host + ":" + strconv.Itoa(port) // host:port

	// open tcp
	conn1, err := net.Dial("tcp", addr)
	if err != nil {
		return err
	}
	defer func() {
		inst.closeThis(conn1)
	}()

	// open tls
	conn2 := conn1
	if config1.UseTLS {
		config2 := &tls.Config{}
		conn2 = tls.Client(conn1, config2)
	}
	defer func() {
		inst.closeThis(conn2)
	}()

	// make reader & writer
	rdr := pktline.NewDecoder(conn2)
	wtr := pktline.NewEncoder(conn2)

	inst.innerReader = rdr
	inst.innerWriter = wtr
	inst.conn1 = conn1
	inst.conn2 = conn2
	conn1 = nil
	conn2 = nil
	return nil
}

func (inst *defaultPktlineConnection) closeThis(c io.Closer) error {
	if c == nil {
		return nil
	}
	return c.Close()
}

func (inst *defaultPktlineConnection) Close() error {
	c1 := inst.conn1
	c2 := inst.conn2
	inst.conn1 = nil
	inst.conn2 = nil
	inst.closeThis(c2)
	inst.closeThis(c1)
	return nil
}

func (inst *defaultPktlineConnection) Read() (*pktline.Packet, error) {
	return inst.innerReader.Read()
}

func (inst *defaultPktlineConnection) Write(p *pktline.Packet) error {
	return inst.innerWriter.Write(p)
}
