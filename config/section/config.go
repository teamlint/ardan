package section

// Config 配置
type Config struct {
	App       *App      // 应用程序
	Server    *Server   // 服务器配置
	Databases Databases // 数据库引擎列表
	Caches    Caches    // 缓存引擎列表
}
