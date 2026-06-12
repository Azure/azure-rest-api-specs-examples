package armcomputeschedule_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/computeschedule/armcomputeschedule/v2"
)

// Generated from example definition: 2026-04-15-preview/ScheduledActions_VirtualMachinesGetOperationErrors_MinimumSet_Gen.json
func ExampleScheduledActionsClient_VirtualMachinesGetOperationErrors_scheduledActionsVirtualMachinesGetOperationErrorsMinimumSet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcomputeschedule.NewClientFactory("732116BD-AF31-4E74-9283-B387C44B4A44", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewScheduledActionsClient().VirtualMachinesGetOperationErrors(ctx, "eastus2", armcomputeschedule.GetOperationErrorsRequest{
		OperationIDs: []*string{
			to.Ptr("01234567-89ab-cdef-0123-456789abcdef"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcomputeschedule.ScheduledActionsClientVirtualMachinesGetOperationErrorsResponse{
	// 	GetOperationErrorsResponse: armcomputeschedule.GetOperationErrorsResponse{
	// 		Results: []*armcomputeschedule.OperationErrorsResult{
	// 			{
	// 			},
	// 		},
	// 	},
	// }
}
