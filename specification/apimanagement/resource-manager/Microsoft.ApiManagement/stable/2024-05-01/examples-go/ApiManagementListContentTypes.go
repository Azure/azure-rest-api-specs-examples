package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e436160e64c0f8d7fb20d662be2712f71f0a7ef5/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/ApiManagementListContentTypes.json
func ExampleContentTypeClient_NewListByServicePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewContentTypeClient().NewListByServicePager("rg1", "apimService1", nil)
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
		// page.ContentTypeCollection = armapimanagement.ContentTypeCollection{
		// 	Value: []*armapimanagement.ContentTypeContract{
		// 		{
		// 			Name: to.Ptr("page"),
		// 			Type: to.Ptr("Microsoft.ApiManagement/service/contentTypes"),
		// 			ID: to.Ptr("/contentTypes/page"),
		// 			Properties: &armapimanagement.ContentTypeContractProperties{
		// 				Name: to.Ptr("Page"),
		// 				Schema: map[string]any{
		// 					"additionalProperties": false,
		// 					"properties":map[string]any{
		// 						"en_us":map[string]any{
		// 							"type": "object",
		// 							"additionalProperties": false,
		// 							"properties":map[string]any{
		// 								"description":map[string]any{
		// 									"type": "string",
		// 									"description": "Page description. This property gets included in SEO attributes.",
		// 									"indexed": true,
		// 									"title": "Description",
		// 								},
		// 								"documentId":map[string]any{
		// 									"type": "string",
		// 									"description": "Reference to page content document.",
		// 									"title": "Document ID",
		// 								},
		// 								"keywords":map[string]any{
		// 									"type": "string",
		// 									"description": "Page keywords. This property gets included in SEO attributes.",
		// 									"indexed": true,
		// 									"title": "Keywords",
		// 								},
		// 								"permalink":map[string]any{
		// 									"type": "string",
		// 									"description": "Page permalink, e.g. '/about'.",
		// 									"indexed": true,
		// 									"title": "Permalink",
		// 								},
		// 								"title":map[string]any{
		// 									"type": "string",
		// 									"description": "Page title. This property gets included in SEO attributes.",
		// 									"indexed": true,
		// 									"title": "Title",
		// 								},
		// 							},
		// 							"required":[]any{
		// 								"title",
		// 								"permalink",
		// 								"documentId",
		// 							},
		// 						},
		// 					},
		// 				},
		// 				Description: to.Ptr("A regular page"),
		// 				Version: to.Ptr("1.0.0"),
		// 			},
		// 	}},
		// }
	}
}
