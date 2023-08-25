package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4cd95123fb961c68740565a1efcaa5e43bd35802/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementGetContentType.json
func ExampleContentTypeClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewContentTypeClient().Get(ctx, "rg1", "apimService1", "page", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ContentTypeContract = armapimanagement.ContentTypeContract{
	// 	Name: to.Ptr("page"),
	// 	Type: to.Ptr("Microsoft.ApiManagement/service/contentTypes"),
	// 	ID: to.Ptr("/subscriptions/subid/resourcegroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/contentTypes/page"),
	// 	Properties: &armapimanagement.ContentTypeContractProperties{
	// 		Name: to.Ptr("Page"),
	// 		Schema: map[string]any{
	// 			"additionalProperties": false,
	// 			"properties":map[string]any{
	// 				"en_us":map[string]any{
	// 					"type": "object",
	// 					"additionalProperties": false,
	// 					"properties":map[string]any{
	// 						"description":map[string]any{
	// 							"type": "string",
	// 							"description": "Page description. This property gets included in SEO attributes.",
	// 							"indexed": true,
	// 							"title": "Description",
	// 						},
	// 						"documentId":map[string]any{
	// 							"type": "string",
	// 							"description": "Reference to page content document.",
	// 							"title": "Document ID",
	// 						},
	// 						"keywords":map[string]any{
	// 							"type": "string",
	// 							"description": "Page keywords. This property gets included in SEO attributes.",
	// 							"indexed": true,
	// 							"title": "Keywords",
	// 						},
	// 						"permalink":map[string]any{
	// 							"type": "string",
	// 							"description": "Page permalink, e.g. '/about'.",
	// 							"indexed": true,
	// 							"title": "Permalink",
	// 						},
	// 						"title":map[string]any{
	// 							"type": "string",
	// 							"description": "Page title. This property gets included in SEO attributes.",
	// 							"indexed": true,
	// 							"title": "Title",
	// 						},
	// 					},
	// 					"required":[]any{
	// 						"title",
	// 						"permalink",
	// 						"documentId",
	// 					},
	// 				},
	// 			},
	// 		},
	// 		Description: to.Ptr("A regular page"),
	// 		Version: to.Ptr("1.0.0"),
	// 	},
	// }
}
