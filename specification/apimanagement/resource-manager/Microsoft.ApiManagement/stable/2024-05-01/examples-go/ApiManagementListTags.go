package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e436160e64c0f8d7fb20d662be2712f71f0a7ef5/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/ApiManagementListTags.json
func ExampleTagClient_NewListByServicePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewTagClient().NewListByServicePager("rg1", "apimService1", &armapimanagement.TagClientListByServiceOptions{Filter: nil,
		Top:   nil,
		Skip:  nil,
		Scope: nil,
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
		// page.TagCollection = armapimanagement.TagCollection{
		// 	Count: to.Ptr[int64](2),
		// 	Value: []*armapimanagement.TagContract{
		// 		{
		// 			Name: to.Ptr("5600b59375ff190048020001"),
		// 			Type: to.Ptr("Microsoft.ApiManagement/service/tags"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/tags/5600b59375ff190048020001"),
		// 			Properties: &armapimanagement.TagContractProperties{
		// 				DisplayName: to.Ptr("tag1"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("5600b59375ff190048020002"),
		// 			Type: to.Ptr("Microsoft.ApiManagement/service/tags"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/tags/5600b59375ff190048020002"),
		// 			Properties: &armapimanagement.TagContractProperties{
		// 				DisplayName: to.Ptr("tag2"),
		// 			},
		// 	}},
		// }
	}
}
