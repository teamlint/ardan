package section

// App 应用程序
type App struct {
	Title      string                 // 应用程序标题
	Copyright  string                 // 版权信息
	Debug      bool                   // 是否开启调试
	TimeFormat string                 // 时间格式
	Charset    string                 // 字符集
	Settings   map[string]interface{} // 应用配置
}
