package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4cd95123fb961c68740565a1efcaa5e43bd35802/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementGetContentTypeContentItem.json
func ExampleContentItemClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewContentItemClient().Get(ctx, "rg1", "apimService1", "page", "4e3cf6a5-574a-ba08-1f23-2e7a38faa6d8", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ContentItemContract = armapimanagement.ContentItemContract{
	// 	Name: to.Ptr("4e3cf6a5-574a-ba08-1f23-2e7a38faa6d8"),
	// 	Type: to.Ptr("Microsoft.ApiManagement/service/contentTypes/contentItems"),
	// 	ID: to.Ptr("/contentTypes/page/contentItems/4e3cf6a5-574a-ba08-1f23-2e7a38faa6d8"),
	// 	Properties: map[string]any{
	// 		"en_us": map[string]any{
	// 			"description": "Short story about the company.",
	// 			"documentId": "contentTypes/document/contentItems/4e3cf6a5-574a-ba08-1f23-2e7a38faa6d8",
	// 			"keywords": "company, about",
	// 			"permalink": "/about",
	// 			"title": "About",
	// 		},
	// 	},
	// }
}
