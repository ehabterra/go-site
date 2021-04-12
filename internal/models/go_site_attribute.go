package models

type GoSiteAttribute struct {
	SiteID   string  `json:"site_id" gorm:"foreignKey:Site; type:CHAR(32); primaryKey"`
	Key      string  `json:"key" gorm:"type:VARCHAR(255); primaryKey; not null"`
	ValueStr *string `json:"value_str" gorm:"type:VARCHAR(255)"`
	ValueInt *int    `json:"value_int" gorm:"type:int"`
}
