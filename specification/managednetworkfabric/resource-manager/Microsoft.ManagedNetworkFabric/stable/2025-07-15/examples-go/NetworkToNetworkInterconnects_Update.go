package armmanagednetworkfabric_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managednetworkfabric/armmanagednetworkfabric/v2"
)

// Generated from example definition: 2025-07-15/NetworkToNetworkInterconnects_Update.json
func ExampleNetworkToNetworkInterconnectsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagednetworkfabric.NewClientFactory("0000ABCD-0A0B-0000-0000-000000ABCDEF", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewNetworkToNetworkInterconnectsClient().BeginUpdate(ctx, "example-rg", "example-nf", "example-nni", armmanagednetworkfabric.NetworkToNetworkInterconnectPatch{
		Properties: &armmanagednetworkfabric.NetworkToNetworkInterconnectPatchProperties{
			Layer2Configuration: &armmanagednetworkfabric.Layer2ConfigurationPatch{
				Mtu: to.Ptr[int32](1500),
				Interfaces: []*string{
					to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/networkDevices/example-networkDevice/networkInterfaces/example-networkInterface"),
				},
			},
			OptionBLayer3Configuration: &armmanagednetworkfabric.OptionBLayer3ConfigurationPatchProperties{
				PrimaryIPv4Prefix:   to.Ptr("20.0.0.12/29"),
				PrimaryIPv6Prefix:   to.Ptr("4FFE:FFFF:0:CD30::a8/127"),
				SecondaryIPv4Prefix: to.Ptr("20.0.0.14/29"),
				SecondaryIPv6Prefix: to.Ptr("6FFE:FFFF:0:CD30::ac/127"),
				PeerASN:             to.Ptr[int64](2345),
				VlanID:              to.Ptr[int32](1235),
				PeLoopbackIPAddress: []*string{
					to.Ptr("10.0.0.1"),
				},
				BmpConfiguration: &armmanagednetworkfabric.NniBmpPatchProperties{
					ConfigurationState: to.Ptr(armmanagednetworkfabric.BmpConfigurationStateEnabled),
				},
				PrefixLimits: []*armmanagednetworkfabric.OptionBLayer3PrefixLimitPatchProperties{
					{
						MaximumRoutes: to.Ptr[int32](1),
					},
				},
			},
			NpbStaticRouteConfiguration: &armmanagednetworkfabric.NpbStaticRouteConfigurationPatch{
				BfdConfiguration: &armmanagednetworkfabric.BfdPatchConfiguration{
					IntervalInMilliSeconds: to.Ptr[int32](300),
					Multiplier:             to.Ptr[int32](10),
				},
				IPv4Routes: []*armmanagednetworkfabric.StaticRoutePatchProperties{
					{
						Prefix: to.Ptr("10.0.0.1/24"),
						NextHop: []*string{
							to.Ptr("10.0.0.1"),
						},
					},
				},
				IPv6Routes: []*armmanagednetworkfabric.StaticRoutePatchProperties{
					{
						Prefix: to.Ptr("fe80::/64"),
						NextHop: []*string{
							to.Ptr("fe80::1"),
						},
					},
				},
			},
			StaticRouteConfiguration: &armmanagednetworkfabric.NniStaticRoutePatchConfiguration{
				BfdConfiguration: &armmanagednetworkfabric.BfdPatchConfiguration{
					IntervalInMilliSeconds: to.Ptr[int32](300),
					Multiplier:             to.Ptr[int32](10),
				},
				IPv4Routes: []*armmanagednetworkfabric.StaticRoutePatchProperties{
					{
						Prefix: to.Ptr("10.0.0.1"),
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
			ImportRoutePolicy: &armmanagednetworkfabric.ImportRoutePolicyInformationPatch{
				ImportIPv4RoutePolicyID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/routePolicies/example-routePolicy"),
				ImportIPv6RoutePolicyID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/routePolicies/example-routePolicy"),
			},
			ExportRoutePolicy: &armmanagednetworkfabric.ExportRoutePolicyInformationPatch{
				ExportIPv4RoutePolicyID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/routePolicies/example-routePolicy"),
				ExportIPv6RoutePolicyID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/routePolicies/example-routePolicy"),
			},
			IngressACLID:  to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/accessControlLists/example-acl"),
			EgressACLID:   to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/accessControlLists/example-acl"),
			MicroBfdState: to.Ptr(armmanagednetworkfabric.MicroBfdStateEnabled),
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
	// res = armmanagednetworkfabric.NetworkToNetworkInterconnectsClientUpdateResponse{
	// 	NetworkToNetworkInterconnect: &armmanagednetworkfabric.NetworkToNetworkInterconnect{
	// 		Properties: &armmanagednetworkfabric.NetworkToNetworkInterconnectProperties{
	// 			NniType: to.Ptr(armmanagednetworkfabric.NniTypeCE),
	// 			IsManagementType: to.Ptr(armmanagednetworkfabric.IsManagementTypeTrue),
	// 			UseOptionB: to.Ptr(armmanagednetworkfabric.BooleanEnumPropertyTrue),
	// 			Layer2Configuration: &armmanagednetworkfabric.Layer2Configuration{
	// 				Mtu: to.Ptr[int32](1500),
	// 				Interfaces: []*string{
	// 					to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/networkDevices/example-networkDevice/networkInterfaces/example-networkInterface"),
	// 				},
	// 			},
	// 			OptionBLayer3Configuration: &armmanagednetworkfabric.OptionBLayer3Configuration{
	// 				PrimaryIPv4Prefix: to.Ptr("10.0.0.12/30"),
	// 				PrimaryIPv6Prefix: to.Ptr("4FFE:FFFF:0:CD30::a8/127"),
	// 				SecondaryIPv4Prefix: to.Ptr("40.0.0.14/30"),
	// 				SecondaryIPv6Prefix: to.Ptr("6FFE:FFFF:0:CD30::ac/127"),
	// 				PeerASN: to.Ptr[int64](61234),
	// 				VlanID: to.Ptr[int32](1234),
	// 				FabricASN: to.Ptr[int64](17),
	// 				PeLoopbackIPAddress: []*string{
	// 					to.Ptr("10.0.0.1"),
	// 				},
	// 				BmpConfiguration: &armmanagednetworkfabric.NniBmpProperties{
	// 					ConfigurationState: to.Ptr(armmanagednetworkfabric.BmpConfigurationStateEnabled),
	// 				},
	// 			},
	// 			NpbStaticRouteConfiguration: &armmanagednetworkfabric.NpbStaticRouteConfiguration{
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
	// 						Prefix: to.Ptr("fe80::/64"),
	// 						NextHop: []*string{
	// 							to.Ptr("fe80::1"),
	// 						},
	// 					},
	// 				},
	// 			},
	// 			StaticRouteConfiguration: &armmanagednetworkfabric.NniStaticRouteConfiguration{
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
	// 			ImportRoutePolicy: &armmanagednetworkfabric.ImportRoutePolicyInformation{
	// 				ImportIPv4RoutePolicyID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/routePolicies/example-routePolicy"),
	// 				ImportIPv6RoutePolicyID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/routePolicies/example-routePolicy"),
	// 			},
	// 			ExportRoutePolicy: &armmanagednetworkfabric.ExportRoutePolicyInformation{
	// 				ExportIPv4RoutePolicyID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/routePolicies/example-routePolicy"),
	// 				ExportIPv6RoutePolicyID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/routePolicies/example-routePolicy"),
	// 			},
	// 			IngressACLID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/accessControlLists/example-acl"),
	// 			EgressACLID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/accessControlLists/example-acl"),
	// 			MicroBfdState: to.Ptr(armmanagednetworkfabric.MicroBfdStateEnabled),
	// 			ConditionalDefaultRouteConfiguration: &armmanagednetworkfabric.ConditionalDefaultRouteProperties{
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
	// 						Prefix: to.Ptr("fe80::/64"),
	// 						NextHop: []*string{
	// 							to.Ptr("fe80::1"),
	// 						},
	// 					},
	// 				},
	// 			},
	// 			ConfigurationState: to.Ptr(armmanagednetworkfabric.ConfigurationStateSucceeded),
	// 			ProvisioningState: to.Ptr(armmanagednetworkfabric.ProvisioningStateAccepted),
	// 			AdministrativeState: to.Ptr(armmanagednetworkfabric.AdministrativeStateEnabled),
	// 		},
	// 		ID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/networkFabrics/example-fabric/networkToNetworkInterconnects/example-nni"),
	// 		Name: to.Ptr("example-nni"),
	// 		Type: to.Ptr("microsoft.managedNetworkFabric/networkFabrics/networkToNetworkInterconnects"),
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
