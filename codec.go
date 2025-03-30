package encoding

type Codec interface {
	// Name 名称
	Name() string
	// Aliases 后缀别名
	Aliases() []string
	// Encode 编码
	Encode(v any) ([]byte, error)
	// Decode 解码
	Decode(data []byte, v any) error
}

// defaultCodec 编码器
type defaultCodec struct {
	name    string
	aliases []string
	encode  Marshal
	decode  Unmarshal
}

// NewCodec 返回一个新的编解码器
func NewCodec(name string, enc Marshal, dec Unmarshal, aliases ...string) Codec {
	return &defaultCodec{
		name:    name,
		encode:  enc,
		decode:  dec,
		aliases: aliases,
	}
}

// Name 名称
func (c *defaultCodec) Name() string {
	return c.name
}

// Encode 编码
func (c *defaultCodec) Encode(v any) ([]byte, error) {
	return c.encode(v)
}

// Decode 解码
func (c *defaultCodec) Decode(data []byte, v any) error {
	return c.decode(data, v)
}

// Aliases 别名
func (c *defaultCodec) Aliases() []string {
	return c.aliases
}
