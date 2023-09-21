package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/80c21c17b4a7aa57f637ee594f7cfd653255a7e0/specification/network/resource-manager/Microsoft.Network/stable/2023-05-01/examples/VpnSitePut.json
func ExampleVPNSitesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVPNSitesClient().BeginCreateOrUpdate(ctx, "rg1", "vpnSite1", armnetwork.VPNSite{
		Location: to.Ptr("West US"),
		Tags: map[string]*string{
			"key1": to.Ptr("value1"),
		},
		Properties: &armnetwork.VPNSiteProperties{
			AddressSpace: &armnetwork.AddressSpace{
				AddressPrefixes: []*string{
					to.Ptr("10.0.0.0/16")},
			},
			IsSecuritySite: to.Ptr(false),
			O365Policy: &armnetwork.O365PolicyProperties{
				BreakOutCategories: &armnetwork.O365BreakOutCategoryPolicies{
					Default:  to.Ptr(false),
					Allow:    to.Ptr(true),
					Optimize: to.Ptr(true),
				},
			},
			VirtualWan: &armnetwork.SubResource{
				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualWANs/wan1"),
			},
			VPNSiteLinks: []*armnetwork.VPNSiteLink{
				{
					Name: to.Ptr("vpnSiteLink1"),
					Properties: &armnetwork.VPNSiteLinkProperties{
						BgpProperties: &armnetwork.VPNLinkBgpSettings{
							Asn:               to.Ptr[int64](1234),
							BgpPeeringAddress: to.Ptr("192.168.0.0"),
						},
						Fqdn:      to.Ptr("link1.vpnsite1.contoso.com"),
						IPAddress: to.Ptr("50.50.50.56"),
						LinkProperties: &armnetwork.VPNLinkProviderProperties{
							LinkProviderName: to.Ptr("vendor1"),
							LinkSpeedInMbps:  to.Ptr[int32](0),
						},
					},
				}},
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
	// res.VPNSite = armnetwork.VPNSite{
	// 	Name: to.Ptr("vpnSite1"),
	// 	Type: to.Ptr("Microsoft.Network/vpnSites"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/vpnSites/vpnSite1"),
	// 	Location: to.Ptr("West US"),
	// 	Tags: map[string]*string{
	// 		"key1": to.Ptr("value1"),
	// 	},
	// 	Etag: to.Ptr("w/\\00000000-0000-0000-0000-000000000000\\"),
	// 	Properties: &armnetwork.VPNSiteProperties{
	// 		AddressSpace: &armnetwork.AddressSpace{
	// 			AddressPrefixes: []*string{
	// 				to.Ptr("10.0.0.0/16")},
	// 			},
	// 			DeviceProperties: &armnetwork.DeviceProperties{
	// 				LinkSpeedInMbps: to.Ptr[int32](0),
	// 			},
	// 			IsSecuritySite: to.Ptr(false),
	// 			ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 			VirtualWan: &armnetwork.SubResource{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualWANs/wan1"),
	// 			},
	// 			VPNSiteLinks: []*armnetwork.VPNSiteLink{
	// 				{
	// 					ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/vpnSites/vpnSite1/vpnSiteLinks/vpnSiteLink1"),
	// 					Name: to.Ptr("vpnSiteLink1"),
	// 					Type: to.Ptr("Microsoft.Network/vpnSites/vpnSiteLinks"),
	// 					Etag: to.Ptr("W/\"00000000-0000-0000-0000-000000000000\""),
	// 					Properties: &armnetwork.VPNSiteLinkProperties{
	// 						BgpProperties: &armnetwork.VPNLinkBgpSettings{
	// 							Asn: to.Ptr[int64](1234),
	// 							BgpPeeringAddress: to.Ptr("192.168.0.0"),
	// 						},
	// 						IPAddress: to.Ptr("50.50.50.56"),
	// 						LinkProperties: &armnetwork.VPNLinkProviderProperties{
	// 							LinkProviderName: to.Ptr("vendor1"),
	// 							LinkSpeedInMbps: to.Ptr[int32](0),
	// 						},
	// 						ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 					},
	// 			}},
	// 		},
	// 	}
}
