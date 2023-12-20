package armpostgresqlflexibleservers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresqlflexibleservers/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/90115af9fda46f323e5c42c274f2b376108d1d47/specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2023-06-01-preview/examples/ServerGet.json
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
	res, err := clientFactory.NewServersClient().Get(ctx, "testrg", "pgtestsvc1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Server = armpostgresqlflexibleservers.Server{
	// 	Name: to.Ptr("pgtestsvc1"),
	// 	Type: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers"),
	// 	ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforPostgreSQL/flexibleServers/pgtestsvc1"),
	// 	Location: to.Ptr("westus"),
	// 	Tags: map[string]*string{
	// 		"ElasticServer": to.Ptr("1"),
	// 	},
	// 	Properties: &armpostgresqlflexibleservers.ServerProperties{
	// 		AdministratorLogin: to.Ptr("cloudsa"),
	// 		AuthConfig: &armpostgresqlflexibleservers.AuthConfig{
	// 			ActiveDirectoryAuth: to.Ptr(armpostgresqlflexibleservers.ActiveDirectoryAuthEnumDisabled),
	// 			PasswordAuth: to.Ptr(armpostgresqlflexibleservers.PasswordAuthEnumEnabled),
	// 		},
	// 		AvailabilityZone: to.Ptr("1"),
	// 		Backup: &armpostgresqlflexibleservers.Backup{
	// 			BackupRetentionDays: to.Ptr[int32](7),
	// 			EarliestRestoreDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-05-26T23:07:25.816Z"); return t}()),
	// 			GeoRedundantBackup: to.Ptr(armpostgresqlflexibleservers.GeoRedundantBackupEnumDisabled),
	// 		},
	// 		FullyQualifiedDomainName: to.Ptr("pgtestsvc1.postgres.database.azure.com"),
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
	// 		MinorVersion: to.Ptr("6"),
	// 		Network: &armpostgresqlflexibleservers.Network{
	// 			PublicNetworkAccess: to.Ptr(armpostgresqlflexibleservers.ServerPublicNetworkAccessStateEnabled),
	// 		},
	// 		State: to.Ptr(armpostgresqlflexibleservers.ServerStateReady),
	// 		Storage: &armpostgresqlflexibleservers.Storage{
	// 			AutoGrow: to.Ptr(armpostgresqlflexibleservers.StorageAutoGrowDisabled),
	// 			Iops: to.Ptr[int32](2300),
	// 			StorageSizeGB: to.Ptr[int32](512),
	// 			Tier: to.Ptr(armpostgresqlflexibleservers.AzureManagedDiskPerformanceTiersP20),
	// 		},
	// 		Version: to.Ptr(armpostgresqlflexibleservers.ServerVersionTwelve),
	// 	},
	// 	SKU: &armpostgresqlflexibleservers.SKU{
	// 		Name: to.Ptr("Standard_D4s_v3"),
	// 		Tier: to.Ptr(armpostgresqlflexibleservers.SKUTierGeneralPurpose),
	// 	},
	// }
}
