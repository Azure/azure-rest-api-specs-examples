package armconsumption_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/consumption/armconsumption/v2"
)

// Generated from example definition: 2024-08-01/ReservationRecommendationsBySubscription.json
func ExampleReservationRecommendationsClient_NewListPager_reservationRecommendationsBySubscriptionLegacy() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armconsumption.NewClientFactory("<subscriptionID>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewReservationRecommendationsClient().NewListPager("subscriptions/00000000-0000-0000-0000-000000000000", nil)
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
		// page = armconsumption.ReservationRecommendationsClientListResponse{
		// 	ReservationRecommendationsListResult: armconsumption.ReservationRecommendationsListResult{
		// 		NextLink: to.Ptr("https://management.azure.com/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Consumption/reservationRecommendations?api-version=2024-08-01&$skiptoken=AQAAAA%3D%3D&"),
		// 		Value: []armconsumption.ReservationRecommendationClassification{
		// 			&armconsumption.LegacyReservationRecommendation{
		// 				Name: to.Ptr("reservationRecommendations1"),
		// 				Type: to.Ptr("Microsoft.Consumption/reservationRecommendations"),
		// 				ID: to.Ptr("subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Consumption/reservationRecommendations/reservationRecommendations1"),
		// 				Kind: to.Ptr(armconsumption.ReservationRecommendationKindLegacy),
		// 				Location: to.Ptr("northeurope"),
		// 				Properties: &armconsumption.LegacySingleScopeReservationRecommendationProperties{
		// 					FirstUsageDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-03-06T00:00:00Z"); return t}()),
		// 					LastUsageDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-03-10T00:00:00Z"); return t}()),
		// 					LookBackPeriod: to.Ptr("Last7Days"),
		// 					MeterID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 					Scope: to.Ptr("Single"),
		// 					SubscriptionID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 					Term: to.Ptr("P1Y"),
		// 					TotalHours: to.Ptr[int32](827),
		// 				},
		// 				SKU: to.Ptr("Standard_DS1_v2"),
		// 			},
		// 			&armconsumption.LegacyReservationRecommendation{
		// 				Name: to.Ptr("reservationRecommendations2"),
		// 				Type: to.Ptr("Microsoft.Consumption/reservationRecommendations"),
		// 				ID: to.Ptr("subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Consumption/reservationRecommendations/reservationRecommendations2"),
		// 				Kind: to.Ptr(armconsumption.ReservationRecommendationKindLegacy),
		// 				Location: to.Ptr("northeurope"),
		// 				Properties: &armconsumption.LegacySingleScopeReservationRecommendationProperties{
		// 					FirstUsageDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-03-06T00:00:00Z"); return t}()),
		// 					LastUsageDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-03-15T00:00:00Z"); return t}()),
		// 					LookBackPeriod: to.Ptr("Last7Days"),
		// 					MeterID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 					Scope: to.Ptr("Single"),
		// 					SubscriptionID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 					Term: to.Ptr("P3Y"),
		// 					TotalHours: to.Ptr[int32](927),
		// 				},
		// 				SKU: to.Ptr("Standard_DS1_v2"),
		// 			},
		// 		},
		// 	},
		// }
	}
}
