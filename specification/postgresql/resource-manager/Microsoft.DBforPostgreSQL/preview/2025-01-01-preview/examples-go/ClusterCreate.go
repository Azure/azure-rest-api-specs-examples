package armpostgresqlflexibleservers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresqlflexibleservers/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b1f4d539964453ce8008e4b069e59885e12ba441/specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2025-01-01-preview/examples/ClusterCreate.json
func ExampleServersClient_BeginCreate_clusterCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpostgresqlflexibleservers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewServersClient().BeginCreate(ctx, "testrg", "pgtestcluster", armpostgresqlflexibleservers.Server{
		Location: to.Ptr("westus"),
		Properties: &armpostgresqlflexibleservers.ServerProperties{
			AdministratorLogin:         to.Ptr("login"),
			AdministratorLoginPassword: to.Ptr("Password1"),
			Backup: &armpostgresqlflexibleservers.Backup{
				BackupRetentionDays: to.Ptr[int32](7),
				GeoRedundantBackup:  to.Ptr(armpostgresqlflexibleservers.GeoRedundantBackupEnumDisabled),
			},
			Cluster: &armpostgresqlflexibleservers.Cluster{
				ClusterSize: to.Ptr[int32](2),
			},
			CreateMode: to.Ptr(armpostgresqlflexibleservers.CreateModeCreate),
			HighAvailability: &armpostgresqlflexibleservers.HighAvailability{
				Mode: to.Ptr(armpostgresqlflexibleservers.HighAvailabilityModeDisabled),
			},
			Network: &armpostgresqlflexibleservers.Network{
				PublicNetworkAccess: to.Ptr(armpostgresqlflexibleservers.ServerPublicNetworkAccessStateDisabled),
			},
			Storage: &armpostgresqlflexibleservers.Storage{
				AutoGrow:      to.Ptr(armpostgresqlflexibleservers.StorageAutoGrowDisabled),
				StorageSizeGB: to.Ptr[int32](256),
				Tier:          to.Ptr(armpostgresqlflexibleservers.AzureManagedDiskPerformanceTiersP15),
			},
			Version: to.Ptr(armpostgresqlflexibleservers.ServerVersionSixteen),
		},
		SKU: &armpostgresqlflexibleservers.SKU{
			Name: to.Ptr("Standard_D4ds_v5"),
			Tier: to.Ptr(armpostgresqlflexibleservers.SKUTierGeneralPurpose),
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
	// 	Name: to.Ptr("pgtestcluster"),
	// 	Type: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers"),
	// 	ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforPostgreSQL/flexibleServers/pgtestcluster"),
	// 	SystemData: &armpostgresqlflexibleservers.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-10-10T18:31:50.730Z"); return t}()),
	// 	},
	// 	Location: to.Ptr("westus"),
	// 	Properties: &armpostgresqlflexibleservers.ServerProperties{
	// 		AdministratorLogin: to.Ptr("login"),
	// 		Backup: &armpostgresqlflexibleservers.Backup{
	// 			BackupRetentionDays: to.Ptr[int32](7),
	// 			EarliestRestoreDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-10-10T18:40:39.045Z"); return t}()),
	// 			GeoRedundantBackup: to.Ptr(armpostgresqlflexibleservers.GeoRedundantBackupEnumDisabled),
	// 		},
	// 		Cluster: &armpostgresqlflexibleservers.Cluster{
	// 			ClusterSize: to.Ptr[int32](2),
	// 		},
	// 		DataEncryption: &armpostgresqlflexibleservers.DataEncryption{
	// 			Type: to.Ptr(armpostgresqlflexibleservers.ArmServerKeyTypeSystemManaged),
	// 		},
	// 		HighAvailability: &armpostgresqlflexibleservers.HighAvailability{
	// 			Mode: to.Ptr(armpostgresqlflexibleservers.HighAvailabilityModeDisabled),
	// 		},
	// 		MaintenanceWindow: &armpostgresqlflexibleservers.MaintenanceWindow{
	// 			CustomWindow: to.Ptr("Disabled"),
	// 			DayOfWeek: to.Ptr[int32](0),
	// 			StartHour: to.Ptr[int32](0),
	// 			StartMinute: to.Ptr[int32](0),
	// 		},
	// 		Network: &armpostgresqlflexibleservers.Network{
	// 			PublicNetworkAccess: to.Ptr(armpostgresqlflexibleservers.ServerPublicNetworkAccessStateDisabled),
	// 		},
	// 		PrivateEndpointConnections: []*armpostgresqlflexibleservers.PrivateEndpointConnection{
	// 		},
	// 		Replica: &armpostgresqlflexibleservers.Replica{
	// 			ReplicationState: to.Ptr(armpostgresqlflexibleservers.ReplicationStateActive),
	// 			Role: to.Ptr(armpostgresqlflexibleservers.ReplicationRolePrimary),
	// 		},
	// 		ReplicationRole: to.Ptr(armpostgresqlflexibleservers.ReplicationRolePrimary),
	// 		State: to.Ptr(armpostgresqlflexibleservers.ServerStateReady),
	// 		Storage: &armpostgresqlflexibleservers.Storage{
	// 			AutoGrow: to.Ptr(armpostgresqlflexibleservers.StorageAutoGrowDisabled),
	// 			Iops: to.Ptr[int32](1100),
	// 			StorageSizeGB: to.Ptr[int32](256),
	// 			Tier: to.Ptr(armpostgresqlflexibleservers.AzureManagedDiskPerformanceTiersP15),
	// 		},
	// 		Version: to.Ptr(armpostgresqlflexibleservers.ServerVersionSixteen),
	// 	},
	// 	SKU: &armpostgresqlflexibleservers.SKU{
	// 		Name: to.Ptr("Standard_D4ds_v5"),
	// 		Tier: to.Ptr(armpostgresqlflexibleservers.SKUTierGeneralPurpose),
	// 	},
	// }
}
