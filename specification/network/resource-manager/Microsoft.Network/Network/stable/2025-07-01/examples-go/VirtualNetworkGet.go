package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v10"
)

// Generated from example definition: 2025-07-01/VirtualNetworkGet.json
func ExampleVirtualNetworksClient_Get_getVirtualNetwork() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVirtualNetworksClient().Get(ctx, "rg1", "test-vnet", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armnetwork.VirtualNetworksClientGetResponse{
	// 	VirtualNetwork: armnetwork.VirtualNetwork{
	// 		Name: to.Ptr("test-vnet"),
	// 		Type: to.Ptr("Microsoft.Network/virtualNetworks"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/test-vnet"),
	// 		Location: to.Ptr("westus"),
	// 		Properties: &armnetwork.VirtualNetworkPropertiesFormat{
	// 			AddressSpace: &armnetwork.AddressSpace{
	// 				AddressPrefixes: []*string{
	// 					to.Ptr("10.0.0.0/16"),
	// 				},
	// 			},
	// 			DefaultPublicNatGateway: &armnetwork.SubResource{
	// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/natGateways/test-ngw"),
	// 			},
	// 			ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 			Subnets: []*armnetwork.Subnet{
	// 				{
	// 					Name: to.Ptr("subnet1"),
	// 					ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/test-vnet/subnets/subnet1"),
	// 					Properties: &armnetwork.SubnetPropertiesFormat{
	// 						AddressPrefix: to.Ptr("10.0.1.0/24"),
	// 						IPConfigurations: []*armnetwork.IPConfiguration{
	// 							{
	// 								ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/frontendIPConfigurations/fe"),
	// 							},
	// 						},
	// 						ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 					},
	// 				},
	// 			},
	// 			VirtualNetworkPeerings: []*armnetwork.VirtualNetworkPeering{
	// 			},
	// 		},
	// 	},
	// }
}
