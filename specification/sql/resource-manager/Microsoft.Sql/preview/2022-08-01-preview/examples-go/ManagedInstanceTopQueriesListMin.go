package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/fe78d8f1e7bd86c778c7e1cafd52cb0e9fec67ef/specification/sql/resource-manager/Microsoft.Sql/preview/2022-08-01-preview/examples/ManagedInstanceTopQueriesListMin.json
func ExampleManagedInstancesClient_NewListByManagedInstancePager_obtainListOfInstancesTopResourceConsumingQueriesMinimalRequestAndResponse() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewManagedInstancesClient().NewListByManagedInstancePager("sqlcrudtest-7398", "sqlcrudtest-4645", &armsql.ManagedInstancesClientListByManagedInstanceOptions{NumberOfQueries: nil,
		Databases:           nil,
		StartTime:           nil,
		EndTime:             nil,
		Interval:            nil,
		AggregationFunction: nil,
		ObservationMetric:   nil,
	})
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
		// page.TopQueriesListResult = armsql.TopQueriesListResult{
		// 	Value: []*armsql.TopQueries{
		// 		{
		// 			AggregationFunction: to.Ptr("sum"),
		// 			EndTime: to.Ptr("03/11/2020 12:24:07"),
		// 			IntervalType: to.Ptr(armsql.QueryTimeGrainTypePT1H),
		// 			NumberOfQueries: to.Ptr[int32](5),
		// 			ObservationMetric: to.Ptr("cpu"),
		// 			Queries: []*armsql.QueryStatisticsProperties{
		// 			},
		// 			StartTime: to.Ptr("03/10/2020 12:00:00"),
		// 	}},
		// }
	}
}
