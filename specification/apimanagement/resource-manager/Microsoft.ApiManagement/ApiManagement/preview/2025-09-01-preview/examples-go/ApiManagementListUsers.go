package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v4"
)

// Generated from example definition: 2025-09-01-preview/ApiManagementListUsers.json
func ExampleUserClient_NewListByServicePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewUserClient().NewListByServicePager("rg1", "apimService1", nil)
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
		// page = armapimanagement.UserClientListByServiceResponse{
		// 	UserCollection: armapimanagement.UserCollection{
		// 		Count: to.Ptr[int64](3),
		// 		NextLink: to.Ptr(""),
		// 		Value: []*armapimanagement.UserContract{
		// 			{
		// 				Name: to.Ptr("1"),
		// 				Type: to.Ptr("Microsoft.ApiManagement/service/users"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/users/1"),
		// 				Properties: &armapimanagement.UserContractProperties{
		// 					Email: to.Ptr("admin@live.com"),
		// 					FirstName: to.Ptr("Administrator"),
		// 					Identities: []*armapimanagement.UserIdentityContract{
		// 						{
		// 							ID: to.Ptr("admin@live.com"),
		// 							Provider: to.Ptr("Azure"),
		// 						},
		// 					},
		// 					LastName: to.Ptr(""),
		// 					RegistrationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2015-09-22T01:57:39.677Z"); return t}()),
		// 					State: to.Ptr(armapimanagement.UserStateActive),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("56eaec62baf08b06e46d27fd"),
		// 				Type: to.Ptr("Microsoft.ApiManagement/service/users"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/users/56eaec62baf08b06e46d27fd"),
		// 				Properties: &armapimanagement.UserContractProperties{
		// 					Email: to.Ptr("foo.bar.83@gmail.com"),
		// 					FirstName: to.Ptr("foo"),
		// 					Identities: []*armapimanagement.UserIdentityContract{
		// 						{
		// 							ID: to.Ptr("foo.bar.83@gmail.com"),
		// 							Provider: to.Ptr("Basic"),
		// 						},
		// 					},
		// 					LastName: to.Ptr("bar"),
		// 					RegistrationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-03-17T17:41:56.327Z"); return t}()),
		// 					State: to.Ptr(armapimanagement.UserStateActive),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("5931a75ae4bbd512a88c680b"),
		// 				Type: to.Ptr("Microsoft.ApiManagement/service/users"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/users/5931a75ae4bbd512a88c680b"),
		// 				Properties: &armapimanagement.UserContractProperties{
		// 					Email: to.Ptr("foobar@outlook.com"),
		// 					FirstName: to.Ptr("foo"),
		// 					Identities: []*armapimanagement.UserIdentityContract{
		// 						{
		// 							ID: to.Ptr("*************"),
		// 							Provider: to.Ptr("Microsoft"),
		// 						},
		// 					},
		// 					LastName: to.Ptr("bar"),
		// 					RegistrationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-06-02T17:58:50.357Z"); return t}()),
		// 					State: to.Ptr(armapimanagement.UserStateActive),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
