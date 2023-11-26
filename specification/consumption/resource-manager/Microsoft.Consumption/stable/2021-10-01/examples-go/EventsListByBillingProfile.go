package armconsumption_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/consumption/armconsumption"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/EventsListByBillingProfile.json
func ExampleEventsClient_NewListByBillingProfilePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armconsumption.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewEventsClient().NewListByBillingProfilePager("1234:5678", "4268", "2019-09-01", "2019-10-31", nil)
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
		// page.Events = armconsumption.Events{
		// 	Value: []*armconsumption.EventSummary{
		// 		{
		// 			Name: to.Ptr("event1"),
		// 			Type: to.Ptr("Microsoft.Consumption/events"),
		// 			ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/1234:5678/billingProfiles/4268/providers/Microsoft.Consumption/events/event1"),
		// 			Properties: &armconsumption.EventProperties{
		// 				Description: to.Ptr("Settled invoice #312033"),
		// 				Adjustments: &armconsumption.Amount{
		// 					Currency: to.Ptr("USD"),
		// 					Value: to.Ptr[float64](0),
		// 				},
		// 				AdjustmentsInBillingCurrency: &armconsumption.AmountWithExchangeRate{
		// 					ExchangeRate: to.Ptr[float64](5000),
		// 					ExchangeRateMonth: to.Ptr[int32](1),
		// 				},
		// 				BillingCurrency: to.Ptr("USD"),
		// 				Charges: &armconsumption.Amount{
		// 					Currency: to.Ptr("USD"),
		// 					Value: to.Ptr[float64](500),
		// 				},
		// 				ChargesInBillingCurrency: &armconsumption.AmountWithExchangeRate{
		// 					ExchangeRate: to.Ptr[float64](5000),
		// 					ExchangeRateMonth: to.Ptr[int32](1),
		// 				},
		// 				ClosedBalance: &armconsumption.Amount{
		// 					Currency: to.Ptr("USD"),
		// 					Value: to.Ptr[float64](500),
		// 				},
		// 				ClosedBalanceInBillingCurrency: &armconsumption.AmountWithExchangeRate{
		// 					ExchangeRate: to.Ptr[float64](5000),
		// 					ExchangeRateMonth: to.Ptr[int32](1),
		// 				},
		// 				CreditCurrency: to.Ptr("USD"),
		// 				CreditExpired: &armconsumption.Amount{
		// 					Currency: to.Ptr("USD"),
		// 					Value: to.Ptr[float64](0),
		// 				},
		// 				CreditExpiredInBillingCurrency: &armconsumption.AmountWithExchangeRate{
		// 					ExchangeRate: to.Ptr[float64](5000),
		// 					ExchangeRateMonth: to.Ptr[int32](1),
		// 				},
		// 				EventType: to.Ptr(armconsumption.EventTypeSettledCharges),
		// 				InvoiceNumber: to.Ptr("3301"),
		// 				NewCredit: &armconsumption.Amount{
		// 					Currency: to.Ptr("USD"),
		// 					Value: to.Ptr[float64](0),
		// 				},
		// 				NewCreditInBillingCurrency: &armconsumption.AmountWithExchangeRate{
		// 					ExchangeRate: to.Ptr[float64](5000),
		// 					ExchangeRateMonth: to.Ptr[int32](1),
		// 				},
		// 				Reseller: &armconsumption.Reseller{
		// 					ResellerDescription: to.Ptr("Reseller information"),
		// 					ResellerID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/1234:5678/billingProfiles/2468/providers/Microsoft.Consumption/reseller/reseller1"),
		// 				},
		// 				TransactionDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-07-01T00:00:00.000Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("event2"),
		// 			Type: to.Ptr("Microsoft.Consumption/events"),
		// 			ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/1234:5678/billingProfiles/4268/providers/Microsoft.Consumption/events/event2"),
		// 			Properties: &armconsumption.EventProperties{
		// 				Description: to.Ptr("New credits added"),
		// 				Adjustments: &armconsumption.Amount{
		// 					Currency: to.Ptr("USD"),
		// 					Value: to.Ptr[float64](0),
		// 				},
		// 				AdjustmentsInBillingCurrency: &armconsumption.AmountWithExchangeRate{
		// 					ExchangeRate: to.Ptr[float64](5000),
		// 					ExchangeRateMonth: to.Ptr[int32](1),
		// 				},
		// 				BillingCurrency: to.Ptr("USD"),
		// 				CanceledCredit: &armconsumption.Amount{
		// 					Currency: to.Ptr("USD"),
		// 					Value: to.Ptr[float64](5000),
		// 				},
		// 				Charges: &armconsumption.Amount{
		// 					Currency: to.Ptr("USD"),
		// 					Value: to.Ptr[float64](0),
		// 				},
		// 				ChargesInBillingCurrency: &armconsumption.AmountWithExchangeRate{
		// 					ExchangeRate: to.Ptr[float64](5000),
		// 					ExchangeRateMonth: to.Ptr[int32](1),
		// 				},
		// 				ClosedBalance: &armconsumption.Amount{
		// 					Currency: to.Ptr("USD"),
		// 					Value: to.Ptr[float64](900),
		// 				},
		// 				ClosedBalanceInBillingCurrency: &armconsumption.AmountWithExchangeRate{
		// 					ExchangeRate: to.Ptr[float64](5000),
		// 					ExchangeRateMonth: to.Ptr[int32](1),
		// 				},
		// 				CreditCurrency: to.Ptr("USD"),
		// 				CreditExpired: &armconsumption.Amount{
		// 					Currency: to.Ptr("USD"),
		// 					Value: to.Ptr[float64](0),
		// 				},
		// 				CreditExpiredInBillingCurrency: &armconsumption.AmountWithExchangeRate{
		// 					ExchangeRate: to.Ptr[float64](5000),
		// 					ExchangeRateMonth: to.Ptr[int32](1),
		// 				},
		// 				EventType: to.Ptr(armconsumption.EventTypeNewCredit),
		// 				InvoiceNumber: to.Ptr("3302"),
		// 				NewCredit: &armconsumption.Amount{
		// 					Currency: to.Ptr("USD"),
		// 					Value: to.Ptr[float64](400),
		// 				},
		// 				NewCreditInBillingCurrency: &armconsumption.AmountWithExchangeRate{
		// 					ExchangeRate: to.Ptr[float64](5000),
		// 					ExchangeRateMonth: to.Ptr[int32](1),
		// 				},
		// 				Reseller: &armconsumption.Reseller{
		// 					ResellerDescription: to.Ptr("Reseller information"),
		// 					ResellerID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/1234:5678/billingProfiles/2468/providers/Microsoft.Consumption/reseller/reseller1"),
		// 				},
		// 				TransactionDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-08-01T00:00:00.000Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("event3"),
		// 			Type: to.Ptr("Microsoft.Consumption/events"),
		// 			ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/1234:5678/billingProfiles/4268/providers/Microsoft.Consumption/events/event3"),
		// 			Properties: &armconsumption.EventProperties{
		// 				Description: to.Ptr("Credits Expired"),
		// 				Adjustments: &armconsumption.Amount{
		// 					Currency: to.Ptr("USD"),
		// 					Value: to.Ptr[float64](0),
		// 				},
		// 				AdjustmentsInBillingCurrency: &armconsumption.AmountWithExchangeRate{
		// 					ExchangeRate: to.Ptr[float64](5000),
		// 					ExchangeRateMonth: to.Ptr[int32](1),
		// 				},
		// 				BillingCurrency: to.Ptr("USD"),
		// 				CanceledCredit: &armconsumption.Amount{
		// 					Currency: to.Ptr("USD"),
		// 					Value: to.Ptr[float64](5000),
		// 				},
		// 				Charges: &armconsumption.Amount{
		// 					Currency: to.Ptr("USD"),
		// 					Value: to.Ptr[float64](0),
		// 				},
		// 				ChargesInBillingCurrency: &armconsumption.AmountWithExchangeRate{
		// 					ExchangeRate: to.Ptr[float64](5000),
		// 					ExchangeRateMonth: to.Ptr[int32](1),
		// 				},
		// 				ClosedBalance: &armconsumption.Amount{
		// 					Currency: to.Ptr("USD"),
		// 					Value: to.Ptr[float64](600),
		// 				},
		// 				ClosedBalanceInBillingCurrency: &armconsumption.AmountWithExchangeRate{
		// 					ExchangeRate: to.Ptr[float64](5000),
		// 					ExchangeRateMonth: to.Ptr[int32](1),
		// 				},
		// 				CreditCurrency: to.Ptr("USD"),
		// 				CreditExpired: &armconsumption.Amount{
		// 					Currency: to.Ptr("USD"),
		// 					Value: to.Ptr[float64](300),
		// 				},
		// 				CreditExpiredInBillingCurrency: &armconsumption.AmountWithExchangeRate{
		// 					ExchangeRate: to.Ptr[float64](5000),
		// 					ExchangeRateMonth: to.Ptr[int32](1),
		// 				},
		// 				EventType: to.Ptr(armconsumption.EventType("ExpiredCredit")),
		// 				InvoiceNumber: to.Ptr(""),
		// 				NewCredit: &armconsumption.Amount{
		// 					Currency: to.Ptr("USD"),
		// 					Value: to.Ptr[float64](0),
		// 				},
		// 				NewCreditInBillingCurrency: &armconsumption.AmountWithExchangeRate{
		// 					ExchangeRate: to.Ptr[float64](5000),
		// 					ExchangeRateMonth: to.Ptr[int32](1),
		// 				},
		// 				Reseller: &armconsumption.Reseller{
		// 					ResellerDescription: to.Ptr("Reseller information"),
		// 					ResellerID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/1234:5678/billingProfiles/2468/providers/Microsoft.Consumption/reseller/reseller1"),
		// 				},
		// 				TransactionDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-09-01T00:00:00.000Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("event4"),
		// 			Type: to.Ptr("Microsoft.Consumption/events"),
		// 			ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/1234:5678/billingProfiles/4268/providers/Microsoft.Consumption/events/event4"),
		// 			Properties: &armconsumption.EventProperties{
		// 				Description: to.Ptr("Settled invoice #212033"),
		// 				Adjustments: &armconsumption.Amount{
		// 					Currency: to.Ptr("USD"),
		// 					Value: to.Ptr[float64](-200),
		// 				},
		// 				AdjustmentsInBillingCurrency: &armconsumption.AmountWithExchangeRate{
		// 					ExchangeRate: to.Ptr[float64](5000),
		// 					ExchangeRateMonth: to.Ptr[int32](1),
		// 				},
		// 				BillingCurrency: to.Ptr("USD"),
		// 				CanceledCredit: &armconsumption.Amount{
		// 					Currency: to.Ptr("USD"),
		// 					Value: to.Ptr[float64](5000),
		// 				},
		// 				Charges: &armconsumption.Amount{
		// 					Currency: to.Ptr("USD"),
		// 					Value: to.Ptr[float64](300),
		// 				},
		// 				ChargesInBillingCurrency: &armconsumption.AmountWithExchangeRate{
		// 					ExchangeRate: to.Ptr[float64](5000),
		// 					ExchangeRateMonth: to.Ptr[int32](1),
		// 				},
		// 				ClosedBalance: &armconsumption.Amount{
		// 					Currency: to.Ptr("USD"),
		// 					Value: to.Ptr[float64](700),
		// 				},
		// 				ClosedBalanceInBillingCurrency: &armconsumption.AmountWithExchangeRate{
		// 					ExchangeRate: to.Ptr[float64](5000),
		// 					ExchangeRateMonth: to.Ptr[int32](1),
		// 				},
		// 				CreditCurrency: to.Ptr("USD"),
		// 				CreditExpired: &armconsumption.Amount{
		// 					Currency: to.Ptr("USD"),
		// 					Value: to.Ptr[float64](100),
		// 				},
		// 				CreditExpiredInBillingCurrency: &armconsumption.AmountWithExchangeRate{
		// 					ExchangeRate: to.Ptr[float64](5000),
		// 					ExchangeRateMonth: to.Ptr[int32](1),
		// 				},
		// 				EventType: to.Ptr(armconsumption.EventTypeSettledCharges),
		// 				InvoiceNumber: to.Ptr("3303"),
		// 				NewCredit: &armconsumption.Amount{
		// 					Currency: to.Ptr("USD"),
		// 					Value: to.Ptr[float64](300),
		// 				},
		// 				NewCreditInBillingCurrency: &armconsumption.AmountWithExchangeRate{
		// 					ExchangeRate: to.Ptr[float64](5000),
		// 					ExchangeRateMonth: to.Ptr[int32](1),
		// 				},
		// 				Reseller: &armconsumption.Reseller{
		// 					ResellerDescription: to.Ptr("Reseller information"),
		// 					ResellerID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/1234:5678/billingProfiles/2468/providers/Microsoft.Consumption/reseller/reseller1"),
		// 				},
		// 				TransactionDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-10-01T00:00:00.000Z"); return t}()),
		// 			},
		// 	}},
		// }
	}
}
