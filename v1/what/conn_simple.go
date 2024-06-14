package what

import (
	"io"
	"net/http"

	"github.com/starter-go/base/util"
	"github.com/starter-go/pktline"
)

// SimpleConnection ...
type SimpleConnection interface {
	Connection

	io.Reader

	io.Writer
}

////////////////////////////////////////////////////////////////////////////////

// SimpleConnector ...
type SimpleConnector struct {
}

func (inst *SimpleConnector) _impl() Connector {
	return inst
}

// Registration ...
func (inst *SimpleConnector) Registration() *ConnectorRegistration {
	return &ConnectorRegistration{
		Enabled:   true,
		Priority:  0,
		Name:      "SimpleConnector",
		Mode:      ModeSimple,
		Connector: inst,
		Support:   nil,
	}
}

// Connect ...
func (inst *SimpleConnector) Connect(c *Context, p *ConnectionParams) (Connection, error) {

	req := &HypertextRequest{
		Method: http.MethodPost,
		URL:    "/connect",
	}
	req.Headers.Set("service", p.Service)
	req.Headers.Set("mode", p.Mode.String())
	req.Headers.Set("location", p.URL)

	if p.DoGet {
		req.URL = p.URL
		req.Method = http.MethodGet
	}

	////////////////////

	p1 := &ConnectionParams{
		Mode: ModeHypertext,
	}
	conn1, err := c.Agent.Connect(p1)
	if err != nil {
		return nil, err
	}
	defer func() {
		util.Close(conn1)
	}()

	conn2 := conn1.(HypertextConnection)
	resp, err := conn2.Do(req)
	if err != nil {
		return nil, err
	}

	conn3 := resp.Connection
	defer func() {
		err = util.Close(conn3)
		HandleError(err)
	}()

	err = resp.Error()
	if err != nil {
		return nil, err
	}

	conn4 := &defaultSimpleConnection{
		inner: conn3,
	}
	conn1 = nil
	conn2 = nil
	conn3 = nil
	return conn4, nil
}

////////////////////////////////////////////////////////////////////////////////

type defaultSimpleConnection struct {
	inner PktlineConnection
}

func (inst *defaultSimpleConnection) _impl() SimpleConnection {
	return inst
}

func (inst *defaultSimpleConnection) Close() error {
	conn := inst.inner
	inst.inner = nil
	if conn != nil {
		return conn.Close()
	}
	return nil
}

func (inst *defaultSimpleConnection) Read(dst []byte) (int, error) {
	p, err := inst.inner.Read()
	if err != nil {
		return 0, err
	}
	src := p.Body
	cnt := copy(dst, src)
	return cnt, nil
}

func (inst *defaultSimpleConnection) Write(b []byte) (int, error) {

	pack := &pktline.Packet{
		Head: "simple:/data",
		Body: b,
	}

	err := inst.inner.Write(pack)
	return len(b), err
}
