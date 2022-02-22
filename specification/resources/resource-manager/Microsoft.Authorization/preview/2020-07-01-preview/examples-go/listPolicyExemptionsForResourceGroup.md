Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fresources%2Farmpolicy%2Fv0.3.1/sdk/resourcemanager/resources/armpolicy/README.md) on how to add the SDK to your project and authenticate.

```go
package armpolicy_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armpolicy"
)

// x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/preview/2020-07-01-preview/examples/listPolicyExemptionsForResourceGroup.json
func ExampleExemptionsClient_ListForResourceGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armpolicy.NewExemptionsClient("<subscription-id>", cred, nil)
	pager := client.ListForResourceGroup("<resource-group-name>",
		&armpolicy.ExemptionsClientListForResourceGroupOptions{Filter: to.StringPtr("<filter>")})
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
