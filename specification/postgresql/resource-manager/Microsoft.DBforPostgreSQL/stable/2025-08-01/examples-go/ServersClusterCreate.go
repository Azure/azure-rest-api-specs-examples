package armpostgresqlflexibleservers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresqlflexibleservers/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e96c24570a484cff13d153fb472f812878866a39/specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2025-08-01/examples/ServersClusterCreate.json
func ExampleServersClient_BeginCreateOrUpdate_createANewElasticCluster() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpostgresqlflexibleservers.NewClientFactory("<subscription-id>", cred, nil)
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
				Mode: to.Ptr(armpostgresqlflexibleservers.HighAvailabilityMode("Disabled")),
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
