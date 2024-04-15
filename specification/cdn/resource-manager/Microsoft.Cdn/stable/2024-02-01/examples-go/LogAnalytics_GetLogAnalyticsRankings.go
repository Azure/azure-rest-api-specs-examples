package armcdn_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/92de53a5f1e0e03c94b40475d2135d97148ed014/specification/cdn/resource-manager/Microsoft.Cdn/stable/2024-02-01/examples/LogAnalytics_GetLogAnalyticsRankings.json
func ExampleLogAnalyticsClient_GetLogAnalyticsRankings() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcdn.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewLogAnalyticsClient().GetLogAnalyticsRankings(ctx, "RG", "profile1", []armcdn.LogRanking{
		armcdn.LogRankingURL}, []armcdn.LogRankingMetric{
		armcdn.LogRankingMetricClientRequestCount}, 5, func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-04T06:49:27.554Z"); return t }(), func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-04T09:49:27.554Z"); return t }(), &armcdn.LogAnalyticsClientGetLogAnalyticsRankingsOptions{CustomDomains: []string{}})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.RankingsResponse = armcdn.RankingsResponse{
	// 	DateTimeBegin: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-04T06:49:27.554Z"); return t}()),
	// 	DateTimeEnd: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-04T09:49:27.554Z"); return t}()),
	// 	Tables: []*armcdn.RankingsResponseTablesItem{
	// 		{
	// 			Data: []*armcdn.RankingsResponseTablesPropertiesItemsItem{
	// 				{
	// 					Name: to.Ptr("https://testdomain.com/favicon.png"),
	// 					Metrics: []*armcdn.RankingsResponseTablesPropertiesItemsMetricsItem{
	// 						{
	// 							Metric: to.Ptr("clientRequestCount"),
	// 							Percentage: to.Ptr[float32](8.28133862733976),
	// 							Value: to.Ptr[int64](2336),
	// 					}},
	// 				},
	// 				{
	// 					Name: to.Ptr("https://testdomain.com/js/app.js"),
	// 					Metrics: []*armcdn.RankingsResponseTablesPropertiesItemsMetricsItem{
	// 						{
	// 							Metric: to.Ptr("clientRequestCount"),
	// 							Percentage: to.Ptr[float32](7.586500283607488),
	// 							Value: to.Ptr[int64](2140),
	// 					}},
	// 				},
	// 				{
	// 					Name: to.Ptr("https://testdomain.com/js/lang/en.js"),
	// 					Metrics: []*armcdn.RankingsResponseTablesPropertiesItemsMetricsItem{
	// 						{
	// 							Metric: to.Ptr("clientRequestCount"),
	// 							Percentage: to.Ptr[float32](5.445263754963131),
	// 							Value: to.Ptr[int64](1536),
	// 					}},
	// 				},
	// 				{
	// 					Name: to.Ptr("https://testdomain.com/js/lib.js"),
	// 					Metrics: []*armcdn.RankingsResponseTablesPropertiesItemsMetricsItem{
	// 						{
	// 							Metric: to.Ptr("clientRequestCount"),
	// 							Percentage: to.Ptr[float32](5.246738513896767),
	// 							Value: to.Ptr[int64](1480),
	// 					}},
	// 				},
	// 				{
	// 					Name: to.Ptr("https://cdn.exam.net/css/lib.css"),
	// 					Metrics: []*armcdn.RankingsResponseTablesPropertiesItemsMetricsItem{
	// 						{
	// 							Metric: to.Ptr("clientRequestCount"),
	// 							Percentage: to.Ptr[float32](5.147475893363584),
	// 							Value: to.Ptr[int64](1452),
	// 					}},
	// 			}},
	// 			Ranking: to.Ptr("url"),
	// 	}},
	// }
}
