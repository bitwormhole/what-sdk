package what

import "time"

// ConnectionParams ...
type ConnectionParams struct {
	Service string
	Mode    Mode
	URL     string // a station node URL
	DoGet   bool   // for listener to get connection
	Timeout time.Duration
}

// ListenerParams ...
type ListenerParams struct {
	Service string
	Mode    Mode
}
