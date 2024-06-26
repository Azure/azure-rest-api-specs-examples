package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c767823fdfd9d5e96bad245e3ea4d14d94a716bb/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementListUserSubscriptions.json
func ExampleUserSubscriptionClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewUserSubscriptionClient().NewListPager("rg1", "apimService1", "57681833a40f7eb6c49f6acf", &armapimanagement.UserSubscriptionClientListOptions{Filter: nil,
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
		// page.SubscriptionCollection = armapimanagement.SubscriptionCollection{
		// 	Value: []*armapimanagement.SubscriptionContract{
		// 		{
		// 			Name: to.Ptr("57681850a40f7eb6c49f6ae3"),
		// 			Type: to.Ptr("Microsoft.ApiManagement/service/users/subscriptions"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/users/57681833a40f7eb6c49f6acf/subscriptions/57681850a40f7eb6c49f6ae3"),
		// 			Properties: &armapimanagement.SubscriptionContractProperties{
		// 				CreatedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-06-20T16:22:39.547Z"); return t}()),
		// 				DisplayName: to.Ptr("57681850a40f7eb6c49f6ae5"),
		// 				OwnerID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/users/57681833a40f7eb6c49f6acf"),
		// 				Scope: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/products/5768181ea40f7eb6c49f6ac7"),
		// 				StartDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-06-20T00:00:00Z"); return t}()),
		// 				State: to.Ptr(armapimanagement.SubscriptionStateActive),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("57681850a40f7eb6c49f6b2b"),
		// 			Type: to.Ptr("Microsoft.ApiManagement/service/users/subscriptions"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/users/57681833a40f7eb6c49f6acf/subscriptions/57681850a40f7eb6c49f6b2b"),
		// 			Properties: &armapimanagement.SubscriptionContractProperties{
		// 				CreatedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-06-20T16:22:41.103Z"); return t}()),
		// 				DisplayName: to.Ptr("57681850a40f7eb6c49f6b2d"),
		// 				OwnerID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/users/57681833a40f7eb6c49f6acf"),
		// 				Scope: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/products/5768181ea40f7eb6c49f6ac7"),
		// 				StartDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-06-20T00:00:00Z"); return t}()),
		// 				State: to.Ptr(armapimanagement.SubscriptionStateActive),
		// 			},
		// 	}},
		// }
	}
}
