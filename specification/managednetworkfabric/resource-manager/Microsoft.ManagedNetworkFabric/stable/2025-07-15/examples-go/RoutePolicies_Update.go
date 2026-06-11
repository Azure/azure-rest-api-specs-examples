package armmanagednetworkfabric_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managednetworkfabric/armmanagednetworkfabric/v2"
)

// Generated from example definition: 2025-07-15/RoutePolicies_Update.json
func ExampleRoutePoliciesClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagednetworkfabric.NewClientFactory("0000ABCD-0A0B-0000-0000-000000ABCDEF", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewRoutePoliciesClient().BeginUpdate(ctx, "example-rg", "example-routePolicy", armmanagednetworkfabric.RoutePolicyPatch{
		Tags: map[string]*string{
			"keyId": to.Ptr("keyValue"),
		},
		Properties: &armmanagednetworkfabric.RoutePolicyPatchableProperties{
			DefaultAction: to.Ptr(armmanagednetworkfabric.CommunityActionTypesPermit),
			Statements: []*armmanagednetworkfabric.RoutePolicyStatementPatchProperties{
				{
					Annotation:     to.Ptr("annotation"),
					SequenceNumber: to.Ptr[int64](7),
					Condition: &armmanagednetworkfabric.StatementConditionPatchProperties{
						IPCommunityIDs: []*string{
							to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/ipCommunities/example-ipCommunity"),
						},
						IPExtendedCommunityIDs: []*string{
							to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/ipExtendedCommunities/example-ipExtendedCommunity"),
						},
						Type:       to.Ptr(armmanagednetworkfabric.RoutePolicyConditionTypeOr),
						IPPrefixID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/ipPrefixes/example-ipPrefix"),
					},
					Action: &armmanagednetworkfabric.StatementActionPatchProperties{
						LocalPreference: to.Ptr[int64](20),
						ActionType:      to.Ptr(armmanagednetworkfabric.RoutePolicyActionTypePermit),
						IPCommunityProperties: &armmanagednetworkfabric.ActionIPCommunityPatchProperties{
							Add: &armmanagednetworkfabric.IPCommunityIDList{
								IPCommunityIDs: []*string{
									to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/ipCommunities/example-ipCommunity"),
								},
							},
							Delete: &armmanagednetworkfabric.IPCommunityIDList{
								IPCommunityIDs: []*string{
									to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/ipCommunities/example-ipCommunity"),
								},
							},
							Set: &armmanagednetworkfabric.IPCommunityIDList{
								IPCommunityIDs: []*string{
									to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/ipCommunities/example-ipCommunity"),
								},
							},
						},
						IPExtendedCommunityProperties: &armmanagednetworkfabric.ActionIPExtendedCommunityPatchProperties{
							Add: &armmanagednetworkfabric.IPExtendedCommunityIDList{
								IPExtendedCommunityIDs: []*string{
									to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/ipExtendedCommunities/example-ipExtendedCommunity"),
								},
							},
							Delete: &armmanagednetworkfabric.IPExtendedCommunityIDList{
								IPExtendedCommunityIDs: []*string{
									to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/ipExtendedCommunities/example-ipExtendedCommunity"),
								},
							},
							Set: &armmanagednetworkfabric.IPExtendedCommunityIDList{
								IPExtendedCommunityIDs: []*string{
									to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/ipExtendedCommunities/example-ipExtendedCommunity"),
								},
							},
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
	// res = armmanagednetworkfabric.RoutePoliciesClientUpdateResponse{
	// 	RoutePolicy: &armmanagednetworkfabric.RoutePolicy{
	// 		Properties: &armmanagednetworkfabric.RoutePolicyProperties{
	// 			Annotation: to.Ptr("annotation"),
	// 			DefaultAction: to.Ptr(armmanagednetworkfabric.CommunityActionTypesPermit),
	// 			Statements: []*armmanagednetworkfabric.RoutePolicyStatementProperties{
	// 				{
	// 					Annotation: to.Ptr("annotation"),
	// 					SequenceNumber: to.Ptr[int64](7),
	// 					Condition: &armmanagednetworkfabric.StatementConditionProperties{
	// 						IPCommunityIDs: []*string{
	// 							to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/ipCommunities/example-ipCommunity"),
	// 						},
	// 						IPExtendedCommunityIDs: []*string{
	// 							to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/ipExtendedCommunities/example-ipExtendedCommunity"),
	// 						},
	// 						Type: to.Ptr(armmanagednetworkfabric.RoutePolicyConditionTypeOr),
	// 						IPPrefixID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/ipPrefixes/example-ipPrefix"),
	// 					},
	// 					Action: &armmanagednetworkfabric.StatementActionProperties{
	// 						LocalPreference: to.Ptr[int64](20),
	// 						ActionType: to.Ptr(armmanagednetworkfabric.RoutePolicyActionTypePermit),
	// 						IPCommunityProperties: &armmanagednetworkfabric.ActionIPCommunityProperties{
	// 							Add: &armmanagednetworkfabric.IPCommunityIDList{
	// 								IPCommunityIDs: []*string{
	// 									to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/ipCommunities/example-ipCommunity"),
	// 								},
	// 							},
	// 							Delete: &armmanagednetworkfabric.IPCommunityIDList{
	// 								IPCommunityIDs: []*string{
	// 									to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/ipCommunities/example-ipCommunity"),
	// 								},
	// 							},
	// 							Set: &armmanagednetworkfabric.IPCommunityIDList{
	// 								IPCommunityIDs: []*string{
	// 									to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/ipCommunities/example-ipCommunity"),
	// 								},
	// 							},
	// 						},
	// 						IPExtendedCommunityProperties: &armmanagednetworkfabric.ActionIPExtendedCommunityProperties{
	// 							Add: &armmanagednetworkfabric.IPExtendedCommunityIDList{
	// 								IPExtendedCommunityIDs: []*string{
	// 									to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/ipExtendedCommunities/example-ipExtendedCommunity"),
	// 								},
	// 							},
	// 							Delete: &armmanagednetworkfabric.IPExtendedCommunityIDList{
	// 								IPExtendedCommunityIDs: []*string{
	// 									to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/ipExtendedCommunities/example-ipExtendedCommunity"),
	// 								},
	// 							},
	// 							Set: &armmanagednetworkfabric.IPExtendedCommunityIDList{
	// 								IPExtendedCommunityIDs: []*string{
	// 									to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/ipExtendedCommunities/example-ipExtendedCommunity"),
	// 								},
	// 							},
	// 						},
	// 					},
	// 				},
	// 			},
	// 			NetworkFabricID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/networkFabrics/example-fabric"),
	// 			AddressFamilyType: to.Ptr(armmanagednetworkfabric.AddressFamilyTypeIPv4),
	// 			ConfigurationState: to.Ptr(armmanagednetworkfabric.ConfigurationStateSucceeded),
	// 			ProvisioningState: to.Ptr(armmanagednetworkfabric.ProvisioningStateAccepted),
	// 			AdministrativeState: to.Ptr(armmanagednetworkfabric.AdministrativeStateEnabled),
	// 		},
	// 		Tags: map[string]*string{
	// 			"keyId": to.Ptr("keyValue"),
	// 		},
	// 		Location: to.Ptr("eastus"),
	// 		ID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/routePolicies/example-routePolicy"),
	// 		Name: to.Ptr("example-routePolicy"),
	// 		Type: to.Ptr("microsoft.managednetworkfabric/routePolicies"),
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
