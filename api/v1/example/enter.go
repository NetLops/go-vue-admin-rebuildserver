package example

import "rebuildServer/service"

type ApiGroup struct {
	ExcelApi
}

var (
	excelService = service.ServiceGroupApp.ExampleServiceGroup.ExcelService
)
