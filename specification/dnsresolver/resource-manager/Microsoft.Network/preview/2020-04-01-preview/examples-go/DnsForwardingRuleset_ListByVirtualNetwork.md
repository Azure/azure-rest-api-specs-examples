Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fdnsresolver%2Farmdnsresolver%2Fv0.1.0/sdk/resourcemanager/dnsresolver/armdnsresolver/README.md) on how to add the SDK to your project and authenticate.

```go
package armdnsresolver_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dnsresolver/armdnsresolver"
)

// x-ms-original-file: specification/dnsresolver/resource-manager/Microsoft.Network/preview/2020-04-01-preview/examples/DnsForwardingRuleset_ListByVirtualNetwork.json
func ExampleDNSForwardingRulesetsClient_ListByVirtualNetwork() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdnsresolver.NewDNSForwardingRulesetsClient("<subscription-id>", cred, nil)
	pager := client.ListByVirtualNetwork("<resource-group-name>",
		"<virtual-network-name>",
		&armdnsresolver.DNSForwardingRulesetsClientListByVirtualNetworkOptions{Top: nil})
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
