package employeecontroller

import (
	"fmt"
	"math"
	"strconv"

	"github.com/DhikaNino/backend-fleetifyid-challenge/config"

	"github.com/DhikaNino/backend-fleetifyid-challenge/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Index(c *fiber.Ctx) error {
	var employee []models.Employee
	var total int64

	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	search := c.Query("search", "")

	offset := (page - 1) * limit

	query := config.DB.Preload("Departement").Model(&models.Employee{})

	if search != "" {
		query = query.Where("name LIKE ?", "%"+search+"%")
	}

	query.Count(&total)

	query.Offset(offset).Limit(limit).Find(&employee)

	lastPage := int(math.Ceil(float64(total) / float64(limit)))

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": employee,
		"meta": fiber.Map{
			"total":     total,
			"page":      page,
			"last_page": lastPage,
			"limit":     limit,
			"search":    search,
		},
	})
}

func Show(c *fiber.Ctx) error {
	employeeID := c.Params("employee_id")

	var employee models.Employee
	if err := config.DB.Preload("Departement").Where("employee_id = ?", employeeID).First(&employee).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Data tidak ada!",
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Terjadi kesalahan server",
		})
	}

	return c.JSON(fiber.Map{
		"data": employee,
	})
}

func Create(c *fiber.Ctx) error {
	var employee models.Employee
	if err := c.BodyParser(&employee); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	var idTerakhir models.Employee
	if err := config.DB.Order("id desc").First(&idTerakhir).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
			})
		}
	}

	nomorUrut := 1
	if idTerakhir.ID != 0 {
		nomorUrut = int(idTerakhir.ID) + 1
	}

	employee.EmployeeID = fmt.Sprintf("KRY-%05d", nomorUrut)

	if err := config.DB.Create(&employee).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	var createdEmployee models.Employee
	if err := config.DB.Preload("Departement").First(&createdEmployee, employee.ID).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"data": createdEmployee,
	})
}

func Update(c *fiber.Ctx) error {
	employeeID := c.Params("employee_id")

	var input models.Employee
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	var employee models.Employee
	if err := config.DB.Where("employee_id = ?", employeeID).First(&employee).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Data tidak ditemukan!",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Server Error!",
		})
	}

	if err := config.DB.Model(&employee).
		Omit("employee_id").
		Updates(input).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal mengubah data!",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Data berhasil diubah!",
		"data":    employee,
	})
}

func Delete(c *fiber.Ctx) error {
	employeeID := c.Params("employee_id")

	var employee models.Employee
	result := config.DB.Where("employee_id = ?", employeeID).Delete(&employee)
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Tidak dapat menghapus data!",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Data berhasil dihapus!",
	})
}
