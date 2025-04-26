package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e436160e64c0f8d7fb20d662be2712f71f0a7ef5/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/ApiManagementListApiTagDescriptions.json
func ExampleAPITagDescriptionClient_NewListByServicePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAPITagDescriptionClient().NewListByServicePager("rg1", "apimService1", "57d2ef278aa04f0888cba3f3", &armapimanagement.APITagDescriptionClientListByServiceOptions{Filter: nil,
		Top:  nil,
		Skip: nil,
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
		// page.TagDescriptionCollection = armapimanagement.TagDescriptionCollection{
		// 	Count: to.Ptr[int64](1),
		// 	Value: []*armapimanagement.TagDescriptionContract{
		// 		{
		// 			Name: to.Ptr("5600b539c53f5b0062060002"),
		// 			Type: to.Ptr("Microsoft.ApiManagement/service/tags"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/tagDescriptions/5600b539c53f5b0062060002"),
		// 			Properties: &armapimanagement.TagDescriptionContractProperties{
		// 				ExternalDocsDescription: to.Ptr("some additional info"),
		// 				ExternalDocsURL: to.Ptr("http://some_url.com"),
		// 				DisplayName: to.Ptr("tag1"),
		// 				TagID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/tags/5600b539c53f5b0062060002"),
		// 			},
		// 	}},
		// }
	}
}
