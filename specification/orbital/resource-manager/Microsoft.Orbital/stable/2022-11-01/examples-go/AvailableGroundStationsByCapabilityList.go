package armorbital_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/orbital/armorbital/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1e7b408f3323e7f5424745718fe62c7a043a2337/specification/orbital/resource-manager/Microsoft.Orbital/stable/2022-11-01/examples/AvailableGroundStationsByCapabilityList.json
func ExampleAvailableGroundStationsClient_NewListByCapabilityPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armorbital.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAvailableGroundStationsClient().NewListByCapabilityPager(armorbital.CapabilityParameterEarthObservation, nil)
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
		// page.AvailableGroundStationListResult = armorbital.AvailableGroundStationListResult{
		// 	Value: []*armorbital.AvailableGroundStation{
		// 		{
		// 			Name: to.Ptr("EASTUS2_0"),
		// 			Type: to.Ptr("Microsoft.Orbital/availableGroundStations"),
		// 			ID: to.Ptr("/subscriptions/c1be1141-a7c9-4aac-9608-3c2e2f1152c3/providers/Microsoft.Orbital/availableGroundStations/EASTUS2_0"),
		// 			Location: to.Ptr("eastus2"),
		// 			Properties: &armorbital.AvailableGroundStationProperties{
		// 				AltitudeMeters: to.Ptr[float32](1641),
		// 				City: to.Ptr("Atlanta"),
		// 				LatitudeDegrees: to.Ptr[float32](33.74831),
		// 				LongitudeDegrees: to.Ptr[float32](-84.39111),
		// 				ProviderName: to.Ptr("Microsoft"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("EASTUS_1"),
		// 			Type: to.Ptr("Microsoft.Orbital/availableGroundStations"),
		// 			ID: to.Ptr("/subscriptions/c1be1141-a7c9-4aac-9608-3c2e2f1152c3/providers/Microsoft.Orbital/availableGroundStations/EASTUS_1"),
		// 			Location: to.Ptr("eastus2"),
		// 			Properties: &armorbital.AvailableGroundStationProperties{
		// 				AltitudeMeters: to.Ptr[float32](128),
		// 				City: to.Ptr("Reston"),
		// 				LatitudeDegrees: to.Ptr[float32](38.9586307),
		// 				LongitudeDegrees: to.Ptr[float32](-77.357002),
		// 				ProviderName: to.Ptr("Microsoft"),
		// 			},
		// 	}},
		// }
	}
}
