package armstorage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage/v4"
)

// Generated from example definition: 2025-08-01/LocalUserCreateNFSv3Enabled.json
func ExampleLocalUsersClient_CreateOrUpdate_createNfSv3EnabledLocalUser() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorage.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewLocalUsersClient().CreateOrUpdate(ctx, "res6977", "sto2527", "user1", armstorage.LocalUser{
		Properties: &armstorage.LocalUserProperties{
			ExtendedGroups: []*int32{
				to.Ptr[int32](1001),
				to.Ptr[int32](1005),
				to.Ptr[int32](2005),
			},
			IsNFSv3Enabled: to.Ptr(true),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armstorage.LocalUsersClientCreateOrUpdateResponse{
	// 	LocalUser: armstorage.LocalUser{
	// 		Name: to.Ptr("user1"),
	// 		Type: to.Ptr("Microsoft.Storage/storageAccounts/localUsers"),
	// 		ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/res6977/providers/Microsoft.Storage/storageAccounts/sto2527/loalUsers/user1"),
	// 		Properties: &armstorage.LocalUserProperties{
	// 			AllowACLAuthorization: to.Ptr(true),
	// 			ExtendedGroups: []*int32{
	// 				to.Ptr[int32](1001),
	// 				to.Ptr[int32](1005),
	// 				to.Ptr[int32](2005),
	// 			},
	// 			GroupID: to.Ptr[int32](2000),
	// 			HomeDirectory: to.Ptr("homedirectory"),
	// 			IsNFSv3Enabled: to.Ptr(true),
	// 			PermissionScopes: []*armstorage.PermissionScope{
	// 				{
	// 					Permissions: to.Ptr("rwd"),
	// 					ResourceName: to.Ptr("share1"),
	// 					Service: to.Ptr("file"),
	// 				},
	// 				{
	// 					Permissions: to.Ptr("rw"),
	// 					ResourceName: to.Ptr("share2"),
	// 					Service: to.Ptr("file"),
	// 				},
	// 			},
	// 			Sid: to.Ptr("S-1-2-0-125132-153423-36235-1000"),
	// 			SSHAuthorizedKeys: []*armstorage.SSHPublicKey{
	// 				{
	// 					Description: to.Ptr("key name"),
	// 					Key: to.Ptr("ssh-rsa keykeykeykeykey="),
	// 				},
	// 			},
	// 			UserID: to.Ptr[int32](1000),
	// 		},
	// 	},
	// }
}
