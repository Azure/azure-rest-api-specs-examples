package armstorage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/0baf811c3c76c87b3c127d098519bd97141222dd/specification/storage/resource-manager/Microsoft.Storage/stable/2023-01-01/examples/LocalUserUpdate.json
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
			HasSharedKey:   to.Ptr(false),
			HasSSHKey:      to.Ptr(false),
			HasSSHPassword: to.Ptr(false),
			HomeDirectory:  to.Ptr("homedirectory2"),
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
	// 		HasSharedKey: to.Ptr(false),
	// 		HasSSHKey: to.Ptr(false),
	// 		HasSSHPassword: to.Ptr(false),
	// 		HomeDirectory: to.Ptr("homedirectory2"),
	// 		Sid: to.Ptr("S-1-2-0-3528686663-1788730862-2791910117-1000"),
	// 	},
	// }
}
