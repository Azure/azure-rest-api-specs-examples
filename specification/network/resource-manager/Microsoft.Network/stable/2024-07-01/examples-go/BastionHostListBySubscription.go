package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c58fa033619b12c7cfa8a0ec5a9bf03bb18869ab/specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/BastionHostListBySubscription.json
func ExampleBastionHostsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewBastionHostsClient().NewListPager(nil)
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
		// page.BastionHostListResult = armnetwork.BastionHostListResult{
		// 	Value: []*armnetwork.BastionHost{
		// 		{
		// 			Name: to.Ptr("bastionhosttenant"),
		// 			Type: to.Ptr("Microsoft.Network/bastionHosts"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/bastionHosts/bastionhosttenant"),
		// 			Location: to.Ptr("West US"),
		// 			Etag: to.Ptr("w/\\00000000-0000-0000-0000-000000000000\\"),
		// 			Properties: &armnetwork.BastionHostPropertiesFormat{
		// 				DisableCopyPaste: to.Ptr(false),
		// 				DNSName: to.Ptr("bst-9d89d361-100e-4c01-b92d-466548c476dc.bastion.azure.com"),
		// 				EnableIPConnect: to.Ptr(false),
		// 				EnableKerberos: to.Ptr(false),
		// 				EnablePrivateOnlyBastion: to.Ptr(false),
		// 				EnableSessionRecording: to.Ptr(false),
		// 				EnableShareableLink: to.Ptr(false),
		// 				EnableTunneling: to.Ptr(false),
		// 				IPConfigurations: []*armnetwork.BastionHostIPConfiguration{
		// 					{
		// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/bastionHosts/bastionhosttenant/bastionHostIpConfigurations/bastionHostIpConfiguration"),
		// 						Name: to.Ptr("bastionHostIpConfiguration"),
		// 						Type: to.Ptr("Microsoft.Network/bastionHosts/bastionHostIpConfigurations"),
		// 						Etag: to.Ptr("w/\\00000000-0000-0000-0000-000000000000\\"),
		// 						Properties: &armnetwork.BastionHostIPConfigurationPropertiesFormat{
		// 							PrivateIPAllocationMethod: to.Ptr(armnetwork.IPAllocationMethodDynamic),
		// 							ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 							PublicIPAddress: &armnetwork.SubResource{
		// 								ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/publicIPAddresses/pipName"),
		// 							},
		// 							Subnet: &armnetwork.SubResource{
		// 								ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnet/subnets/AzureBastionSubnet"),
		// 							},
		// 						},
		// 				}},
		// 				ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 				ScaleUnits: to.Ptr[int32](2),
		// 			},
		// 			SKU: &armnetwork.SKU{
		// 				Name: to.Ptr(armnetwork.BastionHostSKUNameStandard),
		// 			},
		// 			Zones: []*string{
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("bastionhostdeveloper"),
		// 			Type: to.Ptr("Microsoft.Network/bastionHosts"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg2/providers/Microsoft.Network/bastionHosts/bastionhostdeveloper"),
		// 			Location: to.Ptr("West US"),
		// 			Etag: to.Ptr("w/\\00000000-0000-0000-0000-000000000000\\"),
		// 			Properties: &armnetwork.BastionHostPropertiesFormat{
		// 				DNSName: to.Ptr("omnibrain.uswest.bastionglobal.azure.com"),
		// 				IPConfigurations: []*armnetwork.BastionHostIPConfiguration{
		// 				},
		// 				NetworkACLs: &armnetwork.BastionHostPropertiesFormatNetworkACLs{
		// 					IPRules: []*armnetwork.IPRule{
		// 						{
		// 							AddressPrefix: to.Ptr("1.1.1.1/16"),
		// 					}},
		// 				},
		// 				ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 				VirtualNetwork: &armnetwork.SubResource{
		// 					ID: to.Ptr("/subscriptions/subid/resourceGroups/rg2/providers/Microsoft.Network/virtualNetworks/vnet2"),
		// 				},
		// 			},
		// 			SKU: &armnetwork.SKU{
		// 				Name: to.Ptr(armnetwork.BastionHostSKUNameDeveloper),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("bastionhost3"),
		// 			Type: to.Ptr("Microsoft.Network/bastionHosts"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/bastionHosts/bastionhost3"),
		// 			Location: to.Ptr("West US"),
		// 			Etag: to.Ptr("w/\\00000000-0000-0000-0000-000000000000\\"),
		// 			Properties: &armnetwork.BastionHostPropertiesFormat{
		// 				DisableCopyPaste: to.Ptr(false),
		// 				DNSName: to.Ptr("bst-9d89d361-100e-4c01-b92d-466548c476dc.bastion.azure.com"),
		// 				EnableIPConnect: to.Ptr(false),
		// 				EnableKerberos: to.Ptr(false),
		// 				EnablePrivateOnlyBastion: to.Ptr(false),
		// 				EnableSessionRecording: to.Ptr(true),
		// 				EnableShareableLink: to.Ptr(false),
		// 				EnableTunneling: to.Ptr(false),
		// 				IPConfigurations: []*armnetwork.BastionHostIPConfiguration{
		// 					{
		// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/bastionHosts/bastionhost3/bastionHostIpConfigurations/IpConf"),
		// 						Name: to.Ptr("bastionHostIpConfiguration"),
		// 						Type: to.Ptr("Microsoft.Network/bastionHosts/bastionHostIpConfigurations"),
		// 						Etag: to.Ptr("w/\\00000000-0000-0000-0000-000000000000\\"),
		// 						Properties: &armnetwork.BastionHostIPConfigurationPropertiesFormat{
		// 							PrivateIPAllocationMethod: to.Ptr(armnetwork.IPAllocationMethodDynamic),
		// 							ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 							PublicIPAddress: &armnetwork.SubResource{
		// 								ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/publicIPAddresses/pipName3"),
		// 							},
		// 							Subnet: &armnetwork.SubResource{
		// 								ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnet3/subnets/AzureBastionSubnet"),
		// 							},
		// 						},
		// 				}},
		// 				ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 				ScaleUnits: to.Ptr[int32](2),
		// 			},
		// 			SKU: &armnetwork.SKU{
		// 				Name: to.Ptr(armnetwork.BastionHostSKUNamePremium),
		// 			},
		// 			Zones: []*string{
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("bastionhost4"),
		// 			Type: to.Ptr("Microsoft.Network/bastionHosts"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/bastionHosts/bastionhost4"),
		// 			Location: to.Ptr("West US"),
		// 			Etag: to.Ptr("w/\\00000000-0000-0000-0000-000000000000\\"),
		// 			Properties: &armnetwork.BastionHostPropertiesFormat{
		// 				DisableCopyPaste: to.Ptr(false),
		// 				DNSName: to.Ptr("bst-9d89d361-100e-4c01-b92d-466548c476dc.bastion.azure.com"),
		// 				EnableIPConnect: to.Ptr(false),
		// 				EnableKerberos: to.Ptr(false),
		// 				EnablePrivateOnlyBastion: to.Ptr(true),
		// 				EnableSessionRecording: to.Ptr(false),
		// 				EnableShareableLink: to.Ptr(false),
		// 				EnableTunneling: to.Ptr(false),
		// 				IPConfigurations: []*armnetwork.BastionHostIPConfiguration{
		// 					{
		// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/bastionHosts/bastionhost4/bastionHostIpConfigurations/IpConf"),
		// 						Name: to.Ptr("bastionHostIpConfiguration"),
		// 						Type: to.Ptr("Microsoft.Network/bastionHosts/bastionHostIpConfigurations"),
		// 						Etag: to.Ptr("w/\\00000000-0000-0000-0000-000000000000\\"),
		// 						Properties: &armnetwork.BastionHostIPConfigurationPropertiesFormat{
		// 							PrivateIPAllocationMethod: to.Ptr(armnetwork.IPAllocationMethodDynamic),
		// 							ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 							Subnet: &armnetwork.SubResource{
		// 								ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnet4/subnets/AzureBastionSubnet"),
		// 							},
		// 						},
		// 				}},
		// 				ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 				ScaleUnits: to.Ptr[int32](2),
		// 			},
		// 			SKU: &armnetwork.SKU{
		// 				Name: to.Ptr(armnetwork.BastionHostSKUNamePremium),
		// 			},
		// 			Zones: []*string{
		// 			},
		// 	}},
		// }
	}
}
