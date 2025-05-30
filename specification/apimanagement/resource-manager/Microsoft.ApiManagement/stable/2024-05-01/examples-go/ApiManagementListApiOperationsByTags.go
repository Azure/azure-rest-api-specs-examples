package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e436160e64c0f8d7fb20d662be2712f71f0a7ef5/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/ApiManagementListApiOperationsByTags.json
func ExampleOperationClient_NewListByTagsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewOperationClient().NewListByTagsPager("rg1", "apimService1", "a1", &armapimanagement.OperationClientListByTagsOptions{Filter: nil,
		Top:                        nil,
		Skip:                       nil,
		IncludeNotTaggedOperations: nil,
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
		// page.TagResourceCollection = armapimanagement.TagResourceCollection{
		// 	Count: to.Ptr[int64](1),
		// 	Value: []*armapimanagement.TagResourceContract{
		// 		{
		// 			Operation: &armapimanagement.OperationTagResourceContractProperties{
		// 				Name: to.Ptr("Create resource"),
		// 				Method: to.Ptr("POST"),
		// 				Description: to.Ptr("A demonstration of a POST call based on the echo backend above. The request body is expected to contain JSON-formatted data (see example below). A policy is used to automatically transform any request sent in JSON directly to XML. In a real-world scenario this could be used to enable modern clients to speak to a legacy backend."),
		// 				APIName: to.Ptr("Echo API"),
		// 				APIRevision: to.Ptr("1"),
		// 				ID: to.Ptr("/apis/echo-api/operations/create-resource"),
		// 				URLTemplate: to.Ptr("/resource"),
		// 			},
		// 			Tag: &armapimanagement.TagResourceContractProperties{
		// 				Name: to.Ptr("awesomeTag"),
		// 				ID: to.Ptr("/tags/apitag123"),
		// 			},
		// 	}},
		// }
	}
}
