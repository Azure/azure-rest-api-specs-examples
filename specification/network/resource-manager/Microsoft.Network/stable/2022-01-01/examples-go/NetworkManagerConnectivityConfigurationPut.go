package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/NetworkManagerConnectivityConfigurationPut.json
func ExampleConnectivityConfigurationsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armnetwork.NewConnectivityConfigurationsClient("subscriptionA", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx, "myResourceGroup", "testNetworkManager", "myTestConnectivityConfig", armnetwork.ConnectivityConfiguration{
		Properties: &armnetwork.ConnectivityConfigurationProperties{
			Description: to.Ptr("Sample Configuration"),
			AppliesToGroups: []*armnetwork.ConnectivityGroupItem{
				{
					GroupConnectivity: to.Ptr(armnetwork.GroupConnectivityNone),
					IsGlobal:          to.Ptr(armnetwork.IsGlobalFalse),
					NetworkGroupID:    to.Ptr("subscriptions/subscriptionA/resourceGroups/myResourceGroup/providers/Microsoft.Network/networkManagers/testNetworkManager/networkGroups/group1"),
					UseHubGateway:     to.Ptr(armnetwork.UseHubGatewayTrue),
				}},
			ConnectivityTopology:  to.Ptr(armnetwork.ConnectivityTopologyHubAndSpoke),
			DeleteExistingPeering: to.Ptr(armnetwork.DeleteExistingPeeringTrue),
			Hubs: []*armnetwork.Hub{
				{
					ResourceID:   to.Ptr("subscriptions/subscriptionA/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/myTestConnectivityConfig"),
					ResourceType: to.Ptr("Microsoft.Network/virtualNetworks"),
				}},
			IsGlobal: to.Ptr(armnetwork.IsGlobalTrue),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
