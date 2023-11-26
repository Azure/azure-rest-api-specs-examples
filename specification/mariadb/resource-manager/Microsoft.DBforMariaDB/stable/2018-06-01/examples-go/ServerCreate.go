package armmariadb_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mariadb/armmariadb"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/mariadb/resource-manager/Microsoft.DBforMariaDB/stable/2018-06-01/examples/ServerCreate.json
func ExampleServersClient_BeginCreate_createANewServer() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmariadb.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewServersClient().BeginCreate(ctx, "testrg", "mariadbtestsvc4", armmariadb.ServerForCreate{
		Location: to.Ptr("westus"),
		Properties: &armmariadb.ServerPropertiesForDefaultCreate{
			CreateMode:        to.Ptr(armmariadb.CreateModeDefault),
			MinimalTLSVersion: to.Ptr(armmariadb.MinimalTLSVersionEnumTLS12),
			SSLEnforcement:    to.Ptr(armmariadb.SSLEnforcementEnumEnabled),
			StorageProfile: &armmariadb.StorageProfile{
				BackupRetentionDays: to.Ptr[int32](7),
				GeoRedundantBackup:  to.Ptr(armmariadb.GeoRedundantBackupEnabled),
				StorageMB:           to.Ptr[int32](128000),
			},
			AdministratorLogin:         to.Ptr("cloudsa"),
			AdministratorLoginPassword: to.Ptr("<administratorLoginPassword>"),
		},
		SKU: &armmariadb.SKU{
			Name:     to.Ptr("GP_Gen5_2"),
			Capacity: to.Ptr[int32](2),
			Family:   to.Ptr("Gen5"),
			Tier:     to.Ptr(armmariadb.SKUTierGeneralPurpose),
		},
		Tags: map[string]*string{
			"ElasticServer": to.Ptr("1"),
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
	// res.Server = armmariadb.Server{
	// 	Name: to.Ptr("mariadbtestsvc4"),
	// 	Type: to.Ptr("Microsoft.DBforMariaDB/servers"),
	// 	ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforMariaDB/servers/mariadbtestsvc4"),
	// 	Location: to.Ptr("westus"),
	// 	Tags: map[string]*string{
	// 		"elasticServer": to.Ptr("1"),
	// 	},
	// 	Properties: &armmariadb.ServerProperties{
	// 		AdministratorLogin: to.Ptr("cloudsa"),
	// 		EarliestRestoreDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-03-14T18:02:41.577Z"); return t}()),
	// 		FullyQualifiedDomainName: to.Ptr("mariadbtestsvc4.mariadb.database.azure.com"),
	// 		SSLEnforcement: to.Ptr(armmariadb.SSLEnforcementEnumEnabled),
	// 		StorageProfile: &armmariadb.StorageProfile{
	// 			BackupRetentionDays: to.Ptr[int32](7),
	// 			GeoRedundantBackup: to.Ptr(armmariadb.GeoRedundantBackupEnabled),
	// 			StorageMB: to.Ptr[int32](128000),
	// 		},
	// 		UserVisibleState: to.Ptr(armmariadb.ServerStateReady),
	// 		Version: to.Ptr(armmariadb.ServerVersionTen3),
	// 	},
	// 	SKU: &armmariadb.SKU{
	// 		Name: to.Ptr("GP_Gen5_2"),
	// 		Capacity: to.Ptr[int32](2),
	// 		Family: to.Ptr("Gen5"),
	// 		Tier: to.Ptr(armmariadb.SKUTierGeneralPurpose),
	// 	},
	// }
}
