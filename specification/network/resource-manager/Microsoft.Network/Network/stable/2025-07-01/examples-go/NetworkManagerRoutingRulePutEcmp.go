package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v10"
)

// Generated from example definition: 2025-07-01/NetworkManagerRoutingRulePutEcmp.json
func ExampleRoutingRulesClient_CreateOrUpdate_createAnEcmpRoutingRule() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRoutingRulesClient().CreateOrUpdate(ctx, "rg1", "testNetworkManager", "myTestRoutingConfig", "testRuleCollection", "SampleEcmpRoutingRule", armnetwork.RoutingRule{
		Properties: &armnetwork.RoutingRulePropertiesFormat{
			Description: to.Ptr("This is a sample ECMP routing rule with multiple next hop IP addresses"),
			Destination: &armnetwork.RoutingRuleRouteDestination{
				Type:               to.Ptr(armnetwork.RoutingRuleDestinationTypeAddressPrefix),
				DestinationAddress: to.Ptr("10.0.0.0/16"),
			},
			NextHop: &armnetwork.RoutingRuleNextHop{
				NextHopType:    to.Ptr(armnetwork.RoutingRuleNextHopTypeVirtualAppliance),
				NextHopAddress: to.Ptr("10.1.0.4,10.1.0.5,10.1.0.6"),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armnetwork.RoutingRulesClientCreateOrUpdateResponse{
	// 	RoutingRule: armnetwork.RoutingRule{
	// 		Name: to.Ptr("SampleEcmpRoutingRule"),
	// 		Type: to.Ptr("Microsoft.Network/networkManagers/routingConfigurations/ruleCollections/rules"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/networkManagers/testNetworkManager/routingConfigurations/myTestRoutingConfig/ruleCollections/testRuleCollection/rules/SampleEcmpRoutingRule"),
	// 		Properties: &armnetwork.RoutingRulePropertiesFormat{
	// 			Description: to.Ptr("This is a sample ECMP routing rule with multiple next hop IP addresses"),
	// 			Destination: &armnetwork.RoutingRuleRouteDestination{
	// 				Type: to.Ptr(armnetwork.RoutingRuleDestinationTypeAddressPrefix),
	// 				DestinationAddress: to.Ptr("10.0.0.0/16"),
	// 			},
	// 			NextHop: &armnetwork.RoutingRuleNextHop{
	// 				NextHopType: to.Ptr(armnetwork.RoutingRuleNextHopTypeVirtualAppliance),
	// 				NextHopAddress: to.Ptr("10.1.0.4,10.1.0.5,10.1.0.6"),
	// 			},
	// 			ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 			ResourceGUID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 		},
	// 		SystemData: &armnetwork.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-11T18:52:27Z"); return t}()),
	// 			CreatedBy: to.Ptr("b69a9388-9488-4534-b470-7ec6d41beef5"),
	// 			CreatedByType: to.Ptr(armnetwork.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-11T18:52:27Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("b69a9388-9488-4534-b470-7ec6d41beef5"),
	// 			LastModifiedByType: to.Ptr(armnetwork.CreatedByTypeUser),
	// 		},
	// 	},
	// }
}
