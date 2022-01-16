Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fcontainerregistry%2Farmcontainerregistry%2Fv0.3.0/sdk/resourcemanager/containerregistry/armcontainerregistry/README.md) on how to add the SDK to your project and authenticate.

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

// x-ms-original-file: specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2019-06-01-preview/examples/TaskRunsCreate.json
func ExampleTaskRunsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcontainerregistry.NewTaskRunsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreate(ctx,
		"<resource-group-name>",
		"<registry-name>",
		"<task-run-name>",
		armcontainerregistry.TaskRun{
			Properties: &armcontainerregistry.TaskRunProperties{
				ForceUpdateTag: to.StringPtr("<force-update-tag>"),
				RunRequest: &armcontainerregistry.EncodedTaskRunRequest{
					Type:                 to.StringPtr("<type>"),
					Credentials:          &armcontainerregistry.Credentials{},
					EncodedTaskContent:   to.StringPtr("<encoded-task-content>"),
					EncodedValuesContent: to.StringPtr("<encoded-values-content>"),
					Platform: &armcontainerregistry.PlatformProperties{
						Architecture: armcontainerregistry.Architecture("amd64").ToPtr(),
						OS:           armcontainerregistry.OS("Linux").ToPtr(),
					},
					Values: []*armcontainerregistry.SetValue{},
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
	log.Printf("Response result: %#v\n", res.TaskRunsClientCreateResult)
}
```
