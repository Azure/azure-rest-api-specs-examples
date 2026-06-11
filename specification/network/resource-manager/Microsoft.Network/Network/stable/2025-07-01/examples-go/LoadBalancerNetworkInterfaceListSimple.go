package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v10"
)

// Generated from example definition: 2025-07-01/LoadBalancerNetworkInterfaceListSimple.json
func ExampleLoadBalancerNetworkInterfacesClient_NewListPager_loadBalancerNetworkInterfaceListSimple() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewLoadBalancerNetworkInterfacesClient().NewListPager("testrg", "lb", nil)
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
		// page = armnetwork.LoadBalancerNetworkInterfacesClientListResponse{
		// 	InterfaceListResult: armnetwork.InterfaceListResult{
		// 		Value: []*armnetwork.Interface{
		// 			{
		// 				Name: to.Ptr("mynic"),
		// 				Type: to.Ptr("Microsoft.Network/networkInterfaces"),
		// 				Etag: to.Ptr("W/\\\"00000000-0000-0000-0000-000000000000\\\""),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testrg/providers/Microsoft.Network/networkInterfaces/mynic"),
		// 				Location: to.Ptr("westus"),
		// 				Properties: &armnetwork.InterfacePropertiesFormat{
		// 					DNSSettings: &armnetwork.InterfaceDNSSettings{
		// 						AppliedDNSServers: []*string{
		// 						},
		// 						DNSServers: []*string{
		// 						},
		// 					},
		// 					EnableAcceleratedNetworking: to.Ptr(false),
		// 					EnableIPForwarding: to.Ptr(false),
		// 					IPConfigurations: []*armnetwork.InterfaceIPConfiguration{
		// 						{
		// 							Name: to.Ptr("ipconfig1"),
		// 							Etag: to.Ptr("W/\\\"00000000-0000-0000-0000-000000000000\\\""),
		// 							ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testrg/providers/Microsoft.Network/networkInterfaces/mynic/ipConfigurations/ipconfig1"),
		// 							Properties: &armnetwork.InterfaceIPConfigurationPropertiesFormat{
		// 								LoadBalancerBackendAddressPools: []*armnetwork.BackendAddressPool{
		// 									{
		// 										ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testrg/providers/Microsoft.Network/loadBalancers/lb/backendAddressPools/bepool1"),
		// 									},
		// 								},
		// 								LoadBalancerInboundNatRules: []*armnetwork.InboundNatRule{
		// 									{
		// 										ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testrg/providers/Microsoft.Network/loadBalancers/lb/inboundNatRules/inbound1"),
		// 									},
		// 								},
		// 								PrivateIPAddress: to.Ptr("10.0.1.4"),
		// 								PrivateIPAddressVersion: to.Ptr(armnetwork.IPVersionIPv4),
		// 								PrivateIPAllocationMethod: to.Ptr(armnetwork.IPAllocationMethodDynamic),
		// 								ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 								Subnet: &armnetwork.Subnet{
		// 									ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testrg/providers/Microsoft.Network/virtualNetworks/myVirtualNetwork/subnets/frontendSubnet"),
		// 								},
		// 							},
		// 						},
		// 					},
		// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 					ResourceGUID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
