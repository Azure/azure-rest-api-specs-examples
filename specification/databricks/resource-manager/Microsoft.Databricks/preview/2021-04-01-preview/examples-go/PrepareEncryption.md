Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fdatabricks%2Farmdatabricks%2Fv0.3.0/sdk/resourcemanager/databricks/armdatabricks/README.md) on how to add the SDK to your project and authenticate.

```go
package armdatabricks_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databricks/armdatabricks"
)

// x-ms-original-file: specification/databricks/resource-manager/Microsoft.Databricks/preview/2021-04-01-preview/examples/PrepareEncryption.json
func ExampleWorkspacesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdatabricks.NewWorkspacesClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		armdatabricks.Workspace{
			Location: to.StringPtr("<location>"),
			Properties: &armdatabricks.WorkspaceProperties{
				ManagedResourceGroupID: to.StringPtr("<managed-resource-group-id>"),
				Parameters: &armdatabricks.WorkspaceCustomParameters{
					PrepareEncryption: &armdatabricks.WorkspaceCustomBooleanParameter{
						Value: to.BoolPtr(true),
					},
				},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.WorkspacesClientCreateOrUpdateResult)
}
```
