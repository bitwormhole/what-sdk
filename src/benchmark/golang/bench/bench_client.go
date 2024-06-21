package bench

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/bitwormhole/what-sdk/v1/what"
	"github.com/starter-go/application"
	"github.com/starter-go/base/lang"
	"github.com/starter-go/vlog"
)

// Client ...
type Client struct {

	//starter:component

	Agent Agent //starter:inject("#")

	ServiceName string //starter:inject("${benchmark.client.service}")
	Enabled     bool   //starter:inject("${benchmark.client.enabled}")
	Interval    int    //starter:inject("${benchmark.client.interval}")
	RequestX    int    //starter:inject("${benchmark.client.request-x}")
	ResponseX   int    //starter:inject("${benchmark.client.response-x}")
	PayloadSize int    //starter:inject("${benchmark.client.payload-size}")
	RemoteURL   string //starter:inject("${benchmark.client.remote-url}")

}

func (inst *Client) _impl() application.Lifecycle {
	return inst
}

// Life ...
func (inst *Client) Life() *application.Life {

	if !inst.Enabled {
		return &application.Life{}
	}

	return &application.Life{
		OnStart: inst.start,
		OnLoop:  inst.loop,
	}
}

func (inst *Client) loop() error {
	for {
		time.Sleep(time.Second)
	}
	// return nil
}

func (inst *Client) checkProperties() error {

	// for string values

	if inst.ServiceName == "" {
		return fmt.Errorf("bad property: ServiceName is empty")
	}

	if inst.RemoteURL == "" {
		return fmt.Errorf("bad property: RemoteURL is empty")
	}

	// for int values
	checkIntRange := func(name string, value, min, max int) error {
		if (min <= value) && (value <= max) {
			return nil
		}
		const f = "the int property value is out of range: [name:'%s' value:%d min:%d max:%d]"
		return fmt.Errorf(f, name, value, min, max)
	}

	e1 := checkIntRange("interval", inst.Interval, 10, 10000)
	e2 := checkIntRange("request-x", inst.RequestX, 0, 1000)
	e3 := checkIntRange("response-x", inst.ResponseX, 0, 1000)
	e4 := checkIntRange("payload-size", inst.PayloadSize, 0, 65536)

	errlist := []error{e1, e2, e3, e4}
	for _, err := range errlist {
		if err != nil {
			return err
		}
	}
	return nil
}

func (inst *Client) start() error {

	err := inst.checkProperties()
	if err != nil {
		return err
	}

	go func() {
		time.Sleep(time.Second * 3) // delay
		err := inst.run()
		what.HandleError(err)
	}()
	return nil
}

func (inst *Client) run() error {

	p := &what.ConnectionParams{
		Mode:    what.ModeSimple,
		Service: inst.ServiceName,
		URL:     inst.RemoteURL,
	}

	agent, err := inst.Agent.GetAgent(IndexClient)
	if err != nil {
		return err
	}

	conn, err := agent.Connect(p)
	if err != nil {
		return err
	}
	defer func() {
		conn.Close()
	}()

	conn2 := conn.(what.SimpleConnection)
	task := &clientSideTask{
		conn:   conn2,
		client: inst,
	}
	task.init()
	return task.run()
}

////////////////////////////////////////////////////////////////////////////////

type clientSideTask struct {
	conn       what.SimpleConnection
	client     *Client
	counter    Counter
	remoteSnap Counter
}

func (inst *clientSideTask) init() error {

	now := lang.Now()
	inst.counter.Now = now
	inst.counter.Startup = now

	return nil
}

func (inst *clientSideTask) run() error {

	client := inst.client
	interval := lang.Milliseconds(client.Interval)
	step := interval.Duration()
	countTimer := time.Second

	go func() {
		err := inst.pullDownStream()
		what.HandleError(err)
	}()

	for {

		err := inst.sendPayloadRequestGroup()
		if err != nil {
			return err
		}

		if countTimer > time.Second {
			countTimer -= time.Second
			inst.sendCountRequest()
		}

		countTimer += step
		time.Sleep(step)
	}
}

