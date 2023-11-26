package armorbital_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/orbital/armorbital/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1e7b408f3323e7f5424745718fe62c7a043a2337/specification/orbital/resource-manager/Microsoft.Orbital/stable/2022-11-01/examples/AvailableContactsList.json
func ExampleSpacecraftsClient_BeginListAvailableContacts() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armorbital.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewSpacecraftsClient().BeginListAvailableContacts(ctx, "contoso-Rgp", "CONTOSO_SAT", armorbital.ContactParameters{
		ContactProfile: &armorbital.ContactParametersContactProfile{
			ID: to.Ptr("/subscriptions/c1be1141-a7c9-4aac-9608-3c2e2f1152c3/resourceGroups/contoso-Rgp/providers/Microsoft.Orbital/contactProfiles/CONTOSO-CP"),
		},
		EndTime:           to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-11-02T11:30:00.000Z"); return t }()),
		GroundStationName: to.Ptr("EASTUS2_0"),
		StartTime:         to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-11-01T11:30:00.000Z"); return t }()),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	for res.More() {
		page, err := res.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.AvailableContactsListResult = armorbital.AvailableContactsListResult{
		// 	Value: []*armorbital.AvailableContacts{
		// 		{
		// 			GroundStationName: to.Ptr("EASTUS2_0"),
		// 			Properties: &armorbital.AvailableContactsProperties{
		// 				EndAzimuthDegrees: to.Ptr[float32](330.489627),
		// 				EndElevationDegrees: to.Ptr[float32](5.040625),
		// 				MaximumElevationDegrees: to.Ptr[float32](26.617297),
		// 				RxEndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-11-01T12:05:25.000Z"); return t}()),
		// 				RxStartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-11-01T11:55:01.000Z"); return t}()),
		// 				StartAzimuthDegrees: to.Ptr[float32](201.340472),
		// 				StartElevationDegrees: to.Ptr[float32](5),
		// 				TxEndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-11-01T12:05:25.000Z"); return t}()),
		// 				TxStartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-11-01T11:55:01.000Z"); return t}()),
		// 			},
		// 			Spacecraft: &armorbital.AvailableContactsSpacecraft{
		// 				ID: to.Ptr("/subscriptions/c1be1141-a7c9-4aac-9608-3c2e2f1152c3/resourceGroups/contoso-Rgp/providers/Microsoft.Orbital/spacecrafts/CONTOSO_SAT"),
		// 			},
		// 		},
		// 		{
		// 			GroundStationName: to.Ptr("EASTUS2_0"),
		// 			Properties: &armorbital.AvailableContactsProperties{
		// 				EndAzimuthDegrees: to.Ptr[float32](345.848482),
		// 				EndElevationDegrees: to.Ptr[float32](5.048656),
		// 				MaximumElevationDegrees: to.Ptr[float32](85.9796),
		// 				RxEndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-11-02T11:10:45.000Z"); return t}()),
		// 				RxStartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-11-02T10:58:30.000Z"); return t}()),
		// 				StartAzimuthDegrees: to.Ptr[float32](165.758896),
		// 				StartElevationDegrees: to.Ptr[float32](5),
		// 				TxEndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-11-02T11:10:45.000Z"); return t}()),
		// 				TxStartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-11-02T10:58:30.000Z"); return t}()),
		// 			},
		// 			Spacecraft: &armorbital.AvailableContactsSpacecraft{
		// 				ID: to.Ptr("/subscriptions/c1be1141-a7c9-4aac-9608-3c2e2f1152c3/resourceGroups/contoso-Rgp/providers/Microsoft.Orbital/spacecrafts/CONTOSO_SAT"),
		// 			},
		// 	}},
		// }
	}
}
