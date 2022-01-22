Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fvideoanalyzer%2Farmvideoanalyzer%2Fv0.2.0/sdk/resourcemanager/videoanalyzer/armvideoanalyzer/README.md) on how to add the SDK to your project and authenticate.

```go
package armvideoanalyzer_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/videoanalyzer/armvideoanalyzer"
)

// x-ms-original-file: specification/videoanalyzer/resource-manager/Microsoft.Media/preview/2021-11-01-preview/examples/live-pipeline-patch.json
func ExampleLivePipelinesClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armvideoanalyzer.NewLivePipelinesClient("<subscription-id>", cred, nil)
	res, err := client.Update(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<live-pipeline-name>",
		armvideoanalyzer.LivePipelineUpdate{
			Properties: &armvideoanalyzer.LivePipelinePropertiesUpdate{
				Description: to.StringPtr("<description>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.LivePipelinesClientUpdateResult)
}
```
