package armconsumption_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/consumption/armconsumption/v2"
)

// Generated from example definition: 2024-08-01/ReservationRecommendationsByBillingAccountFilterByScope.json
func ExampleReservationRecommendationsClient_NewListPager_reservationRecommendationsByBillingAccountFilterForScope() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armconsumption.NewClientFactory("<subscriptionID>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewReservationRecommendationsClient().NewListPager("providers/Microsoft.Billing/billingAccounts/123456", &armconsumption.ReservationRecommendationsClientListOptions{
		Filter: to.Ptr("properties/scope eq 'Single'")})
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
		// 		Value: []armconsumption.ReservationRecommendationClassification{
		// 			&armconsumption.LegacyReservationRecommendation{
		// 				Name: to.Ptr("71fd91a7-13b9-4ade-bb11-85cfd6422d9d"),
		// 				Type: to.Ptr("Microsoft.Consumption/reservationRecommendations"),
		// 				ID: to.Ptr("billingAccount/123456/providers/Microsoft.Consumption/reservationRecommendations/71fd91a7-13b9-4ade-bb11-85cfd6422d9d"),
		// 				Kind: to.Ptr(armconsumption.ReservationRecommendationKindLegacy),
		// 				Location: to.Ptr("westus3"),
		// 				Properties: &armconsumption.LegacySingleScopeReservationRecommendationProperties{
		// 					CostWithNoReservedInstances: to.Ptr[float64](0.332),
		// 					FirstUsageDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-04-30T07:00:00Z"); return t}()),
		// 					InstanceFlexibilityGroup: to.Ptr("BS Series High Memory"),
		// 					InstanceFlexibilityRatio: to.Ptr[float32](8),
		// 					LookBackPeriod: to.Ptr("Last7Days"),
		// 					MeterID: to.Ptr("9a3781ce-d0dc-5f76-99d7-29eb5aec447f"),
		// 					NetSavings: to.Ptr[float64](0.13725114155251145),
		// 					NormalizedSize: to.Ptr("Standard_B1ms"),
		// 					RecommendedQuantity: to.Ptr[float64](1),
		// 					RecommendedQuantityNormalized: to.Ptr[float32](8),
		// 					ResourceType: to.Ptr("virtualmachines"),
		// 					Scope: to.Ptr("Single"),
		// 					SKUProperties: []*armconsumption.SKUProperty{
		// 					},
		// 					SubscriptionID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 					Term: to.Ptr("P1Y"),
		// 					TotalCostWithReservedInstances: to.Ptr[float64](0.19474885844748857),
		// 				},
		// 				SKU: to.Ptr("Standard_B4ms"),
		// 			},
		// 			&armconsumption.LegacyReservationRecommendation{
		// 				Name: to.Ptr("904b99c2-baf3-4bff-98ff-a96238ccbc96"),
		// 				Type: to.Ptr("Microsoft.Consumption/reservationRecommendations"),
		// 				ID: to.Ptr("billingAccount/123456/providers/Microsoft.Consumption/reservationRecommendations/904b99c2-baf3-4bff-98ff-a96238ccbc96"),
		// 				Kind: to.Ptr(armconsumption.ReservationRecommendationKindLegacy),
		// 				Location: to.Ptr("westus3"),
		// 				Properties: &armconsumption.LegacySingleScopeReservationRecommendationProperties{
		// 					CostWithNoReservedInstances: to.Ptr[float64](0.332),
		// 					FirstUsageDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-04-30T07:00:00Z"); return t}()),
		// 					InstanceFlexibilityGroup: to.Ptr("BS Series High Memory"),
		// 					InstanceFlexibilityRatio: to.Ptr[float32](8),
		// 					LookBackPeriod: to.Ptr("Last7Days"),
		// 					MeterID: to.Ptr("9a3781ce-d0dc-5f76-99d7-29eb5aec447f"),
		// 					NetSavings: to.Ptr[float64](0.20688584474885846),
		// 					NormalizedSize: to.Ptr("Standard_B1ms"),
		// 					RecommendedQuantity: to.Ptr[float64](1),
		// 					RecommendedQuantityNormalized: to.Ptr[float32](8),
		// 					ResourceType: to.Ptr("virtualmachines"),
		// 					Scope: to.Ptr("Single"),
		// 					SKUProperties: []*armconsumption.SKUProperty{
		// 					},
		// 					SubscriptionID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 					Term: to.Ptr("P3Y"),
		// 					TotalCostWithReservedInstances: to.Ptr[float64](0.12511415525114156),
		// 				},
		// 				SKU: to.Ptr("Standard_B4ms"),
		// 			},
		// 		},
		// 	},
		// }
	}
}
