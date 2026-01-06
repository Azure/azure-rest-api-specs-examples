package armoperationalinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/operationalinsights/armoperationalinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/53d56e4ec74156c450d1e51745a971d3f2031dd7/specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/OperationalInsights/stable/2025-07-01/examples/QueryPackQueriesPut.json
func ExampleQueriesClient_Put() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armoperationalinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewQueriesClient().Put(ctx, "my-resource-group", "my-querypack", "a449f8af-8e64-4b3a-9b16-5a7165ff98c4", armoperationalinsights.LogAnalyticsQueryPackQuery{
		Properties: &armoperationalinsights.LogAnalyticsQueryPackQueryProperties{
			Description: to.Ptr("my description"),
			Body:        to.Ptr("let newExceptionsTimeRange = 1d;\nlet timeRangeToCheckBefore = 7d;\nexceptions\n| where timestamp < ago(timeRangeToCheckBefore)\n| summarize count() by problemId\n| join kind= rightanti (\nexceptions\n| where timestamp >= ago(newExceptionsTimeRange)\n| extend stack = tostring(details[0].rawStack)\n| summarize count(), dcount(user_AuthenticatedId), min(timestamp), max(timestamp), any(stack) by problemId  \n) on problemId \n| order by  count_ desc\n"),
			DisplayName: to.Ptr("Exceptions - New in the last 24 hours"),
			Related: &armoperationalinsights.LogAnalyticsQueryPackQueryPropertiesRelated{
				Categories: []*string{
					to.Ptr("analytics")},
			},
			Tags: map[string][]*string{
				"my-label": {
					to.Ptr("label1")},
				"my-other-label": {
					to.Ptr("label2")},
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.LogAnalyticsQueryPackQuery = armoperationalinsights.LogAnalyticsQueryPackQuery{
	// 	Name: to.Ptr("a449f8af-8e64-4b3a-9b16-5a7165ff98c4"),
	// 	Type: to.Ptr("microsoft.operationalinsights/queryPacks/queries"),
	// 	ID: to.Ptr("/subscriptions/86dc51d3-92ed-4d7e-947a-775ea79b4918/resourceGroups/my-resource-group/providers/microsoft.operationalinsights/queryPacks/my-querypack/queries/a449f8af-8e64-4b3a-9b16-5a7165ff98c4"),
	// 	SystemData: &armoperationalinsights.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-03T01:01:01.107Z"); return t}()),
	// 		CreatedBy: to.Ptr("string"),
	// 		CreatedByType: to.Ptr(armoperationalinsights.CreatedByTypeApplication),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-04T02:03:01.197Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("string"),
	// 		LastModifiedByType: to.Ptr(armoperationalinsights.CreatedByTypeApplication),
	// 	},
	// 	Properties: &armoperationalinsights.LogAnalyticsQueryPackQueryProperties{
	// 		Description: to.Ptr("my description"),
	// 		Author: to.Ptr("1809f206-263a-46f7-942d-4572c156b7e7"),
	// 		Body: to.Ptr("let newExceptionsTimeRange = 1d;\nlet timeRangeToCheckBefore = 7d;\nexceptions\n| where timestamp < ago(timeRangeToCheckBefore)\n| summarize count() by problemId\n| join kind= rightanti (\nexceptions\n| where timestamp >= ago(newExceptionsTimeRange)\n| extend stack = tostring(details[0].rawStack)\n| summarize count(), dcount(user_AuthenticatedId), min(timestamp), max(timestamp), any(stack) by problemId  \n) on problemId \n| order by  count_ desc\n"),
	// 		DisplayName: to.Ptr("Exceptions - New in the last 24 hours"),
	// 		ID: to.Ptr("a449f8af-8e64-4b3a-9b16-5a7165ff98c4"),
	// 		Related: &armoperationalinsights.LogAnalyticsQueryPackQueryPropertiesRelated{
	// 			Categories: []*string{
	// 				to.Ptr("analytics")},
	// 			},
	// 			Tags: map[string][]*string{
	// 				"my-label": []*string{
	// 					to.Ptr("label1")},
	// 					"my-other-label": []*string{
	// 						to.Ptr("label2")},
	// 					},
	// 					TimeCreated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-08-15T10:30:32.574Z"); return t}()),
	// 					TimeModified: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-08-15T10:30:32.574Z"); return t}()),
	// 				},
	// 			}
}
