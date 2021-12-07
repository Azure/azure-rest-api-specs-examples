Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fcontainerinstance%2Farmcontainerinstance%2Fv0.1.0/sdk/resourcemanager/containerinstance/armcontainerinstance/README.md) on how to add the SDK to your project and authenticate.

```go
package armcontainerinstance_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerinstance/armcontainerinstance"
)

// x-ms-original-file: specification/containerinstance/resource-manager/Microsoft.ContainerInstance/stable/2021-09-01/examples/ContainerGroupUsage.json
func ExampleLocationClient_ListUsage() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcontainerinstance.NewLocationClient("<subscription-id>", cred, nil)
	_, err = client.ListUsage(ctx,
		"<location>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}
```
