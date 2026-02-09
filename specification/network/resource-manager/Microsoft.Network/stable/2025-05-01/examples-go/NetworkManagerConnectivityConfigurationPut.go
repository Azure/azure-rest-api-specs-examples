package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v9"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/72410da64f6e945db1e1f1af220e077ba5bdb857/specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/NetworkManagerConnectivityConfigurationPut.json
func ExampleConnectivityConfigurationsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewConnectivityConfigurationsClient().CreateOrUpdate(ctx, "myResourceGroup", "testNetworkManager", "myTestConnectivityConfig", armnetwork.ConnectivityConfiguration{
		Properties: &armnetwork.ConnectivityConfigurationProperties{
			Description: to.Ptr("Sample Configuration"),
			AppliesToGroups: []*armnetwork.ConnectivityGroupItem{
				{
					GroupConnectivity: to.Ptr(armnetwork.GroupConnectivityNone),
					IsGlobal:          to.Ptr(armnetwork.IsGlobalFalse),
					NetworkGroupID:    to.Ptr("subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Network/networkManagers/testNetworkManager/networkGroups/group1"),
					UseHubGateway:     to.Ptr(armnetwork.UseHubGatewayTrue),
				}},
			ConnectivityCapabilities: &armnetwork.ConnectivityConfigurationPropertiesConnectivityCapabilities{
				ConnectedGroupAddressOverlap:        to.Ptr(armnetwork.ConnectedGroupAddressOverlapAllowed),
				ConnectedGroupPrivateEndpointsScale: to.Ptr(armnetwork.ConnectedGroupPrivateEndpointsScaleStandard),
				PeeringEnforcement:                  to.Ptr(armnetwork.PeeringEnforcementUnenforced),
			},
			ConnectivityTopology:  to.Ptr(armnetwork.ConnectivityTopologyHubAndSpoke),
			DeleteExistingPeering: to.Ptr(armnetwork.DeleteExistingPeeringTrue),
			Hubs: []*armnetwork.Hub{
				{
					ResourceID:   to.Ptr("subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/myTestConnectivityConfig"),
					ResourceType: to.Ptr("Microsoft.Network/virtualNetworks"),
				}},
			IsGlobal: to.Ptr(armnetwork.IsGlobalTrue),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ConnectivityConfiguration = armnetwork.ConnectivityConfiguration{
	// 	Name: to.Ptr("myTestConnectivityConfig"),
	// 	Type: to.Ptr("Microsoft.Network/networkManagers/connectivityConfigurations"),
	// 	ID: to.Ptr("subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Network/networkManagers/testNetworkManager/connectivityConfigurations/myTestConnectivityConfig"),
	// 	Properties: &armnetwork.ConnectivityConfigurationProperties{
	// 		Description: to.Ptr("Sample Connectivity Configuration"),
	// 		AppliesToGroups: []*armnetwork.ConnectivityGroupItem{
	// 			{
	// 				GroupConnectivity: to.Ptr(armnetwork.GroupConnectivityNone),
	// 				IsGlobal: to.Ptr(armnetwork.IsGlobalFalse),
	// 				NetworkGroupID: to.Ptr("subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Network/networkManagers/testNetworkManager/groups/group1"),
	// 				UseHubGateway: to.Ptr(armnetwork.UseHubGatewayTrue),
	// 		}},
	// 		ConnectivityCapabilities: &armnetwork.ConnectivityConfigurationPropertiesConnectivityCapabilities{
	// 			ConnectedGroupAddressOverlap: to.Ptr(armnetwork.ConnectedGroupAddressOverlapAllowed),
	// 			ConnectedGroupPrivateEndpointsScale: to.Ptr(armnetwork.ConnectedGroupPrivateEndpointsScaleStandard),
	// 			PeeringEnforcement: to.Ptr(armnetwork.PeeringEnforcementUnenforced),
	// 		},
	// 		ConnectivityTopology: to.Ptr(armnetwork.ConnectivityTopologyHubAndSpoke),
	// 		DeleteExistingPeering: to.Ptr(armnetwork.DeleteExistingPeeringTrue),
	// 		Hubs: []*armnetwork.Hub{
	// 			{
	// 				ResourceID: to.Ptr("subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/myTestConnectivityConfig"),
	// 				ResourceType: to.Ptr("Microsoft.Network/virtualNetworks"),
	// 		}},
	// 		IsGlobal: to.Ptr(armnetwork.IsGlobalTrue),
	// 		ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 		ResourceGUID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 	},
	// 	SystemData: &armnetwork.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-11T18:52:27.000Z"); return t}()),
	// 		CreatedBy: to.Ptr("b69a9388-9488-4534-b470-7ec6d41beef5"),
	// 		CreatedByType: to.Ptr(armnetwork.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-11T18:52:27.000Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("b69a9388-9488-4534-b470-7ec6d41beef5"),
	// 		LastModifiedByType: to.Ptr(armnetwork.CreatedByTypeUser),
	// 	},
	// }
}
