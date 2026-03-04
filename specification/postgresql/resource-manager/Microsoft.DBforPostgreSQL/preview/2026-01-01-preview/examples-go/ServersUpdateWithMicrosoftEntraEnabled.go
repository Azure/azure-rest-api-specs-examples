package armpostgresqlflexibleservers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresqlflexibleservers/v6"
)

// Generated from example definition: 2026-01-01-preview/ServersUpdateWithMicrosoftEntraEnabled.json
func ExampleServersClient_BeginUpdate_updateAnExistingServerWithMicrosoftEntraAuthenticationEnabled() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpostgresqlflexibleservers.NewClientFactory("ffffffff-ffff-ffff-ffff-ffffffffffff", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewServersClient().BeginUpdate(ctx, "exampleresourcegroup", "exampleserver", armpostgresqlflexibleservers.ServerForPatch{
		Properties: &armpostgresqlflexibleservers.ServerPropertiesForPatch{
			AdministratorLoginPassword: to.Ptr("examplenewpassword"),
			AuthConfig: &armpostgresqlflexibleservers.AuthConfigForPatch{
				ActiveDirectoryAuth: to.Ptr(armpostgresqlflexibleservers.MicrosoftEntraAuthEnabled),
				PasswordAuth:        to.Ptr(armpostgresqlflexibleservers.PasswordBasedAuthEnabled),
				TenantID:            to.Ptr("tttttt-tttt-tttt-tttt-tttttttttttt"),
			},
			Backup: &armpostgresqlflexibleservers.BackupForPatch{
				BackupRetentionDays: to.Ptr[int32](20),
			},
			CreateMode: to.Ptr(armpostgresqlflexibleservers.CreateModeForPatchUpdate),
			Storage: &armpostgresqlflexibleservers.Storage{
				AutoGrow:      to.Ptr(armpostgresqlflexibleservers.StorageAutoGrowDisabled),
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
