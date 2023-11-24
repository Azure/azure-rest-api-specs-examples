package armstorage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/0baf811c3c76c87b3c127d098519bd97141222dd/specification/storage/resource-manager/Microsoft.Storage/stable/2023-01-01/examples/StorageAccountEnableAD.json
func ExampleAccountsClient_Update_storageAccountEnableAd() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorage.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAccountsClient().Update(ctx, "res9407", "sto8596", armstorage.AccountUpdateParameters{
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
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Account = armstorage.Account{
	// 	Name: to.Ptr("sto8596"),
	// 	Type: to.Ptr("Microsoft.Storage/storageAccounts"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/res9407/providers/Microsoft.Storage/storageAccounts/sto8596"),
	// 	Location: to.Ptr("eastus2(stage)"),
	// 	Tags: map[string]*string{
	// 		"key1": to.Ptr("value1"),
	// 		"key2": to.Ptr("value2"),
	// 	},
	// 	Kind: to.Ptr(armstorage.KindStorage),
	// 	Properties: &armstorage.AccountProperties{
	// 		AzureFilesIdentityBasedAuthentication: &armstorage.AzureFilesIdentityBasedAuthentication{
	// 			ActiveDirectoryProperties: &armstorage.ActiveDirectoryProperties{
	// 				AccountType: to.Ptr(armstorage.ActiveDirectoryPropertiesAccountTypeUser),
	// 				AzureStorageSid: to.Ptr("S-1-5-21-2400535526-2334094090-2402026252-0012"),
	// 				DomainGUID: to.Ptr("aebfc118-9fa9-4732-a21f-d98e41a77ae1"),
	// 				DomainName: to.Ptr("adtest.com"),
	// 				DomainSid: to.Ptr("S-1-5-21-2400535526-2334094090-2402026252"),
	// 				ForestName: to.Ptr("adtest.com"),
	// 				NetBiosDomainName: to.Ptr("adtest.com"),
	// 				SamAccountName: to.Ptr("sam12498"),
	// 			},
	// 			DirectoryServiceOptions: to.Ptr(armstorage.DirectoryServiceOptionsAD),
	// 		},
	// 		CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-06-01T02:42:41.763Z"); return t}()),
	// 		PrimaryEndpoints: &armstorage.Endpoints{
	// 			Blob: to.Ptr("https://sto8596.blob.core.windows.net/"),
	// 			Dfs: to.Ptr("https://sto8596.dfs.core.windows.net/"),
	// 			File: to.Ptr("https://sto8596.file.core.windows.net/"),
	// 			Queue: to.Ptr("https://sto8596.queue.core.windows.net/"),
	// 			Table: to.Ptr("https://sto8596.table.core.windows.net/"),
	// 			Web: to.Ptr("https://sto8596.web.core.windows.net/"),
	// 		},
	// 		PrimaryLocation: to.Ptr("eastus2(stage)"),
	// 		ProvisioningState: to.Ptr(armstorage.ProvisioningStateSucceeded),
	// 		SecondaryLocation: to.Ptr("northcentralus(stage)"),
	// 		StatusOfPrimary: to.Ptr(armstorage.AccountStatusAvailable),
	// 		StatusOfSecondary: to.Ptr(armstorage.AccountStatusAvailable),
	// 		EnableHTTPSTrafficOnly: to.Ptr(false),
	// 	},
	// 	SKU: &armstorage.SKU{
	// 		Name: to.Ptr(armstorage.SKUNameStandardGRS),
	// 		Tier: to.Ptr(armstorage.SKUTierStandard),
	// 	},
	// }
}
