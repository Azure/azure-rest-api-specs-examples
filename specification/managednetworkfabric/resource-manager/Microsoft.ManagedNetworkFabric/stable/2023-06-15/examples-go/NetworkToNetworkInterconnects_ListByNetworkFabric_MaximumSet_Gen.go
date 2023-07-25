package armmanagednetworkfabric_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managednetworkfabric/armmanagednetworkfabric"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/925ba149e17454ce91ecd3f9f4134effb2f97844/specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/NetworkToNetworkInterconnects_ListByNetworkFabric_MaximumSet_Gen.json
func ExampleNetworkToNetworkInterconnectsClient_NewListByNetworkFabricPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagednetworkfabric.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewNetworkToNetworkInterconnectsClient().NewListByNetworkFabricPager("example-rg", "example-fabric", nil)
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
		// page.NetworkToNetworkInterconnectsList = armmanagednetworkfabric.NetworkToNetworkInterconnectsList{
		// 	Value: []*armmanagednetworkfabric.NetworkToNetworkInterconnect{
		// 		{
		// 			Name: to.Ptr("example-nni"),
		// 			Type: to.Ptr("microsoft.managedNetworkFabric/networkFabrics/networkToNetworkInterconnects"),
		// 			ID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/networkFabrics/example-fabric/networkToNetworkInterconnects/example-nni"),
		// 			SystemData: &armmanagednetworkfabric.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-07T09:53:31.314Z"); return t}()),
		// 				CreatedBy: to.Ptr("user@mail.com"),
		// 				CreatedByType: to.Ptr(armmanagednetworkfabric.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-07T09:53:31.314Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("user@mail.com"),
		// 				LastModifiedByType: to.Ptr(armmanagednetworkfabric.CreatedByTypeUser),
		// 			},
		// 			Properties: &armmanagednetworkfabric.NetworkToNetworkInterconnectProperties{
		// 				AdministrativeState: to.Ptr(armmanagednetworkfabric.AdministrativeStateEnabled),
		// 				ConfigurationState: to.Ptr(armmanagednetworkfabric.ConfigurationStateSucceeded),
		// 				EgressACLID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/accessControlLists/example-acl"),
		// 				ExportRoutePolicy: &armmanagednetworkfabric.ExportRoutePolicyInformation{
		// 					ExportIPv4RoutePolicyID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/routePolicies/example-routePolicy"),
		// 					ExportIPv6RoutePolicyID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/routePolicies/example-routePolicy"),
		// 				},
		// 				ImportRoutePolicy: &armmanagednetworkfabric.ImportRoutePolicyInformation{
		// 					ImportIPv4RoutePolicyID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/routePolicies/example-routePolicy"),
		// 					ImportIPv6RoutePolicyID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/routePolicies/example-routePolicy"),
		// 				},
		// 				IngressACLID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/accessControlLists/example-acl"),
		// 				IsManagementType: to.Ptr(armmanagednetworkfabric.IsManagementTypeTrue),
		// 				Layer2Configuration: &armmanagednetworkfabric.Layer2Configuration{
		// 					Interfaces: []*string{
		// 						to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/networkDevices/example-networkDevice/networkInterfaces/example-networkInterface")},
		// 						Mtu: to.Ptr[int32](1500),
		// 					},
		// 					NniType: to.Ptr(armmanagednetworkfabric.NniTypeCE),
		// 					NpbStaticRouteConfiguration: &armmanagednetworkfabric.NpbStaticRouteConfiguration{
		// 						BfdConfiguration: &armmanagednetworkfabric.BfdConfiguration{
		// 							AdministrativeState: to.Ptr(armmanagednetworkfabric.BfdAdministrativeStateEnabled),
		// 							IntervalInMilliSeconds: to.Ptr[int32](300),
		// 							Multiplier: to.Ptr[int32](25),
		// 						},
		// 						IPv4Routes: []*armmanagednetworkfabric.StaticRouteProperties{
		// 							{
		// 								NextHop: []*string{
		// 									to.Ptr("21.20.20.20")},
		// 									Prefix: to.Ptr("20.0.0.12/30"),
		// 							}},
		// 							IPv6Routes: []*armmanagednetworkfabric.StaticRouteProperties{
		// 								{
		// 									NextHop: []*string{
		// 										to.Ptr("4FFE:FFFF:0:CD30::ac")},
		// 										Prefix: to.Ptr("3FFE:FFFF:0:CD30::ac/127"),
		// 								}},
		// 							},
		// 							OptionBLayer3Configuration: &armmanagednetworkfabric.NetworkToNetworkInterconnectPropertiesOptionBLayer3Configuration{
		// 								PrimaryIPv4Prefix: to.Ptr("10.0.0.12/30"),
		// 								PrimaryIPv6Prefix: to.Ptr("4FFE:FFFF:0:CD30::a8/127"),
		// 								SecondaryIPv4Prefix: to.Ptr("40.0.0.14/30"),
		// 								SecondaryIPv6Prefix: to.Ptr("6FFE:FFFF:0:CD30::ac/127"),
		// 								FabricASN: to.Ptr[int64](17),
		// 								PeerASN: to.Ptr[int64](61234),
		// 								VlanID: to.Ptr[int32](1234),
		// 							},
		// 							ProvisioningState: to.Ptr(armmanagednetworkfabric.ProvisioningStateSucceeded),
		// 							UseOptionB: to.Ptr(armmanagednetworkfabric.BooleanEnumPropertyTrue),
		// 						},
		// 				}},
		// 			}
	}
}
