package armmonitor_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
)

// Generated from example definition: 2024-02-01/GetMetricError.json
func ExampleMetricsClient_List_getMetricWithError() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmonitor.NewClientFactory("<subscriptionID>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewMetricsClient().List(ctx, "subscriptions/ac41e21f-afd6-4a79-8070-f01eba278f97/resourceGroups/todking/providers/Microsoft.DocumentDb/databaseAccounts/tk-cosmos-mongo", &armmonitor.MetricsClientListOptions{
		AutoAdjustTimegrain: to.Ptr(true),
		ValidateDimensions:  to.Ptr(false),
		Aggregation:         to.Ptr("average"),
		Interval:            to.Ptr("FULL"),
		Metricnames:         to.Ptr("MongoRequestsCount,MongoRequests"),
		Metricnamespace:     to.Ptr("microsoft.documentdb/databaseaccounts"),
		Timespan:            to.Ptr("2021-06-07T21:51:00Z/2021-06-08T01:51:00Z")})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armmonitor.MetricsClientListResponse{
	// 	Response: armmonitor.Response{
	// 		Cost: to.Ptr[int32](239),
	// 		Interval: to.Ptr("PT4H"),
	// 		Namespace: to.Ptr("microsoft.documentdb/databaseaccounts"),
	// 		Resourceregion: to.Ptr("westus2"),
	// 		Timespan: to.Ptr("2021-06-07T21:51:00Z/2021-06-08T01:51:00Z"),
	// 		Value: []*armmonitor.Metric{
	// 			{
	// 				Name: &armmonitor.LocalizableString{
	// 					LocalizedValue: to.Ptr("(deprecated) Mongo Request Rate"),
	// 					Value: to.Ptr("MongoRequestsCount"),
	// 				},
	// 				Type: to.Ptr("Microsoft.Insights/metrics"),
	// 				ErrorCode: to.Ptr("InvalidSamplingType"),
	// 				ErrorMessage: to.Ptr("Sampling type is not found. Metric:CosmosDBCustomer,AzureMonitor,MongoRequests, SamplingType:NullableAverage."),
	// 				ID: to.Ptr("/subscriptions/ac41e21f-afd6-4a79-8070-f01eba278f97/resourceGroups/todking/providers/Microsoft.DocumentDb/databaseAccounts/tk-cosmos-mongo/providers/Microsoft.Insights/metrics/MongoRequestsCount"),
	// 				Timeseries: []*armmonitor.TimeSeriesElement{
	// 				},
	// 				Unit: to.Ptr(armmonitor.MetricUnitCountPerSecond),
	// 			},
	// 			{
	// 				Name: &armmonitor.LocalizableString{
	// 					LocalizedValue: to.Ptr("Mongo Requests"),
	// 					Value: to.Ptr("MongoRequests"),
	// 				},
	// 				Type: to.Ptr("Microsoft.Insights/metrics"),
	// 				DisplayDescription: to.Ptr("Number of Mongo Requests Made"),
	// 				ErrorCode: to.Ptr("Success"),
	// 				ID: to.Ptr("/subscriptions/ac41e21f-afd6-4a79-8070-f01eba278f97/resourceGroups/todking/providers/Microsoft.DocumentDb/databaseAccounts/tk-cosmos-mongo/providers/Microsoft.Insights/metrics/MongoRequests"),
	// 				Timeseries: []*armmonitor.TimeSeriesElement{
	// 					{
	// 						Data: []*armmonitor.MetricValue{
	// 							{
	// 								Average: to.Ptr[float64](0),
	// 								TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-07T21:51:00Z"); return t}()),
	// 							},
	// 						},
	// 						Metadatavalues: []*armmonitor.MetadataValue{
	// 						},
	// 					},
	// 				},
	// 				Unit: to.Ptr(armmonitor.MetricUnitCount),
	// 			},
	// 		},
	// 	},
	// }
}
