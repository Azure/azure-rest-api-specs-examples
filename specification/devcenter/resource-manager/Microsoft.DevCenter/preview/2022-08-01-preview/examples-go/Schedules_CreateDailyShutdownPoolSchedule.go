package armdevcenter_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devcenter/armdevcenter"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2022-08-01-preview/examples/Schedules_CreateDailyShutdownPoolSchedule.json
func ExampleSchedulesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdevcenter.NewSchedulesClient("{subscriptionId}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"rg1",
		"DevProject",
		"DevPool",
		"autoShutdown",
		armdevcenter.Schedule{
			Properties: &armdevcenter.ScheduleProperties{
				Type:      to.Ptr(armdevcenter.ScheduledTypeStopDevBox),
				Frequency: to.Ptr(armdevcenter.ScheduledFrequencyDaily),
				State:     to.Ptr(armdevcenter.EnableStatusEnabled),
				Time:      to.Ptr("17:30"),
				TimeZone:  to.Ptr("America/Los_Angeles"),
			},
		},
		&armdevcenter.SchedulesClientBeginCreateOrUpdateOptions{Top: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
