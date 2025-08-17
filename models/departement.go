package models

type Departement struct {
	ID              int64  `gorm:"primaryKey" json:"id"`
	DepartementName string `gorm:"type:varchar(255)" json:"departement_name"`
	MaxClockInTime  string `json:"max_clock_in_time"`
	MaxClockOutTime string `json:"max_clock_out_time"`
}

func (Departement) TableName() string {
	return "departement"
}
