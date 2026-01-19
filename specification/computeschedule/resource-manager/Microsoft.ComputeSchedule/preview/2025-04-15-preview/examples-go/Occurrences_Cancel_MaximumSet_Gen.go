package armcomputeschedule_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/computeschedule/armcomputeschedule"
)

// Generated from example definition: 2025-04-15-preview/Occurrences_Cancel_MaximumSet_Gen.json
func ExampleOccurrencesClient_Cancel() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcomputeschedule.NewClientFactory("CB26D7CB-3E27-465F-99C8-EAF7A4118245", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewOccurrencesClient().Cancel(ctx, "rgcomputeschedule", "myScheduledAction", "CB26D7CB-3E27-465F-99C8-EAF7A4118245", armcomputeschedule.CancelOccurrenceRequest{
		ResourceIDs: []*string{
			to.Ptr("/subscriptions/CB26D7CB-3E27-465F-99C8-EAF7A4118245/resourceGroups/myRg/providers/Microsoft.Compute/virtualMachines/myVm"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcomputeschedule.OccurrencesClientCancelResponse{
	// 	RecurringActionsResourceOperationResult: &armcomputeschedule.RecurringActionsResourceOperationResult{
	// 		TotalResources: to.Ptr[int32](11),
	// 		ResourcesStatuses: []*armcomputeschedule.ResourceStatus{
	// 			{
	// 				ResourceID: to.Ptr("/subscriptions/CB26D7CB-3E27-465F-99C8-EAF7A4118245/resourceGroups/myRg/providers/Microsoft.Compute/virtualMachines/myVm"),
	// 				Status: to.Ptr(armcomputeschedule.ResourceOperationStatusSucceeded),
	// 				Error: &armcomputeschedule.Error{
	// 					Code: to.Ptr("baxjmkbhoatqcj"),
	// 					Message: to.Ptr("chapcwfkqymeof"),
	// 					Target: to.Ptr("mkirmorowetsigohjamvk"),
	// 					Details: []*armcomputeschedule.Error{
	// 					},
	// 					Innererror: &armcomputeschedule.InnerError{
	// 						Code: to.Ptr("cgalioufsabcwatbxa"),
	// 					},
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
