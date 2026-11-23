package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v12"
)

// Generated from example definition: 2026-05-01/InboundSecurityRuleSourceIPProtocolPut.json
func ExampleInboundSecurityRuleClient_BeginCreateOrUpdate_createOrUpdateAPermanentInternetIngressRuleWithSourceIPProtocolLoadDistribution() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewInboundSecurityRuleClient().BeginCreateOrUpdate(ctx, "rg1", "nva", "rule1", armnetwork.InboundSecurityRule{
		Properties: &armnetwork.InboundSecurityRuleProperties{
			RuleType: to.Ptr(armnetwork.InboundSecurityRuleTypePermanent),
			Rules: []*armnetwork.InboundSecurityRules{
				{
					Name: to.Ptr("inboundRule1"),
					AppliesOn: []*string{
						to.Ptr("slbip1"),
					},
					DestinationPortRange: to.Ptr[int32](443),
					SourceAddressPrefix:  to.Ptr("50.20.121.5/32"),
					Protocol:             to.Ptr(armnetwork.InboundSecurityRulesProtocolTCP),
					LoadBalancerConfiguration: &armnetwork.VirtualApplianceLoadBalancerConfiguration{
						LoadDistribution:     to.Ptr(armnetwork.NetworkVirtualApplianceLoadDistributionSourceIPProtocol),
						IdleTimeoutInMinutes: to.Ptr[int32](15),
					},
				},
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armnetwork.InboundSecurityRuleClientCreateOrUpdateResponse{
	// 	InboundSecurityRule: armnetwork.InboundSecurityRule{
	// 		Name: to.Ptr("rule1"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/networkVirtualAppliances/nva/InboundSecurityRules/rule1"),
	// 		Properties: &armnetwork.InboundSecurityRuleProperties{
	// 			ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 			RuleType: to.Ptr(armnetwork.InboundSecurityRuleTypePermanent),
	// 			Rules: []*armnetwork.InboundSecurityRules{
	// 				{
	// 					Name: to.Ptr("inboundRule1"),
	// 					AppliesOn: []*string{
	// 						to.Ptr("slbip1"),
	// 					},
	// 					DestinationPortRange: to.Ptr[int32](443),
	// 					SourceAddressPrefix: to.Ptr("50.20.121.5/32"),
	// 					Protocol: to.Ptr(armnetwork.InboundSecurityRulesProtocolTCP),
	// 					LoadBalancerConfiguration: &armnetwork.VirtualApplianceLoadBalancerConfiguration{
	// 						LoadDistribution: to.Ptr(armnetwork.NetworkVirtualApplianceLoadDistributionSourceIPProtocol),
	// 						IdleTimeoutInMinutes: to.Ptr[int32](15),
	// 					},
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
