package models

type GoSite struct {
	SiteID     string            `json:"site_id" gorm:"type:CHAR(32); primaryKey"`
	Name       string            `json:"name" gorm:"type:VARCHAR(255); not null"`
	Active     *bool             `json:"active" gorm:"default:true; not null"`
	Attributes []GoSiteAttribute `json:"attributes" gorm:"foreignKey:SiteID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
