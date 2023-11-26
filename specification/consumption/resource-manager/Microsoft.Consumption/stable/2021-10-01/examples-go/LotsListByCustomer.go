package armconsumption_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/consumption/armconsumption"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/LotsListByCustomer.json
func ExampleLotsClient_NewListByCustomerPager_lotsListByCustomer() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armconsumption.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewLotsClient().NewListByCustomerPager("1234:5678", "1234:5678", &armconsumption.LotsClientListByCustomerOptions{Filter: nil})
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
		// 			ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/1234/customers/456/providers/Microsoft.Consumption/lots/lot1"),
		// 			Properties: &armconsumption.LotProperties{
		// 				ClosedBalance: &armconsumption.Amount{
		// 					Currency: to.Ptr("USD"),
		// 					Value: to.Ptr[float64](60.9),
		// 				},
		// 				ExpirationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-05-01T00:00:00.000Z"); return t}()),
		// 				OriginalAmount: &armconsumption.Amount{
		// 					Currency: to.Ptr("USD"),
		// 					Value: to.Ptr[float64](5000),
		// 				},
		// 				PoNumber: to.Ptr("3524"),
		// 				Source: to.Ptr(armconsumption.LotSourcePurchasedCredit),
		// 				StartDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-05-01T00:00:00.000Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("lot2"),
		// 			Type: to.Ptr("Microsoft.Consumption/lots"),
		// 			ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/1234/customers/456/providers/Microsoft.Consumption/lots/lot2"),
		// 			Properties: &armconsumption.LotProperties{
		// 				ClosedBalance: &armconsumption.Amount{
		// 					Currency: to.Ptr("USD"),
		// 					Value: to.Ptr[float64](80.9),
		// 				},
		// 				ExpirationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-12-31T00:00:00.000Z"); return t}()),
		// 				OriginalAmount: &armconsumption.Amount{
		// 					Currency: to.Ptr("USD"),
		// 					Value: to.Ptr[float64](6000),
		// 				},
		// 				PoNumber: to.Ptr("31224"),
		// 				Source: to.Ptr(armconsumption.LotSourcePurchasedCredit),
		// 				StartDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-05-01T00:00:00.000Z"); return t}()),
		// 			},
		// 	}},
		// }
	}
}
