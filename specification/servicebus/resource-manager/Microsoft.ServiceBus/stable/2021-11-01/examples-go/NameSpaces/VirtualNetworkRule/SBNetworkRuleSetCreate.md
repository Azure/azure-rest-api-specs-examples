Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fservicebus%2Farmservicebus%2Fv0.3.1/sdk/resourcemanager/servicebus/armservicebus/README.md) on how to add the SDK to your project and authenticate.

```go
package armservicebus_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicebus/armservicebus"
)

// x-ms-original-file: specification/servicebus/resource-manager/Microsoft.ServiceBus/stable/2021-11-01/examples/NameSpaces/VirtualNetworkRule/SBNetworkRuleSetCreate.json
func ExampleNamespacesClient_CreateOrUpdateNetworkRuleSet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armservicebus.NewNamespacesClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdateNetworkRuleSet(ctx,
		"<resource-group-name>",
		"<namespace-name>",
		armservicebus.NetworkRuleSet{
			Properties: &armservicebus.NetworkRuleSetProperties{
				DefaultAction: armservicebus.DefaultAction("Deny").ToPtr(),
				IPRules: []*armservicebus.NWRuleSetIPRules{
					{
						Action: armservicebus.NetworkRuleIPAction("Allow").ToPtr(),
						IPMask: to.StringPtr("<ipmask>"),
					},
					{
						Action: armservicebus.NetworkRuleIPAction("Allow").ToPtr(),
						IPMask: to.StringPtr("<ipmask>"),
					},
					{
						Action: armservicebus.NetworkRuleIPAction("Allow").ToPtr(),
						IPMask: to.StringPtr("<ipmask>"),
					},
					{
						Action: armservicebus.NetworkRuleIPAction("Allow").ToPtr(),
						IPMask: to.StringPtr("<ipmask>"),
					},
					{
						Action: armservicebus.NetworkRuleIPAction("Allow").ToPtr(),
						IPMask: to.StringPtr("<ipmask>"),
					}},
				VirtualNetworkRules: []*armservicebus.NWRuleSetVirtualNetworkRules{
					{
						IgnoreMissingVnetServiceEndpoint: to.BoolPtr(true),
						Subnet: &armservicebus.Subnet{
							ID: to.StringPtr("<id>"),
						},
					},
					{
						IgnoreMissingVnetServiceEndpoint: to.BoolPtr(false),
						Subnet: &armservicebus.Subnet{
							ID: to.StringPtr("<id>"),
						},
					},
					{
						IgnoreMissingVnetServiceEndpoint: to.BoolPtr(false),
						Subnet: &armservicebus.Subnet{
							ID: to.StringPtr("<id>"),
						},
					}},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.NamespacesClientCreateOrUpdateNetworkRuleSetResult)
}
```
