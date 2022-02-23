Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fmonitor%2Farmmonitor%2Fv0.4.1/sdk/resourcemanager/monitor/armmonitor/README.md) on how to add the SDK to your project and authenticate.

```go
package armmonitor_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
)

// x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/stable/2018-01-01/examples/GetMetric.json
func ExampleMetricsClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmonitor.NewMetricsClient(cred, nil)
	res, err := client.List(ctx,
		"<resource-uri>",
		&armmonitor.MetricsClientListOptions{Timespan: to.StringPtr("<timespan>"),
			Interval:        to.StringPtr("<interval>"),
			Metricnames:     nil,
			Aggregation:     to.StringPtr("<aggregation>"),
			Top:             to.Int32Ptr(3),
			Orderby:         to.StringPtr("<orderby>"),
			Filter:          to.StringPtr("<filter>"),
			ResultType:      nil,
			Metricnamespace: to.StringPtr("<metricnamespace>"),
		})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.MetricsClientListResult)
}
```
