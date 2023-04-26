package armcosmos_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b065afacb99e7ec65787383550ee233e0e02a6e6/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2023-03-15/examples/CosmosDBRestoreDatabaseAccountCreateUpdate.json
func ExampleDatabaseAccountsClient_BeginCreateOrUpdate_cosmosDbRestoreDatabaseAccountCreateUpdateJson() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDatabaseAccountsClient().BeginCreateOrUpdate(ctx, "rg1", "ddb1", armcosmos.DatabaseAccountCreateUpdateParameters{
		Location: to.Ptr("westus"),
		Tags:     map[string]*string{},
		Kind:     to.Ptr(armcosmos.DatabaseAccountKindGlobalDocumentDB),
		Properties: &armcosmos.DatabaseAccountCreateUpdateProperties{
			APIProperties: &armcosmos.APIProperties{
				ServerVersion: to.Ptr(armcosmos.ServerVersionThree2),
			},
			BackupPolicy: &armcosmos.ContinuousModeBackupPolicy{
				Type: to.Ptr(armcosmos.BackupPolicyTypeContinuous),
			},
			ConsistencyPolicy: &armcosmos.ConsistencyPolicy{
				DefaultConsistencyLevel: to.Ptr(armcosmos.DefaultConsistencyLevelBoundedStaleness),
				MaxIntervalInSeconds:    to.Ptr[int32](10),
				MaxStalenessPrefix:      to.Ptr[int64](200),
			},
			CreateMode:               to.Ptr(armcosmos.CreateModeRestore),
			DatabaseAccountOfferType: to.Ptr("Standard"),
			EnableAnalyticalStorage:  to.Ptr(true),
			EnableFreeTier:           to.Ptr(false),
			KeyVaultKeyURI:           to.Ptr("https://myKeyVault.vault.azure.net"),
			Locations: []*armcosmos.Location{
				{
					FailoverPriority: to.Ptr[int32](0),
					IsZoneRedundant:  to.Ptr(false),
					LocationName:     to.Ptr("southcentralus"),
				}},
			MinimalTLSVersion: to.Ptr(armcosmos.MinimalTLSVersionTLS),
			RestoreParameters: &armcosmos.RestoreParameters{
				DatabasesToRestore: []*armcosmos.DatabaseRestoreResource{
					{
						CollectionNames: []*string{
							to.Ptr("collection1"),
							to.Ptr("collection2")},
						DatabaseName: to.Ptr("db1"),
					},
					{
						CollectionNames: []*string{
							to.Ptr("collection3"),
							to.Ptr("collection4")},
						DatabaseName: to.Ptr("db2"),
					}},
				RestoreMode:           to.Ptr(armcosmos.RestoreModePointInTime),
				RestoreSource:         to.Ptr("/subscriptions/subid/providers/Microsoft.DocumentDB/locations/westus/restorableDatabaseAccounts/1a97b4bb-f6a0-430e-ade1-638d781830cc"),
				RestoreTimestampInUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-03-11T22:05:09Z"); return t }()),
			},
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
	// res.DatabaseAccountGetResults = armcosmos.DatabaseAccountGetResults{
	// 	Name: to.Ptr("ddb1"),
	// 	Type: to.Ptr("Microsoft.DocumentDB/databaseAccounts"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.DocumentDB/databaseAccounts/ddb1"),
	// 	Location: to.Ptr("West US"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Kind: to.Ptr(armcosmos.DatabaseAccountKindGlobalDocumentDB),
	// 	Properties: &armcosmos.DatabaseAccountGetProperties{
	// 		APIProperties: &armcosmos.APIProperties{
	// 		},
	// 		ConsistencyPolicy: &armcosmos.ConsistencyPolicy{
	// 			DefaultConsistencyLevel: to.Ptr(armcosmos.DefaultConsistencyLevelSession),
	// 			MaxIntervalInSeconds: to.Ptr[int32](5),
	// 			MaxStalenessPrefix: to.Ptr[int64](100),
	// 		},
	// 		CreateMode: to.Ptr(armcosmos.CreateModeRestore),
	// 		DatabaseAccountOfferType: to.Ptr("Standard"),
	// 		DisableKeyBasedMetadataWriteAccess: to.Ptr(false),
	// 		EnableAnalyticalStorage: to.Ptr(false),
	// 		EnableFreeTier: to.Ptr(false),
	// 		FailoverPolicies: []*armcosmos.FailoverPolicy{
	// 			{
	// 				FailoverPriority: to.Ptr[int32](0),
	// 				ID: to.Ptr("ddb1-southcentralus"),
	// 				LocationName: to.Ptr("South Central US"),
	// 		}},
	// 		InstanceID: to.Ptr("d9b26648-2f53-4541-b3d8-3044f4f9810d"),
	// 		IPRules: []*armcosmos.IPAddressOrRange{
	// 		},
	// 		KeysMetadata: &armcosmos.DatabaseAccountKeysMetadata{
	// 			PrimaryMasterKey: &armcosmos.AccountKeyMetadata{
	// 				GenerationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-02-25T20:30:11Z"); return t}()),
	// 			},
	// 			PrimaryReadonlyMasterKey: &armcosmos.AccountKeyMetadata{
	// 				GenerationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-02-25T20:30:11Z"); return t}()),
	// 			},
	// 			SecondaryMasterKey: &armcosmos.AccountKeyMetadata{
	// 				GenerationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-02-25T20:30:11Z"); return t}()),
	// 			},
	// 			SecondaryReadonlyMasterKey: &armcosmos.AccountKeyMetadata{
	// 				GenerationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-02-25T20:30:11Z"); return t}()),
	// 			},
	// 		},
	// 		Locations: []*armcosmos.Location{
	// 			{
	// 				FailoverPriority: to.Ptr[int32](0),
	// 				ID: to.Ptr("ddb1-southcentralus"),
	// 				IsZoneRedundant: to.Ptr(false),
	// 				LocationName: to.Ptr("South Central US"),
	// 				ProvisioningState: to.Ptr("Succeeded"),
	// 		}},
	// 		MinimalTLSVersion: to.Ptr(armcosmos.MinimalTLSVersionTLS),
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		ReadLocations: []*armcosmos.Location{
	// 			{
	// 				FailoverPriority: to.Ptr[int32](0),
	// 				ID: to.Ptr("ddb1-southcentralus"),
	// 				IsZoneRedundant: to.Ptr(false),
	// 				LocationName: to.Ptr("South Central US"),
	// 				ProvisioningState: to.Ptr("Succeeded"),
	// 		}},
	// 		WriteLocations: []*armcosmos.Location{
	// 			{
	// 				FailoverPriority: to.Ptr[int32](0),
	// 				ID: to.Ptr("ddb1-southcentralus"),
	// 				IsZoneRedundant: to.Ptr(false),
	// 				LocationName: to.Ptr("South Central US"),
	// 				ProvisioningState: to.Ptr("Succeeded"),
	// 		}},
	// 	},
	// 	SystemData: &armcosmos.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-03-12T22:05:09Z"); return t}()),
	// 	},
	// }
}
