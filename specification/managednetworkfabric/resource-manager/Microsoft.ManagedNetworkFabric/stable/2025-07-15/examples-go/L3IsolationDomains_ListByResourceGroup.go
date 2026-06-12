package armmanagednetworkfabric_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managednetworkfabric/armmanagednetworkfabric/v2"
)

// Generated from example definition: 2025-07-15/L3IsolationDomains_ListByResourceGroup.json
func ExampleL3IsolationDomainsClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagednetworkfabric.NewClientFactory("0000ABCD-0A0B-0000-0000-000000ABCDEF", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewL3IsolationDomainsClient().NewListByResourceGroupPager("example-rg", nil)
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
		// page = armmanagednetworkfabric.L3IsolationDomainsClientListByResourceGroupResponse{
		// 	L3IsolationDomainsListResult: armmanagednetworkfabric.L3IsolationDomainsListResult{
		// 		Value: []*armmanagednetworkfabric.L3IsolationDomain{
		// 			{
		// 				Properties: &armmanagednetworkfabric.L3IsolationDomainProperties{
		// 					Annotation: to.Ptr("annotation"),
		// 					RedistributeConnectedSubnets: to.Ptr(armmanagednetworkfabric.RedistributeConnectedSubnetsTrue),
		// 					RedistributeStaticRoutes: to.Ptr(armmanagednetworkfabric.RedistributeStaticRoutesTrue),
		// 					AggregateRouteConfiguration: &armmanagednetworkfabric.AggregateRouteConfiguration{
		// 						IPv4Routes: []*armmanagednetworkfabric.AggregateRoute{
		// 							{
		// 								Prefix: to.Ptr("10.0.0.0/24"),
		// 							},
		// 						},
		// 						IPv6Routes: []*armmanagednetworkfabric.AggregateRoute{
		// 							{
		// 								Prefix: to.Ptr("3FFE:FFFF:0:CD30::a0/29"),
		// 							},
		// 						},
		// 					},
		// 					ConnectedSubnetRoutePolicy: &armmanagednetworkfabric.ConnectedSubnetRoutePolicy{
		// 						ExportRoutePolicy: &armmanagednetworkfabric.L3ExportRoutePolicy{
		// 							ExportIPv4RoutePolicyID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/routePolicies/example-routePolicy"),
		// 							ExportIPv6RoutePolicyID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/routePolicies/example-routePolicy"),
		// 						},
		// 					},
		// 					StaticRouteRoutePolicy: &armmanagednetworkfabric.StaticRouteRoutePolicy{
		// 						ExportRoutePolicy: &armmanagednetworkfabric.L3ExportRoutePolicy{
		// 							ExportIPv4RoutePolicyID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/routePolicies/example-routePolicy"),
		// 							ExportIPv6RoutePolicyID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/routePolicies/example-routePolicy"),
		// 						},
		// 					},
		// 					NetworkFabricID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/networkFabrics/example-fabric"),
		// 					UniqueRdConfiguration: &armmanagednetworkfabric.L3UniqueRouteDistinguisherProperties{
		// 						UniqueRds: []*string{
		// 							to.Ptr("uniquerd-1"),
		// 						},
		// 					},
		// 					V4RoutePrefixLimit: &armmanagednetworkfabric.RoutePrefixLimitProperties{
		// 						HardLimit: to.Ptr[int32](1000),
		// 						Threshold: to.Ptr[int32](50),
		// 					},
		// 					V6RoutePrefixLimit: &armmanagednetworkfabric.RoutePrefixLimitProperties{
		// 						HardLimit: to.Ptr[int32](1000),
		// 						Threshold: to.Ptr[int32](50),
		// 					},
		// 					LastOperation: &armmanagednetworkfabric.LastOperationProperties{
		// 						Details: to.Ptr("Succeeded"),
		// 					},
		// 					ConfigurationState: to.Ptr(armmanagednetworkfabric.ConfigurationStateSucceeded),
		// 					ProvisioningState: to.Ptr(armmanagednetworkfabric.ProvisioningStateAccepted),
		// 					AdministrativeState: to.Ptr(armmanagednetworkfabric.AdministrativeStateEnabled),
		// 				},
		// 				Identity: &armmanagednetworkfabric.ManagedServiceIdentity{
		// 					PrincipalID: to.Ptr("0000ABCD-0A0B-0000-0000-000000ABCDEF"),
		// 					TenantID: to.Ptr("0000ABCD-0A0B-0000-0000-000000ABCDEF"),
		// 					Type: to.Ptr(armmanagednetworkfabric.ManagedServiceIdentityTypeNone),
		// 					UserAssignedIdentities: map[string]*armmanagednetworkfabric.UserAssignedIdentity{
		// 						"key3673": &armmanagednetworkfabric.UserAssignedIdentity{
		// 							PrincipalID: to.Ptr("0000ABCD-0A0B-0000-0000-000000ABCDEF"),
		// 							ClientID: to.Ptr("0000ABCD-0A0B-0000-0000-000000ABCDEF"),
		// 						},
		// 					},
		// 				},
		// 				Tags: map[string]*string{
		// 					"KeyId": to.Ptr("KeyValue"),
		// 				},
		// 				Location: to.Ptr("eastus"),
		// 				ID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/l3IsolationDomains/example-l3domain"),
		// 				Name: to.Ptr("example-l3domain"),
		// 				Type: to.Ptr("microsoft.managednetworkfabric/l3isolationdomains"),
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
