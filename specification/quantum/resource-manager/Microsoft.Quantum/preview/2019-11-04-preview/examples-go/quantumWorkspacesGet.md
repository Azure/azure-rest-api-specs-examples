Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fquantum%2Farmquantum%2Fv0.2.0/sdk/resourcemanager/quantum/armquantum/README.md) on how to add the SDK to your project and authenticate.

```go
package armquantum_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/quantum/armquantum"
)

// x-ms-original-file: specification/quantum/resource-manager/Microsoft.Quantum/preview/2019-11-04-preview/examples/quantumWorkspacesGet.json
func ExampleWorkspacesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armquantum.NewWorkspacesClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.WorkspacesClientGetResult)
}
```
