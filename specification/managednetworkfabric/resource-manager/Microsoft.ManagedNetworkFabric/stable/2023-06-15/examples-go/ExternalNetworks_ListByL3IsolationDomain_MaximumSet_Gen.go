package armmanagednetworkfabric_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managednetworkfabric/armmanagednetworkfabric"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/925ba149e17454ce91ecd3f9f4134effb2f97844/specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/ExternalNetworks_ListByL3IsolationDomain_MaximumSet_Gen.json
func ExampleExternalNetworksClient_NewListByL3IsolationDomainPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagednetworkfabric.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewExternalNetworksClient().NewListByL3IsolationDomainPager("example-rg", "example-l3domain", nil)
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
		// page.ExternalNetworksList = armmanagednetworkfabric.ExternalNetworksList{
		// 	Value: []*armmanagednetworkfabric.ExternalNetwork{
		// 		{
		// 			Name: to.Ptr("example-externalnetwork"),
		// 			Type: to.Ptr("microsoft.managednetworkfabric/l3IsolationDomains/externalnetworks"),
		// 			ID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/l3IsolationDomains/example-l3domain/externalNetworks/example-externalnetwork"),
		// 			SystemData: &armmanagednetworkfabric.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-09T04:51:41.251Z"); return t}()),
		// 				CreatedBy: to.Ptr("email@address.com"),
		// 				CreatedByType: to.Ptr(armmanagednetworkfabric.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-09T04:51:41.251Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("UserId"),
		// 				LastModifiedByType: to.Ptr(armmanagednetworkfabric.CreatedByTypeUser),
		// 			},
		// 			Properties: &armmanagednetworkfabric.ExternalNetworkProperties{
		// 				Annotation: to.Ptr("annotation"),
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
		// 				AdministrativeState: to.Ptr(armmanagednetworkfabric.AdministrativeStateEnabled),
		// 				ConfigurationState: to.Ptr(armmanagednetworkfabric.ConfigurationStateSucceeded),
		// 				NetworkToNetworkInterconnectID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/networkFabrics/example-fabric/networkToNetworkInterconnects/example-nni"),
		// 				OptionAProperties: &armmanagednetworkfabric.ExternalNetworkPropertiesOptionAProperties{
		// 					BfdConfiguration: &armmanagednetworkfabric.BfdConfiguration{
		// 						AdministrativeState: to.Ptr(armmanagednetworkfabric.BfdAdministrativeStateEnabled),
		// 						IntervalInMilliSeconds: to.Ptr[int32](300),
		// 						Multiplier: to.Ptr[int32](15),
		// 					},
		// 					EgressACLID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/accessControlLists/example-acl"),
		// 					FabricASN: to.Ptr[int64](1234),
		// 					IngressACLID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/accessControlLists/example-acl"),
		// 					Mtu: to.Ptr[int32](1500),
		// 					PeerASN: to.Ptr[int64](65047),
		// 					VlanID: to.Ptr[int32](1001),
		// 					PrimaryIPv4Prefix: to.Ptr("10.1.1.0/30"),
		// 					PrimaryIPv6Prefix: to.Ptr("3FFE:FFFF:0:CD30::a0/126"),
		// 					SecondaryIPv4Prefix: to.Ptr("10.1.1.4/30"),
		// 					SecondaryIPv6Prefix: to.Ptr("3FFE:FFFF:0:CD30::a4/126"),
		// 				},
		// 				OptionBProperties: &armmanagednetworkfabric.L3OptionBProperties{
		// 					ExportRouteTargets: []*string{
		// 						to.Ptr("65046:10039")},
		// 						ImportRouteTargets: []*string{
		// 							to.Ptr("65046:10039")},
		// 							RouteTargets: &armmanagednetworkfabric.RouteTargetInformation{
		// 								ExportIPv4RouteTargets: []*string{
		// 									to.Ptr("65046:10039")},
		// 									ExportIPv6RouteTargets: []*string{
		// 										to.Ptr("65046:10039")},
		// 										ImportIPv4RouteTargets: []*string{
		// 											to.Ptr("65046:10039")},
		// 											ImportIPv6RouteTargets: []*string{
		// 												to.Ptr("65046:10039")},
		// 											},
		// 										},
		// 										PeeringOption: to.Ptr(armmanagednetworkfabric.PeeringOptionOptionA),
		// 										ProvisioningState: to.Ptr(armmanagednetworkfabric.ProvisioningStateSucceeded),
		// 									},
		// 							}},
		// 						}
	}
}
