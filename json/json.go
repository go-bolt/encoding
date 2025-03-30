package json

import (
	"encoding/json"

	"github.com/go-bolt/encoding"
)

const Name = "json"

var (
	// marshal 序列化
	marshal = json.Marshal
	// unmarshal 反序列化
	unmarshal = json.Unmarshal
	// Codec 编码器
	Codec = encoding.NewCodec(Name, marshal, unmarshal)
)

// Use 设置自定义序列化/反序列化函数
func Use(marshalFn encoding.Marshal, unmarshalFn encoding.Unmarshal) {
	marshal = marshalFn
	unmarshal = unmarshalFn
}
