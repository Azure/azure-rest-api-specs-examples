Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fazurestackhci%2Farmazurestackhci%2Fv0.2.0/sdk/resourcemanager/azurestackhci/armazurestackhci/README.md) on how to add the SDK to your project and authenticate.

```go
package armazurestackhci_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azurestackhci/armazurestackhci"
)

// x-ms-original-file: specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/stable/2021-09-01/examples/PatchExtension.json
func ExampleExtensionsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armazurestackhci.NewExtensionsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginUpdate(ctx,
		"<resource-group-name>",
		"<cluster-name>",
		"<arc-setting-name>",
		"<extension-name>",
		armazurestackhci.Extension{
			Properties: &armazurestackhci.ExtensionProperties{
				ExtensionParameters: &armazurestackhci.ExtensionParameters{
					Type:      to.StringPtr("<type>"),
					Publisher: to.StringPtr("<publisher>"),
					Settings: map[string]interface{}{
						"workspaceId": "xx",
					},
					TypeHandlerVersion: to.StringPtr("<type-handler-version>"),
				},
			},
		},
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
