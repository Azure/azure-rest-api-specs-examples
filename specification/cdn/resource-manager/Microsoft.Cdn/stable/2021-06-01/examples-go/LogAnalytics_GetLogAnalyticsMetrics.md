Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fcdn%2Farmcdn%2Fv1.0.0/sdk/resourcemanager/cdn/armcdn/README.md) on how to add the SDK to your project and authenticate.

```go
package armcdn_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/LogAnalytics_GetLogAnalyticsMetrics.json
func ExampleLogAnalyticsClient_GetLogAnalyticsMetrics() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcdn.NewLogAnalyticsClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.GetLogAnalyticsMetrics(ctx,
		"RG",
		"profile1",
		[]armcdn.LogMetric{
			armcdn.LogMetricClientRequestCount},
		func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-04T04:30:00.000Z"); return t }(),
		func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-04T05:00:00.000Z"); return t }(),
		armcdn.LogMetricsGranularityPT5M,
		[]string{
			"customdomain1.azurecdn.net",
			"customdomain2.azurecdn.net"},
		[]string{
			"https"},
		&armcdn.LogAnalyticsClientGetLogAnalyticsMetricsOptions{GroupBy: []armcdn.LogMetricsGroupBy{
			armcdn.LogMetricsGroupByProtocol},
			Continents:       []string{},
			CountryOrRegions: []string{},
		})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
```
