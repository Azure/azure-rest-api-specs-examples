package armstorage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/220ad9c6554fc7d6d10a89bdb441c1e3b36e3285/specification/storage/resource-manager/Microsoft.Storage/stable/2023-05-01/examples/LocalUserUpdate.json
func ExampleLocalUsersClient_CreateOrUpdate_updateLocalUser() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorage.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewLocalUsersClient().CreateOrUpdate(ctx, "res6977", "sto2527", "user1", armstorage.LocalUser{
		Properties: &armstorage.LocalUserProperties{
			AllowACLAuthorization: to.Ptr(false),
			ExtendedGroups: []*int32{
				to.Ptr[int32](1001),
				to.Ptr[int32](1005),
				to.Ptr[int32](2005)},
			GroupID:        to.Ptr[int32](3000),
			HasSharedKey:   to.Ptr(false),
			HasSSHKey:      to.Ptr(false),
			HasSSHPassword: to.Ptr(false),
			HomeDirectory:  to.Ptr("homedirectory2"),
			IsNFSv3Enabled: to.Ptr(true),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.LocalUser = armstorage.LocalUser{
	// 	Name: to.Ptr("user1"),
	// 	Type: to.Ptr("Microsoft.Storage/storageAccounts/localUsers"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/res6977/providers/Microsoft.Storage/storageAccounts/sto2527/loalUsers/user1"),
	// 	Properties: &armstorage.LocalUserProperties{
	// 		AllowACLAuthorization: to.Ptr(false),
	// 		ExtendedGroups: []*int32{
	// 			to.Ptr[int32](1001),
	// 			to.Ptr[int32](1005),
	// 			to.Ptr[int32](2005)},
	// 			GroupID: to.Ptr[int32](3000),
	// 			HasSharedKey: to.Ptr(false),
	// 			HasSSHKey: to.Ptr(false),
	// 			HasSSHPassword: to.Ptr(false),
	// 			HomeDirectory: to.Ptr("homedirectory2"),
	// 			IsNFSv3Enabled: to.Ptr(true),
	// 			Sid: to.Ptr("S-1-2-0-3528686663-1788730862-2791910117-1000"),
	// 			UserID: to.Ptr[int32](1000),
	// 		},
	// 	}
}
