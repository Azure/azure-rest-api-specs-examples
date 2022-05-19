Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fiotcentral%2Farmiotcentral%2Fv1.0.0/sdk/resourcemanager/iotcentral/armiotcentral/README.md) on how to add the SDK to your project and authenticate.

```go
package armiotcentral_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iotcentral/armiotcentral"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/iotcentral/resource-manager/Microsoft.IoTCentral/stable/2021-06-01/examples/Operations_List.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armiotcentral.NewOperationsClient(cred, nil)
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
