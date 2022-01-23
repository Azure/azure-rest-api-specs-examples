Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fcdn%2Farmcdn%2Fv0.3.0/sdk/resourcemanager/cdn/armcdn/README.md) on how to add the SDK to your project and authenticate.

```go
package armcdn_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn"
)

// x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/LogAnalytics_GetWafLogAnalyticsMetrics.json
func ExampleLogAnalyticsClient_GetWafLogAnalyticsMetrics() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcdn.NewLogAnalyticsClient("<subscription-id>", cred, nil)
	res, err := client.GetWafLogAnalyticsMetrics(ctx,
		"<resource-group-name>",
		"<profile-name>",
		[]armcdn.WafMetric{
			armcdn.WafMetric("clientRequestCount")},
		func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-04T06:49:27.554Z"); return t }(),
		func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-04T09:49:27.554Z"); return t }(),
		armcdn.WafGranularity("PT5M"),
		&armcdn.LogAnalyticsClientGetWafLogAnalyticsMetricsOptions{Actions: []armcdn.WafAction{
			armcdn.WafAction("block"),
			armcdn.WafAction("log")},
			GroupBy:   []armcdn.WafRankingGroupBy{},
			RuleTypes: []armcdn.WafRuleType{},
		})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.LogAnalyticsClientGetWafLogAnalyticsMetricsResult)
}
```
