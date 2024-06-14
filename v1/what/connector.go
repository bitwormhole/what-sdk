package what

// Connector ...
type Connector interface {
	Connect(c *Context, p *ConnectionParams) (Connection, error)

	Registration() *ConnectorRegistration
}

// ConnectorRegistration ...
type ConnectorRegistration struct {
	Enabled   bool
	Priority  int
	Name      string
	Mode      Mode
	Connector Connector
	Support   func(*ConnectionParams) bool
}
