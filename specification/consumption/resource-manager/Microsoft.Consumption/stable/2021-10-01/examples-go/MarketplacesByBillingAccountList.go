package armconsumption_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/consumption/armconsumption"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/MarketplacesByBillingAccountList.json
func ExampleMarketplacesClient_NewListPager_billingAccountMarketplacesList() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armconsumption.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewMarketplacesClient().NewListPager("providers/Microsoft.Billing/billingAccounts/123456", &armconsumption.MarketplacesClientListOptions{Filter: nil,
		Top:       nil,
		Skiptoken: nil,
	})
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
		// page.MarketplacesListResult = armconsumption.MarketplacesListResult{
		// 	Value: []*armconsumption.Marketplace{
		// 		{
		// 			Name: to.Ptr("marketplaceId1"),
		// 			Type: to.Ptr("Microsoft.Consumption/marketPlaces"),
		// 			ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/123456/providers/Microsoft.Billing/billingPeriods/201702/providers/Microsoft.Consumption/marketplaces/marketplaceId1"),
		// 			Tags: map[string]*string{
		// 				"dev": to.Ptr("tools"),
		// 				"env": to.Ptr("newcrp"),
		// 			},
		// 			Properties: &armconsumption.MarketplaceProperties{
		// 				AccountName: to.Ptr("Account1"),
		// 				AdditionalProperties: to.Ptr("additionalProperties"),
		// 				BillingPeriodID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/123456/providers/Microsoft.Billing/billingPeriods/201702"),
		// 				ConsumedQuantity: to.Ptr[float64](0.00328),
		// 				CostCenter: to.Ptr("Center1"),
		// 				Currency: to.Ptr("USD"),
		// 				DepartmentName: to.Ptr("Department1"),
		// 				InstanceID: to.Ptr("/subscriptions/subid/resourceGroups/Default-Web-eastasia/providers/Microsoft.Web/sites/shared1"),
		// 				InstanceName: to.Ptr("shared1"),
		// 				IsEstimated: to.Ptr(false),
		// 				IsRecurringCharge: to.Ptr(false),
		// 				OfferName: to.Ptr("offer1"),
		// 				OrderNumber: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 				PlanName: to.Ptr("plan1"),
		// 				PretaxCost: to.Ptr[float64](0.67),
		// 				PublisherName: to.Ptr("xyz"),
		// 				ResourceGroup: to.Ptr("TEST"),
		// 				ResourceRate: to.Ptr[float64](0.24),
		// 				SubscriptionGUID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 				SubscriptionName: to.Ptr("azure subscription"),
		// 				UnitOfMeasure: to.Ptr("10 Hours"),
		// 				UsageEnd: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-02-13T23:59:59.000Z"); return t}()),
		// 				UsageStart: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-02-13T00:00:00.000Z"); return t}()),
		// 			},
		// 	}},
		// }
	}
}
