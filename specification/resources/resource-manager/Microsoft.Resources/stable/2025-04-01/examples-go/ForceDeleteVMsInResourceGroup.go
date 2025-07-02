package armresources_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4a2bb0762eaad11e725516708483598e0c12cabb/specification/resources/resource-manager/Microsoft.Resources/stable/2025-04-01/examples/ForceDeleteVMsInResourceGroup.json
func ExampleResourceGroupsClient_BeginDelete_forceDeleteAllTheVirtualMachinesInAResourceGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armresources.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewResourceGroupsClient().BeginDelete(ctx, "my-resource-group", &armresources.ResourceGroupsClientBeginDeleteOptions{ForceDeletionTypes: to.Ptr("Microsoft.Compute/virtualMachines")})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
