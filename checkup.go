package checkup

import (
	"context"
	"errors"
	"github.com/sandwich-go/checkup/protocol/gen/golang/internal_command"
)

// Packet 传输的包体
type Packet struct {
	Uri string `json:"uri"`
	Raw []byte `json:"raw"`
}

var (
	// ErrUnknownPacket Decode 时候，若不合法，则抛出该错误
	ErrUnknownPacket = errors.New("unknown checkup packet")
	ErrHandleRequest = errors.New("handle checkup request error")
)

// Is 是否为有效的消息
type Is = bool

// Codec 序列化/反序列化
type Codec interface {
	// Marshal 序列化
	Marshal(v any) ([]byte, error)
	// Unmarshal 反序列化
	Unmarshal(data []byte, v any) error
}

// Handler 处理器
type Handler interface {
	// IsRequestPath 是否是请求的 path
	IsRequestPath(string) Is
	// RequestBytes  Checkup 请求的字节数组
	RequestBytes() []byte
	// HandleIfRequestBytes 如果是 Checkup 请求的字节数组，则处理
	HandleIfRequestBytes(ctx context.Context, in []byte) ([]byte, Is)
	// HandleResponseBytes 处理 Checkup 的响应字节数组
	HandleResponseBytes(in []byte) (*internal_command.CmdCheckup, error)
}
