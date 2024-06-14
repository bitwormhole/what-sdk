package bench

import (
	"fmt"

	"github.com/bitwormhole/what-sdk/v1/what"
)

const (
	IndexClient = 1
	IndexServer = 2
)

// Agent ... 配置并代理 what.Agent
type Agent interface {
	GetAgent(index int) (what.Agent, error)
}

////////////////////////////////////////////////////////////////////////////////

// ProxyAgent ... 实现: Agent
type ProxyAgent struct {

	//starter:component

	_as func(Agent) //starter:as("#")

	Host1 string //starter:inject("${benchmark.core1.host}")
	Port1 int    //starter:inject("${benchmark.core1.port}")

	Host2 string //starter:inject("${benchmark.core2.host}")
	Port2 int    //starter:inject("${benchmark.core2.port}")

}

func (inst *ProxyAgent) _impl() Agent {
	return inst
}

// GetAgent ...
func (inst *ProxyAgent) GetAgent(index int) (what.Agent, error) {
	switch index {
	case 1:
		return inst.loadAgent1()
	case 2:
		return inst.loadAgent2()
	default:
	}
	return nil, fmt.Errorf("bad index(value=%d) of core config", index)
}

func (inst *ProxyAgent) loadAgent1() (what.Agent, error) {
	cfg := &what.Configuration{}
	cfg.Host = inst.Host1
	cfg.Port = inst.Port1
	return what.New(cfg)
}

func (inst *ProxyAgent) loadAgent2() (what.Agent, error) {
	cfg := &what.Configuration{}
	cfg.Host = inst.Host2
	cfg.Port = inst.Port2
	return what.New(cfg)
}
