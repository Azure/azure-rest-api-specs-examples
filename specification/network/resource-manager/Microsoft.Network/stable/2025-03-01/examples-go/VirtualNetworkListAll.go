package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v8"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b24c97bfc136b01dd46a1c8ddcecd0bb5a1ab152/specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/VirtualNetworkListAll.json
func ExampleVirtualNetworksClient_NewListAllPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewVirtualNetworksClient().NewListAllPager(nil)
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
		// page.VirtualNetworkListResult = armnetwork.VirtualNetworkListResult{
		// 	Value: []*armnetwork.VirtualNetwork{
		// 		{
		// 			Name: to.Ptr("vnet1"),
		// 			Type: to.Ptr("Microsoft.Network/virtualNetworks"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnet1"),
		// 			Location: to.Ptr("westus"),
		// 			Properties: &armnetwork.VirtualNetworkPropertiesFormat{
		// 				AddressSpace: &armnetwork.AddressSpace{
		// 					AddressPrefixes: []*string{
		// 						to.Ptr("10.0.0.0/8")},
		// 					},
		// 					DefaultPublicNatGateway: &armnetwork.SubResource{
		// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/natGateways/test-ngw"),
		// 					},
		// 					DhcpOptions: &armnetwork.DhcpOptions{
		// 						DNSServers: []*string{
		// 						},
		// 					},
		// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 					Subnets: []*armnetwork.Subnet{
		// 						{
		// 							ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnet1/subnets/test-1"),
		// 							Name: to.Ptr("test-1"),
		// 							Properties: &armnetwork.SubnetPropertiesFormat{
		// 								AddressPrefix: to.Ptr("10.0.0.0/24"),
		// 								ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 							},
		// 					}},
		// 					VirtualNetworkPeerings: []*armnetwork.VirtualNetworkPeering{
		// 					},
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("vnet2"),
		// 				Type: to.Ptr("Microsoft.Network/virtualNetworks"),
		// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg2/providers/Microsoft.Network/virtualNetworks/vnet2"),
		// 				Location: to.Ptr("westus"),
		// 				Properties: &armnetwork.VirtualNetworkPropertiesFormat{
		// 					AddressSpace: &armnetwork.AddressSpace{
		// 						AddressPrefixes: []*string{
		// 							to.Ptr("10.0.0.0/16")},
		// 						},
		// 						DefaultPublicNatGateway: &armnetwork.SubResource{
		// 							ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/natGateways/test-ngw"),
		// 						},
		// 						DhcpOptions: &armnetwork.DhcpOptions{
		// 							DNSServers: []*string{
		// 								to.Ptr("8.8.8.8")},
		// 							},
		// 							ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 							Subnets: []*armnetwork.Subnet{
		// 							},
		// 							VirtualNetworkPeerings: []*armnetwork.VirtualNetworkPeering{
		// 							},
		// 						},
		// 				}},
		// 			}
	}
}
