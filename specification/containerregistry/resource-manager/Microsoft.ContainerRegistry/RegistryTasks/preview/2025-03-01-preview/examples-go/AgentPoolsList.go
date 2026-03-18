package armcontainerregistrytasks_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerregistry/armcontainerregistrytasks"
)

// Generated from example definition: 2025-03-01-preview/AgentPoolsList.json
func ExampleAgentPoolsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerregistrytasks.NewClientFactory("4385cf00-2d3a-425a-832f-f4285b1c9dce", cred, nil)
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
		// page = armcontainerregistrytasks.AgentPoolsClientListResponse{
		// 	AgentPoolListResult: armcontainerregistrytasks.AgentPoolListResult{
		// 		Value: []*armcontainerregistrytasks.AgentPool{
		// 			{
		// 				Name: to.Ptr("myAgentPool"),
		// 				Type: to.Ptr("Microsoft.ContainerRegistry/registries/agentPools"),
		// 				ID: to.Ptr("/subscriptions/f9d7ebed-adbd-4cb4-b973-aaf82c136138/resourceGroups/huanwudfwestgroup/providers/Microsoft.ContainerRegistry/registries/huanglidfwest01/agentPools/testagent26"),
		// 				Location: to.Ptr("WESTUS"),
		// 				Properties: &armcontainerregistrytasks.AgentPoolProperties{
		// 					Count: to.Ptr[int32](1),
		// 					OS: to.Ptr(armcontainerregistrytasks.OSLinux),
		// 					ProvisioningState: to.Ptr(armcontainerregistrytasks.ProvisioningStateSucceeded),
		// 					Tier: to.Ptr("S1"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
