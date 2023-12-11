package workflow

import (
	"context"

	"github.com/dsha256/feesapi/dto"
	"github.com/dsha256/feesapi/entity"
)

type Activities struct {
	Entity *entity.Client
}

type OpenBillParams struct {
	UUID string
}

func (a *Activities) OpenBill(ctx context.Context, params OpenBillParams) (dto.Bill, error) {
	bill, err := a.Entity.Bill.Create().SetID(params.UUID).Save(ctx)
	if err != nil {
		return dto.Bill{}, err
	}

	return dto.NewBillFromEntityBill(bill), nil
}
