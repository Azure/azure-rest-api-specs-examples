package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4cd95123fb961c68740565a1efcaa5e43bd35802/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementListProducts.json
func ExampleProductClient_NewListByServicePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewProductClient().NewListByServicePager("rg1", "apimService1", &armapimanagement.ProductClientListByServiceOptions{Filter: nil,
		Top:          nil,
		Skip:         nil,
		ExpandGroups: nil,
		Tags:         nil,
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
		// page.ProductCollection = armapimanagement.ProductCollection{
		// 	Count: to.Ptr[int64](3),
		// 	Value: []*armapimanagement.ProductContract{
		// 		{
		// 			Name: to.Ptr("kjoshiarmtemplateCert1"),
		// 			Type: to.Ptr("Microsoft.ApiManagement/service/products"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/products/kjoshiarmtemplateCert1"),
		// 			Properties: &armapimanagement.ProductContractProperties{
		// 				Description: to.Ptr("Development Product"),
		// 				State: to.Ptr(armapimanagement.ProductStatePublished),
		// 				SubscriptionRequired: to.Ptr(false),
		// 				DisplayName: to.Ptr("Dev"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("starter"),
		// 			Type: to.Ptr("Microsoft.ApiManagement/service/products"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/products/starter"),
		// 			Properties: &armapimanagement.ProductContractProperties{
		// 				Description: to.Ptr("Subscribers will be able to run 5 calls/minute up to a maximum of 100 calls/week."),
		// 				ApprovalRequired: to.Ptr(false),
		// 				State: to.Ptr(armapimanagement.ProductStatePublished),
		// 				SubscriptionRequired: to.Ptr(true),
		// 				SubscriptionsLimit: to.Ptr[int32](1),
		// 				Terms: to.Ptr(""),
		// 				DisplayName: to.Ptr("Starter"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("unlimited"),
		// 			Type: to.Ptr("Microsoft.ApiManagement/service/products"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/products/unlimited"),
		// 			Properties: &armapimanagement.ProductContractProperties{
		// 				Description: to.Ptr("Subscribers have completely unlimited access to the API. Administrator approval is required."),
		// 				ApprovalRequired: to.Ptr(true),
		// 				State: to.Ptr(armapimanagement.ProductStatePublished),
		// 				SubscriptionRequired: to.Ptr(true),
		// 				SubscriptionsLimit: to.Ptr[int32](1),
		// 				DisplayName: to.Ptr("Unlimited"),
		// 			},
		// 	}},
		// }
	}
}
