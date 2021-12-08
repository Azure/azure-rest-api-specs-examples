Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fdevops%2Farmdevops%2Fv0.1.0/sdk/resourcemanager/devops/armdevops/README.md) on how to add the SDK to your project and authenticate.

```go
package armdevops_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devops/armdevops"
)

// x-ms-original-file: specification/devops/resource-manager/Microsoft.DevOps/preview/2019-07-01-preview/examples/UpdateAzurePipeline.json
func ExamplePipelinesClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdevops.NewPipelinesClient("<subscription-id>", cred, nil)
	res, err := client.Update(ctx,
		"<resource-group-name>",
		"<pipeline-name>",
		armdevops.PipelineUpdateParameters{
			Tags: map[string]*string{
				"tagKey": to.StringPtr("tagvalue"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Pipeline.ID: %s\n", *res.ID)
}
```
