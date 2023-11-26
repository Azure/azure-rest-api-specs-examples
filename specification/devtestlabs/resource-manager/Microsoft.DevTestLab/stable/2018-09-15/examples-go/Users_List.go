package armdevtestlabs_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devtestlabs/armdevtestlabs"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/Users_List.json
func ExampleUsersClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdevtestlabs.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewUsersClient().NewListPager("resourceGroupName", "{devtestlabName}", &armdevtestlabs.UsersClientListOptions{Expand: nil,
		Filter:  nil,
		Top:     nil,
		Orderby: nil,
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
		// page.UserList = armdevtestlabs.UserList{
		// 	Value: []*armdevtestlabs.User{
		// 		{
		// 			Name: to.Ptr("{userName}"),
		// 			Type: to.Ptr("Microsoft.DevTestLab/labs/users"),
		// 			ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/resourceGroupName/providers/Microsoft.DevTestLab/labs/{labName}/users/{userName}"),
		// 			Location: to.Ptr("{location}"),
		// 			Tags: map[string]*string{
		// 				"tagName1": to.Ptr("tagValue1"),
		// 			},
		// 			Properties: &armdevtestlabs.UserProperties{
		// 				CreatedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-10-02T01:40:48.173Z"); return t}()),
		// 				Identity: &armdevtestlabs.UserIdentity{
		// 					AppID: to.Ptr("{appId}"),
		// 					ObjectID: to.Ptr("{objectId}"),
		// 					PrincipalID: to.Ptr("{principalId}"),
		// 					PrincipalName: to.Ptr("{principalName}"),
		// 					TenantID: to.Ptr("{tenantId}"),
		// 				},
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				SecretStore: &armdevtestlabs.UserSecretStore{
		// 					KeyVaultID: to.Ptr("{keyVaultId}"),
		// 					KeyVaultURI: to.Ptr("{keyVaultUri}"),
		// 				},
		// 				UniqueIdentifier: to.Ptr("{uniqueIdentifier}"),
		// 			},
		// 	}},
		// }
	}
}
