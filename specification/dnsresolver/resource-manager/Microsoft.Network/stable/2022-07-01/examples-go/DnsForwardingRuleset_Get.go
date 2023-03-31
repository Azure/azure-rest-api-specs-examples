package armdnsresolver_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dnsresolver/armdnsresolver"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b749953e21e5c3f275d839862323920ef7bf716e/specification/dnsresolver/resource-manager/Microsoft.Network/stable/2022-07-01/examples/DnsForwardingRuleset_Get.json
func ExampleDNSForwardingRulesetsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdnsresolver.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDNSForwardingRulesetsClient().Get(ctx, "sampleResourceGroup", "sampleDnsForwardingRuleset", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DNSForwardingRuleset = armdnsresolver.DNSForwardingRuleset{
	// 	Name: to.Ptr("sampleDnsForwardingRuleset"),
	// 	Type: to.Ptr("Microsoft.Network/dnsForwardingRulesets"),
	// 	ID: to.Ptr("/subscriptions/abdd4249-9f34-4cc6-8e42-c2e32110603e/resourceGroups/sampleResourceGroup/providers/Microsoft.Network/dnsForwardingRulesets/sampleDnsForwardingRuleset"),
	// 	Location: to.Ptr("westus2"),
	// 	Tags: map[string]*string{
	// 		"key1": to.Ptr("value1"),
	// 	},
	// 	Etag: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 	Properties: &armdnsresolver.DNSForwardingRulesetProperties{
	// 		DNSResolverOutboundEndpoints: []*armdnsresolver.SubResource{
	// 			{
	// 				ID: to.Ptr("/subscriptions/abdd4249-9f34-4cc6-8e42-c2e32110603e/resourceGroups/sampleResourceGroup/providers/Microsoft.Network/dnsResolvers/sampleDnsResolver/outboundEndpoints/sampleOutboundEndpoint0"),
	// 			},
	// 			{
	// 				ID: to.Ptr("/subscriptions/abdd4249-9f34-4cc6-8e42-c2e32110603e/resourceGroups/sampleResourceGroup/providers/Microsoft.Network/dnsResolvers/sampleDnsResolver/outboundEndpoints/sampleOutboundEndpoint1"),
	// 		}},
	// 		ProvisioningState: to.Ptr(armdnsresolver.ProvisioningStateSucceeded),
	// 		ResourceGUID: to.Ptr("a7e1a32c-498c-401c-a805-5bc3518257b8"),
	// 	},
	// 	SystemData: &armdnsresolver.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-03T01:01:01.1075056Z"); return t}()),
	// 		CreatedByType: to.Ptr(armdnsresolver.CreatedByTypeApplication),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-04T02:03:01.1974346Z"); return t}()),
	// 		LastModifiedByType: to.Ptr(armdnsresolver.CreatedByTypeApplication),
	// 	},
	// }
}
