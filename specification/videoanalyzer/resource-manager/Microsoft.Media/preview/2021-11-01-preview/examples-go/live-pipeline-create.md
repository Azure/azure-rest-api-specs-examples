Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fvideoanalyzer%2Farmvideoanalyzer%2Fv0.1.0/sdk/resourcemanager/videoanalyzer/armvideoanalyzer/README.md) on how to add the SDK to your project and authenticate.

```go
package armvideoanalyzer_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/videoanalyzer/armvideoanalyzer"
)

// x-ms-original-file: specification/videoanalyzer/resource-manager/Microsoft.Media/preview/2021-11-01-preview/examples/live-pipeline-create.json
func ExampleLivePipelinesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armvideoanalyzer.NewLivePipelinesClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<live-pipeline-name>",
		armvideoanalyzer.LivePipeline{
			Properties: &armvideoanalyzer.LivePipelineProperties{
				Description: to.StringPtr("<description>"),
				BitrateKbps: to.Int32Ptr(500),
				Parameters: []*armvideoanalyzer.ParameterDefinition{
					{
						Name:  to.StringPtr("<name>"),
						Value: to.StringPtr("<value>"),
					}},
				TopologyName: to.StringPtr("<topology-name>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("LivePipeline.ID: %s\n", *res.ID)
}
```
