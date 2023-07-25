package armmanagednetworkfabric_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managednetworkfabric/armmanagednetworkfabric"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/925ba149e17454ce91ecd3f9f4134effb2f97844/specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/L3IsolationDomains_Create_MaximumSet_Gen.json
func ExampleL3IsolationDomainsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagednetworkfabric.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewL3IsolationDomainsClient().BeginCreate(ctx, "example-rg", "example-l3domain", armmanagednetworkfabric.L3IsolationDomain{
		Location: to.Ptr("eastus"),
		Tags: map[string]*string{
			"keyID": to.Ptr("KeyValue"),
		},
		Properties: &armmanagednetworkfabric.L3IsolationDomainProperties{
			Annotation: to.Ptr("annotation"),
			AggregateRouteConfiguration: &armmanagednetworkfabric.AggregateRouteConfiguration{
				IPv4Routes: []*armmanagednetworkfabric.AggregateRoute{
					{
						Prefix: to.Ptr("10.0.0.0/24"),
					}},
				IPv6Routes: []*armmanagednetworkfabric.AggregateRoute{
					{
						Prefix: to.Ptr("3FFE:FFFF:0:CD30::a0/29"),
					}},
			},
			ConnectedSubnetRoutePolicy: &armmanagednetworkfabric.ConnectedSubnetRoutePolicy{
				ExportRoutePolicy: &armmanagednetworkfabric.L3ExportRoutePolicy{
					ExportIPv4RoutePolicyID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/routePolicies/example-routePolicy"),
					ExportIPv6RoutePolicyID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/routePolicies/example-routePolicy"),
				},
				ExportRoutePolicyID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/routePolicies/routePolicyName"),
			},
			RedistributeConnectedSubnets: to.Ptr(armmanagednetworkfabric.RedistributeConnectedSubnetsTrue),
			RedistributeStaticRoutes:     to.Ptr(armmanagednetworkfabric.RedistributeStaticRoutesFalse),
			NetworkFabricID:              to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/networkFabrics/example-fabric"),
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
	// res.L3IsolationDomain = armmanagednetworkfabric.L3IsolationDomain{
	// 	Name: to.Ptr("example-l3domain"),
	// 	Type: to.Ptr("microsoft.managednetworkfabric/l3isolationdomains"),
	// 	ID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/l3IsolationDomains/example-l3domain"),
	// 	SystemData: &armmanagednetworkfabric.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-05T18:37:10.310Z"); return t}()),
	// 		CreatedBy: to.Ptr("email@address.com"),
	// 		CreatedByType: to.Ptr(armmanagednetworkfabric.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-05T18:37:10.310Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("UserId"),
	// 		LastModifiedByType: to.Ptr(armmanagednetworkfabric.CreatedByTypeUser),
	// 	},
	// 	Location: to.Ptr("eastus"),
	// 	Tags: map[string]*string{
	// 		"keyID": to.Ptr("KeyValue"),
	// 	},
	// 	Properties: &armmanagednetworkfabric.L3IsolationDomainProperties{
	// 		Annotation: to.Ptr("annotation"),
	// 		AggregateRouteConfiguration: &armmanagednetworkfabric.AggregateRouteConfiguration{
	// 			IPv4Routes: []*armmanagednetworkfabric.AggregateRoute{
	// 				{
	// 					Prefix: to.Ptr("10.0.0.0/24"),
	// 			}},
	// 			IPv6Routes: []*armmanagednetworkfabric.AggregateRoute{
	// 				{
	// 					Prefix: to.Ptr("3FFE:FFFF:0:CD30::a0/29"),
	// 			}},
	// 		},
	// 		ConnectedSubnetRoutePolicy: &armmanagednetworkfabric.ConnectedSubnetRoutePolicy{
	// 			ExportRoutePolicy: &armmanagednetworkfabric.L3ExportRoutePolicy{
	// 				ExportIPv4RoutePolicyID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/routePolicies/example-routePolicy"),
	// 				ExportIPv6RoutePolicyID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/routePolicies/example-routePolicy"),
	// 			},
	// 			ExportRoutePolicyID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/routePolicies/routePolicyName"),
	// 		},
	// 		RedistributeConnectedSubnets: to.Ptr(armmanagednetworkfabric.RedistributeConnectedSubnetsTrue),
	// 		RedistributeStaticRoutes: to.Ptr(armmanagednetworkfabric.RedistributeStaticRoutesFalse),
	// 		AdministrativeState: to.Ptr(armmanagednetworkfabric.AdministrativeStateEnabled),
	// 		ConfigurationState: to.Ptr(armmanagednetworkfabric.ConfigurationStateSucceeded),
	// 		NetworkFabricID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/networkFabrics/example-fabric"),
	// 		ProvisioningState: to.Ptr(armmanagednetworkfabric.ProvisioningStateSucceeded),
	// 	},
	// }
}
