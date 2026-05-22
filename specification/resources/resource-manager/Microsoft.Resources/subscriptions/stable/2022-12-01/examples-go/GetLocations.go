package armsubscriptions_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armsubscriptions/v2"
)

// Generated from example definition: 2022-12-01/GetLocations.json
func ExampleClient_NewListLocationsPager_getLocationsWithASubscriptionId() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsubscriptions.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewClient().NewListLocationsPager("a1ffc958-d2c7-493e-9f1e-125a0477f536", nil)
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
		// page = armsubscriptions.ClientListLocationsResponse{
		// 	LocationListResult: armsubscriptions.LocationListResult{
		// 		Value: []*armsubscriptions.Location{
		// 			{
		// 				Name: to.Ptr("eastus"),
		// 				Type: to.Ptr(armsubscriptions.LocationTypeRegion),
		// 				AvailabilityZoneMappings: []*armsubscriptions.AvailabilityZoneMappings{
		// 					{
		// 						LogicalZone: to.Ptr("1"),
		// 						PhysicalZone: to.Ptr("eastus-az1"),
		// 					},
		// 					{
		// 						LogicalZone: to.Ptr("2"),
		// 						PhysicalZone: to.Ptr("eastus-az3"),
		// 					},
		// 					{
		// 						LogicalZone: to.Ptr("3"),
		// 						PhysicalZone: to.Ptr("eastus-az2"),
		// 					},
		// 				},
		// 				DisplayName: to.Ptr("East US"),
		// 				ID: to.Ptr("/subscriptions/a1ffc958-d2c7-493e-9f1e-125a0477f536/locations/eastus"),
		// 				Metadata: &armsubscriptions.LocationMetadata{
		// 					Geography: to.Ptr("United States"),
		// 					GeographyGroup: to.Ptr("US"),
		// 					Latitude: to.Ptr("37.3719"),
		// 					Longitude: to.Ptr("-79.8164"),
		// 					PairedRegion: []*armsubscriptions.PairedRegion{
		// 						{
		// 							Name: to.Ptr("westus"),
		// 							ID: to.Ptr("/subscriptions/a1ffc958-d2c7-493e-9f1e-125a0477f536/locations/westus"),
		// 						},
		// 					},
		// 					PhysicalLocation: to.Ptr("Virginia"),
		// 					RegionCategory: to.Ptr(armsubscriptions.RegionCategoryRecommended),
		// 					RegionType: to.Ptr(armsubscriptions.RegionTypePhysical),
		// 				},
		// 				RegionalDisplayName: to.Ptr("(US) East US"),
		// 			},
		// 			{
		// 				Name: to.Ptr("eastus2"),
		// 				Type: to.Ptr(armsubscriptions.LocationTypeRegion),
		// 				AvailabilityZoneMappings: []*armsubscriptions.AvailabilityZoneMappings{
		// 					{
		// 						LogicalZone: to.Ptr("1"),
		// 						PhysicalZone: to.Ptr("eastus2-az1"),
		// 					},
		// 					{
		// 						LogicalZone: to.Ptr("2"),
		// 						PhysicalZone: to.Ptr("eastus2-az3"),
		// 					},
		// 					{
		// 						LogicalZone: to.Ptr("3"),
		// 						PhysicalZone: to.Ptr("eastus2-az2"),
		// 					},
		// 				},
		// 				DisplayName: to.Ptr("East US 2"),
		// 				ID: to.Ptr("/subscriptions/a1ffc958-d2c7-493e-9f1e-125a0477f536/locations/eastus2"),
		// 				Metadata: &armsubscriptions.LocationMetadata{
		// 					Geography: to.Ptr("United States"),
		// 					GeographyGroup: to.Ptr("US"),
		// 					Latitude: to.Ptr("36.6681"),
		// 					Longitude: to.Ptr("-78.3889"),
		// 					PairedRegion: []*armsubscriptions.PairedRegion{
		// 						{
		// 							Name: to.Ptr("centralus"),
		// 							ID: to.Ptr("/subscriptions/a1ffc958-d2c7-493e-9f1e-125a0477f536/locations/centralus"),
		// 						},
		// 					},
		// 					PhysicalLocation: to.Ptr("Virginia"),
		// 					RegionCategory: to.Ptr(armsubscriptions.RegionCategoryRecommended),
		// 					RegionType: to.Ptr(armsubscriptions.RegionTypePhysical),
		// 				},
		// 				RegionalDisplayName: to.Ptr("(US) East US 2"),
		// 			},
		// 		},
		// 	},
		// }
	}
}
