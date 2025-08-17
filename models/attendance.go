package models

import "time"

type Attendance struct {
	ID           int64      `gorm:"primaryKey" json:"id"`
	EmployeeID   string     `gorm:"type:varchar(50);" json:"employee_id"`
	AttendanceID string     `gorm:"type:varchar(100);unique;" json:"attendance_id"`
	ClockIn      *time.Time `json:"clock_in"`
	ClockOut     *time.Time `json:"clock_out"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
}

func (Attendance) TableName() string {
	return "attendance"
}
