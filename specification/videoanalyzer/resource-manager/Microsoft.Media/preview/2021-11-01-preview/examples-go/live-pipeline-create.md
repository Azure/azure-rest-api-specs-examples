Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fvideoanalyzer%2Farmvideoanalyzer%2Fv0.4.1/sdk/resourcemanager/videoanalyzer/armvideoanalyzer/README.md) on how to add the SDK to your project and authenticate.

```go
package armvideoanalyzer_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/videoanalyzer/armvideoanalyzer"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/videoanalyzer/resource-manager/Microsoft.Media/preview/2021-11-01-preview/examples/live-pipeline-create.json
func ExampleLivePipelinesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armvideoanalyzer.NewLivePipelinesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<live-pipeline-name>",
		armvideoanalyzer.LivePipeline{
			Properties: &armvideoanalyzer.LivePipelineProperties{
				Description: to.Ptr("<description>"),
				BitrateKbps: to.Ptr[int32](500),
				Parameters: []*armvideoanalyzer.ParameterDefinition{
					{
						Name:  to.Ptr("<name>"),
						Value: to.Ptr("<value>"),
					}},
				TopologyName: to.Ptr("<topology-name>"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
```
