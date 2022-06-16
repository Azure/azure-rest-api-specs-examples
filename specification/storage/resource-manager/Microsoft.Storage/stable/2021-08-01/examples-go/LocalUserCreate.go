package armstorage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
)

// x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2021-08-01/examples/LocalUserCreate.json
func ExampleLocalUsersClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armstorage.NewLocalUsersClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<username>",
		armstorage.LocalUser{
			Properties: &armstorage.LocalUserProperties{
				HasSSHPassword: to.BoolPtr(true),
				HomeDirectory:  to.StringPtr("<home-directory>"),
				PermissionScopes: []*armstorage.PermissionScope{
					{
						Permissions:  to.StringPtr("<permissions>"),
						ResourceName: to.StringPtr("<resource-name>"),
						Service:      to.StringPtr("<service>"),
					},
					{
						Permissions:  to.StringPtr("<permissions>"),
						ResourceName: to.StringPtr("<resource-name>"),
						Service:      to.StringPtr("<service>"),
					}},
				SSHAuthorizedKeys: []*armstorage.SSHPublicKey{
					{
						Description: to.StringPtr("<description>"),
						Key:         to.StringPtr("<key>"),
					}},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.LocalUsersClientCreateOrUpdateResult)
}
