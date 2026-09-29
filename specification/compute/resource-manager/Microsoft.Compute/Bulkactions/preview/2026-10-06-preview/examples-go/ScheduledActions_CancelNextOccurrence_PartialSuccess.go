package armbulkactions_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armbulkactions"
)

// Generated from example definition: 2026-10-06-preview/ScheduledActions_CancelNextOccurrence_PartialSuccess.json
func ExampleScheduledActionsClient_BeginCancelNextOccurrence_threeResponseWithPartialResultsWhenCancelingTheNextRecurringScheduledActionOccurrence() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbulkactions.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewScheduledActionsClient().BeginCancelNextOccurrence(ctx, "example-rg", "weekday-start", armbulkactions.CancelOccurrenceRequest{
		ResourceIDs: []*string{
			to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-rg/providers/Microsoft.Compute/virtualMachines/web-vm-01"),
			to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-rg/providers/Microsoft.Compute/virtualMachines/web-vm-02"),
		},
	}, nil)
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
	// res = armbulkactions.ScheduledActionsClientCancelNextOccurrenceResponse{
	// 	ResourceOperationResponse: armbulkactions.ResourceOperationResponse{
	// 		TotalResources: to.Ptr[int32](2),
	// 		ResourcesStatuses: []*armbulkactions.ResourceStatus{
	// 			{
	// 				ResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-rg/providers/Microsoft.Compute/virtualMachines/web-vm-02"),
	// 				Status: to.Ptr(armbulkactions.ResourceOperationStatusFailed),
	// 				Error: &armbulkactions.Error{
	// 					Code: to.Ptr("OccurrenceCancelResourceRejected"),
	// 					Message: to.Ptr("operation already activated; cannot cancel"),
	// 					Target: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-rg/providers/Microsoft.Compute/virtualMachines/web-vm-02"),
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
