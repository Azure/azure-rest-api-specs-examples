package armpostgresqlflexibleservers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresqlflexibleservers/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b1f4d539964453ce8008e4b069e59885e12ba441/specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2025-01-01-preview/examples/ReplicasListByServer.json
func ExampleReplicasClient_NewListByServerPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpostgresqlflexibleservers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewReplicasClient().NewListByServerPager("testrg", "sourcepgservername", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.ServerListResult = armpostgresqlflexibleservers.ServerListResult{
		// 	Value: []*armpostgresqlflexibleservers.Server{
		// 		{
		// 			Name: to.Ptr("pgtestsvc5rep"),
		// 			Type: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforPostgreSQL/flexibleServers/pgtestsvc5rep"),
		// 			Location: to.Ptr("westus"),
		// 			Tags: map[string]*string{
		// 				"ElasticServer": to.Ptr("1"),
		// 			},
		// 			Properties: &armpostgresqlflexibleservers.ServerProperties{
		// 				AdministratorLogin: to.Ptr("login"),
		// 				AuthConfig: &armpostgresqlflexibleservers.AuthConfig{
		// 					ActiveDirectoryAuth: to.Ptr(armpostgresqlflexibleservers.ActiveDirectoryAuthEnumDisabled),
		// 					PasswordAuth: to.Ptr(armpostgresqlflexibleservers.PasswordAuthEnumEnabled),
		// 				},
		// 				AvailabilityZone: to.Ptr("2"),
		// 				Backup: &armpostgresqlflexibleservers.Backup{
		// 					BackupRetentionDays: to.Ptr[int32](7),
		// 					EarliestRestoreDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-03-03T00:28:17.727Z"); return t}()),
		// 					GeoRedundantBackup: to.Ptr(armpostgresqlflexibleservers.GeoRedundantBackupEnumDisabled),
		// 				},
		// 				DataEncryption: &armpostgresqlflexibleservers.DataEncryption{
		// 					Type: to.Ptr(armpostgresqlflexibleservers.ArmServerKeyTypeSystemManaged),
		// 				},
		// 				FullyQualifiedDomainName: to.Ptr("pgtestsvc5rep.postgres.database.azure.com"),
		// 				HighAvailability: &armpostgresqlflexibleservers.HighAvailability{
		// 					Mode: to.Ptr(armpostgresqlflexibleservers.HighAvailabilityModeDisabled),
		// 					State: to.Ptr(armpostgresqlflexibleservers.ServerHAStateNotEnabled),
		// 				},
		// 				MaintenanceWindow: &armpostgresqlflexibleservers.MaintenanceWindow{
		// 					CustomWindow: to.Ptr("Disabled"),
		// 					DayOfWeek: to.Ptr[int32](0),
		// 					StartHour: to.Ptr[int32](0),
		// 					StartMinute: to.Ptr[int32](0),
		// 				},
		// 				MinorVersion: to.Ptr("6"),
		// 				Network: &armpostgresqlflexibleservers.Network{
		// 					PublicNetworkAccess: to.Ptr(armpostgresqlflexibleservers.ServerPublicNetworkAccessStateEnabled),
		// 				},
		// 				Replica: &armpostgresqlflexibleservers.Replica{
		// 					Capacity: to.Ptr[int32](0),
		// 					ReplicationState: to.Ptr(armpostgresqlflexibleservers.ReplicationStateActive),
		// 					Role: to.Ptr(armpostgresqlflexibleservers.ReplicationRoleAsyncReplica),
		// 				},
		// 				ReplicaCapacity: to.Ptr[int32](0),
		// 				ReplicationRole: to.Ptr(armpostgresqlflexibleservers.ReplicationRoleAsyncReplica),
		// 				State: to.Ptr(armpostgresqlflexibleservers.ServerStateReady),
		// 				Storage: &armpostgresqlflexibleservers.Storage{
		// 					AutoGrow: to.Ptr(armpostgresqlflexibleservers.StorageAutoGrowDisabled),
		// 					Iops: to.Ptr[int32](2300),
		// 					StorageSizeGB: to.Ptr[int32](512),
		// 					Tier: to.Ptr(armpostgresqlflexibleservers.AzureManagedDiskPerformanceTiersP20),
		// 				},
		// 				Version: to.Ptr(armpostgresqlflexibleservers.ServerVersionSixteen),
		// 			},
		// 			SKU: &armpostgresqlflexibleservers.SKU{
		// 				Name: to.Ptr("Standard_D4ds_v5"),
		// 				Tier: to.Ptr(armpostgresqlflexibleservers.SKUTierGeneralPurpose),
		// 			},
		// 	}},
		// }
	}
}
