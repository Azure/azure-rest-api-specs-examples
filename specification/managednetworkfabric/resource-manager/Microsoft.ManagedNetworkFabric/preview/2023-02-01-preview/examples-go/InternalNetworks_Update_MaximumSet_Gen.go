package armmanagednetworkfabric_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managednetworkfabric/armmanagednetworkfabric"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d03c1964cb76ffd6884d10a1871bbe779a2f68ef/specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/preview/2023-02-01-preview/examples/InternalNetworks_Update_MaximumSet_Gen.json
func ExampleInternalNetworksClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagednetworkfabric.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewInternalNetworksClient().BeginUpdate(ctx, "resourceGroupName", "example-l3domain", "example-internalnetwork", armmanagednetworkfabric.InternalNetworkPatch{
		Properties: &armmanagednetworkfabric.InternalNetworkPatchProperties{
			BgpConfiguration: &armmanagednetworkfabric.BgpConfiguration{
				AllowAS:               to.Ptr[int32](1),
				AllowASOverride:       to.Ptr(armmanagednetworkfabric.AllowASOverrideEnable),
				BfdConfiguration:      &armmanagednetworkfabric.BfdConfiguration{},
				DefaultRouteOriginate: to.Ptr(armmanagednetworkfabric.BooleanEnumPropertyTrue),
				IPv4ListenRangePrefixes: []*string{
					to.Ptr("10.1.0.0/25")},
				IPv4NeighborAddress: []*armmanagednetworkfabric.NeighborAddress{
					{
						Address: to.Ptr("10.1.0.0"),
					}},
				IPv6ListenRangePrefixes: []*string{
					to.Ptr("2fff::/66")},
				IPv6NeighborAddress: []*armmanagednetworkfabric.NeighborAddress{
					{
						Address: to.Ptr("2fff::"),
					}},
				PeerASN: to.Ptr[int32](6),
			},
			ConnectedIPv4Subnets: []*armmanagednetworkfabric.ConnectedSubnet{
				{
					Prefix: to.Ptr("10.0.0.0/24"),
				}},
			ConnectedIPv6Subnets: []*armmanagednetworkfabric.ConnectedSubnet{
				{
					Prefix: to.Ptr("3FFE:FFFF:0:CD30::a0/29"),
				}},
			ExportRoutePolicyID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.ManagedNetworkFabric/routePolicies/routePolicyName2"),
			ImportRoutePolicyID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.ManagedNetworkFabric/routePolicies/routePolicyName1"),
			Mtu:                 to.Ptr[int32](1500),
			StaticRouteConfiguration: &armmanagednetworkfabric.StaticRouteConfiguration{
				BfdConfiguration: &armmanagednetworkfabric.BfdConfiguration{},
				IPv4Routes: []*armmanagednetworkfabric.StaticRouteProperties{
					{
						NextHop: []*string{
							to.Ptr("10.0.0.1")},
						Prefix: to.Ptr("10.1.0.0/24"),
					}},
				IPv6Routes: []*armmanagednetworkfabric.StaticRouteProperties{
					{
						NextHop: []*string{
							to.Ptr("2ffe::1")},
						Prefix: to.Ptr("2fff::/64"),
					}},
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
	// res.InternalNetwork = armmanagednetworkfabric.InternalNetwork{
	// 	Name: to.Ptr("example-externalnetwork"),
	// 	Type: to.Ptr("microsoft.managednetworkfabric/example-externalnetwork"),
	// 	ID: to.Ptr("/subscriptions/xxxxxx/resourceGroups/resourcegroupname/providers/Microsoft.ManagedNetworkFabric/example-l3domain/externalNetworks/example-externalnetwork"),
	// 	SystemData: &armmanagednetworkfabric.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-02-09T18:35:44.070Z"); return t}()),
	// 		CreatedBy: to.Ptr("email@address.com"),
	// 		CreatedByType: to.Ptr(armmanagednetworkfabric.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-02-09T18:35:44.070Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("UserId"),
	// 		LastModifiedByType: to.Ptr(armmanagednetworkfabric.CreatedByTypeUser),
	// 	},
	// 	Properties: &armmanagednetworkfabric.InternalNetworkProperties{
	// 		Annotation: to.Ptr("dojarxck"),
	// 		BgpConfiguration: &armmanagednetworkfabric.BgpConfiguration{
	// 			AllowAS: to.Ptr[int32](1),
	// 			AllowASOverride: to.Ptr(armmanagednetworkfabric.AllowASOverrideEnable),
	// 			BfdConfiguration: &armmanagednetworkfabric.BfdConfiguration{
	// 				AdministrativeState: to.Ptr(armmanagednetworkfabric.EnabledDisabledStateEnabled),
	// 				Interval: to.Ptr[int32](4),
	// 				Multiplier: to.Ptr[int32](13),
	// 			},
	// 			DefaultRouteOriginate: to.Ptr(armmanagednetworkfabric.BooleanEnumPropertyTrue),
	// 			FabricASN: to.Ptr[int32](21),
	// 			IPv4ListenRangePrefixes: []*string{
	// 				to.Ptr("10.1.0.0/25")},
	// 				IPv4NeighborAddress: []*armmanagednetworkfabric.NeighborAddress{
	// 					{
	// 						Address: to.Ptr("10.1.0.0"),
	// 						OperationalState: to.Ptr("Provisioned"),
	// 				}},
	// 				IPv6ListenRangePrefixes: []*string{
	// 					to.Ptr("2fff::/66")},
	// 					IPv6NeighborAddress: []*armmanagednetworkfabric.NeighborAddress{
	// 						{
	// 							Address: to.Ptr("2fff::"),
	// 							OperationalState: to.Ptr("Provisioned"),
	// 					}},
	// 					PeerASN: to.Ptr[int32](6),
	// 				},
	// 				ConnectedIPv4Subnets: []*armmanagednetworkfabric.ConnectedSubnet{
	// 					{
	// 						Prefix: to.Ptr("10.0.0.0/24"),
	// 				}},
	// 				ConnectedIPv6Subnets: []*armmanagednetworkfabric.ConnectedSubnet{
	// 					{
	// 						Prefix: to.Ptr("3FFE:FFFF:0:CD30::a0/29"),
	// 				}},
	// 				ExportRoutePolicyID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.ManagedNetworkFabric/routePolicies/routePolicyName2"),
	// 				ImportRoutePolicyID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.ManagedNetworkFabric/routePolicies/routePolicyName1"),
	// 				Mtu: to.Ptr[int32](1500),
	// 				StaticRouteConfiguration: &armmanagednetworkfabric.StaticRouteConfiguration{
	// 					BfdConfiguration: &armmanagednetworkfabric.BfdConfiguration{
	// 						AdministrativeState: to.Ptr(armmanagednetworkfabric.EnabledDisabledStateEnabled),
	// 						Interval: to.Ptr[int32](4),
	// 						Multiplier: to.Ptr[int32](13),
	// 					},
	// 					IPv4Routes: []*armmanagednetworkfabric.StaticRouteProperties{
	// 						{
	// 							NextHop: []*string{
	// 								to.Ptr("10.0.0.1")},
	// 								Prefix: to.Ptr("10.1.0.0/24"),
	// 						}},
	// 						IPv6Routes: []*armmanagednetworkfabric.StaticRouteProperties{
	// 							{
	// 								NextHop: []*string{
	// 									to.Ptr("2ffe::1")},
	// 									Prefix: to.Ptr("2fff::/64"),
	// 							}},
	// 						},
	// 						AdministrativeState: to.Ptr(armmanagednetworkfabric.EnabledDisabledStateEnabled),
	// 						BfdDisabledOnResources: []*string{
	// 							to.Ptr("string")},
	// 							BfdForStaticRoutesDisabledOnResources: []*string{
	// 								to.Ptr("string")},
	// 								BgpDisabledOnResources: []*string{
	// 									to.Ptr("string")},
	// 									DisabledOnResources: []*string{
	// 										to.Ptr("string")},
	// 										ProvisioningState: to.Ptr(armmanagednetworkfabric.ProvisioningStateSucceeded),
	// 										VlanID: to.Ptr[int32](501),
	// 									},
	// 								}
}
