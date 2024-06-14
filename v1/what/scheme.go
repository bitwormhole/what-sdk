package what

// Mode ... 表示连接模式
type Mode string

// ...
const (
	ModePktline   Mode = "pktline"
	ModeHypertext Mode = "hypertext"

	ModeSimple Mode = "simple"
	ModeBLOB   Mode = "blob"
	ModeFile   Mode = "file"
	ModeStream Mode = "stream"
)

func (s Mode) String() string {
	return string(s)
}
