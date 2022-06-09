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
	}
	ctx := context.Background()
	client, err := armstorage.NewAccountsClient("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Update(ctx,
		"res9407",
		"sto8596",
		armstorage.AccountUpdateParameters{
			Properties: &armstorage.AccountPropertiesUpdateParameters{
				AzureFilesIdentityBasedAuthentication: &armstorage.AzureFilesIdentityBasedAuthentication{
					ActiveDirectoryProperties: &armstorage.ActiveDirectoryProperties{
						AccountType:       to.Ptr(armstorage.ActiveDirectoryPropertiesAccountTypeUser),
						AzureStorageSid:   to.Ptr("S-1-5-21-2400535526-2334094090-2402026252-0012"),
						DomainGUID:        to.Ptr("aebfc118-9fa9-4732-a21f-d98e41a77ae1"),
						DomainName:        to.Ptr("adtest.com"),
						DomainSid:         to.Ptr("S-1-5-21-2400535526-2334094090-2402026252"),
						ForestName:        to.Ptr("adtest.com"),
						NetBiosDomainName: to.Ptr("adtest.com"),
						SamAccountName:    to.Ptr("sam12498"),
					},
					DirectoryServiceOptions: to.Ptr(armstorage.DirectoryServiceOptionsAD),
				},
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
