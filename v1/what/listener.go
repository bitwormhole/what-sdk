package what

import (
	"io"
	"net/http"
	"time"

	"github.com/starter-go/base/util"
	"github.com/starter-go/pktline"
)

// Listener 代表一个 what 服务监听器
type Listener interface {
	io.Closer

	Accept() (Connection, error)
}

////////////////////////////////////////////////////////////////////////////////

type myListenerBuilder struct {
	context *Context
}

func (inst *myListenerBuilder) create(p *ListenerParams) (Listener, error) {

	cp := &ConnectionParams{
		Mode: ModeHypertext,
	}

	req := &HypertextRequest{
		Method: http.MethodPost,
		URL:    "/listen",
	}
	req.Headers.Set("service", p.Service)
	req.Headers.Set("mode", p.Mode.String())

	conn1, err := inst.context.Agent.Connect(cp)
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

	err = resp.Error()
	if err != nil {
		return nil, err
	}

	conn1 = nil
	return inst.makeResultOK(resp, p)
}

func (inst *myListenerBuilder) makeResultOK(resp *HypertextResponse, lp *ListenerParams) (Listener, error) {
	conn := resp.Connection
	li := &myListener{
		context: inst.context,
		cl:      conn,
		pkt:     conn,
		reader:  conn,
		writer:  conn,
	}
	li.params = *lp
	return li, nil
}

////////////////////////////////////////////////////////////////////////////////

type myListener struct {
	context *Context
	cl      io.Closer
	pkt     PktlineConnection
	reader  pktline.Reader
	writer  pktline.Writer
	params  ListenerParams
}

func (inst *myListener) _impl() Listener {
	return inst
}

func (inst *myListener) Accept() (Connection, error) {

	if inst.cl == nil {
		return nil, io.EOF
	}

	p, err := inst.reader.Read()
	if err != nil {
		return nil, err
	}

	return inst.handleConnection(p)
}

func (inst *myListener) Close() error {
	cl := inst.cl
	inst.cl = nil
	if cl == nil {
		return nil
	}
	return cl.Close()
}

func (inst *myListener) handleConnection(p *pktline.Packet) (Connection, error) {

	url := p.Head

	cp := &ConnectionParams{
		Service: inst.params.Service,
		Mode:    inst.params.Mode,
		URL:     url,
		DoGet:   true,
		Timeout: time.Second * 10,
	}

	return inst.context.Agent.Connect(cp)
}
