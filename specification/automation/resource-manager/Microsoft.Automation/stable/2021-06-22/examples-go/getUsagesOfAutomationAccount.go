package armautomation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/432872fac1d0f8edcae98a0e8504afc0ee302710/specification/automation/resource-manager/Microsoft.Automation/stable/2021-06-22/examples/getUsagesOfAutomationAccount.json
func ExampleUsagesClient_NewListByAutomationAccountPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armautomation.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewUsagesClient().NewListByAutomationAccountPager("rg", "myAutomationAccount11", nil)
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
		// page.UsageListResult = armautomation.UsageListResult{
		// 	Value: []*armautomation.Usage{
		// 		{
		// 			Name: &armautomation.UsageCounterName{
		// 				LocalizedValue: to.Ptr("AccountUsage"),
		// 				Value: to.Ptr("AccountUsage"),
		// 			},
		// 			CurrentValue: to.Ptr[float64](0),
		// 			Limit: to.Ptr[int64](500),
		// 			ThrottleStatus: to.Ptr("NotThrottled"),
		// 			Unit: to.Ptr("Minute"),
		// 		},
		// 		{
		// 			Name: &armautomation.UsageCounterName{
		// 				LocalizedValue: to.Ptr("SubscriptionUsage"),
		// 				Value: to.Ptr("SubscriptionUsage"),
		// 			},
		// 			CurrentValue: to.Ptr[float64](429),
		// 			Limit: to.Ptr[int64](500),
		// 			ThrottleStatus: to.Ptr("NotThrottled"),
		// 			Unit: to.Ptr("Minute"),
		// 		},
		// 		{
		// 			Name: &armautomation.UsageCounterName{
		// 				LocalizedValue: to.Ptr("DscSubscriptionUsage"),
		// 				Value: to.Ptr("DscSubscriptionUsage"),
		// 			},
		// 			CurrentValue: to.Ptr[float64](8),
		// 			Limit: to.Ptr[int64](5),
		// 			ThrottleStatus: to.Ptr("ThrottledAtSubscriptionLevel"),
		// 			Unit: to.Ptr("Count"),
		// 	}},
		// }
	}
}
