package armdnsresolver_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dnsresolver/armdnsresolver"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d0d3a9b4fe0fce880fded7a617e71f84406bacbd/specification/dnsresolver/resource-manager/Microsoft.Network/stable/2025-05-01/examples/DnsSecurityRule_Patch.json
func ExampleDNSSecurityRulesClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdnsresolver.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDNSSecurityRulesClient().BeginUpdate(ctx, "sampleResourceGroup", "sampleDnsResolverPolicy", "sampleDnsSecurityRule", armdnsresolver.DNSSecurityRulePatch{
		Properties: &armdnsresolver.DNSSecurityRulePatchProperties{
			DNSSecurityRuleState: to.Ptr(armdnsresolver.DNSSecurityRuleStateDisabled),
		},
		Tags: map[string]*string{
			"key1": to.Ptr("value1"),
		},
	}, &armdnsresolver.DNSSecurityRulesClientBeginUpdateOptions{IfMatch: nil})
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
	// res.DNSSecurityRule = armdnsresolver.DNSSecurityRule{
	// 	Name: to.Ptr("sampleDnsSecurityRule"),
	// 	Type: to.Ptr("Microsoft.Network/dnsResolverPolicies/dnsSecurityRules"),
	// 	ID: to.Ptr("/subscriptions/abdd4249-9f34-4cc6-8e42-c2e32110603e/resourceGroups/sampleResourceGroup/providers/Microsoft.Network/dnsResolverPolicies/sampleDnsResolverPolicy/dnsSecurityRules/sampleDnsSecurityRule"),
	// 	SystemData: &armdnsresolver.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-01T01:01:01.107Z"); return t}()),
	// 		CreatedByType: to.Ptr(armdnsresolver.CreatedByTypeApplication),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-02T02:03:01.197Z"); return t}()),
	// 		LastModifiedByType: to.Ptr(armdnsresolver.CreatedByTypeApplication),
	// 	},
	// 	Location: to.Ptr("westus2"),
	// 	Tags: map[string]*string{
	// 		"key1": to.Ptr("value1"),
	// 	},
	// 	Etag: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 	Properties: &armdnsresolver.DNSSecurityRuleProperties{
	// 		Action: &armdnsresolver.DNSSecurityRuleAction{
	// 			ActionType: to.Ptr(armdnsresolver.ActionTypeAllow),
	// 		},
	// 		DNSResolverDomainLists: []*armdnsresolver.SubResource{
	// 			{
	// 				ID: to.Ptr("/subscriptions/abdd4249-9f34-4cc6-8e42-c2e32110603e/resourceGroups/sampleResourceGroup/providers/Microsoft.Network/dnsResolverDomainLists/sampleDnsResolverDomainList2"),
	// 		}},
	// 		DNSSecurityRuleState: to.Ptr(armdnsresolver.DNSSecurityRuleStateDisabled),
	// 		Priority: to.Ptr[int32](100),
	// 		ProvisioningState: to.Ptr(armdnsresolver.ProvisioningStateSucceeded),
	// 	},
	// }
}
