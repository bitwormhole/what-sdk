package what

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/starter-go/pktline"
)

// HypertextHeaders ...
type HypertextHeaders struct {
	table map[string]string
}

// Get ...
func (inst *HypertextHeaders) Get(name string) string {
	name = strings.ToLower(name)
	t := inst.table
	if t == nil {
		return ""
	}
	return t[name]
}

// Set ...
func (inst *HypertextHeaders) Set(name, value string) {
	name = strings.ToLower(name)
	t := inst.table
	if t == nil {
		t = make(map[string]string)
		inst.table = t
	}
	t[name] = value
}

////////////////////////////////////////////////////////////////////////////////

// HypertextRequest ...
type HypertextRequest struct {
	Protocol string
	Method   string
	URL      string
	Headers  HypertextHeaders

	Content []byte
}

////////////////////////////////////////////////////////////////////////////////

// HypertextResponse ...
type HypertextResponse struct {
	Protocol   string
	Request    *HypertextRequest
	Status     int
	Message    string
	Headers    HypertextHeaders
	Connection PktlineConnection

	Content []byte
}

// Error 如果响应结果有误，封装成 error; 否则返回 nil
func (inst *HypertextResponse) Error() error {
	code := inst.Status
	if code == http.StatusOK {
		return nil
	}
	err := inst.Headers.Get("error")
	msg := inst.Message
	const f = "HTTP %d %s; error:[%s]"
	return fmt.Errorf(f, code, msg, err)
}

////////////////////////////////////////////////////////////////////////////////

// HypertextConnection ...
type HypertextConnection interface {
	Connection

	Do(req *HypertextRequest) (*HypertextResponse, error)
}

////////////////////////////////////////////////////////////////////////////////

// HypertextConnector ...
type HypertextConnector struct {
}

func (inst *HypertextConnector) _impl() Connector {
	return inst
}

// Registration ...
func (inst *HypertextConnector) Registration() *ConnectorRegistration {
	return &ConnectorRegistration{
		Name:      "HypertextConnector",
		Mode:      ModeHypertext,
		Enabled:   true,
		Priority:  0,
		Connector: inst,
		Support:   nil,
	}
}

// Connect ...
func (inst *HypertextConnector) Connect(c *Context, p *ConnectionParams) (Connection, error) {
	conn := &innerHypertextConnection{
		ctx:    c,
		params: p,
	}
	return conn, nil
}

////////////////////////////////////////////////////////////////////////////////

type innerHypertextConnection struct {
	ctx    *Context
	params *ConnectionParams
}

func (inst *innerHypertextConnection) _impl() HypertextConnection {
	return inst
}

func (inst *innerHypertextConnection) Close() error {
	return nil
}

func (inst *innerHypertextConnection) Do(req *HypertextRequest) (*HypertextResponse, error) {

	conn1, err := inst.openPktlineConn()
	if err != nil {
		return nil, err
	}
	defer func() {
		if conn1 != nil {
			conn1.Close()
		}
	}()

	// send request
	err = inst.prepareRequest(req)
	if err != nil {
		return nil, err
	}
	err = inst.sendRequest(conn1, req)
	if err != nil {
		return nil, err
	}

	// wait response
	resp := &HypertextResponse{}
	err = inst.waitResponse(conn1, req, resp)
	if err != nil {
		return nil, err
	}

	resp.Connection = conn1
	conn1 = nil
	return resp, nil
}

func (inst *innerHypertextConnection) openPktlineConn() (PktlineConnection, error) {

	p1 := inst.params
	p2 := &ConnectionParams{}
	*p2 = *p1
	p2.Mode = ModePktline

	agent := inst.ctx.Agent
	conn1, err := agent.Connect(p2)
	if err != nil {
		return nil, err
	}

	conn2 := conn1.(PktlineConnection)
	return conn2, nil
}

func (inst *innerHypertextConnection) prepareRequest(req *HypertextRequest) error {
	location := req.URL
	u, err := url.Parse(location)
	if err != nil {
		return err
	}

	host := u.Host

	u.Scheme = ""
	u.User = nil
	u.Host = ""

	req.URL = u.String()
	req.Headers.Set("host", host)
	return nil
}

func (inst *innerHypertextConnection) sendRequest(conn PktlineConnection, req *HypertextRequest) error {

	pro := req.Protocol
	method := req.Method
	url := req.URL
	ua := "what-sdk"

	if pro == "" {
		pro = "PKT-HTTP/1.0"
	}
	if method == "" {
		method = http.MethodGet
	}
	if url == "" {
		url = "/"
	}

	req.Headers.Set("User-Agent", ua)

	// prepare want-line
	builder := &strings.Builder{}
	builder.WriteString(method)
	builder.WriteRune(' ')
	builder.WriteString(url)
	builder.WriteRune(' ')
	builder.WriteString(pro)
	builder.WriteRune('\n')

	// prepare headers
	headers := req.Headers.table
	for name, value := range headers {
		builder.WriteString(name)
		builder.WriteRune(':')
		builder.WriteString(value)
		builder.WriteRune('\n')
	}

	// make packet
	pack := &pktline.Packet{}
	pack.Head = builder.String()
	pack.Body = req.Content

	// send
	return conn.Write(pack)
}

func (inst *innerHypertextConnection) waitResponse(conn PktlineConnection, req *HypertextRequest, resp *HypertextResponse) error {

	pack, err := conn.Read()
	if err != nil {
		return err
	}

	// parse headers
	strStatusLine := ""
	headers := &resp.Headers
	str := pack.Head
	rows := strings.Split(str, "\n")
	for index, row := range rows {
		if index == 0 {
			strStatusLine = row
			continue
		}
		parts := strings.SplitN(row, ":", 2)
		if len(parts) == 2 {
			name := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			headers.Set(name, value)
		}
	}

	// parse status-line
	parts := strings.SplitN(strStatusLine, " ", 3)
	if len(parts) == 3 {
		pro := parts[0]
		code := parts[1]
		msg := parts[2]
		status, _ := strconv.Atoi(code)
		resp.Protocol = pro
		resp.Status = status
		resp.Message = msg
	}

	// result
	resp.Connection = conn
	resp.Request = req
	resp.Content = pack.Body
	return nil
}
