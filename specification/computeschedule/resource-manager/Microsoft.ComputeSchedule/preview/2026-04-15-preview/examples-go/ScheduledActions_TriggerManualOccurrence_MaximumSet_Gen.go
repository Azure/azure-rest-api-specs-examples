package armcomputeschedule_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/computeschedule/armcomputeschedule/v2"
)

// Generated from example definition: 2026-04-15-preview/ScheduledActions_TriggerManualOccurrence_MaximumSet_Gen.json
func ExampleScheduledActionsClient_TriggerManualOccurrence() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcomputeschedule.NewClientFactory("B4D97119-C802-4ADE-AF16-D9F92C9962C4", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewScheduledActionsClient().TriggerManualOccurrence(ctx, "rgcomputeschedule", "my-scheduled-action", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcomputeschedule.ScheduledActionsClientTriggerManualOccurrenceResponse{
	// 	Occurrence: armcomputeschedule.Occurrence{
	// 		Properties: &armcomputeschedule.OccurrenceProperties{
	// 			ScheduledTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-04-01T18:53:47.024Z"); return t}()),
	// 			ResultSummary: &armcomputeschedule.OccurrenceResultSummary{
	// 				Total: to.Ptr[int32](18),
	// 				Statuses: []*armcomputeschedule.ResourceResultSummary{
	// 					{
	// 						Code: to.Ptr("Success"),
	// 						Count: to.Ptr[int32](3),
	// 						ErrorDetails: &armcomputeschedule.Error{
	// 							Code: to.Ptr("ResourceNotFound"),
	// 							Message: to.Ptr("The specified resource was not found."),
	// 							Target: to.Ptr("virtualMachines/myVm"),
	// 							Details: []*armcomputeschedule.Error{
	// 							},
	// 							Innererror: &armcomputeschedule.InnerError{
	// 								Code: to.Ptr("ResourceNotFoundError"),
	// 							},
	// 						},
	// 					},
	// 				},
	// 			},
	// 			ProvisioningState: to.Ptr(armcomputeschedule.OccurrenceStateCreated),
	// 		},
	// 		ID: to.Ptr("/subscriptions/B4D97119-C802-4ADE-AF16-D9F92C9962C4/resourceGroups/rgcomputeschedule/providers/Microsoft.ComputeSchedule/scheduledActions/my-scheduled-action/occurrences/12345678-1234-1234-1234-123456789012"),
	// 		Name: to.Ptr("12345678-1234-1234-1234-123456789012"),
	// 		Type: to.Ptr("Microsoft.ComputeSchedule/scheduledActions/occurrences"),
	// 		SystemData: &armcomputeschedule.SystemData{
	// 			CreatedBy: to.Ptr("user@contoso.com"),
	// 			CreatedByType: to.Ptr(armcomputeschedule.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-04-01T18:53:38.591Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("user@contoso.com"),
	// 			LastModifiedByType: to.Ptr(armcomputeschedule.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-04-01T18:53:38.591Z"); return t}()),
	// 		},
	// 	},
	// }
}
