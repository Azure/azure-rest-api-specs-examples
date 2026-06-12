package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v10"
)

// Generated from example definition: 2025-07-01/NetworkManagerConnectivityConfigurationList.json
func ExampleConnectivityConfigurationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewConnectivityConfigurationsClient().NewListPager("myResourceGroup", "testNetworkManager", nil)
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
		// page = armnetwork.ConnectivityConfigurationsClientListResponse{
		// 	ConnectivityConfigurationListResult: armnetwork.ConnectivityConfigurationListResult{
		// 		NextLink: to.Ptr("https://management.azure.com/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Network/managedNetworks/testNetworkManager/connectivityConfigurations?api-version=2025-07-01&$skipToken=10"),
		// 		Value: []*armnetwork.ConnectivityConfiguration{
		// 			{
		// 				Name: to.Ptr("myTestConnectivityConfig"),
		// 				Type: to.Ptr("Microsoft.Network/networkManagers/connectivityConfigurations"),
		// 				ID: to.Ptr("subscription/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Network/networkManagers/testNetworkManager/connectivityConfigurations/myTestConnectivityConfig"),
		// 				Properties: &armnetwork.ConnectivityConfigurationProperties{
		// 					Description: to.Ptr("Sample Configuration"),
		// 					AppliesToGroups: []*armnetwork.ConnectivityGroupItem{
		// 						{
		// 							GroupConnectivity: to.Ptr(armnetwork.GroupConnectivityNone),
		// 							IsGlobal: to.Ptr(armnetwork.IsGlobalFalse),
		// 							NetworkGroupID: to.Ptr("subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Network/networkManagers/testNetworkManager/netwrokGroups/group1"),
		// 							UseHubGateway: to.Ptr(armnetwork.UseHubGatewayTrue),
		// 						},
		// 					},
		// 					ConnectivityCapabilities: &armnetwork.ConnectivityConfigurationPropertiesConnectivityCapabilities{
		// 						ConnectedGroupAddressOverlap: to.Ptr(armnetwork.ConnectedGroupAddressOverlapAllowed),
		// 						ConnectedGroupPrivateEndpointsScale: to.Ptr(armnetwork.ConnectedGroupPrivateEndpointsScaleStandard),
		// 						PeeringEnforcement: to.Ptr(armnetwork.PeeringEnforcementUnenforced),
		// 					},
		// 					ConnectivityTopology: to.Ptr(armnetwork.ConnectivityTopologyHubAndSpoke),
		// 					DeleteExistingPeering: to.Ptr(armnetwork.DeleteExistingPeeringTrue),
		// 					Hubs: []*armnetwork.Hub{
		// 						{
		// 							ResourceID: to.Ptr("subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/myHubVnet"),
		// 							ResourceType: to.Ptr("Microsoft.Network/virtualNetworks"),
		// 						},
		// 					},
		// 					IsGlobal: to.Ptr(armnetwork.IsGlobalTrue),
		// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 					ResourceGUID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 				},
		// 				SystemData: &armnetwork.SystemData{
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-11T18:52:27Z"); return t}()),
		// 					CreatedBy: to.Ptr("b69a9388-9488-4534-b470-7ec6d41beef5"),
		// 					CreatedByType: to.Ptr(armnetwork.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-11T18:52:27Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("b69a9388-9488-4534-b470-7ec6d41beef5"),
		// 					LastModifiedByType: to.Ptr(armnetwork.CreatedByTypeUser),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
