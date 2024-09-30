package armcomputeschedule_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/computeschedule/armcomputeschedule"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/8a287afb3721dee0d88f11502ec123470bc52a28/specification/computeschedule/resource-manager/Microsoft.ComputeSchedule/preview/2024-08-15-preview/examples/ScheduledActions_VirtualMachinesExecuteStart_MaximumSet_Gen.json
func ExampleScheduledActionsClient_VirtualMachinesExecuteStart_scheduledActionsVirtualMachinesExecuteStartGeneratedByMaximumSetRuleGeneratedByMaximumSetRule() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcomputeschedule.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewScheduledActionsClient().VirtualMachinesExecuteStart(ctx, "ysfrwcfmfsh", armcomputeschedule.ExecuteStartRequest{
		Correlationid: to.Ptr("23230d2f-1dca-4610-afb4-dd25eec1f34"),
		ExecutionParameters: &armcomputeschedule.ExecutionParameters{
			OptimizationPreference: to.Ptr(armcomputeschedule.OptimizationPreferenceCost),
			RetryPolicy: &armcomputeschedule.RetryPolicy{
				RetryCount:           to.Ptr[int32](30),
				RetryWindowInMinutes: to.Ptr[int32](27),
			},
		},
		Resources: &armcomputeschedule.Resources{
			IDs: []*string{
				to.Ptr("/subscriptions/YourSubscriptionId/resourceGroups/YourResourceGroupName/providers/Microsoft.Compute/virtualMachines/testResource4")},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.StartResourceOperationResponse = armcomputeschedule.StartResourceOperationResponse{
	// 	Type: to.Ptr("gxmnjtgu"),
	// 	Description: to.Ptr("raacd"),
	// 	Location: to.Ptr("uvlidhowwv"),
	// 	Results: []*armcomputeschedule.ResourceOperation{
	// 		{
	// 			ErrorCode: to.Ptr("pliurpaykwv"),
	// 			ErrorDetails: to.Ptr("mbqwkpxl"),
	// 			Operation: &armcomputeschedule.ResourceOperationDetails{
	// 				OperationID: to.Ptr("23480d2f-1dca-4610-afb4-dd25eec1f34"),
	// 				CompletedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-07-12T18:06:52.974Z"); return t}()),
	// 				Deadline: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-07-12T18:06:52.974Z"); return t}()),
	// 				DeadlineType: to.Ptr(armcomputeschedule.DeadlineTypeUnknown),
	// 				OpType: to.Ptr(armcomputeschedule.ResourceOperationTypeUnknown),
	// 				ResourceID: to.Ptr("/subscriptions/YourSubscriptionId/resourceGroups/YourResourceGroupName/providers/Microsoft.Compute/virtualMachines/testResource4"),
	// 				ResourceOperationError: &armcomputeschedule.ResourceOperationError{
	// 					ErrorCode: to.Ptr("fticyqukvkillwd"),
	// 					ErrorDetails: to.Ptr("yimgxqrkp"),
	// 				},
	// 				RetryPolicy: &armcomputeschedule.RetryPolicy{
	// 					RetryCount: to.Ptr[int32](30),
	// 					RetryWindowInMinutes: to.Ptr[int32](27),
	// 				},
	// 				State: to.Ptr(armcomputeschedule.OperationStateUnknown),
	// 				SubscriptionID: to.Ptr("52C81249-550F-459E-9B6E-5BAB6EE62227"),
	// 				TimeZone: to.Ptr("lsoalxijilbrqum"),
	// 			},
	// 			ResourceID: to.Ptr("/subscriptions/YourSubscriptionId/resourceGroups/YourResourceGroupName/providers/Microsoft.Compute/virtualMachines/testResource4"),
	// 	}},
	// }
}
