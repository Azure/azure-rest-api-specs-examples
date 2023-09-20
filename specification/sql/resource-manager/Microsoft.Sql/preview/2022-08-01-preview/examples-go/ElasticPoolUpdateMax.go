package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c78b5d8bd3aff2d82a5f034d9164b1a9ac030e09/specification/sql/resource-manager/Microsoft.Sql/preview/2022-08-01-preview/examples/ElasticPoolUpdateMax.json
func ExampleElasticPoolsClient_BeginUpdate_updateAnElasticPoolWithAllParameter() {
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
			LicenseType: to.Ptr(armsql.ElasticPoolLicenseTypeLicenseIncluded),
			PerDatabaseSettings: &armsql.ElasticPoolPerDatabaseSettings{
				MaxCapacity: to.Ptr[float64](1),
				MinCapacity: to.Ptr[float64](0.25),
			},
			ZoneRedundant: to.Ptr(true),
		},
		SKU: &armsql.SKU{
			Name:     to.Ptr("BC_Gen4"),
			Capacity: to.Ptr[int32](2),
			Tier:     to.Ptr("BusinessCritical"),
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
	// 		CreationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-02-10T01:27:21.32Z"); return t}()),
	// 		LicenseType: to.Ptr(armsql.ElasticPoolLicenseTypeLicenseIncluded),
	// 		MaxSizeBytes: to.Ptr[int64](5242880000),
	// 		PerDatabaseSettings: &armsql.ElasticPoolPerDatabaseSettings{
	// 			MaxCapacity: to.Ptr[float64](1),
	// 			MinCapacity: to.Ptr[float64](0.25),
	// 		},
	// 		State: to.Ptr(armsql.ElasticPoolStateReady),
	// 		ZoneRedundant: to.Ptr(true),
	// 	},
	// 	SKU: &armsql.SKU{
	// 		Name: to.Ptr("BC_Gen4"),
	// 		Capacity: to.Ptr[int32](2),
	// 		Tier: to.Ptr("BusinessCritical"),
	// 	},
	// }
}
