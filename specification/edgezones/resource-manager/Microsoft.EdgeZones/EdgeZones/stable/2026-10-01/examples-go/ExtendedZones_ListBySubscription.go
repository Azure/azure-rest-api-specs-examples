package armedgezones_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/edgezones/armedgezones"
)

// Generated from example definition: 2026-10-01/ExtendedZones_ListBySubscription.json
func ExampleExtendedZonesClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armedgezones.NewClientFactory("a1ffc958-d2c7-493e-9f1e-125a0477f536", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewExtendedZonesClient().NewListBySubscriptionPager(nil)
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
		// page = armedgezones.ExtendedZonesClientListBySubscriptionResponse{
		// 	ExtendedZoneListResult: armedgezones.ExtendedZoneListResult{
		// 		Value: []*armedgezones.ExtendedZone{
		// 			{
		// 				ID: to.Ptr("/subscriptions/a1ffc958-d2c7-493e-9f1e-125a0477f536/providers/Microsoft.EdgeZones/extendedZones/redmond"),
		// 				Name: to.Ptr("redmond"),
		// 				Type: to.Ptr("Microsoft.EdgeZones/extendedZones"),
		// 				Properties: &armedgezones.ExtendedZoneProperties{
		// 					ProvisioningState: to.Ptr(armedgezones.ProvisioningStateSucceeded),
		// 					RegistrationState: to.Ptr(armedgezones.RegistrationStateNotRegistered),
		// 					DisplayName: to.Ptr("Redmond"),
		// 					RegionalDisplayName: to.Ptr("(US) Redmond"),
		// 					RegionType: to.Ptr("Physical"),
		// 					RegionCategory: to.Ptr("Other"),
		// 					Geography: to.Ptr("usa"),
		// 					GeographyGroup: to.Ptr("US"),
		// 					Longitude: to.Ptr("-122.03197"),
		// 					Latitude: to.Ptr("47.69106"),
		// 					HomeLocation: to.Ptr("westus"),
		// 				},
		// 			},
		// 			{
		// 				ID: to.Ptr("/subscriptions/a1ffc958-d2c7-493e-9f1e-125a0477f536/providers/Microsoft.EdgeZones/extendedZones/losangeles"),
		// 				Name: to.Ptr("losangeles"),
		// 				Type: to.Ptr("Microsoft.EdgeZones/extendedZones"),
		// 				Properties: &armedgezones.ExtendedZoneProperties{
		// 					ProvisioningState: to.Ptr(armedgezones.ProvisioningStateSucceeded),
		// 					RegistrationState: to.Ptr(armedgezones.RegistrationStateNotRegistered),
		// 					DisplayName: to.Ptr("Los Angeles"),
		// 					RegionalDisplayName: to.Ptr("(US) Los Angeles"),
		// 					RegionType: to.Ptr("Physical"),
		// 					RegionCategory: to.Ptr("Other"),
		// 					Geography: to.Ptr("usa"),
		// 					GeographyGroup: to.Ptr("US"),
		// 					Longitude: to.Ptr("-118.23537"),
		// 					Latitude: to.Ptr("34.058414"),
		// 					HomeLocation: to.Ptr("westus"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
