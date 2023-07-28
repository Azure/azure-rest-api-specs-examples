package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/fe78d8f1e7bd86c778c7e1cafd52cb0e9fec67ef/specification/sql/resource-manager/Microsoft.Sql/preview/2022-08-01-preview/examples/ElasticPoolGetWithPreferredEnclaveType.json
func ExampleElasticPoolsClient_Get_getAnElasticPoolWithPreferredEnclaveTypeParameter() {
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
	// 	Kind: to.Ptr("vcore,pool"),
	// 	Properties: &armsql.ElasticPoolProperties{
	// 		CreationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-08-26T03:46:20.57Z"); return t}()),
	// 		HighAvailabilityReplicaCount: to.Ptr[int32](2),
	// 		LicenseType: to.Ptr(armsql.ElasticPoolLicenseTypeLicenseIncluded),
	// 		MaintenanceConfigurationID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/providers/Microsoft.Maintenance/publicMaintenanceConfigurations/SQL_Default"),
	// 		MaxSizeBytes: to.Ptr[int64](0),
	// 		PerDatabaseSettings: &armsql.ElasticPoolPerDatabaseSettings{
	// 			MaxCapacity: to.Ptr[float64](4),
	// 			MinCapacity: to.Ptr[float64](0),
	// 		},
	// 		PreferredEnclaveType: to.Ptr(armsql.AlwaysEncryptedEnclaveTypeVBS),
	// 		State: to.Ptr(armsql.ElasticPoolStateReady),
	// 		ZoneRedundant: to.Ptr(false),
	// 	},
	// 	SKU: &armsql.SKU{
	// 		Name: to.Ptr("GP_Gen5"),
	// 		Capacity: to.Ptr[int32](4),
	// 		Family: to.Ptr("Gen5"),
	// 		Tier: to.Ptr("GeneralPurpose"),
	// 	},
	// }
}
