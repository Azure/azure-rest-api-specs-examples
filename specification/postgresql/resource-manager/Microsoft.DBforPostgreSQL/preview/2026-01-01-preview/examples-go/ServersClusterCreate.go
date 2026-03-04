package armpostgresqlflexibleservers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresqlflexibleservers/v6"
)

// Generated from example definition: 2026-01-01-preview/ServersClusterCreate.json
func ExampleServersClient_BeginCreateOrUpdate_createANewElasticCluster() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpostgresqlflexibleservers.NewClientFactory("ffffffff-ffff-ffff-ffff-ffffffffffff", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewServersClient().BeginCreateOrUpdate(ctx, "exampleresourcegroup", "exampleserver", armpostgresqlflexibleservers.Server{
		Location: to.Ptr("eastus"),
		Properties: &armpostgresqlflexibleservers.ServerProperties{
			AdministratorLogin:         to.Ptr("examplelogin"),
			AdministratorLoginPassword: to.Ptr("examplepassword"),
			Backup: &armpostgresqlflexibleservers.Backup{
				BackupRetentionDays: to.Ptr[int32](7),
				GeoRedundantBackup:  to.Ptr(armpostgresqlflexibleservers.GeographicallyRedundantBackupDisabled),
			},
			Cluster: &armpostgresqlflexibleservers.Cluster{
				ClusterSize:         to.Ptr[int32](2),
				DefaultDatabaseName: to.Ptr("clusterdb"),
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
				Tier:          to.Ptr(armpostgresqlflexibleservers.AzureManagedDiskPerformanceTierP15),
			},
			Version: to.Ptr(armpostgresqlflexibleservers.PostgresMajorVersionSixteen),
		},
		SKU: &armpostgresqlflexibleservers.SKU{
			Name: to.Ptr("Standard_D4ds_v5"),
			Tier: to.Ptr(armpostgresqlflexibleservers.SKUTierGeneralPurpose),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
