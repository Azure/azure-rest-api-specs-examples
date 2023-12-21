package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/01e99457ccf5613a95d5b2960d31a12f84018863/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ManagedInstanceQueryStatisticsListMin.json
func ExampleManagedDatabaseQueriesClient_NewListByQueryPager_obtainQueryExecutionStatisticsMinimalExampleWithOnlyMandatoryRequestParameters() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewManagedDatabaseQueriesClient().NewListByQueryPager("sqlcrudtest-7398", "sqlcrudtest-4645", "database_1", "42", &armsql.ManagedDatabaseQueriesClientListByQueryOptions{StartTime: nil,
		EndTime:  nil,
		Interval: to.Ptr(armsql.QueryTimeGrainTypePT1H),
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
		// page.ManagedInstanceQueryStatistics = armsql.ManagedInstanceQueryStatistics{
		// 	Value: []*armsql.QueryStatistics{
		// 		{
		// 			Type: to.Ptr("Microsoft.Sql/managedInstances/databases/queries/statistics"),
		// 			ID: to.Ptr("28"),
		// 			Properties: &armsql.QueryStatisticsProperties{
		// 				DatabaseName: to.Ptr("db1"),
		// 				EndTime: to.Ptr("03/11/2020 14:00:30"),
		// 				Intervals: []*armsql.QueryMetricInterval{
		// 					{
		// 						ExecutionCount: to.Ptr[int64](1),
		// 						IntervalStartTime: to.Ptr("03/11/2020 11:00:00"),
		// 						IntervalType: to.Ptr(armsql.QueryTimeGrainTypePT1H),
		// 						Metrics: []*armsql.QueryMetricProperties{
		// 							{
		// 								Name: to.Ptr("cpu"),
		// 								Avg: to.Ptr[float64](0.03824320138888889),
		// 								DisplayName: to.Ptr("Cpu"),
		// 								Max: to.Ptr[float64](0.03824320138888889),
		// 								Min: to.Ptr[float64](0.03824320138888889),
		// 								Stdev: to.Ptr[float64](0),
		// 								Sum: to.Ptr[float64](0.03824320138888889),
		// 								Unit: to.Ptr(armsql.QueryMetricUnitTypePercentage),
		// 								Value: to.Ptr[float64](0),
		// 							},
		// 							{
		// 								Name: to.Ptr("io"),
		// 								Avg: to.Ptr[float64](0.0001013888888888889),
		// 								DisplayName: to.Ptr("Physical Io Reads"),
		// 								Max: to.Ptr[float64](0.0001013888888888889),
		// 								Min: to.Ptr[float64](0.0001013888888888889),
		// 								Stdev: to.Ptr[float64](0),
		// 								Sum: to.Ptr[float64](0.0001013888888888889),
		// 								Unit: to.Ptr(armsql.QueryMetricUnitTypePercentage),
		// 								Value: to.Ptr[float64](0),
		// 							},
		// 							{
		// 								Name: to.Ptr("logIo"),
		// 								Avg: to.Ptr[float64](0),
		// 								DisplayName: to.Ptr("Log Writes"),
		// 								Max: to.Ptr[float64](0),
		// 								Min: to.Ptr[float64](0),
		// 								Stdev: to.Ptr[float64](0),
		// 								Sum: to.Ptr[float64](0),
		// 								Unit: to.Ptr(armsql.QueryMetricUnitTypePercentage),
		// 								Value: to.Ptr[float64](0),
		// 							},
		// 							{
		// 								Name: to.Ptr("memory"),
		// 								Avg: to.Ptr[float64](8336),
		// 								DisplayName: to.Ptr("Memory consumption"),
		// 								Max: to.Ptr[float64](8336),
		// 								Min: to.Ptr[float64](8336),
		// 								Stdev: to.Ptr[float64](0),
		// 								Sum: to.Ptr[float64](8336),
		// 								Unit: to.Ptr(armsql.QueryMetricUnitTypeKB),
		// 								Value: to.Ptr[float64](0),
		// 							},
		// 							{
		// 								Name: to.Ptr("duration"),
		// 								Avg: to.Ptr[float64](11091296),
		// 								DisplayName: to.Ptr("Query duration"),
		// 								Max: to.Ptr[float64](11091296),
		// 								Min: to.Ptr[float64](11091296),
		// 								Stdev: to.Ptr[float64](0),
		// 								Sum: to.Ptr[float64](11091296),
		// 								Unit: to.Ptr(armsql.QueryMetricUnitTypeMicroseconds),
		// 								Value: to.Ptr[float64](0),
		// 						}},
		// 				}},
		// 				QueryID: to.Ptr("28"),
		// 				StartTime: to.Ptr("03/10/2020 14:00:30"),
		// 			},
		// 	}},
		// }
	}
}
