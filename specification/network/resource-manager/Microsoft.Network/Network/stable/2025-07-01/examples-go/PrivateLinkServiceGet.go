package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v10"
)

// Generated from example definition: 2025-07-01/PrivateLinkServiceGet.json
func ExamplePrivateLinkServicesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPrivateLinkServicesClient().Get(ctx, "rg1", "testPls", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armnetwork.PrivateLinkServicesClientGetResponse{
	// 	PrivateLinkService: armnetwork.PrivateLinkService{
	// 		Name: to.Ptr("testPls"),
	// 		Type: to.Ptr("Microsoft.Network/privateLinkServices"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/privateLinkServices/testPls"),
	// 		Location: to.Ptr("eastus"),
	// 		Properties: &armnetwork.PrivateLinkServiceProperties{
	// 			Alias: to.Ptr("ContosoService.{guid}.azure.privatelinkservice"),
	// 			AutoApproval: &armnetwork.PrivateLinkServicePropertiesAutoApproval{
	// 				Subscriptions: []*string{
	// 					to.Ptr("subscription1"),
	// 					to.Ptr("subscription2"),
	// 				},
	// 			},
	// 			Fqdns: []*string{
	// 				to.Ptr("fqdn1"),
	// 				to.Ptr("fqdn2"),
	// 				to.Ptr("fqdn3"),
	// 			},
	// 			IPConfigurations: []*armnetwork.PrivateLinkServiceIPConfiguration{
	// 				{
	// 					Name: to.Ptr("ipconfig1"),
	// 					ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/privateLinkServices/testPls/ipConfigurations/ipconfig1"),
	// 					Properties: &armnetwork.PrivateLinkServiceIPConfigurationProperties{
	// 						PrivateIPAddress: to.Ptr("10.0.1.4"),
	// 						PrivateIPAddressVersion: to.Ptr(armnetwork.IPVersionIPv4),
	// 						PrivateIPAllocationMethod: to.Ptr(armnetwork.IPAllocationMethodStatic),
	// 						Subnet: &armnetwork.Subnet{
	// 							ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnetlb/subnets/subnetlb"),
	// 						},
	// 					},
	// 				},
	// 			},
	// 			LoadBalancerFrontendIPConfigurations: []*armnetwork.FrontendIPConfiguration{
	// 				{
	// 					ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/frontendIPConfigurations/fe-lb"),
	// 				},
	// 			},
	// 			NetworkInterfaces: []*armnetwork.Interface{
	// 				{
	// 					ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/provders/Microsoft.Network/networkInterfaces/testPls.nic.abcd1234"),
	// 				},
	// 			},
	// 			PrivateEndpointConnections: []*armnetwork.PrivateEndpointConnection{
	// 				{
	// 					Name: to.Ptr("privateEndpointConnection"),
	// 					ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/privateLinkServices/testPls/privateEndpointConnections/privateEndpointConnection"),
	// 					Properties: &armnetwork.PrivateEndpointConnectionProperties{
	// 						PrivateEndpoint: &armnetwork.PrivateEndpoint{
	// 							ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/privateEndpoints/testPe"),
	// 						},
	// 						PrivateLinkServiceConnectionState: &armnetwork.PrivateLinkServiceConnectionState{
	// 							Description: to.Ptr("approved it for some reason."),
	// 							Status: to.Ptr("Approved"),
	// 						},
	// 					},
	// 				},
	// 			},
	// 			ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 			Visibility: &armnetwork.PrivateLinkServicePropertiesVisibility{
	// 				Subscriptions: []*string{
	// 					to.Ptr("subscription1"),
	// 					to.Ptr("subscription2"),
	// 					to.Ptr("subscription3"),
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
