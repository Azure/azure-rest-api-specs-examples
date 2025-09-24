package armstorage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/260ed6a52537921f53a18ffaf4020e3b4d510367/specification/storage/resource-manager/Microsoft.Storage/stable/2025-01-01/examples/StorageAccountCreateDisallowPublicNetworkAccess.json
func ExampleAccountsClient_BeginCreate_storageAccountCreateDisallowPublicNetworkAccess() {
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
		ExtendedLocation: &armstorage.ExtendedLocation{
			Name: to.Ptr("losangeles001"),
			Type: to.Ptr(armstorage.ExtendedLocationTypesEdgeZone),
		},
		Kind:     to.Ptr(armstorage.KindStorage),
		Location: to.Ptr("eastus"),
		Properties: &armstorage.AccountPropertiesCreateParameters{
			AllowBlobPublicAccess: to.Ptr(false),
			AllowSharedKeyAccess:  to.Ptr(true),
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
			IsHnsEnabled: to.Ptr(true),
			KeyPolicy: &armstorage.KeyPolicy{
				KeyExpirationPeriodInDays: to.Ptr[int32](20),
			},
			MinimumTLSVersion:   to.Ptr(armstorage.MinimumTLSVersionTLS12),
			PublicNetworkAccess: to.Ptr(armstorage.PublicNetworkAccessDisabled),
			RoutingPreference: &armstorage.RoutingPreference{
				PublishInternetEndpoints:  to.Ptr(true),
				PublishMicrosoftEndpoints: to.Ptr(true),
				RoutingChoice:             to.Ptr(armstorage.RoutingChoiceMicrosoftRouting),
			},
			SasPolicy: &armstorage.SasPolicy{
				ExpirationAction:    to.Ptr(armstorage.ExpirationActionLog),
				SasExpirationPeriod: to.Ptr("1.15:59:59"),
			},
		},
		SKU: &armstorage.SKU{
			Name: to.Ptr(armstorage.SKUNameStandardGRS),
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
	// 	ExtendedLocation: &armstorage.ExtendedLocation{
	// 		Name: to.Ptr("losangeles001"),
	// 		Type: to.Ptr(armstorage.ExtendedLocationTypesEdgeZone),
	// 	},
	// 	Kind: to.Ptr(armstorage.KindStorage),
	// 	Properties: &armstorage.AccountProperties{
	// 		AllowBlobPublicAccess: to.Ptr(false),
	// 		AllowSharedKeyAccess: to.Ptr(true),
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
	// 		IsHnsEnabled: to.Ptr(true),
	// 		KeyCreationTime: &armstorage.KeyCreationTime{
	// 			Key1: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-03-18T04:42:22.432Z"); return t}()),
	// 			Key2: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-03-18T04:42:22.432Z"); return t}()),
	// 		},
	// 		KeyPolicy: &armstorage.KeyPolicy{
	// 			KeyExpirationPeriodInDays: to.Ptr[int32](20),
	// 		},
	// 		MinimumTLSVersion: to.Ptr(armstorage.MinimumTLSVersionTLS12),
	// 		PrimaryEndpoints: &armstorage.Endpoints{
	// 			Blob: to.Ptr("https://sto4445.blob.core.windows.net/"),
	// 			Dfs: to.Ptr("https://sto4445.dfs.core.windows.net/"),
	// 			File: to.Ptr("https://sto4445.file.core.windows.net/"),
	// 			InternetEndpoints: &armstorage.AccountInternetEndpoints{
	// 				Blob: to.Ptr("https://sto4445-internetrouting.blob.core.windows.net/"),
	// 				Dfs: to.Ptr("https://sto4445-internetrouting.dfs.core.windows.net/"),
	// 				File: to.Ptr("https://sto4445-internetrouting.file.core.windows.net/"),
	// 				Web: to.Ptr("https://sto4445-internetrouting.web.core.windows.net/"),
	// 			},
	// 			MicrosoftEndpoints: &armstorage.AccountMicrosoftEndpoints{
	// 				Blob: to.Ptr("https://sto4445-microsoftrouting.blob.core.windows.net/"),
	// 				Dfs: to.Ptr("https://sto4445-microsoftrouting.dfs.core.windows.net/"),
	// 				File: to.Ptr("https://sto4445-microsoftrouting.file.core.windows.net/"),
	// 				Queue: to.Ptr("https://sto4445-microsoftrouting.queue.core.windows.net/"),
	// 				Table: to.Ptr("https://sto4445-microsoftrouting.table.core.windows.net/"),
	// 				Web: to.Ptr("https://sto4445-microsoftrouting.web.core.windows.net/"),
	// 			},
	// 			Queue: to.Ptr("https://sto4445.queue.core.windows.net/"),
	// 			Table: to.Ptr("https://sto4445.table.core.windows.net/"),
	// 			Web: to.Ptr("https://sto4445.web.core.windows.net/"),
	// 		},
	// 		PrimaryLocation: to.Ptr("eastus2euap"),
	// 		ProvisioningState: to.Ptr(armstorage.ProvisioningStateSucceeded),
	// 		PublicNetworkAccess: to.Ptr(armstorage.PublicNetworkAccessDisabled),
	// 		RoutingPreference: &armstorage.RoutingPreference{
	// 			PublishInternetEndpoints: to.Ptr(true),
	// 			PublishMicrosoftEndpoints: to.Ptr(true),
	// 			RoutingChoice: to.Ptr(armstorage.RoutingChoiceMicrosoftRouting),
	// 		},
	// 		SasPolicy: &armstorage.SasPolicy{
	// 			ExpirationAction: to.Ptr(armstorage.ExpirationActionLog),
	// 			SasExpirationPeriod: to.Ptr("1.15:59:59"),
	// 		},
	// 		SecondaryLocation: to.Ptr("centraluseuap"),
	// 		StatusOfPrimary: to.Ptr(armstorage.AccountStatusAvailable),
	// 		StatusOfSecondary: to.Ptr(armstorage.AccountStatusAvailable),
	// 		EnableHTTPSTrafficOnly: to.Ptr(true),
	// 	},
	// 	SKU: &armstorage.SKU{
	// 		Name: to.Ptr(armstorage.SKUNameStandardGRS),
	// 		Tier: to.Ptr(armstorage.SKUTierStandard),
	// 	},
	// }
}
