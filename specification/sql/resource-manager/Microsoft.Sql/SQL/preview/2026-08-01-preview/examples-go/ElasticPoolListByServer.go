package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: 2026-08-01-preview/ElasticPoolListByServer.json
func ExampleElasticPoolsClient_NewListByServerPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewElasticPoolsClient().NewListByServerPager("sqlcrudtest-2369", "sqlcrudtest-8069", nil)
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
		// page = armsql.ElasticPoolsClientListByServerResponse{
		// 	ElasticPoolListResult: armsql.ElasticPoolListResult{
		// 		Value: []*armsql.ElasticPool{
		// 			{
		// 				Name: to.Ptr("sqlcrudtest-2729"),
		// 				Type: to.Ptr("Microsoft.Sql/servers/elasticPools"),
		// 				ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/sqlcrudtest-2369/providers/Microsoft.Sql/servers/sqlcrudtest-8069/elasticPools/sqlcrudtest-2729"),
		// 				Location: to.Ptr("Japan East"),
		// 				Properties: &armsql.ElasticPoolProperties{
		// 					CreationDate: to.Ptr(time.Date(2017, time.February, 10, 1, 27, 21, 320000000, time.UTC)),
		// 					LicenseType: to.Ptr(armsql.ElasticPoolLicenseTypeLicenseIncluded),
		// 					MaxSizeBytes: to.Ptr[int64](5242880000),
		// 					PerDatabaseSettings: &armsql.ElasticPoolPerDatabaseSettings{
		// 						MaxCapacity: to.Ptr[float64](1),
		// 						MinCapacity: to.Ptr[float64](0.25),
		// 					},
		// 					State: to.Ptr(armsql.ElasticPoolStateReady),
		// 					ZoneRedundant: to.Ptr(true),
		// 				},
		// 				SKU: &armsql.SKU{
		// 					Name: to.Ptr("GP_Gen4_2"),
		// 					Capacity: to.Ptr[int32](2),
		// 					Tier: to.Ptr("GeneralPurpose"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("sqlcrudtest-3191"),
		// 				Type: to.Ptr("Microsoft.Sql/servers/elasticPools"),
		// 				ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/sqlcrudtest-2369/providers/Microsoft.Sql/servers/sqlcrudtest-8069/elasticPools/sqlcrudtest-3191"),
		// 				Location: to.Ptr("Japan East"),
		// 				Properties: &armsql.ElasticPoolProperties{
		// 					CreationDate: to.Ptr(time.Date(2017, time.February, 10, 1, 26, 26, 450000000, time.UTC)),
		// 					MaxSizeBytes: to.Ptr[int64](5242880000),
		// 					PerDatabaseSettings: &armsql.ElasticPoolPerDatabaseSettings{
		// 						MaxCapacity: to.Ptr[float64](5),
		// 						MinCapacity: to.Ptr[float64](0),
		// 					},
		// 					State: to.Ptr(armsql.ElasticPoolStateReady),
		// 				},
		// 				SKU: &armsql.SKU{
		// 					Name: to.Ptr("BasicPool"),
		// 					Capacity: to.Ptr[int32](50),
		// 					Tier: to.Ptr("Basic"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("sqlcrudtest-8102"),
		// 				Type: to.Ptr("Microsoft.Sql/servers/elasticPools"),
		// 				ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/sqlcrudtest-2369/providers/Microsoft.Sql/servers/sqlcrudtest-8069/elasticPools/sqlcrudtest-8102"),
		// 				Location: to.Ptr("Japan East"),
		// 				Properties: &armsql.ElasticPoolProperties{
		// 					CreationDate: to.Ptr(time.Date(2017, time.February, 10, 1, 25, 25, 33000000, time.UTC)),
		// 					MaxSizeBytes: to.Ptr[int64](5242880000),
		// 					PerDatabaseSettings: &armsql.ElasticPoolPerDatabaseSettings{
		// 						MaxCapacity: to.Ptr[float64](5),
		// 						MinCapacity: to.Ptr[float64](0),
		// 					},
		// 					State: to.Ptr(armsql.ElasticPoolStateReady),
		// 				},
		// 				SKU: &armsql.SKU{
		// 					Name: to.Ptr("BasicPool"),
		// 					Capacity: to.Ptr[int32](50),
		// 					Tier: to.Ptr("Basic"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
