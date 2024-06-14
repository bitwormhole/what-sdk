package units

import (
	"crypto/sha256"
	"time"

	"github.com/bitwormhole/what-sdk/v1/what"
	"github.com/starter-go/application"
)

// TryListenSimple ...
type TryListenSimple struct {

	//starter:component

	EngineHost string //starter:inject("${debug.engine.host}")
	EnginePort int    //starter:inject("${debug.engine.port}")

	RemoteService string //starter:inject("${debug.case.listen-as-simple.service}")
	RemoteURL     string //starter:inject("${debug.case.listen-as-simple.url}")
	Enabled       bool   //starter:inject("${debug.case.listen-as-simple.enabled}")

}

func (inst *TryListenSimple) _impl() application.Lifecycle {
	return inst
}

// Life ...
func (inst *TryListenSimple) Life() *application.Life {

	if !inst.Enabled {
		return &application.Life{}
	}

	return &application.Life{
		OnLoop:  inst.loop,
		OnStart: inst.start,
	}
}

func (inst *TryListenSimple) loop() error {
	for {
		time.Sleep(time.Second)
	}
}

func (inst *TryListenSimple) start() error {
	go func() {
		time.Sleep(time.Second) // delay
		err := inst.run()
		what.HandleError(err)
	}()
	return nil
}

func (inst *TryListenSimple) run() error {

	cfg := &what.Configuration{
		Host:   inst.EngineHost,
		Port:   inst.EnginePort,
		UseTLS: false,
	}

	p := &what.ListenerParams{
		Mode:    what.ModeSimple,
		Service: inst.RemoteService,
	}

	agent, err := what.New(cfg)
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

	conn, err := li.Accept()
	if err != nil {
		return err
	}
	defer func() {
		conn.Close()
	}()

	conn2 := conn.(what.SimpleConnection)
	buffer := make([]byte, 1024*2)

	for {
		cb, err := conn2.Read(buffer)
		if err != nil {
			return err
		}
		data := buffer[0:cb]
		err = inst.handleRxData(data, conn2)
		if err != nil {
			return err
		}
	}
	// return nil
}

func (inst *TryListenSimple) handleRxData(data []byte, conn what.SimpleConnection) error {
	sum := sha256.Sum256(data)
	_, err := conn.Write(sum[:])
	return err
}
