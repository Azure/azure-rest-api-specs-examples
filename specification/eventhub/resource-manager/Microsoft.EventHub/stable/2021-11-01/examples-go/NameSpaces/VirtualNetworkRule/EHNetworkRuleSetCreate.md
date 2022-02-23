Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Feventhub%2Farmeventhub%2Fv0.3.1/sdk/resourcemanager/eventhub/armeventhub/README.md) on how to add the SDK to your project and authenticate.

```go
package armeventhub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventhub/armeventhub"
)

// x-ms-original-file: specification/eventhub/resource-manager/Microsoft.EventHub/stable/2021-11-01/examples/NameSpaces/VirtualNetworkRule/EHNetworkRuleSetCreate.json
func ExampleNamespacesClient_CreateOrUpdateNetworkRuleSet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armeventhub.NewNamespacesClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdateNetworkRuleSet(ctx,
		"<resource-group-name>",
		"<namespace-name>",
		armeventhub.NetworkRuleSet{
			Properties: &armeventhub.NetworkRuleSetProperties{
				DefaultAction: armeventhub.DefaultAction("Deny").ToPtr(),
				IPRules: []*armeventhub.NWRuleSetIPRules{
					{
						Action: armeventhub.NetworkRuleIPAction("Allow").ToPtr(),
						IPMask: to.StringPtr("<ipmask>"),
					},
					{
						Action: armeventhub.NetworkRuleIPAction("Allow").ToPtr(),
						IPMask: to.StringPtr("<ipmask>"),
					},
					{
						Action: armeventhub.NetworkRuleIPAction("Allow").ToPtr(),
						IPMask: to.StringPtr("<ipmask>"),
					},
					{
						Action: armeventhub.NetworkRuleIPAction("Allow").ToPtr(),
						IPMask: to.StringPtr("<ipmask>"),
					},
					{
						Action: armeventhub.NetworkRuleIPAction("Allow").ToPtr(),
						IPMask: to.StringPtr("<ipmask>"),
					}},
				VirtualNetworkRules: []*armeventhub.NWRuleSetVirtualNetworkRules{
					{
						IgnoreMissingVnetServiceEndpoint: to.BoolPtr(true),
						Subnet: &armeventhub.Subnet{
							ID: to.StringPtr("<id>"),
						},
					},
					{
						IgnoreMissingVnetServiceEndpoint: to.BoolPtr(false),
						Subnet: &armeventhub.Subnet{
							ID: to.StringPtr("<id>"),
						},
					},
					{
						IgnoreMissingVnetServiceEndpoint: to.BoolPtr(false),
						Subnet: &armeventhub.Subnet{
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
