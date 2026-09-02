package armcontainerservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v9"
)

// Generated from example definition: 2026-06-02-preview/AgentPoolsCreate_FlexNode.json
func ExampleAgentPoolsClient_BeginCreateOrUpdate_createFlexNodeAgentPool() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerservice.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAgentPoolsClient().BeginCreateOrUpdate(ctx, "rg1", "clustername1", "flexnode1", armcontainerservice.AgentPool{
		Properties: &armcontainerservice.ManagedClusterAgentPoolProfileProperties{
			Mode:                to.Ptr(armcontainerservice.AgentPoolModeUser),
			OrchestratorVersion: to.Ptr("1.32"),
			Type:                to.Ptr(armcontainerservice.AgentPoolTypeFlexNodes),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcontainerservice.AgentPoolsClientCreateOrUpdateResponse{
	// 	AgentPool: armcontainerservice.AgentPool{
	// 		Name: to.Ptr("flexnode1"),
	// 		Type: to.Ptr("Microsoft.ContainerService/managedClusters/agentPools"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ContainerService/managedClusters/clustername1/agentPools/flexnode1"),
	// 		Properties: &armcontainerservice.ManagedClusterAgentPoolProfileProperties{
	// 			Mode: to.Ptr(armcontainerservice.AgentPoolModeUser),
	// 			OrchestratorVersion: to.Ptr("1.32"),
	// 			ProvisioningState: to.Ptr("Succeeded"),
	// 			Type: to.Ptr(armcontainerservice.AgentPoolTypeFlexNodes),
	// 		},
	// 	},
	// }
}
