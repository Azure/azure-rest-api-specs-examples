package armmariadb_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mariadb/armmariadb"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/mariadb/resource-manager/Microsoft.DBforMariaDB/stable/2018-06-01/examples/ReplicasListByServer.json
func ExampleReplicasClient_NewListByServerPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmariadb.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewReplicasClient().NewListByServerPager("TestGroup", "testmaster", nil)
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
		// page.ServerListResult = armmariadb.ServerListResult{
		// 	Value: []*armmariadb.Server{
		// 		{
		// 			Name: to.Ptr("testserver"),
		// 			Type: to.Ptr("Microsoft.DBforMariaDB/servers"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestGroup/providers/Microsoft.DBforMariaDB/servers/testserver"),
		// 			Location: to.Ptr("northeurope"),
		// 			Tags: map[string]*string{
		// 				"elasticServer": to.Ptr("1"),
		// 			},
		// 			Properties: &armmariadb.ServerProperties{
		// 				AdministratorLogin: to.Ptr("cloudsa"),
		// 				EarliestRestoreDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-11T23:56:54.300Z"); return t}()),
		// 				FullyQualifiedDomainName: to.Ptr("testserver.mariadb.database.azure.com"),
		// 				MasterServerID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestGroup/providers/Microsoft.DBforMariaDB/servers/testmaster"),
		// 				ReplicaCapacity: to.Ptr[int32](0),
		// 				ReplicationRole: to.Ptr("Replica"),
		// 				SSLEnforcement: to.Ptr(armmariadb.SSLEnforcementEnumEnabled),
		// 				StorageProfile: &armmariadb.StorageProfile{
		// 					BackupRetentionDays: to.Ptr[int32](35),
		// 					GeoRedundantBackup: to.Ptr(armmariadb.GeoRedundantBackupEnabled),
		// 					StorageMB: to.Ptr[int32](256000),
		// 				},
		// 				UserVisibleState: to.Ptr(armmariadb.ServerStateReady),
		// 				Version: to.Ptr(armmariadb.ServerVersionTen2),
		// 			},
		// 			SKU: &armmariadb.SKU{
		// 				Name: to.Ptr("GP_Gen4_2"),
		// 				Capacity: to.Ptr[int32](2),
		// 				Family: to.Ptr("Gen4"),
		// 				Tier: to.Ptr(armmariadb.SKUTierGeneralPurpose),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("testserver1"),
		// 			Type: to.Ptr("Microsoft.DBforMariaDB/servers"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestGroup/providers/Microsoft.DBforMariaDB/servers/testserver1"),
		// 			Location: to.Ptr("northeurope"),
		// 			Tags: map[string]*string{
		// 				"elasticServer": to.Ptr("1"),
		// 			},
		// 			Properties: &armmariadb.ServerProperties{
		// 				AdministratorLogin: to.Ptr("cloudsa"),
		// 				EarliestRestoreDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-11T23:56:54.300Z"); return t}()),
		// 				FullyQualifiedDomainName: to.Ptr("testserver1.mariadb.database.azure.com"),
		// 				MasterServerID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestGroup/providers/Microsoft.DBforMariaDB/servers/testmaster"),
		// 				ReplicaCapacity: to.Ptr[int32](0),
		// 				ReplicationRole: to.Ptr("Replica"),
		// 				SSLEnforcement: to.Ptr(armmariadb.SSLEnforcementEnumEnabled),
		// 				StorageProfile: &armmariadb.StorageProfile{
		// 					BackupRetentionDays: to.Ptr[int32](35),
		// 					GeoRedundantBackup: to.Ptr(armmariadb.GeoRedundantBackupEnabled),
		// 					StorageMB: to.Ptr[int32](256000),
		// 				},
		// 				UserVisibleState: to.Ptr(armmariadb.ServerStateReady),
		// 				Version: to.Ptr(armmariadb.ServerVersionTen2),
		// 			},
		// 			SKU: &armmariadb.SKU{
		// 				Name: to.Ptr("GP_Gen4_2"),
		// 				Capacity: to.Ptr[int32](2),
		// 				Family: to.Ptr("Gen4"),
		// 				Tier: to.Ptr(armmariadb.SKUTierGeneralPurpose),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("testserver2"),
		// 			Type: to.Ptr("Microsoft.DBforMariaDB/servers"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestGroup/providers/Microsoft.DBforMariaDB/servers/testserver2"),
		// 			Location: to.Ptr("northeurope"),
		// 			Tags: map[string]*string{
		// 				"elasticServer": to.Ptr("1"),
		// 			},
		// 			Properties: &armmariadb.ServerProperties{
		// 				AdministratorLogin: to.Ptr("cloudsa"),
		// 				EarliestRestoreDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-11T23:56:54.300Z"); return t}()),
		// 				FullyQualifiedDomainName: to.Ptr("testserver2.mariadb.database.azure.com"),
		// 				MasterServerID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestGroup/providers/Microsoft.DBforMariaDB/servers/testmaster"),
		// 				ReplicaCapacity: to.Ptr[int32](0),
		// 				ReplicationRole: to.Ptr("Replica"),
		// 				SSLEnforcement: to.Ptr(armmariadb.SSLEnforcementEnumEnabled),
		// 				StorageProfile: &armmariadb.StorageProfile{
		// 					BackupRetentionDays: to.Ptr[int32](35),
		// 					GeoRedundantBackup: to.Ptr(armmariadb.GeoRedundantBackupEnabled),
		// 					StorageMB: to.Ptr[int32](256000),
		// 				},
		// 				UserVisibleState: to.Ptr(armmariadb.ServerStateReady),
		// 				Version: to.Ptr(armmariadb.ServerVersionTen2),
		// 			},
		// 			SKU: &armmariadb.SKU{
		// 				Name: to.Ptr("GP_Gen4_2"),
		// 				Capacity: to.Ptr[int32](2),
		// 				Family: to.Ptr("Gen4"),
		// 				Tier: to.Ptr(armmariadb.SKUTierGeneralPurpose),
		// 			},
		// 	}},
		// }
	}
}
