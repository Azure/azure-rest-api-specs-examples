package armdnsresolver_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dnsresolver/armdnsresolver/v2"
)

// Generated from example definition: 2025-10-01-preview/DnsForwardingRuleset_ListBySubscription.json
func ExampleDNSForwardingRulesetsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdnsresolver.NewClientFactory("abdd4249-9f34-4cc6-8e42-c2e32110603e", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDNSForwardingRulesetsClient().NewListPager(nil)
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
		// page = armdnsresolver.DNSForwardingRulesetsClientListResponse{
		// 	DNSForwardingRulesetListResult: armdnsresolver.DNSForwardingRulesetListResult{
		// 		NextLink: to.Ptr("https://management.azure.com/api/mresolver/subscriptions/abdd4249-9f34-4cc6-8e42-c2e32110603e/providers/Microsoft.Network/dnsForwardingRulesets?$skipToken=skipToken&api-version=2025-10-01-preview"),
		// 		Value: []*armdnsresolver.DNSForwardingRuleset{
		// 			{
		// 				Name: to.Ptr("sampleDnsForwardingRuleset"),
		// 				Type: to.Ptr("Microsoft.Network/dnsForwardingRulesets"),
		// 				Etag: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 				ID: to.Ptr("/subscriptions/abdd4249-9f34-4cc6-8e42-c2e32110603e/resourceGroups/sampleResourceGroup/providers/Microsoft.Network/dnsForwardingRulesets/sampleDnsForwardingRuleset"),
		// 				Location: to.Ptr("westus2"),
		// 				Properties: &armdnsresolver.DNSForwardingRulesetProperties{
		// 					DNSResolverOutboundEndpoints: []*armdnsresolver.SubResource{
		// 						{
		// 							ID: to.Ptr("/subscriptions/abdd4249-9f34-4cc6-8e42-c2e32110603e/resourceGroups/sampleResourceGroup/providers/Microsoft.Network/dnsResolvers/sampleDnsResolver/outboundEndpoints/sampleOutboundEndpoint0"),
		// 						},
		// 						{
		// 							ID: to.Ptr("/subscriptions/abdd4249-9f34-4cc6-8e42-c2e32110603e/resourceGroups/sampleResourceGroup/providers/Microsoft.Network/dnsResolvers/sampleDnsResolver/outboundEndpoints/sampleOutboundEndpoint1"),
		// 						},
		// 					},
		// 					ProvisioningState: to.Ptr(armdnsresolver.ProvisioningStateSucceeded),
		// 					ResourceGUID: to.Ptr("a7e1a32c-498c-401c-a805-5bc3518257b8"),
		// 				},
		// 				SystemData: &armdnsresolver.SystemData{
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-03T01:01:01.1075056Z"); return t}()),
		// 					CreatedByType: to.Ptr(armdnsresolver.CreatedByTypeApplication),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-04T02:03:01.1974346Z"); return t}()),
		// 					LastModifiedByType: to.Ptr(armdnsresolver.CreatedByTypeApplication),
		// 				},
		// 				Tags: map[string]*string{
		// 					"key1": to.Ptr("value1"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("sampleDnsForwardingRuleset"),
		// 				Type: to.Ptr("Microsoft.Network/dnsForwardingRulesets"),
		// 				Etag: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 				ID: to.Ptr("/subscriptions/abdd4249-9f34-4cc6-8e42-c2e32110603e/resourceGroups/sampleResourceGroup/providers/Microsoft.Network/dnsForwardingRulesets/sampleDnsForwardingRuleset1"),
		// 				Location: to.Ptr("westus2"),
		// 				Properties: &armdnsresolver.DNSForwardingRulesetProperties{
		// 					DNSResolverOutboundEndpoints: []*armdnsresolver.SubResource{
		// 						{
		// 							ID: to.Ptr("/subscriptions/abdd4249-9f34-4cc6-8e42-c2e32110603e/resourceGroups/sampleResourceGroup/providers/Microsoft.Network/dnsResolvers/sampleDnsResolver/outboundEndpoints/sampleOutboundEndpoint2"),
		// 						},
		// 						{
		// 							ID: to.Ptr("/subscriptions/abdd4249-9f34-4cc6-8e42-c2e32110603e/resourceGroups/sampleResourceGroup/providers/Microsoft.Network/dnsResolvers/sampleDnsResolver/outboundEndpoints/sampleOutboundEndpoint3"),
		// 						},
		// 					},
		// 					ProvisioningState: to.Ptr(armdnsresolver.ProvisioningStateSucceeded),
		// 					ResourceGUID: to.Ptr("a7e1a32c-498c-401c-a805-5bc3518257b8"),
		// 				},
		// 				SystemData: &armdnsresolver.SystemData{
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-04T01:01:01.1075056Z"); return t}()),
		// 					CreatedByType: to.Ptr(armdnsresolver.CreatedByTypeApplication),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-05T02:03:01.1974346Z"); return t}()),
		// 					LastModifiedByType: to.Ptr(armdnsresolver.CreatedByTypeApplication),
		// 				},
		// 				Tags: map[string]*string{
		// 					"key1": to.Ptr("value1"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
