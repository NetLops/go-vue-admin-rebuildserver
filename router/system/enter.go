package system

type RouterGroup struct {
	ApiRouter
	BaseRouter
	InitRouter
	JwtRouter
	UserRouter
}
