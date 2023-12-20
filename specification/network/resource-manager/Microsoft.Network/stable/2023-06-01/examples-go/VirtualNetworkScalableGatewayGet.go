package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/639ecfad68419328658bd4cfe7094af4ce472be2/specification/network/resource-manager/Microsoft.Network/stable/2023-06-01/examples/VirtualNetworkScalableGatewayGet.json
func ExampleVirtualNetworkGatewaysClient_Get_getVirtualNetworkScalableGateway() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
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
	// res.VirtualNetworkGateway = armnetwork.VirtualNetworkGateway{
	// 	Name: to.Ptr("ergw"),
	// 	Type: to.Ptr("Microsoft.Network/virtualNetworkGateways"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkGateways/ergw"),
	// 	Location: to.Ptr("centralus"),
	// 	Etag: to.Ptr("W/\"00000000-0000-0000-0000-000000000000\""),
	// 	Properties: &armnetwork.VirtualNetworkGatewayPropertiesFormat{
	// 		Active: to.Ptr(false),
	// 		AllowRemoteVnetTraffic: to.Ptr(false),
	// 		AllowVirtualWanTraffic: to.Ptr(false),
	// 		AutoScaleConfiguration: &armnetwork.VirtualNetworkGatewayAutoScaleConfiguration{
	// 			Bounds: &armnetwork.VirtualNetworkGatewayAutoScaleBounds{
	// 				Max: to.Ptr[int32](3),
	// 				Min: to.Ptr[int32](2),
	// 			},
	// 		},
	// 		DisableIPSecReplayProtection: to.Ptr(false),
	// 		EnableBgp: to.Ptr(false),
	// 		EnableBgpRouteTranslationForNat: to.Ptr(false),
	// 		EnablePrivateIPAddress: to.Ptr(false),
	// 		GatewayType: to.Ptr(armnetwork.VirtualNetworkGatewayTypeExpressRoute),
	// 		IPConfigurations: []*armnetwork.VirtualNetworkGatewayIPConfiguration{
	// 			{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkGateways/ergw/ipConfigurations/default"),
	// 				Name: to.Ptr("gwipconfig1"),
	// 				Etag: to.Ptr("W/\"00000000-0000-0000-0000-000000000000\""),
	// 				Properties: &armnetwork.VirtualNetworkGatewayIPConfigurationPropertiesFormat{
	// 					PrivateIPAllocationMethod: to.Ptr(armnetwork.IPAllocationMethodStatic),
	// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 					PublicIPAddress: &armnetwork.SubResource{
	// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/publicIPAddresses/gwpip"),
	// 					},
	// 					Subnet: &armnetwork.SubResource{
	// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnet1/subnets/GatewaySubnet"),
	// 					},
	// 				},
	// 		}},
	// 		NatRules: []*armnetwork.VirtualNetworkGatewayNatRule{
	// 			{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkGateways/ergw/natRules/natRule1"),
	// 				Name: to.Ptr("natRule1"),
	// 				Etag: to.Ptr("W/\"00ae2b69-88e7-4b3a-b66a-cfa2244e0801\""),
	// 				Properties: &armnetwork.VirtualNetworkGatewayNatRuleProperties{
	// 					Type: to.Ptr(armnetwork.VPNNatRuleTypeStatic),
	// 					ExternalMappings: []*armnetwork.VPNNatRuleMapping{
	// 						{
	// 							AddressSpace: to.Ptr("50.0.0.0/24"),
	// 					}},
	// 					InternalMappings: []*armnetwork.VPNNatRuleMapping{
	// 						{
	// 							AddressSpace: to.Ptr("10.10.0.0/24"),
	// 					}},
	// 					Mode: to.Ptr(armnetwork.VPNNatRuleModeEgressSnat),
	// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 				},
	// 			},
	// 			{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkGateways/ergw/natRules/natRule2"),
	// 				Name: to.Ptr("natRule2"),
	// 				Etag: to.Ptr("W/\"00ae2b69-88e7-4b3a-b66a-cfa2244e0801\""),
	// 				Properties: &armnetwork.VirtualNetworkGatewayNatRuleProperties{
	// 					Type: to.Ptr(armnetwork.VPNNatRuleTypeStatic),
	// 					ExternalMappings: []*armnetwork.VPNNatRuleMapping{
	// 						{
	// 							AddressSpace: to.Ptr("30.0.0.0/24"),
	// 					}},
	// 					InternalMappings: []*armnetwork.VPNNatRuleMapping{
	// 						{
	// 							AddressSpace: to.Ptr("20.10.0.0/24"),
	// 					}},
	// 					Mode: to.Ptr(armnetwork.VPNNatRuleModeIngressSnat),
	// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 				},
	// 		}},
	// 		ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 		ResourceGUID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 		SKU: &armnetwork.VirtualNetworkGatewaySKU{
	// 			Name: to.Ptr(armnetwork.VirtualNetworkGatewaySKUNameErGwScale),
	// 			Capacity: to.Ptr[int32](2),
	// 			Tier: to.Ptr(armnetwork.VirtualNetworkGatewaySKUTierErGwScale),
	// 		},
	// 		VPNGatewayGeneration: to.Ptr(armnetwork.VPNGatewayGenerationNone),
	// 		VPNType: to.Ptr(armnetwork.VPNTypePolicyBased),
	// 	},
	// }
}
