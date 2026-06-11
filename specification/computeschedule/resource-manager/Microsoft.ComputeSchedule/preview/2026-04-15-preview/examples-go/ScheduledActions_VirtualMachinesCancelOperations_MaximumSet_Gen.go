package armcomputeschedule_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/computeschedule/armcomputeschedule/v2"
)

// Generated from example definition: 2026-04-15-preview/ScheduledActions_VirtualMachinesCancelOperations_MaximumSet_Gen.json
func ExampleScheduledActionsClient_VirtualMachinesCancelOperations_scheduledActionsVirtualMachinesCancelOperationsMaximumSet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcomputeschedule.NewClientFactory("732116BD-AF31-4E74-9283-B387C44B4A44", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewScheduledActionsClient().VirtualMachinesCancelOperations(ctx, "eastus2", armcomputeschedule.CancelOperationsRequest{
		OperationIDs: []*string{
			to.Ptr("01234567-89ab-cdef-0123-456789abcdef"),
		},
		Correlationid: to.Ptr("a1b2c3d4-e5f6-7890-abcd-ef1234567890"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcomputeschedule.ScheduledActionsClientVirtualMachinesCancelOperationsResponse{
	// 	CancelOperationsResponse: armcomputeschedule.CancelOperationsResponse{
	// 		Results: []*armcomputeschedule.ResourceOperation{
	// 			{
	// 				ResourceID: to.Ptr("/subscriptions/732116BD-AF31-4E74-9283-B387C44B4A44/resourceGroups/rgcomputeschedule/providers/Microsoft.Compute/virtualMachines/myVm"),
	// 				ErrorCode: to.Ptr("Success"),
	// 				ErrorDetails: to.Ptr(""),
	// 				Operation: &armcomputeschedule.ResourceOperationDetails{
	// 					OperationID: to.Ptr("ead40c08-1b98-40cf-848d-fe373146b724"),
	// 					ResourceID: to.Ptr("/subscriptions/732116BD-AF31-4E74-9283-B387C44B4A44/resourceGroups/rgcomputeschedule/providers/Microsoft.Compute/virtualMachines/myVm"),
	// 					OpType: to.Ptr(armcomputeschedule.ResourceOperationTypeUnknown),
	// 					SubscriptionID: to.Ptr("732116BD-AF31-4E74-9283-B387C44B4A44"),
	// 					Deadline: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-03-12T02:39:42.468Z"); return t}()),
	// 					DeadlineType: to.Ptr(armcomputeschedule.DeadlineTypeUnknown),
	// 					State: to.Ptr(armcomputeschedule.OperationStateUnknown),
	// 					Timezone: to.Ptr("UTC"),
	// 					TimeZone: to.Ptr("UTC"),
	// 					ResourceOperationError: &armcomputeschedule.ResourceOperationError{
	// 						ErrorCode: to.Ptr(""),
	// 						ErrorDetails: to.Ptr(""),
	// 					},
	// 					FallbackOperationInfo: &armcomputeschedule.FallbackOperationInfo{
	// 						LastOpType: to.Ptr(armcomputeschedule.ResourceOperationTypeUnknown),
	// 						Status: to.Ptr("Succeeded"),
	// 						Error: &armcomputeschedule.ResourceOperationError{
	// 							ErrorCode: to.Ptr(""),
	// 							ErrorDetails: to.Ptr(""),
	// 						},
	// 					},
	// 					CompletedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-03-12T02:39:42.468Z"); return t}()),
	// 					RetryPolicy: &armcomputeschedule.RetryPolicy{
	// 						RetryCount: to.Ptr[int32](3),
	// 						RetryWindowInMinutes: to.Ptr[int32](30),
	// 						OnFailureAction: to.Ptr(armcomputeschedule.ResourceOperationTypeUnknown),
	// 					},
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
