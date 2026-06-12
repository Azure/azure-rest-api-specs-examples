package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v10"
)

// Generated from example definition: 2025-07-01/NetworkVirtualApplianceVnetNetworkProfilePut.json
func ExampleVirtualAppliancesClient_BeginCreateOrUpdate_createNvaInVNetWithPrivateNicPublicNicIncludingNetworkProfile() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualAppliancesClient().BeginCreateOrUpdate(ctx, "rg1", "nva", armnetwork.VirtualAppliance{
		Identity: &armnetwork.ManagedServiceIdentity{
			Type: to.Ptr(armnetwork.ResourceIdentityTypeUserAssigned),
			UserAssignedIdentities: map[string]*armnetwork.ManagedServiceIdentityUserAssignedIdentities{
				"/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/rg1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identity1": {},
			},
		},
		Location: to.Ptr("West US"),
		Properties: &armnetwork.VirtualAppliancePropertiesFormat{
			BootStrapConfigurationBlobs: []*string{
				to.Ptr("https://csrncvhdstorage1.blob.core.windows.net/csrncvhdstoragecont/csrbootstrapconfig"),
			},
			CloudInitConfigurationBlobs: []*string{
				to.Ptr("https://csrncvhdstorage1.blob.core.windows.net/csrncvhdstoragecont/csrcloudinitconfig"),
			},
			NetworkProfile: &armnetwork.VirtualAppliancePropertiesFormatNetworkProfile{
				NetworkInterfaceConfigurations: []*armnetwork.VirtualApplianceNetworkInterfaceConfiguration{
					{
						NicType: to.Ptr(armnetwork.NicTypeInRequestPublicNic),
						Properties: &armnetwork.VirtualApplianceNetworkInterfaceConfigurationProperties{
							IPConfigurations: []*armnetwork.VirtualApplianceIPConfiguration{
								{
									Name: to.Ptr("myPrimaryPublicIpConfig"),
									Properties: &armnetwork.VirtualApplianceIPConfigurationProperties{
										Primary: to.Ptr(true),
									},
								},
								{
									Name: to.Ptr("myNonPrimaryPublicIpConfig"),
									Properties: &armnetwork.VirtualApplianceIPConfigurationProperties{
										Primary: to.Ptr(false),
									},
								},
							},
						},
					},
					{
						NicType: to.Ptr(armnetwork.NicTypeInRequestPrivateNic),
						Properties: &armnetwork.VirtualApplianceNetworkInterfaceConfigurationProperties{
							IPConfigurations: []*armnetwork.VirtualApplianceIPConfiguration{
								{
									Name: to.Ptr("myPrimaryPrivateIpConfig"),
									Properties: &armnetwork.VirtualApplianceIPConfigurationProperties{
										Primary: to.Ptr(true),
									},
								},
								{
									Name: to.Ptr("myNonPrimaryPrivateIpConfig"),
									Properties: &armnetwork.VirtualApplianceIPConfigurationProperties{
										Primary: to.Ptr(false),
									},
								},
							},
						},
					},
				},
			},
			NvaInterfaceConfigurations: []*armnetwork.NvaInterfaceConfigurationsProperties{
				{
					Name: to.Ptr("dataInterface"),
					Type: []*armnetwork.NvaNicType{
						to.Ptr(armnetwork.NvaNicTypePrivateNic),
					},
					Subnet: &armnetwork.NvaInVnetSubnetReferenceProperties{
						ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnet1/subnets/subnet1"),
					},
				},
				{
					Name: to.Ptr("managementInterface"),
					Type: []*armnetwork.NvaNicType{
						to.Ptr(armnetwork.NvaNicTypePublicNic),
					},
					Subnet: &armnetwork.NvaInVnetSubnetReferenceProperties{
						ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnet1/subnets/subnet2"),
					},
				},
			},
			NvaSKU: &armnetwork.VirtualApplianceSKUProperties{
				BundledScaleUnit:   to.Ptr("1"),
				MarketPlaceVersion: to.Ptr("latest"),
				Vendor:             to.Ptr("Cisco SDWAN"),
			},
			VirtualApplianceAsn: to.Ptr[int64](10000),
		},
		Tags: map[string]*string{
			"key1": to.Ptr("value1"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armnetwork.VirtualAppliancesClientCreateOrUpdateResponse{
	// 	VirtualAppliance: armnetwork.VirtualAppliance{
	// 		Name: to.Ptr("nva"),
	// 		Type: to.Ptr("Microsoft.Network/networkVirtualAppliances"),
	// 		Etag: to.Ptr("w/\\00000000-0000-0000-0000-000000000000\\"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/networkVirtualAppliances/nva"),
	// 		Identity: &armnetwork.ManagedServiceIdentity{
	// 			Type: to.Ptr(armnetwork.ResourceIdentityTypeUserAssigned),
	// 			UserAssignedIdentities: map[string]*armnetwork.ManagedServiceIdentityUserAssignedIdentities{
	// 				"/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/rg1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identity1": &armnetwork.ManagedServiceIdentityUserAssignedIdentities{
	// 				},
	// 			},
	// 		},
	// 		Location: to.Ptr("West US"),
	// 		Properties: &armnetwork.VirtualAppliancePropertiesFormat{
	// 			AddressPrefix: to.Ptr("192.168.1.0/16"),
	// 			BootStrapConfigurationBlobs: []*string{
	// 				to.Ptr("https://csrncvhdstorage1.blob.core.windows.net/csrncvhdstoragecont/csrbootstrapconfig"),
	// 			},
	// 			CloudInitConfigurationBlobs: []*string{
	// 				to.Ptr("https://csrncvhdstorage1.blob.core.windows.net/csrncvhdstoragecont/csrcloudinitconfig"),
	// 			},
	// 			InboundSecurityRules: []*armnetwork.SubResource{
	// 				{
	// 					ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/networkVirtualAppliances/nva/InboundSecurityRules/rule1"),
	// 				},
	// 			},
	// 			NetworkProfile: &armnetwork.VirtualAppliancePropertiesFormatNetworkProfile{
	// 				NetworkInterfaceConfigurations: []*armnetwork.VirtualApplianceNetworkInterfaceConfiguration{
	// 					{
	// 						NicType: to.Ptr(armnetwork.NicTypeInRequestPublicNic),
	// 						Properties: &armnetwork.VirtualApplianceNetworkInterfaceConfigurationProperties{
	// 							IPConfigurations: []*armnetwork.VirtualApplianceIPConfiguration{
	// 								{
	// 									Name: to.Ptr("myPrimaryPublicIpConfig"),
	// 									Properties: &armnetwork.VirtualApplianceIPConfigurationProperties{
	// 										Primary: to.Ptr(true),
	// 									},
	// 								},
	// 								{
	// 									Name: to.Ptr("myNonPrimaryPublicIpConfig"),
	// 									Properties: &armnetwork.VirtualApplianceIPConfigurationProperties{
	// 										Primary: to.Ptr(false),
	// 									},
	// 								},
	// 							},
	// 						},
	// 					},
	// 					{
	// 						NicType: to.Ptr(armnetwork.NicTypeInRequestPrivateNic),
	// 						Properties: &armnetwork.VirtualApplianceNetworkInterfaceConfigurationProperties{
	// 							IPConfigurations: []*armnetwork.VirtualApplianceIPConfiguration{
	// 								{
	// 									Name: to.Ptr("myPrimaryPrivateIpConfig"),
	// 									Properties: &armnetwork.VirtualApplianceIPConfigurationProperties{
	// 										Primary: to.Ptr(true),
	// 									},
	// 								},
	// 								{
	// 									Name: to.Ptr("myNonPrimaryPrivateIpConfig"),
	// 									Properties: &armnetwork.VirtualApplianceIPConfigurationProperties{
	// 										Primary: to.Ptr(false),
	// 									},
	// 								},
	// 							},
	// 						},
	// 					},
	// 				},
	// 			},
	// 			NvaInterfaceConfigurations: []*armnetwork.NvaInterfaceConfigurationsProperties{
	// 				{
	// 					Name: to.Ptr("dataInterface"),
	// 					Type: []*armnetwork.NvaNicType{
	// 						to.Ptr(armnetwork.NvaNicTypePrivateNic),
	// 					},
	// 					Subnet: &armnetwork.NvaInVnetSubnetReferenceProperties{
	// 						ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnet1/subnets/subnet1"),
	// 					},
	// 				},
	// 				{
	// 					Name: to.Ptr("managementInterface"),
	// 					Type: []*armnetwork.NvaNicType{
	// 						to.Ptr(armnetwork.NvaNicTypePublicNic),
	// 					},
	// 					Subnet: &armnetwork.NvaInVnetSubnetReferenceProperties{
	// 						ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnet1/subnets/subnet2"),
	// 					},
	// 				},
	// 			},
	// 			NvaSKU: &armnetwork.VirtualApplianceSKUProperties{
	// 				BundledScaleUnit: to.Ptr("1"),
	// 				MarketPlaceVersion: to.Ptr("latest"),
	// 				Vendor: to.Ptr("Cisco SDWAN"),
	// 			},
	// 			PrivateIPAddress: to.Ptr("192.168.12.9"),
	// 			ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 			VirtualApplianceAsn: to.Ptr[int64](10000),
	// 			VirtualApplianceNics: []*armnetwork.VirtualApplianceNicProperties{
	// 				{
	// 					Name: to.Ptr("myPrimaryPublicIpConfig"),
	// 					InstanceName: to.Ptr("nva_0"),
	// 					NicType: to.Ptr(armnetwork.NicTypeInResponsePublicNic),
	// 					PrivateIPAddress: to.Ptr("192.168.12.1"),
	// 					PublicIPAddress: to.Ptr("40.30.2.2"),
	// 				},
	// 				{
	// 					Name: to.Ptr("myNonPrimaryPublicIpConfig"),
	// 					InstanceName: to.Ptr("nva_0"),
	// 					NicType: to.Ptr(armnetwork.NicTypeInResponsePublicNic),
	// 					PrivateIPAddress: to.Ptr("192.168.12.2"),
	// 					PublicIPAddress: to.Ptr("40.30.2.3"),
	// 				},
	// 				{
	// 					Name: to.Ptr("myPrimaryPrivateIpConfig"),
	// 					InstanceName: to.Ptr("nva_0"),
	// 					NicType: to.Ptr(armnetwork.NicTypeInResponsePrivateNic),
	// 					PrivateIPAddress: to.Ptr("192.168.12.3"),
	// 					PublicIPAddress: to.Ptr(""),
	// 				},
	// 				{
	// 					Name: to.Ptr("myNonPrimaryPrivateIpConfig"),
	// 					InstanceName: to.Ptr("nva_0"),
	// 					NicType: to.Ptr(armnetwork.NicTypeInResponsePrivateNic),
	// 					PrivateIPAddress: to.Ptr("192.168.12.4"),
	// 					PublicIPAddress: to.Ptr(""),
	// 				},
	// 				{
	// 					Name: to.Ptr("myPrimaryPublicIpConfig"),
	// 					InstanceName: to.Ptr("nva_1"),
	// 					NicType: to.Ptr(armnetwork.NicTypeInResponsePublicNic),
	// 					PrivateIPAddress: to.Ptr("192.168.12.5"),
	// 					PublicIPAddress: to.Ptr("40.30.2.4"),
	// 				},
	// 				{
	// 					Name: to.Ptr("myNonPrimaryPublicIpConfig"),
	// 					InstanceName: to.Ptr("nva_1"),
	// 					NicType: to.Ptr(armnetwork.NicTypeInResponsePublicNic),
	// 					PrivateIPAddress: to.Ptr("192.168.12.6"),
	// 					PublicIPAddress: to.Ptr("40.30.2.5"),
	// 				},
	// 				{
	// 					Name: to.Ptr("myPrimaryPrivateIpConfig"),
	// 					InstanceName: to.Ptr("nva_1"),
	// 					NicType: to.Ptr(armnetwork.NicTypeInResponsePrivateNic),
	// 					PrivateIPAddress: to.Ptr("192.168.12.7"),
	// 					PublicIPAddress: to.Ptr(""),
	// 				},
	// 				{
	// 					Name: to.Ptr("myNonPrimaryPrivateIpConfig"),
	// 					InstanceName: to.Ptr("nva_1"),
	// 					NicType: to.Ptr(armnetwork.NicTypeInResponsePrivateNic),
	// 					PrivateIPAddress: to.Ptr("192.168.12.8"),
	// 					PublicIPAddress: to.Ptr(""),
	// 				},
	// 			},
	// 		},
	// 		Tags: map[string]*string{
	// 			"key1": to.Ptr("value1"),
	// 		},
	// 	},
	// }
}
