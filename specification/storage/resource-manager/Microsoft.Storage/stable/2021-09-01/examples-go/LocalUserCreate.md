Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fstorage%2Farmstorage%2Fv0.6.0/sdk/resourcemanager/storage/armstorage/README.md) on how to add the SDK to your project and authenticate.

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
		return
	}
	ctx := context.Background()
	client, err := armstorage.NewLocalUsersClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<username>",
		armstorage.LocalUser{
			Properties: &armstorage.LocalUserProperties{
				HasSSHPassword: to.Ptr(true),
				HomeDirectory:  to.Ptr("<home-directory>"),
				PermissionScopes: []*armstorage.PermissionScope{
					{
						Permissions:  to.Ptr("<permissions>"),
						ResourceName: to.Ptr("<resource-name>"),
						Service:      to.Ptr("<service>"),
					},
					{
						Permissions:  to.Ptr("<permissions>"),
						ResourceName: to.Ptr("<resource-name>"),
						Service:      to.Ptr("<service>"),
					}},
				SSHAuthorizedKeys: []*armstorage.SSHPublicKey{
					{
						Description: to.Ptr("<description>"),
						Key:         to.Ptr("<key>"),
					}},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
