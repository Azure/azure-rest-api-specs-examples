Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fdesktopvirtualization%2Farmdesktopvirtualization%2Fv0.2.0/sdk/resourcemanager/desktopvirtualization/armdesktopvirtualization/README.md) on how to add the SDK to your project and authenticate.

```go
package armdesktopvirtualization_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/desktopvirtualization/armdesktopvirtualization"
)

// x-ms-original-file: specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/preview/2021-09-03-preview/examples/Workspace_Update.json
func ExampleWorkspacesClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdesktopvirtualization.NewWorkspacesClient("<subscription-id>", cred, nil)
	res, err := client.Update(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		&armdesktopvirtualization.WorkspacesClientUpdateOptions{Workspace: &armdesktopvirtualization.WorkspacePatch{
			Properties: &armdesktopvirtualization.WorkspacePatchProperties{
				Description:         to.StringPtr("<description>"),
				FriendlyName:        to.StringPtr("<friendly-name>"),
				PublicNetworkAccess: armdesktopvirtualization.PublicNetworkAccess("Enabled").ToPtr(),
			},
			Tags: map[string]*string{
				"tag1": to.StringPtr("value1"),
				"tag2": to.StringPtr("value2"),
			},
		},
		})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.WorkspacesClientUpdateResult)
}
```
