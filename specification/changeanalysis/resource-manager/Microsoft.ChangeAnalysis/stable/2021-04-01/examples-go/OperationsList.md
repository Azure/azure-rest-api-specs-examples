Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fchangeanalysis%2Farmchangeanalysis%2Fv0.1.0/sdk/resourcemanager/changeanalysis/armchangeanalysis/README.md) on how to add the SDK to your project and authenticate.

```go
package armchangeanalysis_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/changeanalysis/armchangeanalysis"
)

// x-ms-original-file: specification/changeanalysis/resource-manager/Microsoft.ChangeAnalysis/stable/2021-04-01/examples/OperationsList.json
func ExampleOperationsClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armchangeanalysis.NewOperationsClient(cred, nil)
	pager := client.List(&armchangeanalysis.OperationsListOptions{SkipToken: nil})
	for pager.NextPage(ctx) {
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
	}
}
```
