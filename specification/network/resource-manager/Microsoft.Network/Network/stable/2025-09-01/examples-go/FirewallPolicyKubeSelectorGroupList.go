package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v11"
)

// Generated from example definition: 2025-09-01/FirewallPolicyKubeSelectorGroupList.json
func ExampleFirewallPolicyKubeSelectorGroupsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewFirewallPolicyKubeSelectorGroupsClient().NewListPager("rg1", "firewallPolicy", nil)
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
		// page = armnetwork.FirewallPolicyKubeSelectorGroupsClientListResponse{
		// 	FirewallPolicyKubeSelectorGroupListResult: armnetwork.FirewallPolicyKubeSelectorGroupListResult{
		// 		Value: []*armnetwork.FirewallPolicyKubeSelectorGroup{
		// 			{
		// 				Name: to.Ptr("kubeSelectorGroup1"),
		// 				Etag: to.Ptr("W/\"72090554-7e3b-43f2-80ad-99a9020dcb11\""),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/firewallPolicies/firewallPolicy/kubeSelectorGroups/kubeSelectorGroup1"),
		// 				Properties: &armnetwork.FirewallPolicyKubeSelectorGroupProperties{
		// 					PodSelector: &armnetwork.KubeLabelSelector{
		// 						MatchLabels: map[string]*string{
		// 							"app": to.Ptr("web"),
		// 							"env": to.Ptr("production"),
		// 						},
		// 						MatchExpressions: []*armnetwork.LabelSelectorExpression{
		// 							{
		// 								Key: to.Ptr("tier"),
		// 								Operator: to.Ptr(armnetwork.LabelSelectorOperatorIn),
		// 								Values: []*string{
		// 									to.Ptr("frontend"),
		// 									to.Ptr("backend"),
		// 								},
		// 							},
		// 						},
		// 					},
		// 					NamespaceSelector: &armnetwork.KubeLabelSelector{
		// 						MatchLabels: map[string]*string{
		// 							"kubernetes.io/metadata.name": to.Ptr("production"),
		// 						},
		// 					},
		// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
