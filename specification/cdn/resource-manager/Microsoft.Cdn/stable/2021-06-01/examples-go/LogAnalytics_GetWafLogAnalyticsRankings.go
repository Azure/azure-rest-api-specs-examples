package armcdn_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/LogAnalytics_GetWafLogAnalyticsRankings.json
func ExampleLogAnalyticsClient_GetWafLogAnalyticsRankings() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcdn.NewLogAnalyticsClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.GetWafLogAnalyticsRankings(ctx,
		"RG",
		"profile1",
		[]armcdn.WafMetric{
			armcdn.WafMetricClientRequestCount},
		func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-04T06:49:27.554Z"); return t }(),
		func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-04T09:49:27.554Z"); return t }(),
		5,
		[]armcdn.WafRankingType{
			armcdn.WafRankingTypeRuleID},
		&armcdn.LogAnalyticsClientGetWafLogAnalyticsRankingsOptions{Actions: []armcdn.WafAction{},
			RuleTypes: []armcdn.WafRuleType{},
		})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
