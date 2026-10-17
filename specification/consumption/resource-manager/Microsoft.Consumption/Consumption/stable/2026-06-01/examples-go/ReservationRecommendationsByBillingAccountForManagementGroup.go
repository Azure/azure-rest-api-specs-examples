package armconsumption_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/consumption/armconsumption/v2"
)

// Generated from example definition: 2026-06-01/ReservationRecommendationsByBillingAccountForManagementGroup.json
func ExampleReservationRecommendationsClient_NewListPager_reservationRecommendationsByBillingAccountForManagementGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armconsumption.NewClientFactory("<subscriptionID>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewReservationRecommendationsClient().NewListPager("providers/Microsoft.Billing/billingAccounts/00000000", &armconsumption.ReservationRecommendationsClientListOptions{
		Filter: to.Ptr("properties/scope eq 'ManagementGroup' and properties/managementGroupId eq '/providers/Microsoft.Management/managementGroups/00000000-0000-0000-0000-000000000000'")})
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
		// 		NextLink: to.Ptr("https://management.azure.com/providers/Microsoft.Billing/billingAccounts/00000000/providers/Microsoft.Consumption/reservationRecommendations?api-version=2026-06-01&$filter=properties/managementGroupId+eq+'/providers/Microsoft.Management/managementGroups/00000000-0000-0000-0000-000000000000'&$skiptoken=AQAAAA%3D%3D&"),
		// 		Value: []armconsumption.ReservationRecommendationClassification{
		// 			&armconsumption.LegacyReservationRecommendation{
		// 				Name: to.Ptr("reservationRecommendations1"),
		// 				Type: to.Ptr("Microsoft.Consumption/reservationRecommendations"),
		// 				ID: to.Ptr("providers/Microsoft.Billing/billingAccounts/00000000/providers/Microsoft.Consumption/reservationRecommendations/reservationRecommendations1"),
		// 				Kind: to.Ptr(armconsumption.ReservationRecommendationKindLegacy),
		// 				Location: to.Ptr("northeurope"),
		// 				Properties: &armconsumption.LegacyManagementGroupScopeReservationRecommendationProperties{
		// 					CostWithNoReservedInstances: to.Ptr[float64](0),
		// 					FirstUsageDate: to.Ptr(time.Date(2026, time.March, 6, 0, 0, 0, 0, time.UTC)),
		// 					LastUsageDate: to.Ptr(time.Date(2026, time.March, 10, 0, 0, 0, 0, time.UTC)),
		// 					LookBackPeriod: to.Ptr("Last7Days"),
		// 					ManagementGroupID: to.Ptr("/providers/Microsoft.Management/managementGroups/00000000-0000-0000-0000-000000000000"),
		// 					TenantID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 					MeterID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 					NetSavings: to.Ptr[float64](4.634521202630137),
		// 					RecommendedQuantity: to.Ptr[float64](1),
		// 					Scope: to.Ptr("ManagementGroup"),
		// 					Term: to.Ptr("P1Y"),
		// 					TotalCostWithReservedInstances: to.Ptr[float64](0),
		// 					TotalHours: to.Ptr[int32](827),
		// 				},
		// 				SKU: to.Ptr("Standard_DS1_v2"),
		// 			},
		// 			&armconsumption.ModernReservationRecommendation{
		// 				Name: to.Ptr("reservationRecommendations2"),
		// 				Type: to.Ptr("Microsoft.Consumption/reservationRecommendations"),
		// 				ID: to.Ptr("providers/Microsoft.Billing/billingAccounts/00000000/providers/Microsoft.Consumption/reservationRecommendations/reservationRecommendations2"),
		// 				Kind: to.Ptr(armconsumption.ReservationRecommendationKindModern),
		// 				Location: to.Ptr("northeurope"),
		// 				Properties: &armconsumption.ModernManagementGroupScopeReservationRecommendationProperties{
		// 					CostWithNoReservedInstances: &armconsumption.Amount{
		// 						Currency: to.Ptr("USD"),
		// 						Value: to.Ptr[float64](0),
		// 					},
		// 					FirstUsageDate: to.Ptr(time.Date(2026, time.March, 6, 0, 0, 0, 0, time.UTC)),
		// 					LastUsageDate: to.Ptr(time.Date(2026, time.March, 15, 0, 0, 0, 0, time.UTC)),
		// 					LookBackPeriod: to.Ptr[int32](7),
		// 					ManagementGroupID: to.Ptr("/providers/Microsoft.Management/managementGroups/00000000-0000-0000-0000-000000000000"),
		// 					TenantID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 					MeterID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 					NetSavings: &armconsumption.Amount{
		// 						Currency: to.Ptr("USD"),
		// 						Value: to.Ptr[float64](7.289315723178081),
		// 					},
		// 					RecommendedQuantity: to.Ptr[float64](1),
		// 					Scope: to.Ptr("ManagementGroup"),
		// 					SKUName: to.Ptr("Standard_DS1_v2"),
		// 					Term: to.Ptr("P3Y"),
		// 					TotalCostWithReservedInstances: &armconsumption.Amount{
		// 						Currency: to.Ptr("USD"),
		// 						Value: to.Ptr[float64](0),
		// 					},
		// 					TotalHours: to.Ptr[int32](927),
		// 				},
		// 				SKU: to.Ptr("Standard_DS1_v2"),
		// 			},
		// 		},
		// 	},
		// }
	}
}
