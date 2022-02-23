Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fnetwork%2Farmnetwork%2Fv0.3.1/sdk/resourcemanager/network/armnetwork/README.md) on how to add the SDK to your project and authenticate.

```go
package armnetwork_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
)

// x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/VpnSitePut.json
func ExampleVPNSitesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armnetwork.NewVPNSitesClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<vpn-site-name>",
		armnetwork.VPNSite{
			Location: to.StringPtr("<location>"),
			Tags: map[string]*string{
				"key1": to.StringPtr("value1"),
			},
			Properties: &armnetwork.VPNSiteProperties{
				AddressSpace: &armnetwork.AddressSpace{
					AddressPrefixes: []*string{
						to.StringPtr("10.0.0.0/16")},
				},
				IsSecuritySite: to.BoolPtr(false),
				O365Policy: &armnetwork.O365PolicyProperties{
					BreakOutCategories: &armnetwork.O365BreakOutCategoryPolicies{
						Default:  to.BoolPtr(false),
						Allow:    to.BoolPtr(true),
						Optimize: to.BoolPtr(true),
					},
				},
				VirtualWan: &armnetwork.SubResource{
					ID: to.StringPtr("<id>"),
				},
				VPNSiteLinks: []*armnetwork.VPNSiteLink{
					{
						Name: to.StringPtr("<name>"),
						Properties: &armnetwork.VPNSiteLinkProperties{
							BgpProperties: &armnetwork.VPNLinkBgpSettings{
								Asn:               to.Int64Ptr(1234),
								BgpPeeringAddress: to.StringPtr("<bgp-peering-address>"),
							},
							Fqdn:      to.StringPtr("<fqdn>"),
							IPAddress: to.StringPtr("<ipaddress>"),
							LinkProperties: &armnetwork.VPNLinkProviderProperties{
								LinkProviderName: to.StringPtr("<link-provider-name>"),
								LinkSpeedInMbps:  to.Int32Ptr(0),
							},
						},
					}},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.VPNSitesClientCreateOrUpdateResult)
}
```
