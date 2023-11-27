package armcdn_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7b551033155a63739b6d28f79b9c07569f6179b8/specification/cdn/resource-manager/Microsoft.Cdn/stable/2023-05-01/examples/LogAnalytics_GetWafLogAnalyticsRankings.json
func ExampleLogAnalyticsClient_GetWafLogAnalyticsRankings() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcdn.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewLogAnalyticsClient().GetWafLogAnalyticsRankings(ctx, "RG", "profile1", []armcdn.WafMetric{
		armcdn.WafMetricClientRequestCount}, func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-04T06:49:27.554Z"); return t }(), func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-04T09:49:27.554Z"); return t }(), 5, []armcdn.WafRankingType{
		armcdn.WafRankingTypeRuleID}, &armcdn.LogAnalyticsClientGetWafLogAnalyticsRankingsOptions{Actions: []armcdn.WafAction{},
		RuleTypes: []armcdn.WafRuleType{},
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.WafRankingsResponse = armcdn.WafRankingsResponse{
	// 	Data: []*armcdn.WafRankingsResponseDataItem{
	// 		{
	// 			GroupValues: []*string{
	// 				to.Ptr("BlockRateLimit")},
	// 				Metrics: []*armcdn.ComponentsKpo1PjSchemasWafrankingsresponsePropertiesDataItemsPropertiesMetricsItems{
	// 					{
	// 						Metric: to.Ptr("clientRequestCount"),
	// 						Percentage: to.Ptr[float64](0),
	// 						Value: to.Ptr[int64](1268),
	// 				}},
	// 		}},
	// 		DateTimeBegin: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-04T06:49:27.554Z"); return t}()),
	// 		DateTimeEnd: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-04T09:49:27.554Z"); return t}()),
	// 		Groups: []*string{
	// 			to.Ptr("ruleId")},
	// 		}
}
