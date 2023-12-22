package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/01e99457ccf5613a95d5b2960d31a12f84018863/specification/sql/resource-manager/Microsoft.Sql/preview/2022-08-01-preview/examples/ManagedInstanceTopQueriesList.json
func ExampleManagedInstancesClient_NewListByManagedInstancePager_obtainListOfInstancesTopResourceConsumingQueries() {
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
		Interval:            to.Ptr(armsql.QueryTimeGrainTypePT1H),
		AggregationFunction: nil,
		ObservationMetric:   to.Ptr(armsql.MetricTypeDuration),
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
		// 			EndTime: to.Ptr("03/05/2020 13:00:00"),
		// 			IntervalType: to.Ptr(armsql.QueryTimeGrainTypeP1D),
		// 			NumberOfQueries: to.Ptr[int32](5),
		// 			ObservationMetric: to.Ptr("cpu"),
		// 			Queries: []*armsql.QueryStatisticsProperties{
		// 				{
		// 					DatabaseName: to.Ptr("db1"),
		// 					Intervals: []*armsql.QueryMetricInterval{
		// 						{
		// 							ExecutionCount: to.Ptr[int64](1),
		// 							IntervalStartTime: to.Ptr("03/03/2020 00:00:00"),
		// 							Metrics: []*armsql.QueryMetricProperties{
		// 								{
		// 									Name: to.Ptr("cpu"),
		// 									DisplayName: to.Ptr("Cpu"),
		// 									Unit: to.Ptr(armsql.QueryMetricUnitTypePercentage),
		// 									Value: to.Ptr[float64](0.0015841714409722222),
		// 								},
		// 								{
		// 									Name: to.Ptr("io"),
		// 									DisplayName: to.Ptr("Physical Io Reads"),
		// 									Unit: to.Ptr(armsql.QueryMetricUnitTypePercentage),
		// 									Value: to.Ptr[float64](0.000004340277777777778),
		// 								},
		// 								{
		// 									Name: to.Ptr("logIo"),
		// 									DisplayName: to.Ptr("Log Writes"),
		// 									Unit: to.Ptr(armsql.QueryMetricUnitTypePercentage),
		// 									Value: to.Ptr[float64](0),
		// 								},
		// 								{
		// 									Name: to.Ptr("memory"),
		// 									DisplayName: to.Ptr("Memory consumption"),
		// 									Unit: to.Ptr(armsql.QueryMetricUnitTypeKB),
		// 									Value: to.Ptr[float64](8336),
		// 								},
		// 								{
		// 									Name: to.Ptr("duration"),
		// 									DisplayName: to.Ptr("Query duration"),
		// 									Unit: to.Ptr(armsql.QueryMetricUnitTypeMicroseconds),
		// 									Value: to.Ptr[float64](11306905),
		// 							}},
		// 					}},
		// 					QueryID: to.Ptr("25"),
		// 				},
		// 				{
		// 					DatabaseName: to.Ptr("db1"),
		// 					Intervals: []*armsql.QueryMetricInterval{
		// 						{
		// 							ExecutionCount: to.Ptr[int64](1),
		// 							IntervalStartTime: to.Ptr("03/03/2020 00:00:00"),
		// 							Metrics: []*armsql.QueryMetricProperties{
		// 								{
		// 									Name: to.Ptr("cpu"),
		// 									DisplayName: to.Ptr("Cpu"),
		// 									Unit: to.Ptr(armsql.QueryMetricUnitTypePercentage),
		// 									Value: to.Ptr[float64](0.0009521432291666667),
		// 								},
		// 								{
		// 									Name: to.Ptr("io"),
		// 									DisplayName: to.Ptr("Physical Io Reads"),
		// 									Unit: to.Ptr(armsql.QueryMetricUnitTypePercentage),
		// 									Value: to.Ptr[float64](2.3148148148148148e-7),
		// 								},
		// 								{
		// 									Name: to.Ptr("logIo"),
		// 									DisplayName: to.Ptr("Log Writes"),
		// 									Unit: to.Ptr(armsql.QueryMetricUnitTypePercentage),
		// 									Value: to.Ptr[float64](0),
		// 								},
		// 								{
		// 									Name: to.Ptr("memory"),
		// 									DisplayName: to.Ptr("Memory consumption"),
		// 									Unit: to.Ptr(armsql.QueryMetricUnitTypeKB),
		// 									Value: to.Ptr[float64](1024),
		// 								},
		// 								{
		// 									Name: to.Ptr("duration"),
		// 									DisplayName: to.Ptr("Query duration"),
		// 									Unit: to.Ptr(armsql.QueryMetricUnitTypeMicroseconds),
		// 									Value: to.Ptr[float64](6620020),
		// 							}},
		// 					}},
		// 					QueryID: to.Ptr("21"),
		// 				},
		// 				{
		// 					DatabaseName: to.Ptr("db3"),
		// 					Intervals: []*armsql.QueryMetricInterval{
		// 						{
		// 							ExecutionCount: to.Ptr[int64](104),
		// 							IntervalStartTime: to.Ptr("03/04/2020 00:00:00"),
		// 							Metrics: []*armsql.QueryMetricProperties{
		// 								{
		// 									Name: to.Ptr("cpu"),
		// 									DisplayName: to.Ptr("Cpu"),
		// 									Unit: to.Ptr(armsql.QueryMetricUnitTypePercentage),
		// 									Value: to.Ptr[float64](0.0008006611689814815),
		// 								},
		// 								{
		// 									Name: to.Ptr("io"),
		// 									DisplayName: to.Ptr("Physical Io Reads"),
		// 									Unit: to.Ptr(armsql.QueryMetricUnitTypePercentage),
		// 									Value: to.Ptr[float64](0),
		// 								},
		// 								{
		// 									Name: to.Ptr("logIo"),
		// 									DisplayName: to.Ptr("Log Writes"),
		// 									Unit: to.Ptr(armsql.QueryMetricUnitTypePercentage),
		// 									Value: to.Ptr[float64](0),
		// 								},
		// 								{
		// 									Name: to.Ptr("memory"),
		// 									DisplayName: to.Ptr("Memory consumption"),
		// 									Unit: to.Ptr(armsql.QueryMetricUnitTypeKB),
		// 									Value: to.Ptr[float64](0),
		// 								},
		// 								{
		// 									Name: to.Ptr("duration"),
		// 									DisplayName: to.Ptr("Query duration"),
		// 									Unit: to.Ptr(armsql.QueryMetricUnitTypeMicroseconds),
		// 									Value: to.Ptr[float64](5543088),
		// 							}},
		// 					}},
		// 					QueryID: to.Ptr("3"),
		// 				},
		// 				{
		// 					DatabaseName: to.Ptr("db2"),
		// 					Intervals: []*armsql.QueryMetricInterval{
		// 						{
		// 							ExecutionCount: to.Ptr[int64](89),
		// 							IntervalStartTime: to.Ptr("03/03/2020 00:00:00"),
		// 							Metrics: []*armsql.QueryMetricProperties{
		// 								{
		// 									Name: to.Ptr("cpu"),
		// 									DisplayName: to.Ptr("Cpu"),
		// 									Unit: to.Ptr(armsql.QueryMetricUnitTypePercentage),
		// 									Value: to.Ptr[float64](0.0006882543402777778),
		// 								},
		// 								{
		// 									Name: to.Ptr("io"),
		// 									DisplayName: to.Ptr("Physical Io Reads"),
		// 									Unit: to.Ptr(armsql.QueryMetricUnitTypePercentage),
		// 									Value: to.Ptr[float64](0),
		// 								},
		// 								{
		// 									Name: to.Ptr("logIo"),
		// 									DisplayName: to.Ptr("Log Writes"),
		// 									Unit: to.Ptr(armsql.QueryMetricUnitTypePercentage),
		// 									Value: to.Ptr[float64](0),
		// 								},
		// 								{
		// 									Name: to.Ptr("memory"),
		// 									DisplayName: to.Ptr("Memory consumption"),
		// 									Unit: to.Ptr(armsql.QueryMetricUnitTypeKB),
		// 									Value: to.Ptr[float64](0),
		// 								},
		// 								{
		// 									Name: to.Ptr("duration"),
		// 									DisplayName: to.Ptr("Query duration"),
		// 									Unit: to.Ptr(armsql.QueryMetricUnitTypeMicroseconds),
		// 									Value: to.Ptr[float64](4761877),
		// 							}},
		// 					}},
		// 					QueryID: to.Ptr("3"),
		// 				},
		// 				{
		// 					DatabaseName: to.Ptr("db3"),
		// 					Intervals: []*armsql.QueryMetricInterval{
		// 						{
		// 							ExecutionCount: to.Ptr[int64](1),
		// 							IntervalStartTime: to.Ptr("03/04/2020 00:00:00"),
		// 							Metrics: []*armsql.QueryMetricProperties{
		// 								{
		// 									Name: to.Ptr("cpu"),
		// 									DisplayName: to.Ptr("Cpu"),
		// 									Unit: to.Ptr(armsql.QueryMetricUnitTypePercentage),
		// 									Value: to.Ptr[float64](0.0006220661168981482),
		// 								},
		// 								{
		// 									Name: to.Ptr("io"),
		// 									DisplayName: to.Ptr("Physical Io Reads"),
		// 									Unit: to.Ptr(armsql.QueryMetricUnitTypePercentage),
		// 									Value: to.Ptr[float64](0),
		// 								},
		// 								{
		// 									Name: to.Ptr("logIo"),
		// 									DisplayName: to.Ptr("Log Writes"),
		// 									Unit: to.Ptr(armsql.QueryMetricUnitTypePercentage),
		// 									Value: to.Ptr[float64](0),
		// 								},
		// 								{
		// 									Name: to.Ptr("memory"),
		// 									DisplayName: to.Ptr("Memory consumption"),
		// 									Unit: to.Ptr(armsql.QueryMetricUnitTypeKB),
		// 									Value: to.Ptr[float64](1024),
		// 								},
		// 								{
		// 									Name: to.Ptr("duration"),
		// 									DisplayName: to.Ptr("Query duration"),
		// 									Unit: to.Ptr(armsql.QueryMetricUnitTypeMicroseconds),
		// 									Value: to.Ptr[float64](4454161),
		// 							}},
		// 					}},
		// 					QueryID: to.Ptr("22"),
		// 			}},
		// 			StartTime: to.Ptr("03/01/2020 00:00:00"),
		// 	}},
		// }
	}
}
