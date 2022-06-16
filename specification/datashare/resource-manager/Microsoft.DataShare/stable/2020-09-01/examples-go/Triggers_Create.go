package armdatashare_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datashare/armdatashare"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/datashare/resource-manager/Microsoft.DataShare/stable/2020-09-01/examples/Triggers_Create.json
func ExampleTriggersClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdatashare.NewTriggersClient("433a8dfd-e5d5-4e77-ad86-90acdc75eb1a", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreate(ctx,
		"SampleResourceGroup",
		"Account1",
		"ShareSubscription1",
		"Trigger1",
		&armdatashare.ScheduledTrigger{
			Kind: to.Ptr(armdatashare.TriggerKindScheduleBased),
			Properties: &armdatashare.ScheduledTriggerProperties{
				RecurrenceInterval:  to.Ptr(armdatashare.RecurrenceIntervalDay),
				SynchronizationMode: to.Ptr(armdatashare.SynchronizationModeIncremental),
				SynchronizationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-11-14T04:47:52.9614956Z"); return t }()),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
