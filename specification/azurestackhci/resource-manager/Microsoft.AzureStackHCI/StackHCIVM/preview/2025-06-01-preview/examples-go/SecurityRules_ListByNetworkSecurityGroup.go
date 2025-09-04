package armazurestackhcivm_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azurestackhci/armazurestackhcivm"
)

// Generated from example definition: 2025-06-01-preview/SecurityRules_ListByNetworkSecurityGroup.json
func ExampleSecurityRulesClient_NewListByNetworkSecurityGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armazurestackhcivm.NewClientFactory("fd3c3665-1729-4b7b-9a38-238e83b0f98b", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSecurityRulesClient().NewListByNetworkSecurityGroupPager("testrg", "testnsg", nil)
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
		// page = armazurestackhcivm.SecurityRulesClientListByNetworkSecurityGroupResponse{
		// 	SecurityRuleListResult: armazurestackhcivm.SecurityRuleListResult{
		// 		Value: []*armazurestackhcivm.SecurityRule{
		// 			{
		// 				Name: to.Ptr("rule1"),
		// 				ID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.AzureStackHCI/networkSecurityGroups/testnsg/securityRule/rule1"),
		// 				Properties: &armazurestackhcivm.SecurityRuleProperties{
		// 					Access: to.Ptr(armazurestackhcivm.SecurityRuleAccessAllow),
		// 					DestinationAddressPrefixes: []*string{
		// 						to.Ptr("*"),
		// 					},
		// 					DestinationPortRanges: []*string{
		// 						to.Ptr("80"),
		// 					},
		// 					Direction: to.Ptr(armazurestackhcivm.SecurityRuleDirectionInbound),
		// 					Priority: to.Ptr[int32](130),
		// 					SourceAddressPrefixes: []*string{
		// 						to.Ptr("*"),
		// 					},
		// 					SourcePortRanges: []*string{
		// 						to.Ptr("*"),
		// 					},
		// 					Protocol: to.Ptr(armazurestackhcivm.SecurityRuleProtocolAsterisk),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
