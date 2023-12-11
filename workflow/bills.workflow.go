package workflow

import (
	"time"

	"github.com/dsha256/feesapi/dto"
	"github.com/dsha256/feesapi/shared"
	"github.com/mitchellh/mapstructure"
	"go.temporal.io/sdk/workflow"
)

var commonActivityOptions = workflow.ActivityOptions{
	StartToCloseTimeout:    time.Second * 15,
	ScheduleToCloseTimeout: time.Second * 15,
	ScheduleToStartTimeout: time.Second * 5,
}

type BillWorkflowExecutionParams struct {
	UUID string
}

func BillWorkflow(ctx workflow.Context, params BillWorkflowExecutionParams) error {
	logger := workflow.GetLogger(ctx)

	currentBillState := dto.Bill{}

	var activities *Activities

	err := workflow.SetQueryHandler(ctx, "getBill", func(input []byte) (dto.Bill, error) {
		return currentBillState, nil
	})
	if err != nil {
		logger.Info("SetQueryHandler failed.", "Error", err)
		return err
	}

	openBillChannel := workflow.GetSignalChannel(ctx, shared.SignalChannels.OpenBillChannel)
	checkoutBillChannel := workflow.GetSignalChannel(ctx, shared.SignalChannels.CheckoutBillChannel)
	checkedOut := false

	var receiverErr error
	for {
		selector := workflow.NewSelector(ctx)
		selector.AddReceive(openBillChannel, func(c workflow.ReceiveChannel, _ bool) {
			var signal any
			c.Receive(ctx, &signal)
			var message shared.OpenBillSignal
			if err := mapstructure.Decode(signal, &message); err != nil {
				logger.Error("invalid signal type", err.Error())
				receiverErr = err
				return
			}

			ctx = workflow.WithActivityOptions(ctx, commonActivityOptions)
			params := OpenBillParams{UUID: params.UUID}
			var bill dto.Bill
			err := workflow.ExecuteActivity(ctx, activities.OpenBill, params).Get(ctx, &bill)
			if err != nil {
				logger.Error("can not create bill %v", err)
				receiverErr = err
				return
			}

			currentBillState = bill
		})

		selector.AddReceive(checkoutBillChannel, func(c workflow.ReceiveChannel, _ bool) {
			var signal interface{}
			c.Receive(ctx, &signal)
			var message shared.CheckoutSignal
			err := mapstructure.Decode(signal, &message)
			if err != nil {
				logger.Error("invalid signal type %v", err)
				receiverErr = err
				return
			}

			checkedOut = true
		})

		selector.Select(ctx)

		if checkedOut {
			break
		}
	}

	return receiverErr
}
