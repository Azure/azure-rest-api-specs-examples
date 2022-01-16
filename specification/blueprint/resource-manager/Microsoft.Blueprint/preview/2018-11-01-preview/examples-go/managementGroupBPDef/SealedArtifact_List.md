Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fblueprint%2Farmblueprint%2Fv0.2.0/sdk/resourcemanager/blueprint/armblueprint/README.md) on how to add the SDK to your project and authenticate.

```go
package armblueprint_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/blueprint/armblueprint"
)

// x-ms-original-file: specification/blueprint/resource-manager/Microsoft.Blueprint/preview/2018-11-01-preview/examples/managementGroupBPDef/SealedArtifact_List.json
func ExamplePublishedArtifactsClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armblueprint.NewPublishedArtifactsClient(cred, nil)
	pager := client.List("<resource-scope>",
		"<blueprint-name>",
		"<version-id>",
		nil)
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
