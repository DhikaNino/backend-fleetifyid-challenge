package models

import "time"

type AttendanceHistory struct {
	ID             int64     `gorm:"primaryKey" json:"id"`
	EmployeeID     string    `gorm:"type:varchar(50)" json:"employee_id"`
	AttendanceID   string    `gorm:"type:varchar(100)" json:"attendance_id"`
	DateAttendance time.Time `json:"date_attendance"`
	AttendanceType int8      `gorm:"type:tinyint" json:"attendance_type"`
	Description    string    `gorm:"type:text" json:"description"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func (AttendanceHistory) TableName() string {
	return "attendance_history"
}

type AttendanceHistoryRow struct {
	EmployeeID      string    `json:"employee_id"`
	EmployeeName    string    `json:"employee_name"`
	DepartementName string    `json:"departement_name"`
	DateAttendance  time.Time `json:"date_attendance"`
	AttendanceType  int8      `json:"attendance_type"`
	Description     string    `json:"description"`
	MaxClockInTime  string    `json:"max_clock_in_time"`
	MaxClockOutTime string    `json:"max_clock_out_time"`
	StatusKetepatan string    `json:"status_ketepatan"`
}
