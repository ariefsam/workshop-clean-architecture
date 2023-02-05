package controller

var SpendingService interface {
	Create(id, name string, amount float64) (err error)
}
