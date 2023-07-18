package armpanngfw_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/paloaltonetworksngfw/armpanngfw"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/310a0100f5b020c1900c527a6aa70d21992f078a/specification/paloaltonetworks/resource-manager/PaloAltoNetworks.Cloudngfw/stable/2022-08-29/examples/Firewalls_Update_MaximumSet_Gen.json
func ExampleFirewallsClient_Update_firewallsUpdateMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpanngfw.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewFirewallsClient().Update(ctx, "firewall-rg", "firewall1", armpanngfw.FirewallResourceUpdate{
		Identity: &armpanngfw.AzureResourceManagerManagedIdentityProperties{
			Type: to.Ptr(armpanngfw.ManagedIdentityTypeNone),
			UserAssignedIdentities: map[string]*armpanngfw.AzureResourceManagerUserAssignedIdentity{
				"key16": {
					ClientID:    to.Ptr("aaaa"),
					PrincipalID: to.Ptr("aaaaaaaaaaaaaaa"),
				},
			},
		},
		Properties: &armpanngfw.FirewallResourceUpdateProperties{
			AssociatedRulestack: &armpanngfw.RulestackDetails{
				Location:    to.Ptr("eastus"),
				ResourceID:  to.Ptr("aaaaaaaaaa"),
				RulestackID: to.Ptr("aaaaaaaaaaaaaaaa"),
			},
			DNSSettings: &armpanngfw.DNSSettings{
				DNSServers: []*armpanngfw.IPAddress{
					{
						Address:    to.Ptr("20.22.92.111"),
						ResourceID: to.Ptr("/subscriptions/01c7d41f-afaf-464e-8a8b-5c6f9f98cee8/resourceGroups/mj-liftr-integration/providers/Microsoft.Network/publicIPAddresses/mj-liftr-integration-egressNatIp1"),
					}},
				EnableDNSProxy: to.Ptr(armpanngfw.DNSProxyDISABLED),
				EnabledDNSType: to.Ptr(armpanngfw.EnabledDNSTypeCUSTOM),
			},
			FrontEndSettings: []*armpanngfw.FrontendSetting{
				{
					Name: to.Ptr("frontendsetting11"),
					BackendConfiguration: &armpanngfw.EndpointConfiguration{
						Address: &armpanngfw.IPAddress{
							Address:    to.Ptr("20.22.32.136"),
							ResourceID: to.Ptr("/subscriptions/01c7d41f-afaf-464e-8a8b-5c6f9f98cee8/resourceGroups/mj-liftr-integration/providers/Microsoft.Network/publicIPAddresses/mj-liftr-integration-frontendSettingIp2"),
						},
						Port: to.Ptr("80"),
					},
					FrontendConfiguration: &armpanngfw.EndpointConfiguration{
						Address: &armpanngfw.IPAddress{
							Address:    to.Ptr("20.22.91.251"),
							ResourceID: to.Ptr("/subscriptions/01c7d41f-afaf-464e-8a8b-5c6f9f98cee8/resourceGroups/mj-liftr-integration/providers/Microsoft.Network/publicIPAddresses/mj-liftr-integration-frontendSettingIp1"),
						},
						Port: to.Ptr("80"),
					},
					Protocol: to.Ptr(armpanngfw.ProtocolTypeTCP),
				}},
			IsPanoramaManaged: to.Ptr(armpanngfw.BooleanEnumTRUE),
			MarketplaceDetails: &armpanngfw.MarketplaceDetails{
				MarketplaceSubscriptionStatus: to.Ptr(armpanngfw.MarketplaceSubscriptionStatusPendingFulfillmentStart),
				OfferID:                       to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
				PublisherID:                   to.Ptr("aaaa"),
			},
			NetworkProfile: &armpanngfw.NetworkProfile{
				EgressNatIP: []*armpanngfw.IPAddress{
					{
						Address:    to.Ptr("20.22.92.111"),
						ResourceID: to.Ptr("/subscriptions/01c7d41f-afaf-464e-8a8b-5c6f9f98cee8/resourceGroups/mj-liftr-integration/providers/Microsoft.Network/publicIPAddresses/mj-liftr-integration-egressNatIp1"),
					}},
				EnableEgressNat: to.Ptr(armpanngfw.EgressNatENABLED),
				NetworkType:     to.Ptr(armpanngfw.NetworkTypeVNET),
				PublicIPs: []*armpanngfw.IPAddress{
					{
						Address:    to.Ptr("20.22.92.11"),
						ResourceID: to.Ptr("/subscriptions/01c7d41f-afaf-464e-8a8b-5c6f9f98cee8/resourceGroups/mj-liftr-integration/providers/Microsoft.Network/publicIPAddresses/mj-liftr-integration-PublicIp1"),
					}},
				VnetConfiguration: &armpanngfw.VnetConfiguration{
					IPOfTrustSubnetForUdr: &armpanngfw.IPAddress{
						Address:    to.Ptr("10.1.1.0/24"),
						ResourceID: to.Ptr("/subscriptions/2bf4a339-294d-4c25-b0b2-ef649e9f5c27/resourceGroups/os-liftr-integration/providers/Microsoft.Network/virtualNetworks/os-liftr-integration-vnet/subnets/os-liftr-integration-untrust-subnet"),
					},
					TrustSubnet: &armpanngfw.IPAddressSpace{
						AddressSpace: to.Ptr("10.1.1.0/24"),
						ResourceID:   to.Ptr("/subscriptions/2bf4a339-294d-4c25-b0b2-ef649e9f5c27/resourceGroups/os-liftr-integration/providers/Microsoft.Network/virtualNetworks/os-liftr-integration-vnet/subnets/os-liftr-integration-trust-subnet"),
					},
					UnTrustSubnet: &armpanngfw.IPAddressSpace{
						AddressSpace: to.Ptr("10.1.1.0/24"),
						ResourceID:   to.Ptr("/subscriptions/2bf4a339-294d-4c25-b0b2-ef649e9f5c27/resourceGroups/os-liftr-integration/providers/Microsoft.Network/virtualNetworks/os-liftr-integration-vnet/subnets/os-liftr-integration-untrust-subnet"),
					},
					Vnet: &armpanngfw.IPAddressSpace{
						AddressSpace: to.Ptr("10.1.0.0/16"),
						ResourceID:   to.Ptr("/subscriptions/2bf4a339-294d-4c25-b0b2-ef649e9f5c27/resourceGroups/os-liftr-integration/providers/Microsoft.Network/virtualNetworks/os-liftr-integration-vnet"),
					},
				},
				VwanConfiguration: &armpanngfw.VwanConfiguration{
					IPOfTrustSubnetForUdr: &armpanngfw.IPAddress{
						Address:    to.Ptr("10.1.1.0/24"),
						ResourceID: to.Ptr("/subscriptions/2bf4a339-294d-4c25-b0b2-ef649e9f5c27/resourceGroups/os-liftr-integration/providers/Microsoft.Network/virtualNetworks/os-liftr-integration-vnet/subnets/os-liftr-integration-untrust-subnet"),
					},
					NetworkVirtualApplianceID: to.Ptr("2bf4a339-294d-4c25-b0b2-ef649e9f5c12"),
					TrustSubnet: &armpanngfw.IPAddressSpace{
						AddressSpace: to.Ptr("10.1.1.0/24"),
						ResourceID:   to.Ptr("/subscriptions/2bf4a339-294d-4c25-b0b2-ef649e9f5c27/resourceGroups/os-liftr-integration/providers/Microsoft.Network/virtualNetworks/os-liftr-integration-vnet/subnets/os-liftr-integration-trust-subnet"),
					},
					UnTrustSubnet: &armpanngfw.IPAddressSpace{
						AddressSpace: to.Ptr("10.1.1.0/24"),
						ResourceID:   to.Ptr("/subscriptions/2bf4a339-294d-4c25-b0b2-ef649e9f5c27/resourceGroups/os-liftr-integration/providers/Microsoft.Network/virtualNetworks/os-liftr-integration-vnet/subnets/os-liftr-integration-untrust-subnet"),
					},
					VHub: &armpanngfw.IPAddressSpace{
						AddressSpace: to.Ptr("10.1.1.0/24"),
						ResourceID:   to.Ptr("/subscriptions/2bf4a339-294d-4c25-b0b2-ef649e9f5c27/resourceGroups/os-liftr-integration/providers/Microsoft.Network/virtualNetworks/os-liftr-integration-vnet/subnets/os-liftr-integration-untrust-subnet"),
					},
				},
			},
			PanEtag: to.Ptr("2bf4a339-294d-4c25-b0b2-ef649e9f5c12"),
			PanoramaConfig: &armpanngfw.PanoramaConfig{
				ConfigString: to.Ptr("bas64EncodedString"),
			},
			PlanData: &armpanngfw.PlanData{
				BillingCycle: to.Ptr(armpanngfw.BillingCycleWEEKLY),
				PlanID:       to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
				UsageType:    to.Ptr(armpanngfw.UsageTypePAYG),
			},
		},
		Tags: map[string]*string{
			"tagName": to.Ptr("value"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.FirewallResource = armpanngfw.FirewallResource{
	// 	Name: to.Ptr("aaaaaaaaaaaaa"),
	// 	Type: to.Ptr("aaaaaa"),
	// 	ID: to.Ptr("/subscriptions/2bf4a339-294d-4c25-b0b2-ef649e9f5c27/resourceGroups/acctest4136/providers/PaloAltoNetworks.Cloudngfw/firewalls/armid1"),
	// 	SystemData: &armpanngfw.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-09-09T05:08:24.229Z"); return t}()),
	// 		CreatedBy: to.Ptr("praval"),
	// 		CreatedByType: to.Ptr(armpanngfw.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-09-09T05:08:24.229Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("praval"),
	// 		LastModifiedByType: to.Ptr(armpanngfw.CreatedByTypeUser),
	// 	},
	// 	Location: to.Ptr("eastus"),
	// 	Tags: map[string]*string{
	// 		"tagName": to.Ptr("value"),
	// 	},
	// 	Identity: &armpanngfw.AzureResourceManagerManagedIdentityProperties{
	// 		Type: to.Ptr(armpanngfw.ManagedIdentityTypeNone),
	// 		PrincipalID: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaa"),
	// 		TenantID: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
	// 		UserAssignedIdentities: map[string]*armpanngfw.AzureResourceManagerUserAssignedIdentity{
	// 			"key16": &armpanngfw.AzureResourceManagerUserAssignedIdentity{
	// 				ClientID: to.Ptr("aaaa"),
	// 				PrincipalID: to.Ptr("aaaaaaaaaaaaaaa"),
	// 			},
	// 		},
	// 	},
	// 	Properties: &armpanngfw.FirewallDeploymentProperties{
	// 		AssociatedRulestack: &armpanngfw.RulestackDetails{
	// 			Location: to.Ptr("eastus"),
	// 			ResourceID: to.Ptr("aaaaaaaaaa"),
	// 			RulestackID: to.Ptr("aaaaaaaaaaaaaaaa"),
	// 		},
	// 		DNSSettings: &armpanngfw.DNSSettings{
	// 			DNSServers: []*armpanngfw.IPAddress{
	// 				{
	// 					Address: to.Ptr("20.22.92.111"),
	// 					ResourceID: to.Ptr("/subscriptions/01c7d41f-afaf-464e-8a8b-5c6f9f98cee8/resourceGroups/mj-liftr-integration/providers/Microsoft.Network/publicIPAddresses/mj-liftr-integration-egressNatIp1"),
	// 			}},
	// 			EnableDNSProxy: to.Ptr(armpanngfw.DNSProxyDISABLED),
	// 			EnabledDNSType: to.Ptr(armpanngfw.EnabledDNSTypeCUSTOM),
	// 		},
	// 		FrontEndSettings: []*armpanngfw.FrontendSetting{
	// 			{
	// 				Name: to.Ptr("frontendsetting11"),
	// 				BackendConfiguration: &armpanngfw.EndpointConfiguration{
	// 					Address: &armpanngfw.IPAddress{
	// 						Address: to.Ptr("20.22.32.136"),
	// 						ResourceID: to.Ptr("/subscriptions/01c7d41f-afaf-464e-8a8b-5c6f9f98cee8/resourceGroups/mj-liftr-integration/providers/Microsoft.Network/publicIPAddresses/mj-liftr-integration-frontendSettingIp2"),
	// 					},
	// 					Port: to.Ptr("80"),
	// 				},
	// 				FrontendConfiguration: &armpanngfw.EndpointConfiguration{
	// 					Address: &armpanngfw.IPAddress{
	// 						Address: to.Ptr("20.22.91.251"),
	// 						ResourceID: to.Ptr("/subscriptions/01c7d41f-afaf-464e-8a8b-5c6f9f98cee8/resourceGroups/mj-liftr-integration/providers/Microsoft.Network/publicIPAddresses/mj-liftr-integration-frontendSettingIp1"),
	// 					},
	// 					Port: to.Ptr("80"),
	// 				},
	// 				Protocol: to.Ptr(armpanngfw.ProtocolTypeTCP),
	// 		}},
	// 		IsPanoramaManaged: to.Ptr(armpanngfw.BooleanEnumTRUE),
	// 		MarketplaceDetails: &armpanngfw.MarketplaceDetails{
	// 			MarketplaceSubscriptionID: to.Ptr("aa"),
	// 			MarketplaceSubscriptionStatus: to.Ptr(armpanngfw.MarketplaceSubscriptionStatusPendingFulfillmentStart),
	// 			OfferID: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
	// 			PublisherID: to.Ptr("aaaa"),
	// 		},
	// 		NetworkProfile: &armpanngfw.NetworkProfile{
	// 			EgressNatIP: []*armpanngfw.IPAddress{
	// 				{
	// 					Address: to.Ptr("20.22.92.111"),
	// 					ResourceID: to.Ptr("/subscriptions/01c7d41f-afaf-464e-8a8b-5c6f9f98cee8/resourceGroups/mj-liftr-integration/providers/Microsoft.Network/publicIPAddresses/mj-liftr-integration-egressNatIp1"),
	// 			}},
	// 			EnableEgressNat: to.Ptr(armpanngfw.EgressNatENABLED),
	// 			NetworkType: to.Ptr(armpanngfw.NetworkTypeVNET),
	// 			PublicIPs: []*armpanngfw.IPAddress{
	// 				{
	// 					Address: to.Ptr("20.22.92.11"),
	// 					ResourceID: to.Ptr("/subscriptions/01c7d41f-afaf-464e-8a8b-5c6f9f98cee8/resourceGroups/mj-liftr-integration/providers/Microsoft.Network/publicIPAddresses/mj-liftr-integration-PublicIp1"),
	// 			}},
	// 			VnetConfiguration: &armpanngfw.VnetConfiguration{
	// 				IPOfTrustSubnetForUdr: &armpanngfw.IPAddress{
	// 					Address: to.Ptr("10.1.1.0/24"),
	// 					ResourceID: to.Ptr("/subscriptions/2bf4a339-294d-4c25-b0b2-ef649e9f5c27/resourceGroups/os-liftr-integration/providers/Microsoft.Network/virtualNetworks/os-liftr-integration-vnet/subnets/os-liftr-integration-untrust-subnet"),
	// 				},
	// 				TrustSubnet: &armpanngfw.IPAddressSpace{
	// 					AddressSpace: to.Ptr("10.1.1.0/24"),
	// 					ResourceID: to.Ptr("/subscriptions/2bf4a339-294d-4c25-b0b2-ef649e9f5c27/resourceGroups/os-liftr-integration/providers/Microsoft.Network/virtualNetworks/os-liftr-integration-vnet/subnets/os-liftr-integration-trust-subnet"),
	// 				},
	// 				UnTrustSubnet: &armpanngfw.IPAddressSpace{
	// 					AddressSpace: to.Ptr("10.1.1.0/24"),
	// 					ResourceID: to.Ptr("/subscriptions/2bf4a339-294d-4c25-b0b2-ef649e9f5c27/resourceGroups/os-liftr-integration/providers/Microsoft.Network/virtualNetworks/os-liftr-integration-vnet/subnets/os-liftr-integration-untrust-subnet"),
	// 				},
	// 				Vnet: &armpanngfw.IPAddressSpace{
	// 					AddressSpace: to.Ptr("10.1.0.0/16"),
	// 					ResourceID: to.Ptr("/subscriptions/2bf4a339-294d-4c25-b0b2-ef649e9f5c27/resourceGroups/os-liftr-integration/providers/Microsoft.Network/virtualNetworks/os-liftr-integration-vnet"),
	// 				},
	// 			},
	// 			VwanConfiguration: &armpanngfw.VwanConfiguration{
	// 				IPOfTrustSubnetForUdr: &armpanngfw.IPAddress{
	// 					Address: to.Ptr("10.1.1.0/24"),
	// 					ResourceID: to.Ptr("/subscriptions/2bf4a339-294d-4c25-b0b2-ef649e9f5c27/resourceGroups/os-liftr-integration/providers/Microsoft.Network/virtualNetworks/os-liftr-integration-vnet/subnets/os-liftr-integration-untrust-subnet"),
	// 				},
	// 				NetworkVirtualApplianceID: to.Ptr("2bf4a339-294d-4c25-b0b2-ef649e9f5c12"),
	// 				TrustSubnet: &armpanngfw.IPAddressSpace{
	// 					AddressSpace: to.Ptr("10.1.1.0/24"),
	// 					ResourceID: to.Ptr("/subscriptions/2bf4a339-294d-4c25-b0b2-ef649e9f5c27/resourceGroups/os-liftr-integration/providers/Microsoft.Network/virtualNetworks/os-liftr-integration-vnet/subnets/os-liftr-integration-trust-subnet"),
	// 				},
	// 				UnTrustSubnet: &armpanngfw.IPAddressSpace{
	// 					AddressSpace: to.Ptr("10.1.1.0/24"),
	// 					ResourceID: to.Ptr("/subscriptions/2bf4a339-294d-4c25-b0b2-ef649e9f5c27/resourceGroups/os-liftr-integration/providers/Microsoft.Network/virtualNetworks/os-liftr-integration-vnet/subnets/os-liftr-integration-untrust-subnet"),
	// 				},
	// 				VHub: &armpanngfw.IPAddressSpace{
	// 					AddressSpace: to.Ptr("10.1.1.0/24"),
	// 					ResourceID: to.Ptr("/subscriptions/2bf4a339-294d-4c25-b0b2-ef649e9f5c27/resourceGroups/os-liftr-integration/providers/Microsoft.Network/virtualNetworks/os-liftr-integration-vnet/subnets/os-liftr-integration-untrust-subnet"),
	// 				},
	// 			},
	// 		},
	// 		PanEtag: to.Ptr("2bf4a339-294d-4c25-b0b2-ef649e9f5c12"),
	// 		PanoramaConfig: &armpanngfw.PanoramaConfig{
	// 			CgName: to.Ptr("PanoramaCollectorGroup"),
	// 			ConfigString: to.Ptr("bas64EncodedString"),
	// 			DgName: to.Ptr("PanoramaDeviceGroup"),
	// 			HostName: to.Ptr("hostname"),
	// 			PanoramaServer: to.Ptr("10.25.1.1"),
	// 			PanoramaServer2: to.Ptr("10.20.1.1"),
	// 			TplName: to.Ptr("PanoramaTemplateStack"),
	// 			VMAuthKey: to.Ptr("SSH_AUTH_KEY"),
	// 		},
	// 		PlanData: &armpanngfw.PlanData{
	// 			BillingCycle: to.Ptr(armpanngfw.BillingCycleWEEKLY),
	// 			EffectiveDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-09-13T00:46:05.283Z"); return t}()),
	// 			PlanID: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
	// 			UsageType: to.Ptr(armpanngfw.UsageTypePAYG),
	// 		},
	// 		ProvisioningState: to.Ptr(armpanngfw.ProvisioningStateSucceeded),
	// 	},
	// }
}
