package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v10"
)

// Generated from example definition: 2025-07-01/VirtualNetworkScalableGatewayGet.json
func ExampleVirtualNetworkGatewaysClient_Get_getVirtualNetworkScalableGateway() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVirtualNetworkGatewaysClient().Get(ctx, "rg1", "ergw", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armnetwork.VirtualNetworkGatewaysClientGetResponse{
	// 	VirtualNetworkGateway: armnetwork.VirtualNetworkGateway{
	// 		Name: to.Ptr("ergw"),
	// 		Type: to.Ptr("Microsoft.Network/virtualNetworkGateways"),
	// 		Etag: to.Ptr("W/\"00000000-0000-0000-0000-000000000000\""),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkGateways/ergw"),
	// 		Location: to.Ptr("centralus"),
	// 		Properties: &armnetwork.VirtualNetworkGatewayPropertiesFormat{
	// 			Active: to.Ptr(false),
	// 			AdminState: to.Ptr(armnetwork.AdminStateEnabled),
	// 			AllowRemoteVnetTraffic: to.Ptr(false),
	// 			AllowVirtualWanTraffic: to.Ptr(false),
	// 			AutoScaleConfiguration: &armnetwork.VirtualNetworkGatewayAutoScaleConfiguration{
	// 				Bounds: &armnetwork.VirtualNetworkGatewayAutoScaleBounds{
	// 					Max: to.Ptr[int32](3),
	// 					Min: to.Ptr[int32](2),
	// 				},
	// 			},
	// 			DisableIPSecReplayProtection: to.Ptr(false),
	// 			EnableBgp: to.Ptr(false),
	// 			EnableBgpRouteTranslationForNat: to.Ptr(false),
	// 			GatewayType: to.Ptr(armnetwork.VirtualNetworkGatewayTypeExpressRoute),
	// 			IPConfigurations: []*armnetwork.VirtualNetworkGatewayIPConfiguration{
	// 				{
	// 					Name: to.Ptr("gwipconfig1"),
	// 					Etag: to.Ptr("W/\"00000000-0000-0000-0000-000000000000\""),
	// 					ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkGateways/ergw/ipConfigurations/default"),
	// 					Properties: &armnetwork.VirtualNetworkGatewayIPConfigurationPropertiesFormat{
	// 						PrivateIPAllocationMethod: to.Ptr(armnetwork.IPAllocationMethodDynamic),
	// 						ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 						PublicIPAddress: &armnetwork.SubResource{
	// 							ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/publicIPAddresses/gwpip"),
	// 						},
	// 						Subnet: &armnetwork.SubResource{
	// 							ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnet1/subnets/GatewaySubnet"),
	// 						},
	// 					},
	// 				},
	// 			},
	// 			NatRules: []*armnetwork.VirtualNetworkGatewayNatRule{
	// 			},
	// 			ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 			ResourceGUID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 			SKU: &armnetwork.VirtualNetworkGatewaySKU{
	// 				Name: to.Ptr(armnetwork.VirtualNetworkGatewaySKUNameErGwScale),
	// 				Tier: to.Ptr(armnetwork.VirtualNetworkGatewaySKUTierErGwScale),
	// 			},
	// 			VirtualNetworkGatewayPolicyGroups: []*armnetwork.VirtualNetworkGatewayPolicyGroup{
	// 			},
	// 			VPNType: to.Ptr(armnetwork.VPNTypePolicyBased),
	// 		},
	// 	},
	// }
}
