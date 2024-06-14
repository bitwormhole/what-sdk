package bench

import (
	"encoding/json"

	"github.com/starter-go/base/lang"
)

// Counter ... 收发数据计数器
type Counter struct {
	RxBytes int64     `json:"rx_bytes"`
	RxPacks int64     `json:"rx_packs"`
	TxBytes int64     `json:"tx_bytes"`
	TxPacks int64     `json:"tx_packs"`
	Startup lang.Time `json:"startup"`
	Now     lang.Time `json:"now"`
}

// Bytes ... 把 Counter 编码为 JSON 形式
func (inst *Counter) Bytes() []byte {
	j, err := json.Marshal(inst)
	if err != nil {
		str := "{}"
		return []byte(str)
	}
	return j
}

// ParseCounter ... 解析 JSON 形式的 Counter
func ParseCounter(b []byte) (*Counter, error) {
	cnt := &Counter{}
	err := json.Unmarshal(b, cnt)
	if err != nil {
		return nil, err
	}
	return cnt, nil
}
