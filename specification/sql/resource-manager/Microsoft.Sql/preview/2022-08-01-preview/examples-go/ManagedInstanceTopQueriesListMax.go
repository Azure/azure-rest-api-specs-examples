package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/fe78d8f1e7bd86c778c7e1cafd52cb0e9fec67ef/specification/sql/resource-manager/Microsoft.Sql/preview/2022-08-01-preview/examples/ManagedInstanceTopQueriesListMax.json
func ExampleManagedInstancesClient_NewListByManagedInstancePager_obtainListOfInstancesTopResourceConsumingQueriesFullBlownRequestAndResponse() {
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
		Databases:           to.Ptr("db1,db2"),
		StartTime:           to.Ptr("2020-03-10T12:00:00Z"),
		EndTime:             to.Ptr("2020-03-12T12:00:00Z"),
		Interval:            to.Ptr(armsql.QueryTimeGrainTypeP1D),
		AggregationFunction: nil,
		ObservationMetric:   to.Ptr(armsql.MetricTypeCPU),
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
		// 			EndTime: to.Ptr("03/12/2020 13:00:00"),
		// 			IntervalType: to.Ptr(armsql.QueryTimeGrainTypeP1D),
		// 			NumberOfQueries: to.Ptr[int32](5),
		// 			ObservationMetric: to.Ptr("cpu"),
		// 			Queries: []*armsql.QueryStatisticsProperties{
		// 				{
		// 					DatabaseName: to.Ptr("db1"),
		// 					Intervals: []*armsql.QueryMetricInterval{
		// 						{
		// 							ExecutionCount: to.Ptr[int64](1),
		// 							IntervalStartTime: to.Ptr("03/11/2020 00:00:00"),
		// 							Metrics: []*armsql.QueryMetricProperties{
		// 								{
		// 									Name: to.Ptr("cpu"),
		// 									DisplayName: to.Ptr("Cpu"),
		// 									Unit: to.Ptr(armsql.QueryMetricUnitTypePercentage),
		// 									Value: to.Ptr[float64](0.0015934667245370371),
		// 								},
		// 								{
		// 									Name: to.Ptr("io"),
		// 									DisplayName: to.Ptr("Physical Io Reads"),
		// 									Unit: to.Ptr(armsql.QueryMetricUnitTypePercentage),
		// 									Value: to.Ptr[float64](0.000004224537037037037),
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
		// 									Value: to.Ptr[float64](11091296),
		// 							}},
		// 					}},
		// 					QueryID: to.Ptr("28"),
		// 				},
		// 				{
		// 					DatabaseName: to.Ptr("db1"),
		// 					Intervals: []*armsql.QueryMetricInterval{
		// 						{
		// 							ExecutionCount: to.Ptr[int64](1),
		// 							IntervalStartTime: to.Ptr("03/11/2020 00:00:00"),
		// 							Metrics: []*armsql.QueryMetricProperties{
		// 								{
		// 									Name: to.Ptr("cpu"),
		// 									DisplayName: to.Ptr("Cpu"),
		// 									Unit: to.Ptr(armsql.QueryMetricUnitTypePercentage),
		// 									Value: to.Ptr[float64](0.0009522783564814815),
		// 								},
		// 								{
		// 									Name: to.Ptr("io"),
		// 									DisplayName: to.Ptr("Physical Io Reads"),
		// 									Unit: to.Ptr(armsql.QueryMetricUnitTypePercentage),
		// 									Value: to.Ptr[float64](1.7361111111111112e-7),
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
		// 									Value: to.Ptr[float64](6625562),
		// 							}},
		// 					}},
		// 					QueryID: to.Ptr("24"),
		// 				},
		// 				{
		// 					DatabaseName: to.Ptr("db1"),
		// 					Intervals: []*armsql.QueryMetricInterval{
		// 						{
		// 							ExecutionCount: to.Ptr[int64](82),
		// 							IntervalStartTime: to.Ptr("03/11/2020 00:00:00"),
		// 							Metrics: []*armsql.QueryMetricProperties{
		// 								{
		// 									Name: to.Ptr("cpu"),
		// 									DisplayName: to.Ptr("Cpu"),
		// 									Unit: to.Ptr(armsql.QueryMetricUnitTypePercentage),
		// 									Value: to.Ptr[float64](0.0007183139467592593),
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
		// 									Value: to.Ptr[float64](4970199),
		// 							}},
		// 					}},
		// 					QueryID: to.Ptr("3"),
		// 				},
		// 				{
		// 					DatabaseName: to.Ptr("db1"),
		// 					Intervals: []*armsql.QueryMetricInterval{
		// 						{
		// 							ExecutionCount: to.Ptr[int64](1),
		// 							IntervalStartTime: to.Ptr("03/11/2020 00:00:00"),
		// 							Metrics: []*armsql.QueryMetricProperties{
		// 								{
		// 									Name: to.Ptr("cpu"),
		// 									DisplayName: to.Ptr("Cpu"),
		// 									Unit: to.Ptr(armsql.QueryMetricUnitTypePercentage),
		// 									Value: to.Ptr[float64](0.0006494454571759259),
		// 								},
		// 								{
		// 									Name: to.Ptr("io"),
		// 									DisplayName: to.Ptr("Physical Io Reads"),
		// 									Unit: to.Ptr(armsql.QueryMetricUnitTypePercentage),
		// 									Value: to.Ptr[float64](0.000005034722222222222),
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
		// 									Value: to.Ptr[float64](4530668),
		// 							}},
		// 					}},
		// 					QueryID: to.Ptr("29"),
		// 				},
		// 				{
		// 					DatabaseName: to.Ptr("db2"),
		// 					Intervals: []*armsql.QueryMetricInterval{
		// 						{
		// 							ExecutionCount: to.Ptr[int64](1),
		// 							IntervalStartTime: to.Ptr("03/11/2020 00:00:00"),
		// 							Metrics: []*armsql.QueryMetricProperties{
		// 								{
		// 									Name: to.Ptr("cpu"),
		// 									DisplayName: to.Ptr("Cpu"),
		// 									Unit: to.Ptr(armsql.QueryMetricUnitTypePercentage),
		// 									Value: to.Ptr[float64](0.0006275368923611112),
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
		// 									Value: to.Ptr[float64](4349943),
		// 							}},
		// 					}},
		// 					QueryID: to.Ptr("25"),
		// 			}},
		// 			StartTime: to.Ptr("03/10/2020 00:00:00"),
		// 	}},
		// }
	}
}
