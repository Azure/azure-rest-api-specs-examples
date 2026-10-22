package armcontainerservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v9"
)

// Generated from example definition: 2026-08-02-preview/AgentFabricMemberships_CreateOrUpdate.json
func ExampleAgentFabricMembershipsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerservice.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAgentFabricMembershipsClient().BeginCreateOrUpdate(ctx, "rg1", "clustername1", "default", armcontainerservice.AgentFabricMembership{
		ManagedBy: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.NetworkSecurity/agentFabrics/fabric1"),
		Properties: &armcontainerservice.AgentFabricMembershipProperties{
			SKU: to.Ptr(armcontainerservice.AgentFabricMembershipSKUNameStandard),
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
	// res = armcontainerservice.AgentFabricMembershipsClientCreateOrUpdateResponse{
	// 	AgentFabricMembership: armcontainerservice.AgentFabricMembership{
	// 		Name: to.Ptr("default"),
	// 		Type: to.Ptr("Microsoft.ContainerService/managedClusters/agentFabricMemberships"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ContainerService/managedClusters/clustername1/agentFabricMemberships/default"),
	// 		ManagedBy: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.NetworkSecurity/agentFabrics/fabric1"),
	// 		Properties: &armcontainerservice.AgentFabricMembershipProperties{
	// 			ProvisioningState: to.Ptr(armcontainerservice.AgentFabricMembershipProvisioningStateSucceeded),
	// 			SKU: to.Ptr(armcontainerservice.AgentFabricMembershipSKUNameStandard),
	// 		},
	// 	},
	// }
}
