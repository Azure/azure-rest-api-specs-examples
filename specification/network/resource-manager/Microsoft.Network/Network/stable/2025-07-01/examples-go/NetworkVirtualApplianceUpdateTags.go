package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v10"
)

// Generated from example definition: 2025-07-01/NetworkVirtualApplianceUpdateTags.json
func ExampleVirtualAppliancesClient_UpdateTags() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVirtualAppliancesClient().UpdateTags(ctx, "rg1", "nva", armnetwork.TagsObject{
		Tags: map[string]*string{
			"key1": to.Ptr("value1"),
			"key2": to.Ptr("value2"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armnetwork.VirtualAppliancesClientUpdateTagsResponse{
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
	// 			NvaSKU: &armnetwork.VirtualApplianceSKUProperties{
	// 				BundledScaleUnit: to.Ptr("1"),
	// 				MarketPlaceVersion: to.Ptr("12.1"),
	// 				Vendor: to.Ptr("Cisco SDWAN"),
	// 			},
	// 			ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 			VirtualApplianceAsn: to.Ptr[int64](10000),
	// 			VirtualApplianceNics: []*armnetwork.VirtualApplianceNicProperties{
	// 				{
	// 					Name: to.Ptr("managementNic"),
	// 					PrivateIPAddress: to.Ptr("192.168.12.1"),
	// 					PublicIPAddress: to.Ptr("40.30.2.2"),
	// 				},
	// 				{
	// 					Name: to.Ptr("privateNic-1"),
	// 					PrivateIPAddress: to.Ptr("192.168.12.2"),
	// 				},
	// 			},
	// 			VirtualApplianceSites: []*armnetwork.SubResource{
	// 				{
	// 					ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/networtkVirtualAppliances/nva/virtualApplianceSites/site1"),
	// 				},
	// 			},
	// 			VirtualHub: &armnetwork.SubResource{
	// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/hub1"),
	// 			},
	// 		},
	// 		Tags: map[string]*string{
	// 			"key1": to.Ptr("value1"),
	// 			"key2": to.Ptr("value2"),
	// 		},
	// 	},
	// }
}
