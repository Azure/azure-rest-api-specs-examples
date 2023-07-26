package armmanagednetworkfabric_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managednetworkfabric/armmanagednetworkfabric"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/925ba149e17454ce91ecd3f9f4134effb2f97844/specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/InternalNetworks_ListByL3IsolationDomain_MaximumSet_Gen.json
func ExampleInternalNetworksClient_NewListByL3IsolationDomainPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagednetworkfabric.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewInternalNetworksClient().NewListByL3IsolationDomainPager("example-rg", "example-l3domain", nil)
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
		// page.InternalNetworksList = armmanagednetworkfabric.InternalNetworksList{
		// 	Value: []*armmanagednetworkfabric.InternalNetwork{
		// 		{
		// 			Name: to.Ptr("example-internalnetwork"),
		// 			Type: to.Ptr("microsoft.managednetworkfabric/l3IsolationDomains/internalnetworks"),
		// 			ID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/l3IsolationDomains/example-l3domain/internalNetworks/example-internalnetwork"),
		// 			SystemData: &armmanagednetworkfabric.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-09T04:51:41.251Z"); return t}()),
		// 				CreatedBy: to.Ptr("email@address.com"),
		// 				CreatedByType: to.Ptr(armmanagednetworkfabric.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-09T04:51:41.251Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("UserId"),
		// 				LastModifiedByType: to.Ptr(armmanagednetworkfabric.CreatedByTypeUser),
		// 			},
		// 			Properties: &armmanagednetworkfabric.InternalNetworkProperties{
		// 				Annotation: to.Ptr("annotation"),
		// 				Extension: to.Ptr(armmanagednetworkfabric.ExtensionNoExtension),
		// 				ConnectedIPv4Subnets: []*armmanagednetworkfabric.ConnectedSubnet{
		// 					{
		// 						Annotation: to.Ptr("annotation"),
		// 						Prefix: to.Ptr("10.0.0.0/24"),
		// 				}},
		// 				ConnectedIPv6Subnets: []*armmanagednetworkfabric.ConnectedSubnet{
		// 					{
		// 						Annotation: to.Ptr("annotation"),
		// 						Prefix: to.Ptr("3FFE:FFFF:0:CD30::a0/29"),
		// 				}},
		// 				EgressACLID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/accessControlLists/example-acl"),
		// 				ExportRoutePolicy: &armmanagednetworkfabric.ExportRoutePolicy{
		// 					ExportIPv4RoutePolicyID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/routePolicies/routePolicyName"),
		// 					ExportIPv6RoutePolicyID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/routePolicies/routePolicyName"),
		// 				},
		// 				ExportRoutePolicyID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/routePolicies/routePolicyName"),
		// 				ImportRoutePolicy: &armmanagednetworkfabric.ImportRoutePolicy{
		// 					ImportIPv4RoutePolicyID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/routePolicies/routePolicyName"),
		// 					ImportIPv6RoutePolicyID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/routePolicies/routePolicyName"),
		// 				},
		// 				ImportRoutePolicyID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/routePolicies/routePolicyName"),
		// 				IngressACLID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/accessControlLists/example-acl"),
		// 				IsMonitoringEnabled: to.Ptr(armmanagednetworkfabric.IsMonitoringEnabledTrue),
		// 				Mtu: to.Ptr[int32](1500),
		// 				AdministrativeState: to.Ptr(armmanagednetworkfabric.AdministrativeStateEnabled),
		// 				BgpConfiguration: &armmanagednetworkfabric.InternalNetworkPropertiesBgpConfiguration{
		// 					Annotation: to.Ptr("annotation"),
		// 					AllowAS: to.Ptr[int32](10),
		// 					AllowASOverride: to.Ptr(armmanagednetworkfabric.AllowASOverrideEnable),
		// 					BfdConfiguration: &armmanagednetworkfabric.BfdConfiguration{
		// 						AdministrativeState: to.Ptr(armmanagednetworkfabric.BfdAdministrativeStateEnabled),
		// 						IntervalInMilliSeconds: to.Ptr[int32](300),
		// 						Multiplier: to.Ptr[int32](5),
		// 					},
		// 					DefaultRouteOriginate: to.Ptr(armmanagednetworkfabric.BooleanEnumPropertyTrue),
		// 					FabricASN: to.Ptr[int64](20),
		// 					IPv4ListenRangePrefixes: []*string{
		// 						to.Ptr("10.1.0.0/25")},
		// 						IPv4NeighborAddress: []*armmanagednetworkfabric.NeighborAddress{
		// 							{
		// 								Address: to.Ptr("10.1.0.0"),
		// 								ConfigurationState: to.Ptr(armmanagednetworkfabric.ConfigurationStateSucceeded),
		// 						}},
		// 						IPv6ListenRangePrefixes: []*string{
		// 							to.Ptr("2fff::/66")},
		// 							IPv6NeighborAddress: []*armmanagednetworkfabric.NeighborAddress{
		// 								{
		// 									Address: to.Ptr("2fff::"),
		// 									ConfigurationState: to.Ptr(armmanagednetworkfabric.ConfigurationStateSucceeded),
		// 							}},
		// 							PeerASN: to.Ptr[int64](61234),
		// 						},
		// 						ConfigurationState: to.Ptr(armmanagednetworkfabric.ConfigurationStateSucceeded),
		// 						ProvisioningState: to.Ptr(armmanagednetworkfabric.ProvisioningStateSucceeded),
		// 						StaticRouteConfiguration: &armmanagednetworkfabric.InternalNetworkPropertiesStaticRouteConfiguration{
		// 							Extension: to.Ptr(armmanagednetworkfabric.ExtensionNoExtension),
		// 							BfdConfiguration: &armmanagednetworkfabric.BfdConfiguration{
		// 								AdministrativeState: to.Ptr(armmanagednetworkfabric.BfdAdministrativeStateEnabled),
		// 								IntervalInMilliSeconds: to.Ptr[int32](300),
		// 								Multiplier: to.Ptr[int32](15),
		// 							},
		// 							IPv4Routes: []*armmanagednetworkfabric.StaticRouteProperties{
		// 								{
		// 									NextHop: []*string{
		// 										to.Ptr("10.0.0.1")},
		// 										Prefix: to.Ptr("jffgck"),
		// 								}},
		// 								IPv6Routes: []*armmanagednetworkfabric.StaticRouteProperties{
		// 									{
		// 										NextHop: []*string{
		// 											to.Ptr("3ffe::1")},
		// 											Prefix: to.Ptr("2fff::/64"),
		// 									}},
		// 								},
		// 								VlanID: to.Ptr[int32](755),
		// 							},
		// 					}},
		// 				}
	}
}
