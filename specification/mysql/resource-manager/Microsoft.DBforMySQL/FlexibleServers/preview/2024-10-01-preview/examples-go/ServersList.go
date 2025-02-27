package armmysqlflexibleservers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysqlflexibleservers/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e26b89bcbec9eed5026c01416e481408b2a1ca1a/specification/mysql/resource-manager/Microsoft.DBforMySQL/FlexibleServers/preview/2024-10-01-preview/examples/ServersList.json
func ExampleServersClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmysqlflexibleservers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewServersClient().NewListPager(nil)
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
		// page.ServerListResult = armmysqlflexibleservers.ServerListResult{
		// 	Value: []*armmysqlflexibleservers.Server{
		// 		{
		// 			Name: to.Ptr("mysqltestserver1"),
		// 			Type: to.Ptr("Microsoft.DBforMySQL/flexibleServers"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestGroup/providers/Microsoft.DBforMySQL/flexibleServers/mysqltestserver1"),
		// 			Location: to.Ptr("Southeast Asia"),
		// 			Tags: map[string]*string{
		// 				"num": to.Ptr("1"),
		// 			},
		// 			Properties: &armmysqlflexibleservers.ServerProperties{
		// 				AdministratorLogin: to.Ptr("cloudsa"),
		// 				AvailabilityZone: to.Ptr("1"),
		// 				Backup: &armmysqlflexibleservers.Backup{
		// 					BackupIntervalHours: to.Ptr[int32](24),
		// 					BackupRetentionDays: to.Ptr[int32](7),
		// 					EarliestRestoreDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-17T07:08:17.425Z"); return t}()),
		// 					GeoRedundantBackup: to.Ptr(armmysqlflexibleservers.EnableStatusEnumDisabled),
		// 				},
		// 				DatabasePort: to.Ptr[int32](3306),
		// 				FullVersion: to.Ptr("5.7.44"),
		// 				FullyQualifiedDomainName: to.Ptr("mysqltestserver1.database.mysql.azure.com"),
		// 				HighAvailability: &armmysqlflexibleservers.HighAvailability{
		// 					Mode: to.Ptr(armmysqlflexibleservers.HighAvailabilityModeDisabled),
		// 					State: to.Ptr(armmysqlflexibleservers.HighAvailabilityStateNotEnabled),
		// 				},
		// 				MaintenancePolicy: &armmysqlflexibleservers.MaintenancePolicy{
		// 					PatchStrategy: to.Ptr(armmysqlflexibleservers.PatchStrategyRegular),
		// 				},
		// 				MaintenanceWindow: &armmysqlflexibleservers.MaintenanceWindow{
		// 					CustomWindow: to.Ptr("Disabled"),
		// 					DayOfWeek: to.Ptr[int32](0),
		// 					StartHour: to.Ptr[int32](0),
		// 					StartMinute: to.Ptr[int32](0),
		// 				},
		// 				Network: &armmysqlflexibleservers.Network{
		// 					PublicNetworkAccess: to.Ptr(armmysqlflexibleservers.EnableStatusEnumEnabled),
		// 				},
		// 				ReplicaCapacity: to.Ptr[int32](10),
		// 				ReplicationRole: to.Ptr(armmysqlflexibleservers.ReplicationRoleNone),
		// 				State: to.Ptr(armmysqlflexibleservers.ServerStateReady),
		// 				Storage: &armmysqlflexibleservers.Storage{
		// 					AutoGrow: to.Ptr(armmysqlflexibleservers.EnableStatusEnumEnabled),
		// 					Iops: to.Ptr[int32](369),
		// 					StorageRedundancy: to.Ptr(armmysqlflexibleservers.StorageRedundancyEnumLocalRedundancy),
		// 					StorageSizeGB: to.Ptr[int32](23),
		// 					StorageSKU: to.Ptr("Premium_LRS"),
		// 				},
		// 				Version: to.Ptr(armmysqlflexibleservers.ServerVersionFive7),
		// 			},
		// 			SKU: &armmysqlflexibleservers.MySQLServerSKU{
		// 				Name: to.Ptr("Standard_B1ms"),
		// 				Tier: to.Ptr(armmysqlflexibleservers.ServerSKUTierBurstable),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("mysqltestserver2"),
		// 			Type: to.Ptr("Microsoft.DBforMySQL/flexibleServers"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestGroup2/providers/Microsoft.DBforMySQL/flexibleServers/mysqltestserver2"),
		// 			Location: to.Ptr("Southeast Asia"),
		// 			Tags: map[string]*string{
		// 				"num": to.Ptr("1"),
		// 			},
		// 			Properties: &armmysqlflexibleservers.ServerProperties{
		// 				AdministratorLogin: to.Ptr("cloudsa"),
		// 				AvailabilityZone: to.Ptr("2"),
		// 				Backup: &armmysqlflexibleservers.Backup{
		// 					BackupIntervalHours: to.Ptr[int32](24),
		// 					BackupRetentionDays: to.Ptr[int32](7),
		// 					EarliestRestoreDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-17T07:08:17.425Z"); return t}()),
		// 					GeoRedundantBackup: to.Ptr(armmysqlflexibleservers.EnableStatusEnumDisabled),
		// 				},
		// 				DatabasePort: to.Ptr[int32](3306),
		// 				FullVersion: to.Ptr("5.7.44"),
		// 				FullyQualifiedDomainName: to.Ptr("mysqltestserver2.mysql.database.azure.com"),
		// 				HighAvailability: &armmysqlflexibleservers.HighAvailability{
		// 					Mode: to.Ptr(armmysqlflexibleservers.HighAvailabilityModeDisabled),
		// 					State: to.Ptr(armmysqlflexibleservers.HighAvailabilityStateNotEnabled),
		// 				},
		// 				MaintenancePolicy: &armmysqlflexibleservers.MaintenancePolicy{
		// 					PatchStrategy: to.Ptr(armmysqlflexibleservers.PatchStrategy("Default")),
		// 				},
		// 				MaintenanceWindow: &armmysqlflexibleservers.MaintenanceWindow{
		// 					CustomWindow: to.Ptr("Disabled"),
		// 					DayOfWeek: to.Ptr[int32](0),
		// 					StartHour: to.Ptr[int32](0),
		// 					StartMinute: to.Ptr[int32](0),
		// 				},
		// 				Network: &armmysqlflexibleservers.Network{
		// 					PublicNetworkAccess: to.Ptr(armmysqlflexibleservers.EnableStatusEnumEnabled),
		// 				},
		// 				ReplicaCapacity: to.Ptr[int32](10),
		// 				ReplicationRole: to.Ptr(armmysqlflexibleservers.ReplicationRoleNone),
		// 				State: to.Ptr(armmysqlflexibleservers.ServerStateReady),
		// 				Storage: &armmysqlflexibleservers.Storage{
		// 					AutoGrow: to.Ptr(armmysqlflexibleservers.EnableStatusEnumEnabled),
		// 					Iops: to.Ptr[int32](369),
		// 					StorageRedundancy: to.Ptr(armmysqlflexibleservers.StorageRedundancyEnumLocalRedundancy),
		// 					StorageSizeGB: to.Ptr[int32](23),
		// 					StorageSKU: to.Ptr("Premium_LRS"),
		// 				},
		// 				Version: to.Ptr(armmysqlflexibleservers.ServerVersionFive7),
		// 			},
		// 			SKU: &armmysqlflexibleservers.MySQLServerSKU{
		// 				Name: to.Ptr("Standard_D2ds_v4"),
		// 				Tier: to.Ptr(armmysqlflexibleservers.ServerSKUTierGeneralPurpose),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("mysqltestserver3"),
		// 			Type: to.Ptr("Microsoft.DBforMySQL/flexibleServers"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestGroup3/providers/Microsoft.DBforMySQL/flexibleServers/mysqltestserver3"),
		// 			Location: to.Ptr("Southeast Asia"),
		// 			Tags: map[string]*string{
		// 				"num": to.Ptr("1"),
		// 			},
		// 			Properties: &armmysqlflexibleservers.ServerProperties{
		// 				AdministratorLogin: to.Ptr("cloudsa"),
		// 				AvailabilityZone: to.Ptr("1"),
		// 				Backup: &armmysqlflexibleservers.Backup{
		// 					BackupIntervalHours: to.Ptr[int32](24),
		// 					BackupRetentionDays: to.Ptr[int32](7),
		// 					EarliestRestoreDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-24T06:28:19.061Z"); return t}()),
		// 					GeoRedundantBackup: to.Ptr(armmysqlflexibleservers.EnableStatusEnumDisabled),
		// 				},
		// 				DatabasePort: to.Ptr[int32](3306),
		// 				FullVersion: to.Ptr("5.7.44"),
		// 				FullyQualifiedDomainName: to.Ptr("mysqltestserver3.mysql.database.azure.com"),
		// 				HighAvailability: &armmysqlflexibleservers.HighAvailability{
		// 					Mode: to.Ptr(armmysqlflexibleservers.HighAvailabilityModeDisabled),
		// 					State: to.Ptr(armmysqlflexibleservers.HighAvailabilityStateNotEnabled),
		// 				},
		// 				MaintenancePolicy: &armmysqlflexibleservers.MaintenancePolicy{
		// 					PatchStrategy: to.Ptr(armmysqlflexibleservers.PatchStrategy("Default")),
		// 				},
		// 				MaintenanceWindow: &armmysqlflexibleservers.MaintenanceWindow{
		// 					CustomWindow: to.Ptr("Disabled"),
		// 					DayOfWeek: to.Ptr[int32](0),
		// 					StartHour: to.Ptr[int32](0),
		// 					StartMinute: to.Ptr[int32](0),
		// 				},
		// 				Network: &armmysqlflexibleservers.Network{
		// 					PublicNetworkAccess: to.Ptr(armmysqlflexibleservers.EnableStatusEnumEnabled),
		// 				},
		// 				ReplicaCapacity: to.Ptr[int32](10),
		// 				ReplicationRole: to.Ptr(armmysqlflexibleservers.ReplicationRoleNone),
		// 				State: to.Ptr(armmysqlflexibleservers.ServerStateReady),
		// 				Storage: &armmysqlflexibleservers.Storage{
		// 					AutoGrow: to.Ptr(armmysqlflexibleservers.EnableStatusEnumEnabled),
		// 					Iops: to.Ptr[int32](600),
		// 					StorageRedundancy: to.Ptr(armmysqlflexibleservers.StorageRedundancyEnumZoneRedundancy),
		// 					StorageSizeGB: to.Ptr[int32](100),
		// 					StorageSKU: to.Ptr("Premium_ZRS"),
		// 				},
		// 				Version: to.Ptr(armmysqlflexibleservers.ServerVersionFive7),
		// 			},
		// 			SKU: &armmysqlflexibleservers.MySQLServerSKU{
		// 				Name: to.Ptr("Standard_E2ds_v4"),
		// 				Tier: to.Ptr(armmysqlflexibleservers.ServerSKUTierMemoryOptimized),
		// 			},
		// 	}},
		// }
	}
}
