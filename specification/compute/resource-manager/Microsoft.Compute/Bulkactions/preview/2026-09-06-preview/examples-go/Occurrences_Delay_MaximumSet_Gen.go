package armbulkactions_test

import (
	"context"
	"log"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armbulkactions"
)

// Generated from example definition: 2026-09-06-preview/Occurrences_Delay_MaximumSet_Gen.json
func ExampleOccurrencesClient_BeginDelay() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbulkactions.NewClientFactory("CB26D7CB-3E27-465F-99C8-EAF7A4118245", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewOccurrencesClient().BeginDelay(ctx, "rgcompute", "myScheduledAction", "67b5bada-4772-43fc-8dbb-402476d98a45", armbulkactions.DelayRequest{
		Delay: to.Ptr(time.Date(2026, time.August, 5, 17, 0, 0, 0, time.FixedZone("", -25200))),
		ResourceIDs: []*string{
			to.Ptr("/subscriptions/CB26D7CB-3E27-465F-99C8-EAF7A4118245/resourceGroups/rgcompute/providers/Microsoft.Compute/virtualMachines/myVm"),
			to.Ptr("/subscriptions/CB26D7CB-3E27-465F-99C8-EAF7A4118245/resourceGroups/rgcompute/providers/Microsoft.Compute/virtualMachines/myVm2"),
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
	// res = armbulkactions.OccurrencesClientDelayResponse{
	// 	ResourceOperationResponse: armbulkactions.ResourceOperationResponse{
	// 		TotalResources: to.Ptr[int32](2),
	// 		ResourcesStatuses: []*armbulkactions.ResourceStatus{
	// 			{
	// 				ResourceID: to.Ptr("/subscriptions/CB26D7CB-3E27-465F-99C8-EAF7A4118245/resourceGroups/rgcompute/providers/Microsoft.Compute/virtualMachines/myVm"),
	// 				Status: to.Ptr(armbulkactions.ResourceOperationStatusSucceeded),
	// 			},
	// 			{
	// 				ResourceID: to.Ptr("/subscriptions/CB26D7CB-3E27-465F-99C8-EAF7A4118245/resourceGroups/rgcompute/providers/Microsoft.Compute/virtualMachines/myVm2"),
	// 				Status: to.Ptr(armbulkactions.ResourceOperationStatusFailed),
	// 				Error: &armbulkactions.Error{
	// 					Code: to.Ptr("InternalServerError"),
	// 					Message: to.Ptr("An internal error occurred."),
	// 					Target: to.Ptr("virtualMachines"),
	// 					Details: []*armbulkactions.Error{
	// 					},
	// 					Innererror: &armbulkactions.InnerError{
	// 						Code: to.Ptr("InnerErrorCode"),
	// 					},
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
