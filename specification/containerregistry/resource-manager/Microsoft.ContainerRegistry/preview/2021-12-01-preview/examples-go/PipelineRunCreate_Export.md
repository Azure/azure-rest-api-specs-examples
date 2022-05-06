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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2021-12-01-preview/examples/PipelineRunCreate_Export.json
func ExamplePipelineRunsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armcontainerregistry.NewPipelineRunsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreate(ctx,
		"<resource-group-name>",
		"<registry-name>",
		"<pipeline-run-name>",
		armcontainerregistry.PipelineRun{
			Properties: &armcontainerregistry.PipelineRunProperties{
				Request: &armcontainerregistry.PipelineRunRequest{
					Artifacts: []*string{
						to.Ptr("sourceRepository/hello-world"),
						to.Ptr("sourceRepository2@sha256:00000000000000000000000000000000000")},
					PipelineResourceID: to.Ptr("<pipeline-resource-id>"),
					Target: &armcontainerregistry.PipelineRunTargetProperties{
						Name: to.Ptr("<name>"),
						Type: to.Ptr(armcontainerregistry.PipelineRunTargetTypeAzureStorageBlob),
					},
				},
			},
		},
		&armcontainerregistry.PipelineRunsClientBeginCreateOptions{ResumeToken: ""})
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
