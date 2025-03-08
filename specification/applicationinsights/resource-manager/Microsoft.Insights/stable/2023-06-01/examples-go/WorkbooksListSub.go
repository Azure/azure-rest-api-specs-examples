package armapplicationinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/applicationinsights/armapplicationinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/8a0168458930c86636a76bcd7acfdc9c81291bfc/specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2023-06-01/examples/WorkbooksListSub.json
func ExampleWorkbooksClient_NewListBySubscriptionPager_workbooksListSub() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapplicationinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewWorkbooksClient().NewListBySubscriptionPager(armapplicationinsights.CategoryTypeWorkbook, &armapplicationinsights.WorkbooksClientListBySubscriptionOptions{Tags: []string{},
		CanFetchContent: nil,
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
		// page.WorkbooksListResult = armapplicationinsights.WorkbooksListResult{
		// 	Value: []*armapplicationinsights.Workbook{
		// 		{
		// 			Name: to.Ptr("deadb33f-5e0d-4064-8ebb-1a4ed0313eb2"),
		// 			Type: to.Ptr("Microsoft.Insights/workbooks"),
		// 			ID: to.Ptr("/subscriptions/6b643656-33eb-422f-aee8-3ac145d124af/resourcegroups/my-resource-group/providers/Microsoft.Insights/workbooks/deadb33f-5e0d-4064-8ebb-1a4ed0313eb2"),
		// 			Location: to.Ptr("westus"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Kind: to.Ptr(armapplicationinsights.WorkbookSharedTypeKindShared),
		// 			Properties: &armapplicationinsights.WorkbookProperties{
		// 				Description: to.Ptr("Sample workbook"),
		// 				Category: to.Ptr("workbook"),
		// 				DisplayName: to.Ptr("My Workbook 1"),
		// 				Revision: to.Ptr("1e2f8435b98248febee70c64ac22e1bb"),
		// 				SerializedData: to.Ptr("{\"version\":\"Notebook/1.0\",\"items\":[{\"type\":1,\"content\":\"{\"json\":\"## New workbook\\r\\n---\\r\\n\\r\\nWelcome to your new workbook.  This area will display text formatted as markdown.\\r\\n\\r\\n\\r\\nWe've included a basic analytics query to get you started. Use the `Edit` button below each section to configure it or add more sections.\"}\",\"halfWidth\":null,\"conditionalVisibility\":null},{\"type\":3,\"content\":\"{\"version\":\"KqlItem/1.0\",\"query\":\"union withsource=TableName *\\n| summarize Count=count() by TableName\\n| render barchart\",\"showQuery\":false,\"size\":1,\"aggregation\":0,\"showAnnotations\":false}\",\"halfWidth\":null,\"conditionalVisibility\":null}],\"isLocked\":false}"),
		// 				SourceID: to.Ptr("/subscriptions/6b643656-33eb-422f-aee8-3ac145d124af/resourcegroups/my-resource-group/providers/Microsoft.Web/sites/MyApp"),
		// 				TimeModified: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-10-20T22:00:26.422Z"); return t}()),
		// 				UserID: to.Ptr("userId"),
		// 				Version: to.Ptr("Notebook/1.0"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("c0deea5e-3344-40f2-96f8-6f8e1c3b5722"),
		// 			Type: to.Ptr("Microsoft.Insights/workbooks"),
		// 			ID: to.Ptr("/subscriptions/6b643656-33eb-422f-aee8-3ac145d124af/resourcegroups/my-resource-group/providers/Microsoft.Insights/workbooks/c0deea5e-3344-40f2-96f8-6f8e1c3b5722"),
		// 			Location: to.Ptr("westus"),
		// 			Tags: map[string]*string{
		// 				"TagSample01": to.Ptr("sample01"),
		// 				"TagSample02": to.Ptr("sample02"),
		// 			},
		// 			Kind: to.Ptr(armapplicationinsights.WorkbookSharedTypeKindShared),
		// 			Properties: &armapplicationinsights.WorkbookProperties{
		// 				Description: to.Ptr("Sample workbook"),
		// 				Category: to.Ptr("workbook"),
		// 				DisplayName: to.Ptr("My Workbook 2"),
		// 				Revision: to.Ptr("1e2f8435b98248febee70c64ac22e1bc"),
		// 				SerializedData: to.Ptr("{\"version\":\"Notebook/1.0\",\"items\":[{\"type\":1,\"content\":\"{\"json\":\"## New workbook\\r\\n---\\r\\n\\r\\nWelcome to your new workbook.  This area will display text formatted as markdown.\\r\\n\\r\\n\\r\\nWe've included a basic analytics query to get you started. Use the `Edit` button below each section to configure it or add more sections.\"}\",\"halfWidth\":null,\"conditionalVisibility\":null},{\"type\":3,\"content\":\"{\"version\":\"KqlItem/1.0\",\"query\":\"union withsource=TableName *\\n| summarize Count=count() by TableName\\n| render barchart\",\"showQuery\":false,\"size\":1,\"aggregation\":0,\"showAnnotations\":false}\",\"halfWidth\":null,\"conditionalVisibility\":null}],\"isLocked\":false}"),
		// 				SourceID: to.Ptr("/subscriptions/6b643656-33eb-422f-aee8-3ac145d124af/resourcegroups/my-resource-group/providers/Microsoft.Web/sites/MyApp"),
		// 				TimeModified: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-10-20T22:00:26.422Z"); return t}()),
		// 				UserID: to.Ptr("userId"),
		// 				Version: to.Ptr("Notebook/1.0"),
		// 			},
		// 	}},
		// }
	}
}
