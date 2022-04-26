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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/storage/resource-manager/Microsoft.Storage/stable/2021-09-01/examples/StorageAccountEnableAD.json
func ExampleAccountsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armstorage.NewAccountsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.Update(ctx,
		"<resource-group-name>",
		"<account-name>",
		armstorage.AccountUpdateParameters{
			Properties: &armstorage.AccountPropertiesUpdateParameters{
				AzureFilesIdentityBasedAuthentication: &armstorage.AzureFilesIdentityBasedAuthentication{
					ActiveDirectoryProperties: &armstorage.ActiveDirectoryProperties{
						AccountType:       to.Ptr(armstorage.ActiveDirectoryPropertiesAccountTypeUser),
						AzureStorageSid:   to.Ptr("<azure-storage-sid>"),
						DomainGUID:        to.Ptr("<domain-guid>"),
						DomainName:        to.Ptr("<domain-name>"),
						DomainSid:         to.Ptr("<domain-sid>"),
						ForestName:        to.Ptr("<forest-name>"),
						NetBiosDomainName: to.Ptr("<net-bios-domain-name>"),
						SamAccountName:    to.Ptr("<sam-account-name>"),
					},
					DirectoryServiceOptions: to.Ptr(armstorage.DirectoryServiceOptionsAD),
				},
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
