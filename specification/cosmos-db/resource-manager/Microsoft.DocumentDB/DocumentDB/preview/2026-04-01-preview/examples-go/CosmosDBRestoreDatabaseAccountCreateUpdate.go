package armcosmos_test

import (
	"context"
	"log"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v4"
)

// Generated from example definition: 2026-04-01-preview/CosmosDBRestoreDatabaseAccountCreateUpdate.json
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
		Kind:     to.Ptr(armcosmos.DatabaseAccountKindGlobalDocumentDB),
		Location: to.Ptr("westus"),
		Properties: &armcosmos.DatabaseAccountCreateUpdateProperties{
			APIProperties: &armcosmos.APIProperties{
				ServerVersion: to.Ptr(armcosmos.ServerVersionThree2),
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
			CreateMode:               to.Ptr(armcosmos.CreateModeRestore),
			DatabaseAccountOfferType: to.Ptr("Standard"),
			EnableAnalyticalStorage:  to.Ptr(true),
			EnableFreeTier:           to.Ptr(false),
			EnableMaterializedViews:  to.Ptr(false),
			KeyVaultKeyURI:           to.Ptr("https://myKeyVault.vault.azure.net"),
			Locations: []*armcosmos.Location{
				{
					FailoverPriority: to.Ptr[int32](0),
					IsZoneRedundant:  to.Ptr(false),
					LocationName:     to.Ptr("southcentralus"),
				},
			},
			EnforceHierarchicalPartitionKeyIDLastLevel: to.Ptr(false),
			MinimalTLSVersion:                          to.Ptr(armcosmos.MinimalTLSVersionTLS),
			RestoreParameters: &armcosmos.RestoreParameters{
				DatabasesToRestore: []*armcosmos.DatabaseRestoreResource{
					{
						CollectionNames: []*string{
							to.Ptr("collection1"),
							to.Ptr("collection2"),
						},
						DatabaseName: to.Ptr("db1"),
					},
					{
						CollectionNames: []*string{
							to.Ptr("collection3"),
							to.Ptr("collection4"),
						},
						DatabaseName: to.Ptr("db2"),
					},
				},
				RestoreMode:            to.Ptr(armcosmos.RestoreModePointInTime),
				RestoreSource:          to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/providers/Microsoft.DocumentDB/locations/westus/restorableDatabaseAccounts/1a97b4bb-f6a0-430e-ade1-638d781830cc"),
				RestoreTimestampInUTC:  to.Ptr(time.Date(2021, time.March, 11, 22, 5, 9, 0, time.UTC)),
				RestoreWithTTLDisabled: to.Ptr(false),
				SourceBackupLocation:   to.Ptr("westus"),
			},
		},
		Tags: map[string]*string{},
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
	// 		Name: to.Ptr("ddb1"),
	// 		Type: to.Ptr("Microsoft.DocumentDB/databaseAccounts"),
	// 		ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/rg1/providers/Microsoft.DocumentDB/databaseAccounts/ddb1"),
	// 		Kind: to.Ptr(armcosmos.DatabaseAccountKindGlobalDocumentDB),
	// 		Location: to.Ptr("West US"),
	// 		Properties: &armcosmos.DatabaseAccountGetProperties{
	// 			APIProperties: &armcosmos.APIProperties{
	// 			},
	// 			BackupPolicy: &armcosmos.ContinuousModeBackupPolicy{
	// 				Type: to.Ptr(armcosmos.BackupPolicyTypeContinuous),
	// 				ContinuousModeProperties: &armcosmos.ContinuousModeProperties{
	// 					Tier: to.Ptr(armcosmos.ContinuousTierContinuous30Days),
	// 				},
	// 			},
	// 			CapacityMode: to.Ptr(armcosmos.CapacityModeProvisioned),
	// 			ConsistencyPolicy: &armcosmos.ConsistencyPolicy{
	// 				DefaultConsistencyLevel: to.Ptr(armcosmos.DefaultConsistencyLevelSession),
	// 				MaxIntervalInSeconds: to.Ptr[int32](5),
	// 				MaxStalenessPrefix: to.Ptr[int64](100),
	// 			},
	// 			CreateMode: to.Ptr(armcosmos.CreateModeRestore),
	// 			DatabaseAccountOfferType: to.Ptr("Standard"),
	// 			DisableKeyBasedMetadataWriteAccess: to.Ptr(false),
	// 			EnableAnalyticalStorage: to.Ptr(false),
	// 			EnableFreeTier: to.Ptr(false),
	// 			EnableMaterializedViews: to.Ptr(false),
	// 			FailoverPolicies: []*armcosmos.FailoverPolicy{
	// 				{
	// 					FailoverPriority: to.Ptr[int32](0),
	// 					ID: to.Ptr("ddb1-southcentralus"),
	// 					LocationName: to.Ptr("South Central US"),
	// 				},
	// 			},
	// 			InstanceID: to.Ptr("d9b26648-2f53-4541-b3d8-3044f4f9810d"),
	// 			IPRules: []*armcosmos.IPAddressOrRange{
	// 			},
	// 			KeysMetadata: &armcosmos.DatabaseAccountKeysMetadata{
	// 				PrimaryMasterKey: &armcosmos.AccountKeyMetadata{
	// 					GenerationTime: to.Ptr(time.Date(2022, time.February, 25, 20, 30, 11, 0, time.UTC)),
	// 					ApproximateLastUsageTime: to.Ptr(time.Date(2024, time.October, 9, 11, 45, 33, 0, time.UTC)),
	// 				},
	// 				PrimaryReadonlyMasterKey: &armcosmos.AccountKeyMetadata{
	// 					GenerationTime: to.Ptr(time.Date(2022, time.February, 25, 20, 30, 11, 0, time.UTC)),
	// 					ApproximateLastUsageTime: to.Ptr(time.Date(2024, time.October, 9, 11, 45, 33, 0, time.UTC)),
	// 				},
	// 				SecondaryMasterKey: &armcosmos.AccountKeyMetadata{
	// 					GenerationTime: to.Ptr(time.Date(2022, time.February, 25, 20, 30, 11, 0, time.UTC)),
	// 					ApproximateLastUsageTime: to.Ptr(time.Date(2024, time.October, 9, 11, 45, 33, 0, time.UTC)),
	// 				},
	// 				SecondaryReadonlyMasterKey: &armcosmos.AccountKeyMetadata{
	// 					GenerationTime: to.Ptr(time.Date(2022, time.February, 25, 20, 30, 11, 0, time.UTC)),
	// 					ApproximateLastUsageTime: to.Ptr(time.Date(2024, time.October, 9, 11, 45, 33, 0, time.UTC)),
	// 				},
	// 			},
	// 			Locations: []*armcosmos.Location{
	// 				{
	// 					FailoverPriority: to.Ptr[int32](0),
	// 					ID: to.Ptr("ddb1-southcentralus"),
	// 					IsZoneRedundant: to.Ptr(false),
	// 					LocationName: to.Ptr("South Central US"),
	// 					ProvisioningState: to.Ptr("Initializing"),
	// 				},
	// 			},
	// 			EnforceHierarchicalPartitionKeyIDLastLevel: to.Ptr(false),
	// 			MinimalTLSVersion: to.Ptr(armcosmos.MinimalTLSVersionTLS),
	// 			NetworkACLBypass: to.Ptr(armcosmos.NetworkACLBypassNone),
	// 			NetworkACLBypassResourceIDs: []*string{
	// 			},
	// 			ProvisioningState: to.Ptr("Initializing"),
	// 			ReadLocations: []*armcosmos.Location{
	// 				{
	// 					FailoverPriority: to.Ptr[int32](0),
	// 					ID: to.Ptr("ddb1-southcentralus"),
	// 					IsZoneRedundant: to.Ptr(false),
	// 					LocationName: to.Ptr("South Central US"),
	// 					ProvisioningState: to.Ptr("Initializing"),
	// 				},
	// 			},
	// 			WriteLocations: []*armcosmos.Location{
	// 				{
	// 					FailoverPriority: to.Ptr[int32](0),
	// 					ID: to.Ptr("ddb1-southcentralus"),
	// 					IsZoneRedundant: to.Ptr(false),
	// 					LocationName: to.Ptr("South Central US"),
	// 					ProvisioningState: to.Ptr("Initializing"),
	// 				},
	// 			},
	// 		},
	// 		SystemData: &armcosmos.SystemData{
	// 			CreatedAt: to.Ptr(time.Date(2021, time.March, 12, 22, 5, 9, 0, time.UTC)),
	// 		},
	// 		Tags: map[string]*string{
	// 		},
	// 	},
	// }
}
