package armcomputeschedule_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/computeschedule/armcomputeschedule"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/8a287afb3721dee0d88f11502ec123470bc52a28/specification/computeschedule/resource-manager/Microsoft.ComputeSchedule/preview/2024-08-15-preview/examples/ScheduledActions_VirtualMachinesExecuteDeallocate_MinimumSet_Gen.json
func ExampleScheduledActionsClient_VirtualMachinesExecuteDeallocate_scheduledActionsVirtualMachinesExecuteDeallocateGeneratedByMaximumSetRuleGeneratedByMinimumSetRule() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcomputeschedule.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewScheduledActionsClient().VirtualMachinesExecuteDeallocate(ctx, "irbtacrhjgfpyvhvjo", armcomputeschedule.ExecuteDeallocateRequest{
		Correlationid:       to.Ptr("01080d2f-1dca-4610-afb4-dd25eddf1f3c1"),
		ExecutionParameters: &armcomputeschedule.ExecutionParameters{},
		Resources: &armcomputeschedule.Resources{
			IDs: []*string{
				to.Ptr("/subscriptions/YourSubscriptionId/resourceGroups/YourResourceGroupName/providers/Microsoft.Compute/virtualMachines/testResource2")},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DeallocateResourceOperationResponse = armcomputeschedule.DeallocateResourceOperationResponse{
	// 	Type: to.Ptr("eokrhxypminzcgplibblmwojvjd"),
	// 	Description: to.Ptr("ndhb"),
	// 	Location: to.Ptr("qbbvrdisvjgxirnzdvdsodkt"),
	// }
}
