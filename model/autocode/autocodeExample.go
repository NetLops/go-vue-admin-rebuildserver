package autocode

import "rebuildServer/global"

type AutoCodeExample struct {
	global.GVA_MODEL
	AutoCodeExampleField string `json:"autoCodeExampleField" form:"autoCodeExampleField" gorm:"column:auto_code_example_field;comment:仅作示例条目无实际作用"` // 展示值
}
