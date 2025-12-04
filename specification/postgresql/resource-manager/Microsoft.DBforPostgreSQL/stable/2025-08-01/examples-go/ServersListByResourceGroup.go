package armpostgresqlflexibleservers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresqlflexibleservers/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e96c24570a484cff13d153fb472f812878866a39/specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2025-08-01/examples/ServersListByResourceGroup.json
func ExampleServersClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpostgresqlflexibleservers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewServersClient().NewListByResourceGroupPager("exampleresourcegroup", nil)
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
		// page.ServerList = armpostgresqlflexibleservers.ServerList{
		// 	Value: []*armpostgresqlflexibleservers.Server{
		// 		{
		// 			Name: to.Ptr("exampleserver1"),
		// 			Type: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/exampleresourcegroup/providers/Microsoft.DBforPostgreSQL/flexibleServers/exampleserver1"),
		// 			SystemData: &armpostgresqlflexibleservers.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-06-01T18:30:22.123Z"); return t}()),
		// 			},
		// 			Location: to.Ptr("eastus"),
		// 			Tags: map[string]*string{
		// 				"VnetServer": to.Ptr("1"),
		// 			},
		// 			Properties: &armpostgresqlflexibleservers.ServerProperties{
		// 				AdministratorLogin: to.Ptr("exampleadministratorlogin"),
		// 				AuthConfig: &armpostgresqlflexibleservers.AuthConfig{
		// 					ActiveDirectoryAuth: to.Ptr(armpostgresqlflexibleservers.MicrosoftEntraAuthEnabled),
		// 					PasswordAuth: to.Ptr(armpostgresqlflexibleservers.PasswordBasedAuthEnabled),
		// 					TenantID: to.Ptr("tttttt-tttt-tttt-tttt-tttttttttttt"),
		// 				},
		// 				AvailabilityZone: to.Ptr("1"),
		// 				Backup: &armpostgresqlflexibleservers.Backup{
		// 					BackupRetentionDays: to.Ptr[int32](7),
		// 					EarliestRestoreDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-06-01T18:35:22.123Z"); return t}()),
		// 					GeoRedundantBackup: to.Ptr(armpostgresqlflexibleservers.GeographicallyRedundantBackupEnabled),
		// 				},
		// 				DataEncryption: &armpostgresqlflexibleservers.DataEncryption{
		// 					Type: to.Ptr(armpostgresqlflexibleservers.DataEncryptionTypeSystemManaged),
		// 				},
		// 				FullyQualifiedDomainName: to.Ptr("exampleserver1.postgres.database.azure.com"),
		// 				HighAvailability: &armpostgresqlflexibleservers.HighAvailability{
		// 					Mode: to.Ptr(armpostgresqlflexibleservers.HighAvailabilityModeZoneRedundant),
		// 					StandbyAvailabilityZone: to.Ptr("2"),
		// 					State: to.Ptr(armpostgresqlflexibleservers.HighAvailabilityStateHealthy),
		// 				},
		// 				MaintenanceWindow: &armpostgresqlflexibleservers.MaintenanceWindow{
		// 					CustomWindow: to.Ptr("Disabled"),
		// 					DayOfWeek: to.Ptr[int32](0),
		// 					StartHour: to.Ptr[int32](0),
		// 					StartMinute: to.Ptr[int32](0),
		// 				},
		// 				MinorVersion: to.Ptr("5"),
		// 				Network: &armpostgresqlflexibleservers.Network{
		// 					DelegatedSubnetResourceID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/exampleresourcegroup/providers/Microsoft.Network/virtualNetworks/examplevirtualnetwork/subnets/examplesubnet"),
		// 					PrivateDNSZoneArmResourceID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/exampleresourcegroup/providers/Microsoft.Network/privateDnsZones/exampleserver1.private.postgres.database.azure.com"),
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
		// 					Tier: to.Ptr(armpostgresqlflexibleservers.AzureManagedDiskPerformanceTierP20),
		// 				},
		// 				Version: to.Ptr(armpostgresqlflexibleservers.PostgresMajorVersionSeventeen),
		// 			},
		// 			SKU: &armpostgresqlflexibleservers.SKU{
		// 				Name: to.Ptr("Standard_D4ds_v5"),
		// 				Tier: to.Ptr(armpostgresqlflexibleservers.SKUTierGeneralPurpose),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("exampleserver2"),
		// 			Type: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/exampleresourcegroup/providers/Microsoft.DBforPostgreSQL/flexibleServers/exampleserver2"),
		// 			SystemData: &armpostgresqlflexibleservers.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-06-01T19:30:22.123Z"); return t}()),
		// 			},
		// 			Location: to.Ptr("eastus"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Identity: &armpostgresqlflexibleservers.UserAssignedIdentity{
		// 				Type: to.Ptr(armpostgresqlflexibleservers.IdentityTypeSystemAssignedUserAssigned),
		// 				PrincipalID: to.Ptr("rrrrrrrr-rrrr-rrrr-rrrr-rrrrrrrrrrrr"),
		// 				TenantID: to.Ptr("tttttt-tttt-tttt-tttt-tttttttttttt"),
		// 				UserAssignedIdentities: map[string]*armpostgresqlflexibleservers.UserIdentity{
		// 					"/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/exampleresourcegroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/exampleprimaryidentity": &armpostgresqlflexibleservers.UserIdentity{
		// 						ClientID: to.Ptr("tttttt-tttt-tttt-tttt-tttttttttttt"),
		// 						PrincipalID: to.Ptr("pppppppp-pppp-pppp-pppp-pppppppppppp"),
		// 					},
		// 				},
		// 			},
		// 			Properties: &armpostgresqlflexibleservers.ServerProperties{
		// 				AdministratorLogin: to.Ptr("exampleadministratorlogin"),
		// 				AuthConfig: &armpostgresqlflexibleservers.AuthConfig{
		// 					ActiveDirectoryAuth: to.Ptr(armpostgresqlflexibleservers.MicrosoftEntraAuthEnabled),
		// 					PasswordAuth: to.Ptr(armpostgresqlflexibleservers.PasswordBasedAuthEnabled),
		// 					TenantID: to.Ptr("tttttt-tttt-tttt-tttt-tttttttttttt"),
		// 				},
		// 				AvailabilityZone: to.Ptr("1"),
		// 				Backup: &armpostgresqlflexibleservers.Backup{
		// 					BackupRetentionDays: to.Ptr[int32](7),
		// 					EarliestRestoreDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-06-01T19:35:22.123Z"); return t}()),
		// 					GeoRedundantBackup: to.Ptr(armpostgresqlflexibleservers.GeographicallyRedundantBackupDisabled),
		// 				},
		// 				DataEncryption: &armpostgresqlflexibleservers.DataEncryption{
		// 					Type: to.Ptr(armpostgresqlflexibleservers.DataEncryptionTypeAzureKeyVault),
		// 					PrimaryEncryptionKeyStatus: to.Ptr(armpostgresqlflexibleservers.EncryptionKeyStatusValid),
		// 					PrimaryKeyURI: to.Ptr("https://examplekeyvault.vault.azure.net/keys/examplekey/xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"),
		// 					PrimaryUserAssignedIdentityID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/exampleresourcegroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/exampleprimaryidentity"),
		// 				},
		// 				FullyQualifiedDomainName: to.Ptr("exampleserver2.postgres.database.azure.com"),
		// 				HighAvailability: &armpostgresqlflexibleservers.HighAvailability{
		// 					Mode: to.Ptr(armpostgresqlflexibleservers.HighAvailabilityMode("Disabled")),
		// 					State: to.Ptr(armpostgresqlflexibleservers.HighAvailabilityStateNotEnabled),
		// 				},
		// 				MaintenanceWindow: &armpostgresqlflexibleservers.MaintenanceWindow{
		// 					CustomWindow: to.Ptr("Disabled"),
		// 					DayOfWeek: to.Ptr[int32](0),
		// 					StartHour: to.Ptr[int32](0),
		// 					StartMinute: to.Ptr[int32](0),
		// 				},
		// 				MinorVersion: to.Ptr("5"),
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
		// 					Tier: to.Ptr(armpostgresqlflexibleservers.AzureManagedDiskPerformanceTierP10),
		// 				},
		// 				Version: to.Ptr(armpostgresqlflexibleservers.PostgresMajorVersionSeventeen),
		// 			},
		// 			SKU: &armpostgresqlflexibleservers.SKU{
		// 				Name: to.Ptr("Standard_D4ds_v5"),
		// 				Tier: to.Ptr(armpostgresqlflexibleservers.SKUTierGeneralPurpose),
		// 			},
		// 	}},
		// }
	}
}
