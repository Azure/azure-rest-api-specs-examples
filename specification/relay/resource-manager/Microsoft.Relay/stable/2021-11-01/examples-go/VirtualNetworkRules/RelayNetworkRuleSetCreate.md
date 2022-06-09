```go
package armrelay_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/relay/armrelay"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/relay/resource-manager/Microsoft.Relay/stable/2021-11-01/examples/VirtualNetworkRules/RelayNetworkRuleSetCreate.json
func ExampleNamespacesClient_CreateOrUpdateNetworkRuleSet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armrelay.NewNamespacesClient("Subscription", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdateNetworkRuleSet(ctx,
		"ResourceGroup",
		"example-RelayNamespace-6019",
		armrelay.NetworkRuleSet{
			Properties: &armrelay.NetworkRuleSetProperties{
				DefaultAction: to.Ptr(armrelay.DefaultActionDeny),
				IPRules: []*armrelay.NWRuleSetIPRules{
					{
						Action: to.Ptr(armrelay.NetworkRuleIPActionAllow),
						IPMask: to.Ptr("1.1.1.1"),
					},
					{
						Action: to.Ptr(armrelay.NetworkRuleIPActionAllow),
						IPMask: to.Ptr("1.1.1.2"),
					},
					{
						Action: to.Ptr(armrelay.NetworkRuleIPActionAllow),
						IPMask: to.Ptr("1.1.1.3"),
					},
					{
						Action: to.Ptr(armrelay.NetworkRuleIPActionAllow),
						IPMask: to.Ptr("1.1.1.4"),
					},
					{
						Action: to.Ptr(armrelay.NetworkRuleIPActionAllow),
						IPMask: to.Ptr("1.1.1.5"),
					}},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Frelay%2Farmrelay%2Fv1.0.0/sdk/resourcemanager/relay/armrelay/README.md) on how to add the SDK to your project and authenticate.
