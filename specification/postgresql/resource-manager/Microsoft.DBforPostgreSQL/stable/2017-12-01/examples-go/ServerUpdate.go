package armpostgresql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresql"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c767823fdfd9d5e96bad245e3ea4d14d94a716bb/specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2017-12-01/examples/ServerUpdate.json
func ExampleServersClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpostgresql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewServersClient().BeginUpdate(ctx, "testrg", "pgtestsvc4", armpostgresql.ServerUpdateParameters{
		Properties: &armpostgresql.ServerUpdateParametersProperties{
			AdministratorLoginPassword: to.Ptr("<administratorLoginPassword>"),
			MinimalTLSVersion:          to.Ptr(armpostgresql.MinimalTLSVersionEnumTLS12),
			SSLEnforcement:             to.Ptr(armpostgresql.SSLEnforcementEnumEnabled),
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
	// res.Server = armpostgresql.Server{
	// 	Name: to.Ptr("pgtestsvc4"),
	// 	Type: to.Ptr("Microsoft.DBforPostgreSQL/servers"),
	// 	ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforPostgreSQL/servers/pgtestsvc4"),
	// 	Location: to.Ptr("westus"),
	// 	Tags: map[string]*string{
	// 		"ElasticServer": to.Ptr("1"),
	// 	},
	// 	Properties: &armpostgresql.ServerProperties{
	// 		AdministratorLogin: to.Ptr("cloudsa"),
	// 		EarliestRestoreDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-03-14T21:08:24.637Z"); return t}()),
	// 		FullyQualifiedDomainName: to.Ptr("pgtestsvc4.postgres.database.azure.com"),
	// 		MinimalTLSVersion: to.Ptr(armpostgresql.MinimalTLSVersionEnumTLS12),
	// 		SSLEnforcement: to.Ptr(armpostgresql.SSLEnforcementEnumEnabled),
	// 		StorageProfile: &armpostgresql.StorageProfile{
	// 			BackupRetentionDays: to.Ptr[int32](7),
	// 			GeoRedundantBackup: to.Ptr(armpostgresql.GeoRedundantBackupDisabled),
	// 			StorageMB: to.Ptr[int32](128000),
	// 		},
	// 		UserVisibleState: to.Ptr(armpostgresql.ServerStateReady),
	// 		Version: to.Ptr(armpostgresql.ServerVersionNine6),
	// 	},
	// 	SKU: &armpostgresql.SKU{
	// 		Name: to.Ptr("B_Gen4_2"),
	// 		Capacity: to.Ptr[int32](2),
	// 		Family: to.Ptr("Gen4"),
	// 		Tier: to.Ptr(armpostgresql.SKUTierBasic),
	// 	},
	// }
}
