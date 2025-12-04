package armpostgresqlflexibleservers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresqlflexibleservers/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e96c24570a484cff13d153fb472f812878866a39/specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2025-08-01/examples/ServersUpdate.json
func ExampleServersClient_BeginUpdate_updateAnExistingServer() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpostgresqlflexibleservers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewServersClient().BeginUpdate(ctx, "exampleresourcegroup", "exampleserver", armpostgresqlflexibleservers.ServerForPatch{
		Properties: &armpostgresqlflexibleservers.ServerPropertiesForPatch{
			AdministratorLoginPassword: to.Ptr("examplenewpassword"),
			Backup: &armpostgresqlflexibleservers.BackupForPatch{
				BackupRetentionDays: to.Ptr[int32](20),
			},
			CreateMode: to.Ptr(armpostgresqlflexibleservers.CreateModeForPatchUpdate),
			Storage: &armpostgresqlflexibleservers.Storage{
				AutoGrow:      to.Ptr(armpostgresqlflexibleservers.StorageAutoGrowEnabled),
				StorageSizeGB: to.Ptr[int32](1024),
				Tier:          to.Ptr(armpostgresqlflexibleservers.AzureManagedDiskPerformanceTierP30),
			},
		},
		SKU: &armpostgresqlflexibleservers.SKUForPatch{
			Name: to.Ptr("Standard_D8s_v3"),
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
