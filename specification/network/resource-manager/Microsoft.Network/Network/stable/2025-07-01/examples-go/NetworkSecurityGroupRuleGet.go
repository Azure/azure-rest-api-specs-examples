package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v10"
)

// Generated from example definition: 2025-07-01/NetworkSecurityGroupRuleGet.json
func ExampleSecurityRulesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSecurityRulesClient().Get(ctx, "rg1", "testnsg", "rule1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armnetwork.SecurityRulesClientGetResponse{
	// 	SecurityRule: armnetwork.SecurityRule{
	// 		Name: to.Ptr("rule1"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/networkSecurityGroups/testnsg/securityRules/rule1"),
	// 		Properties: &armnetwork.SecurityRulePropertiesFormat{
	// 			Access: to.Ptr(armnetwork.SecurityRuleAccessAllow),
	// 			DestinationAddressPrefix: to.Ptr("*"),
	// 			DestinationPortRange: to.Ptr("80"),
	// 			Direction: to.Ptr(armnetwork.SecurityRuleDirectionInbound),
	// 			Priority: to.Ptr[int32](130),
	// 			ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 			SourceAddressPrefix: to.Ptr("*"),
	// 			SourcePortRange: to.Ptr("*"),
	// 			Protocol: to.Ptr(armnetwork.SecurityRuleProtocolAsterisk),
	// 		},
	// 	},
	// }
}
