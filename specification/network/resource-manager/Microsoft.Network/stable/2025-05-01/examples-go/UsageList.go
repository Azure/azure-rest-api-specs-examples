package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v9"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/72410da64f6e945db1e1f1af220e077ba5bdb857/specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/UsageList.json
func ExampleUsagesClient_NewListPager_listUsages() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewUsagesClient().NewListPager("westus", nil)
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
		// page.UsagesListResult = armnetwork.UsagesListResult{
		// 	Value: []*armnetwork.Usage{
		// 		{
		// 			Name: &armnetwork.UsageName{
		// 				LocalizedValue: to.Ptr("Virtual Networks"),
		// 				Value: to.Ptr("VirtualNetworks"),
		// 			},
		// 			CurrentValue: to.Ptr[int64](8),
		// 			ID: to.Ptr("/subscriptions/subid/providers/Microsoft.Network/locations/westus/usages/VirtualNetworks"),
		// 			Limit: to.Ptr[int64](50),
		// 			Unit: to.Ptr(armnetwork.UsageUnitCount),
		// 		},
		// 		{
		// 			Name: &armnetwork.UsageName{
		// 				LocalizedValue: to.Ptr("Static Public IP Addresses"),
		// 				Value: to.Ptr("StaticPublicIPAddresses"),
		// 			},
		// 			CurrentValue: to.Ptr[int64](3),
		// 			ID: to.Ptr("/subscriptions/subid/providers/Microsoft.Network/locations/westus/usages/StaticPublicIPAddresses"),
		// 			Limit: to.Ptr[int64](20),
		// 			Unit: to.Ptr(armnetwork.UsageUnitCount),
		// 		},
		// 		{
		// 			Name: &armnetwork.UsageName{
		// 				LocalizedValue: to.Ptr("Network Security Groups"),
		// 				Value: to.Ptr("NetworkSecurityGroups"),
		// 			},
		// 			CurrentValue: to.Ptr[int64](1),
		// 			ID: to.Ptr("/subscriptions/subid/providers/Microsoft.Network/locations/westus/usages/NetworkSecurityGroups"),
		// 			Limit: to.Ptr[int64](100),
		// 			Unit: to.Ptr(armnetwork.UsageUnitCount),
		// 		},
		// 		{
		// 			Name: &armnetwork.UsageName{
		// 				LocalizedValue: to.Ptr("Public IP Addresses"),
		// 				Value: to.Ptr("PublicIPAddresses"),
		// 			},
		// 			CurrentValue: to.Ptr[int64](8),
		// 			ID: to.Ptr("/subscriptions/subid/providers/Microsoft.Network/locations/westus/usages/PublicIPAddresses"),
		// 			Limit: to.Ptr[int64](60),
		// 			Unit: to.Ptr(armnetwork.UsageUnitCount),
		// 		},
		// 		{
		// 			Name: &armnetwork.UsageName{
		// 				LocalizedValue: to.Ptr("Network Interfaces"),
		// 				Value: to.Ptr("NetworkInterfaces"),
		// 			},
		// 			CurrentValue: to.Ptr[int64](2),
		// 			ID: to.Ptr("/subscriptions/subid/providers/Microsoft.Network/locations/westus/usages/NetworkInterfaces"),
		// 			Limit: to.Ptr[int64](350),
		// 			Unit: to.Ptr(armnetwork.UsageUnitCount),
		// 		},
		// 		{
		// 			Name: &armnetwork.UsageName{
		// 				LocalizedValue: to.Ptr("Load Balancers"),
		// 				Value: to.Ptr("LoadBalancers"),
		// 			},
		// 			CurrentValue: to.Ptr[int64](2),
		// 			ID: to.Ptr("/subscriptions/subid/providers/Microsoft.Network/locations/westus/usages/LoadBalancers"),
		// 			Limit: to.Ptr[int64](100),
		// 			Unit: to.Ptr(armnetwork.UsageUnitCount),
		// 		},
		// 		{
		// 			Name: &armnetwork.UsageName{
		// 				LocalizedValue: to.Ptr("Application Gateways"),
		// 				Value: to.Ptr("ApplicationGateways"),
		// 			},
		// 			CurrentValue: to.Ptr[int64](1),
		// 			ID: to.Ptr("/subscriptions/subid/providers/Microsoft.Network/locations/westus/usages/ApplicationGateways"),
		// 			Limit: to.Ptr[int64](50),
		// 			Unit: to.Ptr(armnetwork.UsageUnitCount),
		// 		},
		// 		{
		// 			Name: &armnetwork.UsageName{
		// 				LocalizedValue: to.Ptr("Route Tables"),
		// 				Value: to.Ptr("RouteTables"),
		// 			},
		// 			CurrentValue: to.Ptr[int64](0),
		// 			ID: to.Ptr("/subscriptions/subid/providers/Microsoft.Network/locations/westus/usages/RouteTables"),
		// 			Limit: to.Ptr[int64](100),
		// 			Unit: to.Ptr(armnetwork.UsageUnitCount),
		// 		},
		// 		{
		// 			Name: &armnetwork.UsageName{
		// 				LocalizedValue: to.Ptr("Route Filters"),
		// 				Value: to.Ptr("RouteFilters"),
		// 			},
		// 			CurrentValue: to.Ptr[int64](0),
		// 			ID: to.Ptr("/subscriptions/subid/providers/Microsoft.Network/locations/westus/usages/RouteFilters"),
		// 			Limit: to.Ptr[int64](1000),
		// 			Unit: to.Ptr(armnetwork.UsageUnitCount),
		// 		},
		// 		{
		// 			Name: &armnetwork.UsageName{
		// 				LocalizedValue: to.Ptr("Network Watchers"),
		// 				Value: to.Ptr("NetworkWatchers"),
		// 			},
		// 			CurrentValue: to.Ptr[int64](0),
		// 			ID: to.Ptr("/subscriptions/subid/providers/Microsoft.Network/locations/westus/usages/NetworkWatchers"),
		// 			Limit: to.Ptr[int64](1),
		// 			Unit: to.Ptr(armnetwork.UsageUnitCount),
		// 		},
		// 		{
		// 			Name: &armnetwork.UsageName{
		// 				LocalizedValue: to.Ptr("Packet Captures"),
		// 				Value: to.Ptr("PacketCaptures"),
		// 			},
		// 			CurrentValue: to.Ptr[int64](0),
		// 			ID: to.Ptr("/subscriptions/subid/providers/Microsoft.Network/locations/westus/usages/PacketCaptures"),
		// 			Limit: to.Ptr[int64](10),
		// 			Unit: to.Ptr(armnetwork.UsageUnitCount),
		// 		},
		// 		{
		// 			Name: &armnetwork.UsageName{
		// 				LocalizedValue: to.Ptr("DNS servers per Virtual Network"),
		// 				Value: to.Ptr("DnsServersPerVirtualNetwork"),
		// 			},
		// 			CurrentValue: to.Ptr[int64](0),
		// 			ID: to.Ptr("/subscriptions/subid/providers/Microsoft.Network/locations/westus/usages/DnsServersPerVirtualNetwork"),
		// 			Limit: to.Ptr[int64](9),
		// 			Unit: to.Ptr(armnetwork.UsageUnitCount),
		// 		},
		// 		{
		// 			Name: &armnetwork.UsageName{
		// 				LocalizedValue: to.Ptr("Subnets per Virtual Network"),
		// 				Value: to.Ptr("SubnetsPerVirtualNetwork"),
		// 			},
		// 			CurrentValue: to.Ptr[int64](0),
		// 			ID: to.Ptr("/subscriptions/subid/providers/Microsoft.Network/locations/westus/usages/SubnetsPerVirtualNetwork"),
		// 			Limit: to.Ptr[int64](1000),
		// 			Unit: to.Ptr(armnetwork.UsageUnitCount),
		// 		},
		// 		{
		// 			Name: &armnetwork.UsageName{
		// 				LocalizedValue: to.Ptr("IP Configurations per Virtual Network"),
		// 				Value: to.Ptr("IPConfigurationsPerVirtualNetwork"),
		// 			},
		// 			CurrentValue: to.Ptr[int64](0),
		// 			ID: to.Ptr("/subscriptions/subid/providers/Microsoft.Network/locations/westus/usages/IPConfigurationsPerVirtualNetwork"),
		// 			Limit: to.Ptr[int64](4096),
		// 			Unit: to.Ptr(armnetwork.UsageUnitCount),
		// 		},
		// 		{
		// 			Name: &armnetwork.UsageName{
		// 				LocalizedValue: to.Ptr("Peerings per Virtual Network"),
		// 				Value: to.Ptr("PeeringsPerVirtualNetwork"),
		// 			},
		// 			CurrentValue: to.Ptr[int64](0),
		// 			ID: to.Ptr("/subscriptions/subid/providers/Microsoft.Network/locations/westus/usages/PeeringsPerVirtualNetwork"),
		// 			Limit: to.Ptr[int64](10),
		// 			Unit: to.Ptr(armnetwork.UsageUnitCount),
		// 		},
		// 		{
		// 			Name: &armnetwork.UsageName{
		// 				LocalizedValue: to.Ptr("Security rules per Network Security Group"),
		// 				Value: to.Ptr("SecurityRulesPerNetworkSecurityGroup"),
		// 			},
		// 			CurrentValue: to.Ptr[int64](0),
		// 			ID: to.Ptr("/subscriptions/subid/providers/Microsoft.Network/locations/westus/usages/SecurityRulesPerNetworkSecurityGroup"),
		// 			Limit: to.Ptr[int64](200),
		// 			Unit: to.Ptr(armnetwork.UsageUnitCount),
		// 		},
		// 		{
		// 			Name: &armnetwork.UsageName{
		// 				LocalizedValue: to.Ptr("Security rules addresses or ports per Network Security Group"),
		// 				Value: to.Ptr("SecurityRuleAddressesOrPortsPerNetworkSecurityGroup"),
		// 			},
		// 			CurrentValue: to.Ptr[int64](0),
		// 			ID: to.Ptr("/subscriptions/subid/providers/Microsoft.Network/locations/westus/usages/SecurityRuleAddressesOrPortsPerNetworkSecurityGroup"),
		// 			Limit: to.Ptr[int64](2000),
		// 			Unit: to.Ptr(armnetwork.UsageUnitCount),
		// 		},
		// 		{
		// 			Name: &armnetwork.UsageName{
		// 				LocalizedValue: to.Ptr("Inbound Rules per Load Balancer"),
		// 				Value: to.Ptr("InboundRulesPerLoadBalancer"),
		// 			},
		// 			CurrentValue: to.Ptr[int64](0),
		// 			ID: to.Ptr("/subscriptions/subid/providers/Microsoft.Network/locations/westus/usages/InboundRulesPerLoadBalancer"),
		// 			Limit: to.Ptr[int64](150),
		// 			Unit: to.Ptr(armnetwork.UsageUnitCount),
		// 		},
		// 		{
		// 			Name: &armnetwork.UsageName{
		// 				LocalizedValue: to.Ptr("Frontend IP Configurations per Load Balancer"),
		// 				Value: to.Ptr("FrontendIPConfigurationPerLoadBalancer"),
		// 			},
		// 			CurrentValue: to.Ptr[int64](0),
		// 			ID: to.Ptr("/subscriptions/subid/providers/Microsoft.Network/locations/westus/usages/FrontendIPConfigurationPerLoadBalancer"),
		// 			Limit: to.Ptr[int64](10),
		// 			Unit: to.Ptr(armnetwork.UsageUnitCount),
		// 		},
		// 		{
		// 			Name: &armnetwork.UsageName{
		// 				LocalizedValue: to.Ptr("Outbound Rules per Load Balancer"),
		// 				Value: to.Ptr("outboundRulesPerLoadBalancer"),
		// 			},
		// 			CurrentValue: to.Ptr[int64](0),
		// 			ID: to.Ptr("/subscriptions/subid/providers/Microsoft.Network/locations/westus/usages/outboundRulesPerLoadBalancer"),
		// 			Limit: to.Ptr[int64](5),
		// 			Unit: to.Ptr(armnetwork.UsageUnitCount),
		// 		},
		// 		{
		// 			Name: &armnetwork.UsageName{
		// 				LocalizedValue: to.Ptr("Routes per Route Table"),
		// 				Value: to.Ptr("RoutesPerRouteTable"),
		// 			},
		// 			CurrentValue: to.Ptr[int64](0),
		// 			ID: to.Ptr("/subscriptions/subid/providers/Microsoft.Network/locations/westus/usages/RoutesPerRouteTable"),
		// 			Limit: to.Ptr[int64](100),
		// 			Unit: to.Ptr(armnetwork.UsageUnitCount),
		// 		},
		// 		{
		// 			Name: &armnetwork.UsageName{
		// 				LocalizedValue: to.Ptr("Secondary IP Configurations per Network Interface"),
		// 				Value: to.Ptr("SecondaryIPConfigurationsPerNetworkInterface"),
		// 			},
		// 			CurrentValue: to.Ptr[int64](0),
		// 			ID: to.Ptr("/subscriptions/subid/providers/Microsoft.Network/locations/westus/usages/SecondaryIPConfigurationsPerNetworkInterface"),
		// 			Limit: to.Ptr[int64](256),
		// 			Unit: to.Ptr(armnetwork.UsageUnitCount),
		// 		},
		// 		{
		// 			Name: &armnetwork.UsageName{
		// 				LocalizedValue: to.Ptr("Inbound rules per Network Interface"),
		// 				Value: to.Ptr("InboundRulesPerNetworkInterface"),
		// 			},
		// 			CurrentValue: to.Ptr[int64](0),
		// 			ID: to.Ptr("/subscriptions/subid/providers/Microsoft.Network/locations/westus/usages/InboundRulesPerNetworkInterface"),
		// 			Limit: to.Ptr[int64](500),
		// 			Unit: to.Ptr(armnetwork.UsageUnitCount),
		// 		},
		// 		{
		// 			Name: &armnetwork.UsageName{
		// 				LocalizedValue: to.Ptr("Route filter rules per Route Filter"),
		// 				Value: to.Ptr("RouteFilterRulesPerRouteFilter"),
		// 			},
		// 			CurrentValue: to.Ptr[int64](0),
		// 			ID: to.Ptr("/subscriptions/subid/providers/Microsoft.Network/locations/westus/usages/RouteFilterRulesPerRouteFilter"),
		// 			Limit: to.Ptr[int64](1),
		// 			Unit: to.Ptr(armnetwork.UsageUnitCount),
		// 		},
		// 		{
		// 			Name: &armnetwork.UsageName{
		// 				LocalizedValue: to.Ptr("Route filters per Express route BGP Peering"),
		// 				Value: to.Ptr("RouteFiltersPerExpressRouteBgpPeering"),
		// 			},
		// 			CurrentValue: to.Ptr[int64](0),
		// 			ID: to.Ptr("/subscriptions/subid/providers/Microsoft.Network/locations/westus/usages/RouteFiltersPerExpressRouteBgpPeering"),
		// 			Limit: to.Ptr[int64](1),
		// 			Unit: to.Ptr(armnetwork.UsageUnitCount),
		// 	}},
		// }
	}
}
