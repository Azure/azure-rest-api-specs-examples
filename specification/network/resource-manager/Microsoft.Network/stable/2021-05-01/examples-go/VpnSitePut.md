Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fnetwork%2Farmnetwork%2Fv0.5.0/sdk/resourcemanager/network/armnetwork/README.md) on how to add the SDK to your project and authenticate.

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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/VpnSitePut.json
func ExampleVPNSitesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armnetwork.NewVPNSitesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<vpn-site-name>",
		armnetwork.VPNSite{
			Location: to.Ptr("<location>"),
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
					ID: to.Ptr("<id>"),
				},
				VPNSiteLinks: []*armnetwork.VPNSiteLink{
					{
						Name: to.Ptr("<name>"),
						Properties: &armnetwork.VPNSiteLinkProperties{
							BgpProperties: &armnetwork.VPNLinkBgpSettings{
								Asn:               to.Ptr[int64](1234),
								BgpPeeringAddress: to.Ptr("<bgp-peering-address>"),
							},
							Fqdn:      to.Ptr("<fqdn>"),
							IPAddress: to.Ptr("<ipaddress>"),
							LinkProperties: &armnetwork.VPNLinkProviderProperties{
								LinkProviderName: to.Ptr("<link-provider-name>"),
								LinkSpeedInMbps:  to.Ptr[int32](0),
							},
						},
					}},
			},
		},
		&armnetwork.VPNSitesClientBeginCreateOrUpdateOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
