package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c58fa033619b12c7cfa8a0ec5a9bf03bb18869ab/specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/GetApplicationGatewayWafDynamicManifests.json
func ExampleApplicationGatewayWafDynamicManifestsClient_NewGetPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewApplicationGatewayWafDynamicManifestsClient().NewGetPager("westus", nil)
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
		// page.ApplicationGatewayWafDynamicManifestResultList = armnetwork.ApplicationGatewayWafDynamicManifestResultList{
		// 	Value: []*armnetwork.ApplicationGatewayWafDynamicManifestResult{
		// 		{
		// 			Name: to.Ptr("default"),
		// 			Type: to.Ptr("Microsoft.Network/applicationGatewayWafDynamicManifest"),
		// 			ID: to.Ptr("/subscriptions/subid/providers/Microsoft.Network/applicationGatewayWafDynamicManifests/default"),
		// 			Properties: &armnetwork.ApplicationGatewayWafDynamicManifestPropertiesResult{
		// 				AvailableRuleSets: []*armnetwork.ApplicationGatewayFirewallManifestRuleSet{
		// 					{
		// 						RuleGroups: []*armnetwork.ApplicationGatewayFirewallRuleGroup{
		// 							{
		// 								Description: to.Ptr(""),
		// 								RuleGroupName: to.Ptr("General"),
		// 								Rules: []*armnetwork.ApplicationGatewayFirewallRule{
		// 									{
		// 										Description: to.Ptr("Failed to Parse Request Body."),
		// 										Action: to.Ptr(armnetwork.ApplicationGatewayWafRuleActionTypesAnomalyScoring),
		// 										RuleID: to.Ptr[int32](200002),
		// 										RuleIDString: to.Ptr("200002"),
		// 										State: to.Ptr(armnetwork.ApplicationGatewayWafRuleStateTypesEnabled),
		// 									},
		// 									{
		// 										Description: to.Ptr("Multipart Request Body Strict Validation."),
		// 										Action: to.Ptr(armnetwork.ApplicationGatewayWafRuleActionTypesAnomalyScoring),
		// 										RuleID: to.Ptr[int32](200003),
		// 										RuleIDString: to.Ptr("200003"),
		// 										State: to.Ptr(armnetwork.ApplicationGatewayWafRuleStateTypesEnabled),
		// 									},
		// 									{
		// 										Description: to.Ptr("Possible Multipart Unmatched Boundary."),
		// 										Action: to.Ptr(armnetwork.ApplicationGatewayWafRuleActionTypesAnomalyScoring),
		// 										RuleID: to.Ptr[int32](200004),
		// 										RuleIDString: to.Ptr("200004"),
		// 										State: to.Ptr(armnetwork.ApplicationGatewayWafRuleStateTypesEnabled),
		// 								}},
		// 						}},
		// 						RuleSetType: to.Ptr("OWASP"),
		// 						RuleSetVersion: to.Ptr("3.2"),
		// 						Status: to.Ptr(armnetwork.ApplicationGatewayRuleSetStatusOptions("0")),
		// 						Tiers: []*armnetwork.ApplicationGatewayTierTypes{
		// 							to.Ptr(armnetwork.ApplicationGatewayTierTypesWAFV2)},
		// 					}},
		// 					DefaultRuleSet: &armnetwork.DefaultRuleSetPropertyFormat{
		// 						RuleSetType: to.Ptr("OWASP"),
		// 						RuleSetVersion: to.Ptr("3.2"),
		// 					},
		// 				},
		// 		}},
		// 	}
	}
}
