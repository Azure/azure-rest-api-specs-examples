package armconsumption_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/consumption/armconsumption/v2"
)

// Generated from example definition: 2024-08-01/ReservationRecommendationsByBillingProfile.json
func ExampleReservationRecommendationsClient_NewListPager_reservationRecommendationsByBillingProfileModern() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armconsumption.NewClientFactory("<subscriptionID>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewReservationRecommendationsClient().NewListPager("providers/Microsoft.Billing/billingAccounts/123456/billingProfiles/6420", nil)
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
		// 			&armconsumption.ModernReservationRecommendation{
		// 				Name: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 				Type: to.Ptr("Microsoft.Consumption/reservationRecommendations"),
		// 				ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/123456/billingProfiles/6420/providers/Microsoft.Consumption/reservationRecommendations/00000000-0000-0000-0000-000000000000"),
		// 				Kind: to.Ptr(armconsumption.ReservationRecommendationKindModern),
		// 				Properties: &armconsumption.ModernSingleScopeReservationRecommendationProperties{
		// 					CostWithNoReservedInstances: &armconsumption.Amount{
		// 						Currency: to.Ptr("USD"),
		// 					},
		// 					FirstUsageDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-08-22T00:00:00Z"); return t}()),
		// 					InstanceFlexibilityGroup: to.Ptr("NA"),
		// 					InstanceFlexibilityRatio: to.Ptr[float32](1),
		// 					LastUsageDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-09-22T00:00:00Z"); return t}()),
		// 					Location: to.Ptr("westus2"),
		// 					LookBackPeriod: to.Ptr[int32](60),
		// 					MeterID: to.Ptr("30f7049a-b092-42f4-9173-9ec31ab945ad"),
		// 					NetSavings: &armconsumption.Amount{
		// 						Currency: to.Ptr("USD"),
		// 					},
		// 					NormalizedSize: to.Ptr("SQLDB_BC_Compute_Gen5"),
		// 					RecommendedQuantityNormalized: to.Ptr[float32](35),
		// 					ResourceType: to.Ptr("sqldatabases"),
		// 					Scope: to.Ptr("Single"),
		// 					SKUName: to.Ptr("SQLDB_BC_Compute_Gen5"),
		// 					SubscriptionID: to.Ptr("c6aa8a01-a744-44a7-a4f1-caad17512f27"),
		// 					Term: to.Ptr("P3Y"),
		// 					TotalCostWithReservedInstances: &armconsumption.Amount{
		// 						Currency: to.Ptr("USD"),
		// 					},
		// 					TotalHours: to.Ptr[int32](717),
		// 				},
		// 			},
		// 			&armconsumption.ModernReservationRecommendation{
		// 				Name: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 				Type: to.Ptr("Microsoft.Consumption/reservationRecommendations"),
		// 				ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/123456/billingProfiles/6420/providers/Microsoft.Consumption/reservationRecommendations/00000000-0000-0000-0000-000000000000"),
		// 				Kind: to.Ptr(armconsumption.ReservationRecommendationKindModern),
		// 				Properties: &armconsumption.ModernSingleScopeReservationRecommendationProperties{
		// 					CostWithNoReservedInstances: &armconsumption.Amount{
		// 						Currency: to.Ptr("USD"),
		// 					},
		// 					FirstUsageDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-08-22T00:00:00Z"); return t}()),
		// 					InstanceFlexibilityGroup: to.Ptr("NA"),
		// 					InstanceFlexibilityRatio: to.Ptr[float32](1),
		// 					LastUsageDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-08-27T00:00:00Z"); return t}()),
		// 					Location: to.Ptr("westus2"),
		// 					LookBackPeriod: to.Ptr[int32](60),
		// 					MeterID: to.Ptr("30f7049a-b092-42f4-9173-9ec31ab945ad"),
		// 					NetSavings: &armconsumption.Amount{
		// 						Currency: to.Ptr("USD"),
		// 					},
		// 					NormalizedSize: to.Ptr("SQLDB_BC_Compute_Gen5"),
		// 					RecommendedQuantityNormalized: to.Ptr[float32](35),
		// 					ResourceType: to.Ptr("sqldatabases"),
		// 					Scope: to.Ptr("Single"),
		// 					SKUName: to.Ptr("SQLDB_BC_Compute_Gen5"),
		// 					SubscriptionID: to.Ptr("c6aa8a01-a744-44a7-a4f1-caad17512f27"),
		// 					Term: to.Ptr("P1Y"),
		// 					TotalCostWithReservedInstances: &armconsumption.Amount{
		// 						Currency: to.Ptr("USD"),
		// 					},
		// 					TotalHours: to.Ptr[int32](527),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
