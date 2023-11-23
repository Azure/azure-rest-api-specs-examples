package armquota_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/quota/armquota"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/db9788dde7a0c2c0d82e4fdf5f7b4de3843937e3/specification/quota/resource-manager/Microsoft.Quota/stable/2023-02-01/examples/getNetworkUsages.json
func ExampleUsagesClient_NewListPager_quotasListUsagesForNetwork() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armquota.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewUsagesClient().NewListPager("subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Network/locations/eastus", nil)
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
		// page.UsagesLimits = armquota.UsagesLimits{
		// 	Value: []*armquota.CurrentUsagesBase{
		// 		{
		// 			Name: to.Ptr("VirtualNetworks"),
		// 			Type: to.Ptr("Microsoft.Quota/usages"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Network/locations/eastus/providers/Microsoft.Quota/usages/"),
		// 			Properties: &armquota.UsagesProperties{
		// 				Name: &armquota.ResourceName{
		// 					LocalizedValue: to.Ptr("Virtual Networks"),
		// 					Value: to.Ptr("VirtualNetworks"),
		// 				},
		// 				IsQuotaApplicable: to.Ptr(false),
		// 				Properties: map[string]any{
		// 				},
		// 				Unit: to.Ptr("Count"),
		// 				Usages: &armquota.UsagesObject{
		// 					Value: to.Ptr[int32](10),
		// 				},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("StaticPublicIPAddresses"),
		// 			Type: to.Ptr("Microsoft.Quota/usages"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Network/locations/eastus/providers/Microsoft.Quota/usages/"),
		// 			Properties: &armquota.UsagesProperties{
		// 				Name: &armquota.ResourceName{
		// 					LocalizedValue: to.Ptr("Static Public IP Addresses"),
		// 					Value: to.Ptr("StaticPublicIPAddresses"),
		// 				},
		// 				IsQuotaApplicable: to.Ptr(false),
		// 				Properties: map[string]any{
		// 				},
		// 				Unit: to.Ptr("Count"),
		// 				Usages: &armquota.UsagesObject{
		// 					Value: to.Ptr[int32](10),
		// 				},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("NetworkSecurityGroups"),
		// 			Type: to.Ptr("Microsoft.Quota/usages"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Network/locations/eastus/providers/Microsoft.Quota/usages/"),
		// 			Properties: &armquota.UsagesProperties{
		// 				Name: &armquota.ResourceName{
		// 					LocalizedValue: to.Ptr("Network Security Groups"),
		// 					Value: to.Ptr("NetworkSecurityGroups"),
		// 				},
		// 				IsQuotaApplicable: to.Ptr(false),
		// 				Properties: map[string]any{
		// 				},
		// 				Unit: to.Ptr("Count"),
		// 				Usages: &armquota.UsagesObject{
		// 					Value: to.Ptr[int32](10),
		// 				},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("PublicIPAddresses"),
		// 			Type: to.Ptr("Microsoft.Quota/usages"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Network/locations/eastus/providers/Microsoft.Quota/usages/"),
		// 			Properties: &armquota.UsagesProperties{
		// 				Name: &armquota.ResourceName{
		// 					LocalizedValue: to.Ptr("Public IP Addresses - Basic"),
		// 					Value: to.Ptr("PublicIPAddresses"),
		// 				},
		// 				IsQuotaApplicable: to.Ptr(true),
		// 				Properties: map[string]any{
		// 				},
		// 				ResourceType: to.Ptr("PublicIpAddresses"),
		// 				Unit: to.Ptr("Count"),
		// 				Usages: &armquota.UsagesObject{
		// 					Value: to.Ptr[int32](10),
		// 				},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("PublicIpPrefixes"),
		// 			Type: to.Ptr("Microsoft.Quota/usages"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Network/locations/eastus/providers/Microsoft.Quota/usages/"),
		// 			Properties: &armquota.UsagesProperties{
		// 				Name: &armquota.ResourceName{
		// 					LocalizedValue: to.Ptr("Public Ip Prefixes"),
		// 					Value: to.Ptr("PublicIpPrefixes"),
		// 				},
		// 				IsQuotaApplicable: to.Ptr(false),
		// 				Properties: map[string]any{
		// 				},
		// 				Unit: to.Ptr("Count"),
		// 				Usages: &armquota.UsagesObject{
		// 					Value: to.Ptr[int32](10),
		// 				},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("NatGateways"),
		// 			Type: to.Ptr("Microsoft.Quota/usages"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Network/locations/eastus/providers/Microsoft.Quota/usages/"),
		// 			Properties: &armquota.UsagesProperties{
		// 				Name: &armquota.ResourceName{
		// 					LocalizedValue: to.Ptr("Nat Gateways"),
		// 					Value: to.Ptr("NatGateways"),
		// 				},
		// 				IsQuotaApplicable: to.Ptr(false),
		// 				Properties: map[string]any{
		// 				},
		// 				Unit: to.Ptr("Count"),
		// 				Usages: &armquota.UsagesObject{
		// 					Value: to.Ptr[int32](10),
		// 				},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("NetworkInterfaces"),
		// 			Type: to.Ptr("Microsoft.Quota/usages"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Network/locations/eastus/providers/Microsoft.Quota/usages/"),
		// 			Properties: &armquota.UsagesProperties{
		// 				Name: &armquota.ResourceName{
		// 					LocalizedValue: to.Ptr("Network Interfaces"),
		// 					Value: to.Ptr("NetworkInterfaces"),
		// 				},
		// 				IsQuotaApplicable: to.Ptr(false),
		// 				Properties: map[string]any{
		// 				},
		// 				Unit: to.Ptr("Count"),
		// 				Usages: &armquota.UsagesObject{
		// 					Value: to.Ptr[int32](10),
		// 				},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("PrivateEndpoints"),
		// 			Type: to.Ptr("Microsoft.Quota/usages"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Network/locations/eastus/providers/Microsoft.Quota/usages/"),
		// 			Properties: &armquota.UsagesProperties{
		// 				Name: &armquota.ResourceName{
		// 					LocalizedValue: to.Ptr("Private Endpoints"),
		// 					Value: to.Ptr("PrivateEndpoints"),
		// 				},
		// 				IsQuotaApplicable: to.Ptr(false),
		// 				Properties: map[string]any{
		// 				},
		// 				Unit: to.Ptr("Count"),
		// 				Usages: &armquota.UsagesObject{
		// 					Value: to.Ptr[int32](10),
		// 				},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("PrivateEndpointRedirectMaps"),
		// 			Type: to.Ptr("Microsoft.Quota/usages"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Network/locations/eastus/providers/Microsoft.Quota/usages/"),
		// 			Properties: &armquota.UsagesProperties{
		// 				Name: &armquota.ResourceName{
		// 					LocalizedValue: to.Ptr("Private Endpoint Redirect Maps"),
		// 					Value: to.Ptr("PrivateEndpointRedirectMaps"),
		// 				},
		// 				IsQuotaApplicable: to.Ptr(false),
		// 				Properties: map[string]any{
		// 				},
		// 				Unit: to.Ptr("Count"),
		// 				Usages: &armquota.UsagesObject{
		// 					Value: to.Ptr[int32](10),
		// 				},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("LoadBalancers"),
		// 			Type: to.Ptr("Microsoft.Quota/usages"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Network/locations/eastus/providers/Microsoft.Quota/usages/"),
		// 			Properties: &armquota.UsagesProperties{
		// 				Name: &armquota.ResourceName{
		// 					LocalizedValue: to.Ptr("Load Balancers"),
		// 					Value: to.Ptr("LoadBalancers"),
		// 				},
		// 				IsQuotaApplicable: to.Ptr(false),
		// 				Properties: map[string]any{
		// 				},
		// 				Unit: to.Ptr("Count"),
		// 				Usages: &armquota.UsagesObject{
		// 					Value: to.Ptr[int32](10),
		// 				},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("PrivateLinkServices"),
		// 			Type: to.Ptr("Microsoft.Quota/usages"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Network/locations/eastus/providers/Microsoft.Quota/usages/"),
		// 			Properties: &armquota.UsagesProperties{
		// 				Name: &armquota.ResourceName{
		// 					LocalizedValue: to.Ptr("Private Link Services"),
		// 					Value: to.Ptr("PrivateLinkServices"),
		// 				},
		// 				IsQuotaApplicable: to.Ptr(false),
		// 				Properties: map[string]any{
		// 				},
		// 				Unit: to.Ptr("Count"),
		// 				Usages: &armquota.UsagesObject{
		// 					Value: to.Ptr[int32](10),
		// 				},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("ApplicationGateways"),
		// 			Type: to.Ptr("Microsoft.Quota/usages"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Network/locations/eastus/providers/Microsoft.Quota/usages/"),
		// 			Properties: &armquota.UsagesProperties{
		// 				Name: &armquota.ResourceName{
		// 					LocalizedValue: to.Ptr("Application Gateways"),
		// 					Value: to.Ptr("ApplicationGateways"),
		// 				},
		// 				IsQuotaApplicable: to.Ptr(false),
		// 				Properties: map[string]any{
		// 				},
		// 				Unit: to.Ptr("Count"),
		// 				Usages: &armquota.UsagesObject{
		// 					Value: to.Ptr[int32](10),
		// 				},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("RouteTables"),
		// 			Type: to.Ptr("Microsoft.Quota/usages"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Network/locations/eastus/providers/Microsoft.Quota/usages/"),
		// 			Properties: &armquota.UsagesProperties{
		// 				Name: &armquota.ResourceName{
		// 					LocalizedValue: to.Ptr("Route Tables"),
		// 					Value: to.Ptr("RouteTables"),
		// 				},
		// 				IsQuotaApplicable: to.Ptr(false),
		// 				Properties: map[string]any{
		// 				},
		// 				Unit: to.Ptr("Count"),
		// 				Usages: &armquota.UsagesObject{
		// 					Value: to.Ptr[int32](10),
		// 				},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("RouteFilters"),
		// 			Type: to.Ptr("Microsoft.Quota/usages"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Network/locations/eastus/providers/Microsoft.Quota/usages/"),
		// 			Properties: &armquota.UsagesProperties{
		// 				Name: &armquota.ResourceName{
		// 					LocalizedValue: to.Ptr("Route Filters"),
		// 					Value: to.Ptr("RouteFilters"),
		// 				},
		// 				IsQuotaApplicable: to.Ptr(false),
		// 				Properties: map[string]any{
		// 				},
		// 				Unit: to.Ptr("Count"),
		// 				Usages: &armquota.UsagesObject{
		// 					Value: to.Ptr[int32](10),
		// 				},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("NetworkWatchers"),
		// 			Type: to.Ptr("Microsoft.Quota/usages"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Network/locations/eastus/providers/Microsoft.Quota/usages/"),
		// 			Properties: &armquota.UsagesProperties{
		// 				Name: &armquota.ResourceName{
		// 					LocalizedValue: to.Ptr("Network Watchers"),
		// 					Value: to.Ptr("NetworkWatchers"),
		// 				},
		// 				IsQuotaApplicable: to.Ptr(false),
		// 				Properties: map[string]any{
		// 				},
		// 				Unit: to.Ptr("Count"),
		// 				Usages: &armquota.UsagesObject{
		// 					Value: to.Ptr[int32](10),
		// 				},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("PacketCaptures"),
		// 			Type: to.Ptr("Microsoft.Quota/usages"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Network/locations/eastus/providers/Microsoft.Quota/usages/"),
		// 			Properties: &armquota.UsagesProperties{
		// 				Name: &armquota.ResourceName{
		// 					LocalizedValue: to.Ptr("Packet Captures"),
		// 					Value: to.Ptr("PacketCaptures"),
		// 				},
		// 				IsQuotaApplicable: to.Ptr(false),
		// 				Properties: map[string]any{
		// 				},
		// 				Unit: to.Ptr("Count"),
		// 				Usages: &armquota.UsagesObject{
		// 					Value: to.Ptr[int32](10),
		// 				},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("ApplicationSecurityGroups"),
		// 			Type: to.Ptr("Microsoft.Quota/usages"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Network/locations/eastus/providers/Microsoft.Quota/usages/"),
		// 			Properties: &armquota.UsagesProperties{
		// 				Name: &armquota.ResourceName{
		// 					LocalizedValue: to.Ptr("Application Security Groups."),
		// 					Value: to.Ptr("ApplicationSecurityGroups"),
		// 				},
		// 				IsQuotaApplicable: to.Ptr(false),
		// 				Properties: map[string]any{
		// 				},
		// 				Unit: to.Ptr("Count"),
		// 				Usages: &armquota.UsagesObject{
		// 					Value: to.Ptr[int32](10),
		// 				},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("StandardSkuLoadBalancers"),
		// 			Type: to.Ptr("Microsoft.Quota/usages"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Network/locations/eastus/providers/Microsoft.Quota/usages/"),
		// 			Properties: &armquota.UsagesProperties{
		// 				Name: &armquota.ResourceName{
		// 					LocalizedValue: to.Ptr("Standard Sku Load Balancers"),
		// 					Value: to.Ptr("StandardSkuLoadBalancers"),
		// 				},
		// 				IsQuotaApplicable: to.Ptr(false),
		// 				Properties: map[string]any{
		// 				},
		// 				Unit: to.Ptr("Count"),
		// 				Usages: &armquota.UsagesObject{
		// 					Value: to.Ptr[int32](10),
		// 				},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("StandardSkuPublicIpAddresses"),
		// 			Type: to.Ptr("Microsoft.Quota/usages"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Network/locations/eastus/providers/Microsoft.Quota/usages/"),
		// 			Properties: &armquota.UsagesProperties{
		// 				Name: &armquota.ResourceName{
		// 					LocalizedValue: to.Ptr("Public IP Addresses - Standard"),
		// 					Value: to.Ptr("StandardSkuPublicIpAddresses"),
		// 				},
		// 				IsQuotaApplicable: to.Ptr(true),
		// 				ResourceType: to.Ptr("PublicIpAddresses"),
		// 				Unit: to.Ptr("Count"),
		// 				Usages: &armquota.UsagesObject{
		// 					Value: to.Ptr[int32](10),
		// 				},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("DnsServersPerVirtualNetwork"),
		// 			Type: to.Ptr("Microsoft.Quota/usages"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Network/locations/eastus/providers/Microsoft.Quota/usages/"),
		// 			Properties: &armquota.UsagesProperties{
		// 				Name: &armquota.ResourceName{
		// 					LocalizedValue: to.Ptr("DNS servers per Virtual Network"),
		// 					Value: to.Ptr("DnsServersPerVirtualNetwork"),
		// 				},
		// 				IsQuotaApplicable: to.Ptr(false),
		// 				Properties: map[string]any{
		// 				},
		// 				Unit: to.Ptr("Count"),
		// 				Usages: &armquota.UsagesObject{
		// 					Value: to.Ptr[int32](10),
		// 				},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("MinPublicIpInterNetworkPrefixLength"),
		// 			Type: to.Ptr("Microsoft.Quota/usages"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Network/locations/eastus/providers/Microsoft.Quota/usages/"),
		// 			Properties: &armquota.UsagesProperties{
		// 				Name: &armquota.ResourceName{
		// 					LocalizedValue: to.Ptr("Min Public Ip InterNetwork Prefix Length"),
		// 					Value: to.Ptr("MinPublicIpInterNetworkPrefixLength"),
		// 				},
		// 				IsQuotaApplicable: to.Ptr(true),
		// 				Properties: map[string]any{
		// 				},
		// 				Unit: to.Ptr("Count"),
		// 				Usages: &armquota.UsagesObject{
		// 					Value: to.Ptr[int32](10),
		// 				},
		// 			},
		// 	}},
		// }
	}
}
