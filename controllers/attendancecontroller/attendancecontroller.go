package attendancecontroller

import (
	"net/http"
	"time"

	"github.com/DhikaNino/backend-fleetifyid-challenge/config"
	"github.com/DhikaNino/backend-fleetifyid-challenge/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func Index(c *fiber.Ctx) error {

	results := []models.AttendanceHistoryRow{}

	startDateStr := c.Query("start_date")
	endDateStr := c.Query("end_date")
	deptID := c.Query("departement_id")
	query := config.DB.Table("vw_attendance_history")

	if startDateStr != "" && endDateStr != "" {
		startDate, err1 := time.Parse("2006-01-02", startDateStr)
		endDate, err2 := time.Parse("2006-01-02", endDateStr)
		if err1 != nil || err2 != nil {
			return c.Status(400).JSON(fiber.Map{
				"message": "Format tanggal harus YYYY-MM-DD",
			})
		}
		endDate = endDate.Add(23*time.Hour + 59*time.Minute + 59*time.Second)
		query = query.Where("date_attendance BETWEEN ? AND ?", startDate, endDate)
	}

	if deptID != "" {
		query = query.Where("departement_id = ?", deptID)
	}

	if err := query.Scan(&results).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	for i, r := range results {
		loc, _ := time.LoadLocation("Asia/Jakarta")

		if r.AttendanceType == 1 {

			attendance := r.DateAttendance.In(loc)
			maxIn, _ := time.ParseInLocation("15:04:05", r.MaxClockInTime, loc)
			maxInFull := time.Date(attendance.Year(), attendance.Month(), attendance.Day(),
				maxIn.Hour(), maxIn.Minute(), maxIn.Second(), 0, loc)

			if r.DateAttendance.After(maxInFull) {
				results[i].StatusKetepatan = "Terlambat"
			} else {
				results[i].StatusKetepatan = "Tepat Waktu"
			}

		} else if r.AttendanceType == 2 {
			attendance := r.DateAttendance.In(loc)
			maxOut, _ := time.ParseInLocation("15:04:05", r.MaxClockOutTime, loc)
			maxOutFull := time.Date(attendance.Year(), attendance.Month(), attendance.Day(),
				maxOut.Hour(), maxOut.Minute(), maxOut.Second(), 0, loc)

			if r.DateAttendance.Before(maxOutFull) {
				results[i].StatusKetepatan = "Pulang lebih awal"
			} else {
				results[i].StatusKetepatan = "Pulang sesaui jadwal atau lebih dari jadwal yang ditentukan"
			}
		}
	}

	return c.JSON(fiber.Map{
		"data": results,
	})
}

func Create(c *fiber.Ctx) error {
	var req struct {
		EmployeeID string `json:"employee_id"`
	}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": err.Error()})
	}

	today := time.Now().Format("2006-01-02")
	var existing models.Attendance
	if err := config.DB.
		Where("employee_id = ? AND DATE(clock_in) = ?", req.EmployeeID, today).
		First(&existing).Error; err == nil {
		return c.Status(400).JSON(fiber.Map{"message": "Anda sudah absen masuk hari ini!"})
	}

	now := time.Now()

	attendance := models.Attendance{
		EmployeeID:   req.EmployeeID,
		AttendanceID: "ABSEN-" + uuid.New().String(),
		ClockIn:      &now,
		CreatedAt:    now,
		UpdatedAt:    now,
	}

	if err := config.DB.Create(&attendance).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"message": err.Error()})
	}

	history := models.AttendanceHistory{
		EmployeeID:     req.EmployeeID,
		AttendanceID:   attendance.AttendanceID,
		DateAttendance: *attendance.ClockIn,
		AttendanceType: 1,
		Description:    "Absen masuk",
		CreatedAt:      now,
		UpdatedAt:      now,
	}

	if err := config.DB.Create(&history).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"message": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Absensi masuk berhasil!"})
}

func Update(c *fiber.Ctx) error {
	type Request struct {
		EmployeeID string `json:"employee_id"`
	}
	var req Request
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	var attendance models.Attendance
	today := time.Now().Format("2006-01-02")
	if err := config.DB.Where("employee_id = ? AND DATE(clock_in) = ?", req.EmployeeID, today).First(&attendance).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{"message": "Tidak ada absen masuk hari ini!"})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	if attendance.ClockOut != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Anda sudah absen keluar hari ini!"})
	}

	now := time.Now()
	attendance.ClockOut = &now
	if err := config.DB.Save(&attendance).Error; err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	history := models.AttendanceHistory{
		EmployeeID:     req.EmployeeID,
		AttendanceID:   attendance.AttendanceID,
		DateAttendance: *attendance.ClockOut,
		AttendanceType: 2,
		Description:    "Absen pulang",
		CreatedAt:      now,
		UpdatedAt:      now,
	}

	if err := config.DB.Create(&history).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"message": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Absensi keluar berhasil!"})
}
