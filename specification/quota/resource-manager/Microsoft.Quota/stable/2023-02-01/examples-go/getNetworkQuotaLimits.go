package armquota_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/quota/armquota"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/db9788dde7a0c2c0d82e4fdf5f7b4de3843937e3/specification/quota/resource-manager/Microsoft.Quota/stable/2023-02-01/examples/getNetworkQuotaLimits.json
func ExampleClient_NewListPager_quotasListQuotaLimitsForNetwork() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armquota.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewClient().NewListPager("subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Network/locations/eastus", nil)
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
		// page.Limits = armquota.Limits{
		// 	Value: []*armquota.CurrentQuotaLimitBase{
		// 		{
		// 			Name: to.Ptr("VirtualNetworks"),
		// 			Type: to.Ptr("Microsoft.Quota/Quotas"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Network/locations/eastus/providers/Microsoft.Quota/Quotas/"),
		// 			Properties: &armquota.Properties{
		// 				Name: &armquota.ResourceName{
		// 					LocalizedValue: to.Ptr("Virtual Networks"),
		// 					Value: to.Ptr("VirtualNetworks"),
		// 				},
		// 				IsQuotaApplicable: to.Ptr(false),
		// 				Limit: &armquota.LimitObject{
		// 					LimitObjectType: to.Ptr(armquota.LimitTypeLimitValue),
		// 					Value: to.Ptr[int32](10),
		// 				},
		// 				Properties: map[string]any{
		// 				},
		// 				Unit: to.Ptr("Count"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("StaticPublicIPAddresses"),
		// 			Type: to.Ptr("Microsoft.Quota/Quotas"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Network/locations/eastus/providers/Microsoft.Quota/Quotas/"),
		// 			Properties: &armquota.Properties{
		// 				Name: &armquota.ResourceName{
		// 					LocalizedValue: to.Ptr("Static Public IP Addresses"),
		// 					Value: to.Ptr("StaticPublicIPAddresses"),
		// 				},
		// 				IsQuotaApplicable: to.Ptr(false),
		// 				Limit: &armquota.LimitObject{
		// 					LimitObjectType: to.Ptr(armquota.LimitTypeLimitValue),
		// 					Value: to.Ptr[int32](10),
		// 				},
		// 				Properties: map[string]any{
		// 				},
		// 				Unit: to.Ptr("Count"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("NetworkSecurityGroups"),
		// 			Type: to.Ptr("Microsoft.Quota/Quotas"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Network/locations/eastus/providers/Microsoft.Quota/Quotas/"),
		// 			Properties: &armquota.Properties{
		// 				Name: &armquota.ResourceName{
		// 					LocalizedValue: to.Ptr("Network Security Groups"),
		// 					Value: to.Ptr("NetworkSecurityGroups"),
		// 				},
		// 				IsQuotaApplicable: to.Ptr(false),
		// 				Limit: &armquota.LimitObject{
		// 					LimitObjectType: to.Ptr(armquota.LimitTypeLimitValue),
		// 					Value: to.Ptr[int32](10),
		// 				},
		// 				Properties: map[string]any{
		// 				},
		// 				Unit: to.Ptr("Count"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("PublicIPAddresses"),
		// 			Type: to.Ptr("Microsoft.Quota/Quotas"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Network/locations/eastus/providers/Microsoft.Quota/Quotas/"),
		// 			Properties: &armquota.Properties{
		// 				Name: &armquota.ResourceName{
		// 					LocalizedValue: to.Ptr("Public IP Addresses - Basic"),
		// 					Value: to.Ptr("PublicIPAddresses"),
		// 				},
		// 				IsQuotaApplicable: to.Ptr(true),
		// 				Limit: &armquota.LimitObject{
		// 					LimitObjectType: to.Ptr(armquota.LimitTypeLimitValue),
		// 					Value: to.Ptr[int32](10),
		// 				},
		// 				Properties: map[string]any{
		// 				},
		// 				Unit: to.Ptr("Count"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("PublicIpPrefixes"),
		// 			Type: to.Ptr("Microsoft.Quota/Quotas"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Network/locations/eastus/providers/Microsoft.Quota/Quotas/"),
		// 			Properties: &armquota.Properties{
		// 				Name: &armquota.ResourceName{
		// 					LocalizedValue: to.Ptr("Public Ip Prefixes"),
		// 					Value: to.Ptr("PublicIpPrefixes"),
		// 				},
		// 				IsQuotaApplicable: to.Ptr(false),
		// 				Limit: &armquota.LimitObject{
		// 					LimitObjectType: to.Ptr(armquota.LimitTypeLimitValue),
		// 					Value: to.Ptr[int32](10),
		// 				},
		// 				Properties: map[string]any{
		// 				},
		// 				Unit: to.Ptr("Count"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("NatGateways"),
		// 			Type: to.Ptr("Microsoft.Quota/Quotas"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Network/locations/eastus/providers/Microsoft.Quota/Quotas/"),
		// 			Properties: &armquota.Properties{
		// 				Name: &armquota.ResourceName{
		// 					LocalizedValue: to.Ptr("Nat Gateways"),
		// 					Value: to.Ptr("NatGateways"),
		// 				},
		// 				IsQuotaApplicable: to.Ptr(false),
		// 				Limit: &armquota.LimitObject{
		// 					LimitObjectType: to.Ptr(armquota.LimitTypeLimitValue),
		// 					Value: to.Ptr[int32](10),
		// 				},
		// 				Properties: map[string]any{
		// 				},
		// 				Unit: to.Ptr("Count"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("NetworkInterfaces"),
		// 			Type: to.Ptr("Microsoft.Quota/Quotas"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Network/locations/eastus/providers/Microsoft.Quota/Quotas/"),
		// 			Properties: &armquota.Properties{
		// 				Name: &armquota.ResourceName{
		// 					LocalizedValue: to.Ptr("Network Interfaces"),
		// 					Value: to.Ptr("NetworkInterfaces"),
		// 				},
		// 				IsQuotaApplicable: to.Ptr(false),
		// 				Limit: &armquota.LimitObject{
		// 					LimitObjectType: to.Ptr(armquota.LimitTypeLimitValue),
		// 					Value: to.Ptr[int32](10),
		// 				},
		// 				Properties: map[string]any{
		// 				},
		// 				Unit: to.Ptr("Count"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("PrivateEndpoints"),
		// 			Type: to.Ptr("Microsoft.Quota/Quotas"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Network/locations/eastus/providers/Microsoft.Quota/Quotas/"),
		// 			Properties: &armquota.Properties{
		// 				Name: &armquota.ResourceName{
		// 					LocalizedValue: to.Ptr("Private Endpoints"),
		// 					Value: to.Ptr("PrivateEndpoints"),
		// 				},
		// 				IsQuotaApplicable: to.Ptr(false),
		// 				Limit: &armquota.LimitObject{
		// 					LimitObjectType: to.Ptr(armquota.LimitTypeLimitValue),
		// 					Value: to.Ptr[int32](10),
		// 				},
		// 				Properties: map[string]any{
		// 				},
		// 				Unit: to.Ptr("Count"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("PrivateEndpointRedirectMaps"),
		// 			Type: to.Ptr("Microsoft.Quota/Quotas"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Network/locations/eastus/providers/Microsoft.Quota/Quotas/"),
		// 			Properties: &armquota.Properties{
		// 				Name: &armquota.ResourceName{
		// 					LocalizedValue: to.Ptr("Private Endpoint Redirect Maps"),
		// 					Value: to.Ptr("PrivateEndpointRedirectMaps"),
		// 				},
		// 				IsQuotaApplicable: to.Ptr(false),
		// 				Limit: &armquota.LimitObject{
		// 					LimitObjectType: to.Ptr(armquota.LimitTypeLimitValue),
		// 					Value: to.Ptr[int32](10),
		// 				},
		// 				Properties: map[string]any{
		// 				},
		// 				Unit: to.Ptr("Count"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("LoadBalancers"),
		// 			Type: to.Ptr("Microsoft.Quota/Quotas"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Network/locations/eastus/providers/Microsoft.Quota/Quotas/"),
		// 			Properties: &armquota.Properties{
		// 				Name: &armquota.ResourceName{
		// 					LocalizedValue: to.Ptr("Load Balancers"),
		// 					Value: to.Ptr("LoadBalancers"),
		// 				},
		// 				IsQuotaApplicable: to.Ptr(false),
		// 				Limit: &armquota.LimitObject{
		// 					LimitObjectType: to.Ptr(armquota.LimitTypeLimitValue),
		// 					Value: to.Ptr[int32](10),
		// 				},
		// 				Properties: map[string]any{
		// 				},
		// 				Unit: to.Ptr("Count"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("PrivateLinkServices"),
		// 			Type: to.Ptr("Microsoft.Quota/Quotas"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Network/locations/eastus/providers/Microsoft.Quota/Quotas/"),
		// 			Properties: &armquota.Properties{
		// 				Name: &armquota.ResourceName{
		// 					LocalizedValue: to.Ptr("Private Link Services"),
		// 					Value: to.Ptr("PrivateLinkServices"),
		// 				},
		// 				IsQuotaApplicable: to.Ptr(false),
		// 				Limit: &armquota.LimitObject{
		// 					LimitObjectType: to.Ptr(armquota.LimitTypeLimitValue),
		// 					Value: to.Ptr[int32](10),
		// 				},
		// 				Properties: map[string]any{
		// 				},
		// 				Unit: to.Ptr("Count"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("ApplicationGateways"),
		// 			Type: to.Ptr("Microsoft.Quota/Quotas"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Network/locations/eastus/providers/Microsoft.Quota/Quotas/"),
		// 			Properties: &armquota.Properties{
		// 				Name: &armquota.ResourceName{
		// 					LocalizedValue: to.Ptr("Application Gateways"),
		// 					Value: to.Ptr("ApplicationGateways"),
		// 				},
		// 				IsQuotaApplicable: to.Ptr(false),
		// 				Limit: &armquota.LimitObject{
		// 					LimitObjectType: to.Ptr(armquota.LimitTypeLimitValue),
		// 					Value: to.Ptr[int32](10),
		// 				},
		// 				Properties: map[string]any{
		// 				},
		// 				Unit: to.Ptr("Count"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("RouteTables"),
		// 			Type: to.Ptr("Microsoft.Quota/Quotas"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Network/locations/eastus/providers/Microsoft.Quota/Quotas/"),
		// 			Properties: &armquota.Properties{
		// 				Name: &armquota.ResourceName{
		// 					LocalizedValue: to.Ptr("Route Tables"),
		// 					Value: to.Ptr("RouteTables"),
		// 				},
		// 				IsQuotaApplicable: to.Ptr(false),
		// 				Limit: &armquota.LimitObject{
		// 					LimitObjectType: to.Ptr(armquota.LimitTypeLimitValue),
		// 					Value: to.Ptr[int32](10),
		// 				},
		// 				Properties: map[string]any{
		// 				},
		// 				Unit: to.Ptr("Count"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("RouteFilters"),
		// 			Type: to.Ptr("Microsoft.Quota/Quotas"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Network/locations/eastus/providers/Microsoft.Quota/Quotas/"),
		// 			Properties: &armquota.Properties{
		// 				Name: &armquota.ResourceName{
		// 					LocalizedValue: to.Ptr("Route Filters"),
		// 					Value: to.Ptr("RouteFilters"),
		// 				},
		// 				IsQuotaApplicable: to.Ptr(false),
		// 				Limit: &armquota.LimitObject{
		// 					LimitObjectType: to.Ptr(armquota.LimitTypeLimitValue),
		// 					Value: to.Ptr[int32](10),
		// 				},
		// 				Properties: map[string]any{
		// 				},
		// 				Unit: to.Ptr("Count"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("NetworkWatchers"),
		// 			Type: to.Ptr("Microsoft.Quota/Quotas"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Network/locations/eastus/providers/Microsoft.Quota/Quotas/"),
		// 			Properties: &armquota.Properties{
		// 				Name: &armquota.ResourceName{
		// 					LocalizedValue: to.Ptr("Network Watchers"),
		// 					Value: to.Ptr("NetworkWatchers"),
		// 				},
		// 				IsQuotaApplicable: to.Ptr(false),
		// 				Limit: &armquota.LimitObject{
		// 					LimitObjectType: to.Ptr(armquota.LimitTypeLimitValue),
		// 					Value: to.Ptr[int32](10),
		// 				},
		// 				Properties: map[string]any{
		// 				},
		// 				Unit: to.Ptr("Count"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("PacketCaptures"),
		// 			Type: to.Ptr("Microsoft.Quota/Quotas"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Network/locations/eastus/providers/Microsoft.Quota/Quotas/"),
		// 			Properties: &armquota.Properties{
		// 				Name: &armquota.ResourceName{
		// 					LocalizedValue: to.Ptr("Packet Captures"),
		// 					Value: to.Ptr("PacketCaptures"),
		// 				},
		// 				IsQuotaApplicable: to.Ptr(false),
		// 				Limit: &armquota.LimitObject{
		// 					LimitObjectType: to.Ptr(armquota.LimitTypeLimitValue),
		// 					Value: to.Ptr[int32](10),
		// 				},
		// 				Properties: map[string]any{
		// 				},
		// 				Unit: to.Ptr("Count"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("ApplicationSecurityGroups"),
		// 			Type: to.Ptr("Microsoft.Quota/Quotas"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Network/locations/eastus/providers/Microsoft.Quota/Quotas/"),
		// 			Properties: &armquota.Properties{
		// 				Name: &armquota.ResourceName{
		// 					LocalizedValue: to.Ptr("Application Security Groups."),
		// 					Value: to.Ptr("ApplicationSecurityGroups"),
		// 				},
		// 				IsQuotaApplicable: to.Ptr(false),
		// 				Limit: &armquota.LimitObject{
		// 					LimitObjectType: to.Ptr(armquota.LimitTypeLimitValue),
		// 					Value: to.Ptr[int32](10),
		// 				},
		// 				Properties: map[string]any{
		// 				},
		// 				Unit: to.Ptr("Count"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("DdosProtectionPlans"),
		// 			Type: to.Ptr("Microsoft.Quota/Quotas"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Network/locations/eastus/providers/Microsoft.Quota/Quotas/"),
		// 			Properties: &armquota.Properties{
		// 				Name: &armquota.ResourceName{
		// 					LocalizedValue: to.Ptr("DDoS Protection Plans."),
		// 					Value: to.Ptr("DdosProtectionPlans"),
		// 				},
		// 				IsQuotaApplicable: to.Ptr(false),
		// 				Limit: &armquota.LimitObject{
		// 					LimitObjectType: to.Ptr(armquota.LimitTypeLimitValue),
		// 					Value: to.Ptr[int32](10),
		// 				},
		// 				Properties: map[string]any{
		// 				},
		// 				Unit: to.Ptr("Count"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("DdosCustomPolicies"),
		// 			Type: to.Ptr("Microsoft.Quota/Quotas"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Network/locations/eastus/providers/Microsoft.Quota/Quotas/"),
		// 			Properties: &armquota.Properties{
		// 				Name: &armquota.ResourceName{
		// 					LocalizedValue: to.Ptr("DDoS customized policies"),
		// 					Value: to.Ptr("DdosCustomPolicies"),
		// 				},
		// 				IsQuotaApplicable: to.Ptr(false),
		// 				Limit: &armquota.LimitObject{
		// 					LimitObjectType: to.Ptr(armquota.LimitTypeLimitValue),
		// 					Value: to.Ptr[int32](10),
		// 				},
		// 				Properties: map[string]any{
		// 				},
		// 				Unit: to.Ptr("Count"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("ServiceEndpointPolicies"),
		// 			Type: to.Ptr("Microsoft.Quota/Quotas"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Network/locations/eastus/providers/Microsoft.Quota/Quotas/"),
		// 			Properties: &armquota.Properties{
		// 				Name: &armquota.ResourceName{
		// 					LocalizedValue: to.Ptr("Service Endpoint Policies"),
		// 					Value: to.Ptr("ServiceEndpointPolicies"),
		// 				},
		// 				IsQuotaApplicable: to.Ptr(false),
		// 				Limit: &armquota.LimitObject{
		// 					LimitObjectType: to.Ptr(armquota.LimitTypeLimitValue),
		// 					Value: to.Ptr[int32](10),
		// 				},
		// 				Properties: map[string]any{
		// 				},
		// 				Unit: to.Ptr("Count"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("NetworkIntentPolicies"),
		// 			Type: to.Ptr("Microsoft.Quota/Quotas"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Network/locations/eastus/providers/Microsoft.Quota/Quotas/"),
		// 			Properties: &armquota.Properties{
		// 				Name: &armquota.ResourceName{
		// 					LocalizedValue: to.Ptr("Network Intent Policies"),
		// 					Value: to.Ptr("NetworkIntentPolicies"),
		// 				},
		// 				IsQuotaApplicable: to.Ptr(false),
		// 				Limit: &armquota.LimitObject{
		// 					LimitObjectType: to.Ptr(armquota.LimitTypeLimitValue),
		// 					Value: to.Ptr[int32](10),
		// 				},
		// 				Properties: map[string]any{
		// 				},
		// 				Unit: to.Ptr("Count"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("StandardSkuLoadBalancers"),
		// 			Type: to.Ptr("Microsoft.Quota/Quotas"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Network/locations/eastus/providers/Microsoft.Quota/Quotas/"),
		// 			Properties: &armquota.Properties{
		// 				Name: &armquota.ResourceName{
		// 					LocalizedValue: to.Ptr("Standard Sku Load Balancers"),
		// 					Value: to.Ptr("StandardSkuLoadBalancers"),
		// 				},
		// 				IsQuotaApplicable: to.Ptr(false),
		// 				Limit: &armquota.LimitObject{
		// 					LimitObjectType: to.Ptr(armquota.LimitTypeLimitValue),
		// 					Value: to.Ptr[int32](10),
		// 				},
		// 				Properties: map[string]any{
		// 				},
		// 				Unit: to.Ptr("Count"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("StandardSkuPublicIpAddresses"),
		// 			Type: to.Ptr("Microsoft.Quota/Quotas"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Network/locations/eastus/providers/Microsoft.Quota/Quotas/"),
		// 			Properties: &armquota.Properties{
		// 				Name: &armquota.ResourceName{
		// 					LocalizedValue: to.Ptr("Public IP Addresses - Standard"),
		// 					Value: to.Ptr("StandardSkuPublicIpAddresses"),
		// 				},
		// 				IsQuotaApplicable: to.Ptr(true),
		// 				Limit: &armquota.LimitObject{
		// 					LimitObjectType: to.Ptr(armquota.LimitTypeLimitValue),
		// 					Value: to.Ptr[int32](10),
		// 				},
		// 				Properties: map[string]any{
		// 				},
		// 				Unit: to.Ptr("Count"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("DnsServersPerVirtualNetwork"),
		// 			Type: to.Ptr("Microsoft.Quota/Quotas"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Network/locations/eastus/providers/Microsoft.Quota/Quotas/"),
		// 			Properties: &armquota.Properties{
		// 				Name: &armquota.ResourceName{
		// 					LocalizedValue: to.Ptr("DNS servers per Virtual Network"),
		// 					Value: to.Ptr("DnsServersPerVirtualNetwork"),
		// 				},
		// 				IsQuotaApplicable: to.Ptr(false),
		// 				Limit: &armquota.LimitObject{
		// 					LimitObjectType: to.Ptr(armquota.LimitTypeLimitValue),
		// 					Value: to.Ptr[int32](10),
		// 				},
		// 				Properties: map[string]any{
		// 				},
		// 				Unit: to.Ptr("Count"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("CustomDnsServersPerP2SVpnGateway"),
		// 			Type: to.Ptr("Microsoft.Quota/Quotas"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Network/locations/eastus/providers/Microsoft.Quota/Quotas/"),
		// 			Properties: &armquota.Properties{
		// 				Name: &armquota.ResourceName{
		// 					LocalizedValue: to.Ptr("Custom DNS servers per P2SVpnGateway"),
		// 					Value: to.Ptr("CustomDnsServersPerP2SVpnGateway"),
		// 				},
		// 				IsQuotaApplicable: to.Ptr(false),
		// 				Limit: &armquota.LimitObject{
		// 					LimitObjectType: to.Ptr(armquota.LimitTypeLimitValue),
		// 					Value: to.Ptr[int32](10),
		// 				},
		// 				Properties: map[string]any{
		// 				},
		// 				Unit: to.Ptr("Count"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("SubnetsPerVirtualNetwork"),
		// 			Type: to.Ptr("Microsoft.Quota/Quotas"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Network/locations/eastus/providers/Microsoft.Quota/Quotas/"),
		// 			Properties: &armquota.Properties{
		// 				Name: &armquota.ResourceName{
		// 					LocalizedValue: to.Ptr("Subnets per Virtual Network"),
		// 					Value: to.Ptr("SubnetsPerVirtualNetwork"),
		// 				},
		// 				IsQuotaApplicable: to.Ptr(false),
		// 				Limit: &armquota.LimitObject{
		// 					LimitObjectType: to.Ptr(armquota.LimitTypeLimitValue),
		// 					Value: to.Ptr[int32](10),
		// 				},
		// 				Properties: map[string]any{
		// 				},
		// 				Unit: to.Ptr("Count"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("IPConfigurationsPerVirtualNetwork"),
		// 			Type: to.Ptr("Microsoft.Quota/Quotas"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Network/locations/eastus/providers/Microsoft.Quota/Quotas/"),
		// 			Properties: &armquota.Properties{
		// 				Name: &armquota.ResourceName{
		// 					LocalizedValue: to.Ptr("IP Configurations per Virtual Network"),
		// 					Value: to.Ptr("IPConfigurationsPerVirtualNetwork"),
		// 				},
		// 				IsQuotaApplicable: to.Ptr(false),
		// 				Limit: &armquota.LimitObject{
		// 					LimitObjectType: to.Ptr(armquota.LimitTypeLimitValue),
		// 					Value: to.Ptr[int32](10),
		// 				},
		// 				Properties: map[string]any{
		// 				},
		// 				Unit: to.Ptr("Count"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("PeeringsPerVirtualNetwork"),
		// 			Type: to.Ptr("Microsoft.Quota/Quotas"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Network/locations/eastus/providers/Microsoft.Quota/Quotas/"),
		// 			Properties: &armquota.Properties{
		// 				Name: &armquota.ResourceName{
		// 					LocalizedValue: to.Ptr("Peerings per Virtual Network"),
		// 					Value: to.Ptr("PeeringsPerVirtualNetwork"),
		// 				},
		// 				IsQuotaApplicable: to.Ptr(false),
		// 				Limit: &armquota.LimitObject{
		// 					LimitObjectType: to.Ptr(armquota.LimitTypeLimitValue),
		// 					Value: to.Ptr[int32](10),
		// 				},
		// 				Properties: map[string]any{
		// 				},
		// 				Unit: to.Ptr("Count"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("SecurityRulesPerNetworkSecurityGroup"),
		// 			Type: to.Ptr("Microsoft.Quota/Quotas"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Network/locations/eastus/providers/Microsoft.Quota/Quotas/"),
		// 			Properties: &armquota.Properties{
		// 				Name: &armquota.ResourceName{
		// 					LocalizedValue: to.Ptr("Security rules per Network Security Group"),
		// 					Value: to.Ptr("SecurityRulesPerNetworkSecurityGroup"),
		// 				},
		// 				IsQuotaApplicable: to.Ptr(false),
		// 				Limit: &armquota.LimitObject{
		// 					LimitObjectType: to.Ptr(armquota.LimitTypeLimitValue),
		// 					Value: to.Ptr[int32](10),
		// 				},
		// 				Properties: map[string]any{
		// 				},
		// 				Unit: to.Ptr("Count"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("SecurityRulesPerNetworkIntentPolicy"),
		// 			Type: to.Ptr("Microsoft.Quota/Quotas"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Network/locations/eastus/providers/Microsoft.Quota/Quotas/"),
		// 			Properties: &armquota.Properties{
		// 				Name: &armquota.ResourceName{
		// 					LocalizedValue: to.Ptr("Security rules per Network Intent Policy"),
		// 					Value: to.Ptr("SecurityRulesPerNetworkIntentPolicy"),
		// 				},
		// 				IsQuotaApplicable: to.Ptr(false),
		// 				Limit: &armquota.LimitObject{
		// 					LimitObjectType: to.Ptr(armquota.LimitTypeLimitValue),
		// 					Value: to.Ptr[int32](10),
		// 				},
		// 				Properties: map[string]any{
		// 				},
		// 				Unit: to.Ptr("Count"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("RoutesPerNetworkIntentPolicy"),
		// 			Type: to.Ptr("Microsoft.Quota/Quotas"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Network/locations/eastus/providers/Microsoft.Quota/Quotas/"),
		// 			Properties: &armquota.Properties{
		// 				Name: &armquota.ResourceName{
		// 					LocalizedValue: to.Ptr("Routes per Network Intent Policy"),
		// 					Value: to.Ptr("RoutesPerNetworkIntentPolicy"),
		// 				},
		// 				IsQuotaApplicable: to.Ptr(false),
		// 				Limit: &armquota.LimitObject{
		// 					LimitObjectType: to.Ptr(armquota.LimitTypeLimitValue),
		// 					Value: to.Ptr[int32](10),
		// 				},
		// 				Properties: map[string]any{
		// 				},
		// 				Unit: to.Ptr("Count"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("SecurityRuleAddressesOrPortsPerNetworkSecurityGroup"),
		// 			Type: to.Ptr("Microsoft.Quota/Quotas"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Network/locations/eastus/providers/Microsoft.Quota/Quotas/"),
		// 			Properties: &armquota.Properties{
		// 				Name: &armquota.ResourceName{
		// 					LocalizedValue: to.Ptr("Security rules addresses or ports per Network Security Group"),
		// 					Value: to.Ptr("SecurityRuleAddressesOrPortsPerNetworkSecurityGroup"),
		// 				},
		// 				IsQuotaApplicable: to.Ptr(false),
		// 				Limit: &armquota.LimitObject{
		// 					LimitObjectType: to.Ptr(armquota.LimitTypeLimitValue),
		// 					Value: to.Ptr[int32](10),
		// 				},
		// 				Properties: map[string]any{
		// 				},
		// 				Unit: to.Ptr("Count"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("InboundRulesPerLoadBalancer"),
		// 			Type: to.Ptr("Microsoft.Quota/Quotas"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Network/locations/eastus/providers/Microsoft.Quota/Quotas/"),
		// 			Properties: &armquota.Properties{
		// 				Name: &armquota.ResourceName{
		// 					LocalizedValue: to.Ptr("Inbound Rules per Load Balancer"),
		// 					Value: to.Ptr("InboundRulesPerLoadBalancer"),
		// 				},
		// 				IsQuotaApplicable: to.Ptr(false),
		// 				Limit: &armquota.LimitObject{
		// 					LimitObjectType: to.Ptr(armquota.LimitTypeLimitValue),
		// 					Value: to.Ptr[int32](10),
		// 				},
		// 				Properties: map[string]any{
		// 				},
		// 				Unit: to.Ptr("Count"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("FrontendIPConfigurationPerLoadBalancer"),
		// 			Type: to.Ptr("Microsoft.Quota/Quotas"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Network/locations/eastus/providers/Microsoft.Quota/Quotas/"),
		// 			Properties: &armquota.Properties{
		// 				Name: &armquota.ResourceName{
		// 					LocalizedValue: to.Ptr("Frontend IP Configurations per Load Balancer"),
		// 					Value: to.Ptr("FrontendIPConfigurationPerLoadBalancer"),
		// 				},
		// 				IsQuotaApplicable: to.Ptr(false),
		// 				Limit: &armquota.LimitObject{
		// 					LimitObjectType: to.Ptr(armquota.LimitTypeLimitValue),
		// 					Value: to.Ptr[int32](10),
		// 				},
		// 				Properties: map[string]any{
		// 				},
		// 				Unit: to.Ptr("Count"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("OutboundRulesPerLoadBalancer"),
		// 			Type: to.Ptr("Microsoft.Quota/Quotas"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Network/locations/eastus/providers/Microsoft.Quota/Quotas/"),
		// 			Properties: &armquota.Properties{
		// 				Name: &armquota.ResourceName{
		// 					LocalizedValue: to.Ptr("Outbound Rules per Load Balancer"),
		// 					Value: to.Ptr("OutboundRulesPerLoadBalancer"),
		// 				},
		// 				IsQuotaApplicable: to.Ptr(false),
		// 				Limit: &armquota.LimitObject{
		// 					LimitObjectType: to.Ptr(armquota.LimitTypeLimitValue),
		// 					Value: to.Ptr[int32](10),
		// 				},
		// 				Properties: map[string]any{
		// 				},
		// 				Unit: to.Ptr("Count"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("RoutesPerRouteTable"),
		// 			Type: to.Ptr("Microsoft.Quota/Quotas"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Network/locations/eastus/providers/Microsoft.Quota/Quotas/"),
		// 			Properties: &armquota.Properties{
		// 				Name: &armquota.ResourceName{
		// 					LocalizedValue: to.Ptr("Routes per Route Table"),
		// 					Value: to.Ptr("RoutesPerRouteTable"),
		// 				},
		// 				IsQuotaApplicable: to.Ptr(false),
		// 				Limit: &armquota.LimitObject{
		// 					LimitObjectType: to.Ptr(armquota.LimitTypeLimitValue),
		// 					Value: to.Ptr[int32](10),
		// 				},
		// 				Properties: map[string]any{
		// 				},
		// 				Unit: to.Ptr("Count"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("RoutesWithServiceTagPerRouteTable"),
		// 			Type: to.Ptr("Microsoft.Quota/Quotas"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Network/locations/eastus/providers/Microsoft.Quota/Quotas/"),
		// 			Properties: &armquota.Properties{
		// 				Name: &armquota.ResourceName{
		// 					LocalizedValue: to.Ptr("Routes with service tag per Route Table"),
		// 					Value: to.Ptr("RoutesWithServiceTagPerRouteTable"),
		// 				},
		// 				IsQuotaApplicable: to.Ptr(false),
		// 				Limit: &armquota.LimitObject{
		// 					LimitObjectType: to.Ptr(armquota.LimitTypeLimitValue),
		// 					Value: to.Ptr[int32](10),
		// 				},
		// 				Properties: map[string]any{
		// 				},
		// 				Unit: to.Ptr("Count"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("MinPublicIpInterNetworkPrefixLength"),
		// 			Type: to.Ptr("Microsoft.Quota/Quotas"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Network/locations/eastus/providers/Microsoft.Quota/Quotas/"),
		// 			Properties: &armquota.Properties{
		// 				Name: &armquota.ResourceName{
		// 					LocalizedValue: to.Ptr("Min Public Ip InterNetwork Prefix Length"),
		// 					Value: to.Ptr("MinPublicIpInterNetworkPrefixLength"),
		// 				},
		// 				IsQuotaApplicable: to.Ptr(true),
		// 				Limit: &armquota.LimitObject{
		// 					LimitObjectType: to.Ptr(armquota.LimitTypeLimitValue),
		// 					Value: to.Ptr[int32](10),
		// 				},
		// 				Properties: map[string]any{
		// 				},
		// 				ResourceType: to.Ptr("MinPublicIpInterNetworkPrefixLength"),
		// 				Unit: to.Ptr("Count"),
		// 			},
		// 	}},
		// }
	}
}
