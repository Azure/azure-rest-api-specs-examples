package armcomputeschedule_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/computeschedule/armcomputeschedule"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/8a4eca6b060cf70da696963245656fdc440b666b/specification/computeschedule/resource-manager/Microsoft.ComputeSchedule/stable/2024-10-01/examples/ScheduledActions_VirtualMachinesExecuteHibernate.json
func ExampleScheduledActionsClient_VirtualMachinesExecuteHibernate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcomputeschedule.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewScheduledActionsClient().VirtualMachinesExecuteHibernate(ctx, "eastus2euap", armcomputeschedule.ExecuteHibernateRequest{
		Correlationid: to.Ptr("23480d2f-1dca-4610-afb4-dd25eec1f34r"),
		ExecutionParameters: &armcomputeschedule.ExecutionParameters{
			RetryPolicy: &armcomputeschedule.RetryPolicy{
				RetryCount:           to.Ptr[int32](5),
				RetryWindowInMinutes: to.Ptr[int32](27),
			},
		},
		Resources: &armcomputeschedule.Resources{
			IDs: []*string{
				to.Ptr("/subscriptions/YourSubscriptionId/resourceGroups/YourResourceGroupName/providers/Microsoft.Compute/virtualMachines/testResource3")},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.HibernateResourceOperationResponse = armcomputeschedule.HibernateResourceOperationResponse{
	// 	Type: to.Ptr("VirtualMachine"),
	// 	Description: to.Ptr("Hibernate Resource Request"),
	// 	Location: to.Ptr("eastus2euap"),
	// 	Results: []*armcomputeschedule.ResourceOperation{
	// 		{
	// 			ErrorCode: to.Ptr("null"),
	// 			ErrorDetails: to.Ptr("null"),
	// 			Operation: &armcomputeschedule.ResourceOperationDetails{
	// 				OperationID: to.Ptr("23480d2f-1dca-4610-afb4-dd25eec1f34r"),
	// 				CompletedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-11-01T17:52:53.668Z"); return t}()),
	// 				Deadline: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-11-01T17:52:53.667Z"); return t}()),
	// 				DeadlineType: to.Ptr(armcomputeschedule.DeadlineTypeInitiateAt),
	// 				OpType: to.Ptr(armcomputeschedule.ResourceOperationTypeHibernate),
	// 				ResourceID: to.Ptr("/subscriptions/YourSubscriptionId/resourceGroups/YourResourceGroupName/providers/Microsoft.Compute/virtualMachines/testResource3"),
	// 				ResourceOperationError: &armcomputeschedule.ResourceOperationError{
	// 					ErrorCode: to.Ptr("null"),
	// 					ErrorDetails: to.Ptr("null"),
	// 				},
	// 				RetryPolicy: &armcomputeschedule.RetryPolicy{
	// 					RetryCount: to.Ptr[int32](5),
	// 					RetryWindowInMinutes: to.Ptr[int32](27),
	// 				},
	// 				State: to.Ptr(armcomputeschedule.OperationStateSucceeded),
	// 				SubscriptionID: to.Ptr("D8E30CC0-2763-4FCC-84A8-3C5659281032"),
	// 				TimeZone: to.Ptr("UTC"),
	// 				Timezone: to.Ptr("UTC"),
	// 			},
	// 			ResourceID: to.Ptr("/subscriptions/YourSubscriptionId/resourceGroups/YourResourceGroupName/providers/Microsoft.Compute/virtualMachines/testResource3"),
	// 	}},
	// }
}
