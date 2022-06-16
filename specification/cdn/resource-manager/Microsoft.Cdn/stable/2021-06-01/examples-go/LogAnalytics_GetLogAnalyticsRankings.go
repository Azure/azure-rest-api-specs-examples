package armcdn_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/LogAnalytics_GetLogAnalyticsRankings.json
func ExampleLogAnalyticsClient_GetLogAnalyticsRankings() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcdn.NewLogAnalyticsClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.GetLogAnalyticsRankings(ctx,
		"RG",
		"profile1",
		[]armcdn.LogRanking{
			armcdn.LogRankingURL},
		[]armcdn.LogRankingMetric{
			armcdn.LogRankingMetricClientRequestCount},
		5,
		func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-04T06:49:27.554Z"); return t }(),
		func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-04T09:49:27.554Z"); return t }(),
		&armcdn.LogAnalyticsClientGetLogAnalyticsRankingsOptions{CustomDomains: []string{}})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
