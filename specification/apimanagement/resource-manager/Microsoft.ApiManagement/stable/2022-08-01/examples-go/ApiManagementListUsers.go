package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4cd95123fb961c68740565a1efcaa5e43bd35802/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementListUsers.json
func ExampleUserClient_NewListByServicePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewUserClient().NewListByServicePager("rg1", "apimService1", &armapimanagement.UserClientListByServiceOptions{Filter: nil,
		Top:          nil,
		Skip:         nil,
		ExpandGroups: nil,
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
		// page.UserCollection = armapimanagement.UserCollection{
		// 	Count: to.Ptr[int64](3),
		// 	Value: []*armapimanagement.UserContract{
		// 		{
		// 			Name: to.Ptr("1"),
		// 			Type: to.Ptr("Microsoft.ApiManagement/service/users"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/users/1"),
		// 			Properties: &armapimanagement.UserContractProperties{
		// 				Identities: []*armapimanagement.UserIdentityContract{
		// 					{
		// 						ID: to.Ptr("admin@live.com"),
		// 						Provider: to.Ptr("Azure"),
		// 				}},
		// 				State: to.Ptr(armapimanagement.UserStateActive),
		// 				Email: to.Ptr("admin@live.com"),
		// 				FirstName: to.Ptr("Administrator"),
		// 				LastName: to.Ptr(""),
		// 				RegistrationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2015-09-22T01:57:39.677Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("56eaec62baf08b06e46d27fd"),
		// 			Type: to.Ptr("Microsoft.ApiManagement/service/users"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/users/56eaec62baf08b06e46d27fd"),
		// 			Properties: &armapimanagement.UserContractProperties{
		// 				Identities: []*armapimanagement.UserIdentityContract{
		// 					{
		// 						ID: to.Ptr("foo.bar.83@gmail.com"),
		// 						Provider: to.Ptr("Basic"),
		// 				}},
		// 				State: to.Ptr(armapimanagement.UserStateActive),
		// 				Email: to.Ptr("foo.bar.83@gmail.com"),
		// 				FirstName: to.Ptr("foo"),
		// 				LastName: to.Ptr("bar"),
		// 				RegistrationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-03-17T17:41:56.327Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("5931a75ae4bbd512a88c680b"),
		// 			Type: to.Ptr("Microsoft.ApiManagement/service/users"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/users/5931a75ae4bbd512a88c680b"),
		// 			Properties: &armapimanagement.UserContractProperties{
		// 				Identities: []*armapimanagement.UserIdentityContract{
		// 					{
		// 						ID: to.Ptr("*************"),
		// 						Provider: to.Ptr("Microsoft"),
		// 				}},
		// 				State: to.Ptr(armapimanagement.UserStateActive),
		// 				Email: to.Ptr("foobar@outlook.com"),
		// 				FirstName: to.Ptr("foo"),
		// 				LastName: to.Ptr("bar"),
		// 				RegistrationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-06-02T17:58:50.357Z"); return t}()),
		// 			},
		// 	}},
		// }
	}
}
