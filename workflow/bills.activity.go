package workflow

import (
	"context"

	"github.com/dsha256/feesapi/dto"
	"github.com/dsha256/feesapi/entity"
	"github.com/google/uuid"
)

type Activities struct {
	Entity *entity.Client
}

type OpenBillParams struct {
	UUID uuid.UUID
}

func (a Activities) OpenBill(ctx context.Context, params OpenBillParams) (err error) {
	_, err = a.Entity.Bill.Create().SetID(params.UUID).Save(ctx)

	return
}

type GetBillByUUIDParams struct {
	UUID uuid.UUID
}

func (a Activities) GetBillByUUID(ctx context.Context, params GetBillByUUIDParams) (dtoBill dto.Bill, err error) {
	bill, err := a.Entity.Bill.Create().SetID(params.UUID).Save(ctx)
	if err != nil {
		return dto.Bill{}, err
	}
	dtoBill.FromEntityBill(bill)

	return
}
