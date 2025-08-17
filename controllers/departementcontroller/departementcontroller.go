package departementcontroller

import (
	"math"
	"net/http"
	"strconv"

	"github.com/DhikaNino/backend-fleetifyid-challenge/config"
	"github.com/DhikaNino/backend-fleetifyid-challenge/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Index(c *fiber.Ctx) error {
	var departements []models.Departement
	var total int64

	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	search := c.Query("search", "")

	offset := (page - 1) * limit

	query := config.DB.Model(&models.Departement{})

	if search != "" {
		query = query.Where("departement_name LIKE ?", "%"+search+"%")
	}

	query.Count(&total)

	query.Offset(offset).Limit(limit).Find(&departements)

	lastPage := int(math.Ceil(float64(total) / float64(limit)))

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": departements,
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
	id := c.Params("id")
	var departement models.Departement
	if err := config.DB.First(&departement, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"message": "Data tidak ada!",
			})
		}

		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Data tidak ada!",
		})
	}
	return c.JSON(departement)
}

func Create(c *fiber.Ctx) error {
	var departement models.Departement
	if err := c.BodyParser(&departement); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := config.DB.Create(&departement).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	var createdDepartement models.Departement
	if err := config.DB.First(&createdDepartement, departement.ID).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Berhasil membuat departement baru!",
		"data":    createdDepartement,
	})
}

func Update(c *fiber.Ctx) error {
	id := c.Params("id")

	var departement models.Departement
	if err := c.BodyParser(&departement); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if config.DB.Where("id = ?", id).Updates(&departement).RowsAffected == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Data tidak ditemukan!",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Data berhasil diubah!",
	})
}
func Delete(c *fiber.Ctx) error {

	id := c.Params("id")

	var departement models.Departement
	if config.DB.Delete(&departement, id).RowsAffected == 0 {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"message": "Tidak dapat menghapus data!",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Data berhasil dihapus!",
	})
}
