package armpostgresql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresql"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c767823fdfd9d5e96bad245e3ea4d14d94a716bb/specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2017-12-01/examples/ReplicasListByServer.json
func ExampleReplicasClient_NewListByServerPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpostgresql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewReplicasClient().NewListByServerPager("TestGroup_WestCentralUS", "testserver-master", nil)
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
		// page.ServerListResult = armpostgresql.ServerListResult{
		// 	Value: []*armpostgresql.Server{
		// 		{
		// 			Name: to.Ptr("testserver-replica1"),
		// 			Type: to.Ptr("Microsoft.DBforPostgreSQL/servers"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestGroup_WestCentralUS/providers/Microsoft.DBforPostgreSQL/servers/testserver-replica1"),
		// 			Location: to.Ptr("westcentralus"),
		// 			Properties: &armpostgresql.ServerProperties{
		// 				AdministratorLogin: to.Ptr("postgres"),
		// 				EarliestRestoreDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-12T00:05:03.269Z"); return t}()),
		// 				FullyQualifiedDomainName: to.Ptr("testserver-replica1.postgres.database.azure.com"),
		// 				MasterServerID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestGroup_WestCentralUS/providers/Microsoft.DBforPostgreSQL/servers/testserver-master"),
		// 				ReplicaCapacity: to.Ptr[int32](0),
		// 				ReplicationRole: to.Ptr("Replica"),
		// 				SSLEnforcement: to.Ptr(armpostgresql.SSLEnforcementEnumDisabled),
		// 				StorageProfile: &armpostgresql.StorageProfile{
		// 					BackupRetentionDays: to.Ptr[int32](7),
		// 					GeoRedundantBackup: to.Ptr(armpostgresql.GeoRedundantBackupDisabled),
		// 					StorageMB: to.Ptr[int32](2048000),
		// 				},
		// 				UserVisibleState: to.Ptr(armpostgresql.ServerStateReady),
		// 				Version: to.Ptr(armpostgresql.ServerVersionNine6),
		// 			},
		// 			SKU: &armpostgresql.SKU{
		// 				Name: to.Ptr("GP_Gen4_16"),
		// 				Capacity: to.Ptr[int32](16),
		// 				Family: to.Ptr("Gen4"),
		// 				Tier: to.Ptr(armpostgresql.SKUTierGeneralPurpose),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("testserver-replica2"),
		// 			Type: to.Ptr("Microsoft.DBforPostgreSQL/servers"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestGroup_WestCentralUS/providers/Microsoft.DBforPostgreSQL/servers/testserver-replica2"),
		// 			Location: to.Ptr("westcentralus"),
		// 			Properties: &armpostgresql.ServerProperties{
		// 				AdministratorLogin: to.Ptr("postgres"),
		// 				EarliestRestoreDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-12T00:05:03.269Z"); return t}()),
		// 				FullyQualifiedDomainName: to.Ptr("testserver-replica2.postgres.database.azure.com"),
		// 				MasterServerID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestGroup_WestCentralUS/providers/Microsoft.DBforPostgreSQL/servers/testserver-master"),
		// 				ReplicaCapacity: to.Ptr[int32](0),
		// 				ReplicationRole: to.Ptr("Replica"),
		// 				SSLEnforcement: to.Ptr(armpostgresql.SSLEnforcementEnumDisabled),
		// 				StorageProfile: &armpostgresql.StorageProfile{
		// 					BackupRetentionDays: to.Ptr[int32](7),
		// 					GeoRedundantBackup: to.Ptr(armpostgresql.GeoRedundantBackupDisabled),
		// 					StorageMB: to.Ptr[int32](2048000),
		// 				},
		// 				UserVisibleState: to.Ptr(armpostgresql.ServerStateReady),
		// 				Version: to.Ptr(armpostgresql.ServerVersionNine6),
		// 			},
		// 			SKU: &armpostgresql.SKU{
		// 				Name: to.Ptr("GP_Gen4_16"),
		// 				Capacity: to.Ptr[int32](16),
		// 				Family: to.Ptr("Gen4"),
		// 				Tier: to.Ptr(armpostgresql.SKUTierGeneralPurpose),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("testserver-replica3"),
		// 			Type: to.Ptr("Microsoft.DBforPostgreSQL/servers"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestGroup_WestCentralUS/providers/Microsoft.DBforPostgreSQL/servers/testserver-replica3"),
		// 			Location: to.Ptr("westcentralus"),
		// 			Properties: &armpostgresql.ServerProperties{
		// 				AdministratorLogin: to.Ptr("postgres"),
		// 				EarliestRestoreDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-12T00:05:03.269Z"); return t}()),
		// 				FullyQualifiedDomainName: to.Ptr("testserver-replica3.postgres.database.azure.com"),
		// 				MasterServerID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestGroup_WestCentralUS/providers/Microsoft.DBforPostgreSQL/servers/testserver-master"),
		// 				ReplicaCapacity: to.Ptr[int32](0),
		// 				ReplicationRole: to.Ptr("Replica"),
		// 				SSLEnforcement: to.Ptr(armpostgresql.SSLEnforcementEnumDisabled),
		// 				StorageProfile: &armpostgresql.StorageProfile{
		// 					BackupRetentionDays: to.Ptr[int32](7),
		// 					GeoRedundantBackup: to.Ptr(armpostgresql.GeoRedundantBackupDisabled),
		// 					StorageMB: to.Ptr[int32](2048000),
		// 				},
		// 				UserVisibleState: to.Ptr(armpostgresql.ServerStateReady),
		// 				Version: to.Ptr(armpostgresql.ServerVersionNine6),
		// 			},
		// 			SKU: &armpostgresql.SKU{
		// 				Name: to.Ptr("GP_Gen4_16"),
		// 				Capacity: to.Ptr[int32](16),
		// 				Family: to.Ptr("Gen4"),
		// 				Tier: to.Ptr(armpostgresql.SKUTierGeneralPurpose),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("testserver-replica4"),
		// 			Type: to.Ptr("Microsoft.DBforPostgreSQL/servers"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestGroup_WestCentralUS/providers/Microsoft.DBforPostgreSQL/servers/testserver-replica4"),
		// 			Location: to.Ptr("westcentralus"),
		// 			Properties: &armpostgresql.ServerProperties{
		// 				AdministratorLogin: to.Ptr("postgres"),
		// 				EarliestRestoreDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-12T00:05:03.269Z"); return t}()),
		// 				FullyQualifiedDomainName: to.Ptr("testserver-replica4.postgres.database.azure.com"),
		// 				MasterServerID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestGroup_WestCentralUS/providers/Microsoft.DBforPostgreSQL/servers/testserver-master"),
		// 				ReplicaCapacity: to.Ptr[int32](0),
		// 				ReplicationRole: to.Ptr("Replica"),
		// 				SSLEnforcement: to.Ptr(armpostgresql.SSLEnforcementEnumDisabled),
		// 				StorageProfile: &armpostgresql.StorageProfile{
		// 					BackupRetentionDays: to.Ptr[int32](7),
		// 					GeoRedundantBackup: to.Ptr(armpostgresql.GeoRedundantBackupDisabled),
		// 					StorageMB: to.Ptr[int32](2048000),
		// 				},
		// 				UserVisibleState: to.Ptr(armpostgresql.ServerStateReady),
		// 				Version: to.Ptr(armpostgresql.ServerVersionNine6),
		// 			},
		// 			SKU: &armpostgresql.SKU{
		// 				Name: to.Ptr("GP_Gen4_16"),
		// 				Capacity: to.Ptr[int32](16),
		// 				Family: to.Ptr("Gen4"),
		// 				Tier: to.Ptr(armpostgresql.SKUTierGeneralPurpose),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("testserver-replica5"),
		// 			Type: to.Ptr("Microsoft.DBforPostgreSQL/servers"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestGroup_WestCentralUS/providers/Microsoft.DBforPostgreSQL/servers/testserver-replica5"),
		// 			Location: to.Ptr("westcentralus"),
		// 			Properties: &armpostgresql.ServerProperties{
		// 				AdministratorLogin: to.Ptr("postgres"),
		// 				EarliestRestoreDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-12T00:05:03.269Z"); return t}()),
		// 				FullyQualifiedDomainName: to.Ptr("testserver-replica5.postgres.database.azure.com"),
		// 				MasterServerID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestGroup_WestCentralUS/providers/Microsoft.DBforPostgreSQL/servers/testserver-master"),
		// 				ReplicaCapacity: to.Ptr[int32](0),
		// 				ReplicationRole: to.Ptr("Replica"),
		// 				SSLEnforcement: to.Ptr(armpostgresql.SSLEnforcementEnumDisabled),
		// 				StorageProfile: &armpostgresql.StorageProfile{
		// 					BackupRetentionDays: to.Ptr[int32](7),
		// 					GeoRedundantBackup: to.Ptr(armpostgresql.GeoRedundantBackupDisabled),
		// 					StorageMB: to.Ptr[int32](2048000),
		// 				},
		// 				UserVisibleState: to.Ptr(armpostgresql.ServerStateReady),
		// 				Version: to.Ptr(armpostgresql.ServerVersionNine6),
		// 			},
		// 			SKU: &armpostgresql.SKU{
		// 				Name: to.Ptr("GP_Gen4_16"),
		// 				Capacity: to.Ptr[int32](16),
		// 				Family: to.Ptr("Gen4"),
		// 				Tier: to.Ptr(armpostgresql.SKUTierGeneralPurpose),
		// 			},
		// 	}},
		// }
	}
}
