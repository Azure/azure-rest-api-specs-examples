package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/01e99457ccf5613a95d5b2960d31a12f84018863/specification/sql/resource-manager/Microsoft.Sql/stable/2014-04-01/examples/DatabaseMetricsListWithFilter.json
func ExampleDatabasesClient_NewListMetricsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDatabasesClient().NewListMetricsPager("sqlcrudtest-6730", "sqlcrudtest-9007", "3481", "name/value eq 'cpu_percent' and timeGrain eq '00:10:00' and startTime eq '2017-06-02T18:35:00Z' and endTime eq '2017-06-02T18:55:00Z'", nil)
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
		// page.MetricListResult = armsql.MetricListResult{
		// 	Value: []*armsql.Metric{
		// 		{
		// 			Name: &armsql.MetricName{
		// 				LocalizedValue: to.Ptr("CPU percentage"),
		// 				Value: to.Ptr("cpu_percent"),
		// 			},
		// 			EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-06-02T18:55:00.000Z"); return t}()),
		// 			MetricValues: []*armsql.MetricValue{
		// 				{
		// 					Average: to.Ptr[float64](0),
		// 					Count: to.Ptr[int32](1),
		// 					Maximum: to.Ptr[float64](0),
		// 					Minimum: to.Ptr[float64](0),
		// 					Timestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-06-02T18:30:01.000Z"); return t}()),
		// 					Total: to.Ptr[float64](0),
		// 				},
		// 				{
		// 					Average: to.Ptr[float64](0),
		// 					Count: to.Ptr[int32](1),
		// 					Maximum: to.Ptr[float64](0),
		// 					Minimum: to.Ptr[float64](0),
		// 					Timestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-06-02T18:40:01.000Z"); return t}()),
		// 					Total: to.Ptr[float64](0),
		// 				},
		// 				{
		// 					Average: to.Ptr[float64](0),
		// 					Count: to.Ptr[int32](1),
		// 					Maximum: to.Ptr[float64](0),
		// 					Minimum: to.Ptr[float64](0),
		// 					Timestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-06-02T18:50:01.000Z"); return t}()),
		// 					Total: to.Ptr[float64](0),
		// 			}},
		// 			StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-06-02T18:35:00.000Z"); return t}()),
		// 			TimeGrain: to.Ptr("00:10:00"),
		// 			Unit: to.Ptr(armsql.UnitTypePercent),
		// 	}},
		// }
	}
}
