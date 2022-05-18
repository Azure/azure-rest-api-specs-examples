Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fresources%2Farmchanges%2Fv1.0.0/sdk/resourcemanager/resources/armchanges/README.md) on how to add the SDK to your project and authenticate.

```go
package armchanges_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armchanges"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/resources/resource-manager/Microsoft.Resources/stable/2022-05-01/examples/GetChange.json
func ExampleClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armchanges.NewClient("subscriptionId1", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx,
		"resourceGroup1",
		"resourceProvider1",
		"resourceType1",
		"resourceName1",
		"1d58d72f-0719-4a48-9228-b7ea682885bf",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
```
