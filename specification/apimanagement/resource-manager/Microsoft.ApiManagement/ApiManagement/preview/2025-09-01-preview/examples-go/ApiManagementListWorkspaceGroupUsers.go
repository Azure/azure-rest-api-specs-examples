package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v4"
)

// Generated from example definition: 2025-09-01-preview/ApiManagementListWorkspaceGroupUsers.json
func ExampleWorkspaceGroupUserClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewWorkspaceGroupUserClient().NewListPager("rg1", "apimService1", "wks1", "57d2ef278aa04f0888cba3f3", nil)
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
		// page = armapimanagement.WorkspaceGroupUserClientListResponse{
		// 	UserCollection: armapimanagement.UserCollection{
		// 		Count: to.Ptr[int64](1),
		// 		NextLink: to.Ptr(""),
		// 		Value: []*armapimanagement.UserContract{
		// 			{
		// 				Name: to.Ptr("armTemplateUser1"),
		// 				Type: to.Ptr("Microsoft.ApiManagement/service/workspaces/groups/users"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/workspaces/wks1/users/kjoshiarmTemplateUser1"),
		// 				Properties: &armapimanagement.UserContractProperties{
		// 					Email: to.Ptr("user1@live.com"),
		// 					FirstName: to.Ptr("user1"),
		// 					Identities: []*armapimanagement.UserIdentityContract{
		// 						{
		// 							ID: to.Ptr("user1@live.com"),
		// 							Provider: to.Ptr("Basic"),
		// 						},
		// 					},
		// 					LastName: to.Ptr("lastname1"),
		// 					Note: to.Ptr("note for user 1"),
		// 					RegistrationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-05-31T18:54:41.447Z"); return t}()),
		// 					State: to.Ptr(armapimanagement.UserStateActive),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
