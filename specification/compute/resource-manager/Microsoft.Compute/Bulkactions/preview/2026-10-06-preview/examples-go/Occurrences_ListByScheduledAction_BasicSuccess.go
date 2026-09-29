package armbulkactions_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armbulkactions"
)

// Generated from example definition: 2026-10-06-preview/Occurrences_ListByScheduledAction_BasicSuccess.json
func ExampleOccurrencesClient_NewListByScheduledActionPager_oneListRecurringScheduledActionOccurrences() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbulkactions.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewOccurrencesClient().NewListByScheduledActionPager("example-rg", "weekday-start", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page = armbulkactions.OccurrencesClientListByScheduledActionResponse{
		// 	OccurrenceListResult: armbulkactions.OccurrenceListResult{
		// 		Value: []*armbulkactions.Occurrence{
		// 			{
		// 				Properties: &armbulkactions.OccurrenceProperties{
		// 					ScheduledTime: to.Ptr(time.Date(2026, time.September, 15, 14, 0, 0, 0, time.UTC)),
		// 					ResultSummary: &armbulkactions.OccurrenceResultSummary{
		// 						Total: to.Ptr[int32](2),
		// 						Statuses: []*armbulkactions.ResourceResultSummary{
		// 							{
		// 								Code: to.Ptr("Success"),
		// 								Count: to.Ptr[int32](2),
		// 							},
		// 						},
		// 					},
		// 					ProvisioningState: to.Ptr(armbulkactions.OccurrenceStateSucceeded),
		// 				},
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-rg/providers/Microsoft.Compute/scheduledActions/weekday-start/occurrences/77777777-7777-7777-7777-777777777777"),
		// 				Name: to.Ptr("77777777-7777-7777-7777-777777777777"),
		// 				Type: to.Ptr("Microsoft.Compute/scheduledActions/occurrences"),
		// 			},
		// 			{
		// 				Properties: &armbulkactions.OccurrenceProperties{
		// 					ScheduledTime: to.Ptr(time.Date(2026, time.September, 16, 14, 0, 0, 0, time.UTC)),
		// 					ResultSummary: &armbulkactions.OccurrenceResultSummary{
		// 						Total: to.Ptr[int32](2),
		// 						Statuses: []*armbulkactions.ResourceResultSummary{
		// 							{
		// 								Code: to.Ptr("Success"),
		// 								Count: to.Ptr[int32](2),
		// 							},
		// 						},
		// 					},
		// 					ProvisioningState: to.Ptr(armbulkactions.OccurrenceStateSucceeded),
		// 				},
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-rg/providers/Microsoft.Compute/scheduledActions/weekday-start/occurrences/88888888-8888-8888-8888-888888888888"),
		// 				Name: to.Ptr("88888888-8888-8888-8888-888888888888"),
		// 				Type: to.Ptr("Microsoft.Compute/scheduledActions/occurrences"),
		// 			},
		// 		},
		// 	},
		// }
	}
}
