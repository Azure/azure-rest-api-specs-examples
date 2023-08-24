package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/baac183ffa684d94f697f0fc6f480e02cfb00f3d/specification/network/resource-manager/Microsoft.Network/stable/2023-04-01/examples/NatRuleGet.json
func ExampleNatRulesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewNatRulesClient().Get(ctx, "rg1", "gateway1", "natRule1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.VPNGatewayNatRule = armnetwork.VPNGatewayNatRule{
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/vpnGateways/gateway1/natRules/natRule1"),
	// 	Name: to.Ptr("natRule1"),
	// 	Etag: to.Ptr("w/\\00000000-0000-0000-0000-000000000000\\"),
	// 	Properties: &armnetwork.VPNGatewayNatRuleProperties{
	// 		Type: to.Ptr(armnetwork.VPNNatRuleTypeStatic),
	// 		EgressVPNSiteLinkConnections: []*armnetwork.SubResource{
	// 		},
	// 		ExternalMappings: []*armnetwork.VPNNatRuleMapping{
	// 		},
	// 		IngressVPNSiteLinkConnections: []*armnetwork.SubResource{
	// 		},
	// 		InternalMappings: []*armnetwork.VPNNatRuleMapping{
	// 			{
	// 				AddressSpace: to.Ptr("10.4.0.0/24"),
	// 		}},
	// 		Mode: to.Ptr(armnetwork.VPNNatRuleModeEgressSnat),
	// 		ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 	},
	// }
}
