package armcontainerservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/17b0396c5ce59d801b9eb38fd6c47d149f6db2a2/specification/containerservice/resource-manager/Microsoft.ContainerService/aks/preview/2023-03-02-preview/examples/AgentPools_Stop.json
func ExampleAgentPoolsClient_BeginCreateOrUpdate_stopAgentPool() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAgentPoolsClient().BeginCreateOrUpdate(ctx, "rg1", "clustername1", "agentpool1", armcontainerservice.AgentPool{
		Properties: &armcontainerservice.ManagedClusterAgentPoolProfileProperties{
			PowerState: &armcontainerservice.PowerState{
				Code: to.Ptr(armcontainerservice.CodeStopped),
			},
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
	// res.AgentPool = armcontainerservice.AgentPool{
	// 	Name: to.Ptr("agentpool1"),
	// 	Type: to.Ptr("Microsoft.ContainerService/managedClusters/agentPools"),
	// 	ID: to.Ptr("/subscriptions/subid1/resourcegroups/rg1/providers/Microsoft.ContainerService/managedClusters/clustername1/agentPools/agentpool1"),
	// 	Properties: &armcontainerservice.ManagedClusterAgentPoolProfileProperties{
	// 		Count: to.Ptr[int32](0),
	// 		EnableAutoScaling: to.Ptr(false),
	// 		PowerState: &armcontainerservice.PowerState{
	// 			Code: to.Ptr(armcontainerservice.CodeStopped),
	// 		},
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 	},
	// }
}
