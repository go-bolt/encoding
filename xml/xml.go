package xml

import (
	"encoding/xml"

	"github.com/go-bolt/encoding"
)

const Name = "xml"

var Codec = encoding.NewCodec(Name, Marshal, Unmarshal)

// Marshal 编码
func Marshal(v any) ([]byte, error) {
	return xml.Marshal(v)
}

// Unmarshal 解码
func Unmarshal(data []byte, v any) error {
	return xml.Unmarshal(data, v)
}
