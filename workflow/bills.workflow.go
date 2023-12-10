package workflow

import (
	"github.com/dsha256/feesapi/dto"
	"github.com/dsha256/feesapi/shared"
	"github.com/google/uuid"
	"github.com/mitchellh/mapstructure"
	"go.temporal.io/sdk/workflow"
)

func BillWorkflow(ctx workflow.Context, billUUID uuid.UUID) error {
	logger := workflow.GetLogger(ctx)

	var activities *Activities

	err := workflow.SetQueryHandler(ctx, "getBill", func(input []byte) (dto.Bill, error) {
		var bill dto.Bill
		params := GetBillByUUIDParams{UUID: billUUID}
		err := workflow.ExecuteActivity(ctx, activities.GetBillByUUID, params).Get(ctx, &bill)
		if err != nil {
			logger.Error("error getting bill %v", err)
			return dto.Bill{}, err
		}
		return bill, nil
	})
	if err != nil {
		logger.Info("SetQueryHandler failed.", "Error", err)
		return err
	}

	createBillChannel := workflow.GetSignalChannel(ctx, shared.SignalChannels.OpenBillChannel)
	checkOut := false

	for {
		selector := workflow.NewSelector(ctx)
		selector.AddReceive(createBillChannel, func(c workflow.ReceiveChannel, _ bool) {
			var signal any
			c.Receive(ctx, &signal)

			var message shared.OpenBillSignal
			if err := mapstructure.Decode(signal, &message); err != nil {
				logger.Error("Invalid signal type %v", err)
				return
			}

			err = workflow.ExecuteActivity(ctx, activities.OpenBill, message.UUID).Get(ctx, nil)
			if err != nil {
				logger.Error("Can not create bill %v", err)
				return
			}
		})

		selector.Select(ctx)

		if checkOut {
			break
		}
	}

	return nil
}
