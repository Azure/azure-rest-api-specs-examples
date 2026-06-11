package armcomputeschedule_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/computeschedule/armcomputeschedule/v2"
)

// Generated from example definition: 2026-04-15-preview/Occurrences_Get_MaximumSet_Gen.json
func ExampleOccurrencesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcomputeschedule.NewClientFactory("732116BD-AF31-4E74-9283-B387C44B4A44", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewOccurrencesClient().Get(ctx, "rgcomputeschedule", "scheduled-action-01", "11111111-1111-1111-1111-111111111111", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcomputeschedule.OccurrencesClientGetResponse{
	// 	Occurrence: armcomputeschedule.Occurrence{
	// 		Properties: &armcomputeschedule.OccurrenceProperties{
	// 			ScheduledTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-03-12T02:39:47.285Z"); return t}()),
	// 			ResultSummary: &armcomputeschedule.OccurrenceResultSummary{
	// 				Total: to.Ptr[int32](3),
	// 				Statuses: []*armcomputeschedule.ResourceResultSummary{
	// 					{
	// 						Code: to.Ptr("Success"),
	// 						Count: to.Ptr[int32](2),
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
	// 		ID: to.Ptr("/subscriptions/83C27AB3-A7B9-498B-B165-D9440661474F/resourceGroups/myRg/providers/Microsoft.ComputeSchedule/scheduledActions/scheduled-action-01/occurrences/11111111-1111-1111-1111-111111111111"),
	// 		Name: to.Ptr("11111111-1111-1111-1111-111111111111"),
	// 		Type: to.Ptr("Microsoft.ComputeSchedule/scheduledActions/occurrences"),
	// 		SystemData: &armcomputeschedule.SystemData{
	// 			CreatedBy: to.Ptr("user@contoso.com"),
	// 			CreatedByType: to.Ptr(armcomputeschedule.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-03-12T02:39:41.641Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("user@contoso.com"),
	// 			LastModifiedByType: to.Ptr(armcomputeschedule.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-03-12T02:39:41.641Z"); return t}()),
	// 		},
	// 	},
	// }
}
