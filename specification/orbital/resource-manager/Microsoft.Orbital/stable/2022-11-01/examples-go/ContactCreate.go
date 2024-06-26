package armorbital_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/orbital/armorbital/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1e7b408f3323e7f5424745718fe62c7a043a2337/specification/orbital/resource-manager/Microsoft.Orbital/stable/2022-11-01/examples/ContactCreate.json
func ExampleContactsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armorbital.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewContactsClient().BeginCreate(ctx, "contoso-Rgp", "CONTOSO_SAT", "contact1", armorbital.Contact{
		Properties: &armorbital.ContactsProperties{
			ContactProfile: &armorbital.ContactsPropertiesContactProfile{
				ID: to.Ptr("/subscriptions/c1be1141-a7c9-4aac-9608-3c2e2f1152c3/resourceGroups/contoso-Rgp/providers/Microsoft.Orbital/contactProfiles/CONTOSO-CP"),
			},
			GroundStationName:    to.Ptr("EASTUS2_0"),
			ReservationEndTime:   to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-02-22T11:10:45.000Z"); return t }()),
			ReservationStartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-02-22T10:58:30.000Z"); return t }()),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Contact = armorbital.Contact{
	// 	Name: to.Ptr("contact1"),
	// 	Type: to.Ptr("Microsoft.Orbital/spacecrafts/contacts"),
	// 	ID: to.Ptr("/subscriptions/c1be1141-a7c9-4aac-9608-3c2e2f1152c3/resourceGroups/contoso-Rgp/providers/Microsoft.Orbital/spacecrafts/CONTOSO_SAT/contacts/contact1"),
	// 	Properties: &armorbital.ContactsProperties{
	// 		ContactProfile: &armorbital.ContactsPropertiesContactProfile{
	// 			ID: to.Ptr("/subscriptions/c1be1141-a7c9-4aac-9608-3c2e2f1152c3/resourceGroups/contoso-Rgp/providers/Microsoft.Orbital/contactProfiles/CONTOSO-CP"),
	// 		},
	// 		EndAzimuthDegrees: to.Ptr[float32](345.848482),
	// 		EndElevationDegrees: to.Ptr[float32](5.048656),
	// 		GroundStationName: to.Ptr("EASTUS2_0"),
	// 		MaximumElevationDegrees: to.Ptr[float32](85.9796),
	// 		ReservationEndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-02-22T11:10:45.000Z"); return t}()),
	// 		ReservationStartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-02-22T10:58:30.000Z"); return t}()),
	// 		RxEndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-02-22T11:10:45.000Z"); return t}()),
	// 		RxStartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-02-22T10:58:30.000Z"); return t}()),
	// 		StartAzimuthDegrees: to.Ptr[float32](165.758896),
	// 		StartElevationDegrees: to.Ptr[float32](5),
	// 		Status: to.Ptr(armorbital.ContactsStatusScheduled),
	// 		TxEndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-02-22T11:10:45.000Z"); return t}()),
	// 		TxStartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-02-22T10:58:30.000Z"); return t}()),
	// 	},
	// }
}
