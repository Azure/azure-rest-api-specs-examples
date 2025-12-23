package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v8"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b24c97bfc136b01dd46a1c8ddcecd0bb5a1ab152/specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/NetworkWatcherNetworkConfigurationDiagnostic.json
func ExampleWatchersClient_BeginGetNetworkConfigurationDiagnostic() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewWatchersClient().BeginGetNetworkConfigurationDiagnostic(ctx, "rg1", "nw1", armnetwork.ConfigurationDiagnosticParameters{
		Profiles: []*armnetwork.ConfigurationDiagnosticProfile{
			{
				Destination:     to.Ptr("12.11.12.14"),
				DestinationPort: to.Ptr("12100"),
				Direction:       to.Ptr(armnetwork.DirectionInbound),
				Source:          to.Ptr("10.1.0.4"),
				Protocol:        to.Ptr("TCP"),
			}},
		TargetResourceID: to.Ptr("/subscriptions/subid/resourceGroups/rg2/providers/Microsoft.Compute/virtualMachines/vm1"),
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
	// res.ConfigurationDiagnosticResponse = armnetwork.ConfigurationDiagnosticResponse{
	// 	Results: []*armnetwork.ConfigurationDiagnosticResult{
	// 		{
	// 			NetworkSecurityGroupResult: &armnetwork.SecurityGroupResult{
	// 				EvaluatedNetworkSecurityGroups: []*armnetwork.EvaluatedNetworkSecurityGroup{
	// 					{
	// 						AppliedTo: to.Ptr("/subscriptions/subid/resourceGroups/rg2/providers/Microsoft.Network/virtualNetworks/vnet/subnets/AppSubnet"),
	// 						MatchedRule: &armnetwork.MatchedRule{
	// 							Action: to.Ptr("Allow"),
	// 							RuleName: to.Ptr("UserRule_fe_rule"),
	// 						},
	// 						NetworkSecurityGroupID: to.Ptr("/subscriptions/subid/resourceGroups/rg2/providers/Microsoft.Network/networkSecurityGroups/nsg1"),
	// 						RulesEvaluationResult: []*armnetwork.SecurityRulesEvaluationResult{
	// 							{
	// 								Name: to.Ptr("UserRule_Cleanuptool-Allow-100"),
	// 								DestinationMatched: to.Ptr(true),
	// 								DestinationPortMatched: to.Ptr(false),
	// 								ProtocolMatched: to.Ptr(true),
	// 								SourceMatched: to.Ptr(false),
	// 								SourcePortMatched: to.Ptr(true),
	// 							},
	// 							{
	// 								Name: to.Ptr("UserRule_Cleanuptool-Allow-101"),
	// 								DestinationMatched: to.Ptr(true),
	// 								DestinationPortMatched: to.Ptr(false),
	// 								ProtocolMatched: to.Ptr(true),
	// 								SourceMatched: to.Ptr(true),
	// 								SourcePortMatched: to.Ptr(true),
	// 							},
	// 							{
	// 								Name: to.Ptr("UserRule_Cleanuptool-Allow-102"),
	// 								DestinationMatched: to.Ptr(true),
	// 								DestinationPortMatched: to.Ptr(false),
	// 								ProtocolMatched: to.Ptr(true),
	// 								SourceMatched: to.Ptr(false),
	// 								SourcePortMatched: to.Ptr(true),
	// 							},
	// 							{
	// 								Name: to.Ptr("UserRule_Cleanuptool-Deny-103"),
	// 								DestinationMatched: to.Ptr(true),
	// 								DestinationPortMatched: to.Ptr(false),
	// 								ProtocolMatched: to.Ptr(true),
	// 								SourceMatched: to.Ptr(true),
	// 								SourcePortMatched: to.Ptr(true),
	// 							},
	// 							{
	// 								Name: to.Ptr("UserRule_fe_rule"),
	// 								DestinationMatched: to.Ptr(true),
	// 								DestinationPortMatched: to.Ptr(true),
	// 								ProtocolMatched: to.Ptr(true),
	// 								SourceMatched: to.Ptr(true),
	// 								SourcePortMatched: to.Ptr(true),
	// 						}},
	// 					},
	// 					{
	// 						AppliedTo: to.Ptr("/subscriptions/subid/resourceGroups/rg2/providers/Microsoft.Network/virtualNetworks/vnet/networkInterfaces/nic"),
	// 						MatchedRule: &armnetwork.MatchedRule{
	// 							Action: to.Ptr("Allow"),
	// 							RuleName: to.Ptr("UserRule_fe_rule"),
	// 						},
	// 						NetworkSecurityGroupID: to.Ptr("/subscriptions/61cc8a98-a8be-4bfe-a04e-0b461f93fe35/resourceGroups/NwRgCentralUSEUAP_copy/providers/Microsoft.Network/networkSecurityGroups/AppNSG"),
	// 						RulesEvaluationResult: []*armnetwork.SecurityRulesEvaluationResult{
	// 							{
	// 								Name: to.Ptr("UserRule_fe_rule"),
	// 								DestinationMatched: to.Ptr(true),
	// 								DestinationPortMatched: to.Ptr(true),
	// 								ProtocolMatched: to.Ptr(true),
	// 								SourceMatched: to.Ptr(true),
	// 								SourcePortMatched: to.Ptr(true),
	// 						}},
	// 				}},
	// 				SecurityRuleAccessResult: to.Ptr(armnetwork.SecurityRuleAccessAllow),
	// 			},
	// 			Profile: &armnetwork.ConfigurationDiagnosticProfile{
	// 				Destination: to.Ptr("12.11.12.14"),
	// 				DestinationPort: to.Ptr("12100"),
	// 				Direction: to.Ptr(armnetwork.DirectionInbound),
	// 				Source: to.Ptr("10.1.0.4"),
	// 				Protocol: to.Ptr("TCP"),
	// 			},
	// 	}},
	// }
}
