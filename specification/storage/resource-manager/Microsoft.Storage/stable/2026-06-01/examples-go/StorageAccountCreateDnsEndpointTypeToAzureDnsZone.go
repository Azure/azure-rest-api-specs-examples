package armstorage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage/v4"
)

// Generated from example definition: 2026-06-01/StorageAccountCreateDnsEndpointTypeToAzureDnsZone.json
func ExampleAccountsClient_BeginCreate_storageAccountCreateDnsEndpointTypeToAzureDnsZone() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorage.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
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
			AllowBlobPublicAccess:        to.Ptr(false),
			AllowSharedKeyAccess:         to.Ptr(true),
			DefaultToOAuthAuthentication: to.Ptr(false),
			DNSEndpointType:              to.Ptr(armstorage.DNSEndpointTypeAzureDNSZone),
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
			IsHnsEnabled:  to.Ptr(true),
			IsSftpEnabled: to.Ptr(true),
			KeyPolicy: &armstorage.KeyPolicy{
				KeyExpirationPeriodInDays: to.Ptr[int32](20),
			},
			MinimumTLSVersion: to.Ptr(armstorage.MinimumTLSVersionTLS12),
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
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armstorage.AccountsClientCreateResponse{
	// 	Account: armstorage.Account{
	// 		Name: to.Ptr("sto4445"),
	// 		Type: to.Ptr("Microsoft.Storage/storageAccounts"),
	// 		ExtendedLocation: &armstorage.ExtendedLocation{
	// 			Name: to.Ptr("losangeles001"),
	// 			Type: to.Ptr(armstorage.ExtendedLocationTypesEdgeZone),
	// 		},
	// 		ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/res9101/providers/Microsoft.Storage/storageAccounts/sto4445"),
	// 		Kind: to.Ptr(armstorage.KindStorage),
	// 		Location: to.Ptr("eastus"),
	// 		Properties: &armstorage.AccountProperties{
	// 			AllowBlobPublicAccess: to.Ptr(false),
	// 			AllowSharedKeyAccess: to.Ptr(true),
	// 			CreationTime: to.Ptr(time.Date(2017, time.May, 24, 13, 25, 33, 486323600, time.UTC)),
	// 			DNSEndpointType: to.Ptr(armstorage.DNSEndpointTypeAzureDNSZone),
	// 			Encryption: &armstorage.Encryption{
	// 				KeySource: to.Ptr(armstorage.KeySourceMicrosoftStorage),
	// 				RequireInfrastructureEncryption: to.Ptr(false),
	// 				Services: &armstorage.EncryptionServices{
	// 					Blob: &armstorage.EncryptionService{
	// 						Enabled: to.Ptr(true),
	// 						KeyType: to.Ptr(armstorage.KeyTypeAccount),
	// 						LastEnabledTime: to.Ptr(time.Date(2019, time.December, 11, 20, 49, 31, 703614000, time.UTC)),
	// 					},
	// 					File: &armstorage.EncryptionService{
	// 						Enabled: to.Ptr(true),
	// 						KeyType: to.Ptr(armstorage.KeyTypeAccount),
	// 						LastEnabledTime: to.Ptr(time.Date(2019, time.December, 11, 20, 49, 31, 703614000, time.UTC)),
	// 					},
	// 				},
	// 			},
	// 			IsHnsEnabled: to.Ptr(true),
	// 			IsSftpEnabled: to.Ptr(true),
	// 			KeyCreationTime: &armstorage.KeyCreationTime{
	// 				Key1: to.Ptr(time.Date(2021, time.March, 18, 4, 42, 22, 432283600, time.UTC)),
	// 				Key2: to.Ptr(time.Date(2021, time.March, 18, 4, 42, 22, 432283600, time.UTC)),
	// 			},
	// 			KeyPolicy: &armstorage.KeyPolicy{
	// 				KeyExpirationPeriodInDays: to.Ptr[int32](20),
	// 			},
	// 			MinimumTLSVersion: to.Ptr(armstorage.MinimumTLSVersionTLS12),
	// 			PrimaryEndpoints: &armstorage.Endpoints{
	// 				Blob: to.Ptr("https://sto4445.z24.blob.storage.azure.net/"),
	// 				Dfs: to.Ptr("https://sto4445.z24.dfs.storage.azure.net/"),
	// 				File: to.Ptr("https://sto4445.z24.file.storage.azure.net/"),
	// 				InternetEndpoints: &armstorage.AccountInternetEndpoints{
	// 					Blob: to.Ptr("https://sto4445-internetrouting.z24.blob.storage.azure.net/"),
	// 					Dfs: to.Ptr("https://sto4445-internetrouting.z24.dfs.storage.azure.net/"),
	// 					File: to.Ptr("https://sto4445-internetrouting.z24.file.storage.azure.net/"),
	// 					Web: to.Ptr("https://sto4445-internetrouting.z24.web.storage.azure.net/"),
	// 				},
	// 				MicrosoftEndpoints: &armstorage.AccountMicrosoftEndpoints{
	// 					Blob: to.Ptr("https://sto4445-microsoftrouting.z24.blob.storage.azure.net/"),
	// 					Dfs: to.Ptr("https://sto4445-microsoftrouting.z24.dfs.storage.azure.net/"),
	// 					File: to.Ptr("https://sto4445-microsoftrouting.z24.file.storage.azure.net/"),
	// 					Queue: to.Ptr("https://sto4445-microsoftrouting.z24.queue.storage.azure.net/"),
	// 					Table: to.Ptr("https://sto4445-microsoftrouting.z24.table.storage.azure.net/"),
	// 					Web: to.Ptr("https://sto4445-microsoftrouting.z24.web.storage.azure.net/"),
	// 				},
	// 				Queue: to.Ptr("https://sto4445.z24.queue.storage.azure.net/"),
	// 				Table: to.Ptr("https://sto4445.z24.table.storage.azure.net/"),
	// 				Web: to.Ptr("https://sto4445.z24.web.storage.azure.net/"),
	// 			},
	// 			PrimaryLocation: to.Ptr("eastus2euap"),
	// 			ProvisioningState: to.Ptr(armstorage.ProvisioningStateSucceeded),
	// 			RoutingPreference: &armstorage.RoutingPreference{
	// 				PublishInternetEndpoints: to.Ptr(true),
	// 				PublishMicrosoftEndpoints: to.Ptr(true),
	// 				RoutingChoice: to.Ptr(armstorage.RoutingChoiceMicrosoftRouting),
	// 			},
	// 			SasPolicy: &armstorage.SasPolicy{
	// 				ExpirationAction: to.Ptr(armstorage.ExpirationActionLog),
	// 				SasExpirationPeriod: to.Ptr("1.15:59:59"),
	// 			},
	// 			SecondaryLocation: to.Ptr("centraluseuap"),
	// 			StatusOfPrimary: to.Ptr(armstorage.AccountStatusAvailable),
	// 			StatusOfSecondary: to.Ptr(armstorage.AccountStatusAvailable),
	// 			EnableHTTPSTrafficOnly: to.Ptr(true),
	// 		},
	// 		SKU: &armstorage.SKU{
	// 			Name: to.Ptr(armstorage.SKUNameStandardGRS),
	// 			Tier: to.Ptr(armstorage.SKUTierStandard),
	// 		},
	// 		Tags: map[string]*string{
	// 			"key1": to.Ptr("value1"),
	// 			"key2": to.Ptr("value2"),
	// 		},
	// 	},
	// }
}
