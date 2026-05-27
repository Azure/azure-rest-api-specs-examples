package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v4"
)

// Generated from example definition: 2025-09-01-preview/ApiManagementListProductsByTags.json
func ExampleProductClient_NewListByTagsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewProductClient().NewListByTagsPager("rg1", "apimService1", nil)
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
		// page = armapimanagement.ProductClientListByTagsResponse{
		// 	TagResourceCollection: armapimanagement.TagResourceCollection{
		// 		Count: to.Ptr[int64](1),
		// 		Value: []*armapimanagement.TagResourceContract{
		// 			{
		// 				Product: &armapimanagement.ProductTagResourceContractProperties{
		// 					Name: to.Ptr("Starter"),
		// 					Description: to.Ptr("Subscribers will be able to run 5 calls/minute up to a maximum of 100 calls/week."),
		// 					ApprovalRequired: to.Ptr(false),
		// 					ID: to.Ptr("/products/starter"),
		// 					State: to.Ptr(armapimanagement.ProductStatePublished),
		// 					SubscriptionRequired: to.Ptr(true),
		// 					SubscriptionsLimit: to.Ptr[int32](1),
		// 					Terms: to.Ptr(""),
		// 				},
		// 				Tag: &armapimanagement.TagResourceContractProperties{
		// 					Name: to.Ptr("awesomeTag"),
		// 					ID: to.Ptr("/tags/apitag123"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
