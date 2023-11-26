package armconsumption_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/consumption/armconsumption"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/LotsListByBillingProfile.json
func ExampleLotsClient_NewListByBillingProfilePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armconsumption.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewLotsClient().NewListByBillingProfilePager("1234:5678", "2468", nil)
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
		// page.Lots = armconsumption.Lots{
		// 	Value: []*armconsumption.LotSummary{
		// 		{
		// 			Name: to.Ptr("lot1"),
		// 			Type: to.Ptr("Microsoft.Consumption/lots"),
		// 			ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/1234:5678/billingProfiles/2468/providers/Microsoft.Consumption/lots/lot1"),
		// 			Properties: &armconsumption.LotProperties{
		// 				BillingCurrency: to.Ptr("USD"),
		// 				ClosedBalance: &armconsumption.Amount{
		// 					Currency: to.Ptr("USD"),
		// 					Value: to.Ptr[float64](60.9),
		// 				},
		// 				ClosedBalanceInBillingCurrency: &armconsumption.AmountWithExchangeRate{
		// 					ExchangeRate: to.Ptr[float64](5000),
		// 					ExchangeRateMonth: to.Ptr[int32](1),
		// 				},
		// 				CreditCurrency: to.Ptr("USD"),
		// 				ExpirationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-10-01T00:00:00.000Z"); return t}()),
		// 				OriginalAmount: &armconsumption.Amount{
		// 					Currency: to.Ptr("USD"),
		// 					Value: to.Ptr[float64](5000),
		// 				},
		// 				OriginalAmountInBillingCurrency: &armconsumption.AmountWithExchangeRate{
		// 					ExchangeRate: to.Ptr[float64](5000),
		// 					ExchangeRateMonth: to.Ptr[int32](1),
		// 				},
		// 				PoNumber: to.Ptr("3524"),
		// 				Reseller: &armconsumption.Reseller{
		// 					ResellerDescription: to.Ptr("Reseller information."),
		// 					ResellerID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/1234:5678/billingProfiles/2468/providers/Microsoft.Consumption/reseller/reseller1"),
		// 				},
		// 				Source: to.Ptr(armconsumption.LotSourcePurchasedCredit),
		// 				StartDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-10-01T00:00:00.000Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("lot2"),
		// 			Type: to.Ptr("Microsoft.Consumption/lots"),
		// 			ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/1234:5678/billingProfiles/2468/providers/Microsoft.Consumption/lots/lot2"),
		// 			Properties: &armconsumption.LotProperties{
		// 				BillingCurrency: to.Ptr("USD"),
		// 				ClosedBalance: &armconsumption.Amount{
		// 					Currency: to.Ptr("USD"),
		// 					Value: to.Ptr[float64](80.9),
		// 				},
		// 				ClosedBalanceInBillingCurrency: &armconsumption.AmountWithExchangeRate{
		// 					ExchangeRate: to.Ptr[float64](5000),
		// 					ExchangeRateMonth: to.Ptr[int32](1),
		// 				},
		// 				CreditCurrency: to.Ptr("USD"),
		// 				ExpirationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-12-31T00:00:00.000Z"); return t}()),
		// 				OriginalAmount: &armconsumption.Amount{
		// 					Currency: to.Ptr("USD"),
		// 					Value: to.Ptr[float64](6000),
		// 				},
		// 				OriginalAmountInBillingCurrency: &armconsumption.AmountWithExchangeRate{
		// 					ExchangeRate: to.Ptr[float64](5000),
		// 					ExchangeRateMonth: to.Ptr[int32](1),
		// 				},
		// 				PoNumber: to.Ptr("31224"),
		// 				Reseller: &armconsumption.Reseller{
		// 					ResellerDescription: to.Ptr("Reseller information."),
		// 					ResellerID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/1234:5678/billingProfiles/2468/providers/Microsoft.Consumption/reseller/reseller2"),
		// 				},
		// 				Source: to.Ptr(armconsumption.LotSourcePurchasedCredit),
		// 				StartDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-10-01T00:00:00.000Z"); return t}()),
		// 			},
		// 	}},
		// }
	}
}
