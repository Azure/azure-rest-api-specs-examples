Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fchaos%2Farmchaos%2Fv0.1.0/sdk/resourcemanager/chaos/armchaos/README.md) on how to add the SDK to your project and authenticate.

```go
package armchaos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/chaos/armchaos"
)

// x-ms-original-file: specification/chaos/resource-manager/Microsoft.Chaos/preview/2021-09-15-preview/examples/ListTargets.json
func ExampleTargetsClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armchaos.NewTargetsClient("<subscription-id>", cred, nil)
	pager := client.List("<resource-group-name>",
		"<parent-provider-namespace>",
		"<parent-resource-type>",
		"<parent-resource-name>",
		&armchaos.TargetsListOptions{ContinuationToken: to.StringPtr("<continuation-token>")})
	for pager.NextPage(ctx) {
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("Target.ID: %s\n", *v.ID)
		}
	}
}
```
