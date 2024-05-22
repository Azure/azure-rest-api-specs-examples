package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/33c4457b1d13f83965f4fe3367dca4a6df898100/specification/sql/resource-manager/Microsoft.Sql/preview/2022-08-01-preview/examples/ElasticPoolUpdateAssignMaintenanceConfiguration.json
func ExampleElasticPoolsClient_BeginUpdate_assignsMaintenanceConfigurationToAnElasticPool() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewElasticPoolsClient().BeginUpdate(ctx, "sqlcrudtest-2369", "sqlcrudtest-8069", "sqlcrudtest-8102", armsql.ElasticPoolUpdate{
		Properties: &armsql.ElasticPoolUpdateProperties{
			MaintenanceConfigurationID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/providers/Microsoft.Maintenance/publicMaintenanceConfigurations/SQL_JapanEast_1"),
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
	// res.ElasticPool = armsql.ElasticPool{
	// 	Name: to.Ptr("sqlcrudtest-8102"),
	// 	Type: to.Ptr("Microsoft.Sql/servers/elasticPools"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/sqlcrudtest-2369/providers/Microsoft.Sql/servers/sqlcrudtest-8069/elasticPools/sqlcrudtest-8102"),
	// 	Location: to.Ptr("Japan East"),
	// 	Properties: &armsql.ElasticPoolProperties{
	// 		CreationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-02-10T01:25:25.033Z"); return t}()),
	// 		MaintenanceConfigurationID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/providers/Microsoft.Maintenance/publicMaintenanceConfigurations/SQL_JapanEast_1"),
	// 		MaxSizeBytes: to.Ptr[int64](5242880000),
	// 		PerDatabaseSettings: &armsql.ElasticPoolPerDatabaseSettings{
	// 			MaxCapacity: to.Ptr[float64](5),
	// 			MinCapacity: to.Ptr[float64](0),
	// 		},
	// 		State: to.Ptr(armsql.ElasticPoolStateReady),
	// 	},
	// 	SKU: &armsql.SKU{
	// 		Name: to.Ptr("BasicPool"),
	// 		Capacity: to.Ptr[int32](50),
	// 		Tier: to.Ptr("Basic"),
	// 	},
	// }
}
