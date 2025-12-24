package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v8"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b24c97bfc136b01dd46a1c8ddcecd0bb5a1ab152/specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/AzureFirewallPut.json
func ExampleAzureFirewallsClient_BeginCreateOrUpdate_createAzureFirewall() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAzureFirewallsClient().BeginCreateOrUpdate(ctx, "rg1", "azurefirewall", armnetwork.AzureFirewall{
		Location: to.Ptr("West US"),
		Tags: map[string]*string{
			"key1": to.Ptr("value1"),
		},
		Properties: &armnetwork.AzureFirewallPropertiesFormat{
			ApplicationRuleCollections: []*armnetwork.AzureFirewallApplicationRuleCollection{
				{
					ID:   to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/azureFirewalls/azurefirewall/applicationRuleCollections/apprulecoll"),
					Name: to.Ptr("apprulecoll"),
					Properties: &armnetwork.AzureFirewallApplicationRuleCollectionPropertiesFormat{
						Action: &armnetwork.AzureFirewallRCAction{
							Type: to.Ptr(armnetwork.AzureFirewallRCActionTypeDeny),
						},
						Priority: to.Ptr[int32](110),
						Rules: []*armnetwork.AzureFirewallApplicationRule{
							{
								Name:        to.Ptr("rule1"),
								Description: to.Ptr("Deny inbound rule"),
								Protocols: []*armnetwork.AzureFirewallApplicationRuleProtocol{
									{
										Port:         to.Ptr[int32](443),
										ProtocolType: to.Ptr(armnetwork.AzureFirewallApplicationRuleProtocolTypeHTTPS),
									}},
								SourceAddresses: []*string{
									to.Ptr("216.58.216.164"),
									to.Ptr("10.0.0.0/24")},
								TargetFqdns: []*string{
									to.Ptr("www.test.com")},
							}},
					},
				}},
			IPConfigurations: []*armnetwork.AzureFirewallIPConfiguration{
				{
					Name: to.Ptr("azureFirewallIpConfiguration"),
					Properties: &armnetwork.AzureFirewallIPConfigurationPropertiesFormat{
						PublicIPAddress: &armnetwork.SubResource{
							ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/publicIPAddresses/pipName"),
						},
						Subnet: &armnetwork.SubResource{
							ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnet2/subnets/AzureFirewallSubnet"),
						},
					},
				}},
			NatRuleCollections: []*armnetwork.AzureFirewallNatRuleCollection{
				{
					ID:   to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/azureFirewalls/azurefirewall/natRuleCollections/natrulecoll"),
					Name: to.Ptr("natrulecoll"),
					Properties: &armnetwork.AzureFirewallNatRuleCollectionProperties{
						Action: &armnetwork.AzureFirewallNatRCAction{
							Type: to.Ptr(armnetwork.AzureFirewallNatRCActionTypeDnat),
						},
						Priority: to.Ptr[int32](112),
						Rules: []*armnetwork.AzureFirewallNatRule{
							{
								Name:        to.Ptr("DNAT-HTTPS-traffic"),
								Description: to.Ptr("D-NAT all outbound web traffic for inspection"),
								DestinationAddresses: []*string{
									to.Ptr("1.2.3.4")},
								DestinationPorts: []*string{
									to.Ptr("443")},
								Protocols: []*armnetwork.AzureFirewallNetworkRuleProtocol{
									to.Ptr(armnetwork.AzureFirewallNetworkRuleProtocolTCP)},
								SourceAddresses: []*string{
									to.Ptr("*")},
								TranslatedAddress: to.Ptr("1.2.3.5"),
								TranslatedPort:    to.Ptr("8443"),
							},
							{
								Name:        to.Ptr("DNAT-HTTP-traffic-With-FQDN"),
								Description: to.Ptr("D-NAT all inbound web traffic for inspection"),
								DestinationAddresses: []*string{
									to.Ptr("1.2.3.4")},
								DestinationPorts: []*string{
									to.Ptr("80")},
								Protocols: []*armnetwork.AzureFirewallNetworkRuleProtocol{
									to.Ptr(armnetwork.AzureFirewallNetworkRuleProtocolTCP)},
								SourceAddresses: []*string{
									to.Ptr("*")},
								TranslatedFqdn: to.Ptr("internalhttpserver"),
								TranslatedPort: to.Ptr("880"),
							}},
					},
				}},
			NetworkRuleCollections: []*armnetwork.AzureFirewallNetworkRuleCollection{
				{
					ID:   to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/azureFirewalls/azurefirewall/networkRuleCollections/netrulecoll"),
					Name: to.Ptr("netrulecoll"),
					Properties: &armnetwork.AzureFirewallNetworkRuleCollectionPropertiesFormat{
						Action: &armnetwork.AzureFirewallRCAction{
							Type: to.Ptr(armnetwork.AzureFirewallRCActionTypeDeny),
						},
						Priority: to.Ptr[int32](112),
						Rules: []*armnetwork.AzureFirewallNetworkRule{
							{
								Name:        to.Ptr("L4-traffic"),
								Description: to.Ptr("Block traffic based on source IPs and ports"),
								DestinationAddresses: []*string{
									to.Ptr("*")},
								DestinationPorts: []*string{
									to.Ptr("443-444"),
									to.Ptr("8443")},
								Protocols: []*armnetwork.AzureFirewallNetworkRuleProtocol{
									to.Ptr(armnetwork.AzureFirewallNetworkRuleProtocolTCP)},
								SourceAddresses: []*string{
									to.Ptr("192.168.1.1-192.168.1.12"),
									to.Ptr("10.1.4.12-10.1.4.255")},
							},
							{
								Name:        to.Ptr("L4-traffic-with-FQDN"),
								Description: to.Ptr("Block traffic based on source IPs and ports to amazon"),
								DestinationFqdns: []*string{
									to.Ptr("www.amazon.com")},
								DestinationPorts: []*string{
									to.Ptr("443-444"),
									to.Ptr("8443")},
								Protocols: []*armnetwork.AzureFirewallNetworkRuleProtocol{
									to.Ptr(armnetwork.AzureFirewallNetworkRuleProtocolTCP)},
								SourceAddresses: []*string{
									to.Ptr("10.2.4.12-10.2.4.255")},
							}},
					},
				}},
			SKU: &armnetwork.AzureFirewallSKU{
				Name: to.Ptr(armnetwork.AzureFirewallSKUNameAZFWVnet),
				Tier: to.Ptr(armnetwork.AzureFirewallSKUTierStandard),
			},
			ThreatIntelMode: to.Ptr(armnetwork.AzureFirewallThreatIntelModeAlert),
		},
		Zones: []*string{},
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
	// res.AzureFirewall = armnetwork.AzureFirewall{
	// 	Name: to.Ptr("azurefirewall"),
	// 	Type: to.Ptr("Microsoft.Network/azureFirewalls"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/azureFirewalls/azurefirewall"),
	// 	Location: to.Ptr("West US"),
	// 	Tags: map[string]*string{
	// 		"key1": to.Ptr("value1"),
	// 	},
	// 	Etag: to.Ptr("w/\\00000000-0000-0000-0000-000000000000\\"),
	// 	Properties: &armnetwork.AzureFirewallPropertiesFormat{
	// 		AdditionalProperties: map[string]*string{
	// 		},
	// 		ApplicationRuleCollections: []*armnetwork.AzureFirewallApplicationRuleCollection{
	// 			{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/azureFirewalls/azurefirewall/applicationRuleCollections/apprulecoll"),
	// 				Name: to.Ptr("apprulecoll"),
	// 				Properties: &armnetwork.AzureFirewallApplicationRuleCollectionPropertiesFormat{
	// 					Action: &armnetwork.AzureFirewallRCAction{
	// 						Type: to.Ptr(armnetwork.AzureFirewallRCActionTypeDeny),
	// 					},
	// 					Priority: to.Ptr[int32](110),
	// 					Rules: []*armnetwork.AzureFirewallApplicationRule{
	// 						{
	// 							Name: to.Ptr("rule1"),
	// 							Description: to.Ptr("Deny inbound rule"),
	// 							Protocols: []*armnetwork.AzureFirewallApplicationRuleProtocol{
	// 								{
	// 									Port: to.Ptr[int32](443),
	// 									ProtocolType: to.Ptr(armnetwork.AzureFirewallApplicationRuleProtocolTypeHTTPS),
	// 							}},
	// 							SourceAddresses: []*string{
	// 								to.Ptr("216.58.216.164"),
	// 								to.Ptr("10.0.0.0/24")},
	// 								TargetFqdns: []*string{
	// 									to.Ptr("www.test.com")},
	// 							}},
	// 						},
	// 				}},
	// 				IPConfigurations: []*armnetwork.AzureFirewallIPConfiguration{
	// 					{
	// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/azureFirewalls/azurefirewall/ipConfigurations/azureFirewallIpConfiguration"),
	// 						Name: to.Ptr("azureFirewallIpConfiguration"),
	// 						Etag: to.Ptr("w/\\00000000-0000-0000-0000-000000000000\\"),
	// 						Properties: &armnetwork.AzureFirewallIPConfigurationPropertiesFormat{
	// 							PrivateIPAddress: to.Ptr("10.0.0.0"),
	// 							ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 							PublicIPAddress: &armnetwork.SubResource{
	// 								ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/publicIPAddresses/pipName"),
	// 							},
	// 							Subnet: &armnetwork.SubResource{
	// 								ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnet2/subnets/AzureFirewallSubnet"),
	// 							},
	// 						},
	// 				}},
	// 				IPGroups: []*armnetwork.AzureFirewallIPGroups{
	// 				},
	// 				NatRuleCollections: []*armnetwork.AzureFirewallNatRuleCollection{
	// 					{
	// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/azureFirewalls/azurefirewall/natRuleCollections/natrulecoll"),
	// 						Name: to.Ptr("natrulecoll"),
	// 						Properties: &armnetwork.AzureFirewallNatRuleCollectionProperties{
	// 							Action: &armnetwork.AzureFirewallNatRCAction{
	// 								Type: to.Ptr(armnetwork.AzureFirewallNatRCActionTypeDnat),
	// 							},
	// 							Priority: to.Ptr[int32](112),
	// 							Rules: []*armnetwork.AzureFirewallNatRule{
	// 								{
	// 									Name: to.Ptr("DNAT-HTTPS-traffic"),
	// 									Description: to.Ptr("D-NAT all outbound web traffic for inspection"),
	// 									DestinationAddresses: []*string{
	// 										to.Ptr("1.2.3.4")},
	// 										DestinationPorts: []*string{
	// 											to.Ptr("443")},
	// 											Protocols: []*armnetwork.AzureFirewallNetworkRuleProtocol{
	// 												to.Ptr(armnetwork.AzureFirewallNetworkRuleProtocolTCP)},
	// 												SourceAddresses: []*string{
	// 													to.Ptr("*")},
	// 													TranslatedAddress: to.Ptr("1.2.3.5"),
	// 													TranslatedPort: to.Ptr("8443"),
	// 												},
	// 												{
	// 													Name: to.Ptr("DNAT-HTTP-traffic-With-FQDN"),
	// 													Description: to.Ptr("D-NAT all inbound web traffic for inspection"),
	// 													DestinationAddresses: []*string{
	// 														to.Ptr("1.2.3.4")},
	// 														DestinationPorts: []*string{
	// 															to.Ptr("80")},
	// 															Protocols: []*armnetwork.AzureFirewallNetworkRuleProtocol{
	// 																to.Ptr(armnetwork.AzureFirewallNetworkRuleProtocolTCP)},
	// 																SourceAddresses: []*string{
	// 																	to.Ptr("*")},
	// 																	TranslatedFqdn: to.Ptr("internalhttpserver"),
	// 																	TranslatedPort: to.Ptr("880"),
	// 															}},
	// 														},
	// 												}},
	// 												NetworkRuleCollections: []*armnetwork.AzureFirewallNetworkRuleCollection{
	// 													{
	// 														ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/azureFirewalls/azurefirewall/networkRuleCollections/netrulecoll"),
	// 														Name: to.Ptr("netrulecoll"),
	// 														Properties: &armnetwork.AzureFirewallNetworkRuleCollectionPropertiesFormat{
	// 															Action: &armnetwork.AzureFirewallRCAction{
	// 																Type: to.Ptr(armnetwork.AzureFirewallRCActionTypeDeny),
	// 															},
	// 															Priority: to.Ptr[int32](112),
	// 															Rules: []*armnetwork.AzureFirewallNetworkRule{
	// 																{
	// 																	Name: to.Ptr("L4-traffic"),
	// 																	Description: to.Ptr("Block traffic based on source IPs and ports"),
	// 																	DestinationAddresses: []*string{
	// 																		to.Ptr("*")},
	// 																		DestinationPorts: []*string{
	// 																			to.Ptr("443-444"),
	// 																			to.Ptr("8443")},
	// 																			Protocols: []*armnetwork.AzureFirewallNetworkRuleProtocol{
	// 																				to.Ptr(armnetwork.AzureFirewallNetworkRuleProtocolTCP)},
	// 																				SourceAddresses: []*string{
	// 																					to.Ptr("192.168.1.1-192.168.1.12"),
	// 																					to.Ptr("10.1.4.12-10.1.4.255")},
	// 																				},
	// 																				{
	// 																					Name: to.Ptr("L4-traffic-with-FQDN"),
	// 																					Description: to.Ptr("Block traffic based on source IPs and ports to amazon"),
	// 																					DestinationFqdns: []*string{
	// 																						to.Ptr("www.amazon.com")},
	// 																						DestinationPorts: []*string{
	// 																							to.Ptr("443-444"),
	// 																							to.Ptr("8443")},
	// 																							Protocols: []*armnetwork.AzureFirewallNetworkRuleProtocol{
	// 																								to.Ptr(armnetwork.AzureFirewallNetworkRuleProtocolTCP)},
	// 																								SourceAddresses: []*string{
	// 																									to.Ptr("10.2.4.12-10.2.4.255")},
	// 																							}},
	// 																						},
	// 																				}},
	// 																				ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 																				SKU: &armnetwork.AzureFirewallSKU{
	// 																					Name: to.Ptr(armnetwork.AzureFirewallSKUNameAZFWVnet),
	// 																					Tier: to.Ptr(armnetwork.AzureFirewallSKUTierStandard),
	// 																				},
	// 																				ThreatIntelMode: to.Ptr(armnetwork.AzureFirewallThreatIntelModeAlert),
	// 																			},
	// 																			Zones: []*string{
	// 																			},
	// 																		}
}
