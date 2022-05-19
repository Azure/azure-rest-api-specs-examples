Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fwindowsiot%2Farmwindowsiot%2Fv1.0.0/sdk/resourcemanager/windowsiot/armwindowsiot/README.md) on how to add the SDK to your project and authenticate.

```go
package armwindowsiot_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/windowsiot/armwindowsiot"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/windowsiot/resource-manager/Microsoft.WindowsIoT/stable/2019-06-01/examples/Service_List.json
func ExampleServicesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armwindowsiot.NewServicesClient("27de630f-e1ee-42de-8849-90def4986454", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListPager(nil)
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}
```
