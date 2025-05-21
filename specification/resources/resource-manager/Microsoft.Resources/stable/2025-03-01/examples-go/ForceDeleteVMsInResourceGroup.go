package armresources_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ada20fa9d6dcc6453f04f32d20fe3cf97293aa88/specification/resources/resource-manager/Microsoft.Resources/stable/2025-03-01/examples/ForceDeleteVMsInResourceGroup.json
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
