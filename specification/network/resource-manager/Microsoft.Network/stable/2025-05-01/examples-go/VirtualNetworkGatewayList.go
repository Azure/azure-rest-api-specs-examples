package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v9"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/72410da64f6e945db1e1f1af220e077ba5bdb857/specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/VirtualNetworkGatewayList.json
func ExampleVirtualNetworkGatewaysClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewVirtualNetworkGatewaysClient().NewListPager("rg1", nil)
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
		// page.VirtualNetworkGatewayListResult = armnetwork.VirtualNetworkGatewayListResult{
		// 	Value: []*armnetwork.VirtualNetworkGateway{
		// 		{
		// 			Name: to.Ptr("vpngw1"),
		// 			Type: to.Ptr("Microsoft.Network/virtualNetworkGateways"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkGateways/vpngw1"),
		// 			Location: to.Ptr("loc1"),
		// 			Etag: to.Ptr("W/\"00000000-0000-0000-0000-000000000000\""),
		// 			Properties: &armnetwork.VirtualNetworkGatewayPropertiesFormat{
		// 				Active: to.Ptr(false),
		// 				AllowRemoteVnetTraffic: to.Ptr(false),
		// 				AllowVirtualWanTraffic: to.Ptr(false),
		// 				BgpSettings: &armnetwork.BgpSettings{
		// 					Asn: to.Ptr[int64](65515),
		// 					BgpPeeringAddress: to.Ptr("10.0.0.14"),
		// 					PeerWeight: to.Ptr[int32](0),
		// 				},
		// 				CustomRoutes: &armnetwork.AddressSpace{
		// 					AddressPrefixes: []*string{
		// 						to.Ptr("101.168.0.6/32")},
		// 					},
		// 					DisableIPSecReplayProtection: to.Ptr(false),
		// 					EnableBgp: to.Ptr(false),
		// 					EnableBgpRouteTranslationForNat: to.Ptr(false),
		// 					EnablePrivateIPAddress: to.Ptr(false),
		// 					GatewayType: to.Ptr(armnetwork.VirtualNetworkGatewayTypeVPN),
		// 					IPConfigurations: []*armnetwork.VirtualNetworkGatewayIPConfiguration{
		// 						{
		// 							ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkGateways/vpngw1/ipConfigurations/default"),
		// 							Name: to.Ptr("default"),
		// 							Etag: to.Ptr("W/\"00000000-0000-0000-0000-000000000000\""),
		// 							Properties: &armnetwork.VirtualNetworkGatewayIPConfigurationPropertiesFormat{
		// 								PrivateIPAllocationMethod: to.Ptr(armnetwork.IPAllocationMethodDynamic),
		// 								ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 								PublicIPAddress: &armnetwork.SubResource{
		// 									ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/publicIPAddresses/vpngw1-ip"),
		// 								},
		// 								Subnet: &armnetwork.SubResource{
		// 									ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnet1/subnets/GatewaySubnet"),
		// 								},
		// 							},
		// 					}},
		// 					NatRules: []*armnetwork.VirtualNetworkGatewayNatRule{
		// 						{
		// 							ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkGateways/vpngw1/natRules/natRule1"),
		// 							Name: to.Ptr("natRule1"),
		// 							Etag: to.Ptr("W/\"00ae2b69-88e7-4b3a-b66a-cfa2244e0801\""),
		// 							Properties: &armnetwork.VirtualNetworkGatewayNatRuleProperties{
		// 								Type: to.Ptr(armnetwork.VPNNatRuleTypeStatic),
		// 								ExternalMappings: []*armnetwork.VPNNatRuleMapping{
		// 									{
		// 										AddressSpace: to.Ptr("50.0.0.0/24"),
		// 								}},
		// 								InternalMappings: []*armnetwork.VPNNatRuleMapping{
		// 									{
		// 										AddressSpace: to.Ptr("10.10.0.0/24"),
		// 								}},
		// 								Mode: to.Ptr(armnetwork.VPNNatRuleModeEgressSnat),
		// 								ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 							},
		// 						},
		// 						{
		// 							ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkGateways/vpngw1/natRules/natRule2"),
		// 							Name: to.Ptr("natRule2"),
		// 							Etag: to.Ptr("W/\"00ae2b69-88e7-4b3a-b66a-cfa2244e0801\""),
		// 							Properties: &armnetwork.VirtualNetworkGatewayNatRuleProperties{
		// 								Type: to.Ptr(armnetwork.VPNNatRuleTypeStatic),
		// 								ExternalMappings: []*armnetwork.VPNNatRuleMapping{
		// 									{
		// 										AddressSpace: to.Ptr("30.0.0.0/24"),
		// 								}},
		// 								InternalMappings: []*armnetwork.VPNNatRuleMapping{
		// 									{
		// 										AddressSpace: to.Ptr("20.10.0.0/24"),
		// 								}},
		// 								Mode: to.Ptr(armnetwork.VPNNatRuleModeIngressSnat),
		// 								ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 							},
		// 					}},
		// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 					ResourceGUID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 					SKU: &armnetwork.VirtualNetworkGatewaySKU{
		// 						Name: to.Ptr(armnetwork.VirtualNetworkGatewaySKUNameVPNGw1),
		// 						Capacity: to.Ptr[int32](2),
		// 						Tier: to.Ptr(armnetwork.VirtualNetworkGatewaySKUTierVPNGw1),
		// 					},
		// 					VirtualNetworkGatewayMigrationStatus: &armnetwork.VirtualNetworkGatewayMigrationStatus{
		// 						ErrorMessage: to.Ptr(""),
		// 						Phase: to.Ptr(armnetwork.VirtualNetworkGatewayMigrationPhase("")),
		// 						State: to.Ptr(armnetwork.VirtualNetworkGatewayMigrationState("")),
		// 					},
		// 					VPNClientConfiguration: &armnetwork.VPNClientConfiguration{
		// 						VPNClientProtocols: []*armnetwork.VPNClientProtocol{
		// 						},
		// 						VPNClientRevokedCertificates: []*armnetwork.VPNClientRevokedCertificate{
		// 						},
		// 						VPNClientRootCertificates: []*armnetwork.VPNClientRootCertificate{
		// 						},
		// 					},
		// 					VPNGatewayGeneration: to.Ptr(armnetwork.VPNGatewayGenerationNone),
		// 					VPNType: to.Ptr(armnetwork.VPNTypeRouteBased),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("vpngw2"),
		// 				Type: to.Ptr("Microsoft.Network/virtualNetworkGateways"),
		// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkGateways/vpngw2"),
		// 				Location: to.Ptr("loc2"),
		// 				Etag: to.Ptr("W/\"00000000-0000-0000-0000-000000000000\""),
		// 				Properties: &armnetwork.VirtualNetworkGatewayPropertiesFormat{
		// 					Active: to.Ptr(false),
		// 					AllowRemoteVnetTraffic: to.Ptr(false),
		// 					AllowVirtualWanTraffic: to.Ptr(false),
		// 					BgpSettings: &armnetwork.BgpSettings{
		// 						Asn: to.Ptr[int64](65515),
		// 						BgpPeeringAddress: to.Ptr("10.1.0.46"),
		// 						PeerWeight: to.Ptr[int32](0),
		// 					},
		// 					CustomRoutes: &armnetwork.AddressSpace{
		// 						AddressPrefixes: []*string{
		// 							to.Ptr("101.168.0.6/32")},
		// 						},
		// 						DisableIPSecReplayProtection: to.Ptr(false),
		// 						EnableBgp: to.Ptr(false),
		// 						EnableBgpRouteTranslationForNat: to.Ptr(false),
		// 						EnablePrivateIPAddress: to.Ptr(true),
		// 						GatewayType: to.Ptr(armnetwork.VirtualNetworkGatewayTypeVPN),
		// 						IPConfigurations: []*armnetwork.VirtualNetworkGatewayIPConfiguration{
		// 							{
		// 								ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkGateways/vpngw2/ipConfigurations/default"),
		// 								Name: to.Ptr("default"),
		// 								Etag: to.Ptr("W/\"00000000-0000-0000-0000-000000000000\""),
		// 								Properties: &armnetwork.VirtualNetworkGatewayIPConfigurationPropertiesFormat{
		// 									PrivateIPAddress: to.Ptr("10.1.0.7"),
		// 									PrivateIPAllocationMethod: to.Ptr(armnetwork.IPAllocationMethodDynamic),
		// 									ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 									PublicIPAddress: &armnetwork.SubResource{
		// 										ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/publicIPAddresses/vpngw2-ip"),
		// 									},
		// 									Subnet: &armnetwork.SubResource{
		// 										ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnet2/subnets/GatewaySubnet"),
		// 									},
		// 								},
		// 						}},
		// 						NatRules: []*armnetwork.VirtualNetworkGatewayNatRule{
		// 							{
		// 								ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkGateways/vpngw2/natRules/natRule1"),
		// 								Name: to.Ptr("natRule1"),
		// 								Etag: to.Ptr("W/\"00ae2b69-88e7-4b3a-b66a-cfa2244e0801\""),
		// 								Properties: &armnetwork.VirtualNetworkGatewayNatRuleProperties{
		// 									Type: to.Ptr(armnetwork.VPNNatRuleTypeStatic),
		// 									ExternalMappings: []*armnetwork.VPNNatRuleMapping{
		// 										{
		// 											AddressSpace: to.Ptr("50.0.0.0/24"),
		// 									}},
		// 									InternalMappings: []*armnetwork.VPNNatRuleMapping{
		// 										{
		// 											AddressSpace: to.Ptr("10.10.0.0/24"),
		// 									}},
		// 									Mode: to.Ptr(armnetwork.VPNNatRuleModeEgressSnat),
		// 									ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 								},
		// 							},
		// 							{
		// 								ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkGateways/vpngw2/natRules/natRule2"),
		// 								Name: to.Ptr("natRule2"),
		// 								Etag: to.Ptr("W/\"00ae2b69-88e7-4b3a-b66a-cfa2244e0801\""),
		// 								Properties: &armnetwork.VirtualNetworkGatewayNatRuleProperties{
		// 									Type: to.Ptr(armnetwork.VPNNatRuleTypeStatic),
		// 									ExternalMappings: []*armnetwork.VPNNatRuleMapping{
		// 										{
		// 											AddressSpace: to.Ptr("30.0.0.0/24"),
		// 									}},
		// 									InternalMappings: []*armnetwork.VPNNatRuleMapping{
		// 										{
		// 											AddressSpace: to.Ptr("20.10.0.0/24"),
		// 									}},
		// 									Mode: to.Ptr(armnetwork.VPNNatRuleModeIngressSnat),
		// 									ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 								},
		// 						}},
		// 						ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 						ResourceGUID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 						SKU: &armnetwork.VirtualNetworkGatewaySKU{
		// 							Name: to.Ptr(armnetwork.VirtualNetworkGatewaySKUNameVPNGw1),
		// 							Capacity: to.Ptr[int32](2),
		// 							Tier: to.Ptr(armnetwork.VirtualNetworkGatewaySKUTierVPNGw1),
		// 						},
		// 						VirtualNetworkGatewayMigrationStatus: &armnetwork.VirtualNetworkGatewayMigrationStatus{
		// 							ErrorMessage: to.Ptr(""),
		// 							Phase: to.Ptr(armnetwork.VirtualNetworkGatewayMigrationPhase("")),
		// 							State: to.Ptr(armnetwork.VirtualNetworkGatewayMigrationState("")),
		// 						},
		// 						VPNClientConfiguration: &armnetwork.VPNClientConfiguration{
		// 							RadiusServers: []*armnetwork.RadiusServer{
		// 								{
		// 									RadiusServerAddress: to.Ptr("10.2.0.0"),
		// 									RadiusServerScore: to.Ptr[int64](20),
		// 							}},
		// 							VPNClientProtocols: []*armnetwork.VPNClientProtocol{
		// 								to.Ptr(armnetwork.VPNClientProtocolOpenVPN)},
		// 								VPNClientRevokedCertificates: []*armnetwork.VPNClientRevokedCertificate{
		// 								},
		// 								VPNClientRootCertificates: []*armnetwork.VPNClientRootCertificate{
		// 								},
		// 							},
		// 							VPNGatewayGeneration: to.Ptr(armnetwork.VPNGatewayGenerationNone),
		// 							VPNType: to.Ptr(armnetwork.VPNTypeRouteBased),
		// 						},
		// 				}},
		// 			}
	}
}
