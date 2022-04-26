Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fvideoanalyzer%2Farmvideoanalyzer%2Fv0.4.0/sdk/resourcemanager/videoanalyzer/armvideoanalyzer/README.md) on how to add the SDK to your project and authenticate.

```go
package armvideoanalyzer_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/videoanalyzer/armvideoanalyzer"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/videoanalyzer/resource-manager/Microsoft.Media/preview/2021-11-01-preview/examples/pipeline-job-patch.json
func ExamplePipelineJobsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armvideoanalyzer.NewPipelineJobsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.Update(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<pipeline-job-name>",
		armvideoanalyzer.PipelineJobUpdate{
			Properties: &armvideoanalyzer.PipelineJobPropertiesUpdate{
				Description: to.Ptr("<description>"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
