package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v4"
)

// Generated from example definition: 2026-04-01-preview/CosmosDBDatabaseAccountCreateMin.json
func ExampleDatabaseAccountsClient_BeginCreateOrUpdate_cosmosDbDatabaseAccountCreateMin() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDatabaseAccountsClient().BeginCreateOrUpdate(ctx, "rg1", "ddb1", armcosmos.DatabaseAccountCreateUpdateParameters{
		Location: to.Ptr("westus"),
		Properties: &armcosmos.DatabaseAccountCreateUpdateProperties{
			DatabaseAccountOfferType: to.Ptr("Standard"),
			Locations: []*armcosmos.Location{
				{
					FailoverPriority: to.Ptr[int32](0),
					LocationName:     to.Ptr("southcentralus"),
					IsZoneRedundant:  to.Ptr(false),
				},
			},
			CreateMode: to.Ptr(armcosmos.CreateModeDefault),
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
	// 			Cors: []*armcosmos.CorsPolicy{
	// 			},
	// 			EnableFreeTier: to.Ptr(false),
	// 			APIProperties: &armcosmos.APIProperties{
	// 			},
	// 			EnableAnalyticalStorage: to.Ptr(false),
	// 			NetworkACLBypass: to.Ptr(armcosmos.NetworkACLBypassNone),
	// 			NetworkACLBypassResourceIDs: []*string{
	// 			},
	// 			InstanceID: to.Ptr("d9b26648-2f53-4541-b3d8-3044f4f9810d"),
	// 			CreateMode: to.Ptr(armcosmos.CreateModeDefault),
	// 			BackupPolicy: &armcosmos.PeriodicModeBackupPolicy{
	// 				Type: to.Ptr(armcosmos.BackupPolicyTypePeriodic),
	// 				PeriodicModeProperties: &armcosmos.PeriodicModeProperties{
	// 					BackupIntervalInMinutes: to.Ptr[int32](240),
	// 					BackupRetentionIntervalInHours: to.Ptr[int32](720),
	// 					BackupStorageRedundancy: to.Ptr(armcosmos.BackupStorageRedundancyGeo),
	// 				},
	// 			},
	// 			CapacityMode: to.Ptr(armcosmos.CapacityModeProvisioned),
	// 			EnableMaterializedViews: to.Ptr(false),
	// 			MinimalTLSVersion: to.Ptr(armcosmos.MinimalTLSVersionTLS),
	// 			KeysMetadata: &armcosmos.DatabaseAccountKeysMetadata{
	// 				PrimaryMasterKey: &armcosmos.AccountKeyMetadata{
	// 					GenerationTime: to.Ptr(time.Date(2022, time.February, 25, 20, 30, 11, 0, time.UTC)),
	// 					ApproximateLastUsageTime: to.Ptr(time.Date(2024, time.October, 9, 11, 45, 33, 0, time.UTC)),
	// 				},
	// 				SecondaryMasterKey: &armcosmos.AccountKeyMetadata{
	// 					GenerationTime: to.Ptr(time.Date(2022, time.February, 25, 20, 30, 11, 0, time.UTC)),
	// 					ApproximateLastUsageTime: to.Ptr(time.Date(2024, time.October, 9, 11, 45, 33, 0, time.UTC)),
	// 				},
	// 				PrimaryReadonlyMasterKey: &armcosmos.AccountKeyMetadata{
	// 					GenerationTime: to.Ptr(time.Date(2022, time.February, 25, 20, 30, 11, 0, time.UTC)),
	// 					ApproximateLastUsageTime: to.Ptr(time.Date(2024, time.October, 9, 11, 45, 33, 0, time.UTC)),
	// 				},
	// 				SecondaryReadonlyMasterKey: &armcosmos.AccountKeyMetadata{
	// 					GenerationTime: to.Ptr(time.Date(2022, time.February, 25, 20, 30, 11, 0, time.UTC)),
	// 					ApproximateLastUsageTime: to.Ptr(time.Date(2024, time.October, 9, 11, 45, 33, 0, time.UTC)),
	// 				},
	// 			},
	// 			ThroughputPoolDedicatedRUs: to.Ptr[int64](0),
	// 			ThroughputPoolMaxConsumableRUs: to.Ptr[int64](0),
	// 		},
	// 		SystemData: &armcosmos.SystemData{
	// 			CreatedAt: to.Ptr(time.Date(2021, time.March, 12, 22, 5, 9, 0, time.UTC)),
	// 		},
	// 	},
	// }
}
