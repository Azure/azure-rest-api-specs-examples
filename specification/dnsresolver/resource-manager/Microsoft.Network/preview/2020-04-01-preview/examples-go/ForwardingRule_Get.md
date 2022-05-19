Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fdnsresolver%2Farmdnsresolver%2Fv0.4.0/sdk/resourcemanager/dnsresolver/armdnsresolver/README.md) on how to add the SDK to your project and authenticate.

```go
package armdnsresolver_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dnsresolver/armdnsresolver"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/dnsresolver/resource-manager/Microsoft.Network/preview/2020-04-01-preview/examples/ForwardingRule_Get.json
func ExampleForwardingRulesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdnsresolver.NewForwardingRulesClient("abdd4249-9f34-4cc6-8e42-c2e32110603e", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx,
		"sampleResourceGroup",
		"sampleDnsForwardingRuleset",
		"sampleForwardingRule",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
```
