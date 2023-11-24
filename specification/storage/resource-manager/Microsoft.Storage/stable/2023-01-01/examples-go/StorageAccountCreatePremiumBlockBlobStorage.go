package armstorage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/0baf811c3c76c87b3c127d098519bd97141222dd/specification/storage/resource-manager/Microsoft.Storage/stable/2023-01-01/examples/StorageAccountCreatePremiumBlockBlobStorage.json
func ExampleAccountsClient_BeginCreate_storageAccountCreatePremiumBlockBlobStorage() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorage.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAccountsClient().BeginCreate(ctx, "res9101", "sto4445", armstorage.AccountCreateParameters{
		Kind:     to.Ptr(armstorage.KindBlockBlobStorage),
		Location: to.Ptr("eastus"),
		Properties: &armstorage.AccountPropertiesCreateParameters{
			AllowSharedKeyAccess: to.Ptr(true),
			Encryption: &armstorage.Encryption{
				KeySource:                       to.Ptr(armstorage.KeySourceMicrosoftStorage),
				RequireInfrastructureEncryption: to.Ptr(false),
				Services: &armstorage.EncryptionServices{
					Blob: &armstorage.EncryptionService{
						Enabled: to.Ptr(true),
						KeyType: to.Ptr(armstorage.KeyTypeAccount),
					},
					File: &armstorage.EncryptionService{
						Enabled: to.Ptr(true),
						KeyType: to.Ptr(armstorage.KeyTypeAccount),
					},
				},
			},
			MinimumTLSVersion: to.Ptr(armstorage.MinimumTLSVersionTLS12),
		},
		SKU: &armstorage.SKU{
			Name: to.Ptr(armstorage.SKUNamePremiumLRS),
		},
		Tags: map[string]*string{
			"key1": to.Ptr("value1"),
			"key2": to.Ptr("value2"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Account = armstorage.Account{
	// 	Name: to.Ptr("sto4445"),
	// 	Type: to.Ptr("Microsoft.Storage/storageAccounts"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/res9101/providers/Microsoft.Storage/storageAccounts/sto4445"),
	// 	Location: to.Ptr("eastus"),
	// 	Tags: map[string]*string{
	// 		"key1": to.Ptr("value1"),
	// 		"key2": to.Ptr("value2"),
	// 	},
	// 	Kind: to.Ptr(armstorage.KindStorage),
	// 	Properties: &armstorage.AccountProperties{
	// 		AccessTier: to.Ptr(armstorage.AccessTierPremium),
	// 		AllowBlobPublicAccess: to.Ptr(false),
	// 		CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-05-24T13:25:33.486Z"); return t}()),
	// 		Encryption: &armstorage.Encryption{
	// 			KeySource: to.Ptr(armstorage.KeySourceMicrosoftStorage),
	// 			RequireInfrastructureEncryption: to.Ptr(false),
	// 			Services: &armstorage.EncryptionServices{
	// 				Blob: &armstorage.EncryptionService{
	// 					Enabled: to.Ptr(true),
	// 					KeyType: to.Ptr(armstorage.KeyTypeAccount),
	// 					LastEnabledTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-12-11T20:49:31.703Z"); return t}()),
	// 				},
	// 				File: &armstorage.EncryptionService{
	// 					Enabled: to.Ptr(true),
	// 					KeyType: to.Ptr(armstorage.KeyTypeAccount),
	// 					LastEnabledTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-12-11T20:49:31.703Z"); return t}()),
	// 				},
	// 			},
	// 		},
	// 		MinimumTLSVersion: to.Ptr(armstorage.MinimumTLSVersionTLS12),
	// 		PrimaryEndpoints: &armstorage.Endpoints{
	// 			Blob: to.Ptr("https://sto4445.blob.core.windows.net/"),
	// 			Dfs: to.Ptr("https://sto4445.dfs.core.windows.net/"),
	// 			Web: to.Ptr("https://sto4445.web.core.windows.net/"),
	// 		},
	// 		PrimaryLocation: to.Ptr("eastus2euap"),
	// 		ProvisioningState: to.Ptr(armstorage.ProvisioningStateSucceeded),
	// 		StatusOfPrimary: to.Ptr(armstorage.AccountStatusAvailable),
	// 		StatusOfSecondary: to.Ptr(armstorage.AccountStatusAvailable),
	// 		EnableHTTPSTrafficOnly: to.Ptr(true),
	// 	},
	// 	SKU: &armstorage.SKU{
	// 		Name: to.Ptr(armstorage.SKUNamePremiumLRS),
	// 		Tier: to.Ptr(armstorage.SKUTierPremium),
	// 	},
	// }
}
