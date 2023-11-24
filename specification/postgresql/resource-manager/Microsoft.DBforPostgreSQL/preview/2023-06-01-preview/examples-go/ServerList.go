package armpostgresqlflexibleservers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresqlflexibleservers/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/bf204aab860f2eb58a9d346b00d44760f2a9b0a2/specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2023-06-01-preview/examples/ServerList.json
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
		// 			Name: to.Ptr("pgtestsvc4"),
		// 			Type: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforPostgreSQL/flexibleServers/pgtestsvc4"),
		// 			Location: to.Ptr("westus"),
		// 			Tags: map[string]*string{
		// 				"ElasticServer": to.Ptr("1"),
		// 			},
		// 			Properties: &armpostgresqlflexibleservers.ServerProperties{
		// 				AdministratorLogin: to.Ptr("cloudsa"),
		// 				AuthConfig: &armpostgresqlflexibleservers.AuthConfig{
		// 					ActiveDirectoryAuth: to.Ptr(armpostgresqlflexibleservers.ActiveDirectoryAuthEnumDisabled),
		// 					PasswordAuth: to.Ptr(armpostgresqlflexibleservers.PasswordAuthEnumEnabled),
		// 				},
		// 				AvailabilityZone: to.Ptr("1"),
		// 				Backup: &armpostgresqlflexibleservers.Backup{
		// 					BackupRetentionDays: to.Ptr[int32](7),
		// 					EarliestRestoreDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-05-26T01:16:58.372Z"); return t}()),
		// 					GeoRedundantBackup: to.Ptr(armpostgresqlflexibleservers.GeoRedundantBackupEnumDisabled),
		// 				},
		// 				FullyQualifiedDomainName: to.Ptr("c7d7483a8ceb.test-private-dns-zone.postgres.database.azure.com"),
		// 				HighAvailability: &armpostgresqlflexibleservers.HighAvailability{
		// 					Mode: to.Ptr(armpostgresqlflexibleservers.HighAvailabilityModeZoneRedundant),
		// 					StandbyAvailabilityZone: to.Ptr("2"),
		// 					State: to.Ptr(armpostgresqlflexibleservers.ServerHAStateHealthy),
		// 				},
		// 				MinorVersion: to.Ptr("6"),
		// 				Network: &armpostgresqlflexibleservers.Network{
		// 					DelegatedSubnetResourceID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.Network/virtualNetworks/test-vnet/subnets/test-vnet-subnet"),
		// 					PrivateDNSZoneArmResourceID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.Network/privateDnsZones/test-private-dns-zone.postgres.database.azure.com"),
		// 					PublicNetworkAccess: to.Ptr(armpostgresqlflexibleservers.ServerPublicNetworkAccessStateDisabled),
		// 				},
		// 				State: to.Ptr(armpostgresqlflexibleservers.ServerStateReady),
		// 				Storage: &armpostgresqlflexibleservers.Storage{
		// 					AutoGrow: to.Ptr(armpostgresqlflexibleservers.StorageAutoGrowDisabled),
		// 					StorageSizeGB: to.Ptr[int32](512),
		// 				},
		// 				Version: to.Ptr(armpostgresqlflexibleservers.ServerVersionTwelve),
		// 			},
		// 			SKU: &armpostgresqlflexibleservers.SKU{
		// 				Name: to.Ptr("Standard_D4s_v3"),
		// 				Tier: to.Ptr(armpostgresqlflexibleservers.SKUTierGeneralPurpose),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("pgtestsvc1"),
		// 			Type: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforPostgreSQL/flexibleServers/pgtestsvc1"),
		// 			Location: to.Ptr("eastus"),
		// 			Tags: map[string]*string{
		// 				"ElasticServer": to.Ptr("1"),
		// 			},
		// 			Properties: &armpostgresqlflexibleservers.ServerProperties{
		// 				AdministratorLogin: to.Ptr("cloudsa"),
		// 				AuthConfig: &armpostgresqlflexibleservers.AuthConfig{
		// 					ActiveDirectoryAuth: to.Ptr(armpostgresqlflexibleservers.ActiveDirectoryAuthEnumDisabled),
		// 					PasswordAuth: to.Ptr(armpostgresqlflexibleservers.PasswordAuthEnumEnabled),
		// 				},
		// 				AvailabilityZone: to.Ptr("1"),
		// 				Backup: &armpostgresqlflexibleservers.Backup{
		// 					BackupRetentionDays: to.Ptr[int32](7),
		// 					EarliestRestoreDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-05-26T23:15:38.813Z"); return t}()),
		// 					GeoRedundantBackup: to.Ptr(armpostgresqlflexibleservers.GeoRedundantBackupEnumDisabled),
		// 				},
		// 				FullyQualifiedDomainName: to.Ptr("pgtestsvc1.postgres.database.azure.com"),
		// 				HighAvailability: &armpostgresqlflexibleservers.HighAvailability{
		// 					Mode: to.Ptr(armpostgresqlflexibleservers.HighAvailabilityModeDisabled),
		// 					State: to.Ptr(armpostgresqlflexibleservers.ServerHAStateNotEnabled),
		// 				},
		// 				MinorVersion: to.Ptr("6"),
		// 				Network: &armpostgresqlflexibleservers.Network{
		// 					PublicNetworkAccess: to.Ptr(armpostgresqlflexibleservers.ServerPublicNetworkAccessStateEnabled),
		// 				},
		// 				State: to.Ptr(armpostgresqlflexibleservers.ServerStateReady),
		// 				Storage: &armpostgresqlflexibleservers.Storage{
		// 					AutoGrow: to.Ptr(armpostgresqlflexibleservers.StorageAutoGrowEnabled),
		// 					Iops: to.Ptr[int32](2300),
		// 					StorageSizeGB: to.Ptr[int32](512),
		// 					Tier: to.Ptr(armpostgresqlflexibleservers.AzureManagedDiskPerformanceTiersP20),
		// 				},
		// 				Version: to.Ptr(armpostgresqlflexibleservers.ServerVersionTwelve),
		// 			},
		// 			SKU: &armpostgresqlflexibleservers.SKU{
		// 				Name: to.Ptr("Standard_D4s_v3"),
		// 				Tier: to.Ptr(armpostgresqlflexibleservers.SKUTierGeneralPurpose),
		// 			},
		// 	}},
		// }
	}
}
