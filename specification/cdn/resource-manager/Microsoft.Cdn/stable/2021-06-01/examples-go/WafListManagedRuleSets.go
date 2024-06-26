package armcdn_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c767823fdfd9d5e96bad245e3ea4d14d94a716bb/specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/WafListManagedRuleSets.json
func ExampleManagedRuleSetsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcdn.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewManagedRuleSetsClient().NewListPager(nil)
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
		// page.ManagedRuleSetDefinitionList = armcdn.ManagedRuleSetDefinitionList{
		// 	Value: []*armcdn.ManagedRuleSetDefinition{
		// 		{
		// 			Name: to.Ptr("DefaultRuleSet_1.0"),
		// 			Type: to.Ptr("Microsoft.Cdn/cdnwebapplicationfirewallmanagedrulesets"),
		// 			ID: to.Ptr("/subscriptions/subid/providers/Microsoft.Cdn/CdnWebApplicationFirewallManagedRuleSets"),
		// 			Properties: &armcdn.ManagedRuleSetDefinitionProperties{
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				RuleGroups: []*armcdn.ManagedRuleGroupDefinition{
		// 					{
		// 						Description: to.Ptr("Description for rule group 1."),
		// 						RuleGroupName: to.Ptr("Group1"),
		// 						Rules: []*armcdn.ManagedRuleDefinition{
		// 							{
		// 								Description: to.Ptr("Generic managed web application firewall rule."),
		// 								RuleID: to.Ptr("GROUP1-0001"),
		// 							},
		// 							{
		// 								Description: to.Ptr("Generic managed web application firewall rule."),
		// 								RuleID: to.Ptr("GROUP1-0002"),
		// 						}},
		// 					},
		// 					{
		// 						Description: to.Ptr("Description for rule group 2."),
		// 						RuleGroupName: to.Ptr("Group2"),
		// 						Rules: []*armcdn.ManagedRuleDefinition{
		// 							{
		// 								Description: to.Ptr("Generic managed web application firewall rule."),
		// 								RuleID: to.Ptr("GROUP2-0001"),
		// 						}},
		// 				}},
		// 				RuleSetType: to.Ptr("DefaultRuleSet"),
		// 				RuleSetVersion: to.Ptr("preview-1.0"),
		// 			},
		// 			SKU: &armcdn.SKU{
		// 				Name: to.Ptr(armcdn.SKUNameStandardMicrosoft),
		// 			},
		// 	}},
		// }
	}
}
