package section

// Databases 数据库引擎列表
type Databases map[string]*Database

// Database 数据库配置
type Database struct {
	DriverName      string // 数据库驱动名称
	ConnString      string // 数据库连接字符串
	ConnMaxLifetime string // 连接生命周期(分钟)
	MaxOpenConns    int    // 最大连接数
	MaxIdleConns    int    // 最大空闲连接数
	Log             bool   // 数据库日志
}
