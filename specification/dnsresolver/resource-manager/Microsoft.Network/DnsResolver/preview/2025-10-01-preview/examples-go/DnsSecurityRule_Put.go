package armdnsresolver_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dnsresolver/armdnsresolver/v2"
)

// Generated from example definition: 2025-10-01-preview/DnsSecurityRule_Put.json
func ExampleDNSSecurityRulesClient_BeginCreateOrUpdate_upsertDnsSecurityRule() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdnsresolver.NewClientFactory("abdd4249-9f34-4cc6-8e42-c2e32110603e", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDNSSecurityRulesClient().BeginCreateOrUpdate(ctx, "sampleResourceGroup", "sampleDnsResolverPolicy", "sampleDnsSecurityRule", armdnsresolver.DNSSecurityRule{
		Location: to.Ptr("westus2"),
		Properties: &armdnsresolver.DNSSecurityRuleProperties{
			Action: &armdnsresolver.DNSSecurityRuleAction{
				ActionType: to.Ptr(armdnsresolver.ActionTypeBlock),
			},
			DNSResolverDomainLists: []*armdnsresolver.SubResource{
				{
					ID: to.Ptr("/subscriptions/abdd4249-9f34-4cc6-8e42-c2e32110603e/resourceGroups/sampleResourceGroup/providers/Microsoft.Network/dnsResolverDomainLists/sampleDnsResolverDomainList"),
				},
			},
			DNSSecurityRuleState: to.Ptr(armdnsresolver.DNSSecurityRuleStateEnabled),
			Priority:             to.Ptr[int32](100),
		},
		Tags: map[string]*string{
			"key1": to.Ptr("value1"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armdnsresolver.DNSSecurityRulesClientCreateOrUpdateResponse{
	// 	DNSSecurityRule: &armdnsresolver.DNSSecurityRule{
	// 		Name: to.Ptr("sampleDnsSecurityRule"),
	// 		Type: to.Ptr("Microsoft.Network/dnsResolverPolicies/dnsSecurityRules"),
	// 		Etag: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 		ID: to.Ptr("/subscriptions/abdd4249-9f34-4cc6-8e42-c2e32110603e/resourceGroups/sampleResourceGroup/providers/Microsoft.Network/dnsResolverPolicies/sampleDnsResolverPolicy/dnsSecurityRules/sampleDnsSecurityRule"),
	// 		Location: to.Ptr("westus2"),
	// 		Properties: &armdnsresolver.DNSSecurityRuleProperties{
	// 			Action: &armdnsresolver.DNSSecurityRuleAction{
	// 				ActionType: to.Ptr(armdnsresolver.ActionTypeBlock),
	// 			},
	// 			DNSResolverDomainLists: []*armdnsresolver.SubResource{
	// 				{
	// 					ID: to.Ptr("/subscriptions/abdd4249-9f34-4cc6-8e42-c2e32110603e/resourceGroups/sampleResourceGroup/providers/Microsoft.Network/dnsResolverDomainLists/sampleDnsResolverDomainList"),
	// 				},
	// 			},
	// 			DNSSecurityRuleState: to.Ptr(armdnsresolver.DNSSecurityRuleStateEnabled),
	// 			Priority: to.Ptr[int32](100),
	// 			ProvisioningState: to.Ptr(armdnsresolver.ProvisioningStateSucceeded),
	// 		},
	// 		SystemData: &armdnsresolver.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-01T01:01:01.1075056Z"); return t}()),
	// 			CreatedByType: to.Ptr(armdnsresolver.CreatedByTypeApplication),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-02T02:03:01.1974346Z"); return t}()),
	// 			LastModifiedByType: to.Ptr(armdnsresolver.CreatedByTypeApplication),
	// 		},
	// 		Tags: map[string]*string{
	// 			"key1": to.Ptr("value1"),
	// 		},
	// 	},
	// }
}
