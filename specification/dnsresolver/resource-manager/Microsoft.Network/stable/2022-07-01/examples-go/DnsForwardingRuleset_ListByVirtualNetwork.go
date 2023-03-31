package armdnsresolver_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dnsresolver/armdnsresolver"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b749953e21e5c3f275d839862323920ef7bf716e/specification/dnsresolver/resource-manager/Microsoft.Network/stable/2022-07-01/examples/DnsForwardingRuleset_ListByVirtualNetwork.json
func ExampleDNSForwardingRulesetsClient_NewListByVirtualNetworkPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdnsresolver.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDNSForwardingRulesetsClient().NewListByVirtualNetworkPager("sampleResourceGroup", "sampleVirtualNetwork", &armdnsresolver.DNSForwardingRulesetsClientListByVirtualNetworkOptions{Top: nil})
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.VirtualNetworkDNSForwardingRulesetListResult = armdnsresolver.VirtualNetworkDNSForwardingRulesetListResult{
		// 	Value: []*armdnsresolver.VirtualNetworkDNSForwardingRuleset{
		// 		{
		// 			ID: to.Ptr("/subscriptions/abdd4249-9f34-4cc6-8e42-c2e32110603e/resourceGroups/sampleResourceGroup/providers/Microsoft.Network/dnsForwardingRulesets/sampleDnsForwardingRuleset"),
		// 			Properties: &armdnsresolver.VirtualNetworkLinkSubResourceProperties{
		// 				VirtualNetworkLink: &armdnsresolver.SubResource{
		// 					ID: to.Ptr("/subscriptions/abdd4249-9f34-4cc6-8e42-c2e32110603e/resourceGroups/sampleResourceGroup/providers/Microsoft.Network/dnsForwardingRuleset/sampleDnsForwardingRuleset/virtualNetworkLinks/sampleVirtualNetworkLink"),
		// 				},
		// 			},
		// 	}},
		// }
	}
}
