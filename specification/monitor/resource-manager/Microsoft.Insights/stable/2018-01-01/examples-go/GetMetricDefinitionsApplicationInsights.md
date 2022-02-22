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

// x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/stable/2018-01-01/examples/GetMetricDefinitionsApplicationInsights.json
func ExampleMetricDefinitionsClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmonitor.NewMetricDefinitionsClient(cred, nil)
	res, err := client.List(ctx,
		"<resource-uri>",
		&armmonitor.MetricDefinitionsClientListOptions{Metricnamespace: to.StringPtr("<metricnamespace>")})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.MetricDefinitionsClientListResult)
}
```
