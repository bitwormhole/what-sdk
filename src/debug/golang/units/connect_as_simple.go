package units

import (
	"time"

	"github.com/bitwormhole/what-sdk/v1/what"
	"github.com/starter-go/application"
	"github.com/starter-go/base/lang"
	"github.com/starter-go/vlog"
)

// TryConnectAsSimple ...
type TryConnectAsSimple struct {

	//starter:component

	EngineHost string //starter:inject("${debug.engine.host}")
	EnginePort int    //starter:inject("${debug.engine.port}")

	RemoteService string //starter:inject("${debug.case.connect-as-simple.service}")
	RemoteURL     string //starter:inject("${debug.case.connect-as-simple.url}")
	Enabled       bool   //starter:inject("${debug.case.connect-as-simple.enabled}")

}

func (inst *TryConnectAsSimple) _impl() application.Lifecycle {
	return inst
}

// Life ...
func (inst *TryConnectAsSimple) Life() *application.Life {

	if !inst.Enabled {
		return &application.Life{}
	}

	return &application.Life{
		OnLoop:  inst.loop,
		OnStart: inst.start,
	}
}

func (inst *TryConnectAsSimple) loop() error {
	for {
		time.Sleep(time.Second)
	}
}

func (inst *TryConnectAsSimple) start() error {
	go func() {
		time.Sleep(time.Second * 3) // delay
		err := inst.run()
		what.HandleError(err)
	}()
	return nil
}

func (inst *TryConnectAsSimple) run() error {

	cfg := &what.Configuration{
		Host:   inst.EngineHost,
		Port:   inst.EnginePort,
		UseTLS: false,
	}

	p := &what.ConnectionParams{
		Mode:    what.ModeSimple,
		Service: inst.RemoteService,
		URL:     inst.RemoteURL,
	}

	agent, err := what.New(cfg)
	if err != nil {
		return err
	}

	conn, err := agent.Connect(p)
	if err != nil {
		return err
	}
	defer conn.Close()

	conn2 := conn.(what.SimpleConnection)
	buffer := make([]byte, 10)

	go func() {
		err := inst.pumpRxDataStream(conn2)
		what.HandleError(err)
	}()

	for {
		cb, err := conn2.Write(buffer)
		if err != nil {
			return err
		}
		if cb == 0 {
			break
		}
		time.Sleep(time.Second)
	}

	return nil
}

func (inst *TryConnectAsSimple) pumpRxDataStream(conn what.SimpleConnection) error {
	buffer := make([]byte, rxBufferSize)
	for {
		cb, err := conn.Read(buffer)
		if err != nil {
			return err
		}
		data := buffer[0:cb]
		hex := lang.HexFromBytes(data)
		vlog.Info("TryConnectAsSimple.pumpRxDataStream(): receive data [%s]", hex.String())
	}
}
