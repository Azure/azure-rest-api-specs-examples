Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fworkloadmonitor%2Farmworkloadmonitor%2Fv0.1.0/sdk/resourcemanager/workloadmonitor/armworkloadmonitor/README.md) on how to add the SDK to your project and authenticate.

```go
package armworkloadmonitor_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/workloadmonitor/armworkloadmonitor"
)

// x-ms-original-file: specification/workloadmonitor/resource-manager/Microsoft.WorkloadMonitor/preview/2020-01-13-preview/examples/OperationList_Get.json
func ExampleOperationsClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armworkloadmonitor.NewOperationsClient(cred, nil)
	pager := client.List(nil)
	for pager.NextPage(ctx) {
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
	}
}
```
