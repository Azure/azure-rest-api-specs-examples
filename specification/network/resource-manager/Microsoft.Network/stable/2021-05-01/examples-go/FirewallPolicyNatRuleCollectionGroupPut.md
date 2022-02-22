Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fnetwork%2Farmnetwork%2Fv0.3.1/sdk/resourcemanager/network/armnetwork/README.md) on how to add the SDK to your project and authenticate.

```go
package armnetwork_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
)

// x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/FirewallPolicyNatRuleCollectionGroupPut.json
func ExampleFirewallPolicyRuleCollectionGroupsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armnetwork.NewFirewallPolicyRuleCollectionGroupsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<firewall-policy-name>",
		"<rule-collection-group-name>",
		armnetwork.FirewallPolicyRuleCollectionGroup{
			Properties: &armnetwork.FirewallPolicyRuleCollectionGroupProperties{
				Priority: to.Int32Ptr(100),
				RuleCollections: []armnetwork.FirewallPolicyRuleCollectionClassification{
					&armnetwork.FirewallPolicyNatRuleCollection{
						Name:               to.StringPtr("<name>"),
						Priority:           to.Int32Ptr(100),
						RuleCollectionType: armnetwork.FirewallPolicyRuleCollectionType("FirewallPolicyNatRuleCollection").ToPtr(),
						Action: &armnetwork.FirewallPolicyNatRuleCollectionAction{
							Type: armnetwork.FirewallPolicyNatRuleCollectionActionType("DNAT").ToPtr(),
						},
						Rules: []armnetwork.FirewallPolicyRuleClassification{
							&armnetwork.NatRule{
								Name:     to.StringPtr("<name>"),
								RuleType: armnetwork.FirewallPolicyRuleType("NatRule").ToPtr(),
								DestinationAddresses: []*string{
									to.StringPtr("152.23.32.23")},
								DestinationPorts: []*string{
									to.StringPtr("8080")},
								IPProtocols: []*armnetwork.FirewallPolicyRuleNetworkProtocol{
									armnetwork.FirewallPolicyRuleNetworkProtocol("TCP").ToPtr(),
									armnetwork.FirewallPolicyRuleNetworkProtocol("UDP").ToPtr()},
								SourceAddresses: []*string{
									to.StringPtr("2.2.2.2")},
								SourceIPGroups: []*string{},
								TranslatedFqdn: to.StringPtr("<translated-fqdn>"),
								TranslatedPort: to.StringPtr("<translated-port>"),
							}},
					}},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.FirewallPolicyRuleCollectionGroupsClientCreateOrUpdateResult)
}
```
