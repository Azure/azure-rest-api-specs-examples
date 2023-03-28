package armapplicationinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/applicationinsights/armapplicationinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7932c2df6c8435d6c0e5cbebbca79bce627d5f06/specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2022-04-01/examples/WorkbookGet1.json
func ExampleWorkbooksClient_Get_workbookGet1() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapplicationinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWorkbooksClient().Get(ctx, "my-resource-group", "deadb33f-5e0d-4064-8ebb-1a4ed0313eb2", &armapplicationinsights.WorkbooksClientGetOptions{CanFetchContent: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Workbook = armapplicationinsights.Workbook{
	// 	Name: to.Ptr("deadb33f-5e0d-4064-8ebb-1a4ed0313eb2"),
	// 	Type: to.Ptr("Microsoft.Insights/workbooks"),
	// 	ID: to.Ptr("/subscriptions/6b643656-33eb-422f-aee8-3ac145d124af/resourcegroups/my-resource-group/providers/microsoft.insights/workbooks/deadb33f-5e0d-4064-8ebb-1a4ed0313eb2"),
	// 	Location: to.Ptr("westus"),
	// 	Tags: map[string]*string{
	// 		"TagSample01": to.Ptr("sample01"),
	// 		"TagSample02": to.Ptr("sample02"),
	// 	},
	// 	Kind: to.Ptr(armapplicationinsights.WorkbookSharedTypeKindShared),
	// 	Properties: &armapplicationinsights.WorkbookProperties{
	// 		Category: to.Ptr("workbook"),
	// 		DisplayName: to.Ptr("Sample workbook"),
	// 		SerializedData: to.Ptr("{\"version\":\"Notebook/1.0\",\"items\":[{\"type\":1,\"content\":\"{\"json\":\"## New workbook\\r\\n---\\r\\n\\r\\nWelcome to your new workbook.  This area will display text formatted as markdown.\\r\\n\\r\\n\\r\\nWe've included a basic analytics query to get you started. Use the `Edit` button below each section to configure it or add more sections.\"}\",\"halfWidth\":null,\"conditionalVisibility\":null},{\"type\":3,\"content\":\"{\"version\":\"KqlItem/1.0\",\"query\":\"union withsource=TableName *\\n| summarize Count=count() by TableName\\n| render barchart\",\"showQuery\":false,\"size\":1,\"aggregation\":0,\"showAnnotations\":false}\",\"halfWidth\":null,\"conditionalVisibility\":null}],\"isLocked\":false}"),
	// 		TimeModified: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-24T06:56:31.6301521Z"); return t}()),
	// 		UserID: to.Ptr("userId"),
	// 		Version: to.Ptr("Notebook/1.0"),
	// 	},
	// }
}
