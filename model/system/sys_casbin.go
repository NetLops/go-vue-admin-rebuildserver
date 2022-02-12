package system

type CasbinModel struct {
	Ptype       string `json:"ptype" gorm:"column:ptype"`
	AuthorityId string `json:"rolename" gorm:"colummn:v0"`
	Path        string `json:"path" gorm:"colummn:v1"`
	Method      string `json:"method" gorm:"colummn:v2"`
}
