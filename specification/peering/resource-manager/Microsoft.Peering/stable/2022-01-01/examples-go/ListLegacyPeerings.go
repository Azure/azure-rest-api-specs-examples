package armpeering_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/peering/armpeering"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/peering/resource-manager/Microsoft.Peering/stable/2022-01-01/examples/ListLegacyPeerings.json
func ExampleLegacyPeeringsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpeering.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewLegacyPeeringsClient().NewListPager("peeringLocation0", armpeering.LegacyPeeringsKindExchange, &armpeering.LegacyPeeringsClientListOptions{Asn: nil})
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
		// page.ListResult = armpeering.ListResult{
		// 	Value: []*armpeering.Peering{
		// 		{
		// 			Name: to.Ptr("peeringName"),
		// 			Type: to.Ptr("Microsoft.Peering/peerings"),
		// 			ID: to.Ptr("/subscriptions/subId/providers/Microsoft.Peering/peerings/peeringName"),
		// 			Kind: to.Ptr(armpeering.KindExchange),
		// 			Location: to.Ptr("centralus"),
		// 			Properties: &armpeering.Properties{
		// 				Exchange: &armpeering.PropertiesExchange{
		// 					Connections: []*armpeering.ExchangeConnection{
		// 						{
		// 							BgpSession: &armpeering.BgpSession{
		// 								MaxPrefixesAdvertisedV4: to.Ptr[int32](1000),
		// 								MaxPrefixesAdvertisedV6: to.Ptr[int32](100),
		// 								MD5AuthenticationKey: to.Ptr("test-md5-auth-key"),
		// 								MicrosoftSessionIPv4Address: to.Ptr("192.168.3.1"),
		// 								MicrosoftSessionIPv6Address: to.Ptr("fd00::1:1"),
		// 								PeerSessionIPv4Address: to.Ptr("192.168.2.1"),
		// 								PeerSessionIPv6Address: to.Ptr("fd00::1"),
		// 								SessionStateV4: to.Ptr(armpeering.SessionStateV4Established),
		// 								SessionStateV6: to.Ptr(armpeering.SessionStateV6Established),
		// 							},
		// 							ConnectionIdentifier: to.Ptr("CE495334-0E94-4E51-8164-8116D6CD284D"),
		// 							ConnectionState: to.Ptr(armpeering.ConnectionStateActive),
		// 							PeeringDBFacilityID: to.Ptr[int32](99999),
		// 						},
		// 						{
		// 							BgpSession: &armpeering.BgpSession{
		// 								MaxPrefixesAdvertisedV4: to.Ptr[int32](1000),
		// 								MaxPrefixesAdvertisedV6: to.Ptr[int32](100),
		// 								MD5AuthenticationKey: to.Ptr("test-md5-auth-key"),
		// 								MicrosoftSessionIPv4Address: to.Ptr("192.168.3.2"),
		// 								MicrosoftSessionIPv6Address: to.Ptr("fd00::1:2"),
		// 								PeerSessionIPv4Address: to.Ptr("192.168.2.2"),
		// 								PeerSessionIPv6Address: to.Ptr("fd00::2"),
		// 								SessionStateV4: to.Ptr(armpeering.SessionStateV4Established),
		// 								SessionStateV6: to.Ptr(armpeering.SessionStateV6Established),
		// 							},
		// 							ConnectionIdentifier: to.Ptr("CDD8E673-CB07-47E6-84DE-3739F778762B"),
		// 							ConnectionState: to.Ptr(armpeering.ConnectionStateActive),
		// 							PeeringDBFacilityID: to.Ptr[int32](99999),
		// 					}},
		// 					PeerAsn: &armpeering.SubResource{
		// 						ID: to.Ptr("65000"),
		// 					},
		// 				},
		// 				PeeringLocation: to.Ptr("peeringLocation0"),
		// 				ProvisioningState: to.Ptr(armpeering.ProvisioningStateSucceeded),
		// 			},
		// 			SKU: &armpeering.SKU{
		// 				Name: to.Ptr("Basic_Exchange_Free"),
		// 				Family: to.Ptr(armpeering.FamilyExchange),
		// 				Size: to.Ptr(armpeering.SizeFree),
		// 				Tier: to.Ptr(armpeering.TierBasic),
		// 			},
		// 	}},
		// }
	}
}
