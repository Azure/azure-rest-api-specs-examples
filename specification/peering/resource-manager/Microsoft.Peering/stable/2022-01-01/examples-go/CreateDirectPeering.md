Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fpeering%2Farmpeering%2Fv0.5.0/sdk/resourcemanager/peering/armpeering/README.md) on how to add the SDK to your project and authenticate.

```go
package armpeering_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/peering/armpeering"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/peering/resource-manager/Microsoft.Peering/stable/2022-01-01/examples/CreateDirectPeering.json
func ExamplePeeringsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armpeering.NewPeeringsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<peering-name>",
		armpeering.Peering{
			Kind:     to.Ptr(armpeering.KindDirect),
			Location: to.Ptr("<location>"),
			Properties: &armpeering.Properties{
				Direct: &armpeering.PropertiesDirect{
					Connections: []*armpeering.DirectConnection{
						{
							BandwidthInMbps: to.Ptr[int32](10000),
							BgpSession: &armpeering.BgpSession{
								MaxPrefixesAdvertisedV4: to.Ptr[int32](1000),
								MaxPrefixesAdvertisedV6: to.Ptr[int32](100),
								MD5AuthenticationKey:    to.Ptr("<md5authentication-key>"),
								SessionPrefixV4:         to.Ptr("<session-prefix-v4>"),
								SessionPrefixV6:         to.Ptr("<session-prefix-v6>"),
							},
							ConnectionIdentifier:   to.Ptr("<connection-identifier>"),
							PeeringDBFacilityID:    to.Ptr[int32](99999),
							SessionAddressProvider: to.Ptr(armpeering.SessionAddressProviderPeer),
							UseForPeeringService:   to.Ptr(false),
						},
						{
							BandwidthInMbps:        to.Ptr[int32](10000),
							ConnectionIdentifier:   to.Ptr("<connection-identifier>"),
							PeeringDBFacilityID:    to.Ptr[int32](99999),
							SessionAddressProvider: to.Ptr(armpeering.SessionAddressProviderMicrosoft),
							UseForPeeringService:   to.Ptr(true),
						}},
					DirectPeeringType: to.Ptr(armpeering.DirectPeeringTypeEdge),
					PeerAsn: &armpeering.SubResource{
						ID: to.Ptr("<id>"),
					},
				},
				PeeringLocation: to.Ptr("<peering-location>"),
			},
			SKU: &armpeering.SKU{
				Name: to.Ptr("<name>"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
