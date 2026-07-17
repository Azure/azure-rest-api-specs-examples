package armcosmos_test

import (
	"context"
	"log"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v4"
)

// Generated from example definition: 2026-03-15/CosmosDBRestoreDatabaseAccountCreateUpdate.json
func ExampleDatabaseAccountsClient_BeginCreateOrUpdate_cosmosDbRestoreDatabaseAccountCreateUpdateJson() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDatabaseAccountsClient().BeginCreateOrUpdate(ctx, "rg1", "ddb1", armcosmos.DatabaseAccountCreateUpdateParameters{
		Location: to.Ptr("westus"),
		Tags:     map[string]*string{},
		Kind:     to.Ptr(armcosmos.DatabaseAccountKindGlobalDocumentDB),
		Properties: &armcosmos.DatabaseAccountCreateUpdateProperties{
			DatabaseAccountOfferType: to.Ptr("Standard"),
			Locations: []*armcosmos.Location{
				{
					FailoverPriority: to.Ptr[int32](0),
					LocationName:     to.Ptr("southcentralus"),
					IsZoneRedundant:  to.Ptr(false),
				},
			},
			CreateMode: to.Ptr(armcosmos.CreateModeRestore),
			RestoreParameters: &armcosmos.RestoreParameters{
				RestoreMode:            to.Ptr(armcosmos.RestoreModePointInTime),
				RestoreSource:          to.Ptr("/subscriptions/subid/providers/Microsoft.DocumentDB/locations/westus/restorableDatabaseAccounts/1a97b4bb-f6a0-430e-ade1-638d781830cc"),
				RestoreTimestampInUTC:  to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-03-11T22:05:09Z"); return t }()),
				RestoreWithTTLDisabled: to.Ptr(false),
				DatabasesToRestore: []*armcosmos.DatabaseRestoreResource{
					{
						DatabaseName: to.Ptr("db1"),
						CollectionNames: []*string{
							to.Ptr("collection1"),
							to.Ptr("collection2"),
						},
					},
					{
						DatabaseName: to.Ptr("db2"),
						CollectionNames: []*string{
							to.Ptr("collection3"),
							to.Ptr("collection4"),
						},
					},
				},
				SourceBackupLocation: to.Ptr("westus"),
			},
			BackupPolicy: &armcosmos.ContinuousModeBackupPolicy{
				Type: to.Ptr(armcosmos.BackupPolicyTypeContinuous),
				ContinuousModeProperties: &armcosmos.ContinuousModeProperties{
					Tier: to.Ptr(armcosmos.ContinuousTierContinuous30Days),
				},
			},
			ConsistencyPolicy: &armcosmos.ConsistencyPolicy{
				DefaultConsistencyLevel: to.Ptr(armcosmos.DefaultConsistencyLevelBoundedStaleness),
				MaxIntervalInSeconds:    to.Ptr[int32](10),
				MaxStalenessPrefix:      to.Ptr[int64](200),
			},
			KeyVaultKeyURI: to.Ptr("https://myKeyVault.vault.azure.net"),
			EnableFreeTier: to.Ptr(false),
			APIProperties: &armcosmos.APIProperties{
				ServerVersion: to.Ptr(armcosmos.ServerVersionThree2),
			},
			EnableAnalyticalStorage:                    to.Ptr(true),
			EnforceHierarchicalPartitionKeyIDLastLevel: to.Ptr(false),
			MinimalTLSVersion:                          to.Ptr(armcosmos.MinimalTLSVersionTLS),
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
	// res = armcosmos.DatabaseAccountsClientCreateOrUpdateResponse{
	// 	DatabaseAccountGetResults: armcosmos.DatabaseAccountGetResults{
	// 		ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.DocumentDB/databaseAccounts/ddb1"),
	// 		Name: to.Ptr("ddb1"),
	// 		Location: to.Ptr("West US"),
	// 		Type: to.Ptr("Microsoft.DocumentDB/databaseAccounts"),
	// 		Kind: to.Ptr(armcosmos.DatabaseAccountKindGlobalDocumentDB),
	// 		Tags: map[string]*string{
	// 		},
	// 		Properties: &armcosmos.DatabaseAccountGetProperties{
	// 			ProvisioningState: to.Ptr("Initializing"),
	// 			IPRules: []*armcosmos.IPAddressOrRange{
	// 			},
	// 			DatabaseAccountOfferType: to.Ptr("Standard"),
	// 			DisableKeyBasedMetadataWriteAccess: to.Ptr(false),
	// 			InstanceID: to.Ptr("d9b26648-2f53-4541-b3d8-3044f4f9810d"),
	// 			CreateMode: to.Ptr(armcosmos.CreateModeRestore),
	// 			ConsistencyPolicy: &armcosmos.ConsistencyPolicy{
	// 				DefaultConsistencyLevel: to.Ptr(armcosmos.DefaultConsistencyLevelSession),
	// 				MaxIntervalInSeconds: to.Ptr[int32](5),
	// 				MaxStalenessPrefix: to.Ptr[int64](100),
	// 			},
	// 			WriteLocations: []*armcosmos.Location{
	// 				{
	// 					ID: to.Ptr("ddb1-southcentralus"),
	// 					LocationName: to.Ptr("South Central US"),
	// 					ProvisioningState: to.Ptr("Initializing"),
	// 					FailoverPriority: to.Ptr[int32](0),
	// 					IsZoneRedundant: to.Ptr(false),
	// 				},
	// 			},
	// 			ReadLocations: []*armcosmos.Location{
	// 				{
	// 					ID: to.Ptr("ddb1-southcentralus"),
	// 					LocationName: to.Ptr("South Central US"),
	// 					ProvisioningState: to.Ptr("Initializing"),
	// 					FailoverPriority: to.Ptr[int32](0),
	// 					IsZoneRedundant: to.Ptr(false),
	// 				},
	// 			},
	// 			Locations: []*armcosmos.Location{
	// 				{
	// 					ID: to.Ptr("ddb1-southcentralus"),
	// 					LocationName: to.Ptr("South Central US"),
	// 					ProvisioningState: to.Ptr("Initializing"),
	// 					FailoverPriority: to.Ptr[int32](0),
	// 					IsZoneRedundant: to.Ptr(false),
	// 				},
	// 			},
	// 			FailoverPolicies: []*armcosmos.FailoverPolicy{
	// 				{
	// 					ID: to.Ptr("ddb1-southcentralus"),
	// 					LocationName: to.Ptr("South Central US"),
	// 					FailoverPriority: to.Ptr[int32](0),
	// 				},
	// 			},
	// 			BackupPolicy: &armcosmos.ContinuousModeBackupPolicy{
	// 				Type: to.Ptr(armcosmos.BackupPolicyTypeContinuous),
	// 				ContinuousModeProperties: &armcosmos.ContinuousModeProperties{
	// 					Tier: to.Ptr(armcosmos.ContinuousTierContinuous30Days),
	// 				},
	// 			},
	// 			EnableFreeTier: to.Ptr(false),
	// 			APIProperties: &armcosmos.APIProperties{
	// 			},
	// 			EnableAnalyticalStorage: to.Ptr(false),
	// 			KeysMetadata: &armcosmos.DatabaseAccountKeysMetadata{
	// 				PrimaryMasterKey: &armcosmos.AccountKeyMetadata{
	// 					GenerationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-02-25T20:30:11Z"); return t}()),
	// 				},
	// 				SecondaryMasterKey: &armcosmos.AccountKeyMetadata{
	// 					GenerationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-02-25T20:30:11Z"); return t}()),
	// 				},
	// 				PrimaryReadonlyMasterKey: &armcosmos.AccountKeyMetadata{
	// 					GenerationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-02-25T20:30:11Z"); return t}()),
	// 				},
	// 				SecondaryReadonlyMasterKey: &armcosmos.AccountKeyMetadata{
	// 					GenerationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-02-25T20:30:11Z"); return t}()),
	// 				},
	// 			},
	// 			EnforceHierarchicalPartitionKeyIDLastLevel: to.Ptr(false),
	// 			MinimalTLSVersion: to.Ptr(armcosmos.MinimalTLSVersionTLS),
	// 		},
	// 		SystemData: &armcosmos.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-03-12T22:05:09Z"); return t}()),
	// 		},
	// 	},
	// }
}
