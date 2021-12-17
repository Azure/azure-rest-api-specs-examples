Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fhybridcompute%2Farmhybridcompute%2Fv0.1.0/sdk/resourcemanager/hybridcompute/armhybridcompute/README.md) on how to add the SDK to your project and authenticate.

```go
package armhybridcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridcompute/armhybridcompute"
)

// x-ms-original-file: specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2021-06-10-preview/examples/PrivateLinkScopesUpdateTagsOnly.json
func ExamplePrivateLinkScopesClient_UpdateTags() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armhybridcompute.NewPrivateLinkScopesClient("<subscription-id>", cred, nil)
	res, err := client.UpdateTags(ctx,
		"<resource-group-name>",
		"<scope-name>",
		armhybridcompute.TagsResource{},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("HybridComputePrivateLinkScope.ID: %s\n", *res.ID)
}
```
