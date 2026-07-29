package armbulkactions_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armbulkactions"
)

// Generated from example definition: 2026-07-06-preview/ScheduledActions_TriggerManualOccurrence_MaximumSet_Gen.json
func ExampleScheduledActionsClient_BeginTriggerManualOccurrence() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbulkactions.NewClientFactory("CB26D7CB-3E27-465F-99C8-EAF7A4118245", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewScheduledActionsClient().BeginTriggerManualOccurrence(ctx, "rgcompute", "myScheduledAction", nil)
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
	// 			ScheduledTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-04-17T00:23:59.243Z"); return t}()),
	// 			ResultSummary: &armbulkactions.OccurrenceResultSummary{
	// 				Total: to.Ptr[int32](25),
	// 				Statuses: []*armbulkactions.ResourceResultSummary{
	// 					{
	// 						Code: to.Ptr("Succeeded"),
	// 						Count: to.Ptr[int32](4),
	// 						ErrorDetails: &armbulkactions.Error{
	// 							Code: to.Ptr("InternalServerError"),
	// 							Message: to.Ptr("An internal error occurred."),
	// 							Target: to.Ptr("virtualMachines"),
	// 							Details: []*armbulkactions.Error{
	// 							},
	// 							Innererror: &armbulkactions.InnerError{
	// 								Code: to.Ptr("InnerErrorCode"),
	// 							},
	// 						},
	// 					},
	// 				},
	// 			},
	// 			ProvisioningState: to.Ptr(armbulkactions.OccurrenceStateCreated),
	// 		},
	// 		ID: to.Ptr("/subscriptions/83C27AB3-A7B9-498B-B165-D9440661474F/resourceGroups/myRg/providers/Microsoft.Compute/scheduledActions/myScheduledAction"),
	// 		Name: to.Ptr("67b5bada-4772-43fc-8dbb-402476d98a45"),
	// 		Type: to.Ptr("Microsoft.Compute/scheduledActions/occurrences"),
	// 		SystemData: &armbulkactions.SystemData{
	// 			CreatedBy: to.Ptr("user@contoso.com"),
	// 			CreatedByType: to.Ptr(armbulkactions.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-04-17T00:23:55.288Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("user@contoso.com"),
	// 			LastModifiedByType: to.Ptr(armbulkactions.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-04-17T00:23:55.288Z"); return t}()),
	// 		},
	// 	},
	// }
}
