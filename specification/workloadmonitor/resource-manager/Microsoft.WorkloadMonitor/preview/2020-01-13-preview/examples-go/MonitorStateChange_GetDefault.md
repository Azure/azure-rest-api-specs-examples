Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fworkloadmonitor%2Farmworkloadmonitor%2Fv0.2.1/sdk/resourcemanager/workloadmonitor/armworkloadmonitor/README.md) on how to add the SDK to your project and authenticate.

```go
package armworkloadmonitor_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/workloadmonitor/armworkloadmonitor"
)

// x-ms-original-file: specification/workloadmonitor/resource-manager/Microsoft.WorkloadMonitor/preview/2020-01-13-preview/examples/MonitorStateChange_GetDefault.json
func ExampleHealthMonitorsClient_GetStateChange() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armworkloadmonitor.NewHealthMonitorsClient(cred, nil)
	res, err := client.GetStateChange(ctx,
		"<subscription-id>",
		"<resource-group-name>",
		"<provider-name>",
		"<resource-collection-name>",
		"<resource-name>",
		"<monitor-id>",
		"<timestamp-unix>",
		&armworkloadmonitor.HealthMonitorsClientGetStateChangeOptions{Expand: nil})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.HealthMonitorsClientGetStateChangeResult)
}
```
