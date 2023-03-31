package armconsumption_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/consumption/armconsumption"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/ReservationRecommendationDetailsByResourceGroup.json
func ExampleReservationRecommendationDetailsClient_Get_reservationRecommendationsByResourceGroupLegacy() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armconsumption.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewReservationRecommendationDetailsClient().Get(ctx, "subscriptions/00000000-0000-0000-0000-00000000/resourceGroups/testGroup", armconsumption.ScopeSingle, "westus", armconsumption.TermP3Y, armconsumption.LookBackPeriodLast30Days, "Standard_DS13_v2", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ReservationRecommendationDetailsModel = armconsumption.ReservationRecommendationDetailsModel{
	// 	Name: to.Ptr("reservationRecommendationDetails"),
	// 	Type: to.Ptr("Microsoft.Consumption/ReservationRecommendationDetails"),
	// 	ID: to.Ptr("subscriptions/00000000-0000-0000-0000-00000000/resourceGroups/testGroup/providers/microsoft.consumption/reservationrecommendationdetails"),
	// 	Properties: &armconsumption.ReservationRecommendationDetailsProperties{
	// 		Currency: to.Ptr("USD"),
	// 		Resource: &armconsumption.ReservationRecommendationDetailsResourceProperties{
	// 			AppliedScopes: []*string{
	// 				to.Ptr("00000000-0000-0000-0000-00000000"),
	// 				to.Ptr("testGroup")},
	// 				OnDemandRate: to.Ptr[float32](0.519),
	// 				Product: to.Ptr("Standard_DS13_v2"),
	// 				Region: to.Ptr("westus"),
	// 				ReservationRate: to.Ptr[float32](0.302549467275493),
	// 				ResourceType: to.Ptr("virtualmachines"),
	// 			},
	// 			ResourceGroup: to.Ptr("testGroup"),
	// 			Savings: &armconsumption.ReservationRecommendationDetailsSavingsProperties{
	// 				CalculatedSavings: []*armconsumption.ReservationRecommendationDetailsCalculatedSavingsProperties{
	// 					{
	// 						OnDemandCost: to.Ptr[float32](368.4813602070006),
	// 						OverageCost: to.Ptr[float32](0),
	// 						Quantity: to.Ptr[float32](2),
	// 						ReservationCost: to.Ptr[float32](429.01514459665),
	// 						Savings: to.Ptr[float32](-60.5337843896494),
	// 						TotalReservationCost: to.Ptr[float32](429.01514459665),
	// 					},
	// 					{
	// 						OnDemandCost: to.Ptr[float32](368.481360207),
	// 						OverageCost: to.Ptr[float32](1.557),
	// 						Quantity: to.Ptr[float32](1),
	// 						ReservationCost: to.Ptr[float32](214.507572298325),
	// 						Savings: to.Ptr[float32](152.416787908675),
	// 						TotalReservationCost: to.Ptr[float32](216.064572298325),
	// 				}},
	// 				LookBackPeriod: to.Ptr[int32](30),
	// 				RecommendedQuantity: to.Ptr[float32](1),
	// 				ReservationOrderTerm: to.Ptr("P3Y"),
	// 				SavingsType: to.Ptr("instance"),
	// 				UnitOfMeasure: to.Ptr("hour"),
	// 			},
	// 			Scope: to.Ptr("Single"),
	// 			Usage: &armconsumption.ReservationRecommendationDetailsUsageProperties{
	// 				FirstConsumptionDate: to.Ptr("2020-02-03T00:00:00"),
	// 				LastConsumptionDate: to.Ptr("2020-03-03T13:00:00"),
	// 				LookBackUnitType: to.Ptr("virtualMachine quantity"),
	// 				UsageData: []*float32{
	// 					to.Ptr[float32](1),
	// 					to.Ptr[float32](1),
	// 					to.Ptr[float32](1),
	// 					to.Ptr[float32](1),
	// 					to.Ptr[float32](1),
	// 					to.Ptr[float32](1)},
	// 					UsageGrain: to.Ptr("hourly"),
	// 				},
	// 			},
	// 		}
}
