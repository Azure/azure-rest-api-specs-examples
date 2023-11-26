package armconsumption_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/consumption/armconsumption"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/ReservationRecommendationsBySubscription.json
func ExampleReservationRecommendationsClient_NewListPager_reservationRecommendationsBySubscriptionLegacy() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armconsumption.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewReservationRecommendationsClient().NewListPager("subscriptions/00000000-0000-0000-0000-000000000000", &armconsumption.ReservationRecommendationsClientListOptions{Filter: nil})
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
		// page.ReservationRecommendationsListResult = armconsumption.ReservationRecommendationsListResult{
		// 	Value: []armconsumption.ReservationRecommendationClassification{
		// 		&armconsumption.LegacyReservationRecommendation{
		// 			Name: to.Ptr("reservationRecommendations1"),
		// 			Type: to.Ptr("Microsoft.Consumption/reservationRecommendations"),
		// 			ID: to.Ptr("subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Consumption/reservationRecommendations/reservationRecommendations1"),
		// 			Location: to.Ptr("northeurope"),
		// 			SKU: to.Ptr("Standard_DS1_v2"),
		// 			Kind: to.Ptr(armconsumption.ReservationRecommendationKindLegacy),
		// 			Properties: &armconsumption.LegacySingleScopeReservationRecommendationProperties{
		// 				CostWithNoReservedInstances: to.Ptr[float64](0),
		// 				FirstUsageDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-03-06T00:00:00.000Z"); return t}()),
		// 				LookBackPeriod: to.Ptr("Last7Days"),
		// 				MeterID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 				NetSavings: to.Ptr[float64](4.634521202630137),
		// 				RecommendedQuantity: to.Ptr[float64](1),
		// 				Scope: to.Ptr("Single"),
		// 				Term: to.Ptr("P1Y"),
		// 				TotalCostWithReservedInstances: to.Ptr[float64](0),
		// 				SubscriptionID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 			},
		// 		},
		// 		&armconsumption.LegacyReservationRecommendation{
		// 			Name: to.Ptr("reservationRecommendations2"),
		// 			Type: to.Ptr("Microsoft.Consumption/reservationRecommendations"),
		// 			ID: to.Ptr("subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Consumption/reservationRecommendations/reservationRecommendations2"),
		// 			Location: to.Ptr("northeurope"),
		// 			SKU: to.Ptr("Standard_DS1_v2"),
		// 			Kind: to.Ptr(armconsumption.ReservationRecommendationKindLegacy),
		// 			Properties: &armconsumption.LegacySingleScopeReservationRecommendationProperties{
		// 				CostWithNoReservedInstances: to.Ptr[float64](0),
		// 				FirstUsageDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-03-06T00:00:00.000Z"); return t}()),
		// 				LookBackPeriod: to.Ptr("Last7Days"),
		// 				MeterID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 				NetSavings: to.Ptr[float64](7.289315723178081),
		// 				RecommendedQuantity: to.Ptr[float64](1),
		// 				Scope: to.Ptr("Single"),
		// 				Term: to.Ptr("P3Y"),
		// 				TotalCostWithReservedInstances: to.Ptr[float64](0),
		// 				SubscriptionID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 			},
		// 	}},
		// }
	}
}
