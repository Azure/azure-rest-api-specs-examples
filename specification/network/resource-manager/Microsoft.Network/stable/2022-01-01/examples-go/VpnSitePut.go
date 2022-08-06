package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/VpnSitePut.json
func ExampleVPNSitesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armnetwork.NewVPNSitesClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx, "rg1", "vpnSite1", armnetwork.VPNSite{
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
	// TODO: use response item
	_ = res
}
