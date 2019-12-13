package route

import(
    "go.uber.org/dig"
    "github.com/gin-gonic/gin"
)

// RouteInfo 路由信息 
type RouteInfo struct {
	Method      string
	Path        string
	// Handler     string
	Handlers gin.HandlersChain
}

// Route 路由
type Route struct{
    dig.Out
    // gin.RouteInfo `group:"route"`
    RouteInfo `group:"route"`
}

// 路由集合
type Routes []Route

// New 创建一个路由
func New(method string,path string,handlers ...gin.HandlerFunc) Route{
    routeInfo:=RouteInfo{Method:method,Path:path,Handlers:handlers}
    return Route{RouteInfo:routeInfo}
}

func Get(path string, handlers ...gin.HandlerFunc) Route {
    get:=RouteInfo{Method:"GET",Path:path,Handlers:handlers}
    return Route{RouteInfo:get}
}
func Post(path string, handlers ...gin.HandlerFunc) Route {
    post:=RouteInfo{Method:"POST",Path:path,Handlers:handlers}
    return Route{RouteInfo:post}
}
func Put(path string, handlers ...gin.HandlerFunc) Route {
    put:=RouteInfo{Method:"PUT",Path:path,Handlers:handlers}
    return Route{RouteInfo:put}
}
func Patch(path string, handlers ...gin.HandlerFunc) Route {
    patch:=RouteInfo{Method:"PATCH",Path:path,Handlers:handlers}
    return Route{RouteInfo:patch}
}
func Head(path string, handlers ...gin.HandlerFunc) Route {
    head:=RouteInfo{Method:"HEAD",Path:path,Handlers:handlers}
    return Route{RouteInfo:head}
}
func Options(path string, handlers ...gin.HandlerFunc) Route {
    options:=RouteInfo{Method:"OPTIONS",Path:path,Handlers:handlers}
    return Route{RouteInfo:options}
}
func Delete(path string, handlers ...gin.HandlerFunc) Route {
    del:=RouteInfo{Method:"DELETE",Path:path,Handlers:handlers}
    return Route{RouteInfo:del}
}
func Connect(path string, handlers ...gin.HandlerFunc) Route {
    connect:=RouteInfo{Method:"CONNECT",Path:path,Handlers:handlers}
    return Route{RouteInfo:connect}
}
func Trace(path string, handlers ...gin.HandlerFunc) Route {
    trace:=RouteInfo{Method:"TRACE",Path:path,Handlers:handlers}
    return Route{RouteInfo:trace}
}
func GetAndPost(path string, handlers ...gin.HandlerFunc) Route {
    gp:=RouteInfo{Method:"GET_POST",Path:path,Handlers:handlers}
    return Route{RouteInfo:gp}
}
func Any(path string, handlers ...gin.HandlerFunc) Route {
    any:=RouteInfo{Method:"ANY",Path:path,Handlers:handlers}
    return Route{RouteInfo:any}
}
