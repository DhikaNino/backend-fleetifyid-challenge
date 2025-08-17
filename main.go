package main

import (
	"log"
	"os"

	"github.com/DhikaNino/backend-fleetifyid-challenge/config"
	"github.com/DhikaNino/backend-fleetifyid-challenge/controllers/attendancecontroller"
	"github.com/DhikaNino/backend-fleetifyid-challenge/controllers/departementcontroller"
	"github.com/DhikaNino/backend-fleetifyid-challenge/controllers/employeecontroller"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error file .env tidak ketemu!")
	}

	config.ConnDatabase()

	app := fiber.New()
	app.Use(cors.New())

	// Ini untuk endpoint utama API
	api := app.Group("/api")

	// Ini untuk endpoint karyawan
	employee := api.Group("/employee")
	employee.Get("/", employeecontroller.Index)
	employee.Get("/:employee_id", employeecontroller.Show)
	employee.Post("/", employeecontroller.Create)
	employee.Put("/:employee_id", employeecontroller.Update)
	employee.Delete("/:employee_id", employeecontroller.Delete)

	// Ini untuk endpoint departemen
	departement := api.Group("/departement")
	departement.Get("/", departementcontroller.Index)
	departement.Get("/:id", departementcontroller.Show)
	departement.Post("/", departementcontroller.Create)
	departement.Put("/:id", departementcontroller.Update)
	departement.Delete("/:id", departementcontroller.Delete)

	// Ini untuk endpoint absensi dan log
	attendance := api.Group("/attendance")
	attendance.Get("/", attendancecontroller.Index)
	attendance.Post("/in", attendancecontroller.Create)
	attendance.Put("/out", attendancecontroller.Update)

	port := os.Getenv("API_PORT")
	if port == "" {
		port = "3000" // port default jika tidak ada port di .env
	}

	app.Listen(":" + port)

}
