package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v8"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b24c97bfc136b01dd46a1c8ddcecd0bb5a1ab152/specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/P2SVpnGatewayList.json
func ExampleP2SVPNGatewaysClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewP2SVPNGatewaysClient().NewListPager(nil)
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
		// page.ListP2SVPNGatewaysResult = armnetwork.ListP2SVPNGatewaysResult{
		// 	Value: []*armnetwork.P2SVPNGateway{
		// 		{
		// 			Name: to.Ptr("p2sVpnGateway1"),
		// 			Type: to.Ptr("Microsoft.Network/p2sVpnGateways"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/P2SvpnGateways/p2sVpnGateway1"),
		// 			Location: to.Ptr("West US"),
		// 			Etag: to.Ptr("w/\\00000000-0000-0000-0000-000000000000\\"),
		// 			Properties: &armnetwork.P2SVPNGatewayProperties{
		// 				CustomDNSServers: []*string{
		// 					to.Ptr("3.3.3.3")},
		// 					IsRoutingPreferenceInternet: to.Ptr(false),
		// 					P2SConnectionConfigurations: []*armnetwork.P2SConnectionConfiguration{
		// 						{
		// 							ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/p2sVpnGateways/p2sVpnGateway1/p2sConnectionConfigurations/P2SConnectionConfig1"),
		// 							Name: to.Ptr("P2SConnectionConfig1"),
		// 							Etag: to.Ptr("w/\\00000000-0000-0000-0000-000000000000\\"),
		// 							Properties: &armnetwork.P2SConnectionConfigurationProperties{
		// 								ConfigurationPolicyGroupAssociations: []*armnetwork.SubResource{
		// 									{
		// 										ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/vpnServerConfigurations/vpnServerConfiguration1/configurationPolicyGroups/policyGroup1"),
		// 								}},
		// 								EnableInternetSecurity: to.Ptr(true),
		// 								PreviousConfigurationPolicyGroupAssociations: []*armnetwork.VPNServerConfigurationPolicyGroup{
		// 									{
		// 										ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/vpnServerConfigurations/vpnServerConfiguration1/vpnServerConfigurationPolicyGroups/policyGroup1"),
		// 										Name: to.Ptr("policyGroup1"),
		// 										Properties: &armnetwork.VPNServerConfigurationPolicyGroupProperties{
		// 											IsDefault: to.Ptr(true),
		// 											PolicyMembers: []*armnetwork.VPNServerConfigurationPolicyGroupMember{
		// 												{
		// 													Name: to.Ptr("policy1"),
		// 													AttributeType: to.Ptr(armnetwork.VPNPolicyMemberAttributeTypeRadiusAzureGroupID),
		// 													AttributeValue: to.Ptr("6ad1bd08"),
		// 											}},
		// 											Priority: to.Ptr[int32](0),
		// 											ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 										},
		// 								}},
		// 								ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 								RoutingConfiguration: &armnetwork.RoutingConfiguration{
		// 									AssociatedRouteTable: &armnetwork.SubResource{
		// 										ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1/hubRouteTables/hubRouteTable1"),
		// 									},
		// 									PropagatedRouteTables: &armnetwork.PropagatedRouteTable{
		// 										IDs: []*armnetwork.SubResource{
		// 											{
		// 												ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1/hubRouteTables/hubRouteTable1"),
		// 											},
		// 											{
		// 												ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1/hubRouteTables/hubRouteTable2"),
		// 											},
		// 											{
		// 												ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1/hubRouteTables/hubRouteTable3"),
		// 										}},
		// 										Labels: []*string{
		// 											to.Ptr("label1"),
		// 											to.Ptr("label2")},
		// 										},
		// 										VnetRoutes: &armnetwork.VnetRoute{
		// 											StaticRoutes: []*armnetwork.StaticRoute{
		// 											},
		// 										},
		// 									},
		// 									VPNClientAddressPool: &armnetwork.AddressSpace{
		// 										AddressPrefixes: []*string{
		// 											to.Ptr("101.3.0.0/16")},
		// 										},
		// 									},
		// 							}},
		// 							ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 							VirtualHub: &armnetwork.SubResource{
		// 								ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1"),
		// 							},
		// 							VPNClientConnectionHealth: &armnetwork.VPNClientConnectionHealth{
		// 								AllocatedIPAddresses: []*string{
		// 									to.Ptr("1.1.1.1"),
		// 									to.Ptr("2.2.2.2")},
		// 									TotalEgressBytesTransferred: to.Ptr[int64](3000),
		// 									TotalIngressBytesTransferred: to.Ptr[int64](2000),
		// 									VPNClientConnectionsCount: to.Ptr[int32](2),
		// 								},
		// 								VPNGatewayScaleUnit: to.Ptr[int32](1),
		// 								VPNServerConfiguration: &armnetwork.SubResource{
		// 									ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/vpnServerConfigurations/vpnServerConfiguration1"),
		// 								},
		// 							},
		// 						},
		// 						{
		// 							Name: to.Ptr("p2sVpnGateway2"),
		// 							Type: to.Ptr("Microsoft.Network/p2sVpnGateways"),
		// 							ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/P2SvpnGateways/p2sVpnGateway2"),
		// 							Location: to.Ptr("West US"),
		// 							Etag: to.Ptr("w/\\00000000-0000-0000-0000-000000000000\\"),
		// 							Properties: &armnetwork.P2SVPNGatewayProperties{
		// 								CustomDNSServers: []*string{
		// 									to.Ptr("4.4.4.4")},
		// 									IsRoutingPreferenceInternet: to.Ptr(false),
		// 									P2SConnectionConfigurations: []*armnetwork.P2SConnectionConfiguration{
		// 										{
		// 											ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/p2sVpnGateways/p2sVpnGateway1/p2sConnectionConfigurations/P2SConnectionConfig1"),
		// 											Name: to.Ptr("P2SConnectionConfig1"),
		// 											Etag: to.Ptr("w/\\00000000-0000-0000-0000-000000000000\\"),
		// 											Properties: &armnetwork.P2SConnectionConfigurationProperties{
		// 												ConfigurationPolicyGroupAssociations: []*armnetwork.SubResource{
		// 													{
		// 														ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/vpnServerConfigurations/vpnServerConfiguration1/configurationPolicyGroups/policyGroup1"),
		// 												}},
		// 												EnableInternetSecurity: to.Ptr(true),
		// 												PreviousConfigurationPolicyGroupAssociations: []*armnetwork.VPNServerConfigurationPolicyGroup{
		// 													{
		// 														ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/vpnServerConfigurations/vpnServerConfiguration1/vpnServerConfigurationPolicyGroups/policyGroup1"),
		// 														Name: to.Ptr("policyGroup1"),
		// 														Properties: &armnetwork.VPNServerConfigurationPolicyGroupProperties{
		// 															IsDefault: to.Ptr(true),
		// 															PolicyMembers: []*armnetwork.VPNServerConfigurationPolicyGroupMember{
		// 																{
		// 																	Name: to.Ptr("policy1"),
		// 																	AttributeType: to.Ptr(armnetwork.VPNPolicyMemberAttributeTypeRadiusAzureGroupID),
		// 																	AttributeValue: to.Ptr("6ad1bd08"),
		// 															}},
		// 															Priority: to.Ptr[int32](0),
		// 															ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 														},
		// 												}},
		// 												ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 												RoutingConfiguration: &armnetwork.RoutingConfiguration{
		// 													AssociatedRouteTable: &armnetwork.SubResource{
		// 														ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub2/hubRouteTables/hubRouteTable1"),
		// 													},
		// 													PropagatedRouteTables: &armnetwork.PropagatedRouteTable{
		// 														IDs: []*armnetwork.SubResource{
		// 															{
		// 																ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub2/hubRouteTables/hubRouteTable1"),
		// 															},
		// 															{
		// 																ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub2/hubRouteTables/hubRouteTable2"),
		// 															},
		// 															{
		// 																ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub2/hubRouteTables/hubRouteTable3"),
		// 														}},
		// 														Labels: []*string{
		// 															to.Ptr("label1"),
		// 															to.Ptr("label2")},
		// 														},
		// 														VnetRoutes: &armnetwork.VnetRoute{
		// 															StaticRoutes: []*armnetwork.StaticRoute{
		// 															},
		// 														},
		// 													},
		// 													VPNClientAddressPool: &armnetwork.AddressSpace{
		// 														AddressPrefixes: []*string{
		// 															to.Ptr("101.4.0.0/16")},
		// 														},
		// 													},
		// 											}},
		// 											ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 											VirtualHub: &armnetwork.SubResource{
		// 												ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub2"),
		// 											},
		// 											VPNClientConnectionHealth: &armnetwork.VPNClientConnectionHealth{
		// 												AllocatedIPAddresses: []*string{
		// 													to.Ptr("1.1.1.1"),
		// 													to.Ptr("2.2.2.2")},
		// 													TotalEgressBytesTransferred: to.Ptr[int64](3000),
		// 													TotalIngressBytesTransferred: to.Ptr[int64](2000),
		// 													VPNClientConnectionsCount: to.Ptr[int32](2),
		// 												},
		// 												VPNGatewayScaleUnit: to.Ptr[int32](1),
		// 												VPNServerConfiguration: &armnetwork.SubResource{
		// 													ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/vpnServerConfigurations/vpnServerConfiguration1"),
		// 												},
		// 											},
		// 									}},
		// 								}
	}
}
