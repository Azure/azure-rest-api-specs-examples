package armpostgresqlflexibleservers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresqlflexibleservers/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b1f4d539964453ce8008e4b069e59885e12ba441/specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2025-01-01-preview/examples/ServerList.json
func ExampleServersClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpostgresqlflexibleservers.NewClientFactory("<subscription-id>", cred, nil)
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
		// page.ServerListResult = armpostgresqlflexibleservers.ServerListResult{
		// 	Value: []*armpostgresqlflexibleservers.Server{
		// 		{
		// 			Name: to.Ptr("pgtest2"),
		// 			Type: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrgn/providers/Microsoft.DBforPostgreSQL/flexibleServers/pgtest2"),
		// 			SystemData: &armpostgresqlflexibleservers.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-03-03T17:44:30.709Z"); return t}()),
		// 			},
		// 			Location: to.Ptr("East US"),
		// 			Tags: map[string]*string{
		// 				"VnetServer": to.Ptr("1"),
		// 			},
		// 			Properties: &armpostgresqlflexibleservers.ServerProperties{
		// 				AdministratorLogin: to.Ptr("login"),
		// 				AuthConfig: &armpostgresqlflexibleservers.AuthConfig{
		// 					ActiveDirectoryAuth: to.Ptr(armpostgresqlflexibleservers.ActiveDirectoryAuthEnumEnabled),
		// 					PasswordAuth: to.Ptr(armpostgresqlflexibleservers.PasswordAuthEnumEnabled),
		// 					TenantID: to.Ptr("tttttt-tttt-tttt-tttt-tttttttttttt"),
		// 				},
		// 				AvailabilityZone: to.Ptr("1"),
		// 				Backup: &armpostgresqlflexibleservers.Backup{
		// 					BackupRetentionDays: to.Ptr[int32](7),
		// 					EarliestRestoreDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-03-03T17:49:56.940Z"); return t}()),
		// 					GeoRedundantBackup: to.Ptr(armpostgresqlflexibleservers.GeoRedundantBackupEnumEnabled),
		// 				},
		// 				DataEncryption: &armpostgresqlflexibleservers.DataEncryption{
		// 					Type: to.Ptr(armpostgresqlflexibleservers.ArmServerKeyTypeSystemManaged),
		// 				},
		// 				FullyQualifiedDomainName: to.Ptr("pgtest2.postgres.database.azure.com"),
		// 				HighAvailability: &armpostgresqlflexibleservers.HighAvailability{
		// 					Mode: to.Ptr(armpostgresqlflexibleservers.HighAvailabilityModeZoneRedundant),
		// 					StandbyAvailabilityZone: to.Ptr("2"),
		// 					State: to.Ptr(armpostgresqlflexibleservers.ServerHAStateHealthy),
		// 				},
		// 				MaintenanceWindow: &armpostgresqlflexibleservers.MaintenanceWindow{
		// 					CustomWindow: to.Ptr("Disabled"),
		// 					DayOfWeek: to.Ptr[int32](0),
		// 					StartHour: to.Ptr[int32](0),
		// 					StartMinute: to.Ptr[int32](0),
		// 				},
		// 				MinorVersion: to.Ptr("8"),
		// 				Network: &armpostgresqlflexibleservers.Network{
		// 					DelegatedSubnetResourceID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrgn/providers/Microsoft.Network/virtualNetworks/test-vnet/subnets/test-subnet"),
		// 					PrivateDNSZoneArmResourceID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrgn/providers/Microsoft.Network/privateDnsZones/pgtest2.private.postgres.database.azure.com"),
		// 					PublicNetworkAccess: to.Ptr(armpostgresqlflexibleservers.ServerPublicNetworkAccessStateDisabled),
		// 				},
		// 				Replica: &armpostgresqlflexibleservers.Replica{
		// 					Capacity: to.Ptr[int32](5),
		// 					Role: to.Ptr(armpostgresqlflexibleservers.ReplicationRolePrimary),
		// 				},
		// 				ReplicaCapacity: to.Ptr[int32](5),
		// 				ReplicationRole: to.Ptr(armpostgresqlflexibleservers.ReplicationRolePrimary),
		// 				State: to.Ptr(armpostgresqlflexibleservers.ServerStateReady),
		// 				Storage: &armpostgresqlflexibleservers.Storage{
		// 					Type: to.Ptr(armpostgresqlflexibleservers.StorageType("")),
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
		// 		},
		// 		{
		// 			Name: to.Ptr("pgtest3"),
		// 			Type: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg6/providers/Microsoft.DBforPostgreSQL/flexibleServers/pgtest3"),
		// 			SystemData: &armpostgresqlflexibleservers.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-03-03T23:41:01.102Z"); return t}()),
		// 			},
		// 			Location: to.Ptr("East US"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armpostgresqlflexibleservers.ServerProperties{
		// 				AdministratorLogin: to.Ptr("login"),
		// 				AuthConfig: &armpostgresqlflexibleservers.AuthConfig{
		// 					ActiveDirectoryAuth: to.Ptr(armpostgresqlflexibleservers.ActiveDirectoryAuthEnumEnabled),
		// 					PasswordAuth: to.Ptr(armpostgresqlflexibleservers.PasswordAuthEnumEnabled),
		// 					TenantID: to.Ptr("tttttt-tttt-tttt-tttt-tttttttttttt"),
		// 				},
		// 				AvailabilityZone: to.Ptr("2"),
		// 				Backup: &armpostgresqlflexibleservers.Backup{
		// 					BackupRetentionDays: to.Ptr[int32](7),
		// 					EarliestRestoreDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-03-03T23:48:24.491Z"); return t}()),
		// 					GeoRedundantBackup: to.Ptr(armpostgresqlflexibleservers.GeoRedundantBackupEnumEnabled),
		// 				},
		// 				DataEncryption: &armpostgresqlflexibleservers.DataEncryption{
		// 					Type: to.Ptr(armpostgresqlflexibleservers.ArmServerKeyTypeSystemManaged),
		// 				},
		// 				FullyQualifiedDomainName: to.Ptr("pgtest3.postgres.database.azure.com"),
		// 				HighAvailability: &armpostgresqlflexibleservers.HighAvailability{
		// 					Mode: to.Ptr(armpostgresqlflexibleservers.HighAvailabilityModeZoneRedundant),
		// 					StandbyAvailabilityZone: to.Ptr("1"),
		// 					State: to.Ptr(armpostgresqlflexibleservers.ServerHAStateHealthy),
		// 				},
		// 				MaintenanceWindow: &armpostgresqlflexibleservers.MaintenanceWindow{
		// 					CustomWindow: to.Ptr("Disabled"),
		// 					DayOfWeek: to.Ptr[int32](0),
		// 					StartHour: to.Ptr[int32](0),
		// 					StartMinute: to.Ptr[int32](0),
		// 				},
		// 				MinorVersion: to.Ptr("8"),
		// 				Network: &armpostgresqlflexibleservers.Network{
		// 					PublicNetworkAccess: to.Ptr(armpostgresqlflexibleservers.ServerPublicNetworkAccessStateEnabled),
		// 				},
		// 				PrivateEndpointConnections: []*armpostgresqlflexibleservers.PrivateEndpointConnection{
		// 				},
		// 				Replica: &armpostgresqlflexibleservers.Replica{
		// 					Capacity: to.Ptr[int32](5),
		// 					Role: to.Ptr(armpostgresqlflexibleservers.ReplicationRolePrimary),
		// 				},
		// 				ReplicaCapacity: to.Ptr[int32](5),
		// 				ReplicationRole: to.Ptr(armpostgresqlflexibleservers.ReplicationRolePrimary),
		// 				State: to.Ptr(armpostgresqlflexibleservers.ServerStateReady),
		// 				Storage: &armpostgresqlflexibleservers.Storage{
		// 					Type: to.Ptr(armpostgresqlflexibleservers.StorageType("")),
		// 					AutoGrow: to.Ptr(armpostgresqlflexibleservers.StorageAutoGrowDisabled),
		// 					Iops: to.Ptr[int32](500),
		// 					StorageSizeGB: to.Ptr[int32](128),
		// 					Tier: to.Ptr(armpostgresqlflexibleservers.AzureManagedDiskPerformanceTiersP10),
		// 				},
		// 				Version: to.Ptr(armpostgresqlflexibleservers.ServerVersionSixteen),
		// 			},
		// 			SKU: &armpostgresqlflexibleservers.SKU{
		// 				Name: to.Ptr("Standard_D4ds_v5"),
		// 				Tier: to.Ptr(armpostgresqlflexibleservers.SKUTierGeneralPurpose),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("pgtest6"),
		// 			Type: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrgn/providers/Microsoft.DBforPostgreSQL/flexibleServers/pgtest6"),
		// 			SystemData: &armpostgresqlflexibleservers.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-11-23T00:13:34.864Z"); return t}()),
		// 			},
		// 			Location: to.Ptr("East US 2"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Identity: &armpostgresqlflexibleservers.UserAssignedIdentity{
		// 				Type: to.Ptr(armpostgresqlflexibleservers.IdentityTypeSystemAssignedUserAssigned),
		// 				PrincipalID: to.Ptr("rrrrrrrr-rrrr-rrrr-rrrr-rrrrrrrrrrrr"),
		// 				TenantID: to.Ptr("tttttt-tttt-tttt-tttt-tttttttttttt"),
		// 				UserAssignedIdentities: map[string]*armpostgresqlflexibleservers.UserIdentity{
		// 					"/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrgn/providers/Microsoft.ManagedIdentity/userAssignedIdentities/test-user": &armpostgresqlflexibleservers.UserIdentity{
		// 						ClientID: to.Ptr("tttttt-tttt-tttt-tttt-tttttttttttt"),
		// 						PrincipalID: to.Ptr("pppppppp-pppp-pppp-pppp-pppppppppppp"),
		// 					},
		// 				},
		// 			},
		// 			Properties: &armpostgresqlflexibleservers.ServerProperties{
		// 				AdministratorLogin: to.Ptr("login"),
		// 				AuthConfig: &armpostgresqlflexibleservers.AuthConfig{
		// 					ActiveDirectoryAuth: to.Ptr(armpostgresqlflexibleservers.ActiveDirectoryAuthEnumEnabled),
		// 					PasswordAuth: to.Ptr(armpostgresqlflexibleservers.PasswordAuthEnumEnabled),
		// 					TenantID: to.Ptr("tttttt-tttt-tttt-tttt-tttttttttttt"),
		// 				},
		// 				AvailabilityZone: to.Ptr("1"),
		// 				Backup: &armpostgresqlflexibleservers.Backup{
		// 					BackupRetentionDays: to.Ptr[int32](7),
		// 					EarliestRestoreDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-02-25T01:24:56.297Z"); return t}()),
		// 					GeoRedundantBackup: to.Ptr(armpostgresqlflexibleservers.GeoRedundantBackupEnumDisabled),
		// 				},
		// 				DataEncryption: &armpostgresqlflexibleservers.DataEncryption{
		// 					Type: to.Ptr(armpostgresqlflexibleservers.ArmServerKeyTypeAzureKeyVault),
		// 					PrimaryEncryptionKeyStatus: to.Ptr(armpostgresqlflexibleservers.KeyStatusEnumValid),
		// 					PrimaryKeyURI: to.Ptr("https://test-vault1.vault.azure.net/keys/test-key1/xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"),
		// 					PrimaryUserAssignedIdentityID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg2/providers/Microsoft.ManagedIdentity/userAssignedIdentities/test-user"),
		// 				},
		// 				FullyQualifiedDomainName: to.Ptr("pgtest6.postgres.database.azure.com"),
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
		// 				MinorVersion: to.Ptr("8"),
		// 				Network: &armpostgresqlflexibleservers.Network{
		// 					PublicNetworkAccess: to.Ptr(armpostgresqlflexibleservers.ServerPublicNetworkAccessStateEnabled),
		// 				},
		// 				PrivateEndpointConnections: []*armpostgresqlflexibleservers.PrivateEndpointConnection{
		// 				},
		// 				Replica: &armpostgresqlflexibleservers.Replica{
		// 					Capacity: to.Ptr[int32](5),
		// 					Role: to.Ptr(armpostgresqlflexibleservers.ReplicationRolePrimary),
		// 				},
		// 				ReplicaCapacity: to.Ptr[int32](5),
		// 				ReplicationRole: to.Ptr(armpostgresqlflexibleservers.ReplicationRolePrimary),
		// 				State: to.Ptr(armpostgresqlflexibleservers.ServerStateReady),
		// 				Storage: &armpostgresqlflexibleservers.Storage{
		// 					Type: to.Ptr(armpostgresqlflexibleservers.StorageType("")),
		// 					AutoGrow: to.Ptr(armpostgresqlflexibleservers.StorageAutoGrowDisabled),
		// 					Iops: to.Ptr[int32](500),
		// 					StorageSizeGB: to.Ptr[int32](128),
		// 					Tier: to.Ptr(armpostgresqlflexibleservers.AzureManagedDiskPerformanceTiersP10),
		// 				},
		// 				Version: to.Ptr(armpostgresqlflexibleservers.ServerVersionSixteen),
		// 			},
		// 			SKU: &armpostgresqlflexibleservers.SKU{
		// 				Name: to.Ptr("Standard_D4ds_v5"),
		// 				Tier: to.Ptr(armpostgresqlflexibleservers.SKUTierGeneralPurpose),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("pgtest9"),
		// 			Type: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg9/providers/Microsoft.DBforPostgreSQL/flexibleServers/pgtest9"),
		// 			SystemData: &armpostgresqlflexibleservers.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-02-26T21:17:37.501Z"); return t}()),
		// 			},
		// 			Location: to.Ptr("East US 2"),
		// 			Tags: map[string]*string{
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
		// 					EarliestRestoreDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-02-26T21:41:19.335Z"); return t}()),
		// 					GeoRedundantBackup: to.Ptr(armpostgresqlflexibleservers.GeoRedundantBackupEnumDisabled),
		// 				},
		// 				DataEncryption: &armpostgresqlflexibleservers.DataEncryption{
		// 					Type: to.Ptr(armpostgresqlflexibleservers.ArmServerKeyTypeSystemManaged),
		// 				},
		// 				FullyQualifiedDomainName: to.Ptr("pgtest9.postgres.database.azure.com"),
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
		// 				MinorVersion: to.Ptr("8"),
		// 				Network: &armpostgresqlflexibleservers.Network{
		// 					PublicNetworkAccess: to.Ptr(armpostgresqlflexibleservers.ServerPublicNetworkAccessStateEnabled),
		// 				},
		// 				PrivateEndpointConnections: []*armpostgresqlflexibleservers.PrivateEndpointConnection{
		// 				},
		// 				Replica: &armpostgresqlflexibleservers.Replica{
		// 					Capacity: to.Ptr[int32](5),
		// 					Role: to.Ptr(armpostgresqlflexibleservers.ReplicationRolePrimary),
		// 				},
		// 				ReplicaCapacity: to.Ptr[int32](5),
		// 				ReplicationRole: to.Ptr(armpostgresqlflexibleservers.ReplicationRolePrimary),
		// 				State: to.Ptr(armpostgresqlflexibleservers.ServerStateReady),
		// 				Storage: &armpostgresqlflexibleservers.Storage{
		// 					Type: to.Ptr(armpostgresqlflexibleservers.StorageTypePremiumV2LRS),
		// 					AutoGrow: to.Ptr(armpostgresqlflexibleservers.StorageAutoGrowDisabled),
		// 					Iops: to.Ptr[int32](3000),
		// 					StorageSizeGB: to.Ptr[int32](512),
		// 					Throughput: to.Ptr[int32](125),
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
