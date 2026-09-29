package armbulkactions_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armbulkactions"
)

// Generated from example definition: 2026-10-06-preview/ScheduledActions_TriggerManualOccurrence_BasicSuccess.json
func ExampleScheduledActionsClient_BeginTriggerManualOccurrence() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbulkactions.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewScheduledActionsClient().BeginTriggerManualOccurrence(ctx, "example-rg", "weekday-start", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armbulkactions.ScheduledActionsClientTriggerManualOccurrenceResponse{
	// 	Occurrence: armbulkactions.Occurrence{
	// 		Properties: &armbulkactions.OccurrenceProperties{
	// 			ScheduledTime: to.Ptr(time.Date(2026, time.September, 1, 2, 0, 0, 0, time.UTC)),
	// 			ResultSummary: &armbulkactions.OccurrenceResultSummary{
	// 				Total: to.Ptr[int32](2),
	// 				Statuses: []*armbulkactions.ResourceResultSummary{
	// 				},
	// 			},
	// 			ProvisioningState: to.Ptr(armbulkactions.OccurrenceStateScheduled),
	// 		},
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-rg/providers/Microsoft.Compute/scheduledActions/weekday-start/occurrences/67b5bada-4772-43fc-8dbb-402476d98a45"),
	// 		Name: to.Ptr("67b5bada-4772-43fc-8dbb-402476d98a45"),
	// 		Type: to.Ptr("Microsoft.Compute/scheduledActions/occurrences"),
	// 	},
	// }
}
