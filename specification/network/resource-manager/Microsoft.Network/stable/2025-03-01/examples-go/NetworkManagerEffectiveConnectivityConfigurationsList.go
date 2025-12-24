package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v8"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b24c97bfc136b01dd46a1c8ddcecd0bb5a1ab152/specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/NetworkManagerEffectiveConnectivityConfigurationsList.json
func ExampleManagementClient_ListNetworkManagerEffectiveConnectivityConfigurations() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewManagementClient().ListNetworkManagerEffectiveConnectivityConfigurations(ctx, "myResourceGroup", "testVirtualNetwork", armnetwork.QueryRequestOptions{
		SkipToken: to.Ptr("FakeSkipTokenCode"),
	}, &armnetwork.ManagementClientListNetworkManagerEffectiveConnectivityConfigurationsOptions{Top: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ManagerEffectiveConnectivityConfigurationListResult = armnetwork.ManagerEffectiveConnectivityConfigurationListResult{
	// 	SkipToken: to.Ptr("FakeSkipTokenCode"),
	// 	Value: []*armnetwork.EffectiveConnectivityConfiguration{
	// 		{
	// 			ConfigurationGroups: []*armnetwork.ConfigurationGroup{
	// 				{
	// 					ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroup/myResourceGroup/providers/Microsoft.Network/networkManagers/testNetworkManager/networkGroups/group1"),
	// 					Properties: &armnetwork.GroupProperties{
	// 						Description: to.Ptr("A group for all test Virtual Networks"),
	// 						ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 					},
	// 			}},
	// 			ID: to.Ptr("subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Network/networkManagers/testNetworkManager/connectivityConfigurations/myTestConnectivityConfig"),
	// 			Properties: &armnetwork.ConnectivityConfigurationProperties{
	// 				Description: to.Ptr("Sample Configuration"),
	// 				AppliesToGroups: []*armnetwork.ConnectivityGroupItem{
	// 					{
	// 						GroupConnectivity: to.Ptr(armnetwork.GroupConnectivityNone),
	// 						IsGlobal: to.Ptr(armnetwork.IsGlobalFalse),
	// 						NetworkGroupID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Network/networkManagers/testNetworkManager/groups/group1"),
	// 						UseHubGateway: to.Ptr(armnetwork.UseHubGatewayTrue),
	// 				}},
	// 				ConnectivityCapabilities: &armnetwork.ConnectivityConfigurationPropertiesConnectivityCapabilities{
	// 					ConnectedGroupAddressOverlap: to.Ptr(armnetwork.ConnectedGroupAddressOverlapAllowed),
	// 					ConnectedGroupPrivateEndpointsScale: to.Ptr(armnetwork.ConnectedGroupPrivateEndpointsScaleStandard),
	// 					PeeringEnforcement: to.Ptr(armnetwork.PeeringEnforcementUnenforced),
	// 				},
	// 				ConnectivityTopology: to.Ptr(armnetwork.ConnectivityTopologyHubAndSpoke),
	// 				DeleteExistingPeering: to.Ptr(armnetwork.DeleteExistingPeeringTrue),
	// 				Hubs: []*armnetwork.Hub{
	// 					{
	// 						ResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/myHubVnet"),
	// 						ResourceType: to.Ptr("Microsoft.Network/virtualNetworks"),
	// 				}},
	// 				IsGlobal: to.Ptr(armnetwork.IsGlobalTrue),
	// 				ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 			},
	// 	}},
	// }
}
