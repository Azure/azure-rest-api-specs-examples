package armmanagednetworkfabric_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managednetworkfabric/armmanagednetworkfabric/v2"
)

// Generated from example definition: 2025-07-15/InternalNetworks_ListByL3IsolationDomain.json
func ExampleInternalNetworksClient_NewListByL3IsolationDomainPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagednetworkfabric.NewClientFactory("0000ABCD-0A0B-0000-0000-000000ABCDEF", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewInternalNetworksClient().NewListByL3IsolationDomainPager("example-rg", "example-l3isd", nil)
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
		// page = armmanagednetworkfabric.InternalNetworksClientListByL3IsolationDomainResponse{
		// 	InternalNetworksList: armmanagednetworkfabric.InternalNetworksList{
		// 		Value: []*armmanagednetworkfabric.InternalNetwork{
		// 			{
		// 				Properties: &armmanagednetworkfabric.InternalNetworkProperties{
		// 					Annotation: to.Ptr("annotation"),
		// 					Mtu: to.Ptr[int32](1500),
		// 					NetworkFabricID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/networkFabrics/example-fabric"),
		// 					ConnectedIPv4Subnets: []*armmanagednetworkfabric.ConnectedSubnet{
		// 						{
		// 							Prefix: to.Ptr("10.0.0.0/24"),
		// 							Annotation: to.Ptr("annotation"),
		// 						},
		// 					},
		// 					ConnectedIPv6Subnets: []*armmanagednetworkfabric.ConnectedSubnet{
		// 						{
		// 							Prefix: to.Ptr("3FFE:FFFF:0:CD30::a0/29"),
		// 							Annotation: to.Ptr("annotation"),
		// 						},
		// 					},
		// 					ImportRoutePolicy: &armmanagednetworkfabric.ImportRoutePolicy{
		// 						ImportIPv4RoutePolicyID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/routePolicies/routePolicyName"),
		// 						ImportIPv6RoutePolicyID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/routePolicies/routePolicyName"),
		// 					},
		// 					ExportRoutePolicy: &armmanagednetworkfabric.ExportRoutePolicy{
		// 						ExportIPv4RoutePolicyID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/routePolicies/routePolicyName"),
		// 						ExportIPv6RoutePolicyID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/routePolicies/routePolicyName"),
		// 					},
		// 					IngressACLID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/accessControlLists/example-acl"),
		// 					EgressACLID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/accessControlLists/example-acl"),
		// 					IsMonitoringEnabled: to.Ptr(armmanagednetworkfabric.IsMonitoringEnabledTrue),
		// 					Extension: to.Ptr(armmanagednetworkfabric.ExtensionNoExtension),
		// 					VlanID: to.Ptr[int32](755),
		// 					BgpConfiguration: &armmanagednetworkfabric.BgpConfiguration{
		// 						Annotation: to.Ptr("annotation"),
		// 						BfdConfiguration: &armmanagednetworkfabric.BfdConfiguration{
		// 							AdministrativeState: to.Ptr(armmanagednetworkfabric.BfdAdministrativeStateEnabled),
		// 							IntervalInMilliSeconds: to.Ptr[int32](300),
		// 							Multiplier: to.Ptr[int32](10),
		// 						},
		// 						DefaultRouteOriginate: to.Ptr(armmanagednetworkfabric.BooleanEnumPropertyTrue),
		// 						AllowAS: to.Ptr[int32](10),
		// 						AllowASOverride: to.Ptr(armmanagednetworkfabric.AllowASOverrideEnable),
		// 						FabricASN: to.Ptr[int64](20),
		// 						PeerASN: to.Ptr[int64](61234),
		// 						IPv4ListenRangePrefixes: []*string{
		// 							to.Ptr("10.1.0.0/25"),
		// 						},
		// 						IPv6ListenRangePrefixes: []*string{
		// 							to.Ptr("2fff::/66"),
		// 						},
		// 						IPv4NeighborAddress: []*armmanagednetworkfabric.NeighborAddress{
		// 							{
		// 								Address: to.Ptr("10.1.0.0"),
		// 								BfdAdministrativeState: to.Ptr(armmanagednetworkfabric.BfdAdministrativeState("Enable")),
		// 								BgpAdministrativeState: to.Ptr(armmanagednetworkfabric.BgpAdministrativeState("Enable")),
		// 								ConfigurationState: to.Ptr(armmanagednetworkfabric.ConfigurationStateSucceeded),
		// 							},
		// 						},
		// 						IPv6NeighborAddress: []*armmanagednetworkfabric.NeighborAddress{
		// 							{
		// 								Address: to.Ptr("2fff::"),
		// 								BfdAdministrativeState: to.Ptr(armmanagednetworkfabric.BfdAdministrativeState("Enable")),
		// 								BgpAdministrativeState: to.Ptr(armmanagednetworkfabric.BgpAdministrativeState("Enable")),
		// 								ConfigurationState: to.Ptr(armmanagednetworkfabric.ConfigurationStateSucceeded),
		// 							},
		// 						},
		// 						BmpConfiguration: &armmanagednetworkfabric.InternalNetworkBmpProperties{
		// 							NeighborIPExclusions: []*string{
		// 								to.Ptr("10.0.0.1/24"),
		// 							},
		// 							BmpConfigurationState: to.Ptr(armmanagednetworkfabric.BmpConfigurationStateEnabled),
		// 							ExportPolicyConfiguration: &armmanagednetworkfabric.BmpExportPolicyProperties{
		// 								ExportPolicies: []*armmanagednetworkfabric.BmpExportPolicy{
		// 									to.Ptr(armmanagednetworkfabric.BmpExportPolicyPrePolicy),
		// 								},
		// 							},
		// 						},
		// 						V4OverV6BgpSession: to.Ptr(armmanagednetworkfabric.V4OverV6BgpSessionStateEnabled),
		// 						V6OverV4BgpSession: to.Ptr(armmanagednetworkfabric.V6OverV4BgpSessionStateEnabled),
		// 					},
		// 					StaticRouteConfiguration: &armmanagednetworkfabric.StaticRouteConfiguration{
		// 						BfdConfiguration: &armmanagednetworkfabric.BfdConfiguration{
		// 							AdministrativeState: to.Ptr(armmanagednetworkfabric.BfdAdministrativeStateEnabled),
		// 							IntervalInMilliSeconds: to.Ptr[int32](300),
		// 							Multiplier: to.Ptr[int32](10),
		// 						},
		// 						IPv4Routes: []*armmanagednetworkfabric.StaticRouteProperties{
		// 							{
		// 								Prefix: to.Ptr("jffgck"),
		// 								NextHop: []*string{
		// 									to.Ptr("10.0.0.1"),
		// 								},
		// 							},
		// 						},
		// 						IPv6Routes: []*armmanagednetworkfabric.StaticRouteProperties{
		// 							{
		// 								Prefix: to.Ptr("2fff::/64"),
		// 								NextHop: []*string{
		// 									to.Ptr("3ffe::1"),
		// 								},
		// 							},
		// 						},
		// 						Extension: to.Ptr(armmanagednetworkfabric.ExtensionNoExtension),
		// 					},
		// 					NativeIPv4PrefixLimit: &armmanagednetworkfabric.NativeIPv4PrefixLimitProperties{
		// 						PrefixLimits: []*armmanagednetworkfabric.PrefixLimitProperties{
		// 							{
		// 								MaximumRoutes: to.Ptr[int32](23),
		// 								Threshold: to.Ptr[int32](7),
		// 								IdleTimeExpiry: to.Ptr[int32](28),
		// 							},
		// 						},
		// 					},
		// 					NativeIPv6PrefixLimit: &armmanagednetworkfabric.NativeIPv6PrefixLimitProperties{
		// 						PrefixLimits: []*armmanagednetworkfabric.PrefixLimitProperties{
		// 							{
		// 								MaximumRoutes: to.Ptr[int32](23),
		// 								Threshold: to.Ptr[int32](7),
		// 								IdleTimeExpiry: to.Ptr[int32](28),
		// 							},
		// 						},
		// 					},
		// 					LastOperation: &armmanagednetworkfabric.LastOperationProperties{
		// 						Details: to.Ptr("Succeeded"),
		// 					},
		// 					ConfigurationState: to.Ptr(armmanagednetworkfabric.ConfigurationStateSucceeded),
		// 					ProvisioningState: to.Ptr(armmanagednetworkfabric.ProvisioningStateAccepted),
		// 					AdministrativeState: to.Ptr(armmanagednetworkfabric.AdministrativeStateEnabled),
		// 				},
		// 				ID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/l3IsolationDomains/example-l3domain/internalNetworks/example-internalnetwork"),
		// 				Name: to.Ptr("example-internalnetwork"),
		// 				Type: to.Ptr("microsoft.managednetworkfabric/l3IsolationDomains/internalnetworks"),
		// 				SystemData: &armmanagednetworkfabric.SystemData{
		// 					CreatedBy: to.Ptr("email@address.com"),
		// 					CreatedByType: to.Ptr(armmanagednetworkfabric.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-09T04:51:41.251Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("UserId"),
		// 					LastModifiedByType: to.Ptr(armmanagednetworkfabric.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-09T04:51:41.251Z"); return t}()),
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://microsoft.com/aukskqxh"),
		// 	},
		// }
	}
}
