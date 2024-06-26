package armconsumption_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/consumption/armconsumption"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/ReservationRecommendationsByBillingProfile.json
func ExampleReservationRecommendationsClient_NewListPager_reservationRecommendationsByBillingProfileModern() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armconsumption.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewReservationRecommendationsClient().NewListPager("providers/Microsoft.Billing/billingAccounts/123456/billingProfiles/6420", &armconsumption.ReservationRecommendationsClientListOptions{Filter: nil})
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
		// 		&armconsumption.ModernReservationRecommendation{
		// 			Name: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 			Type: to.Ptr("Microsoft.Consumption/reservationRecommendations"),
		// 			ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/123456/billingProfiles/6420/providers/Microsoft.Consumption/reservationRecommendations/00000000-0000-0000-0000-000000000000"),
		// 			Location: to.Ptr("westus"),
		// 			SKU: to.Ptr("Standard_DS1_v2"),
		// 			Kind: to.Ptr(armconsumption.ReservationRecommendationKindModern),
		// 			Properties: &armconsumption.ModernReservationRecommendationProperties{
		// 				CostWithNoReservedInstances: &armconsumption.Amount{
		// 					Currency: to.Ptr("USD"),
		// 					Value: to.Ptr[float64](1.4749000882),
		// 				},
		// 				FirstUsageDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-07-07T07:00:00.000Z"); return t}()),
		// 				InstanceFlexibilityGroup: to.Ptr("DSv2 Series"),
		// 				InstanceFlexibilityRatio: to.Ptr[float32](1),
		// 				LookBackPeriod: to.Ptr[int32](7),
		// 				MeterID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 				NetSavings: &armconsumption.Amount{
		// 					Currency: to.Ptr("USD"),
		// 					Value: to.Ptr[float64](0.39824251795372),
		// 				},
		// 				NormalizedSize: to.Ptr("Standard_DS1_v2"),
		// 				RecommendedQuantity: to.Ptr[float64](1),
		// 				RecommendedQuantityNormalized: to.Ptr[float32](1),
		// 				Scope: to.Ptr("Shared"),
		// 				SKUProperties: []*armconsumption.SKUProperty{
		// 					{
		// 						Name: to.Ptr("Cores"),
		// 						Value: to.Ptr("1"),
		// 					},
		// 					{
		// 						Name: to.Ptr("Ram"),
		// 						Value: to.Ptr("1"),
		// 				}},
		// 				Term: to.Ptr("P1Y"),
		// 				TotalCostWithReservedInstances: &armconsumption.Amount{
		// 					Currency: to.Ptr("USD"),
		// 					Value: to.Ptr[float64](1.07665757024628),
		// 				},
		// 			},
		// 		},
		// 		&armconsumption.ModernReservationRecommendation{
		// 			Name: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 			Type: to.Ptr("Microsoft.Consumption/reservationRecommendations"),
		// 			ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/123456/billingProfiles/6420/providers/Microsoft.Consumption/reservationRecommendations/00000000-0000-0000-0000-000000000000"),
		// 			Location: to.Ptr("westus"),
		// 			SKU: to.Ptr("Standard_DS1_v2"),
		// 			Kind: to.Ptr(armconsumption.ReservationRecommendationKindModern),
		// 			Properties: &armconsumption.ModernReservationRecommendationProperties{
		// 				CostWithNoReservedInstances: &armconsumption.Amount{
		// 					Currency: to.Ptr("USD"),
		// 					Value: to.Ptr[float64](1.4896),
		// 				},
		// 				FirstUsageDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-07-07T07:00:00.000Z"); return t}()),
		// 				InstanceFlexibilityGroup: to.Ptr("DSv2 Series"),
		// 				InstanceFlexibilityRatio: to.Ptr[float32](1),
		// 				LookBackPeriod: to.Ptr[int32](7),
		// 				MeterID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 				NetSavings: &armconsumption.Amount{
		// 					Currency: to.Ptr("USD"),
		// 					Value: to.Ptr[float64](0.40694246575372),
		// 				},
		// 				NormalizedSize: to.Ptr("Standard_DS1"),
		// 				RecommendedQuantity: to.Ptr[float64](1),
		// 				RecommendedQuantityNormalized: to.Ptr[float32](1.2),
		// 				Scope: to.Ptr("Shared"),
		// 				SKUProperties: []*armconsumption.SKUProperty{
		// 					{
		// 						Name: to.Ptr("SkuDisplayName"),
		// 						Value: to.Ptr("B"),
		// 					},
		// 					{
		// 						Name: to.Ptr("CPU"),
		// 						Value: to.Ptr("1"),
		// 				}},
		// 				Term: to.Ptr("P1Y"),
		// 				TotalCostWithReservedInstances: &armconsumption.Amount{
		// 					Currency: to.Ptr("USD"),
		// 					Value: to.Ptr[float64](1.08265753424628),
		// 				},
		// 			},
		// 	}},
		// }
	}
}
