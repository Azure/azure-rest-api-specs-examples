Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fapplicationinsights%2Farmapplicationinsights%2Fv0.4.0/sdk/resourcemanager/applicationinsights/armapplicationinsights/README.md) on how to add the SDK to your project and authenticate.

```go
package armapplicationinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/applicationinsights/armapplicationinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2020-11-20/examples/WorkbookTemplateAdd.json
func ExampleWorkbookTemplatesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armapplicationinsights.NewWorkbookTemplatesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<resource-name>",
		armapplicationinsights.WorkbookTemplate{
			Location: to.Ptr("<location>"),
			Properties: &armapplicationinsights.WorkbookTemplateProperties{
				Author: to.Ptr("<author>"),
				Galleries: []*armapplicationinsights.WorkbookTemplateGallery{
					{
						Name:         to.Ptr("<name>"),
						Type:         to.Ptr("<type>"),
						Category:     to.Ptr("<category>"),
						Order:        to.Ptr[int32](100),
						ResourceType: to.Ptr("<resource-type>"),
					}},
				Priority: to.Ptr[int32](1),
				TemplateData: map[string]interface{}{
					"$schema": "https://github.com/Microsoft/Application-Insights-Workbooks/blob/master/schema/workbook.json",
					"items": []interface{}{
						map[string]interface{}{
							"name": "text - 2",
							"type": float64(1),
							"content": map[string]interface{}{
								"json": "## New workbook\n---\n\nWelcome to your new workbook.  This area will display text formatted as markdown.\n\n\nWe've included a basic analytics query to get you started. Use the `Edit` button below each section to configure it or add more sections.",
							},
						},
						map[string]interface{}{
							"name": "query - 2",
							"type": float64(3),
							"content": map[string]interface{}{
								"exportToExcelOptions": "visible",
								"query":                "union withsource=TableName *\n| summarize Count=count() by TableName\n| render barchart",
								"queryType":            float64(0),
								"resourceType":         "microsoft.operationalinsights/workspaces",
								"size":                 float64(1),
								"version":              "KqlItem/1.0",
							},
						},
					},
					"styleSettings": map[string]interface{}{},
					"version":       "Notebook/1.0",
				},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
