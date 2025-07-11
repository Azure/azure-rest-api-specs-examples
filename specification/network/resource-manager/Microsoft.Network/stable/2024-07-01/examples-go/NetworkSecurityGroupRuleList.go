package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c58fa033619b12c7cfa8a0ec5a9bf03bb18869ab/specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/NetworkSecurityGroupRuleList.json
func ExampleSecurityRulesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSecurityRulesClient().NewListPager("rg1", "testnsg", nil)
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
		// page.SecurityRuleListResult = armnetwork.SecurityRuleListResult{
		// 	Value: []*armnetwork.SecurityRule{
		// 		{
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/networkSecurityGroups/testnsg/securityRules/rule1"),
		// 			Name: to.Ptr("rule1"),
		// 			Properties: &armnetwork.SecurityRulePropertiesFormat{
		// 				Access: to.Ptr(armnetwork.SecurityRuleAccessAllow),
		// 				DestinationAddressPrefix: to.Ptr("*"),
		// 				DestinationPortRange: to.Ptr("80"),
		// 				Direction: to.Ptr(armnetwork.SecurityRuleDirectionInbound),
		// 				Priority: to.Ptr[int32](130),
		// 				ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 				SourceAddressPrefix: to.Ptr("*"),
		// 				SourcePortRange: to.Ptr("*"),
		// 				Protocol: to.Ptr(armnetwork.SecurityRuleProtocolAsterisk),
		// 			},
		// 	}},
		// }
	}
}
