package dto

import (
	"time"

	"github.com/dsha256/feesapi/currency"
	"github.com/dsha256/feesapi/entity"
)

type Bill struct {
	ID        string                     `json:"id"`
	Currency  currency.SupportedCurrency `json:"currency"`
	Total     int64                      `json:"total"`
	IsOpen    bool                       `json:"is_open"`
	Items     []*LineItem                `json:"items"`
	ClosedAt  time.Time                  `json:"closed_at"`
	CreatedAt time.Time                  `json:"created_at"`
}

func NewBillFromEntityBill(entityBill *entity.Bill) Bill {
	return Bill{
		ID:        entityBill.ID,
		Currency:  currency.SupportedCurrency(entityBill.Currency.String()),
		Total:     entityBill.Total,
		IsOpen:    entityBill.IsOpen,
		ClosedAt:  entityBill.ClosedAt,
		CreatedAt: entityBill.CreatedAt,
	}
}

type LineItem struct {
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	Price    int64     `json:"price"`
	Quantity int64     `json:"quantity"`
	AddedAt  time.Time `json:"added_at"`
}

type Checkout struct {
	Total int `json:"total"`
}
