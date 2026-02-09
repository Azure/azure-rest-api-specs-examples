package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v9"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/72410da64f6e945db1e1f1af220e077ba5bdb857/specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/PrivateLinkServiceCreate.json
func ExamplePrivateLinkServicesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewPrivateLinkServicesClient().BeginCreateOrUpdate(ctx, "rg1", "testPls", armnetwork.PrivateLinkService{
		Location: to.Ptr("eastus"),
		Properties: &armnetwork.PrivateLinkServiceProperties{
			AutoApproval: &armnetwork.PrivateLinkServicePropertiesAutoApproval{
				Subscriptions: []*string{
					to.Ptr("subscription1"),
					to.Ptr("subscription2")},
			},
			Fqdns: []*string{
				to.Ptr("fqdn1"),
				to.Ptr("fqdn2"),
				to.Ptr("fqdn3")},
			IPConfigurations: []*armnetwork.PrivateLinkServiceIPConfiguration{
				{
					Name: to.Ptr("fe-lb"),
					Properties: &armnetwork.PrivateLinkServiceIPConfigurationProperties{
						PrivateIPAddress:          to.Ptr("10.0.1.4"),
						PrivateIPAddressVersion:   to.Ptr(armnetwork.IPVersionIPv4),
						PrivateIPAllocationMethod: to.Ptr(armnetwork.IPAllocationMethodStatic),
						Subnet: &armnetwork.Subnet{
							ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnetlb/subnets/subnetlb"),
						},
					},
				}},
			LoadBalancerFrontendIPConfigurations: []*armnetwork.FrontendIPConfiguration{
				{
					ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/frontendIPConfigurations/fe-lb"),
				}},
			Visibility: &armnetwork.PrivateLinkServicePropertiesVisibility{
				Subscriptions: []*string{
					to.Ptr("subscription1"),
					to.Ptr("subscription2"),
					to.Ptr("subscription3")},
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
	// res.PrivateLinkService = armnetwork.PrivateLinkService{
	// 	Name: to.Ptr("testPls"),
	// 	ID: to.Ptr("/subscriptions/subId/resourceGroups/rg1/providers/Microsoft.Network/privateLinkServices/testPls"),
	// 	Location: to.Ptr("eastus"),
	// 	Properties: &armnetwork.PrivateLinkServiceProperties{
	// 		Alias: to.Ptr("ContosoService.{guid}.azure.privatelinkservice"),
	// 		AutoApproval: &armnetwork.PrivateLinkServicePropertiesAutoApproval{
	// 			Subscriptions: []*string{
	// 				to.Ptr("subscription1"),
	// 				to.Ptr("subscription2")},
	// 			},
	// 			Fqdns: []*string{
	// 				to.Ptr("fqdn1"),
	// 				to.Ptr("fqdn2"),
	// 				to.Ptr("fqdn3")},
	// 				IPConfigurations: []*armnetwork.PrivateLinkServiceIPConfiguration{
	// 					{
	// 						ID: to.Ptr("/subscriptions/subId/resourceGroups/rg1/providers/Microsoft.Network/privateLinkServices/testPls/ipConfigurations/ipconfig1"),
	// 						Name: to.Ptr("ipconfig1"),
	// 						Properties: &armnetwork.PrivateLinkServiceIPConfigurationProperties{
	// 							PrivateIPAddress: to.Ptr("10.0.1.4"),
	// 							PrivateIPAddressVersion: to.Ptr(armnetwork.IPVersionIPv4),
	// 							PrivateIPAllocationMethod: to.Ptr(armnetwork.IPAllocationMethodStatic),
	// 							Subnet: &armnetwork.Subnet{
	// 								ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnetlb/subnets/subnetlb"),
	// 							},
	// 						},
	// 				}},
	// 				LoadBalancerFrontendIPConfigurations: []*armnetwork.FrontendIPConfiguration{
	// 					{
	// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/frontendIPConfigurations/fe-lb"),
	// 				}},
	// 				NetworkInterfaces: []*armnetwork.Interface{
	// 					{
	// 						ID: to.Ptr("/subscriptions/subId/resourceGroups/rg1/provders/Microsoft.Network/networkInterfaces/testPls.nic.abcd1234"),
	// 				}},
	// 				ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 				Visibility: &armnetwork.PrivateLinkServicePropertiesVisibility{
	// 					Subscriptions: []*string{
	// 						to.Ptr("subscription1"),
	// 						to.Ptr("subscription2"),
	// 						to.Ptr("subscription3")},
	// 					},
	// 				},
	// 			}
}
