package armcomputeschedule_test

import (
	"context"
	"log"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/computeschedule/armcomputeschedule/v2"
)

// Generated from example definition: 2026-04-15-preview/Occurrences_Delay_MaximumSet_Gen.json
func ExampleOccurrencesClient_BeginDelay() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcomputeschedule.NewClientFactory("732116BD-AF31-4E74-9283-B387C44B4A44", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewOccurrencesClient().BeginDelay(ctx, "rgcomputeschedule", "scheduled-action-01", "11111111-1111-1111-1111-111111111111", armcomputeschedule.DelayRequest{
		Delay: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-03-12T02:39:48.148Z"); return t }()),
		ResourceIDs: []*string{
			to.Ptr("/subscriptions/11111111-1111-1111-1111-111111111111/resourceGroups/rgcomputeschedule/providers/Microsoft.Compute/virtualMachines/vm1"),
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
	// res = armcomputeschedule.OccurrencesClientDelayResponse{
	// 	RecurringActionsResourceOperationResult: armcomputeschedule.RecurringActionsResourceOperationResult{
	// 		TotalResources: to.Ptr[int32](4),
	// 		ResourcesStatuses: []*armcomputeschedule.ResourceStatus{
	// 			{
	// 				ResourceID: to.Ptr("/subscriptions/1d04e8f1-ee04-4056-b0b2-718f5bb45b04/resourceGroups/myRg/providers/Microsoft.Compute/virtualMachines/myVm"),
	// 				Status: to.Ptr(armcomputeschedule.ResourceOperationStatusSucceeded),
	// 				Error: &armcomputeschedule.Error{
	// 					Code: to.Ptr("ResourceNotFound"),
	// 					Message: to.Ptr("The specified resource was not found."),
	// 					Target: to.Ptr("virtualMachines/myVm"),
	// 					Details: []*armcomputeschedule.Error{
	// 					},
	// 					Innererror: &armcomputeschedule.InnerError{
	// 						Code: to.Ptr("ResourceNotFoundError"),
	// 					},
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
