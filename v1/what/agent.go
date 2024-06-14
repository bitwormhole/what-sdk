package what

import "fmt"

// Configuration ...
type Configuration struct {
	Host      string
	Port      int
	KeyID     string
	KeySecret string
	UseTLS    bool
}

// Agent 代表一个 what 客户端
type Agent interface {
	Connect(p *ConnectionParams) (Connection, error)

	Listen(p *ListenerParams) (Listener, error)
}

// New 新建一个 Agent
func New(cfg *Configuration) (Agent, error) {

	if cfg == nil {
		return nil, fmt.Errorf("param: config is nil")
	}

	ctx := &Context{}

	ctx.Connectors = append(ctx.Connectors, &PktlineConnector{})
	ctx.Connectors = append(ctx.Connectors, &SimpleConnector{})
	ctx.Connectors = append(ctx.Connectors, &HypertextConnector{})

	ctx.Config = *cfg
	return NewWithContext(ctx)
}

// NewWithContext 新建一个 Agent
func NewWithContext(ctx *Context) (Agent, error) {

	if ctx == nil {
		return nil, fmt.Errorf("param: context is nil")
	}

	a := &defaultAgent{context: ctx}
	if ctx.Agent == nil {
		ctx.Agent = a
	}
	return a, nil
}

////////////////////////////////////////////////////////////////////////////////

type defaultAgent struct {
	context *Context
}

func (inst *defaultAgent) _impl() Agent {
	return inst
}

func (inst *defaultAgent) Connect(p *ConnectionParams) (Connection, error) {

	ctx := inst.context
	all := inst.context.Connectors
	mode := p.Mode

	for _, c := range all {
		reg := c.Registration()
		if reg.Mode == mode {
			return reg.Connector.Connect(ctx, p)
		}
	}

	return nil, fmt.Errorf("no connector supports mode [%s]", mode)
}

func (inst *defaultAgent) Listen(p *ListenerParams) (Listener, error) {
	builder := &myListenerBuilder{
		context: inst.context,
	}
	return builder.create(p)
}
