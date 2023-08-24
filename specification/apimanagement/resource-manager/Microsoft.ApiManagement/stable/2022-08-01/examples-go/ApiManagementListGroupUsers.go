package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4cd95123fb961c68740565a1efcaa5e43bd35802/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementListGroupUsers.json
func ExampleGroupUserClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewGroupUserClient().NewListPager("rg1", "apimService1", "57d2ef278aa04f0888cba3f3", &armapimanagement.GroupUserClientListOptions{Filter: nil,
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
		// page.UserCollection = armapimanagement.UserCollection{
		// 	Count: to.Ptr[int64](1),
		// 	Value: []*armapimanagement.UserContract{
		// 		{
		// 			Name: to.Ptr("armTemplateUser1"),
		// 			Type: to.Ptr("Microsoft.ApiManagement/service/groups/users"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/users/kjoshiarmTemplateUser1"),
		// 			Properties: &armapimanagement.UserContractProperties{
		// 				Identities: []*armapimanagement.UserIdentityContract{
		// 					{
		// 						ID: to.Ptr("user1@live.com"),
		// 						Provider: to.Ptr("Basic"),
		// 				}},
		// 				Note: to.Ptr("note for user 1"),
		// 				State: to.Ptr(armapimanagement.UserStateActive),
		// 				Email: to.Ptr("user1@live.com"),
		// 				FirstName: to.Ptr("user1"),
		// 				LastName: to.Ptr("lastname1"),
		// 				RegistrationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-05-31T18:54:41.447Z"); return t}()),
		// 			},
		// 	}},
		// }
	}
}
