package armpostgresqlflexibleservers_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresqlflexibleservers/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b1f4d539964453ce8008e4b069e59885e12ba441/specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2025-01-01-preview/examples/ServerCreateReplica.json
func ExampleServersClient_BeginCreate_serverCreateReplica() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpostgresqlflexibleservers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewServersClient().BeginCreate(ctx, "testrg", "pgtestsvc5rep", armpostgresqlflexibleservers.Server{
		Location: to.Ptr("westus"),
		Identity: &armpostgresqlflexibleservers.UserAssignedIdentity{
			Type: to.Ptr(armpostgresqlflexibleservers.IdentityTypeUserAssigned),
			UserAssignedIdentities: map[string]*armpostgresqlflexibleservers.UserIdentity{
				"/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testresourcegroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/test-usermanagedidentity": {},
			},
		},
		Properties: &armpostgresqlflexibleservers.ServerProperties{
			CreateMode: to.Ptr(armpostgresqlflexibleservers.CreateModeReplica),
			DataEncryption: &armpostgresqlflexibleservers.DataEncryption{
				Type:                            to.Ptr(armpostgresqlflexibleservers.ArmServerKeyTypeAzureKeyVault),
				GeoBackupKeyURI:                 to.Ptr(""),
				GeoBackupUserAssignedIdentityID: to.Ptr(""),
				PrimaryKeyURI:                   to.Ptr("https://test-kv.vault.azure.net/keys/test-key1/xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"),
				PrimaryUserAssignedIdentityID:   to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testresourcegroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/test-usermanagedidentity"),
			},
			PointInTimeUTC:         to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-27T00:04:59.407Z"); return t }()),
			SourceServerResourceID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforPostgreSQL/flexibleServers/sourcepgservername"),
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
	// res.Server = armpostgresqlflexibleservers.Server{
	// 	Name: to.Ptr("pgtestsvc5rep"),
	// 	Type: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers"),
	// 	ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforPostgreSQL/flexibleServers/pgtestsvc5rep"),
	// 	Location: to.Ptr("westus"),
	// 	Tags: map[string]*string{
	// 		"ElasticServer": to.Ptr("1"),
	// 	},
	// 	Identity: &armpostgresqlflexibleservers.UserAssignedIdentity{
	// 		Type: to.Ptr(armpostgresqlflexibleservers.IdentityTypeUserAssigned),
	// 		TenantID: to.Ptr("cccccccc-cccc-cccc-cccc-cccccccccccc"),
	// 		UserAssignedIdentities: map[string]*armpostgresqlflexibleservers.UserIdentity{
	// 			"/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testresourcegroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/test-usermanagedidentity": &armpostgresqlflexibleservers.UserIdentity{
	// 				ClientID: to.Ptr("cccccccc-cccc-cccc-cccc-cccccccccccc"),
	// 				PrincipalID: to.Ptr("pppppppp-pppp-pppp-pppp-pppppppppppp"),
	// 			},
	// 		},
	// 	},
	// 	Properties: &armpostgresqlflexibleservers.ServerProperties{
	// 		AdministratorLogin: to.Ptr("login"),
	// 		AuthConfig: &armpostgresqlflexibleservers.AuthConfig{
	// 			ActiveDirectoryAuth: to.Ptr(armpostgresqlflexibleservers.ActiveDirectoryAuthEnumDisabled),
	// 			PasswordAuth: to.Ptr(armpostgresqlflexibleservers.PasswordAuthEnumEnabled),
	// 		},
	// 		AvailabilityZone: to.Ptr("2"),
	// 		Backup: &armpostgresqlflexibleservers.Backup{
	// 			BackupRetentionDays: to.Ptr[int32](7),
	// 			EarliestRestoreDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-03-03T00:28:17.727Z"); return t}()),
	// 			GeoRedundantBackup: to.Ptr(armpostgresqlflexibleservers.GeoRedundantBackupEnumDisabled),
	// 		},
	// 		DataEncryption: &armpostgresqlflexibleservers.DataEncryption{
	// 			Type: to.Ptr(armpostgresqlflexibleservers.ArmServerKeyTypeAzureKeyVault),
	// 			PrimaryEncryptionKeyStatus: to.Ptr(armpostgresqlflexibleservers.KeyStatusEnumValid),
	// 			PrimaryKeyURI: to.Ptr("https://test-kv.vault.azure.net/keys/test-key1/xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"),
	// 			PrimaryUserAssignedIdentityID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testresourcegroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/test-usermanagedidentity"),
	// 		},
	// 		FullyQualifiedDomainName: to.Ptr("pgtestsvc5rep.postgres.database.azure.com"),
	// 		HighAvailability: &armpostgresqlflexibleservers.HighAvailability{
	// 			Mode: to.Ptr(armpostgresqlflexibleservers.HighAvailabilityModeDisabled),
	// 			State: to.Ptr(armpostgresqlflexibleservers.ServerHAStateNotEnabled),
	// 		},
	// 		MaintenanceWindow: &armpostgresqlflexibleservers.MaintenanceWindow{
	// 			CustomWindow: to.Ptr("Disabled"),
	// 			DayOfWeek: to.Ptr[int32](0),
	// 			StartHour: to.Ptr[int32](0),
	// 			StartMinute: to.Ptr[int32](0),
	// 		},
	// 		MinorVersion: to.Ptr("6"),
	// 		Network: &armpostgresqlflexibleservers.Network{
	// 			PublicNetworkAccess: to.Ptr(armpostgresqlflexibleservers.ServerPublicNetworkAccessStateEnabled),
	// 		},
	// 		Replica: &armpostgresqlflexibleservers.Replica{
	// 			Capacity: to.Ptr[int32](0),
	// 			ReplicationState: to.Ptr(armpostgresqlflexibleservers.ReplicationStateActive),
	// 			Role: to.Ptr(armpostgresqlflexibleservers.ReplicationRoleAsyncReplica),
	// 		},
	// 		ReplicaCapacity: to.Ptr[int32](0),
	// 		ReplicationRole: to.Ptr(armpostgresqlflexibleservers.ReplicationRoleAsyncReplica),
	// 		SourceServerResourceID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforPostgreSQL/flexibleServers/sourcepgservername"),
	// 		State: to.Ptr(armpostgresqlflexibleservers.ServerStateReady),
	// 		Storage: &armpostgresqlflexibleservers.Storage{
	// 			AutoGrow: to.Ptr(armpostgresqlflexibleservers.StorageAutoGrowDisabled),
	// 			Iops: to.Ptr[int32](2300),
	// 			StorageSizeGB: to.Ptr[int32](512),
	// 			Tier: to.Ptr(armpostgresqlflexibleservers.AzureManagedDiskPerformanceTiersP20),
	// 		},
	// 		Version: to.Ptr(armpostgresqlflexibleservers.ServerVersionSixteen),
	// 	},
	// 	SKU: &armpostgresqlflexibleservers.SKU{
	// 		Name: to.Ptr("Standard_D4ds_v5"),
	// 		Tier: to.Ptr(armpostgresqlflexibleservers.SKUTierGeneralPurpose),
	// 	},
	// }
}
