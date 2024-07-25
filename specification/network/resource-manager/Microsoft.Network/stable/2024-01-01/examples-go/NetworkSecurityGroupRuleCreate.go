package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v6"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/81a4ee5a83ae38620c0e1404793caffe005d26e4/specification/network/resource-manager/Microsoft.Network/stable/2024-01-01/examples/NetworkSecurityGroupRuleCreate.json
func ExampleSecurityRulesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewSecurityRulesClient().BeginCreateOrUpdate(ctx, "rg1", "testnsg", "rule1", armnetwork.SecurityRule{
		Properties: &armnetwork.SecurityRulePropertiesFormat{
			Access:                   to.Ptr(armnetwork.SecurityRuleAccessDeny),
			DestinationAddressPrefix: to.Ptr("11.0.0.0/8"),
			DestinationPortRange:     to.Ptr("8080"),
			Direction:                to.Ptr(armnetwork.SecurityRuleDirectionOutbound),
			Priority:                 to.Ptr[int32](100),
			SourceAddressPrefix:      to.Ptr("10.0.0.0/8"),
			SourcePortRange:          to.Ptr("*"),
			Protocol:                 to.Ptr(armnetwork.SecurityRuleProtocolAsterisk),
		},
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
	// res.SecurityRule = armnetwork.SecurityRule{
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/networkSecurityGroups/testnsg/securityRules/rule1"),
	// 	Name: to.Ptr("rule1"),
	// 	Properties: &armnetwork.SecurityRulePropertiesFormat{
	// 		Access: to.Ptr(armnetwork.SecurityRuleAccessDeny),
	// 		DestinationAddressPrefix: to.Ptr("11.0.0.0/8"),
	// 		DestinationPortRange: to.Ptr("8080"),
	// 		Direction: to.Ptr(armnetwork.SecurityRuleDirectionOutbound),
	// 		Priority: to.Ptr[int32](100),
	// 		ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 		SourceAddressPrefix: to.Ptr("10.0.0.0/8"),
	// 		SourcePortRange: to.Ptr("*"),
	// 		Protocol: to.Ptr(armnetwork.SecurityRuleProtocolAsterisk),
	// 	},
	// }
}
