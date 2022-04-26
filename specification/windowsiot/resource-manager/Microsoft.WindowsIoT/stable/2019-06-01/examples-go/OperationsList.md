Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fwindowsiot%2Farmwindowsiot%2Fv0.4.0/sdk/resourcemanager/windowsiot/armwindowsiot/README.md) on how to add the SDK to your project and authenticate.

```go
package armwindowsiot_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/windowsiot/armwindowsiot"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/windowsiot/resource-manager/Microsoft.WindowsIoT/stable/2019-06-01/examples/OperationsList.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armwindowsiot.NewOperationsClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	pager := client.NewListPager(nil)
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
