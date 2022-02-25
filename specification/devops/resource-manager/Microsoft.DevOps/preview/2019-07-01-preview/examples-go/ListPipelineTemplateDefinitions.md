Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fdevops%2Farmdevops%2Fv0.2.1/sdk/resourcemanager/devops/armdevops/README.md) on how to add the SDK to your project and authenticate.

```go
package armdevops_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devops/armdevops"
)

// x-ms-original-file: specification/devops/resource-manager/Microsoft.DevOps/preview/2019-07-01-preview/examples/ListPipelineTemplateDefinitions.json
func ExamplePipelineTemplateDefinitionsClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdevops.NewPipelineTemplateDefinitionsClient(cred, nil)
	pager := client.List(nil)
	for {
		nextResult := pager.NextPage(ctx)
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		if !nextResult {
			break
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("Pager result: %#v\n", v)
		}
	}
}
```
