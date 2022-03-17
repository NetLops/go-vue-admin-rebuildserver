package system

type RouterGroup struct {
	ApiRouter
	BaseRouter
	InitRouter
	JwtRouter
	UserRouter
	MenuRouter
	SysRouter
	CasbinRouter
	AutoCodeRouter
	AuthorityRouter
	DictionaryRouter
	AutoCodeHistoryRouter
	OperationRecordRouter
	DictionaryDetailRouter
}
