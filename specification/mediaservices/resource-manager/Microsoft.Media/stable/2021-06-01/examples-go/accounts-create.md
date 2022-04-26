Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fmediaservices%2Farmmediaservices%2Fv0.6.0/sdk/resourcemanager/mediaservices/armmediaservices/README.md) on how to add the SDK to your project and authenticate.

```go
package armmediaservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mediaservices/armmediaservices"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-06-01/examples/accounts-create.json
func ExampleClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armmediaservices.NewClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<account-name>",
		armmediaservices.MediaService{
			Location: to.Ptr("<location>"),
			Tags: map[string]*string{
				"key1": to.Ptr("value1"),
				"key2": to.Ptr("value2"),
			},
			Identity: &armmediaservices.MediaServiceIdentity{
				Type: to.Ptr("<type>"),
				UserAssignedIdentities: map[string]*armmediaservices.UserAssignedManagedIdentity{
					"/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/id1": {},
					"/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/id2": {},
				},
			},
			Properties: &armmediaservices.MediaServiceProperties{
				Encryption: &armmediaservices.AccountEncryption{
					Type: to.Ptr(armmediaservices.AccountEncryptionKeyTypeCustomerKey),
					Identity: &armmediaservices.ResourceIdentity{
						UseSystemAssignedIdentity: to.Ptr(false),
						UserAssignedIdentity:      to.Ptr("<user-assigned-identity>"),
					},
					KeyVaultProperties: &armmediaservices.KeyVaultProperties{
						KeyIdentifier: to.Ptr("<key-identifier>"),
					},
				},
				StorageAccounts: []*armmediaservices.StorageAccount{
					{
						Type: to.Ptr(armmediaservices.StorageAccountTypePrimary),
						ID:   to.Ptr("<id>"),
						Identity: &armmediaservices.ResourceIdentity{
							UseSystemAssignedIdentity: to.Ptr(false),
							UserAssignedIdentity:      to.Ptr("<user-assigned-identity>"),
						},
					}},
				StorageAuthentication: to.Ptr(armmediaservices.StorageAuthenticationManagedIdentity),
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