func (inst *clientSideTask) sendCountRequest() error {
	inst.printCounterInfo()
	p := &Pack{
		Type: TypeCountRequest,
	}
	return inst.sendRequest(p)
}

func (inst *clientSideTask) sendPayloadRequest() error {

	client := inst.client
	size := client.PayloadSize
	data := make([]byte, size)

	p := &Pack{
		Type: TypePayloadRequest,
	}
	p.RequestX = client.RequestX
	p.ResponseX = client.ResponseX
	p.body = data

	return inst.sendRequest(p)
}

func (inst *clientSideTask) sendPayloadRequestGroup() error {
	count := inst.client.RequestX
	if count < 1 {
		return nil
	}
	for i := 0; i < count; i++ {
		err := inst.sendPayloadRequest()
		if err != nil {
			return err
		}
	}
	return nil
}

func (inst *clientSideTask) sendRequest(p *Pack) error {
	raw := Marshal(p)
	_, err := inst.conn.Write(raw)

	size := len(raw)
	inst.counter.Now = lang.Now()
	inst.counter.TxBytes += int64(size)
	inst.counter.TxPacks++

	return err
}

func (inst *clientSideTask) pullDownStream() error {
	buffer := make([]byte, 1024*2)
	conn := inst.conn
	for {
		cb, err := conn.Read(buffer)
		if err != nil {
			return err
		}
		data := buffer[0:cb]
		err = inst.handleResponseData(data)
		what.HandleError(err)
	}
}

func (inst *clientSideTask) handleResponseData(data []byte) error {

	size := len(data)
	inst.counter.Now = lang.Now()
	inst.counter.RxBytes += int64(size)
	inst.counter.RxPacks++

	resp, err := Unmarshal(data)
	if err != nil {
		return err
	}

	switch resp.Type {
	case TypeCountResponse:
		return inst.handleCountResponse(resp)
	case TypePayloadResponse:
		return inst.handlePayloadResponse(resp)
	}
	return fmt.Errorf("unsupported response pack type:%s", resp.Type)
}

func (inst *clientSideTask) handlePayloadResponse(resp *Pack) error {
	return nil
}

func (inst *clientSideTask) handleCountResponse(resp *Pack) error {
	cnt, err := ParseCounter(resp.body)
	if err != nil {
		return err
	}
	inst.remoteSnap = *cnt
	return nil
}

func (inst *clientSideTask) printCounterInfo() {

	table := make(map[string]*Counter)
	table["serverside"] = &inst.remoteSnap
	table["clientside"] = &inst.counter

	j, err := json.MarshalIndent(table, "", "\t")
	if err != nil {
		return
	}
	str := string(j)

	speed1 := inst.computeSpeedInfoText(&inst.counter)
	speed2 := inst.computeSpeedInfoText(&inst.remoteSnap)

	vlog.Info("count_info: %s", str)
	vlog.Info("speed(client_side): %s", speed1)
	vlog.Info("speed(server_side): %s", speed2)
}

func (inst *clientSideTask) computeSpeedInfoText(cnt *Counter) string {

	t0 := cnt.Startup
	t1 := cnt.Now

	span := (t1 - t0)
	spanf := float32(span) / 1000
	if span < 1 {
		spanf = 0.0001
	}

	const (
		rxb = "rx.bytes"
		rxp = "rx.packs"
		txb = "tx.bytes"
		txp = "tx.packs"
	)

	keys := []string{rxp, rxb, txp, txb}
	table := map[string]int64{
		rxb: cnt.RxBytes,
		rxp: cnt.RxPacks,
		txb: cnt.TxBytes,
		txp: cnt.TxPacks,
	}
	builder := &strings.Builder{}

	for _, key := range keys {
		value := table[key]
		v2 := float32(value) / spanf
		v3 := int64(v2)
		text := strconv.FormatInt(v3, 10)
		builder.WriteString(" ")
		builder.WriteString(key)
		builder.WriteString(": ")
		builder.WriteString(text)
	}

	return builder.String()
}
