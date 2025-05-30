package armapplicationinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/applicationinsights/armapplicationinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/8a0168458930c86636a76bcd7acfdc9c81291bfc/specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2020-11-20/examples/WorkbookTemplateAdd.json
func ExampleWorkbookTemplatesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapplicationinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWorkbookTemplatesClient().CreateOrUpdate(ctx, "my-resource-group", "testtemplate2", armapplicationinsights.WorkbookTemplate{
		Location: to.Ptr("west us"),
		Properties: &armapplicationinsights.WorkbookTemplateProperties{
			Author: to.Ptr("Contoso"),
			Galleries: []*armapplicationinsights.WorkbookTemplateGallery{
				{
					Name:         to.Ptr("Simple Template"),
					Type:         to.Ptr("tsg"),
					Category:     to.Ptr("Failures"),
					Order:        to.Ptr[int32](100),
					ResourceType: to.Ptr("microsoft.insights/components"),
				}},
			Priority: to.Ptr[int32](1),
			TemplateData: map[string]any{
				"$schema": "https://github.com/Microsoft/Application-Insights-Workbooks/blob/master/schema/workbook.json",
				"items": []any{
					map[string]any{
						"name": "text - 2",
						"type": float64(1),
						"content": map[string]any{
							"json": "## New workbook\n---\n\nWelcome to your new workbook.  This area will display text formatted as markdown.\n\n\nWe've included a basic analytics query to get you started. Use the `Edit` button below each section to configure it or add more sections.",
						},
					},
					map[string]any{
						"name": "query - 2",
						"type": float64(3),
						"content": map[string]any{
							"exportToExcelOptions": "visible",
							"query":                "union withsource=TableName *\n| summarize Count=count() by TableName\n| render barchart",
							"queryType":            float64(0),
							"resourceType":         "microsoft.operationalinsights/workspaces",
							"size":                 float64(1),
							"version":              "KqlItem/1.0",
						},
					},
				},
				"styleSettings": map[string]any{},
				"version":       "Notebook/1.0",
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.WorkbookTemplate = armapplicationinsights.WorkbookTemplate{
	// 	Name: to.Ptr("testtemplate2"),
	// 	Type: to.Ptr("microsoft.insights/workbooktemplates"),
	// 	ID: to.Ptr("/subscriptions/50359d91-7b9d-4823-85af-eb298a61ba95/resourceGroups/testrg/providers/microsoft.insights/workbooktemplates/testtemplate2"),
	// 	Location: to.Ptr("westeurope"),
	// 	Properties: &armapplicationinsights.WorkbookTemplateProperties{
	// 		Author: to.Ptr("Contoso"),
	// 		Galleries: []*armapplicationinsights.WorkbookTemplateGallery{
	// 			{
	// 				Name: to.Ptr("Simple Template"),
	// 				Type: to.Ptr("tsg"),
	// 				Category: to.Ptr("Failures"),
	// 				Order: to.Ptr[int32](100),
	// 				ResourceType: to.Ptr("microsoft.insights/components"),
	// 		}},
	// 		Priority: to.Ptr[int32](1),
	// 		TemplateData: map[string]any{
	// 			"$schema": "https://github.com/Microsoft/Application-Insights-Workbooks/blob/master/schema/workbook.json",
	// 			"items":[]any{
	// 				map[string]any{
	// 					"name": "text - 2",
	// 					"type": float64(1),
	// 					"content":map[string]any{
	// 						"json": "## New workbook\n---\n\nWelcome to your new workbook.  This area will display text formatted as markdown.\n\n\nWe've included a basic analytics query to get you started. Use the `Edit` button below each section to configure it or add more sections.",
	// 					},
	// 				},
	// 				map[string]any{
	// 					"name": "query - 2",
	// 					"type": float64(3),
	// 					"content":map[string]any{
	// 						"exportToExcelOptions": "visible",
	// 						"query": "union withsource=TableName *\n| summarize Count=count() by TableName\n| render barchart",
	// 						"queryType": float64(0),
	// 						"resourceType": "microsoft.operationalinsights/workspaces",
	// 						"size": float64(1),
	// 						"version": "KqlItem/1.0",
	// 					},
	// 				},
	// 			},
	// 			"styleSettings":map[string]any{
	// 			},
	// 			"version": "Notebook/1.0",
	// 		},
	// 	},
	// }
}
