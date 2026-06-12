package armmanagednetworkfabric_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managednetworkfabric/armmanagednetworkfabric/v2"
)

// Generated from example definition: 2025-07-15/ExternalNetworks_Update.json
func ExampleExternalNetworksClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagednetworkfabric.NewClientFactory("0000ABCD-0A0B-0000-0000-000000ABCDEF", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewExternalNetworksClient().BeginUpdate(ctx, "example-rg", "example-externalnetwork", "example-ext", armmanagednetworkfabric.ExternalNetworkPatch{
		Properties: &armmanagednetworkfabric.ExternalNetworkPatchProperties{
			Annotation:                     to.Ptr("annotation1"),
			NetworkToNetworkInterconnectID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/networkFabrics/example-fabric/networkToNetworkInterconnects/example-nni"),
			ImportRoutePolicy: &armmanagednetworkfabric.ImportRoutePolicyPatch{
				ImportIPv4RoutePolicyID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/routePolicies/routePolicyName"),
				ImportIPv6RoutePolicyID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/routePolicies/routePolicyName"),
			},
			ExportRoutePolicy: &armmanagednetworkfabric.ExportRoutePolicyPatch{
				ExportIPv4RoutePolicyID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/routePolicies/routePolicyName"),
				ExportIPv6RoutePolicyID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/routePolicies/routePolicyName"),
			},
			PeeringOption: to.Ptr(armmanagednetworkfabric.PeeringOptionOptionA),
			OptionBProperties: &armmanagednetworkfabric.L3OptionBPatchProperties{
				ImportRouteTargets: []*string{
					to.Ptr("65046:10039"),
				},
				ExportRouteTargets: []*string{
					to.Ptr("65046:10039"),
				},
				RouteTargets: &armmanagednetworkfabric.RouteTargetPatchInformation{
					ImportIPv4RouteTargets: []*string{
						to.Ptr("65046:10050"),
					},
					ImportIPv6RouteTargets: []*string{
						to.Ptr("65046:10050"),
					},
					ExportIPv4RouteTargets: []*string{
						to.Ptr("65046:10050"),
					},
					ExportIPv6RouteTargets: []*string{
						to.Ptr("65046:10050"),
					},
				},
			},
			OptionAProperties: &armmanagednetworkfabric.ExternalNetworkPatchPropertiesOptionAProperties{
				PrimaryIPv4Prefix:   to.Ptr("10.1.1.0/30"),
				PrimaryIPv6Prefix:   to.Ptr("3FFE:FFFF:0:CD30::a0/126"),
				SecondaryIPv4Prefix: to.Ptr("10.1.1.4/30"),
				SecondaryIPv6Prefix: to.Ptr("3FFE:FFFF:0:CD30::a4/126"),
				Mtu:                 to.Ptr[int32](1500),
				VlanID:              to.Ptr[int32](1001),
				PeerASN:             to.Ptr[int64](65047),
				BfdConfiguration: &armmanagednetworkfabric.BfdPatchConfiguration{
					IntervalInMilliSeconds: to.Ptr[int32](300),
					Multiplier:             to.Ptr[int32](10),
				},
				IngressACLID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/accessControlLists/example-acl"),
				EgressACLID:  to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/accessControlLists/example-acl"),
				BmpConfiguration: &armmanagednetworkfabric.ExternalNetworkBmpPatchProperties{
					ConfigurationState: to.Ptr(armmanagednetworkfabric.BmpConfigurationStateEnabled),
				},
				V4OverV6BgpSession: to.Ptr(armmanagednetworkfabric.V4OverV6BgpSessionStateEnabled),
				V6OverV4BgpSession: to.Ptr(armmanagednetworkfabric.V6OverV4BgpSessionStateEnabled),
				NativeIPv4PrefixLimit: &armmanagednetworkfabric.NativeIPv4PrefixLimitPatchProperties{
					PrefixLimits: []*armmanagednetworkfabric.PrefixLimitPatchProperties{
						{
							MaximumRoutes:  to.Ptr[int32](13),
							Threshold:      to.Ptr[int32](24),
							IdleTimeExpiry: to.Ptr[int32](8),
						},
					},
				},
				NativeIPv6PrefixLimit: &armmanagednetworkfabric.NativeIPv6PrefixLimitPatchProperties{
					PrefixLimits: []*armmanagednetworkfabric.PrefixLimitPatchProperties{
						{
							MaximumRoutes:  to.Ptr[int32](13),
							Threshold:      to.Ptr[int32](24),
							IdleTimeExpiry: to.Ptr[int32](8),
						},
					},
				},
			},
			StaticRouteConfiguration: &armmanagednetworkfabric.ExternalNetworkStaticRoutePatchConfiguration{
				BfdConfiguration: &armmanagednetworkfabric.BfdPatchConfiguration{
					IntervalInMilliSeconds: to.Ptr[int32](300),
					Multiplier:             to.Ptr[int32](10),
				},
				IPv4Routes: []*armmanagednetworkfabric.StaticRoutePatchProperties{
					{
						Prefix: to.Ptr("10.0.0.1/14"),
						NextHop: []*string{
							to.Ptr("10.0.0.1"),
						},
					},
				},
				IPv6Routes: []*armmanagednetworkfabric.StaticRoutePatchProperties{
					{
						Prefix: to.Ptr("2fff::/64"),
						NextHop: []*string{
							to.Ptr("3ffe::1"),
						},
					},
				},
			},
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
	// res = armmanagednetworkfabric.ExternalNetworksClientUpdateResponse{
	// 	ExternalNetwork: &armmanagednetworkfabric.ExternalNetwork{
	// 		Properties: &armmanagednetworkfabric.ExternalNetworkProperties{
	// 			Annotation: to.Ptr("annotation"),
	// 			NetworkToNetworkInterconnectID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/networkFabrics/example-fabric/networkToNetworkInterconnects/example-nni"),
	// 			ImportRoutePolicy: &armmanagednetworkfabric.ImportRoutePolicy{
	// 				ImportIPv4RoutePolicyID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/routePolicies/routePolicyName"),
	// 				ImportIPv6RoutePolicyID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/routePolicies/routePolicyName"),
	// 			},
	// 			ExportRoutePolicy: &armmanagednetworkfabric.ExportRoutePolicy{
	// 				ExportIPv4RoutePolicyID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/routePolicies/routePolicyName"),
	// 				ExportIPv6RoutePolicyID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/routePolicies/routePolicyName"),
	// 			},
	// 			PeeringOption: to.Ptr(armmanagednetworkfabric.PeeringOptionOptionA),
	// 			OptionBProperties: &armmanagednetworkfabric.L3OptionBProperties{
	// 				ImportRouteTargets: []*string{
	// 					to.Ptr("65046:10039"),
	// 				},
	// 				ExportRouteTargets: []*string{
	// 					to.Ptr("65046:10039"),
	// 				},
	// 				RouteTargets: &armmanagednetworkfabric.RouteTargetInformation{
	// 					ImportIPv4RouteTargets: []*string{
	// 						to.Ptr("65046:10050"),
	// 					},
	// 					ImportIPv6RouteTargets: []*string{
	// 						to.Ptr("65046:10050"),
	// 					},
	// 					ExportIPv4RouteTargets: []*string{
	// 						to.Ptr("65046:10050"),
	// 					},
	// 					ExportIPv6RouteTargets: []*string{
	// 						to.Ptr("65046:10050"),
	// 					},
	// 				},
	// 			},
	// 			OptionAProperties: &armmanagednetworkfabric.ExternalNetworkPropertiesOptionAProperties{
	// 				PrimaryIPv4Prefix: to.Ptr("10.1.1.0/30"),
	// 				PrimaryIPv6Prefix: to.Ptr("3FFE:FFFF:0:CD30::a0/126"),
	// 				SecondaryIPv4Prefix: to.Ptr("10.1.1.4/30"),
	// 				SecondaryIPv6Prefix: to.Ptr("3FFE:FFFF:0:CD30::a4/126"),
	// 				Mtu: to.Ptr[int32](1500),
	// 				VlanID: to.Ptr[int32](1001),
	// 				FabricASN: to.Ptr[int64](1234),
	// 				PeerASN: to.Ptr[int64](65047),
	// 				BfdConfiguration: &armmanagednetworkfabric.BfdConfiguration{
	// 					AdministrativeState: to.Ptr(armmanagednetworkfabric.BfdAdministrativeStateEnabled),
	// 					IntervalInMilliSeconds: to.Ptr[int32](300),
	// 					Multiplier: to.Ptr[int32](10),
	// 				},
	// 				IngressACLID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/accessControlLists/example-acl"),
	// 				EgressACLID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/accessControlLists/example-acl"),
	// 				BmpConfiguration: &armmanagednetworkfabric.ExternalNetworkBmpProperties{
	// 					ConfigurationState: to.Ptr(armmanagednetworkfabric.BmpConfigurationStateEnabled),
	// 				},
	// 			},
	// 			StaticRouteConfiguration: &armmanagednetworkfabric.ExternalNetworkStaticRouteConfiguration{
	// 				BfdConfiguration: &armmanagednetworkfabric.BfdConfiguration{
	// 					AdministrativeState: to.Ptr(armmanagednetworkfabric.BfdAdministrativeStateEnabled),
	// 					IntervalInMilliSeconds: to.Ptr[int32](300),
	// 					Multiplier: to.Ptr[int32](10),
	// 				},
	// 				IPv4Routes: []*armmanagednetworkfabric.StaticRouteProperties{
	// 					{
	// 						Prefix: to.Ptr("10.0.0.1/24"),
	// 						NextHop: []*string{
	// 							to.Ptr("10.0.0.1"),
	// 						},
	// 					},
	// 				},
	// 				IPv6Routes: []*armmanagednetworkfabric.StaticRouteProperties{
	// 					{
	// 						Prefix: to.Ptr("2fff::/64"),
	// 						NextHop: []*string{
	// 							to.Ptr("3ffe::1"),
	// 						},
	// 					},
	// 				},
	// 			},
	// 			ConfigurationState: to.Ptr(armmanagednetworkfabric.ConfigurationStateSucceeded),
	// 			ProvisioningState: to.Ptr(armmanagednetworkfabric.ProvisioningStateAccepted),
	// 			AdministrativeState: to.Ptr(armmanagednetworkfabric.AdministrativeStateEnabled),
	// 		},
	// 		ID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/l3IsolationDomains/example-l3domain/externalNetworks/example-externalnetwork"),
	// 		Name: to.Ptr("example-externalnetwork"),
	// 		Type: to.Ptr("microsoft.managednetworkfabric/l3IsolationDomains/externalnetworks"),
	// 		SystemData: &armmanagednetworkfabric.SystemData{
	// 			CreatedBy: to.Ptr("email@address.com"),
	// 			CreatedByType: to.Ptr(armmanagednetworkfabric.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-09T04:51:41.251Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("UserId"),
	// 			LastModifiedByType: to.Ptr(armmanagednetworkfabric.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-09T04:51:41.251Z"); return t}()),
	// 		},
	// 	},
	// }
}
