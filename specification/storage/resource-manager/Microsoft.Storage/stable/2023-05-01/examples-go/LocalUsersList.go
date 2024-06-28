package armstorage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/220ad9c6554fc7d6d10a89bdb441c1e3b36e3285/specification/storage/resource-manager/Microsoft.Storage/stable/2023-05-01/examples/LocalUsersList.json
func ExampleLocalUsersClient_NewListPager_listLocalUsers() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorage.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewLocalUsersClient().NewListPager("res6977", "sto2527", &armstorage.LocalUsersClientListOptions{Maxpagesize: nil,
		Filter:  nil,
		Include: nil,
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
		// page.LocalUsers = armstorage.LocalUsers{
		// 	Value: []*armstorage.LocalUser{
		// 		{
		// 			Name: to.Ptr("user1"),
		// 			Type: to.Ptr("Microsoft.Storage/storageAccounts/localUsers"),
		// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/res6977/providers/Microsoft.Storage/storageAccounts/sto2527/loalUsers/user1"),
		// 			Properties: &armstorage.LocalUserProperties{
		// 				AllowACLAuthorization: to.Ptr(true),
		// 				GroupID: to.Ptr[int32](2000),
		// 				HasSharedKey: to.Ptr(true),
		// 				HasSSHKey: to.Ptr(true),
		// 				HasSSHPassword: to.Ptr(true),
		// 				HomeDirectory: to.Ptr("homedirectory"),
		// 				PermissionScopes: []*armstorage.PermissionScope{
		// 					{
		// 						Permissions: to.Ptr("rwd"),
		// 						ResourceName: to.Ptr("share1"),
		// 						Service: to.Ptr("file"),
		// 					},
		// 					{
		// 						Permissions: to.Ptr("rw"),
		// 						ResourceName: to.Ptr("share2"),
		// 						Service: to.Ptr("file"),
		// 				}},
		// 				Sid: to.Ptr("S-1-2-0-125132-153423-36235-1000"),
		// 				UserID: to.Ptr[int32](1000),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("user2"),
		// 			Type: to.Ptr("Microsoft.Storage/storageAccounts/localUsers"),
		// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/res9290/providers/Microsoft.Storage/storageAccounts/sto1590/loalUsers/user2"),
		// 			Properties: &armstorage.LocalUserProperties{
		// 				AllowACLAuthorization: to.Ptr(true),
		// 				GroupID: to.Ptr[int32](2000),
		// 				HasSharedKey: to.Ptr(true),
		// 				HasSSHKey: to.Ptr(false),
		// 				HasSSHPassword: to.Ptr(true),
		// 				PermissionScopes: []*armstorage.PermissionScope{
		// 					{
		// 						Permissions: to.Ptr("rw"),
		// 						ResourceName: to.Ptr("resourcename"),
		// 						Service: to.Ptr("blob"),
		// 				}},
		// 				Sid: to.Ptr("S-1-2-0-533672-235636-66334-1001"),
		// 				UserID: to.Ptr[int32](1001),
		// 			},
		// 	}},
		// }
	}
}
