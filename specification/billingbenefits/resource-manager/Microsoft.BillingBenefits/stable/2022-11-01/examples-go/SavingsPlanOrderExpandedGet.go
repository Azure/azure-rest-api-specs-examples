package armbillingbenefits_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billingbenefits/armbillingbenefits/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/f790e624d0d080b89d962a3bd19c65bc6a6b2f5e/specification/billingbenefits/resource-manager/Microsoft.BillingBenefits/stable/2022-11-01/examples/SavingsPlanOrderExpandedGet.json
func ExampleSavingsPlanOrderClient_Get_savingsPlanOrderWithExpandedPaymentsGet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbillingbenefits.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSavingsPlanOrderClient().Get(ctx, "20000000-0000-0000-0000-000000000000", &armbillingbenefits.SavingsPlanOrderClientGetOptions{Expand: to.Ptr("schedule")})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SavingsPlanOrderModel = armbillingbenefits.SavingsPlanOrderModel{
	// 	Name: to.Ptr("20000000-0000-0000-0000-000000000000"),
	// 	Type: to.Ptr("Microsoft.BillingBenefits/savingsPlanOrders"),
	// 	ID: to.Ptr("/providers/microsoft.billingbenefits/savingsPlanOrders/20000000-0000-0000-0000-000000000000"),
	// 	Properties: &armbillingbenefits.SavingsPlanOrderModelProperties{
	// 		BillingAccountID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/20000000-1000-0000-0000-000000000000:20000000-0000-3000-0000-000000000000_2019-05-31"),
	// 		BillingPlan: to.Ptr(armbillingbenefits.BillingPlanP1M),
	// 		BillingProfileID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/20000000-1000-0000-0000-000000000000:20000000-0000-3000-0000-000000000000_2019-05-31/billingProfiles/KPSV-DWNE-BG7-TGB"),
	// 		BillingScopeID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/1234567/billingSubscriptions/30000000-0000-0000-0000-000000000003"),
	// 		DisplayName: to.Ptr("Compute_SavingsPlan_10-19-2022_11-01"),
	// 		ExpiryDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-10-19T18:03:55.251Z"); return t}()),
	// 		PlanInformation: &armbillingbenefits.BillingPlanInformation{
	// 			NextPaymentDueDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2022-11-19"); return t}()),
	// 			PricingCurrencyTotal: &armbillingbenefits.Price{
	// 				Amount: to.Ptr[float64](26.28),
	// 				CurrencyCode: to.Ptr("USD"),
	// 			},
	// 			StartDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2022-10-19"); return t}()),
	// 			Transactions: []*armbillingbenefits.PaymentDetail{
	// 				{
	// 					BillingCurrencyTotal: &armbillingbenefits.Price{
	// 						Amount: to.Ptr[float64](0.73),
	// 						CurrencyCode: to.Ptr("USD"),
	// 					},
	// 					DueDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2022-10-19"); return t}()),
	// 					PaymentDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2022-10-19"); return t}()),
	// 					PricingCurrencyTotal: &armbillingbenefits.Price{
	// 						Amount: to.Ptr[float64](0.73),
	// 						CurrencyCode: to.Ptr("USD"),
	// 					},
	// 					Status: to.Ptr(armbillingbenefits.PaymentStatus("Completed")),
	// 				},
	// 				{
	// 					DueDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2022-11-19"); return t}()),
	// 					PricingCurrencyTotal: &armbillingbenefits.Price{
	// 						Amount: to.Ptr[float64](0.73),
	// 						CurrencyCode: to.Ptr("USD"),
	// 					},
	// 					Status: to.Ptr(armbillingbenefits.PaymentStatusScheduled),
	// 				},
	// 				{
	// 					DueDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2022-12-19"); return t}()),
	// 					PricingCurrencyTotal: &armbillingbenefits.Price{
	// 						Amount: to.Ptr[float64](0.73),
	// 						CurrencyCode: to.Ptr("USD"),
	// 					},
	// 					Status: to.Ptr(armbillingbenefits.PaymentStatusScheduled),
	// 				},
	// 				{
	// 					DueDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2023-01-19"); return t}()),
	// 					PricingCurrencyTotal: &armbillingbenefits.Price{
	// 						Amount: to.Ptr[float64](0.73),
	// 						CurrencyCode: to.Ptr("USD"),
	// 					},
	// 					Status: to.Ptr(armbillingbenefits.PaymentStatusScheduled),
	// 				},
	// 				{
	// 					DueDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2023-02-19"); return t}()),
	// 					PricingCurrencyTotal: &armbillingbenefits.Price{
	// 						Amount: to.Ptr[float64](0.73),
	// 						CurrencyCode: to.Ptr("USD"),
	// 					},
	// 					Status: to.Ptr(armbillingbenefits.PaymentStatusScheduled),
	// 				},
	// 				{
	// 					DueDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2023-03-19"); return t}()),
	// 					PricingCurrencyTotal: &armbillingbenefits.Price{
	// 						Amount: to.Ptr[float64](0.73),
	// 						CurrencyCode: to.Ptr("USD"),
	// 					},
	// 					Status: to.Ptr(armbillingbenefits.PaymentStatusScheduled),
	// 				},
	// 				{
	// 					DueDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2023-04-19"); return t}()),
	// 					PricingCurrencyTotal: &armbillingbenefits.Price{
	// 						Amount: to.Ptr[float64](0.73),
	// 						CurrencyCode: to.Ptr("USD"),
	// 					},
	// 					Status: to.Ptr(armbillingbenefits.PaymentStatusScheduled),
	// 				},
	// 				{
	// 					DueDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2023-05-19"); return t}()),
	// 					PricingCurrencyTotal: &armbillingbenefits.Price{
	// 						Amount: to.Ptr[float64](0.73),
	// 						CurrencyCode: to.Ptr("USD"),
	// 					},
	// 					Status: to.Ptr(armbillingbenefits.PaymentStatusScheduled),
	// 				},
	// 				{
	// 					DueDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2023-06-19"); return t}()),
	// 					PricingCurrencyTotal: &armbillingbenefits.Price{
	// 						Amount: to.Ptr[float64](0.73),
	// 						CurrencyCode: to.Ptr("USD"),
	// 					},
	// 					Status: to.Ptr(armbillingbenefits.PaymentStatusScheduled),
	// 				},
	// 				{
	// 					DueDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2023-07-19"); return t}()),
	// 					PricingCurrencyTotal: &armbillingbenefits.Price{
	// 						Amount: to.Ptr[float64](0.73),
	// 						CurrencyCode: to.Ptr("USD"),
	// 					},
	// 					Status: to.Ptr(armbillingbenefits.PaymentStatusScheduled),
	// 				},
	// 				{
	// 					DueDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2023-08-19"); return t}()),
	// 					PricingCurrencyTotal: &armbillingbenefits.Price{
	// 						Amount: to.Ptr[float64](0.73),
	// 						CurrencyCode: to.Ptr("USD"),
	// 					},
	// 					Status: to.Ptr(armbillingbenefits.PaymentStatusScheduled),
	// 				},
	// 				{
	// 					DueDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2023-09-19"); return t}()),
	// 					PricingCurrencyTotal: &armbillingbenefits.Price{
	// 						Amount: to.Ptr[float64](0.73),
	// 						CurrencyCode: to.Ptr("USD"),
	// 					},
	// 					Status: to.Ptr(armbillingbenefits.PaymentStatusScheduled),
	// 				},
	// 				{
	// 					DueDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2023-10-19"); return t}()),
	// 					PricingCurrencyTotal: &armbillingbenefits.Price{
	// 						Amount: to.Ptr[float64](0.73),
	// 						CurrencyCode: to.Ptr("USD"),
	// 					},
	// 					Status: to.Ptr(armbillingbenefits.PaymentStatusScheduled),
	// 				},
	// 				{
	// 					DueDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2023-11-19"); return t}()),
	// 					PricingCurrencyTotal: &armbillingbenefits.Price{
	// 						Amount: to.Ptr[float64](0.73),
	// 						CurrencyCode: to.Ptr("USD"),
	// 					},
	// 					Status: to.Ptr(armbillingbenefits.PaymentStatusScheduled),
	// 				},
	// 				{
	// 					DueDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2023-12-19"); return t}()),
	// 					PricingCurrencyTotal: &armbillingbenefits.Price{
	// 						Amount: to.Ptr[float64](0.73),
	// 						CurrencyCode: to.Ptr("USD"),
	// 					},
	// 					Status: to.Ptr(armbillingbenefits.PaymentStatusScheduled),
	// 				},
	// 				{
	// 					DueDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2024-01-19"); return t}()),
	// 					PricingCurrencyTotal: &armbillingbenefits.Price{
	// 						Amount: to.Ptr[float64](0.73),
	// 						CurrencyCode: to.Ptr("USD"),
	// 					},
	// 					Status: to.Ptr(armbillingbenefits.PaymentStatusScheduled),
	// 				},
	// 				{
	// 					DueDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2024-02-19"); return t}()),
	// 					PricingCurrencyTotal: &armbillingbenefits.Price{
	// 						Amount: to.Ptr[float64](0.73),
	// 						CurrencyCode: to.Ptr("USD"),
	// 					},
	// 					Status: to.Ptr(armbillingbenefits.PaymentStatusScheduled),
	// 				},
	// 				{
	// 					DueDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2024-03-19"); return t}()),
	// 					PricingCurrencyTotal: &armbillingbenefits.Price{
	// 						Amount: to.Ptr[float64](0.73),
	// 						CurrencyCode: to.Ptr("USD"),
	// 					},
	// 					Status: to.Ptr(armbillingbenefits.PaymentStatusScheduled),
	// 				},
	// 				{
	// 					DueDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2024-04-19"); return t}()),
	// 					PricingCurrencyTotal: &armbillingbenefits.Price{
	// 						Amount: to.Ptr[float64](0.73),
	// 						CurrencyCode: to.Ptr("USD"),
	// 					},
	// 					Status: to.Ptr(armbillingbenefits.PaymentStatusScheduled),
	// 				},
	// 				{
	// 					DueDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2024-05-19"); return t}()),
	// 					PricingCurrencyTotal: &armbillingbenefits.Price{
	// 						Amount: to.Ptr[float64](0.73),
	// 						CurrencyCode: to.Ptr("USD"),
	// 					},
	// 					Status: to.Ptr(armbillingbenefits.PaymentStatusScheduled),
	// 				},
	// 				{
	// 					DueDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2024-06-19"); return t}()),
	// 					PricingCurrencyTotal: &armbillingbenefits.Price{
	// 						Amount: to.Ptr[float64](0.73),
	// 						CurrencyCode: to.Ptr("USD"),
	// 					},
	// 					Status: to.Ptr(armbillingbenefits.PaymentStatusScheduled),
	// 				},
	// 				{
	// 					DueDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2024-07-19"); return t}()),
	// 					PricingCurrencyTotal: &armbillingbenefits.Price{
	// 						Amount: to.Ptr[float64](0.73),
	// 						CurrencyCode: to.Ptr("USD"),
	// 					},
	// 					Status: to.Ptr(armbillingbenefits.PaymentStatusScheduled),
	// 				},
	// 				{
	// 					DueDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2024-08-19"); return t}()),
	// 					PricingCurrencyTotal: &armbillingbenefits.Price{
	// 						Amount: to.Ptr[float64](0.73),
	// 						CurrencyCode: to.Ptr("USD"),
	// 					},
	// 					Status: to.Ptr(armbillingbenefits.PaymentStatusScheduled),
	// 				},
	// 				{
	// 					DueDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2024-09-19"); return t}()),
	// 					PricingCurrencyTotal: &armbillingbenefits.Price{
	// 						Amount: to.Ptr[float64](0.73),
	// 						CurrencyCode: to.Ptr("USD"),
	// 					},
	// 					Status: to.Ptr(armbillingbenefits.PaymentStatusScheduled),
	// 				},
	// 				{
	// 					DueDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2024-10-19"); return t}()),
	// 					PricingCurrencyTotal: &armbillingbenefits.Price{
	// 						Amount: to.Ptr[float64](0.73),
	// 						CurrencyCode: to.Ptr("USD"),
	// 					},
	// 					Status: to.Ptr(armbillingbenefits.PaymentStatusScheduled),
	// 				},
	// 				{
	// 					DueDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2024-11-19"); return t}()),
	// 					PricingCurrencyTotal: &armbillingbenefits.Price{
	// 						Amount: to.Ptr[float64](0.73),
	// 						CurrencyCode: to.Ptr("USD"),
	// 					},
	// 					Status: to.Ptr(armbillingbenefits.PaymentStatusScheduled),
	// 				},
	// 				{
	// 					DueDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2024-12-19"); return t}()),
	// 					PricingCurrencyTotal: &armbillingbenefits.Price{
	// 						Amount: to.Ptr[float64](0.73),
	// 						CurrencyCode: to.Ptr("USD"),
	// 					},
	// 					Status: to.Ptr(armbillingbenefits.PaymentStatusScheduled),
	// 				},
	// 				{
	// 					DueDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2025-01-19"); return t}()),
	// 					PricingCurrencyTotal: &armbillingbenefits.Price{
	// 						Amount: to.Ptr[float64](0.73),
	// 						CurrencyCode: to.Ptr("USD"),
	// 					},
	// 					Status: to.Ptr(armbillingbenefits.PaymentStatusScheduled),
	// 				},
	// 				{
	// 					DueDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2025-02-19"); return t}()),
	// 					PricingCurrencyTotal: &armbillingbenefits.Price{
	// 						Amount: to.Ptr[float64](0.73),
	// 						CurrencyCode: to.Ptr("USD"),
	// 					},
	// 					Status: to.Ptr(armbillingbenefits.PaymentStatusScheduled),
	// 				},
	// 				{
	// 					DueDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2025-03-19"); return t}()),
	// 					PricingCurrencyTotal: &armbillingbenefits.Price{
	// 						Amount: to.Ptr[float64](0.73),
	// 						CurrencyCode: to.Ptr("USD"),
	// 					},
	// 					Status: to.Ptr(armbillingbenefits.PaymentStatusScheduled),
	// 				},
	// 				{
	// 					DueDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2025-04-19"); return t}()),
	// 					PricingCurrencyTotal: &armbillingbenefits.Price{
	// 						Amount: to.Ptr[float64](0.73),
	// 						CurrencyCode: to.Ptr("USD"),
	// 					},
	// 					Status: to.Ptr(armbillingbenefits.PaymentStatusScheduled),
	// 				},
	// 				{
	// 					DueDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2025-05-19"); return t}()),
	// 					PricingCurrencyTotal: &armbillingbenefits.Price{
	// 						Amount: to.Ptr[float64](0.73),
	// 						CurrencyCode: to.Ptr("USD"),
	// 					},
	// 					Status: to.Ptr(armbillingbenefits.PaymentStatusScheduled),
	// 				},
	// 				{
	// 					DueDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2025-06-19"); return t}()),
	// 					PricingCurrencyTotal: &armbillingbenefits.Price{
	// 						Amount: to.Ptr[float64](0.73),
	// 						CurrencyCode: to.Ptr("USD"),
	// 					},
	// 					Status: to.Ptr(armbillingbenefits.PaymentStatusScheduled),
	// 				},
	// 				{
	// 					DueDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2025-07-19"); return t}()),
	// 					PricingCurrencyTotal: &armbillingbenefits.Price{
	// 						Amount: to.Ptr[float64](0.73),
	// 						CurrencyCode: to.Ptr("USD"),
	// 					},
	// 					Status: to.Ptr(armbillingbenefits.PaymentStatusScheduled),
	// 				},
	// 				{
	// 					DueDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2025-08-19"); return t}()),
	// 					PricingCurrencyTotal: &armbillingbenefits.Price{
	// 						Amount: to.Ptr[float64](0.73),
	// 						CurrencyCode: to.Ptr("USD"),
	// 					},
	// 					Status: to.Ptr(armbillingbenefits.PaymentStatusScheduled),
	// 				},
	// 				{
	// 					DueDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2025-09-19"); return t}()),
	// 					PricingCurrencyTotal: &armbillingbenefits.Price{
	// 						Amount: to.Ptr[float64](0.73),
	// 						CurrencyCode: to.Ptr("USD"),
	// 					},
	// 					Status: to.Ptr(armbillingbenefits.PaymentStatusScheduled),
	// 			}},
	// 		},
	// 		ProvisioningState: to.Ptr(armbillingbenefits.ProvisioningStateSucceeded),
	// 		SavingsPlans: []*string{
	// 			to.Ptr("/providers/Microsoft.BillingBenefits/savingsPlanOrders/20000000-0000-0000-0000-000000000000/savingsPlans/30000000-0000-0000-0000-000000000000")},
	// 			Term: to.Ptr(armbillingbenefits.TermP3Y),
	// 		},
	// 		SKU: &armbillingbenefits.SKU{
	// 			Name: to.Ptr("Compute_Savings_Plan"),
	// 		},
	// 	}
}
