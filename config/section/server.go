package section

// Server 服务器
type Server struct {
	HttpAddr     string // HTTP 服务地址
	GrpcAddr     string // GRPC 服务地址
	NatsAddr     string // NATS 服务地址
	DebugAddr    string // Debug 服务地址
	ReadTimeout  string // 读超时
	WriteTimeout string // 写超时
	IdleTimeout  string // 空闲超时
}
