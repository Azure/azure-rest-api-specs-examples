package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a3913f4b26467aed413cdc907116e99894f08994/specification/sql/resource-manager/Microsoft.Sql/preview/2022-08-01-preview/examples/ElasticPoolGet.json
func ExampleElasticPoolsClient_Get_getAnElasticPool() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewElasticPoolsClient().Get(ctx, "sqlcrudtest-2369", "sqlcrudtest-8069", "sqlcrudtest-8102", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ElasticPool = armsql.ElasticPool{
	// 	Name: to.Ptr("sqlcrudtest-8102"),
	// 	Type: to.Ptr("Microsoft.Sql/servers/elasticPools"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/sqlcrudtest-2369/providers/Microsoft.Sql/servers/sqlcrudtest-8069/elasticPools/sqlcrudtest-8102"),
	// 	Location: to.Ptr("Japan East"),
	// 	Properties: &armsql.ElasticPoolProperties{
	// 		CreationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-10-10T01:25:25.033Z"); return t}()),
	// 		LicenseType: to.Ptr(armsql.ElasticPoolLicenseTypeLicenseIncluded),
	// 		MaintenanceConfigurationID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/providers/Microsoft.Maintenance/publicMaintenanceConfigurations/SQL_JapanEast_1"),
	// 		MaxSizeBytes: to.Ptr[int64](5242880000),
	// 		PerDatabaseSettings: &armsql.ElasticPoolPerDatabaseSettings{
	// 			MaxCapacity: to.Ptr[float64](1),
	// 			MinCapacity: to.Ptr[float64](0.25),
	// 		},
	// 		State: to.Ptr(armsql.ElasticPoolStateReady),
	// 		ZoneRedundant: to.Ptr(true),
	// 	},
	// 	SKU: &armsql.SKU{
	// 		Name: to.Ptr("GP_Gen5_2"),
	// 		Capacity: to.Ptr[int32](2),
	// 		Tier: to.Ptr("GeneralPurpose"),
	// 	},
	// }
}
