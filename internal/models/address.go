package models

type Address struct {
	ID           uint   `gorm:"primary_key"`
	ZipCode      string `gorm:"size:9" json:"zip_code"`
	Street       string `gorm:"size:50" json:"street"`
	Number       string `gorm:"size:10" json:"number"`
	Complement   string `gorm:"size:30" json:"complement"`
	Neighborhood string `gorm:"size:50" json:"neighborhood"`
	City         string `gorm:"size:50" json:"city"`
	State        string `gorm:"size:50" json:"state"`
	PersonID     uint
}
