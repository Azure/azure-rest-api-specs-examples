package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/fe78d8f1e7bd86c778c7e1cafd52cb0e9fec67ef/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ListElasticPoolOperations.json
func ExampleElasticPoolOperationsClient_NewListByElasticPoolPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewElasticPoolOperationsClient().NewListByElasticPoolPager("sqlcrudtestgroup", "sqlcrudtestserver", "testpool", nil)
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
		// page.ElasticPoolOperationListResult = armsql.ElasticPoolOperationListResult{
		// 	Value: []*armsql.ElasticPoolOperation{
		// 		{
		// 			Name: to.Ptr("11111111-1111-1111-1111-111111111111"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/elasticPools/operations"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/sqlcrudtestgroup/providers/Microsoft.Sql/servers/sqlcrudtestserver/elasticPools/testpool/operations/11111111-1111-1111-1111-111111111111"),
		// 			Properties: &armsql.ElasticPoolOperationProperties{
		// 				ElasticPoolName: to.Ptr("testpool"),
		// 				Operation: to.Ptr("UPDATE"),
		// 				OperationFriendlyName: to.Ptr("UPDATE"),
		// 				PercentComplete: to.Ptr[int32](100),
		// 				ServerName: to.Ptr("sqlcrudtestserver"),
		// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-10-30T09:10:08.1Z"); return t}()),
		// 				State: to.Ptr("COMPLETED"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("55555555-5555-5555-5555-555555555555"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/elasticPools/operations"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/sqlcrudtestgroup/providers/Microsoft.Sql/servers/sqlcrudtestserver/elasticPools/testpool/operations/55555555-5555-5555-5555-555555555555"),
		// 			Properties: &armsql.ElasticPoolOperationProperties{
		// 				ElasticPoolName: to.Ptr("testpool"),
		// 				Operation: to.Ptr("UPDATE"),
		// 				OperationFriendlyName: to.Ptr("UPDATE"),
		// 				PercentComplete: to.Ptr[int32](0),
		// 				ServerName: to.Ptr("sqlcrudtestserver"),
		// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-10-30T10:10:08.1Z"); return t}()),
		// 				State: to.Ptr("IN_PROGRESS"),
		// 			},
		// 	}},
		// }
	}
}
