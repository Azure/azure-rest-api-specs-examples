package armpostgresqlflexibleservers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresqlflexibleservers/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e96c24570a484cff13d153fb472f812878866a39/specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2025-08-01/examples/ServersCreateWithMicrosoftEntraEnabledInYourOwnVirtualNetworkWithoutHighAvailability.json
func ExampleServersClient_BeginCreateOrUpdate_createANewServerWithMicrosoftEntraAuthenticationEnabledInYourOwnVirtualNetworkAndWithoutHighAvailability() {
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
			AdministratorLogin:         to.Ptr("exampleadministratorlogin"),
			AdministratorLoginPassword: to.Ptr("examplepassword"),
			AuthConfig: &armpostgresqlflexibleservers.AuthConfig{
				ActiveDirectoryAuth: to.Ptr(armpostgresqlflexibleservers.MicrosoftEntraAuthEnabled),
				PasswordAuth:        to.Ptr(armpostgresqlflexibleservers.PasswordBasedAuthEnabled),
				TenantID:            to.Ptr("tttttt-tttt-tttt-tttt-tttttttttttt"),
			},
			AvailabilityZone: to.Ptr("1"),
			Backup: &armpostgresqlflexibleservers.Backup{
				BackupRetentionDays: to.Ptr[int32](7),
				GeoRedundantBackup:  to.Ptr(armpostgresqlflexibleservers.GeographicallyRedundantBackupDisabled),
			},
			CreateMode: to.Ptr(armpostgresqlflexibleservers.CreateModeCreate),
			DataEncryption: &armpostgresqlflexibleservers.DataEncryption{
				Type: to.Ptr(armpostgresqlflexibleservers.DataEncryptionTypeSystemManaged),
			},
			HighAvailability: &armpostgresqlflexibleservers.HighAvailability{
				Mode: to.Ptr(armpostgresqlflexibleservers.HighAvailabilityMode("Disabled")),
			},
			Network: &armpostgresqlflexibleservers.Network{
				DelegatedSubnetResourceID:   to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/exampleresourcegroup/providers/Microsoft.Network/virtualNetworks/examplevirtualnetwork/subnets/examplesubnet"),
				PrivateDNSZoneArmResourceID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourcegroups/exampleresourcegroup/providers/Microsoft.Network/privateDnsZones/exampleprivatednszone.postgres.database.azure.com"),
			},
			Storage: &armpostgresqlflexibleservers.Storage{
				AutoGrow:      to.Ptr(armpostgresqlflexibleservers.StorageAutoGrowDisabled),
				StorageSizeGB: to.Ptr[int32](512),
				Tier:          to.Ptr(armpostgresqlflexibleservers.AzureManagedDiskPerformanceTierP20),
			},
			Version: to.Ptr(armpostgresqlflexibleservers.PostgresMajorVersionSeventeen),
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
