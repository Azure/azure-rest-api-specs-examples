package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v9"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/72410da64f6e945db1e1f1af220e077ba5bdb857/specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/NatGatewayList.json
func ExampleNatGatewaysClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewNatGatewaysClient().NewListPager("rg1", nil)
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
		// page.NatGatewayListResult = armnetwork.NatGatewayListResult{
		// 	Value: []*armnetwork.NatGateway{
		// 		{
		// 			Name: to.Ptr("test-natGateway"),
		// 			Type: to.Ptr("Microsoft.Network/natGateways"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/natGateway/test-natGateway"),
		// 			Location: to.Ptr("westus"),
		// 			Properties: &armnetwork.NatGatewayPropertiesFormat{
		// 				IdleTimeoutInMinutes: to.Ptr[int32](5),
		// 				ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 				PublicIPAddresses: []*armnetwork.SubResource{
		// 					{
		// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/publicIPAddresses/PublicIpAddress1"),
		// 				}},
		// 				PublicIPPrefixes: []*armnetwork.SubResource{
		// 					{
		// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/publicIPPrefixes/PublicIpPrefix1"),
		// 				}},
		// 				Subnets: []*armnetwork.SubResource{
		// 					{
		// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/test-vnet/subnets/subnet1"),
		// 				}},
		// 			},
		// 			SKU: &armnetwork.NatGatewaySKU{
		// 				Name: to.Ptr(armnetwork.NatGatewaySKUNameStandard),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("test-natGateway2"),
		// 			Type: to.Ptr("Microsoft.Network/natGateways"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg2/providers/Microsoft.Network/natGateway/test-natGateway2"),
		// 			Location: to.Ptr("westus"),
		// 			Properties: &armnetwork.NatGatewayPropertiesFormat{
		// 				IdleTimeoutInMinutes: to.Ptr[int32](5),
		// 				ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 				PublicIPAddresses: []*armnetwork.SubResource{
		// 					{
		// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg2/providers/Microsoft.Network/publicIPAddresses/PublicIpAddress1"),
		// 				}},
		// 				PublicIPPrefixes: []*armnetwork.SubResource{
		// 					{
		// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg2/providers/Microsoft.Network/publicIPPrefixes/PublicIpPrefix1"),
		// 				}},
		// 				Subnets: []*armnetwork.SubResource{
		// 					{
		// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg2/providers/Microsoft.Network/virtualNetworks/test-vnet/subnets/subnet1"),
		// 				}},
		// 			},
		// 			SKU: &armnetwork.NatGatewaySKU{
		// 				Name: to.Ptr(armnetwork.NatGatewaySKUNameStandard),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("test-natGateway3"),
		// 			Type: to.Ptr("Microsoft.Network/natGateways"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg2/providers/Microsoft.Network/natGateway/test-natGateway3"),
		// 			Location: to.Ptr("westus"),
		// 			Properties: &armnetwork.NatGatewayPropertiesFormat{
		// 				IdleTimeoutInMinutes: to.Ptr[int32](5),
		// 				ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 				PublicIPAddresses: []*armnetwork.SubResource{
		// 					{
		// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg2/providers/Microsoft.Network/publicIPAddresses/PublicIpAddress1"),
		// 				}},
		// 				PublicIPAddressesV6: []*armnetwork.SubResource{
		// 					{
		// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg2/providers/Microsoft.Network/publicIPAddresses/PublicIpAddress2"),
		// 				}},
		// 				PublicIPPrefixes: []*armnetwork.SubResource{
		// 					{
		// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg2/providers/Microsoft.Network/publicIPPrefixes/PublicIpPrefix1"),
		// 				}},
		// 				PublicIPPrefixesV6: []*armnetwork.SubResource{
		// 					{
		// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg2/providers/Microsoft.Network/publicIPPrefixes/PublicIpPrefix2"),
		// 				}},
		// 				SourceVirtualNetwork: &armnetwork.SubResource{
		// 					ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/test-vnet"),
		// 				},
		// 				Subnets: []*armnetwork.SubResource{
		// 					{
		// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg2/providers/Microsoft.Network/virtualNetworks/test-vnet/subnets/subnet1"),
		// 				}},
		// 			},
		// 			SKU: &armnetwork.NatGatewaySKU{
		// 				Name: to.Ptr(armnetwork.NatGatewaySKUNameStandardV2),
		// 			},
		// 	}},
		// }
	}
}
