package armpeering_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/peering/armpeering"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/peering/resource-manager/Microsoft.Peering/stable/2022-01-01/examples/CreatePeeringWithExchangeRouteServer.json
func ExamplePeeringsClient_CreateOrUpdate_createAPeeringWithExchangeRouteServer() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpeering.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPeeringsClient().CreateOrUpdate(ctx, "rgName", "peeringName", armpeering.Peering{
		Kind:     to.Ptr(armpeering.KindDirect),
		Location: to.Ptr("eastus"),
		Properties: &armpeering.Properties{
			Direct: &armpeering.PropertiesDirect{
				Connections: []*armpeering.DirectConnection{
					{
						BandwidthInMbps: to.Ptr[int32](10000),
						BgpSession: &armpeering.BgpSession{
							MaxPrefixesAdvertisedV4:     to.Ptr[int32](1000),
							MaxPrefixesAdvertisedV6:     to.Ptr[int32](100),
							MicrosoftSessionIPv4Address: to.Ptr("192.168.0.123"),
							PeerSessionIPv4Address:      to.Ptr("192.168.0.234"),
							SessionPrefixV4:             to.Ptr("192.168.0.0/24"),
						},
						ConnectionIdentifier:   to.Ptr("5F4CB5C7-6B43-4444-9338-9ABC72606C16"),
						PeeringDBFacilityID:    to.Ptr[int32](99999),
						SessionAddressProvider: to.Ptr(armpeering.SessionAddressProviderPeer),
						UseForPeeringService:   to.Ptr(true),
					}},
				DirectPeeringType: to.Ptr(armpeering.DirectPeeringTypeIxRs),
				PeerAsn: &armpeering.SubResource{
					ID: to.Ptr("/subscriptions/subId/providers/Microsoft.Peering/peerAsns/myAsn1"),
				},
			},
			PeeringLocation: to.Ptr("peeringLocation0"),
		},
		SKU: &armpeering.SKU{
			Name: to.Ptr("Premium_Direct_Free"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Peering = armpeering.Peering{
	// 	Name: to.Ptr("peeringName"),
	// 	Type: to.Ptr("Microsoft.Peering/peerings"),
	// 	ID: to.Ptr("/subscriptions/subId/resourceGroups/rgName/providers/Microsoft.Peering/peerings/peeringName"),
	// 	Kind: to.Ptr(armpeering.KindDirect),
	// 	Location: to.Ptr("eastus"),
	// 	Properties: &armpeering.Properties{
	// 		Direct: &armpeering.PropertiesDirect{
	// 			Connections: []*armpeering.DirectConnection{
	// 				{
	// 					BandwidthInMbps: to.Ptr[int32](10000),
	// 					BgpSession: &armpeering.BgpSession{
	// 						MaxPrefixesAdvertisedV4: to.Ptr[int32](1000),
	// 						MaxPrefixesAdvertisedV6: to.Ptr[int32](100),
	// 						MicrosoftSessionIPv4Address: to.Ptr("192.168.0.123"),
	// 						PeerSessionIPv4Address: to.Ptr("192.168.0.234"),
	// 						SessionPrefixV4: to.Ptr("192.168.0.0/24"),
	// 						SessionStateV4: to.Ptr(armpeering.SessionStateV4Established),
	// 						SessionStateV6: to.Ptr(armpeering.SessionStateV6Established),
	// 					},
	// 					ConnectionIdentifier: to.Ptr("5F4CB5C7-6B43-4444-9338-9ABC72606C16"),
	// 					ConnectionState: to.Ptr(armpeering.ConnectionStateActive),
	// 					PeeringDBFacilityID: to.Ptr[int32](99999),
	// 					ProvisionedBandwidthInMbps: to.Ptr[int32](10000),
	// 					SessionAddressProvider: to.Ptr(armpeering.SessionAddressProviderPeer),
	// 					UseForPeeringService: to.Ptr(true),
	// 			}},
	// 			DirectPeeringType: to.Ptr(armpeering.DirectPeeringTypeIxRs),
	// 			PeerAsn: &armpeering.SubResource{
	// 				ID: to.Ptr("/subscriptions/subId/providers/Microsoft.Peering/peerAsns/myAsn1"),
	// 			},
	// 			UseForPeeringService: to.Ptr(true),
	// 		},
	// 		PeeringLocation: to.Ptr("peeringLocation0"),
	// 		ProvisioningState: to.Ptr(armpeering.ProvisioningStateSucceeded),
	// 	},
	// 	SKU: &armpeering.SKU{
	// 		Name: to.Ptr("Premium_Direct_Free"),
	// 		Family: to.Ptr(armpeering.FamilyDirect),
	// 		Size: to.Ptr(armpeering.SizeFree),
	// 		Tier: to.Ptr(armpeering.TierPremium),
	// 	},
	// }
}
