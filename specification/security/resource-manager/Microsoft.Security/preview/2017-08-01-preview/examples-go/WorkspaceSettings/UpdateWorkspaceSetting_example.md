Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsecurity%2Farmsecurity%2Fv0.3.1/sdk/resourcemanager/security/armsecurity/README.md) on how to add the SDK to your project and authenticate.

```go
package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2017-08-01-preview/examples/WorkspaceSettings/UpdateWorkspaceSetting_example.json
func ExampleWorkspaceSettingsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsecurity.NewWorkspaceSettingsClient("<subscription-id>", cred, nil)
	res, err := client.Update(ctx,
		"<workspace-setting-name>",
		armsecurity.WorkspaceSetting{
			Name: to.StringPtr("<name>"),
			Type: to.StringPtr("<type>"),
			ID:   to.StringPtr("<id>"),
			Properties: &armsecurity.WorkspaceSettingProperties{
				WorkspaceID: to.StringPtr("<workspace-id>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.WorkspaceSettingsClientUpdateResult)
}
```
