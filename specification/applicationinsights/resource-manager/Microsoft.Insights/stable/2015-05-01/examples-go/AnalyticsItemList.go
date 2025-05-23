package armapplicationinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/applicationinsights/armapplicationinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/8a0168458930c86636a76bcd7acfdc9c81291bfc/specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2015-05-01/examples/AnalyticsItemList.json
func ExampleAnalyticsItemsClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapplicationinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAnalyticsItemsClient().List(ctx, "my-resource-group", "my-component", armapplicationinsights.ItemScopePathAnalyticsItems, &armapplicationinsights.AnalyticsItemsClientListOptions{Scope: nil,
		Type:           nil,
		IncludeContent: nil,
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ComponentAnalyticsItemArray = []*armapplicationinsights.ComponentAnalyticsItem{
	// 	{
	// 		Content: to.Ptr("//Top 10 countries by traffic in the past 24 hours\nrequests \n | where  timestamp > ago(24h) \n | summarize count() by client_CountryOrRegion\n | top 10 by count_ \n | render piechart"),
	// 		ID: to.Ptr("b753348d-333a-4678-a684-c0e9090713b7"),
	// 		Name: to.Ptr("1"),
	// 		Scope: to.Ptr(armapplicationinsights.ItemScopeUser),
	// 		TimeModified: to.Ptr("2017-06-29T10:27:03Z"),
	// 		Type: to.Ptr(armapplicationinsights.ItemTypeQuery),
	// 		Version: to.Ptr("1.0"),
	// 	},
	// 	{
	// 		Content: to.Ptr("//Top 10 countries by traffic in the past 24 hours\nrequests \n | where  timestamp > ago(24h) \n | summarize count() by client_CountryOrRegion\n | top 10 by count_ \n | render piechart"),
	// 		ID: to.Ptr("0d2f1b19-04b2-4c93-bc6f-2466b23c5284"),
	// 		Name: to.Ptr("4"),
	// 		Scope: to.Ptr(armapplicationinsights.ItemScopeUser),
	// 		TimeModified: to.Ptr("2017-06-29T10:27:13Z"),
	// 		Type: to.Ptr(armapplicationinsights.ItemTypeQuery),
	// 		Version: to.Ptr("1.0"),
	// 	},
	// 	{
	// 		Content: to.Ptr("//Top 10 countries by traffic in the past 24 hours\nrequests \n | where  timestamp > ago(24h) \n | summarize count() by client_CountryOrRegion\n | top 10 by count_ \n | render piechart"),
	// 		ID: to.Ptr("3d17bebb-0b20-4b58-9bbd-22aeed70be51"),
	// 		Name: to.Ptr("2"),
	// 		Scope: to.Ptr(armapplicationinsights.ItemScopeUser),
	// 		TimeModified: to.Ptr("2018-02-10T23:21:05.9952874Z"),
	// 		Type: to.Ptr(armapplicationinsights.ItemTypeQuery),
	// 		Version: to.Ptr("1.0"),
	// 	},
	// 	{
	// 		Content: to.Ptr("//Top 10 countries by traffic in the past 24 hours\nrequests \n | where  timestamp > ago(24h) \n | summarize count() by client_CountryOrRegion\n | top 10 by count_ \n | render piechart"),
	// 		ID: to.Ptr("2be491c6-10d9-4cf6-9490-2a7ce7270c54"),
	// 		Name: to.Ptr("5"),
	// 		Scope: to.Ptr(armapplicationinsights.ItemScopeUser),
	// 		TimeModified: to.Ptr("2017-06-29T10:27:17Z"),
	// 		Type: to.Ptr(armapplicationinsights.ItemTypeQuery),
	// 		Version: to.Ptr("1.0"),
	// 	},
	// 	{
	// 		Content: to.Ptr("//Top 10 countries by traffic in the past 24 hours\nrequests \n | where  timestamp > ago(24h) \n | summarize count() by client_CountryOrRegion\n | top 10 by count_ \n | render piechart"),
	// 		ID: to.Ptr("d8f83601-4a40-4dc1-8516-0a28dcb74420"),
	// 		Name: to.Ptr("8"),
	// 		Scope: to.Ptr(armapplicationinsights.ItemScopeUser),
	// 		TimeCreated: to.Ptr("2018-02-10T23:20:19.0174631Z"),
	// 		TimeModified: to.Ptr("2018-02-10T23:20:19.0174631Z"),
	// 		Type: to.Ptr(armapplicationinsights.ItemTypeQuery),
	// 		Version: to.Ptr("1.0"),
	// 	},
	// 	{
	// 		Content: to.Ptr("let newExceptionsTimeRange = 7d;\nlet timeRangeToCheckBefore = 7d;\nexceptions\n| where timestamp < ago(timeRangeToCheckBefore)\n| summarize count() by problemId\n| join kind= rightanti (\nexceptions\n| where timestamp >= ago(newExceptionsTimeRange)\n| extend stack = tostring(details[0].rawStack)\n| summarize count(), dcount(user_AuthenticatedId), min(timestamp), max(timestamp), any(stack) by problemId  \n) on problemId \n| order by  count_ desc\n"),
	// 		ID: to.Ptr("fd3afe4d-9139-4c76-9b47-81d0fada977b"),
	// 		Name: to.Ptr("Exceptions - New in the last 7 days"),
	// 		Scope: to.Ptr(armapplicationinsights.ItemScopeUser),
	// 		TimeCreated: to.Ptr("2018-02-11T22:05:57.6019354Z"),
	// 		TimeModified: to.Ptr("2018-02-12T11:01:15.5687326Z"),
	// 		Type: to.Ptr(armapplicationinsights.ItemTypeQuery),
	// 		Version: to.Ptr("1.0"),
	// 	},
	// 	{
	// 		Content: to.Ptr("let newExceptionsTimeRange = 1d;\nlet timeRangeToCheckBefore = 7d;\nexceptions\n| where timestamp < ago(timeRangeToCheckBefore)\n| summarize count() by problemId\n| join kind= rightanti (\nexceptions\n| where timestamp >= ago(newExceptionsTimeRange)\n| extend stack = tostring(details[0].rawStack)\n| summarize count(), dcount(user_AuthenticatedId), min(timestamp), max(timestamp), any(stack) by problemId  \n) on problemId \n| order by  count_ desc\n"),
	// 		ID: to.Ptr("3466c160-4a10-4df8-afdf-0007f3f6dee5"),
	// 		Name: to.Ptr("Exceptions - New in the last 24 hours"),
	// 		Scope: to.Ptr(armapplicationinsights.ItemScopeShared),
	// 		TimeCreated: to.Ptr("2018-02-12T11:44:39.2980634Z"),
	// 		TimeModified: to.Ptr("2018-02-14T13:13:19.3381394Z"),
	// 		Type: to.Ptr(armapplicationinsights.ItemTypeQuery),
	// 		Version: to.Ptr("1.0"),
	// }}
}
