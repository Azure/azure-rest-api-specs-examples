Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fstorage%2Farmstorage%2Fv0.3.0/sdk/resourcemanager/storage/armstorage/README.md) on how to add the SDK to your project and authenticate.

```go
package armstorage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
)

// x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2021-08-01/examples/StorageAccountPatchEncryptionScope.json
func ExampleEncryptionScopesClient_Patch() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armstorage.NewEncryptionScopesClient("<subscription-id>", cred, nil)
	res, err := client.Patch(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<encryption-scope-name>",
		armstorage.EncryptionScope{
			EncryptionScopeProperties: &armstorage.EncryptionScopeProperties{
				KeyVaultProperties: &armstorage.EncryptionScopeKeyVaultProperties{
					KeyURI: to.StringPtr("<key-uri>"),
				},
				Source: armstorage.EncryptionScopeSource("Microsoft.KeyVault").ToPtr(),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.EncryptionScopesClientPatchResult)
}
```
