package toml

import (
	"bytes"

	"github.com/BurntSushi/toml"
	"github.com/go-bolt/encoding"
)

// Name 名称
const Name = "toml"

// Codec 编解码器
var Codec = encoding.NewCodec(Name, Marshal, Unmarshal)

// Marshal 编码
func Marshal(v any) ([]byte, error) {
	buffer := &bytes.Buffer{}
	return buffer.Bytes(), toml.NewEncoder(buffer).Encode(v)
}

// Unmarshal 解码
func Unmarshal(data []byte, v any) error {
	return toml.Unmarshal(data, v)
}
