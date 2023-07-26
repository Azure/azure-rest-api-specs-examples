package armmanagednetworkfabric_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managednetworkfabric/armmanagednetworkfabric"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/925ba149e17454ce91ecd3f9f4134effb2f97844/specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/RoutePolicies_Get_MaximumSet_Gen.json
func ExampleRoutePoliciesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagednetworkfabric.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRoutePoliciesClient().Get(ctx, "example-rg", "example-routePolicy", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.RoutePolicy = armmanagednetworkfabric.RoutePolicy{
	// 	Name: to.Ptr("example-routePolicy"),
	// 	Type: to.Ptr("microsoft.managednetworkfabric/routePolicies"),
	// 	ID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/routePolicies/example-routePolicy"),
	// 	SystemData: &armmanagednetworkfabric.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-11T17:48:22.837Z"); return t}()),
	// 		CreatedBy: to.Ptr("user@mail.com"),
	// 		CreatedByType: to.Ptr(armmanagednetworkfabric.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-11T17:48:22.837Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("user@mail.com"),
	// 		LastModifiedByType: to.Ptr(armmanagednetworkfabric.CreatedByTypeUser),
	// 	},
	// 	Location: to.Ptr("eastus"),
	// 	Tags: map[string]*string{
	// 		"keyID": to.Ptr("keyValue"),
	// 	},
	// 	Properties: &armmanagednetworkfabric.RoutePolicyProperties{
	// 		Annotation: to.Ptr("annotation"),
	// 		Statements: []*armmanagednetworkfabric.RoutePolicyStatementProperties{
	// 			{
	// 				Annotation: to.Ptr("annotation"),
	// 				Action: &armmanagednetworkfabric.StatementActionProperties{
	// 					ActionType: to.Ptr(armmanagednetworkfabric.RoutePolicyActionTypePermit),
	// 					IPCommunityProperties: &armmanagednetworkfabric.ActionIPCommunityProperties{
	// 						Add: &armmanagednetworkfabric.IPCommunityIDList{
	// 							IPCommunityIDs: []*string{
	// 								to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/ipCommunities/example-ipCommunity")},
	// 							},
	// 							Delete: &armmanagednetworkfabric.IPCommunityIDList{
	// 								IPCommunityIDs: []*string{
	// 									to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/ipCommunities/example-ipCommunity")},
	// 								},
	// 								Set: &armmanagednetworkfabric.IPCommunityIDList{
	// 									IPCommunityIDs: []*string{
	// 										to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/ipCommunities/example-ipCommunity")},
	// 									},
	// 								},
	// 								IPExtendedCommunityProperties: &armmanagednetworkfabric.ActionIPExtendedCommunityProperties{
	// 									Add: &armmanagednetworkfabric.IPExtendedCommunityIDList{
	// 										IPExtendedCommunityIDs: []*string{
	// 											to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/ipExtendedCommunities/example-ipExtendedCommunity")},
	// 										},
	// 										Delete: &armmanagednetworkfabric.IPExtendedCommunityIDList{
	// 											IPExtendedCommunityIDs: []*string{
	// 												to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/ipExtendedCommunities/example-ipExtendedCommunity")},
	// 											},
	// 											Set: &armmanagednetworkfabric.IPExtendedCommunityIDList{
	// 												IPExtendedCommunityIDs: []*string{
	// 													to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/ipExtendedCommunities/example-ipExtendedCommunity")},
	// 												},
	// 											},
	// 											LocalPreference: to.Ptr[int64](20),
	// 										},
	// 										Condition: &armmanagednetworkfabric.StatementConditionProperties{
	// 											IPCommunityIDs: []*string{
	// 												to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/ipCommunities/example-ipCommunity")},
	// 												IPExtendedCommunityIDs: []*string{
	// 													to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/ipExtendedCommunities/example-ipExtendedCommunity")},
	// 													Type: to.Ptr(armmanagednetworkfabric.RoutePolicyConditionTypeOr),
	// 													IPPrefixID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/ipPrefixes/example-ipPrefix"),
	// 												},
	// 												SequenceNumber: to.Ptr[int64](7),
	// 										}},
	// 										AddressFamilyType: to.Ptr(armmanagednetworkfabric.AddressFamilyTypeIPv4),
	// 										AdministrativeState: to.Ptr(armmanagednetworkfabric.AdministrativeStateEnabled),
	// 										ConfigurationState: to.Ptr(armmanagednetworkfabric.ConfigurationStateSucceeded),
	// 										NetworkFabricID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/networkFabrics/example-fabric"),
	// 										ProvisioningState: to.Ptr(armmanagednetworkfabric.ProvisioningStateSucceeded),
	// 									},
	// 								}
}
