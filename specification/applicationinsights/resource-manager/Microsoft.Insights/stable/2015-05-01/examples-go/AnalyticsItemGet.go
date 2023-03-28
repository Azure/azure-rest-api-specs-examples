package armapplicationinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/applicationinsights/armapplicationinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7932c2df6c8435d6c0e5cbebbca79bce627d5f06/specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2015-05-01/examples/AnalyticsItemGet.json
func ExampleAnalyticsItemsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapplicationinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAnalyticsItemsClient().Get(ctx, "my-resource-group", "my-component", armapplicationinsights.ItemScopePathAnalyticsItems, &armapplicationinsights.AnalyticsItemsClientGetOptions{ID: to.Ptr("3466c160-4a10-4df8-afdf-0007f3f6dee5"),
		Name: nil,
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ComponentAnalyticsItem = armapplicationinsights.ComponentAnalyticsItem{
	// 	Content: to.Ptr("let newExceptionsTimeRange = 1d;\nlet timeRangeToCheckBefore = 7d;\nexceptions\n| where timestamp < ago(timeRangeToCheckBefore)\n| summarize count() by problemId\n| join kind= rightanti (\nexceptions\n| where timestamp >= ago(newExceptionsTimeRange)\n| extend stack = tostring(details[0].rawStack)\n| summarize count(), dcount(user_AuthenticatedId), min(timestamp), max(timestamp), any(stack) by problemId  \n) on problemId \n| order by  count_ desc\n"),
	// 	ID: to.Ptr("3466c160-4a10-4df8-afdf-0007f3f6dee5"),
	// 	Name: to.Ptr("Exceptions - New in the last 24 hours"),
	// 	Scope: to.Ptr(armapplicationinsights.ItemScopeShared),
	// 	TimeCreated: to.Ptr("2018-02-12T11:44:39.2980634Z"),
	// 	TimeModified: to.Ptr("2018-02-14T13:13:19.3381394Z"),
	// 	Type: to.Ptr(armapplicationinsights.ItemTypeQuery),
	// 	Version: to.Ptr("1.0"),
	// }
}
