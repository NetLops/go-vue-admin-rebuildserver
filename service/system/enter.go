package system

type ServiceGroup struct {
	OperationRecordService
	UserService
	JwtService
	InitDBService
	CasbinService
	ApiService
	MenuService
	BaseMenuService
	SystemConfigService
	AutoCodeService
	AuthorityService
	DictionaryService
	AutoCodeHistoryService
	DictionaryDetailService
}
