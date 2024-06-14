package what

import "io"

// Connection ...
type Connection interface {
	io.Closer
}
