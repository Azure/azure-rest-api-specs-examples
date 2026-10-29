package armcdn_test

import (
	"context"
	"log"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn/v4"
)

// Generated from example definition: 2026-10-01/LogAnalytics_GetWafLogAnalyticsRankings.json
func ExampleLogAnalyticsClient_GetWafLogAnalyticsRankings() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcdn.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewLogAnalyticsClient().GetWafLogAnalyticsRankings(ctx, "RG", "profile1", []armcdn.WafMetric{
		armcdn.WafMetricClientRequestCount,
	}, time.Date(2020, time.November, 4, 6, 49, 27, 554000000, time.UTC), time.Date(2020, time.November, 4, 9, 49, 27, 554000000, time.UTC), 5, []armcdn.WafRankingType{
		armcdn.WafRankingTypeRuleID,
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcdn.LogAnalyticsClientGetWafLogAnalyticsRankingsResponse{
	// 	WafRankingsResponse: armcdn.WafRankingsResponse{
	// 		Data: []*armcdn.WafRankingsResponseDataItem{
	// 			{
	// 				GroupValues: []*string{
	// 					to.Ptr("BlockRateLimit"),
	// 				},
	// 				Metrics: []*armcdn.ComponentsKpo1PjSchemasWafrankingsresponsePropertiesDataItemsPropertiesMetricsItems{
	// 					{
	// 						Metric: to.Ptr("clientRequestCount"),
	// 						Percentage: to.Ptr[float64](0),
	// 						Value: to.Ptr[int64](1268),
	// 					},
	// 				},
	// 			},
	// 		},
	// 		DateTimeBegin: to.Ptr(time.Date(2020, time.November, 4, 6, 49, 27, 554000000, time.UTC)),
	// 		DateTimeEnd: to.Ptr(time.Date(2020, time.November, 4, 9, 49, 27, 554000000, time.UTC)),
	// 		Groups: []*string{
	// 			to.Ptr("ruleId"),
	// 		},
	// 	},
	// }
}
