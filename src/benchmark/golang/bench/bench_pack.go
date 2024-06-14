package bench

import (
	"bytes"
	"strconv"
	"strings"
)

// PackType ...
type PackType string

// define PackType(s) ...
const (
	TypeCountRequest    PackType = "count.req"
	TypeCountResponse   PackType = "count.resp"
	TypePayloadRequest  PackType = "payload.req"
	TypePayloadResponse PackType = "payload.resp"
)

func (t PackType) String() string {
	return string(t)
}

////////////////////////////////////////////////////////////////////////////////

// Pack ...
type Pack struct {
	Type      PackType
	RequestX  int
	ResponseX int

	head map[string]string
	body []byte
}

// SetHeader ...
func (inst *Pack) SetHeader(name, value string) {
	t := inst.head
	if t == nil {
		t = make(map[string]string)
		inst.head = t
	}
	name = strings.ToLower(name)
	t[name] = value
}

// GetHeader ...
func (inst *Pack) GetHeader(name string) string {
	t := inst.head
	if t == nil {
		return ""
	}
	name = strings.ToLower(name)
	return t[name]
}

////////////////////////////////////////////////////////////////////////////////

const (
	fieldType = "type"
	fieldX1   = "x1"
	fieldX2   = "x2"
)

// Marshal ...
func Marshal(p *Pack) []byte {
	builder := &bytes.Buffer{}

	// prepare

	x1 := strconv.Itoa(p.RequestX)
	x2 := strconv.Itoa(p.ResponseX)

	p.SetHeader(fieldType, p.Type.String())
	p.SetHeader(fieldX1, x1)
	p.SetHeader(fieldX2, x2)

	// head
	table := p.head
	for k, v := range table {
		builder.WriteString(k)
		builder.WriteByte(':')
		builder.WriteString(v)
		builder.WriteByte('\n')
	}

	// body
	builder.WriteByte(0)
	builder.Write(p.body)

	return builder.Bytes()
}

// Unmarshal ...
func Unmarshal(raw []byte) (*Pack, error) {

	// const
	type AT rune
	const (
		atName  AT = 'n'
		atValue AT = 'v'
	)
	at := atName

	// prepare
	builder := &bytes.Buffer{}
	name := ""
	value := ""
	p := &Pack{}

	// scan
	for i, b := range raw {
		if b == 0 {
			// end of head
			p.body = raw[i+1:]
			break
		} else if (b == ':') && (at == atName) {
			at = atValue
			raw := builder.Bytes()
			name = string(raw)
			builder.Reset()
		} else if (b == '\n') && (at == atValue) {
			at = atName
			raw := builder.Bytes()
			value = string(raw)
			builder.Reset()
			k := strings.TrimSpace(name)
			v := strings.TrimSpace(value)
			p.SetHeader(k, v)
			name = ""
			value = ""
		} else {
			builder.WriteByte(b)
		}
	}

	// 提取特殊 headers

	pt := p.GetHeader(fieldType)
	x1str := p.GetHeader(fieldX1)
	x2str := p.GetHeader(fieldX2)
	x1, _ := strconv.Atoi(x1str)
	x2, _ := strconv.Atoi(x2str)

	p.Type = PackType(pt)
	p.RequestX = x1
	p.ResponseX = x2

	return p, nil
}
