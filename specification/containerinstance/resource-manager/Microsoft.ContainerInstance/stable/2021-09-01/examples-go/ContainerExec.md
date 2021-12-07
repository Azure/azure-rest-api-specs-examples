Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fcontainerinstance%2Farmcontainerinstance%2Fv0.1.0/sdk/resourcemanager/containerinstance/armcontainerinstance/README.md) on how to add the SDK to your project and authenticate.

```go
package armcontainerinstance_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerinstance/armcontainerinstance"
)

// x-ms-original-file: specification/containerinstance/resource-manager/Microsoft.ContainerInstance/stable/2021-09-01/examples/ContainerExec.json
func ExampleContainersClient_ExecuteCommand() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcontainerinstance.NewContainersClient("<subscription-id>", cred, nil)
	_, err = client.ExecuteCommand(ctx,
		"<resource-group-name>",
		"<container-group-name>",
		"<container-name>",
		armcontainerinstance.ContainerExecRequest{
			Command: to.StringPtr("<command>"),
			TerminalSize: &armcontainerinstance.ContainerExecRequestTerminalSize{
				Cols: to.Int32Ptr(12),
				Rows: to.Int32Ptr(12),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
}
```
