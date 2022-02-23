Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fmediaservices%2Farmmediaservices%2Fv0.3.1/sdk/resourcemanager/mediaservices/armmediaservices/README.md) on how to add the SDK to your project and authenticate.

```go
package armmediaservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mediaservices/armmediaservices"
)

// x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-06-01/examples/accounts-create.json
func ExampleClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmediaservices.NewClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<account-name>",
		armmediaservices.MediaService{
			Location: to.StringPtr("<location>"),
			Tags: map[string]*string{
				"key1": to.StringPtr("value1"),
				"key2": to.StringPtr("value2"),
			},
			Identity: &armmediaservices.MediaServiceIdentity{
				Type: to.StringPtr("<type>"),
				UserAssignedIdentities: map[string]*armmediaservices.UserAssignedManagedIdentity{
					"/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/id1": {},
					"/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/id2": {},
				},
			},
			Properties: &armmediaservices.MediaServiceProperties{
				Encryption: &armmediaservices.AccountEncryption{
					Type: armmediaservices.AccountEncryptionKeyType("CustomerKey").ToPtr(),
					Identity: &armmediaservices.ResourceIdentity{
						UseSystemAssignedIdentity: to.BoolPtr(false),
						UserAssignedIdentity:      to.StringPtr("<user-assigned-identity>"),
					},
					KeyVaultProperties: &armmediaservices.KeyVaultProperties{
						KeyIdentifier: to.StringPtr("<key-identifier>"),
					},
				},
				StorageAccounts: []*armmediaservices.StorageAccount{
					{
						Type: armmediaservices.StorageAccountType("Primary").ToPtr(),
						ID:   to.StringPtr("<id>"),
						Identity: &armmediaservices.ResourceIdentity{
							UseSystemAssignedIdentity: to.BoolPtr(false),
							UserAssignedIdentity:      to.StringPtr("<user-assigned-identity>"),
						},
					}},
				StorageAuthentication: armmediaservices.StorageAuthentication("ManagedIdentity").ToPtr(),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ClientCreateOrUpdateResult)
}
```
