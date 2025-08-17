package models

import (
	"time"
)

type Employee struct {
	ID            int64     `gorm:"primaryKey" json:"id"`
	EmployeeID    string    `gorm:"type:varchar(50);unique;not null" json:"employee_id"`
	DepartementID int64     `gorm:"not null" json:"departement_id"`
	Name          string    `gorm:"type:varchar(255);not null" json:"name"`
	Address       string    `gorm:"type:text" json:"address"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`

	Departement Departement `gorm:"foreignKey:DepartementID;references:ID" json:"departement"`
}

func (Employee) TableName() string {
	return "employee"
}
