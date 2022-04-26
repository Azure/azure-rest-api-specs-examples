Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fworkloadmonitor%2Farmworkloadmonitor%2Fv0.4.0/sdk/resourcemanager/workloadmonitor/armworkloadmonitor/README.md) on how to add the SDK to your project and authenticate.

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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/workloadmonitor/resource-manager/Microsoft.WorkloadMonitor/preview/2020-01-13-preview/examples/MonitorHistory_GetDefault.json
func ExampleHealthMonitorsClient_NewListStateChangesPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armworkloadmonitor.NewHealthMonitorsClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	pager := client.NewListStateChangesPager("<subscription-id>",
		"<resource-group-name>",
		"<provider-name>",
		"<resource-collection-name>",
		"<resource-name>",
		"<monitor-id>",
		&armworkloadmonitor.HealthMonitorsClientListStateChangesOptions{Filter: nil,
			Expand:            nil,
			StartTimestampUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-10-19T19:24:14Z"); return t }()),
			EndTimestampUTC:   to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-10-20T01:24:14Z"); return t }()),
		})
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
			return
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}
```
