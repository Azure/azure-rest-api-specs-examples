Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fcontainerservice%2Farmcontainerservice%2Fv0.2.1/sdk/resourcemanager/containerservice/armcontainerservice/README.md) on how to add the SDK to your project and authenticate.

```go
package armcontainerservice_test

import (
	"context"
	"log"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice"
)

// x-ms-original-file: specification/containerservice/resource-manager/Microsoft.ContainerService/stable/2021-09-01/examples/ManagedClustersStart.json
func ExampleManagedClustersClient_BeginStart() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcontainerservice.NewManagedClustersClient("<subscription-id>", cred, nil)
	poller, err := client.BeginStart(ctx,
		"<resource-group-name>",
		"<resource-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}
```
