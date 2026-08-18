package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v11"
)

// Generated from example definition: 2025-09-01/NetworkVirtualApplianceVnetDualStackPut.json
func ExampleVirtualAppliancesClient_BeginCreateOrUpdate_createNvaInVNetForIpv4AndIpv6() {
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
			AddressFamily: []*armnetwork.IPVersion{
				to.Ptr(armnetwork.IPVersionIPv4),
				to.Ptr(armnetwork.IPVersionIPv6),
			},
			BootStrapConfigurationBlobs: []*string{
				to.Ptr("https://csrncvhdstorage1.blob.core.windows.net/csrncvhdstoragecont/csrbootstrapconfig"),
			},
			CloudInitConfigurationBlobs: []*string{
				to.Ptr("https://csrncvhdstorage1.blob.core.windows.net/csrncvhdstoragecont/csrcloudinitconfig"),
			},
			NvaInterfaceConfigurations: []*armnetwork.NvaInterfaceConfigurationsProperties{
				{
					Name: to.Ptr("privateInterface"),
					Type: []*armnetwork.NvaNicType{
						to.Ptr(armnetwork.NvaNicTypePrivateNic),
					},
					Subnet: &armnetwork.NvaInVnetSubnetReferenceProperties{
						ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnet1/subnets/subnet1"),
					},
				},
				{
					Name: to.Ptr("publicInterface"),
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
	// 			AddressFamily: []*armnetwork.IPVersion{
	// 				to.Ptr(armnetwork.IPVersionIPv4),
	// 				to.Ptr(armnetwork.IPVersionIPv6),
	// 			},
	// 			AddressPrefix: to.Ptr("10.26.112.0/25"),
	// 			AddressPrefixV6: to.Ptr("2001:db8:26:5::/64"),
	// 			BootStrapConfigurationBlobs: []*string{
	// 				to.Ptr("https://csrncvhdstorage1.blob.core.windows.net/csrncvhdstoragecont/csrbootstrapconfig"),
	// 			},
	// 			CloudInitConfigurationBlobs: []*string{
	// 				to.Ptr("https://csrncvhdstorage1.blob.core.windows.net/csrncvhdstoragecont/csrcloudinitconfig"),
	// 			},
	// 			NvaInterfaceConfigurations: []*armnetwork.NvaInterfaceConfigurationsProperties{
	// 				{
	// 					Name: to.Ptr("privateInterface"),
	// 					Type: []*armnetwork.NvaNicType{
	// 						to.Ptr(armnetwork.NvaNicTypePrivateNic),
	// 					},
	// 					Subnet: &armnetwork.NvaInVnetSubnetReferenceProperties{
	// 						ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnet1/subnets/subnet1"),
	// 					},
	// 				},
	// 				{
	// 					Name: to.Ptr("publicInterface"),
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
	// 			PrivateIPAddress: to.Ptr("10.26.112.10"),
	// 			PrivateIPAddressV6: to.Ptr("2001:db8:26:4::10"),
	// 			ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 			VirtualApplianceAsn: to.Ptr[int64](10000),
	// 			VirtualApplianceNics: []*armnetwork.VirtualApplianceNicProperties{
	// 				{
	// 					NicType: to.Ptr(armnetwork.NicTypeInResponsePrivateNic),
	// 					Name: to.Ptr("privateInterface-ipconfig"),
	// 					PublicIPAddress: to.Ptr(""),
	// 					PrivateIPAddress: to.Ptr("10.26.112.11"),
	// 					PrivateIPAddressV6: to.Ptr("2001:db8:26:5::11"),
	// 					PublicIPAddressV6: to.Ptr(""),
	// 					InstanceName: to.Ptr("nva_0"),
	// 				},
	// 				{
	// 					NicType: to.Ptr(armnetwork.NicTypeInResponsePublicNic),
	// 					Name: to.Ptr("publicInterface-ipconfig"),
	// 					PublicIPAddress: to.Ptr("20.70.202.149"),
	// 					PrivateIPAddress: to.Ptr("10.26.112.132"),
	// 					PrivateIPAddressV6: to.Ptr("2001:db8:26:6::10"),
	// 					PublicIPAddressV6: to.Ptr("2603:1010:3:17::52"),
	// 					InstanceName: to.Ptr("nva_0"),
	// 				},
	// 				{
	// 					NicType: to.Ptr(armnetwork.NicTypeInResponsePrivateNic),
	// 					Name: to.Ptr("privateInterface-ipconfig"),
	// 					PublicIPAddress: to.Ptr(""),
	// 					PrivateIPAddress: to.Ptr("10.26.112.12"),
	// 					PrivateIPAddressV6: to.Ptr("2001:db8:26:5::5"),
	// 					PublicIPAddressV6: to.Ptr(""),
	// 					InstanceName: to.Ptr("nva_1"),
	// 				},
	// 				{
	// 					NicType: to.Ptr(armnetwork.NicTypeInResponsePublicNic),
	// 					Name: to.Ptr("publicInterface-ipconfig"),
	// 					PublicIPAddress: to.Ptr("20.211.41.245"),
	// 					PrivateIPAddress: to.Ptr("10.26.112.133"),
	// 					PrivateIPAddressV6: to.Ptr("2001:db8:26:6::11"),
	// 					PublicIPAddressV6: to.Ptr("2603:1010:3:2::12"),
	// 					InstanceName: to.Ptr("nva_1"),
	// 				},
	// 			},
	// 		},
	// 		Tags: map[string]*string{
	// 			"key1": to.Ptr("value1"),
	// 		},
	// 	},
	// }
}
