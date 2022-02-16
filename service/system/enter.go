package system

type ServiceGroup struct {
	OperationRecordService
	UserService
	JwtService
	InitDBService
	CasbinService
	ApiService
}
