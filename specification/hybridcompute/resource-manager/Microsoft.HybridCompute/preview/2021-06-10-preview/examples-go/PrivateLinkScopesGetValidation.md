Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fhybridcompute%2Farmhybridcompute%2Fv0.2.0/sdk/resourcemanager/hybridcompute/armhybridcompute/README.md) on how to add the SDK to your project and authenticate.

```go
package armhybridcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridcompute/armhybridcompute"
)

// x-ms-original-file: specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2021-06-10-preview/examples/PrivateLinkScopesGetValidation.json
func ExamplePrivateLinkScopesClient_GetValidationDetails() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armhybridcompute.NewPrivateLinkScopesClient("<subscription-id>", cred, nil)
	res, err := client.GetValidationDetails(ctx,
		"<location>",
		"<private-link-scope-id>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.PrivateLinkScopesClientGetValidationDetailsResult)
}
```
