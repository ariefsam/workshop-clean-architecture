package main

import (
	"expense-tracker/controller"
	"expense-tracker/repository"
	"expense-tracker/service/activity"
	"expense-tracker/service/expense-tracker/spending"
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type spendingGorm struct {
	gorm.Model
	spending.Spending
}

func (s spendingGorm) TableName() string {
	return "spending"
}

type spendingActivityGorm struct {
	gorm.Model
	activity.Activity
}

func (s spendingActivityGorm) TableName() string {
	return "activity"
}

func main() {
	repoService, err := repository.NewSQLite("tbl.db")
	if err != nil {
		log.Println(err)
	}
	modelSpending := spendingGorm{}
	modelActivity := spendingActivityGorm{}
	repoService.AutoMigrate(&modelSpending)
	repoService.AutoMigrate(&modelActivity)

	// repoService := repository.NewMock()
	spendingService := spending.NewSpending(repoService)
	// spendingService.Create("id001", "Belanja bulanan", 1000000)
	controller.SpendingService = spendingService

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Post("/spending/create", controller.CreateSpendingFiber)

	app.Listen(":3000")

}
