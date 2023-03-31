package armconsumption_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/consumption/armconsumption"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/ReservationRecommendationDetailsByBillingProfile.json
func ExampleReservationRecommendationDetailsClient_Get_reservationRecommendationsByBillingProfileModern() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armconsumption.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewReservationRecommendationDetailsClient().Get(ctx, "providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-00000000:00000000-0000-0000-0000-00000000/billingProfiles/00000000-0000-0000-0000-00000000", armconsumption.ScopeShared, "australiaeast", armconsumption.TermP1Y, armconsumption.LookBackPeriodLast07Days, "Standard_B2s", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ReservationRecommendationDetailsModel = armconsumption.ReservationRecommendationDetailsModel{
	// 	Name: to.Ptr("reservationRecommendationDetails"),
	// 	Type: to.Ptr("Microsoft.Consumption/ReservationRecommendationDetails"),
	// 	ID: to.Ptr("providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-00000000:00000000-0000-0000-0000-00000000/billingProfiles/00000000-0000-0000-0000-00000000/providers/microsoft.consumption/reservationrecommendationdetails"),
	// 	Properties: &armconsumption.ReservationRecommendationDetailsProperties{
	// 		Currency: to.Ptr("AUD"),
	// 		Resource: &armconsumption.ReservationRecommendationDetailsResourceProperties{
	// 			AppliedScopes: []*string{
	// 				to.Ptr("00000000-0000-0000-0000-00000000")},
	// 				OnDemandRate: to.Ptr[float32](0.0725),
	// 				Product: to.Ptr("Standard_B2s"),
	// 				Region: to.Ptr("australiaeast"),
	// 				ReservationRate: to.Ptr[float32](0.04414166531788041),
	// 				ResourceType: to.Ptr("virtualmachines"),
	// 			},
	// 			Savings: &armconsumption.ReservationRecommendationDetailsSavingsProperties{
	// 				CalculatedSavings: []*armconsumption.ReservationRecommendationDetailsCalculatedSavingsProperties{
	// 					{
	// 						OnDemandCost: to.Ptr[float32](632.8844720496894),
	// 						OverageCost: to.Ptr[float32](0),
	// 						Quantity: to.Ptr[float32](1),
	// 						ReservationCost: to.Ptr[float32](387.74038815226174),
	// 						Savings: to.Ptr[float32](245.1440838974277),
	// 						TotalReservationCost: to.Ptr[float32](387.74038815226174),
	// 				}},
	// 				LookBackPeriod: to.Ptr[int32](7),
	// 				RecommendedQuantity: to.Ptr[float32](1),
	// 				ReservationOrderTerm: to.Ptr("P1Y"),
	// 				SavingsType: to.Ptr("instance"),
	// 				UnitOfMeasure: to.Ptr("hour"),
	// 			},
	// 			Scope: to.Ptr("Shared"),
	// 			Usage: &armconsumption.ReservationRecommendationDetailsUsageProperties{
	// 				FirstConsumptionDate: to.Ptr("2020-01-19T00:00:00"),
	// 				LastConsumptionDate: to.Ptr("2020-01-25T17:00:00"),
	// 				LookBackUnitType: to.Ptr("virtualMachine quantity"),
	// 				UsageData: []*float32{
	// 					to.Ptr[float32](1),
	// 					to.Ptr[float32](1),
	// 					to.Ptr[float32](1),
	// 					to.Ptr[float32](1),
	// 					to.Ptr[float32](1),
	// 					to.Ptr[float32](1),
	// 					to.Ptr[float32](1),
	// 					to.Ptr[float32](1),
	// 					to.Ptr[float32](1),
	// 					to.Ptr[float32](0)},
	// 					UsageGrain: to.Ptr("hourly"),
	// 				},
	// 			},
	// 		}
}
