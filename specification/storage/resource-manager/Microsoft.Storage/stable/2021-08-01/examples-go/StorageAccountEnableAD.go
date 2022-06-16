package armstorage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
)

// x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2021-08-01/examples/StorageAccountEnableAD.json
func ExampleAccountsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armstorage.NewAccountsClient("<subscription-id>", cred, nil)
	res, err := client.Update(ctx,
		"<resource-group-name>",
		"<account-name>",
		armstorage.AccountUpdateParameters{
			Properties: &armstorage.AccountPropertiesUpdateParameters{
				AzureFilesIdentityBasedAuthentication: &armstorage.AzureFilesIdentityBasedAuthentication{
					ActiveDirectoryProperties: &armstorage.ActiveDirectoryProperties{
						AccountType:       armstorage.ActiveDirectoryPropertiesAccountType("User").ToPtr(),
						AzureStorageSid:   to.StringPtr("<azure-storage-sid>"),
						DomainGUID:        to.StringPtr("<domain-guid>"),
						DomainName:        to.StringPtr("<domain-name>"),
						DomainSid:         to.StringPtr("<domain-sid>"),
						ForestName:        to.StringPtr("<forest-name>"),
						NetBiosDomainName: to.StringPtr("<net-bios-domain-name>"),
						SamAccountName:    to.StringPtr("<sam-account-name>"),
					},
					DirectoryServiceOptions: armstorage.DirectoryServiceOptions("AD").ToPtr(),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.AccountsClientUpdateResult)
}
