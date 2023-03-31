package armdnsresolver_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dnsresolver/armdnsresolver"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b749953e21e5c3f275d839862323920ef7bf716e/specification/dnsresolver/resource-manager/Microsoft.Network/stable/2022-07-01/examples/ForwardingRule_List.json
func ExampleForwardingRulesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdnsresolver.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewForwardingRulesClient().NewListPager("sampleResourceGroup", "sampleDnsForwardingRuleset", &armdnsresolver.ForwardingRulesClientListOptions{Top: nil})
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
		// page.ForwardingRuleListResult = armdnsresolver.ForwardingRuleListResult{
		// 	Value: []*armdnsresolver.ForwardingRule{
		// 		{
		// 			Name: to.Ptr("sampleForwardingRule"),
		// 			Type: to.Ptr("Microsoft.Network/dnsForwardingRulesets/forwardingRules"),
		// 			ID: to.Ptr("/subscriptions/abdd4249-9f34-4cc6-8e42-c2e32110603e/resourceGroups/sampleResourceGroup/providers/Microsoft.Network/dnsForwardingRulesets/sampleDnsForwardingRuleset/forwardingRules/sampleForwardingRule"),
		// 			Etag: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 			Properties: &armdnsresolver.ForwardingRuleProperties{
		// 				DomainName: to.Ptr("contoso.com."),
		// 				ForwardingRuleState: to.Ptr(armdnsresolver.ForwardingRuleStateEnabled),
		// 				Metadata: map[string]*string{
		// 					"additionalProp1": to.Ptr("value1"),
		// 				},
		// 				ProvisioningState: to.Ptr(armdnsresolver.ProvisioningStateSucceeded),
		// 				TargetDNSServers: []*armdnsresolver.TargetDNSServer{
		// 					{
		// 						IPAddress: to.Ptr("10.0.0.1"),
		// 						Port: to.Ptr[int32](53),
		// 					},
		// 					{
		// 						IPAddress: to.Ptr("10.0.0.2"),
		// 						Port: to.Ptr[int32](53),
		// 				}},
		// 			},
		// 			SystemData: &armdnsresolver.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-01T01:01:01.1075056Z"); return t}()),
		// 				CreatedByType: to.Ptr(armdnsresolver.CreatedByTypeApplication),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-02T02:03:01.1974346Z"); return t}()),
		// 				LastModifiedByType: to.Ptr(armdnsresolver.CreatedByTypeApplication),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("sampleForwardingRule1"),
		// 			Type: to.Ptr("Microsoft.Network/dnsForwardingRulesets/forwardingRules"),
		// 			ID: to.Ptr("/subscriptions/abdd4249-9f34-4cc6-8e42-c2e32110603e/resourceGroups/sampleResourceGroup/providers/Microsoft.Network/dnsForwardingRulesets/sampleDnsForwardingRuleset/forwardingRules/sampleForwardingRule1"),
		// 			Etag: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 			Properties: &armdnsresolver.ForwardingRuleProperties{
		// 				DomainName: to.Ptr("foobar.com."),
		// 				ForwardingRuleState: to.Ptr(armdnsresolver.ForwardingRuleStateEnabled),
		// 				Metadata: map[string]*string{
		// 					"additionalProp1": to.Ptr("value1"),
		// 				},
		// 				ProvisioningState: to.Ptr(armdnsresolver.ProvisioningStateSucceeded),
		// 				TargetDNSServers: []*armdnsresolver.TargetDNSServer{
		// 					{
		// 						IPAddress: to.Ptr("10.0.0.3"),
		// 						Port: to.Ptr[int32](53),
		// 					},
		// 					{
		// 						IPAddress: to.Ptr("10.0.0.4"),
		// 						Port: to.Ptr[int32](53),
		// 				}},
		// 			},
		// 			SystemData: &armdnsresolver.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-01T01:01:01.1075056Z"); return t}()),
		// 				CreatedByType: to.Ptr(armdnsresolver.CreatedByTypeApplication),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-02T02:03:01.1974346Z"); return t}()),
		// 				LastModifiedByType: to.Ptr(armdnsresolver.CreatedByTypeApplication),
		// 			},
		// 	}},
		// }
	}
}
