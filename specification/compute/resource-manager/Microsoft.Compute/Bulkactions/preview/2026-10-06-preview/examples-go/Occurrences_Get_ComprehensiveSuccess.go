package armbulkactions_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armbulkactions"
)

// Generated from example definition: 2026-10-06-preview/Occurrences_Get_ComprehensiveSuccess.json
func ExampleOccurrencesClient_Get_twoReadARecurringScheduledActionOccurrenceWithMixedResults() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbulkactions.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewOccurrencesClient().Get(ctx, "example-rg", "weekday-start", "88888888-8888-8888-8888-888888888888", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armbulkactions.OccurrencesClientGetResponse{
	// 	Occurrence: armbulkactions.Occurrence{
	// 		Properties: &armbulkactions.OccurrenceProperties{
	// 			ScheduledTime: to.Ptr(time.Date(2026, time.September, 16, 14, 0, 0, 0, time.UTC)),
	// 			ResultSummary: &armbulkactions.OccurrenceResultSummary{
	// 				Total: to.Ptr[int32](2),
	// 				Statuses: []*armbulkactions.ResourceResultSummary{
	// 					{
	// 						Code: to.Ptr("Success"),
	// 						Count: to.Ptr[int32](1),
	// 					},
	// 					{
	// 						Code: to.Ptr("OperationNotAllowed"),
	// 						Count: to.Ptr[int32](1),
	// 						ErrorDetails: &armbulkactions.Error{
	// 							Code: to.Ptr("OperationNotAllowed"),
	// 							Message: to.Ptr("The virtual machine cannot be started while it is being deallocated."),
	// 							Target: to.Ptr("virtualMachines"),
	// 						},
	// 					},
	// 				},
	// 			},
	// 			ProvisioningState: to.Ptr(armbulkactions.OccurrenceStateFailed),
	// 		},
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-rg/providers/Microsoft.Compute/scheduledActions/weekday-start/occurrences/88888888-8888-8888-8888-888888888888"),
	// 		Name: to.Ptr("88888888-8888-8888-8888-888888888888"),
	// 		Type: to.Ptr("Microsoft.Compute/scheduledActions/occurrences"),
	// 		SystemData: &armbulkactions.SystemData{
	// 			CreatedBy: to.Ptr("user@contoso.com"),
	// 			CreatedByType: to.Ptr(armbulkactions.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(time.Date(2026, time.September, 15, 14, 0, 0, 0, time.UTC)),
	// 			LastModifiedBy: to.Ptr("user@contoso.com"),
	// 			LastModifiedByType: to.Ptr(armbulkactions.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(time.Date(2026, time.September, 16, 14, 5, 0, 0, time.UTC)),
	// 		},
	// 	},
	// }
}
