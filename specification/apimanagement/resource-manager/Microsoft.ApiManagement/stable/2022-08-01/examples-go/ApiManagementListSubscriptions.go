package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4cd95123fb961c68740565a1efcaa5e43bd35802/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementListSubscriptions.json
func ExampleSubscriptionClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSubscriptionClient().NewListPager("rg1", "apimService1", &armapimanagement.SubscriptionClientListOptions{Filter: nil,
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
		// 	Count: to.Ptr[int64](3),
		// 	Value: []*armapimanagement.SubscriptionContract{
		// 		{
		// 			Name: to.Ptr("5600b59475ff190048070001"),
		// 			Type: to.Ptr("Microsoft.ApiManagement/service/subscriptions"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/subscriptions/5600b59475ff190048070001"),
		// 			Properties: &armapimanagement.SubscriptionContractProperties{
		// 				CreatedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2015-09-22T01:57:40.3Z"); return t}()),
		// 				OwnerID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/users/1"),
		// 				Scope: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/products/5600b59475ff190048060001"),
		// 				State: to.Ptr(armapimanagement.SubscriptionStateActive),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("56eaed3dbaf08b06e46d27fe"),
		// 			Type: to.Ptr("Microsoft.ApiManagement/service/subscriptions"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/subscriptions/56eaed3dbaf08b06e46d27fe"),
		// 			Properties: &armapimanagement.SubscriptionContractProperties{
		// 				CreatedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-03-17T17:45:33.837Z"); return t}()),
		// 				DisplayName: to.Ptr("Starter"),
		// 				ExpirationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-04-01T00:00:00Z"); return t}()),
		// 				NotificationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-03-20T00:00:00Z"); return t}()),
		// 				OwnerID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/users/56eaec62baf08b06e46d27fd"),
		// 				Scope: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/products/5600b59475ff190048060001"),
		// 				StartDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-03-17T00:00:00Z"); return t}()),
		// 				State: to.Ptr(armapimanagement.SubscriptionStateActive),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("5931a769d8d14f0ad8ce13b8"),
		// 			Type: to.Ptr("Microsoft.ApiManagement/service/subscriptions"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/subscriptions/5931a769d8d14f0ad8ce13b8"),
		// 			Properties: &armapimanagement.SubscriptionContractProperties{
		// 				CreatedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-06-02T17:59:06.223Z"); return t}()),
		// 				DisplayName: to.Ptr("Unlimited"),
		// 				OwnerID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/users/5931a75ae4bbd512a88c680b"),
		// 				Scope: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/products/5600b59475ff190048060002"),
		// 				State: to.Ptr(armapimanagement.SubscriptionStateSubmitted),
		// 			},
		// 	}},
		// }
	}
}
