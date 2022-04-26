Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fworkloadmonitor%2Farmworkloadmonitor%2Fv0.4.0/sdk/resourcemanager/workloadmonitor/armworkloadmonitor/README.md) on how to add the SDK to your project and authenticate.

```go
package armworkloadmonitor_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/workloadmonitor/armworkloadmonitor"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/workloadmonitor/resource-manager/Microsoft.WorkloadMonitor/preview/2020-01-13-preview/examples/MonitorList_GetDefault.json
func ExampleHealthMonitorsClient_NewListPager() {
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
	pager := client.NewListPager("<subscription-id>",
		"<resource-group-name>",
		"<provider-name>",
		"<resource-collection-name>",
		"<resource-name>",
		&armworkloadmonitor.HealthMonitorsClientListOptions{Filter: nil,
			Expand: nil,
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
