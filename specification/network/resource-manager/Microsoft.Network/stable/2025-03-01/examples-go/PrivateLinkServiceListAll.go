package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v8"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b24c97bfc136b01dd46a1c8ddcecd0bb5a1ab152/specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/PrivateLinkServiceListAll.json
func ExamplePrivateLinkServicesClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPrivateLinkServicesClient().NewListBySubscriptionPager(nil)
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
		// page.PrivateLinkServiceListResult = armnetwork.PrivateLinkServiceListResult{
		// 	Value: []*armnetwork.PrivateLinkService{
		// 		{
		// 			Name: to.Ptr("testPls1"),
		// 			Type: to.Ptr("Microsoft.Network/privateLinkServices"),
		// 			ID: to.Ptr("/subscriptions/subId/resourceGroups/rg1/providers/Microsoft.Network/privateLinkServices/testPls1"),
		// 			Location: to.Ptr("eastus"),
		// 			Properties: &armnetwork.PrivateLinkServiceProperties{
		// 				Alias: to.Ptr("ContosoService.{guid}.azure.privatelinkservice"),
		// 				AutoApproval: &armnetwork.PrivateLinkServicePropertiesAutoApproval{
		// 					Subscriptions: []*string{
		// 						to.Ptr("subscription1"),
		// 						to.Ptr("subscription2")},
		// 					},
		// 					Fqdns: []*string{
		// 						to.Ptr("fqdn1"),
		// 						to.Ptr("fqdn2"),
		// 						to.Ptr("fqdn3")},
		// 						IPConfigurations: []*armnetwork.PrivateLinkServiceIPConfiguration{
		// 							{
		// 								ID: to.Ptr("/subscriptions/subId/resourceGroups/rg1/providers/Microsoft.Network/privateLinkServices/testPls1/ipConfigurations/ipconfig1"),
		// 								Name: to.Ptr("ipconfig1"),
		// 								Properties: &armnetwork.PrivateLinkServiceIPConfigurationProperties{
		// 									PrivateIPAddress: to.Ptr("10.0.1.4"),
		// 									PrivateIPAddressVersion: to.Ptr(armnetwork.IPVersionIPv4),
		// 									PrivateIPAllocationMethod: to.Ptr(armnetwork.IPAllocationMethodStatic),
		// 									Subnet: &armnetwork.Subnet{
		// 										ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnetlb/subnets/subnetlb1"),
		// 									},
		// 								},
		// 						}},
		// 						LoadBalancerFrontendIPConfigurations: []*armnetwork.FrontendIPConfiguration{
		// 							{
		// 								ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/frontendIPConfigurations/fe-lb1"),
		// 						}},
		// 						NetworkInterfaces: []*armnetwork.Interface{
		// 							{
		// 								ID: to.Ptr("/subscriptions/subId/resourceGroups/rg1/provders/Microsoft.Network/networkInterfaces/testPls1.nic.abcd1234"),
		// 						}},
		// 						PrivateEndpointConnections: []*armnetwork.PrivateEndpointConnection{
		// 							{
		// 								ID: to.Ptr("/subscriptions/subId/resourceGroups/rg1/providers/Microsoft.Network/privateLinkServices/testPls1/privateEndpointConnections/pec1"),
		// 								Name: to.Ptr("pec1"),
		// 								Properties: &armnetwork.PrivateEndpointConnectionProperties{
		// 									PrivateEndpoint: &armnetwork.PrivateEndpoint{
		// 										ID: to.Ptr("/subscriptions/subId/resourceGroups/rg1/providers/Microsoft.Network/privateEndpoints/testPe1"),
		// 									},
		// 									PrivateLinkServiceConnectionState: &armnetwork.PrivateLinkServiceConnectionState{
		// 										Description: to.Ptr("approved it for some reason."),
		// 										Status: to.Ptr("Approved"),
		// 									},
		// 								},
		// 						}},
		// 						ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 						Visibility: &armnetwork.PrivateLinkServicePropertiesVisibility{
		// 							Subscriptions: []*string{
		// 								to.Ptr("subscription1"),
		// 								to.Ptr("subscription2"),
		// 								to.Ptr("subscription3")},
		// 							},
		// 						},
		// 					},
		// 					{
		// 						Name: to.Ptr("testPls2"),
		// 						Type: to.Ptr("Microsoft.Network/privateLinkServices"),
		// 						ID: to.Ptr("/subscriptions/subId/resourceGroups/rg2/providers/Microsoft.Network/privateLinkServices/testPls2"),
		// 						Location: to.Ptr("eastus"),
		// 						Properties: &armnetwork.PrivateLinkServiceProperties{
		// 							Alias: to.Ptr("ContosoService.{guid}.azure.privatelinkservice"),
		// 							AutoApproval: &armnetwork.PrivateLinkServicePropertiesAutoApproval{
		// 								Subscriptions: []*string{
		// 									to.Ptr("subscription1"),
		// 									to.Ptr("subscription2")},
		// 								},
		// 								Fqdns: []*string{
		// 									to.Ptr("fqdn1"),
		// 									to.Ptr("fqdn2")},
		// 									IPConfigurations: []*armnetwork.PrivateLinkServiceIPConfiguration{
		// 										{
		// 											ID: to.Ptr("/subscriptions/subId/resourceGroups/rg2/providers/Microsoft.Network/privateLinkServices/testPls2/ipConfigurations/ipconfig2"),
		// 											Name: to.Ptr("ipconfig2"),
		// 											Properties: &armnetwork.PrivateLinkServiceIPConfigurationProperties{
		// 												PrivateIPAddress: to.Ptr("10.0.1.5"),
		// 												PrivateIPAddressVersion: to.Ptr(armnetwork.IPVersionIPv4),
		// 												PrivateIPAllocationMethod: to.Ptr(armnetwork.IPAllocationMethodStatic),
		// 												Subnet: &armnetwork.Subnet{
		// 													ID: to.Ptr("/subscriptions/subid/resourceGroups/rg2/providers/Microsoft.Network/virtualNetworks/vnetlb/subnets/subnetlb2"),
		// 												},
		// 											},
		// 									}},
		// 									LoadBalancerFrontendIPConfigurations: []*armnetwork.FrontendIPConfiguration{
		// 										{
		// 											ID: to.Ptr("/subscriptions/subid/resourceGroups/rg2/providers/Microsoft.Network/loadBalancers/lb/frontendIPConfigurations/fe-lb2"),
		// 									}},
		// 									NetworkInterfaces: []*armnetwork.Interface{
		// 										{
		// 											ID: to.Ptr("/subscriptions/subId/resourceGroups/rg2/provders/Microsoft.Network/networkInterfaces/testPls2.nic.efgh5678"),
		// 									}},
		// 									PrivateEndpointConnections: []*armnetwork.PrivateEndpointConnection{
		// 										{
		// 											ID: to.Ptr("/subscriptions/subId/resourceGroups/rg2/providers/Microsoft.Network/privateLinkServices/testPls2/privateEndpointConnections/pec2"),
		// 											Name: to.Ptr("pec2"),
		// 											Properties: &armnetwork.PrivateEndpointConnectionProperties{
		// 												PrivateEndpoint: &armnetwork.PrivateEndpoint{
		// 													ID: to.Ptr("/subscriptions/subId/resourceGroups/rg2/providers/Microsoft.Network/privateEndpoints/testPe2"),
		// 												},
		// 												PrivateLinkServiceConnectionState: &armnetwork.PrivateLinkServiceConnectionState{
		// 													Description: to.Ptr("approved it for some reason."),
		// 													Status: to.Ptr("Approved"),
		// 												},
		// 											},
		// 									}},
		// 									ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 									Visibility: &armnetwork.PrivateLinkServicePropertiesVisibility{
		// 										Subscriptions: []*string{
		// 											to.Ptr("subscription1"),
		// 											to.Ptr("subscription2")},
		// 										},
		// 									},
		// 							}},
		// 						}
	}
}
