Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fcontainerregistry%2Farmcontainerregistry%2Fv0.5.0/sdk/resourcemanager/containerregistry/armcontainerregistry/README.md) on how to add the SDK to your project and authenticate.

```go
package armcontainerregistry_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerregistry/armcontainerregistry"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2019-06-01-preview/examples/TaskRunsCreate.json
func ExampleTaskRunsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armcontainerregistry.NewTaskRunsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreate(ctx,
		"<resource-group-name>",
		"<registry-name>",
		"<task-run-name>",
		armcontainerregistry.TaskRun{
			Properties: &armcontainerregistry.TaskRunProperties{
				ForceUpdateTag: to.Ptr("<force-update-tag>"),
				RunRequest: &armcontainerregistry.EncodedTaskRunRequest{
					Type:                 to.Ptr("<type>"),
					Credentials:          &armcontainerregistry.Credentials{},
					EncodedTaskContent:   to.Ptr("<encoded-task-content>"),
					EncodedValuesContent: to.Ptr("<encoded-values-content>"),
					Platform: &armcontainerregistry.PlatformProperties{
						Architecture: to.Ptr(armcontainerregistry.ArchitectureAmd64),
						OS:           to.Ptr(armcontainerregistry.OSLinux),
					},
					Values: []*armcontainerregistry.SetValue{},
				},
			},
		},
		&armcontainerregistry.TaskRunsClientBeginCreateOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
