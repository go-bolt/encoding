package encoding

type (
	// Marshal 序列化
	Marshal func(v any) ([]byte, error)
	// Unmarshal 反序列化
	Unmarshal func(data []byte, v any) error
)
