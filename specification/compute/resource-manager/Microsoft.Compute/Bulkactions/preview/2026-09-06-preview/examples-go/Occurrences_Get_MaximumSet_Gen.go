package armbulkactions_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armbulkactions"
)

// Generated from example definition: 2026-09-06-preview/Occurrences_Get_MaximumSet_Gen.json
func ExampleOccurrencesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbulkactions.NewClientFactory("CB26D7CB-3E27-465F-99C8-EAF7A4118245", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewOccurrencesClient().Get(ctx, "rgcompute", "myScheduledAction", "67b5bada-4772-43fc-8dbb-402476d98a45", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armbulkactions.OccurrencesClientGetResponse{
	// 	Occurrence: armbulkactions.Occurrence{
	// 		Properties: &armbulkactions.OccurrenceProperties{
	// 			ScheduledTime: to.Ptr(time.Date(2026, time.August, 5, 0, 30, 0, 0, time.UTC)),
	// 			ResultSummary: &armbulkactions.OccurrenceResultSummary{
	// 				Total: to.Ptr[int32](25),
	// 				Statuses: []*armbulkactions.ResourceResultSummary{
	// 					{
	// 						Code: to.Ptr("Succeeded"),
	// 						Count: to.Ptr[int32](24),
	// 					},
	// 					{
	// 						Code: to.Ptr("Failed"),
	// 						Count: to.Ptr[int32](1),
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
	// 			ProvisioningState: to.Ptr(armbulkactions.OccurrenceStateFailed),
	// 		},
	// 		ID: to.Ptr("/subscriptions/CB26D7CB-3E27-465F-99C8-EAF7A4118245/resourceGroups/rgcompute/providers/Microsoft.Compute/scheduledActions/myScheduledAction/occurrences/67b5bada-4772-43fc-8dbb-402476d98a45"),
	// 		Name: to.Ptr("67b5bada-4772-43fc-8dbb-402476d98a45"),
	// 		Type: to.Ptr("Microsoft.Compute/scheduledActions/occurrences"),
	// 		SystemData: &armbulkactions.SystemData{
	// 			CreatedBy: to.Ptr("user@contoso.com"),
	// 			CreatedByType: to.Ptr(armbulkactions.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(time.Date(2026, time.August, 1, 12, 0, 0, 0, time.UTC)),
	// 			LastModifiedBy: to.Ptr("user@contoso.com"),
	// 			LastModifiedByType: to.Ptr(armbulkactions.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(time.Date(2026, time.August, 5, 0, 31, 0, 0, time.UTC)),
	// 		},
	// 	},
	// }
}
