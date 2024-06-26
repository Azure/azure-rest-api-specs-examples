package armmysql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysql"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2017-12-01/examples/ReplicasListByServer.json
func ExampleReplicasClient_NewListByServerPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmysql.NewClientFactory("<subscription-id>", cred, nil)
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
		// page.ServerListResult = armmysql.ServerListResult{
		// 	Value: []*armmysql.Server{
		// 		{
		// 			Name: to.Ptr("testserver"),
		// 			Type: to.Ptr("Microsoft.DBforMySQL/servers"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestGroup/providers/Microsoft.DBforMySQL/servers/testserver"),
		// 			Location: to.Ptr("northeurope"),
		// 			Tags: map[string]*string{
		// 				"elasticServer": to.Ptr("1"),
		// 			},
		// 			Properties: &armmysql.ServerProperties{
		// 				AdministratorLogin: to.Ptr("cloudsa"),
		// 				EarliestRestoreDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-11T23:56:54.300Z"); return t}()),
		// 				FullyQualifiedDomainName: to.Ptr("testserver.mysql.database.azure.com"),
		// 				MasterServerID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestGroup/providers/Microsoft.DBforMySQL/servers/testmaster"),
		// 				ReplicaCapacity: to.Ptr[int32](0),
		// 				ReplicationRole: to.Ptr("Replica"),
		// 				SSLEnforcement: to.Ptr(armmysql.SSLEnforcementEnumEnabled),
		// 				StorageProfile: &armmysql.StorageProfile{
		// 					BackupRetentionDays: to.Ptr[int32](35),
		// 					GeoRedundantBackup: to.Ptr(armmysql.GeoRedundantBackupEnabled),
		// 					StorageMB: to.Ptr[int32](256000),
		// 				},
		// 				UserVisibleState: to.Ptr(armmysql.ServerStateReady),
		// 				Version: to.Ptr(armmysql.ServerVersionFive6),
		// 			},
		// 			SKU: &armmysql.SKU{
		// 				Name: to.Ptr("GP_Gen4_2"),
		// 				Capacity: to.Ptr[int32](2),
		// 				Family: to.Ptr("Gen4"),
		// 				Tier: to.Ptr(armmysql.SKUTierGeneralPurpose),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("testserver1"),
		// 			Type: to.Ptr("Microsoft.DBforMySQL/servers"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestGroup/providers/Microsoft.DBforMySQL/servers/testserver1"),
		// 			Location: to.Ptr("northeurope"),
		// 			Tags: map[string]*string{
		// 				"elasticServer": to.Ptr("1"),
		// 			},
		// 			Properties: &armmysql.ServerProperties{
		// 				AdministratorLogin: to.Ptr("cloudsa"),
		// 				EarliestRestoreDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-11T23:56:54.300Z"); return t}()),
		// 				FullyQualifiedDomainName: to.Ptr("testserver1.mysql.database.azure.com"),
		// 				MasterServerID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestGroup/providers/Microsoft.DBforMySQL/servers/testmaster"),
		// 				ReplicaCapacity: to.Ptr[int32](0),
		// 				ReplicationRole: to.Ptr("Replica"),
		// 				SSLEnforcement: to.Ptr(armmysql.SSLEnforcementEnumEnabled),
		// 				StorageProfile: &armmysql.StorageProfile{
		// 					BackupRetentionDays: to.Ptr[int32](35),
		// 					GeoRedundantBackup: to.Ptr(armmysql.GeoRedundantBackupEnabled),
		// 					StorageMB: to.Ptr[int32](256000),
		// 				},
		// 				UserVisibleState: to.Ptr(armmysql.ServerStateReady),
		// 				Version: to.Ptr(armmysql.ServerVersionFive6),
		// 			},
		// 			SKU: &armmysql.SKU{
		// 				Name: to.Ptr("GP_Gen4_2"),
		// 				Capacity: to.Ptr[int32](2),
		// 				Family: to.Ptr("Gen4"),
		// 				Tier: to.Ptr(armmysql.SKUTierGeneralPurpose),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("testserver2"),
		// 			Type: to.Ptr("Microsoft.DBforMySQL/servers"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestGroup/providers/Microsoft.DBforMySQL/servers/testserver2"),
		// 			Location: to.Ptr("northeurope"),
		// 			Tags: map[string]*string{
		// 				"elasticServer": to.Ptr("1"),
		// 			},
		// 			Properties: &armmysql.ServerProperties{
		// 				AdministratorLogin: to.Ptr("cloudsa"),
		// 				EarliestRestoreDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-11T23:56:54.300Z"); return t}()),
		// 				FullyQualifiedDomainName: to.Ptr("testserver2.mysql.database.azure.com"),
		// 				MasterServerID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestGroup/providers/Microsoft.DBforMySQL/servers/testmaster"),
		// 				ReplicaCapacity: to.Ptr[int32](0),
		// 				ReplicationRole: to.Ptr("Replica"),
		// 				SSLEnforcement: to.Ptr(armmysql.SSLEnforcementEnumEnabled),
		// 				StorageProfile: &armmysql.StorageProfile{
		// 					BackupRetentionDays: to.Ptr[int32](35),
		// 					GeoRedundantBackup: to.Ptr(armmysql.GeoRedundantBackupEnabled),
		// 					StorageMB: to.Ptr[int32](256000),
		// 				},
		// 				UserVisibleState: to.Ptr(armmysql.ServerStateReady),
		// 				Version: to.Ptr(armmysql.ServerVersionFive6),
		// 			},
		// 			SKU: &armmysql.SKU{
		// 				Name: to.Ptr("GP_Gen4_2"),
		// 				Capacity: to.Ptr[int32](2),
		// 				Family: to.Ptr("Gen4"),
		// 				Tier: to.Ptr(armmysql.SKUTierGeneralPurpose),
		// 			},
		// 	}},
		// }
	}
}
