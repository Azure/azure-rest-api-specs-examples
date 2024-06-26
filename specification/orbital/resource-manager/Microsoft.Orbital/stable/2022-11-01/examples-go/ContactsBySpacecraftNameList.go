package armorbital_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/orbital/armorbital/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1e7b408f3323e7f5424745718fe62c7a043a2337/specification/orbital/resource-manager/Microsoft.Orbital/stable/2022-11-01/examples/ContactsBySpacecraftNameList.json
func ExampleContactsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armorbital.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewContactsClient().NewListPager("contoso-Rgp", "CONTOSO_SAT", &armorbital.ContactsClientListOptions{Skiptoken: nil})
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
		// page.ContactListResult = armorbital.ContactListResult{
		// 	Value: []*armorbital.Contact{
		// 		{
		// 			Name: to.Ptr("contact1"),
		// 			Type: to.Ptr("Microsoft.Orbital/spacecrafts/contacts"),
		// 			ID: to.Ptr("/subscriptions/c1be1141-a7c9-4aac-9608-3c2e2f1152c3/resourceGroups/contoso-Rgp/providers/Microsoft.Orbital/spacecrafts/CONTOSO_SAT/contacts/contact2"),
		// 			Properties: &armorbital.ContactsProperties{
		// 				AntennaConfiguration: &armorbital.ContactsPropertiesAntennaConfiguration{
		// 					DestinationIP: to.Ptr("10.1.1.171"),
		// 					SourceIPs: []*string{
		// 						to.Ptr("10.1.1.4"),
		// 						to.Ptr("10.1.1.5")},
		// 					},
		// 					ContactProfile: &armorbital.ContactsPropertiesContactProfile{
		// 						ID: to.Ptr("/subscriptions/c1be1141-a7c9-4aac-9608-3c2e2f1152c3/resourceGroups/contoso-Rgp/providers/Microsoft.Orbital/contactProfiles/CONTOSO-CP"),
		// 					},
		// 					EndAzimuthDegrees: to.Ptr[float32](330.489627),
		// 					EndElevationDegrees: to.Ptr[float32](5.040625),
		// 					GroundStationName: to.Ptr("EASTUS2_0"),
		// 					MaximumElevationDegrees: to.Ptr[float32](26.617297),
		// 					ProvisioningState: to.Ptr(armorbital.ContactsPropertiesProvisioningState("Succeeded")),
		// 					ReservationEndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-01T12:05:25.000Z"); return t}()),
		// 					ReservationStartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-01T11:55:01.000Z"); return t}()),
		// 					RxEndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-01T12:05:25.000Z"); return t}()),
		// 					RxStartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-01T11:55:01.000Z"); return t}()),
		// 					StartAzimuthDegrees: to.Ptr[float32](201.340472),
		// 					StartElevationDegrees: to.Ptr[float32](5),
		// 					Status: to.Ptr(armorbital.ContactsStatusSucceeded),
		// 					TxEndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-01T12:05:25.000Z"); return t}()),
		// 					TxStartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-01T11:55:01.000Z"); return t}()),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("contact2"),
		// 				Type: to.Ptr("Microsoft.Orbital/spacecrafts/contacts"),
		// 				ID: to.Ptr("/subscriptions/c1be1141-a7c9-4aac-9608-3c2e2f1152c3/resourceGroups/contoso-Rgp/providers/Microsoft.Orbital/spacecrafts/CONTOSO_SAT/contacts/contact1"),
		// 				Properties: &armorbital.ContactsProperties{
		// 					AntennaConfiguration: &armorbital.ContactsPropertiesAntennaConfiguration{
		// 						DestinationIP: to.Ptr("10.1.1.171"),
		// 						SourceIPs: []*string{
		// 							to.Ptr("10.1.1.4"),
		// 							to.Ptr("10.1.1.5")},
		// 						},
		// 						ContactProfile: &armorbital.ContactsPropertiesContactProfile{
		// 							ID: to.Ptr("/subscriptions/c1be1141-a7c9-4aac-9608-3c2e2f1152c3/resourceGroups/contoso-Rgp/providers/Microsoft.Orbital/contactProfiles/CONTOSO-CP"),
		// 						},
		// 						EndAzimuthDegrees: to.Ptr[float32](345.848482),
		// 						EndElevationDegrees: to.Ptr[float32](5.048656),
		// 						GroundStationName: to.Ptr("EASTUS2_0"),
		// 						MaximumElevationDegrees: to.Ptr[float32](85.9796),
		// 						ProvisioningState: to.Ptr(armorbital.ContactsPropertiesProvisioningState("Succeeded")),
		// 						ReservationEndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-02T11:10:45.000Z"); return t}()),
		// 						ReservationStartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-02T10:58:30.000Z"); return t}()),
		// 						RxEndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-02T11:10:45.000Z"); return t}()),
		// 						RxStartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-02T10:58:30.000Z"); return t}()),
		// 						StartAzimuthDegrees: to.Ptr[float32](165.758896),
		// 						StartElevationDegrees: to.Ptr[float32](5),
		// 						Status: to.Ptr(armorbital.ContactsStatusSucceeded),
		// 						TxEndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-02T11:10:45.000Z"); return t}()),
		// 						TxStartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-02T10:58:30.000Z"); return t}()),
		// 					},
		// 			}},
		// 		}
	}
}
