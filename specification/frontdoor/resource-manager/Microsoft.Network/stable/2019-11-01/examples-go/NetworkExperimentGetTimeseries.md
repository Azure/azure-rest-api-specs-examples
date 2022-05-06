Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Ffrontdoor%2Farmfrontdoor%2Fv0.4.0/sdk/resourcemanager/frontdoor/armfrontdoor/README.md) on how to add the SDK to your project and authenticate.

```go
package armfrontdoor_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/frontdoor/armfrontdoor"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/frontdoor/resource-manager/Microsoft.Network/stable/2019-11-01/examples/NetworkExperimentGetTimeseries.json
func ExampleReportsClient_GetTimeseries() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armfrontdoor.NewReportsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.GetTimeseries(ctx,
		"<resource-group-name>",
		"<profile-name>",
		"<experiment-name>",
		func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-07-21T17:32:28Z"); return t }(),
		func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-09-21T17:32:28Z"); return t }(),
		armfrontdoor.TimeseriesAggregationIntervalHourly,
		armfrontdoor.TimeseriesTypeMeasurementCounts,
		&armfrontdoor.ReportsClientGetTimeseriesOptions{Endpoint: nil,
			Country: nil,
		})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
