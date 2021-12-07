Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fazurearcdata%2Farmazurearcdata%2Fv0.1.0/sdk/resourcemanager/azurearcdata/armazurearcdata/README.md) on how to add the SDK to your project and authenticate.

```go
package armazurearcdata_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azurearcdata/armazurearcdata"
)

// x-ms-original-file: specification/azurearcdata/resource-manager/Microsoft.AzureArcData/stable/2021-11-01/examples/DeleteDataController.json
func ExampleDataControllersClient_BeginDeleteDataController() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armazurearcdata.NewDataControllersClient("<subscription-id>", cred, nil)
	poller, err := client.BeginDeleteDataController(ctx,
		"<resource-group-name>",
		"<data-controller-name>",
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
