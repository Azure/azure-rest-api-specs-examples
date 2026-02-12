package armcomputebulkactions_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/computebulkactions/armcomputebulkactions"
)

// Generated from example definition: 2026-02-01-preview/BulkActions_VirtualMachinesExecuteCreate_MinimumSet_Gen.json
func ExampleBulkActionsClient_VirtualMachinesExecuteCreate_bulkActionsVirtualMachinesExecuteCreateMinimumSetGenGeneratedByMinimumSetRuleGeneratedByMinimumSetRule() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcomputebulkactions.NewClientFactory("50352BBD-59F1-4EE2-BA9C-A6E51B0B1B2B", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewBulkActionsClient().VirtualMachinesExecuteCreate(ctx, "eastus2euap", armcomputebulkactions.ExecuteCreateRequest{
		ResourceConfigParameters: &armcomputebulkactions.ResourceProvisionPayload{
			ResourceCount: to.Ptr[int32](1),
		},
		ExecutionParameters: &armcomputebulkactions.ExecutionParameters{},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcomputebulkactions.BulkActionsClientVirtualMachinesExecuteCreateResponse{
	// 	CreateResourceOperationResponse: &armcomputebulkactions.CreateResourceOperationResponse{
	// 		Type: to.Ptr("VirtualMachine"),
	// 		Location: to.Ptr("eastus2euap"),
	// 		Description: to.Ptr("Create Resource Request"),
	// 	},
	// }
}
