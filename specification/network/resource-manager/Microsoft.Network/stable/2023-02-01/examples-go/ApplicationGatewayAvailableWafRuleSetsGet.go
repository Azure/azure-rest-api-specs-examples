package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9975d3476c05bcc6bd9535ad3dfb564e6a168fa5/specification/network/resource-manager/Microsoft.Network/stable/2023-02-01/examples/ApplicationGatewayAvailableWafRuleSetsGet.json
func ExampleApplicationGatewaysClient_ListAvailableWafRuleSets() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewApplicationGatewaysClient().ListAvailableWafRuleSets(ctx, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ApplicationGatewayAvailableWafRuleSetsResult = armnetwork.ApplicationGatewayAvailableWafRuleSetsResult{
	// 	Value: []*armnetwork.ApplicationGatewayFirewallRuleSet{
	// 		{
	// 			Name: to.Ptr("OWASP_3.0"),
	// 			Type: to.Ptr("Microsoft.Network/applicationGatewayAvailableWafRuleSets"),
	// 			ID: to.Ptr("/subscriptions//resourceGroups//providers/Microsoft.Network/applicationGatewayAvailableWafRuleSets/"),
	// 			Properties: &armnetwork.ApplicationGatewayFirewallRuleSetPropertiesFormat{
	// 				ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 				RuleGroups: []*armnetwork.ApplicationGatewayFirewallRuleGroup{
	// 					{
	// 						Description: to.Ptr(""),
	// 						RuleGroupName: to.Ptr("General"),
	// 						Rules: []*armnetwork.ApplicationGatewayFirewallRule{
	// 							{
	// 								Description: to.Ptr("Possible Multipart Unmatched Boundary."),
	// 								RuleID: to.Ptr[int32](200004),
	// 								RuleIDString: to.Ptr("200004"),
	// 						}},
	// 				}},
	// 				RuleSetType: to.Ptr("OWASP"),
	// 				RuleSetVersion: to.Ptr("3.0"),
	// 			},
	// 	}},
	// }
}
