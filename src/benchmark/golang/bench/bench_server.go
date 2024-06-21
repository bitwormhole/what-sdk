package bench

import (
	"fmt"
	"time"

	"github.com/bitwormhole/what-sdk/v1/what"
	"github.com/starter-go/application"
	"github.com/starter-go/base/lang"
)

// Server ...
type Server struct {

	//starter:component

	Agent Agent //starter:inject("#")

	ServiceName string //starter:inject("${benchmark.server.service}")
	Enabled     bool   //starter:inject("${benchmark.server.enabled}")

}

func (inst *Server) _impl() application.Lifecycle {
	return inst
}

// Life ...
func (inst *Server) Life() *application.Life {

	if !inst.Enabled {
		return &application.Life{}
	}

	return &application.Life{
		OnStart: inst.start,
		OnLoop:  inst.loop,
	}
}

func (inst *Server) loop() error {
	for {
		time.Sleep(time.Second)
	}
	// return nil
}

func (inst *Server) start() error {
	go func() {
		err := inst.run()
		what.HandleError(err)
	}()
	return nil
}

func (inst *Server) run() error {

	p := &what.ListenerParams{
		Mode:    what.ModeSimple,
		Service: inst.ServiceName,
	}

	agent, err := inst.Agent.GetAgent(IndexServer)
	if err != nil {
		return err
	}

	li, err := agent.Listen(p)
	if err != nil {
		return err
	}

	defer func() {
		li.Close()
	}()

	for {
		conn, err := li.Accept()
		if err != nil {
			return err
		}
		inst.handleConnection(conn.(what.SimpleConnection))
	}
}

func (inst *Server) handleConnection(conn what.SimpleConnection) {
	task := &serverSideTask{
		conn: conn,
	}
	go func() {
		err := task.run()
		what.HandleError(err)
	}()
}

////////////////////////////////////////////////////////////////////////////////

type serverSideTask struct {
	conn    what.SimpleConnection
	counter Counter
}

func (inst *serverSideTask) run() error {

	defer func() {
		inst.conn.Close()
	}()

	now := lang.Now()
	inst.counter.Startup = now
	inst.counter.Now = now

	buffer := make([]byte, rxBufferSize)

	for {
		cb, err := inst.conn.Read(buffer)
		if err != nil {
			return err
		}

		now = lang.Now()
		inst.counter.Now = now
		inst.counter.RxBytes += int64(cb)
		inst.counter.RxPacks++

		data := buffer[0:cb]
		err = inst.handleRxData(data)
		what.HandleError(err)
	}
	// return nil
}

func (inst *serverSideTask) handleRxData(data1 []byte) error {

	// decode request
	req, err := Unmarshal(data1)
	if err != nil {
		return err
	}

	// handle
	resp, err := inst.handleRequest(req)
	if err != nil {
		return err
	}

	// encode response
	return inst.sendPack(resp)
}

func (inst *serverSideTask) sendPack(p *Pack) error {
	data2 := Marshal(p)
	return inst.sendTxData(data2)
}

func (inst *serverSideTask) sendTxData(data []byte) error {

	size := len(data)
	now := lang.Now()
	inst.counter.Now = now
	inst.counter.TxBytes += int64(size)
	inst.counter.TxPacks++

	_, err := inst.conn.Write(data)
	return err
}

func (inst *serverSideTask) handleRequest(req *Pack) (*Pack, error) {
	switch req.Type {
	case TypePayloadRequest:
		return inst.handlePayloadRequest(req)
	case TypeCountRequest:
		return inst.handleCountRequest(req)
	}
	return nil, fmt.Errorf("unsupported request type: %s", req.Type)
}

func (inst *serverSideTask) handlePayloadRequest(req *Pack) (*Pack, error) {

	resp := &Pack{}
	resp.Type = TypePayloadResponse
	resp.body = req.body
	resp.RequestX = req.RequestX
	resp.ResponseX = req.ResponseX

	count := req.ResponseX
	for i := 0; i < count; i++ {
		inst.sendPack(resp)
	}

	resp.body = nil
	return resp, nil
}

func (inst *serverSideTask) handleCountRequest(req *Pack) (*Pack, error) {

	if req.Type != TypeCountRequest {
		return nil, fmt.Errorf("bad request pack type: %s", req.Type)
	}

	j := inst.counter.Bytes()
	resp := &Pack{}
	resp.Type = TypeCountResponse
	resp.body = j
	return resp, nil
}
