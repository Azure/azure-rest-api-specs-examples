package armconsumption_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/consumption/armconsumption/v2"
)

// Generated from example definition: 2024-08-01/LotsListByCustomerWithFilters.json
func ExampleLotsClient_NewListByCustomerPager_lotsListByCustomerWithFilter() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armconsumption.NewClientFactory("<subscriptionID>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewLotsClient().NewListByCustomerPager("1234:5678", "1234:5678", &armconsumption.LotsClientListByCustomerOptions{
		Filter: to.Ptr("status eq 'active' AND source eq 'consumptioncommitment'")})
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
		// page = armconsumption.LotsClientListByCustomerResponse{
		// 	Lots: armconsumption.Lots{
		// 		Value: []*armconsumption.LotSummary{
		// 			{
		// 				Name: to.Ptr("lot1"),
		// 				Type: to.Ptr("Microsoft.Consumption/lots"),
		// 				ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/1234/customers/456/providers/Microsoft.Consumption/lots/lot1"),
		// 				Properties: &armconsumption.LotProperties{
		// 					ClosedBalance: &armconsumption.Amount{
		// 						Currency: to.Ptr("USD"),
		// 					},
		// 					ExpirationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-05-01T00:00:00Z"); return t}()),
		// 					IsEstimatedBalance: to.Ptr(false),
		// 					OriginalAmount: &armconsumption.Amount{
		// 						Currency: to.Ptr("USD"),
		// 					},
		// 					PoNumber: to.Ptr("3524"),
		// 					Source: to.Ptr(armconsumption.LotSourcePurchasedCredit),
		// 					StartDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-05-01T00:00:00Z"); return t}()),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
