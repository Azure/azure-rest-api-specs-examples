package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/33c4457b1d13f83965f4fe3367dca4a6df898100/specification/sql/resource-manager/Microsoft.Sql/preview/2023-02-01-preview/examples/ListDatabasesByElasticPool.json
func ExampleDatabasesClient_NewListByElasticPoolPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDatabasesClient().NewListByElasticPoolPager("Default-SQL-SouthEastAsia", "testsvr", "pool1", nil)
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
		// page.DatabaseListResult = armsql.DatabaseListResult{
		// 	Value: []*armsql.Database{
		// 		{
		// 			Name: to.Ptr("DB001"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/databases"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default-SQL-SouthEastAsia/providers/Microsoft.Sql/servers/testsvr/databases/DB001"),
		// 			Location: to.Ptr("southeastasia"),
		// 			Kind: to.Ptr("v12.0,user"),
		// 			Properties: &armsql.DatabaseProperties{
		// 				CatalogCollation: to.Ptr(armsql.CatalogCollationTypeSQLLatin1GeneralCP1CIAS),
		// 				Collation: to.Ptr("SQL_Latin1_General_CP1_CI_AS"),
		// 				CreationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-07-12T22:08:39.163Z"); return t}()),
		// 				CurrentServiceObjectiveName: to.Ptr("ElasticPool"),
		// 				DatabaseID: to.Ptr("bfe0735f-bc87-447f-b2c2-481f4b100614"),
		// 				DefaultSecondaryLocation: to.Ptr("North Europe"),
		// 				ElasticPoolID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default-SQL-SouthEastAsia/providers/Microsoft.Sql/servers/testsvr/elasticPools/Pool1"),
		// 				MaxSizeBytes: to.Ptr[int64](268435456000),
		// 				Status: to.Ptr(armsql.DatabaseStatusOnline),
		// 			},
		// 			SKU: &armsql.SKU{
		// 				Name: to.Ptr("ElasticPool"),
		// 				Tier: to.Ptr("Standard"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("DB002"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/databases"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default-SQL-SouthEastAsia/providers/Microsoft.Sql/servers/testsvr/databases/DB002"),
		// 			Location: to.Ptr("southeastasia"),
		// 			Kind: to.Ptr("v12.0,user"),
		// 			Properties: &armsql.DatabaseProperties{
		// 				CatalogCollation: to.Ptr(armsql.CatalogCollationTypeSQLLatin1GeneralCP1CIAS),
		// 				Collation: to.Ptr("SQL_Latin1_General_CP1_CI_AS"),
		// 				CreationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-07-12T22:10:10.773Z"); return t}()),
		// 				CurrentServiceObjectiveName: to.Ptr("ElasticPool"),
		// 				DatabaseID: to.Ptr("82246152-3177-4357-b81c-a16d87ce3593"),
		// 				DefaultSecondaryLocation: to.Ptr("North Europe"),
		// 				ElasticPoolID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default-SQL-SouthEastAsia/providers/Microsoft.Sql/servers/testsvr/elasticPools/Pool1"),
		// 				MaxSizeBytes: to.Ptr[int64](268435456000),
		// 				Status: to.Ptr(armsql.DatabaseStatusOnline),
		// 			},
		// 			SKU: &armsql.SKU{
		// 				Name: to.Ptr("ElasticPool"),
		// 				Tier: to.Ptr("Standard"),
		// 			},
		// 	}},
		// }
	}
}
