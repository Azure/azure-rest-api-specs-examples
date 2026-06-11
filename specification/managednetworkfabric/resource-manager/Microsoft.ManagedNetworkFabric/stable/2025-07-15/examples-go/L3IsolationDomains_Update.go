package armmanagednetworkfabric_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managednetworkfabric/armmanagednetworkfabric/v2"
)

// Generated from example definition: 2025-07-15/L3IsolationDomains_Update.json
func ExampleL3IsolationDomainsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagednetworkfabric.NewClientFactory("0000ABCD-0A0B-0000-0000-000000ABCDEF", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewL3IsolationDomainsClient().BeginUpdate(ctx, "example-rg", "example-l3domain", armmanagednetworkfabric.L3IsolationDomainPatch{
		Tags: map[string]*string{
			"KeyId": to.Ptr("KeyValue"),
		},
		Identity: &armmanagednetworkfabric.ManagedServiceIdentityPatch{
			Type: to.Ptr(armmanagednetworkfabric.ManagedServiceIdentityTypeNone),
			UserAssignedIdentities: map[string]*armmanagednetworkfabric.UserAssignedIdentity{
				"key8793": {},
			},
		},
		Properties: &armmanagednetworkfabric.L3IsolationDomainPatchProperties{
			Annotation:                   to.Ptr("annotation1"),
			RedistributeConnectedSubnets: to.Ptr(armmanagednetworkfabric.RedistributeConnectedSubnetsTrue),
			RedistributeStaticRoutes:     to.Ptr(armmanagednetworkfabric.RedistributeStaticRoutesTrue),
			AggregateRouteConfiguration: &armmanagednetworkfabric.AggregateRoutePatchConfiguration{
				IPv4Routes: []*armmanagednetworkfabric.AggregateRoute{
					{
						Prefix: to.Ptr("10.0.0.0/24"),
					},
				},
				IPv6Routes: []*armmanagednetworkfabric.AggregateRoute{
					{
						Prefix: to.Ptr("3FFE:FFFF:0:CD30::a0/29"),
					},
				},
			},
			ConnectedSubnetRoutePolicy: &armmanagednetworkfabric.ConnectedSubnetRoutePolicyPatch{
				ExportRoutePolicy: &armmanagednetworkfabric.L3ExportRoutePolicyPatch{
					ExportIPv4RoutePolicyID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/routePolicies/example-routePolicy"),
					ExportIPv6RoutePolicyID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/routePolicies/example-routePolicy"),
				},
			},
			StaticRouteRoutePolicy: &armmanagednetworkfabric.StaticRouteRoutePolicyPatch{
				ExportRoutePolicy: &armmanagednetworkfabric.L3ExportRoutePolicyPatch{
					ExportIPv4RoutePolicyID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/routePolicies/example-routePolicy"),
					ExportIPv6RoutePolicyID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/routePolicies/example-routePolicy"),
				},
			},
			V4RoutePrefixLimit: &armmanagednetworkfabric.RoutePrefixLimitPatchProperties{
				HardLimit: to.Ptr[int32](1000),
				Threshold: to.Ptr[int32](50),
			},
			V6RoutePrefixLimit: &armmanagednetworkfabric.RoutePrefixLimitPatchProperties{
				HardLimit: to.Ptr[int32](1000),
				Threshold: to.Ptr[int32](50),
			},
			ExportPolicyConfiguration: &armmanagednetworkfabric.BmpExportPolicyPatchProperties{
				ExportPolicies: []*armmanagednetworkfabric.BmpExportPolicy{
					to.Ptr(armmanagednetworkfabric.BmpExportPolicyPrePolicy),
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
	// res = armmanagednetworkfabric.L3IsolationDomainsClientUpdateResponse{
	// 	L3IsolationDomain: &armmanagednetworkfabric.L3IsolationDomain{
	// 		Properties: &armmanagednetworkfabric.L3IsolationDomainProperties{
	// 			Annotation: to.Ptr("annotation"),
	// 			RedistributeConnectedSubnets: to.Ptr(armmanagednetworkfabric.RedistributeConnectedSubnetsTrue),
	// 			RedistributeStaticRoutes: to.Ptr(armmanagednetworkfabric.RedistributeStaticRoutesTrue),
	// 			AggregateRouteConfiguration: &armmanagednetworkfabric.AggregateRouteConfiguration{
	// 				IPv4Routes: []*armmanagednetworkfabric.AggregateRoute{
	// 					{
	// 						Prefix: to.Ptr("10.0.0.0/24"),
	// 					},
	// 				},
	// 				IPv6Routes: []*armmanagednetworkfabric.AggregateRoute{
	// 					{
	// 						Prefix: to.Ptr("3FFE:FFFF:0:CD30::a0/29"),
	// 					},
	// 				},
	// 			},
	// 			ConnectedSubnetRoutePolicy: &armmanagednetworkfabric.ConnectedSubnetRoutePolicy{
	// 				ExportRoutePolicy: &armmanagednetworkfabric.L3ExportRoutePolicy{
	// 					ExportIPv4RoutePolicyID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/routePolicies/example-routePolicy"),
	// 					ExportIPv6RoutePolicyID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/routePolicies/example-routePolicy"),
	// 				},
	// 			},
	// 			StaticRouteRoutePolicy: &armmanagednetworkfabric.StaticRouteRoutePolicy{
	// 				ExportRoutePolicy: &armmanagednetworkfabric.L3ExportRoutePolicy{
	// 					ExportIPv4RoutePolicyID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/routePolicies/example-routePolicy"),
	// 					ExportIPv6RoutePolicyID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/routePolicies/example-routePolicy"),
	// 				},
	// 			},
	// 			NetworkFabricID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/networkFabrics/example-fabric"),
	// 			UniqueRdConfiguration: &armmanagednetworkfabric.L3UniqueRouteDistinguisherProperties{
	// 				UniqueRds: []*string{
	// 					to.Ptr("uniquerd-1"),
	// 				},
	// 			},
	// 			V4RoutePrefixLimit: &armmanagednetworkfabric.RoutePrefixLimitProperties{
	// 				HardLimit: to.Ptr[int32](1000),
	// 				Threshold: to.Ptr[int32](90),
	// 			},
	// 			V6RoutePrefixLimit: &armmanagednetworkfabric.RoutePrefixLimitProperties{
	// 				HardLimit: to.Ptr[int32](1000),
	// 				Threshold: to.Ptr[int32](90),
	// 			},
	// 			ConfigurationState: to.Ptr(armmanagednetworkfabric.ConfigurationStateSucceeded),
	// 			ProvisioningState: to.Ptr(armmanagednetworkfabric.ProvisioningStateAccepted),
	// 			AdministrativeState: to.Ptr(armmanagednetworkfabric.AdministrativeStateEnabled),
	// 		},
	// 		Identity: &armmanagednetworkfabric.ManagedServiceIdentity{
	// 			PrincipalID: to.Ptr("0000ABCD-0A0B-0000-0000-000000ABCDEF"),
	// 			TenantID: to.Ptr("0000ABCD-0A0B-0000-0000-000000ABCDEF"),
	// 			Type: to.Ptr(armmanagednetworkfabric.ManagedServiceIdentityTypeNone),
	// 			UserAssignedIdentities: map[string]*armmanagednetworkfabric.UserAssignedIdentity{
	// 				"key4876": &armmanagednetworkfabric.UserAssignedIdentity{
	// 					PrincipalID: to.Ptr("0000ABCD-0A0B-0000-0000-000000ABCDEF"),
	// 					ClientID: to.Ptr("0000ABCD-0A0B-0000-0000-000000ABCDEF"),
	// 				},
	// 			},
	// 		},
	// 		Tags: map[string]*string{
	// 			"KeyId": to.Ptr("KeyValue"),
	// 		},
	// 		Location: to.Ptr("eastus"),
	// 		ID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/l3IsolationDomains/example-l3domain"),
	// 		Name: to.Ptr("example-l3domain"),
	// 		Type: to.Ptr("microsoft.managednetworkfabric/l3isolationdomains"),
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
