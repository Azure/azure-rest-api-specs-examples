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
	}
	ctx := context.Background()
	client, err := armpeering.NewPeeringsClient("subId", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx,
		"rgName",
		"peeringName",
		armpeering.Peering{
			Kind:     to.Ptr(armpeering.KindDirect),
			Location: to.Ptr("eastus"),
			Properties: &armpeering.Properties{
				Direct: &armpeering.PropertiesDirect{
					Connections: []*armpeering.DirectConnection{
						{
							BandwidthInMbps: to.Ptr[int32](10000),
							BgpSession: &armpeering.BgpSession{
								MaxPrefixesAdvertisedV4: to.Ptr[int32](1000),
								MaxPrefixesAdvertisedV6: to.Ptr[int32](100),
								MD5AuthenticationKey:    to.Ptr("test-md5-auth-key"),
								SessionPrefixV4:         to.Ptr("192.168.0.0/31"),
								SessionPrefixV6:         to.Ptr("fd00::0/127"),
							},
							ConnectionIdentifier:   to.Ptr("5F4CB5C7-6B43-4444-9338-9ABC72606C16"),
							PeeringDBFacilityID:    to.Ptr[int32](99999),
							SessionAddressProvider: to.Ptr(armpeering.SessionAddressProviderPeer),
							UseForPeeringService:   to.Ptr(false),
						},
						{
							BandwidthInMbps:        to.Ptr[int32](10000),
							ConnectionIdentifier:   to.Ptr("8AB00818-D533-4504-A25A-03A17F61201C"),
							PeeringDBFacilityID:    to.Ptr[int32](99999),
							SessionAddressProvider: to.Ptr(armpeering.SessionAddressProviderMicrosoft),
							UseForPeeringService:   to.Ptr(true),
						}},
					DirectPeeringType: to.Ptr(armpeering.DirectPeeringTypeEdge),
					PeerAsn: &armpeering.SubResource{
						ID: to.Ptr("/subscriptions/subId/providers/Microsoft.Peering/peerAsns/myAsn1"),
					},
				},
				PeeringLocation: to.Ptr("peeringLocation0"),
			},
			SKU: &armpeering.SKU{
				Name: to.Ptr("Basic_Direct_Free"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
