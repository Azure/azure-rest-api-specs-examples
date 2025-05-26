package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c58fa033619b12c7cfa8a0ec5a9bf03bb18869ab/specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/DefaultSecurityRuleGet.json
func ExampleDefaultSecurityRulesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDefaultSecurityRulesClient().Get(ctx, "testrg", "nsg1", "AllowVnetInBound", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SecurityRule = armnetwork.SecurityRule{
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/testrg/providers/Microsoft.Network/networkSecurityGroups/nsg1/defaultSecurityRules/AllowVnetInBound"),
	// 	Name: to.Ptr("AllowVnetInBound"),
	// 	Properties: &armnetwork.SecurityRulePropertiesFormat{
	// 		Description: to.Ptr("Allow inbound traffic from all VMs in VNET"),
	// 		Access: to.Ptr(armnetwork.SecurityRuleAccessAllow),
	// 		DestinationAddressPrefix: to.Ptr("VirtualNetwork"),
	// 		DestinationAddressPrefixes: []*string{
	// 		},
	// 		DestinationPortRange: to.Ptr("*"),
	// 		DestinationPortRanges: []*string{
	// 		},
	// 		Direction: to.Ptr(armnetwork.SecurityRuleDirectionInbound),
	// 		Priority: to.Ptr[int32](65000),
	// 		ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 		SourceAddressPrefix: to.Ptr("VirtualNetwork"),
	// 		SourceAddressPrefixes: []*string{
	// 		},
	// 		SourcePortRange: to.Ptr("*"),
	// 		SourcePortRanges: []*string{
	// 		},
	// 		Protocol: to.Ptr(armnetwork.SecurityRuleProtocolAsterisk),
	// 	},
	// }
}
