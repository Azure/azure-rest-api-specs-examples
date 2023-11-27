package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1b33e81bbdc28fcd6644a1315b8d7b1b6d030590/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2023-03-15-preview/examples/CosmosDBDatabaseAccountCreateMin.json
func ExampleDatabaseAccountsClient_BeginCreateOrUpdate_cosmosDbDatabaseAccountCreateMin() {
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
		Properties: &armcosmos.DatabaseAccountCreateUpdateProperties{
			CreateMode:               to.Ptr(armcosmos.CreateModeDefault),
			DatabaseAccountOfferType: to.Ptr("Standard"),
			Locations: []*armcosmos.Location{
				{
					FailoverPriority: to.Ptr[int32](0),
					IsZoneRedundant:  to.Ptr(false),
					LocationName:     to.Ptr("southcentralus"),
				}},
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
	// 		BackupPolicy: &armcosmos.PeriodicModeBackupPolicy{
	// 			Type: to.Ptr(armcosmos.BackupPolicyTypePeriodic),
	// 			PeriodicModeProperties: &armcosmos.PeriodicModeProperties{
	// 				BackupIntervalInMinutes: to.Ptr[int32](240),
	// 				BackupRetentionIntervalInHours: to.Ptr[int32](720),
	// 				BackupStorageRedundancy: to.Ptr(armcosmos.BackupStorageRedundancyGeo),
	// 			},
	// 		},
	// 		ConsistencyPolicy: &armcosmos.ConsistencyPolicy{
	// 			DefaultConsistencyLevel: to.Ptr(armcosmos.DefaultConsistencyLevelSession),
	// 			MaxIntervalInSeconds: to.Ptr[int32](5),
	// 			MaxStalenessPrefix: to.Ptr[int64](100),
	// 		},
	// 		Cors: []*armcosmos.CorsPolicy{
	// 		},
	// 		CreateMode: to.Ptr(armcosmos.CreateModeDefault),
	// 		DatabaseAccountOfferType: to.Ptr("Standard"),
	// 		DisableKeyBasedMetadataWriteAccess: to.Ptr(false),
	// 		EnableAnalyticalStorage: to.Ptr(false),
	// 		EnableFreeTier: to.Ptr(false),
	// 		EnableMaterializedViews: to.Ptr(false),
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
	// 				GenerationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-02-25T20:30:11.000Z"); return t}()),
	// 			},
	// 			PrimaryReadonlyMasterKey: &armcosmos.AccountKeyMetadata{
	// 				GenerationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-02-25T20:30:11.000Z"); return t}()),
	// 			},
	// 			SecondaryMasterKey: &armcosmos.AccountKeyMetadata{
	// 				GenerationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-02-25T20:30:11.000Z"); return t}()),
	// 			},
	// 			SecondaryReadonlyMasterKey: &armcosmos.AccountKeyMetadata{
	// 				GenerationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-02-25T20:30:11.000Z"); return t}()),
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
	// 		NetworkACLBypass: to.Ptr(armcosmos.NetworkACLBypassNone),
	// 		NetworkACLBypassResourceIDs: []*string{
	// 		},
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
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-03-12T22:05:09.000Z"); return t}()),
	// 	},
	// }
}
