package yaml

import (
	"gopkg.in/yaml.v3"

	"github.com/go-bolt/encoding"
)

// Name 名称
const Name = "yaml"

var Codec = encoding.NewCodec(Name, yaml.Marshal, yaml.Unmarshal, "yml")

// Marshal 编码
func Marshal(v any) ([]byte, error) {
	return yaml.Marshal(v)
}

// Unmarshal 解码
func Unmarshal(data []byte, v any) error {
	return yaml.Unmarshal(data, v)
}
