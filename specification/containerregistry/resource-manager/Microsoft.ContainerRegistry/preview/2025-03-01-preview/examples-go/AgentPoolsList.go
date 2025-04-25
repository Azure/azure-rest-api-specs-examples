package armcontainerregistry_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerregistry/armcontainerregistry"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/dc4c1eaef16e0bc8b1e96c3d1e014deb96259b35/specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2025-03-01-preview/examples/AgentPoolsList.json
func ExampleAgentPoolsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerregistry.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAgentPoolsClient().NewListPager("myResourceGroup", "myRegistry", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.AgentPoolListResult = armcontainerregistry.AgentPoolListResult{
		// 	Value: []*armcontainerregistry.AgentPool{
		// 		{
		// 			Name: to.Ptr("myAgentPool"),
		// 			Type: to.Ptr("Microsoft.ContainerRegistry/registries/agentPools"),
		// 			ID: to.Ptr("/subscriptions/f9d7ebed-adbd-4cb4-b973-aaf82c136138/resourceGroups/huanwudfwestgroup/providers/Microsoft.ContainerRegistry/registries/huanglidfwest01/agentPools/testagent26"),
		// 			Location: to.Ptr("WESTUS"),
		// 			Properties: &armcontainerregistry.AgentPoolProperties{
		// 				Count: to.Ptr[int32](1),
		// 				OS: to.Ptr(armcontainerregistry.OSLinux),
		// 				ProvisioningState: to.Ptr(armcontainerregistry.ProvisioningStateSucceeded),
		// 				Tier: to.Ptr("S1"),
		// 			},
		// 	}},
		// }
	}
}
