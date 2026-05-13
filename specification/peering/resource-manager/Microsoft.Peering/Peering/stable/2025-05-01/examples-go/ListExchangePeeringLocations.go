package armpeering_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/peering/armpeering/v2"
)

// Generated from example definition: 2025-05-01/ListExchangePeeringLocations.json
func ExampleLocationsClient_NewListPager_listExchangePeeringLocations() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpeering.NewClientFactory("subId", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewLocationsClient().NewListPager(armpeering.PeeringLocationsKindExchange, nil)
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
		// page = armpeering.LocationsClientListResponse{
		// 	LocationListResult: armpeering.LocationListResult{
		// 		Value: []*armpeering.Location{
		// 			{
		// 				Name: to.Ptr("peeringLocation1"),
		// 				Type: to.Ptr("Microsoft.Peering/peeringLocations"),
		// 				ID: to.Ptr("/subscriptions/subId/providers/Microsoft.Peering/peeringLocations/peeringLocation1"),
		// 				Kind: to.Ptr(armpeering.KindExchange),
		// 				Properties: &armpeering.LocationProperties{
		// 					Country: to.Ptr("country1"),
		// 					Exchange: &armpeering.LocationPropertiesExchange{
		// 						PeeringFacilities: []*armpeering.ExchangePeeringFacility{
		// 							{
		// 								BandwidthInMbps: to.Ptr[int32](10000),
		// 								ExchangeName: to.Ptr("name1"),
		// 								FacilityIPv4Prefix: to.Ptr("192.168.128.0/17"),
		// 								FacilityIPv6Prefix: to.Ptr("fd00::1000:0/98"),
		// 								MicrosoftIPv4Address: to.Ptr("192.168.131.1"),
		// 								MicrosoftIPv6Address: to.Ptr("fd00::1:1"),
		// 								PeeringDBFacilityID: to.Ptr[int32](99999),
		// 								PeeringDBFacilityLink: to.Ptr("https://www.peeringdb.com/ix/99999"),
		// 							},
		// 							{
		// 								BandwidthInMbps: to.Ptr[int32](10000),
		// 								ExchangeName: to.Ptr("name3"),
		// 								FacilityIPv4Prefix: to.Ptr("192.168.0.0/17"),
		// 								FacilityIPv6Prefix: to.Ptr("fd00::0/98"),
		// 								MicrosoftIPv4Address: to.Ptr("192.168.2.2"),
		// 								MicrosoftIPv6Address: to.Ptr("fd00::2"),
		// 								PeeringDBFacilityID: to.Ptr[int32](99999),
		// 								PeeringDBFacilityLink: to.Ptr("https://www.peeringdb.com/ix/99999"),
		// 							},
		// 						},
		// 					},
		// 					PeeringLocation: to.Ptr("peeringLocation1"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("peeringLocation2"),
		// 				Type: to.Ptr("Microsoft.Peering/peeringLocations"),
		// 				ID: to.Ptr("/subscriptions/subId/providers/Microsoft.Peering/peeringLocations/peeringLocation2"),
		// 				Kind: to.Ptr(armpeering.KindExchange),
		// 				Properties: &armpeering.LocationProperties{
		// 					Country: to.Ptr("country2"),
		// 					Exchange: &armpeering.LocationPropertiesExchange{
		// 						PeeringFacilities: []*armpeering.ExchangePeeringFacility{
		// 							{
		// 								BandwidthInMbps: to.Ptr[int32](10000),
		// 								ExchangeName: to.Ptr("name2"),
		// 								FacilityIPv4Prefix: to.Ptr("192.168.0.0/16"),
		// 								FacilityIPv6Prefix: to.Ptr("fd00::0/98"),
		// 								MicrosoftIPv4Address: to.Ptr("192.168.2.1"),
		// 								MicrosoftIPv6Address: to.Ptr("fd00::2"),
		// 								PeeringDBFacilityID: to.Ptr[int32](99999),
		// 								PeeringDBFacilityLink: to.Ptr("https://www.peeringdb.com/ix/99999"),
		// 							},
		// 						},
		// 					},
		// 					PeeringLocation: to.Ptr("peeringLocation2"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
