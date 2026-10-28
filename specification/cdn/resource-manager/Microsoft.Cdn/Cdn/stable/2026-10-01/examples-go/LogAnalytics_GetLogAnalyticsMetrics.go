package armcdn_test

import (
	"context"
	"log"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn/v4"
)

// Generated from example definition: 2026-10-01/LogAnalytics_GetLogAnalyticsMetrics.json
func ExampleLogAnalyticsClient_GetLogAnalyticsMetrics() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcdn.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewLogAnalyticsClient().GetLogAnalyticsMetrics(ctx, "RG", "profile1", []armcdn.LogMetric{
		armcdn.LogMetricClientRequestCount,
	}, time.Date(2020, time.November, 4, 4, 30, 0, 0, time.UTC), time.Date(2020, time.November, 4, 5, 0, 0, 0, time.UTC), armcdn.LogMetricsGranularityPT5M, []string{
		"customdomain1.azurecdn.net",
		"customdomain2.azurecdn.net",
	}, []string{
		"https",
	}, &armcdn.LogAnalyticsClientGetLogAnalyticsMetricsOptions{
		GroupBy: []armcdn.LogMetricsGroupBy{
			armcdn.LogMetricsGroupByProtocol,
		}})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcdn.LogAnalyticsClientGetLogAnalyticsMetricsResponse{
	// 	MetricsResponse: armcdn.MetricsResponse{
	// 		DateTimeBegin: to.Ptr(time.Date(2020, time.November, 4, 4, 30, 27, 554000000, time.UTC)),
	// 		DateTimeEnd: to.Ptr(time.Date(2020, time.November, 4, 5, 0, 27, 554000000, time.UTC)),
	// 		Granularity: to.Ptr(armcdn.MetricsGranularityPT5M),
	// 		Series: []*armcdn.MetricsResponseSeriesItem{
	// 			{
	// 				Data: []*armcdn.Components1Gs0LlpSchemasMetricsresponsePropertiesSeriesItemsPropertiesDataItems{
	// 					{
	// 						DateTime: to.Ptr(time.Date(2020, time.November, 4, 4, 35, 0, 0, time.UTC)),
	// 						Value: to.Ptr[float32](4250),
	// 					},
	// 					{
	// 						DateTime: to.Ptr(time.Date(2020, time.November, 4, 4, 40, 0, 0, time.UTC)),
	// 						Value: to.Ptr[float32](3120),
	// 					},
	// 					{
	// 						DateTime: to.Ptr(time.Date(2020, time.November, 4, 4, 45, 0, 0, time.UTC)),
	// 						Value: to.Ptr[float32](2221),
	// 					},
	// 					{
	// 						DateTime: to.Ptr(time.Date(2020, time.November, 4, 4, 50, 0, 0, time.UTC)),
	// 						Value: to.Ptr[float32](2466),
	// 					},
	// 					{
	// 						DateTime: to.Ptr(time.Date(2020, time.November, 4, 4, 55, 0, 0, time.UTC)),
	// 						Value: to.Ptr[float32](2654),
	// 					},
	// 					{
	// 						DateTime: to.Ptr(time.Date(2020, time.November, 4, 5, 0, 0, 0, time.UTC)),
	// 						Value: to.Ptr[float32](3565),
	// 					},
	// 				},
	// 				Groups: []*armcdn.MetricsResponseSeriesPropertiesItemsItem{
	// 					{
	// 						Name: to.Ptr("protocol"),
	// 						Value: to.Ptr("https"),
	// 					},
	// 				},
	// 				Metric: to.Ptr("clientRequestCount"),
	// 				Unit: to.Ptr(armcdn.MetricsSeriesUnitCount),
	// 			},
	// 		},
	// 	},
	// }
}
