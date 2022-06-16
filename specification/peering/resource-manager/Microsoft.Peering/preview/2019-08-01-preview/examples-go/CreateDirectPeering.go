package armpeering_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/peering/armpeering"
)

// x-ms-original-file: specification/peering/resource-manager/Microsoft.Peering/preview/2019-08-01-preview/examples/CreateDirectPeering.json
func ExamplePeeringsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armpeering.NewPeeringsClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<peering-name>",
		armpeering.Peering{
			Kind:     armpeering.Kind("Direct").ToPtr(),
			Location: to.StringPtr("<location>"),
			Properties: &armpeering.Properties{
				Direct: &armpeering.PropertiesDirect{
					Connections: []*armpeering.DirectConnection{
						{
							BandwidthInMbps: to.Int32Ptr(10000),
							BgpSession: &armpeering.BgpSession{
								MaxPrefixesAdvertisedV4: to.Int32Ptr(1000),
								MaxPrefixesAdvertisedV6: to.Int32Ptr(100),
								MD5AuthenticationKey:    to.StringPtr("<md5authentication-key>"),
								SessionPrefixV4:         to.StringPtr("<session-prefix-v4>"),
								SessionPrefixV6:         to.StringPtr("<session-prefix-v6>"),
							},
							ConnectionIdentifier:   to.StringPtr("<connection-identifier>"),
							PeeringDBFacilityID:    to.Int32Ptr(99999),
							SessionAddressProvider: armpeering.SessionAddressProvider("Peer").ToPtr(),
							UseForPeeringService:   to.BoolPtr(false),
						},
						{
							BandwidthInMbps: to.Int32Ptr(10000),
							BgpSession: &armpeering.BgpSession{
								MaxPrefixesAdvertisedV4: to.Int32Ptr(1000),
								MaxPrefixesAdvertisedV6: to.Int32Ptr(100),
								MD5AuthenticationKey:    to.StringPtr("<md5authentication-key>"),
								SessionPrefixV4:         to.StringPtr("<session-prefix-v4>"),
								SessionPrefixV6:         to.StringPtr("<session-prefix-v6>"),
							},
							ConnectionIdentifier:   to.StringPtr("<connection-identifier>"),
							PeeringDBFacilityID:    to.Int32Ptr(99999),
							SessionAddressProvider: armpeering.SessionAddressProvider("Microsoft").ToPtr(),
							UseForPeeringService:   to.BoolPtr(true),
						}},
					DirectPeeringType: armpeering.DirectPeeringType("Edge").ToPtr(),
					PeerAsn: &armpeering.SubResource{
						ID: to.StringPtr("<id>"),
					},
					UseForPeeringService: to.BoolPtr(false),
				},
				PeeringLocation: to.StringPtr("<peering-location>"),
			},
			SKU: &armpeering.SKU{
				Name: armpeering.Name("Basic_Direct_Free").ToPtr(),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.PeeringsClientCreateOrUpdateResult)
}
