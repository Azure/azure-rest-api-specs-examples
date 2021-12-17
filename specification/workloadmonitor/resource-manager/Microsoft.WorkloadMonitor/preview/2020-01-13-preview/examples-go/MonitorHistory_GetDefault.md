Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fworkloadmonitor%2Farmworkloadmonitor%2Fv0.1.0/sdk/resourcemanager/workloadmonitor/armworkloadmonitor/README.md) on how to add the SDK to your project and authenticate.

```go
package armworkloadmonitor_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/workloadmonitor/armworkloadmonitor"
)

// x-ms-original-file: specification/workloadmonitor/resource-manager/Microsoft.WorkloadMonitor/preview/2020-01-13-preview/examples/MonitorHistory_GetDefault.json
func ExampleHealthMonitorsClient_ListStateChanges() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armworkloadmonitor.NewHealthMonitorsClient(cred, nil)
	pager := client.ListStateChanges("<subscription-id>",
		"<resource-group-name>",
		"<provider-name>",
		"<resource-collection-name>",
		"<resource-name>",
		"<monitor-id>",
		&armworkloadmonitor.HealthMonitorsListStateChangesOptions{Filter: nil,
			Expand:            nil,
			StartTimestampUTC: to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-10-19T19:24:14Z"); return t }()),
			EndTimestampUTC:   to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-10-20T01:24:14Z"); return t }()),
		})
	for pager.NextPage(ctx) {
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("HealthMonitorStateChange.ID: %s\n", *v.ID)
		}
	}
}
```
