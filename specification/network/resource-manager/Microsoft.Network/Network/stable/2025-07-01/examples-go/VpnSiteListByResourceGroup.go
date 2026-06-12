package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v10"
)

// Generated from example definition: 2025-07-01/VpnSiteListByResourceGroup.json
func ExampleVPNSitesClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewVPNSitesClient().NewListByResourceGroupPager("rg1", nil)
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
		// page = armnetwork.VPNSitesClientListByResourceGroupResponse{
		// 	ListVPNSitesResult: armnetwork.ListVPNSitesResult{
		// 		Value: []*armnetwork.VPNSite{
		// 			{
		// 				Name: to.Ptr("vpnSite1"),
		// 				Type: to.Ptr("Microsoft.Network/vpnSites"),
		// 				Etag: to.Ptr("w/\\00000000-0000-0000-0000-000000000000\\"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/vpnSites/vpnSite1"),
		// 				Location: to.Ptr("West US"),
		// 				Properties: &armnetwork.VPNSiteProperties{
		// 					AddressSpace: &armnetwork.AddressSpace{
		// 						AddressPrefixes: []*string{
		// 							to.Ptr("10.0.0.0/16"),
		// 						},
		// 					},
		// 					DeviceProperties: &armnetwork.DeviceProperties{
		// 						LinkSpeedInMbps: to.Ptr[int32](0),
		// 					},
		// 					IsSecuritySite: to.Ptr(false),
		// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 					VirtualWan: &armnetwork.SubResource{
		// 						ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualWANs/wan1"),
		// 					},
		// 					VPNSiteLinks: []*armnetwork.VPNSiteLink{
		// 						{
		// 							Name: to.Ptr("vpnSiteLink1"),
		// 							Type: to.Ptr("Microsoft.Network/vpnSites/vpnSiteLinks"),
		// 							Etag: to.Ptr("W/\"00000000-0000-0000-0000-000000000000\""),
		// 							ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/vpnSites/vpnSite1/vpnSiteLinks/vpnSiteLink1"),
		// 							Properties: &armnetwork.VPNSiteLinkProperties{
		// 								BgpProperties: &armnetwork.VPNLinkBgpSettings{
		// 									Asn: to.Ptr[int64](1234),
		// 									BgpPeeringAddress: to.Ptr("192.168.0.0"),
		// 								},
		// 								IPAddress: to.Ptr("50.50.50.56"),
		// 								LinkProperties: &armnetwork.VPNLinkProviderProperties{
		// 									LinkSpeedInMbps: to.Ptr[int32](0),
		// 								},
		// 								ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 							},
		// 						},
		// 					},
		// 				},
		// 				Tags: map[string]*string{
		// 					"key1": to.Ptr("value1"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("vpnSite2"),
		// 				Type: to.Ptr("Microsoft.Network/vpnSites"),
		// 				Etag: to.Ptr("w/\\00000000-0000-0000-0000-000000000000\\"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/vpnSites/vpnSite2"),
		// 				Location: to.Ptr("West US"),
		// 				Properties: &armnetwork.VPNSiteProperties{
		// 					AddressSpace: &armnetwork.AddressSpace{
		// 						AddressPrefixes: []*string{
		// 							to.Ptr("10.0.0.0/16"),
		// 						},
		// 					},
		// 					BgpProperties: &armnetwork.BgpSettings{
		// 						Asn: to.Ptr[int64](1234),
		// 						BgpPeeringAddress: to.Ptr("192.168.0.0"),
		// 					},
		// 					DeviceProperties: &armnetwork.DeviceProperties{
		// 						DeviceModel: to.Ptr("model01"),
		// 						DeviceVendor: to.Ptr("vendor1"),
		// 						LinkSpeedInMbps: to.Ptr[int32](200),
		// 					},
		// 					IPAddress: to.Ptr("10.1.0.0"),
		// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 					VirtualWan: &armnetwork.SubResource{
		// 						ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualWANs/wan1"),
		// 					},
		// 				},
		// 				Tags: map[string]*string{
		// 					"key1": to.Ptr("value1"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
