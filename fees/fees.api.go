package fees

import (
	"context"
	"log"

	"encore.dev/rlog"
	"github.com/dsha256/feesapi/currency"
	"github.com/dsha256/feesapi/dto"
	"github.com/dsha256/feesapi/shared"
	"github.com/dsha256/feesapi/workflow"
	"github.com/google/uuid"
	"go.temporal.io/sdk/client"
)

//encore:api public method=POST path=/fees/api/v1/bills
func (f *Fees) OpenBill(ctx context.Context) (*dto.Bill, error) {
	newBillUUID := uuid.New().String()
	log.Println(">>> UUID:", newBillUUID)
	workflowID := "BILL-" + newBillUUID

	options := client.StartWorkflowOptions{
		ID:        workflowID,
		TaskQueue: ServiceTaskQueue,
	}

	state := workflow.BillWorkflowExecutionParams{
		UUID: newBillUUID,
	}
	we, err := f.client.ExecuteWorkflow(context.Background(), options, workflow.BillWorkflow, state)
	if err != nil {
		return &dto.Bill{}, err
	}
	rlog.Info("started workflow", "id", we.GetID(), "run_id", we.GetRunID())

	key := TemporalWorkflowsIDsCacheKey{UUID: newBillUUID}
	err = f.workflowIDCache.Set(ctx, key, workflowID)
	if err != nil {
		return &dto.Bill{}, err
	}

	args := shared.OpenBillSignal{UUID: newBillUUID}
	err = f.client.SignalWorkflow(context.Background(),
		workflowID,
		"",
		shared.SignalChannels.OpenBillChannel,
		args,
	)
	if err != nil {
		return nil, err
	}

	var bill *dto.Bill
	response, err := f.client.QueryWorkflow(context.Background(), workflowID, "", "getBill")
	if err != nil {
		return bill, err
	}

	err = response.Get(&bill)
	if err != nil {
		return nil, err
	}

	//time.Sleep(time.Second * 2)

	//bill, err := f.entity.Bill.Get(ctx, newBillUUID)
	//if err != nil {
	//	return &dto.Bill{}, err
	//}

	return bill, nil
}

//encore:api public method=POST path=/fees/api/v1/bills/:billUUID/checkout
func (f *Fees) Checkout(ctx context.Context, billUUID string) (*dto.Checkout, error) {
	UUID, err := uuid.Parse(billUUID)
	if err != nil {
		return &dto.Checkout{}, err
	}

	//workflowID, _ := workflows["BILL-"+UUID.String()]

	err = f.client.SignalWorkflow(context.Background(), "BILL-"+UUID.String(), "", shared.SignalChannels.CheckoutBillChannel, shared.CheckoutSignal{
		PayingCurrency: currency.USD,
	})
	if err != nil {
		return &dto.Checkout{}, err
	}

	return &dto.Checkout{Total: 10}, nil
}
