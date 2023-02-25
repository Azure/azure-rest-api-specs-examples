package armcontainerregistry_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerregistry/armcontainerregistry"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/270d3cd664cca3ddc8511f92d3851a715e2c61db/specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2019-06-01-preview/examples/AgentPoolsUpdate.json
func ExampleAgentPoolsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcontainerregistry.NewAgentPoolsClient("4385cf00-2d3a-425a-832f-f4285b1c9dce", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginUpdate(ctx, "myResourceGroup", "myRegistry", "myAgentPool", armcontainerregistry.AgentPoolUpdateParameters{
		Properties: &armcontainerregistry.AgentPoolPropertiesUpdateParameters{
			Count: to.Ptr[int32](1),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AgentPool = armcontainerregistry.AgentPool{
	// 	Name: to.Ptr("myAgentPool"),
	// 	Type: to.Ptr("Microsoft.ContainerRegistry/registries/agentPools"),
	// 	ID: to.Ptr("/subscriptions/f9d7ebed-adbd-4cb4-b973-aaf82c136138/resourceGroups/huanwudfwestgroup/providers/Microsoft.ContainerRegistry/registries/huanglidfwest01/agentPools/testagent26"),
	// 	Location: to.Ptr("WESTUS"),
	// 	Properties: &armcontainerregistry.AgentPoolProperties{
	// 		Count: to.Ptr[int32](1),
	// 		OS: to.Ptr(armcontainerregistry.OSLinux),
	// 		ProvisioningState: to.Ptr(armcontainerregistry.ProvisioningStateSucceeded),
	// 		Tier: to.Ptr("S1"),
	// 	},
	// }
}
