package controller

import (
	"encoding/json"
	"expense-tracker/idgenerator"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func CreateSpendingFiber(c *fiber.Ctx) (err error) {
	payload := struct {
		Name   string  `json:"name"`
		Amount float64 `json:"amount"`
	}{}
	c.BodyParser(&payload)

	id := idgenerator.Generate()

	err = SpendingService.Create(id, payload.Name, payload.Amount)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	dataResponse := map[string]any{
		"id":     id,
		"name":   payload.Name,
		"amount": payload.Amount,
	}

	return c.JSON(dataResponse)
}

func CreateSpendingHandler(w http.ResponseWriter, r *http.Request) {
	payload := struct {
		Name   string  `json:"name"`
		Amount float64 `json:"amount"`
	}{}

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	id := idgenerator.Generate()

	err = SpendingService.Create(id, payload.Name, payload.Amount)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	dataResponse := map[string]any{
		"id":     id,
		"name":   payload.Name,
		"amount": payload.Amount,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(dataResponse)

}
