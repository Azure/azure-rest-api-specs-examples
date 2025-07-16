package armcomputeschedule_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/computeschedule/armcomputeschedule"
)

// Generated from example definition: 2025-05-01/ScheduledActions_VirtualMachinesExecuteHibernate_MaximumSet_Gen.json
func ExampleScheduledActionsClient_VirtualMachinesExecuteHibernate_scheduledActionsVirtualMachinesExecuteHibernateMaximumSetGenGeneratedByMaximumSetRule() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcomputeschedule.NewClientFactory("0505D8E4-D41A-48FB-9CA5-4AF8D93BE75F", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewScheduledActionsClient().VirtualMachinesExecuteHibernate(ctx, "pxjjubifupovxakrcjafxrcbgizolx", armcomputeschedule.ExecuteHibernateRequest{
		ExecutionParameters: &armcomputeschedule.ExecutionParameters{
			OptimizationPreference: to.Ptr(armcomputeschedule.OptimizationPreferenceCost),
			RetryPolicy: &armcomputeschedule.RetryPolicy{
				RetryCount:           to.Ptr[int32](25),
				RetryWindowInMinutes: to.Ptr[int32](4),
			},
		},
		Resources: &armcomputeschedule.Resources{
			IDs: []*string{
				to.Ptr("/subscriptions/YourSubscriptionId/resourceGroups/YourResourceGroupName/providers/Microsoft.Compute/virtualMachines/testResource3"),
			},
		},
		Correlationid: to.Ptr("jmdiz"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcomputeschedule.ScheduledActionsClientVirtualMachinesExecuteHibernateResponse{
	// 	HibernateResourceOperationResponse: &armcomputeschedule.HibernateResourceOperationResponse{
	// 		Type: to.Ptr("yrmuumqaqiyotst"),
	// 		Location: to.Ptr("hhioerbsucdqayxk"),
	// 		Results: []*armcomputeschedule.ResourceOperation{
	// 			{
	// 				ResourceID: to.Ptr("/subscriptions/YourSubscriptionId/resourceGroups/YourResourceGroupName/providers/Microsoft.Compute/virtualMachines/testResource3"),
	// 				ErrorCode: to.Ptr("ynukyltendgmn"),
	// 				ErrorDetails: to.Ptr("tifeuh"),
	// 				Operation: &armcomputeschedule.ResourceOperationDetails{
	// 					OperationID: to.Ptr("vppyaxq"),
	// 					ResourceID: to.Ptr("/subscriptions/YourSubscriptionId/resourceGroups/YourResourceGroupName/providers/Microsoft.Compute/virtualMachines/testResource3"),
	// 					OpType: to.Ptr(armcomputeschedule.ResourceOperationTypeUnknown),
	// 					SubscriptionID: to.Ptr("vofvsus"),
	// 					Deadline: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-04-15T19:47:03.591Z"); return t}()),
	// 					DeadlineType: to.Ptr(armcomputeschedule.DeadlineTypeUnknown),
	// 					State: to.Ptr(armcomputeschedule.OperationStateUnknown),
	// 					Timezone: to.Ptr("nwugsooykqggcokphgdj"),
	// 					TimeZone: to.Ptr("qkxnxnumvfqmsmpyccv"),
	// 					ResourceOperationError: &armcomputeschedule.ResourceOperationError{
	// 						ErrorCode: to.Ptr("fagfsojftlff"),
	// 						ErrorDetails: to.Ptr("rtihrkjasrjkllqccuysjrg"),
	// 					},
	// 					CompletedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-04-15T19:47:03.591Z"); return t}()),
	// 					RetryPolicy: &armcomputeschedule.RetryPolicy{
	// 						RetryCount: to.Ptr[int32](25),
	// 						RetryWindowInMinutes: to.Ptr[int32](4),
	// 					},
	// 				},
	// 			},
	// 		},
	// 		Description: to.Ptr("ploigcuqj"),
	// 	},
	// }
}
