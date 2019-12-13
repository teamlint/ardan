package section

// Caches 缓存引擎列表
type Caches map[string]*Cache

// Cache 缓存
type Cache struct {
	Proto        string // 协议
	Addr         string // 地址
	MaxOpenConns int    // 最大连接数
	MaxIdleConns int    // 最大空闲连接数
	ReadTimeout  string // 读超时
	WriteTimeout string // 写超时

}
