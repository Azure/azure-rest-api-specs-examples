package armcomputeschedule_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/computeschedule/armcomputeschedule/v2"
)

// Generated from example definition: 2026-04-15-preview/ScheduledActions_VirtualMachinesSubmitStart_MinimumSet_Gen.json
func ExampleScheduledActionsClient_VirtualMachinesSubmitStart_scheduledActionsVirtualMachinesSubmitStartMaximumSetGeneratedByMinimumSetRule() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcomputeschedule.NewClientFactory("732116BD-AF31-4E74-9283-B387C44B4A44", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewScheduledActionsClient().VirtualMachinesSubmitStart(ctx, "eastus2", armcomputeschedule.SubmitStartRequest{
		Schedule: &armcomputeschedule.Schedule{
			DeadlineType: to.Ptr(armcomputeschedule.DeadlineTypeUnknown),
		},
		ExecutionParameters: &armcomputeschedule.ExecutionParameters{},
		Resources: &armcomputeschedule.Resources{
			IDs: []*string{
				to.Ptr("/subscriptions/732116BD-AF31-4E74-9283-B387C44B4A44/resourceGroups/rgcomputeschedule/providers/Microsoft.Compute/virtualMachines/vm1"),
			},
		},
		Correlationid: to.Ptr("a1b2c3d4-e5f6-7890-abcd-ef1234567890"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcomputeschedule.ScheduledActionsClientVirtualMachinesSubmitStartResponse{
	// 	StartResourceOperationResponse: armcomputeschedule.StartResourceOperationResponse{
	// 		Type: to.Ptr("VirtualMachine"),
	// 		Location: to.Ptr("eastus2"),
	// 		Description: to.Ptr("Submit start operation completed successfully"),
	// 	},
	// }
}
