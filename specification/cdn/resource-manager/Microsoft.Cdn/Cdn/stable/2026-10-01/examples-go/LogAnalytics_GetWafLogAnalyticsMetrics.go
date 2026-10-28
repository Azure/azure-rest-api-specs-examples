package armcdn_test

import (
	"context"
	"log"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn/v4"
)

// Generated from example definition: 2026-10-01/LogAnalytics_GetWafLogAnalyticsMetrics.json
func ExampleLogAnalyticsClient_GetWafLogAnalyticsMetrics() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcdn.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewLogAnalyticsClient().GetWafLogAnalyticsMetrics(ctx, "RG", "profile1", []armcdn.WafMetric{
		armcdn.WafMetricClientRequestCount,
	}, time.Date(2020, time.November, 4, 6, 49, 27, 554000000, time.UTC), time.Date(2020, time.November, 4, 9, 49, 27, 554000000, time.UTC), armcdn.WafGranularityPT5M, &armcdn.LogAnalyticsClientGetWafLogAnalyticsMetricsOptions{
		Actions: []armcdn.WafAction{
			armcdn.WafActionBlock,
			armcdn.WafActionLog,
		}})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcdn.LogAnalyticsClientGetWafLogAnalyticsMetricsResponse{
	// 	WafMetricsResponse: armcdn.WafMetricsResponse{
	// 		DateTimeBegin: to.Ptr(time.Date(2020, time.November, 4, 6, 30, 27, 554000000, time.UTC)),
	// 		DateTimeEnd: to.Ptr(time.Date(2020, time.November, 4, 9, 0, 27, 554000000, time.UTC)),
	// 		Granularity: to.Ptr(armcdn.WafMetricsGranularityPT5M),
	// 		Series: []*armcdn.WafMetricsResponseSeriesItem{
	// 			{
	// 				Data: []*armcdn.Components18OrqelSchemasWafmetricsresponsePropertiesSeriesItemsPropertiesDataItems{
	// 					{
	// 						DateTime: to.Ptr(time.Date(2020, time.November, 4, 7, 5, 0, 0, time.UTC)),
	// 						Value: to.Ptr[float32](2),
	// 					},
	// 					{
	// 						DateTime: to.Ptr(time.Date(2020, time.November, 4, 7, 10, 0, 0, time.UTC)),
	// 						Value: to.Ptr[float32](32),
	// 					},
	// 					{
	// 						DateTime: to.Ptr(time.Date(2020, time.November, 4, 7, 15, 0, 0, time.UTC)),
	// 						Value: to.Ptr[float32](31),
	// 					},
	// 					{
	// 						DateTime: to.Ptr(time.Date(2020, time.November, 4, 7, 20, 0, 0, time.UTC)),
	// 						Value: to.Ptr[float32](63),
	// 					},
	// 					{
	// 						DateTime: to.Ptr(time.Date(2020, time.November, 4, 7, 25, 0, 0, time.UTC)),
	// 						Value: to.Ptr[float32](50),
	// 					},
	// 					{
	// 						DateTime: to.Ptr(time.Date(2020, time.November, 4, 7, 30, 0, 0, time.UTC)),
	// 						Value: to.Ptr[float32](12),
	// 					},
	// 					{
	// 						DateTime: to.Ptr(time.Date(2020, time.November, 4, 7, 35, 0, 0, time.UTC)),
	// 						Value: to.Ptr[float32](8),
	// 					},
	// 					{
	// 						DateTime: to.Ptr(time.Date(2020, time.November, 4, 7, 40, 0, 0, time.UTC)),
	// 						Value: to.Ptr[float32](21),
	// 					},
	// 					{
	// 						DateTime: to.Ptr(time.Date(2020, time.November, 4, 7, 45, 0, 0, time.UTC)),
	// 						Value: to.Ptr[float32](30),
	// 					},
	// 					{
	// 						DateTime: to.Ptr(time.Date(2020, time.November, 4, 7, 50, 0, 0, time.UTC)),
	// 						Value: to.Ptr[float32](18),
	// 					},
	// 					{
	// 						DateTime: to.Ptr(time.Date(2020, time.November, 4, 7, 55, 0, 0, time.UTC)),
	// 						Value: to.Ptr[float32](28),
	// 					},
	// 					{
	// 						DateTime: to.Ptr(time.Date(2020, time.November, 4, 8, 0, 0, 0, time.UTC)),
	// 						Value: to.Ptr[float32](3),
	// 					},
	// 					{
	// 						DateTime: to.Ptr(time.Date(2020, time.November, 4, 8, 5, 0, 0, time.UTC)),
	// 						Value: to.Ptr[float32](58),
	// 					},
	// 					{
	// 						DateTime: to.Ptr(time.Date(2020, time.November, 4, 8, 10, 0, 0, time.UTC)),
	// 						Value: to.Ptr[float32](42),
	// 					},
	// 					{
	// 						DateTime: to.Ptr(time.Date(2020, time.November, 4, 8, 15, 0, 0, time.UTC)),
	// 						Value: to.Ptr[float32](17),
	// 					},
	// 					{
	// 						DateTime: to.Ptr(time.Date(2020, time.November, 4, 8, 20, 0, 0, time.UTC)),
	// 						Value: to.Ptr[float32](21),
	// 					},
	// 					{
	// 						DateTime: to.Ptr(time.Date(2020, time.November, 4, 8, 25, 0, 0, time.UTC)),
	// 						Value: to.Ptr[float32](41),
	// 					},
	// 					{
	// 						DateTime: to.Ptr(time.Date(2020, time.November, 4, 8, 30, 0, 0, time.UTC)),
	// 						Value: to.Ptr[float32](8),
	// 					},
	// 					{
	// 						DateTime: to.Ptr(time.Date(2020, time.November, 4, 8, 35, 0, 0, time.UTC)),
	// 						Value: to.Ptr[float32](15),
	// 					},
	// 					{
	// 						DateTime: to.Ptr(time.Date(2020, time.November, 4, 8, 40, 0, 0, time.UTC)),
	// 						Value: to.Ptr[float32](25),
	// 					},
	// 					{
	// 						DateTime: to.Ptr(time.Date(2020, time.November, 4, 8, 45, 0, 0, time.UTC)),
	// 						Value: to.Ptr[float32](13),
	// 					},
	// 					{
	// 						DateTime: to.Ptr(time.Date(2020, time.November, 4, 8, 50, 0, 0, time.UTC)),
	// 						Value: to.Ptr[float32](17),
	// 					},
	// 					{
	// 						DateTime: to.Ptr(time.Date(2020, time.November, 4, 8, 55, 0, 0, time.UTC)),
	// 						Value: to.Ptr[float32](29),
	// 					},
	// 					{
	// 						DateTime: to.Ptr(time.Date(2020, time.November, 4, 9, 0, 0, 0, time.UTC)),
	// 						Value: to.Ptr[float32](17),
	// 					},
	// 				},
	// 				Groups: []*armcdn.WafMetricsResponseSeriesPropertiesItemsItem{
	// 				},
	// 				Metric: to.Ptr("clientRequestCount"),
	// 				Unit: to.Ptr(armcdn.WafMetricsSeriesUnitCount),
	// 			},
	// 		},
	// 	},
	// }
}
