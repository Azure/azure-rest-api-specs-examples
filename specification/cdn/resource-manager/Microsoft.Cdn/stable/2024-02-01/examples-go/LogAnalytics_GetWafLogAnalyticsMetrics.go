package armcdn_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/92de53a5f1e0e03c94b40475d2135d97148ed014/specification/cdn/resource-manager/Microsoft.Cdn/stable/2024-02-01/examples/LogAnalytics_GetWafLogAnalyticsMetrics.json
func ExampleLogAnalyticsClient_GetWafLogAnalyticsMetrics() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcdn.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewLogAnalyticsClient().GetWafLogAnalyticsMetrics(ctx, "RG", "profile1", []armcdn.WafMetric{
		armcdn.WafMetricClientRequestCount}, func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-04T06:49:27.554Z"); return t }(), func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-04T09:49:27.554Z"); return t }(), armcdn.WafGranularityPT5M, &armcdn.LogAnalyticsClientGetWafLogAnalyticsMetricsOptions{Actions: []armcdn.WafAction{
		armcdn.WafActionBlock,
		armcdn.WafActionLog},
		GroupBy:   []armcdn.WafRankingGroupBy{},
		RuleTypes: []armcdn.WafRuleType{},
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.WafMetricsResponse = armcdn.WafMetricsResponse{
	// 	DateTimeBegin: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-04T06:30:27.554Z"); return t}()),
	// 	DateTimeEnd: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-04T09:00:27.554Z"); return t}()),
	// 	Granularity: to.Ptr(armcdn.WafMetricsGranularityPT5M),
	// 	Series: []*armcdn.WafMetricsResponseSeriesItem{
	// 		{
	// 			Data: []*armcdn.Components18OrqelSchemasWafmetricsresponsePropertiesSeriesItemsPropertiesDataItems{
	// 				{
	// 					DateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-04T07:05:00.000Z"); return t}()),
	// 					Value: to.Ptr[float32](2),
	// 				},
	// 				{
	// 					DateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-04T07:10:00.000Z"); return t}()),
	// 					Value: to.Ptr[float32](32),
	// 				},
	// 				{
	// 					DateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-04T07:15:00.000Z"); return t}()),
	// 					Value: to.Ptr[float32](31),
	// 				},
	// 				{
	// 					DateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-04T07:20:00.000Z"); return t}()),
	// 					Value: to.Ptr[float32](63),
	// 				},
	// 				{
	// 					DateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-04T07:25:00.000Z"); return t}()),
	// 					Value: to.Ptr[float32](50),
	// 				},
	// 				{
	// 					DateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-04T07:30:00.000Z"); return t}()),
	// 					Value: to.Ptr[float32](12),
	// 				},
	// 				{
	// 					DateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-04T07:35:00.000Z"); return t}()),
	// 					Value: to.Ptr[float32](8),
	// 				},
	// 				{
	// 					DateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-04T07:40:00.000Z"); return t}()),
	// 					Value: to.Ptr[float32](21),
	// 				},
	// 				{
	// 					DateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-04T07:45:00.000Z"); return t}()),
	// 					Value: to.Ptr[float32](30),
	// 				},
	// 				{
	// 					DateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-04T07:50:00.000Z"); return t}()),
	// 					Value: to.Ptr[float32](18),
	// 				},
	// 				{
	// 					DateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-04T07:55:00.000Z"); return t}()),
	// 					Value: to.Ptr[float32](28),
	// 				},
	// 				{
	// 					DateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-04T08:00:00.000Z"); return t}()),
	// 					Value: to.Ptr[float32](3),
	// 				},
	// 				{
	// 					DateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-04T08:05:00.000Z"); return t}()),
	// 					Value: to.Ptr[float32](58),
	// 				},
	// 				{
	// 					DateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-04T08:10:00.000Z"); return t}()),
	// 					Value: to.Ptr[float32](42),
	// 				},
	// 				{
	// 					DateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-04T08:15:00.000Z"); return t}()),
	// 					Value: to.Ptr[float32](17),
	// 				},
	// 				{
	// 					DateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-04T08:20:00.000Z"); return t}()),
	// 					Value: to.Ptr[float32](21),
	// 				},
	// 				{
	// 					DateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-04T08:25:00.000Z"); return t}()),
	// 					Value: to.Ptr[float32](41),
	// 				},
	// 				{
	// 					DateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-04T08:30:00.000Z"); return t}()),
	// 					Value: to.Ptr[float32](8),
	// 				},
	// 				{
	// 					DateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-04T08:35:00.000Z"); return t}()),
	// 					Value: to.Ptr[float32](15),
	// 				},
	// 				{
	// 					DateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-04T08:40:00.000Z"); return t}()),
	// 					Value: to.Ptr[float32](25),
	// 				},
	// 				{
	// 					DateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-04T08:45:00.000Z"); return t}()),
	// 					Value: to.Ptr[float32](13),
	// 				},
	// 				{
	// 					DateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-04T08:50:00.000Z"); return t}()),
	// 					Value: to.Ptr[float32](17),
	// 				},
	// 				{
	// 					DateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-04T08:55:00.000Z"); return t}()),
	// 					Value: to.Ptr[float32](29),
	// 				},
	// 				{
	// 					DateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-04T09:00:00.000Z"); return t}()),
	// 					Value: to.Ptr[float32](17),
	// 			}},
	// 			Groups: []*armcdn.WafMetricsResponseSeriesPropertiesItemsItem{
	// 			},
	// 			Metric: to.Ptr("clientRequestCount"),
	// 			Unit: to.Ptr(armcdn.WafMetricsSeriesUnitCount),
	// 	}},
	// }
}
