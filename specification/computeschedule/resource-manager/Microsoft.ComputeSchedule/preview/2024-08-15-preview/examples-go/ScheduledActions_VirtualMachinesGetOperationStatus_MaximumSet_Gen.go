package armcomputeschedule_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/computeschedule/armcomputeschedule"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/8a287afb3721dee0d88f11502ec123470bc52a28/specification/computeschedule/resource-manager/Microsoft.ComputeSchedule/preview/2024-08-15-preview/examples/ScheduledActions_VirtualMachinesGetOperationStatus_MaximumSet_Gen.json
func ExampleScheduledActionsClient_VirtualMachinesGetOperationStatus_scheduledActionsVirtualMachinesGetOperationStatusGeneratedByMinimumSetRuleGeneratedByMaximumSetRule() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcomputeschedule.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewScheduledActionsClient().VirtualMachinesGetOperationStatus(ctx, "htxxasfxaesuecem", armcomputeschedule.GetOperationStatusRequest{
		Correlationid: to.Ptr("01080d2f-1dca-4610-afb4-dd25eec1f3c1"),
		OperationIDs: []*string{
			to.Ptr("01080d2f-1dca-4610-afb4-dd25eec1f3df")},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.GetOperationStatusResponse = armcomputeschedule.GetOperationStatusResponse{
	// 	Results: []*armcomputeschedule.ResourceOperation{
	// 		{
	// 			ErrorCode: to.Ptr("ipyvwgk"),
	// 			ErrorDetails: to.Ptr("dfegatwcb"),
	// 			Operation: &armcomputeschedule.ResourceOperationDetails{
	// 				OperationID: to.Ptr("brvr"),
	// 				CompletedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-08-08T19:00:13.928Z"); return t}()),
	// 				Deadline: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-08-08T19:00:13.928Z"); return t}()),
	// 				DeadlineType: to.Ptr(armcomputeschedule.DeadlineTypeUnknown),
	// 				OpType: to.Ptr(armcomputeschedule.ResourceOperationTypeUnknown),
	// 				ResourceID: to.Ptr("/subscriptions/YourSubscriptionId/resourceGroups/YourResourceGroupName/providers/Microsoft.Compute/virtualMachines/testResource3"),
	// 				ResourceOperationError: &armcomputeschedule.ResourceOperationError{
	// 					ErrorCode: to.Ptr("ryxlbd"),
	// 					ErrorDetails: to.Ptr("vdwoj"),
	// 				},
	// 				RetryPolicy: &armcomputeschedule.RetryPolicy{
	// 					RetryCount: to.Ptr[int32](5),
	// 					RetryWindowInMinutes: to.Ptr[int32](17),
	// 				},
	// 				State: to.Ptr(armcomputeschedule.OperationStateUnknown),
	// 				SubscriptionID: to.Ptr("l"),
	// 				TimeZone: to.Ptr("ggsopagtvrtsjltyyccqgufgm"),
	// 			},
	// 			ResourceID: to.Ptr("/subscriptions/YourSubscriptionId/resourceGroups/YourResourceGroupName/providers/Microsoft.Compute/virtualMachines/testResource3"),
	// 	}},
	// }
}
