package armapplicationinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/applicationinsights/armapplicationinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2015-05-01/examples/WorkbookUpdate.json
func ExampleWorkbooksClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armapplicationinsights.NewWorkbooksClient("6b643656-33eb-422f-aee8-3ac145d124af", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Update(ctx,
		"my-resource-group",
		"deadb33f-5e0d-4064-8ebb-1a4ed0313eb2",
		armapplicationinsights.Workbook{
			Name:     to.Ptr("deadb33f-8bee-4d3b-a059-9be8dac93960"),
			Location: to.Ptr("west us"),
			Tags: map[string]*string{
				"0": to.Ptr("TagSample01"),
				"1": to.Ptr("TagSample02"),
			},
			Properties: &armapplicationinsights.WorkbookProperties{
				Name:           to.Ptr("Blah Blah Blah"),
				Category:       to.Ptr("workbook"),
				SharedTypeKind: to.Ptr(armapplicationinsights.SharedTypeKindShared),
				SerializedData: to.Ptr("{\"version\":\"Notebook/1.0\",\"items\":[{\"type\":1,\"content\":\"{\"json\":\"## New workbook\\r\\n---\\r\\n\\r\\nWelcome to your new workbook.  This area will display text formatted as markdown.\\r\\n\\r\\n\\r\\nWe've included a basic analytics query to get you started. Use the `Edit` button below each section to configure it or add more sections.\"}\",\"halfWidth\":null,\"conditionalVisibility\":null},{\"type\":3,\"content\":\"{\"version\":\"KqlItem/1.0\",\"query\":\"union withsource=TableName *\\n| summarize Count=count() by TableName\\n| render barchart\",\"showQuery\":false,\"size\":1,\"aggregation\":0,\"showAnnotations\":false}\",\"halfWidth\":null,\"conditionalVisibility\":null}],\"isLocked\":false}"),
				UserID:         to.Ptr("userId"),
				Version:        to.Ptr("ME"),
				WorkbookID:     to.Ptr("deadb33f-8bee-4d3b-a059-9be8dac93960"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
