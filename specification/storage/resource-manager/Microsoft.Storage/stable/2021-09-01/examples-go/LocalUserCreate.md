```go
package armstorage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/storage/resource-manager/Microsoft.Storage/stable/2021-09-01/examples/LocalUserCreate.json
func ExampleLocalUsersClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armstorage.NewLocalUsersClient("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx,
		"res6977",
		"sto2527",
		"user1",
		armstorage.LocalUser{
			Properties: &armstorage.LocalUserProperties{
				HasSSHPassword: to.Ptr(true),
				HomeDirectory:  to.Ptr("homedirectory"),
				PermissionScopes: []*armstorage.PermissionScope{
					{
						Permissions:  to.Ptr("rwd"),
						ResourceName: to.Ptr("share1"),
						Service:      to.Ptr("file"),
					},
					{
						Permissions:  to.Ptr("rw"),
						ResourceName: to.Ptr("share2"),
						Service:      to.Ptr("file"),
					}},
				SSHAuthorizedKeys: []*armstorage.SSHPublicKey{
					{
						Description: to.Ptr("key name"),
						Key:         to.Ptr("ssh-rsa keykeykeykeykey="),
					}},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fstorage%2Farmstorage%2Fv1.0.0/sdk/resourcemanager/storage/armstorage/README.md) on how to add the SDK to your project and authenticate.
