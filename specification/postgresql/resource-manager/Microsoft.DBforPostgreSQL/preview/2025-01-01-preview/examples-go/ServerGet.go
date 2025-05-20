package armpostgresqlflexibleservers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresqlflexibleservers/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b1f4d539964453ce8008e4b069e59885e12ba441/specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2025-01-01-preview/examples/ServerGet.json
func ExampleServersClient_Get_serverGet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpostgresqlflexibleservers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewServersClient().Get(ctx, "testrg", "testpgflex", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Server = armpostgresqlflexibleservers.Server{
	// 	Name: to.Ptr("testpgflex"),
	// 	Type: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers"),
	// 	ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforPostgreSQL/flexibleServers/testpgflex"),
	// 	SystemData: &armpostgresqlflexibleservers.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-12-02T21:26:40.243Z"); return t}()),
	// 	},
	// 	Location: to.Ptr("East US"),
	// 	Properties: &armpostgresqlflexibleservers.ServerProperties{
	// 		AdministratorLogin: to.Ptr("login"),
	// 		AuthConfig: &armpostgresqlflexibleservers.AuthConfig{
	// 			ActiveDirectoryAuth: to.Ptr(armpostgresqlflexibleservers.ActiveDirectoryAuthEnumDisabled),
	// 			PasswordAuth: to.Ptr(armpostgresqlflexibleservers.PasswordAuthEnumEnabled),
	// 		},
	// 		AvailabilityZone: to.Ptr("1"),
	// 		Backup: &armpostgresqlflexibleservers.Backup{
	// 			BackupRetentionDays: to.Ptr[int32](7),
	// 			EarliestRestoreDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-02-25T02:11:46.457Z"); return t}()),
	// 			GeoRedundantBackup: to.Ptr(armpostgresqlflexibleservers.GeoRedundantBackupEnumEnabled),
	// 		},
	// 		DataEncryption: &armpostgresqlflexibleservers.DataEncryption{
	// 			Type: to.Ptr(armpostgresqlflexibleservers.ArmServerKeyTypeSystemManaged),
	// 		},
	// 		FullyQualifiedDomainName: to.Ptr("testpgflex.postgres.database.azure.com"),
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
	// 		MinorVersion: to.Ptr("8"),
	// 		Network: &armpostgresqlflexibleservers.Network{
	// 			PublicNetworkAccess: to.Ptr(armpostgresqlflexibleservers.ServerPublicNetworkAccessStateEnabled),
	// 		},
	// 		PrivateEndpointConnections: []*armpostgresqlflexibleservers.PrivateEndpointConnection{
	// 		},
	// 		Replica: &armpostgresqlflexibleservers.Replica{
	// 			Capacity: to.Ptr[int32](5),
	// 			Role: to.Ptr(armpostgresqlflexibleservers.ReplicationRolePrimary),
	// 		},
	// 		ReplicaCapacity: to.Ptr[int32](5),
	// 		ReplicationRole: to.Ptr(armpostgresqlflexibleservers.ReplicationRolePrimary),
	// 		State: to.Ptr(armpostgresqlflexibleservers.ServerStateReady),
	// 		Storage: &armpostgresqlflexibleservers.Storage{
	// 			Type: to.Ptr(armpostgresqlflexibleservers.StorageType("")),
	// 			AutoGrow: to.Ptr(armpostgresqlflexibleservers.StorageAutoGrowDisabled),
	// 			Iops: to.Ptr[int32](500),
	// 			StorageSizeGB: to.Ptr[int32](128),
	// 			Tier: to.Ptr(armpostgresqlflexibleservers.AzureManagedDiskPerformanceTiersP10),
	// 		},
	// 		Version: to.Ptr(armpostgresqlflexibleservers.ServerVersionSixteen),
	// 	},
	// 	SKU: &armpostgresqlflexibleservers.SKU{
	// 		Name: to.Ptr("Standard_D4ds_v5"),
	// 		Tier: to.Ptr(armpostgresqlflexibleservers.SKUTierGeneralPurpose),
	// 	},
	// }
}
