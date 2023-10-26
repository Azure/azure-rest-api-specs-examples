package armcdn_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7b551033155a63739b6d28f79b9c07569f6179b8/specification/cdn/resource-manager/Microsoft.Cdn/stable/2023-05-01/examples/LogAnalytics_GetLogAnalyticsMetrics.json
func ExampleLogAnalyticsClient_GetLogAnalyticsMetrics() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcdn.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewLogAnalyticsClient().GetLogAnalyticsMetrics(ctx, "RG", "profile1", []armcdn.LogMetric{
		armcdn.LogMetricClientRequestCount}, func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-04T04:30:00.000Z"); return t }(), func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-04T05:00:00.000Z"); return t }(), armcdn.LogMetricsGranularityPT5M, []string{
		"customdomain1.azurecdn.net",
		"customdomain2.azurecdn.net"}, []string{
		"https"}, &armcdn.LogAnalyticsClientGetLogAnalyticsMetricsOptions{GroupBy: []armcdn.LogMetricsGroupBy{
		armcdn.LogMetricsGroupByProtocol},
		Continents:       []string{},
		CountryOrRegions: []string{},
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.MetricsResponse = armcdn.MetricsResponse{
	// 	DateTimeBegin: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-04T12:30:27.554+08:00"); return t}()),
	// 	DateTimeEnd: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-04T13:00:27.554+08:00"); return t}()),
	// 	Granularity: to.Ptr(armcdn.MetricsGranularityPT5M),
	// 	Series: []*armcdn.MetricsResponseSeriesItem{
	// 		{
	// 			Data: []*armcdn.Components1Gs0LlpSchemasMetricsresponsePropertiesSeriesItemsPropertiesDataItems{
	// 				{
	// 					DateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-04T04:35:00+00:00"); return t}()),
	// 					Value: to.Ptr[float32](4250),
	// 				},
	// 				{
	// 					DateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-04T04:40:00+00:00"); return t}()),
	// 					Value: to.Ptr[float32](3120),
	// 				},
	// 				{
	// 					DateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-04T04:45:00+00:00"); return t}()),
	// 					Value: to.Ptr[float32](2221),
	// 				},
	// 				{
	// 					DateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-04T04:50:00+00:00"); return t}()),
	// 					Value: to.Ptr[float32](2466),
	// 				},
	// 				{
	// 					DateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-04T04:55:00+00:00"); return t}()),
	// 					Value: to.Ptr[float32](2654),
	// 				},
	// 				{
	// 					DateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-04T05:00:00+00:00"); return t}()),
	// 					Value: to.Ptr[float32](3565),
	// 			}},
	// 			Groups: []*armcdn.MetricsResponseSeriesPropertiesItemsItem{
	// 				{
	// 					Name: to.Ptr("protocol"),
	// 					Value: to.Ptr("https"),
	// 			}},
	// 			Metric: to.Ptr("clientRequestCount"),
	// 			Unit: to.Ptr(armcdn.MetricsSeriesUnitCount),
	// 	}},
	// }
}
