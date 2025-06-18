package armdnsresolver_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dnsresolver/armdnsresolver"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d0d3a9b4fe0fce880fded7a617e71f84406bacbd/specification/dnsresolver/resource-manager/Microsoft.Network/stable/2025-05-01/examples/ForwardingRule_Get.json
func ExampleForwardingRulesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdnsresolver.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewForwardingRulesClient().Get(ctx, "sampleResourceGroup", "sampleDnsForwardingRuleset", "sampleForwardingRule", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ForwardingRule = armdnsresolver.ForwardingRule{
	// 	Name: to.Ptr("sampleForwardingRule"),
	// 	Type: to.Ptr("Microsoft.Network/dnsForwardingRulesets/forwardingRules"),
	// 	ID: to.Ptr("/subscriptions/abdd4249-9f34-4cc6-8e42-c2e32110603e/resourceGroups/sampleResourceGroup/providers/Microsoft.Network/dnsForwardingRulesets/sampleDnsForwardingRuleset/forwardingRules/sampleForwardingRule"),
	// 	SystemData: &armdnsresolver.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-01T01:01:01.107Z"); return t}()),
	// 		CreatedByType: to.Ptr(armdnsresolver.CreatedByTypeApplication),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-02T02:03:01.197Z"); return t}()),
	// 		LastModifiedByType: to.Ptr(armdnsresolver.CreatedByTypeApplication),
	// 	},
	// 	Etag: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 	Properties: &armdnsresolver.ForwardingRuleProperties{
	// 		DomainName: to.Ptr("contoso.com."),
	// 		ForwardingRuleState: to.Ptr(armdnsresolver.ForwardingRuleStateEnabled),
	// 		Metadata: map[string]*string{
	// 			"additionalProp1": to.Ptr("value1"),
	// 		},
	// 		ProvisioningState: to.Ptr(armdnsresolver.ProvisioningStateSucceeded),
	// 		TargetDNSServers: []*armdnsresolver.TargetDNSServer{
	// 			{
	// 				IPAddress: to.Ptr("10.0.0.1"),
	// 				Port: to.Ptr[int32](53),
	// 			},
	// 			{
	// 				IPAddress: to.Ptr("10.0.0.2"),
	// 				Port: to.Ptr[int32](53),
	// 		}},
	// 	},
	// }
}
