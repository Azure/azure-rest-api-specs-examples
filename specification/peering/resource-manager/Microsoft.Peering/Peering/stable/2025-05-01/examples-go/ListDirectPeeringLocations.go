package armpeering_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/peering/armpeering/v2"
)

// Generated from example definition: 2025-05-01/ListDirectPeeringLocations.json
func ExampleLocationsClient_NewListPager_listDirectPeeringLocations() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpeering.NewClientFactory("subId", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewLocationsClient().NewListPager(armpeering.PeeringLocationsKindDirect, nil)
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
		// 				Kind: to.Ptr(armpeering.KindDirect),
		// 				Properties: &armpeering.LocationProperties{
		// 					Country: to.Ptr("country1"),
		// 					Direct: &armpeering.LocationPropertiesDirect{
		// 						BandwidthOffers: []*armpeering.BandwidthOffer{
		// 							{
		// 								OfferName: to.Ptr("10Gbps"),
		// 								ValueInMbps: to.Ptr[int32](10000),
		// 							},
		// 							{
		// 								OfferName: to.Ptr("100Gbps"),
		// 								ValueInMbps: to.Ptr[int32](100000),
		// 							},
		// 						},
		// 						PeeringFacilities: []*armpeering.DirectPeeringFacility{
		// 							{
		// 								Address: to.Ptr("address1"),
		// 								DirectPeeringType: to.Ptr(armpeering.DirectPeeringTypeEdge),
		// 								PeeringDBFacilityID: to.Ptr[int32](99999),
		// 								PeeringDBFacilityLink: to.Ptr("https://www.peeringdb.com/fac/99999"),
		// 							},
		// 							{
		// 								Address: to.Ptr("address3"),
		// 								DirectPeeringType: to.Ptr(armpeering.DirectPeeringTypeCdn),
		// 								PeeringDBFacilityID: to.Ptr[int32](99999),
		// 								PeeringDBFacilityLink: to.Ptr("https://www.peeringdb.com/fac/99999"),
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
		// 				Kind: to.Ptr(armpeering.KindDirect),
		// 				Properties: &armpeering.LocationProperties{
		// 					Country: to.Ptr("country2"),
		// 					Direct: &armpeering.LocationPropertiesDirect{
		// 						BandwidthOffers: []*armpeering.BandwidthOffer{
		// 							{
		// 								OfferName: to.Ptr("10Gbps"),
		// 								ValueInMbps: to.Ptr[int32](10000),
		// 							},
		// 							{
		// 								OfferName: to.Ptr("100Gbps"),
		// 								ValueInMbps: to.Ptr[int32](100000),
		// 							},
		// 						},
		// 						PeeringFacilities: []*armpeering.DirectPeeringFacility{
		// 							{
		// 								Address: to.Ptr("address2"),
		// 								DirectPeeringType: to.Ptr(armpeering.DirectPeeringTypeEdge),
		// 								PeeringDBFacilityID: to.Ptr[int32](99999),
		// 								PeeringDBFacilityLink: to.Ptr("https://www.peeringdb.com/fac/99999"),
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
