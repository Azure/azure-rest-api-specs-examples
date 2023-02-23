package armreservations_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/reservations/armreservations/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/reservations/resource-manager/Microsoft.Capacity/stable/2022-11-01/examples/CalculateReservationOrder.json
func ExampleReservationOrderClient_Calculate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armreservations.NewReservationOrderClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Calculate(ctx, armreservations.PurchaseRequest{
		Location: to.Ptr("westus"),
		Properties: &armreservations.PurchaseRequestProperties{
			AppliedScopeType: to.Ptr(armreservations.AppliedScopeTypeShared),
			BillingPlan:      to.Ptr(armreservations.ReservationBillingPlanMonthly),
			BillingScopeID:   to.Ptr("/subscriptions/ed3a1871-612d-abcd-a849-c2542a68be83"),
			DisplayName:      to.Ptr("TestReservationOrder"),
			Quantity:         to.Ptr[int32](1),
			ReservedResourceProperties: &armreservations.PurchaseRequestPropertiesReservedResourceProperties{
				InstanceFlexibility: to.Ptr(armreservations.InstanceFlexibilityOn),
			},
			ReservedResourceType: to.Ptr(armreservations.ReservedResourceTypeVirtualMachines),
			Term:                 to.Ptr(armreservations.ReservationTermP1Y),
		},
		SKU: &armreservations.SKUName{
			Name: to.Ptr("standard_D1"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CalculatePriceResponse = armreservations.CalculatePriceResponse{
	// 	Properties: &armreservations.CalculatePriceResponseProperties{
	// 		BillingCurrencyTotal: &armreservations.CalculatePriceResponsePropertiesBillingCurrencyTotal{
	// 			Amount: to.Ptr[float64](46),
	// 			CurrencyCode: to.Ptr("USD"),
	// 		},
	// 		PaymentSchedule: []*armreservations.PaymentDetail{
	// 			{
	// 				BillingCurrencyTotal: &armreservations.Price{
	// 					Amount: to.Ptr[float64](40),
	// 					CurrencyCode: to.Ptr("EUR"),
	// 				},
	// 				DueDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2019-05-14"); return t}()),
	// 				PricingCurrencyTotal: &armreservations.Price{
	// 					Amount: to.Ptr[float64](46),
	// 					CurrencyCode: to.Ptr("USD"),
	// 				},
	// 			},
	// 			{
	// 				DueDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2019-06-14"); return t}()),
	// 				PricingCurrencyTotal: &armreservations.Price{
	// 					Amount: to.Ptr[float64](46),
	// 					CurrencyCode: to.Ptr("USD"),
	// 				},
	// 			},
	// 			{
	// 				DueDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2019-07-14"); return t}()),
	// 				PricingCurrencyTotal: &armreservations.Price{
	// 					Amount: to.Ptr[float64](46),
	// 					CurrencyCode: to.Ptr("USD"),
	// 				},
	// 		}},
	// 		PricingCurrencyTotal: &armreservations.CalculatePriceResponsePropertiesPricingCurrencyTotal{
	// 			Amount: to.Ptr[float32](46),
	// 			CurrencyCode: to.Ptr("USD"),
	// 		},
	// 		ReservationOrderID: to.Ptr("6d9cec54-7de8-abcd-9de7-80f5d634f2d2"),
	// 		SKUDescription: to.Ptr("standard_D1"),
	// 		SKUTitle: to.Ptr("Reserved VM Instance, Standard_D1, US West, 1 Year"),
	// 	},
	// }
}
