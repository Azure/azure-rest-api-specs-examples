package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v6"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4883fa5dbf6f2c9093fac8ce334547e9dfac68fa/specification/network/resource-manager/Microsoft.Network/stable/2024-03-01/examples/VirtualNetworkGatewayUpdateTags.json
func ExampleVirtualNetworkGatewaysClient_BeginUpdateTags() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualNetworkGatewaysClient().BeginUpdateTags(ctx, "rg1", "vpngw", armnetwork.TagsObject{
		Tags: map[string]*string{
			"tag1": to.Ptr("value1"),
			"tag2": to.Ptr("value2"),
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
	// res.VirtualNetworkGateway = armnetwork.VirtualNetworkGateway{
	// 	Name: to.Ptr("vpngw"),
	// 	Type: to.Ptr("Microsoft.Network/virtualNetworkGateways"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkGateways/vpngw"),
	// 	Location: to.Ptr("westus"),
	// 	Tags: map[string]*string{
	// 		"tag1": to.Ptr("value1"),
	// 		"tag2": to.Ptr("value2"),
	// 	},
	// 	Identity: &armnetwork.ManagedServiceIdentity{
	// 		Type: to.Ptr(armnetwork.ResourceIdentityTypeUserAssigned),
	// 		UserAssignedIdentities: map[string]*armnetwork.Components1Jq1T4ISchemasManagedserviceidentityPropertiesUserassignedidentitiesAdditionalproperties{
	// 			"/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identity1": &armnetwork.Components1Jq1T4ISchemasManagedserviceidentityPropertiesUserassignedidentitiesAdditionalproperties{
	// 			},
	// 		},
	// 	},
	// 	Properties: &armnetwork.VirtualNetworkGatewayPropertiesFormat{
	// 		Active: to.Ptr(false),
	// 		AllowRemoteVnetTraffic: to.Ptr(false),
	// 		AllowVirtualWanTraffic: to.Ptr(false),
	// 		BgpSettings: &armnetwork.BgpSettings{
	// 			Asn: to.Ptr[int64](65515),
	// 			BgpPeeringAddress: to.Ptr("10.0.0.254"),
	// 			BgpPeeringAddresses: []*armnetwork.IPConfigurationBgpPeeringAddress{
	// 				{
	// 					CustomBgpIPAddresses: []*string{
	// 						to.Ptr("169.254.21.10")},
	// 						DefaultBgpIPAddresses: []*string{
	// 							to.Ptr("10.3.1.254")},
	// 							IPConfigurationID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkGateways/vpngw/ipConfigurations/gwipconfig1"),
	// 							TunnelIPAddresses: []*string{
	// 								to.Ptr("52.161.10.135")},
	// 						}},
	// 						PeerWeight: to.Ptr[int32](0),
	// 					},
	// 					CustomRoutes: &armnetwork.AddressSpace{
	// 						AddressPrefixes: []*string{
	// 							to.Ptr("101.168.0.6/32")},
	// 						},
	// 						DisableIPSecReplayProtection: to.Ptr(false),
	// 						EnableBgp: to.Ptr(false),
	// 						EnableBgpRouteTranslationForNat: to.Ptr(false),
	// 						GatewayType: to.Ptr(armnetwork.VirtualNetworkGatewayTypeVPN),
	// 						IPConfigurations: []*armnetwork.VirtualNetworkGatewayIPConfiguration{
	// 							{
	// 								ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkGateways/vpngw/ipConfigurations/default"),
	// 								Name: to.Ptr("default"),
	// 								Properties: &armnetwork.VirtualNetworkGatewayIPConfigurationPropertiesFormat{
	// 									PrivateIPAllocationMethod: to.Ptr(armnetwork.IPAllocationMethodDynamic),
	// 									ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 									PublicIPAddress: &armnetwork.SubResource{
	// 										ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/publicIPAddresses/testpub1"),
	// 									},
	// 									Subnet: &armnetwork.SubResource{
	// 										ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnet2/subnets/GatewaySubnet"),
	// 									},
	// 								},
	// 						}},
	// 						NatRules: []*armnetwork.VirtualNetworkGatewayNatRule{
	// 						},
	// 						ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 						ResourceGUID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 						SKU: &armnetwork.VirtualNetworkGatewaySKU{
	// 							Name: to.Ptr(armnetwork.VirtualNetworkGatewaySKUNameVPNGw1),
	// 							Capacity: to.Ptr[int32](2),
	// 							Tier: to.Ptr(armnetwork.VirtualNetworkGatewaySKUTierVPNGw1),
	// 						},
	// 						VPNGatewayGeneration: to.Ptr(armnetwork.VPNGatewayGenerationNone),
	// 						VPNType: to.Ptr(armnetwork.VPNTypeRouteBased),
	// 					},
	// 				}
}
