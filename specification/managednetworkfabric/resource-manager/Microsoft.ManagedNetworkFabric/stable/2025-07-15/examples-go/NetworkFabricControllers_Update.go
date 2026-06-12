package armmanagednetworkfabric_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managednetworkfabric/armmanagednetworkfabric/v2"
)

// Generated from example definition: 2025-07-15/NetworkFabricControllers_Update.json
func ExampleNetworkFabricControllersClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagednetworkfabric.NewClientFactory("0000ABCD-0A0B-0000-0000-000000ABCDEF", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewNetworkFabricControllersClient().BeginUpdate(ctx, "example-rg", "example-networkController", armmanagednetworkfabric.NetworkFabricControllerPatch{
		Tags: map[string]*string{
			"keyId": to.Ptr("KeyValue"),
		},
		Identity: &armmanagednetworkfabric.ManagedServiceIdentityPatch{
			Type: to.Ptr(armmanagednetworkfabric.ManagedServiceIdentityTypeNone),
			UserAssignedIdentities: map[string]*armmanagednetworkfabric.UserAssignedIdentity{
				"key1402": {},
			},
		},
		Properties: &armmanagednetworkfabric.NetworkFabricControllerPatchProperties{
			InfrastructureExpressRouteConnections: []*armmanagednetworkfabric.ExpressRouteConnectionInformation{
				{
					ExpressRouteCircuitID:        to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.Network/expressRouteCircuits/expressRouteCircuitName"),
					ExpressRouteAuthorizationKey: to.Ptr("xxx-xxx-xxx"),
				},
			},
			WorkloadExpressRouteConnections: []*armmanagednetworkfabric.ExpressRouteConnectionInformation{
				{
					ExpressRouteCircuitID:        to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.Network/expressRouteCircuits/expressRouteCircuitName"),
					ExpressRouteAuthorizationKey: to.Ptr("xxx-xxx-xxx"),
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
	// res = armmanagednetworkfabric.NetworkFabricControllersClientUpdateResponse{
	// 	NetworkFabricController: &armmanagednetworkfabric.NetworkFabricController{
	// 		Properties: &armmanagednetworkfabric.NetworkFabricControllerProperties{
	// 			Annotation: to.Ptr("annotation"),
	// 			InfrastructureExpressRouteConnections: []*armmanagednetworkfabric.ExpressRouteConnectionInformation{
	// 				{
	// 					ExpressRouteCircuitID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.Network/expressRouteCircuits/expressRouteCircuitName"),
	// 				},
	// 			},
	// 			WorkloadExpressRouteConnections: []*armmanagednetworkfabric.ExpressRouteConnectionInformation{
	// 				{
	// 					ExpressRouteCircuitID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.Network/expressRouteCircuits/expressRouteCircuitName"),
	// 				},
	// 			},
	// 			InfrastructureServices: &armmanagednetworkfabric.ControllerServices{
	// 				IPv4AddressSpaces: []*string{
	// 					to.Ptr("172.253.0.0/19"),
	// 				},
	// 				IPv6AddressSpaces: []*string{
	// 				},
	// 			},
	// 			WorkloadServices: &armmanagednetworkfabric.ControllerServices{
	// 				IPv4AddressSpaces: []*string{
	// 					to.Ptr("172.253.28.0/22"),
	// 				},
	// 				IPv6AddressSpaces: []*string{
	// 				},
	// 			},
	// 			ManagedResourceGroupConfiguration: &armmanagednetworkfabric.ManagedResourceGroupConfiguration{
	// 				Name: to.Ptr("managedResourceGroupName"),
	// 				Location: to.Ptr("eastus"),
	// 			},
	// 			NetworkFabricIDs: []*string{
	// 				to.Ptr("/subscriptions/xxxxxx/resourceGroups/resourcegroupname/providers/Microsoft.ManagedNetworkFabric/networkFabrics/eaxmple-fabric"),
	// 			},
	// 			IsWorkloadManagementNetworkEnabled: to.Ptr(armmanagednetworkfabric.IsWorkloadManagementNetworkEnabledTrue),
	// 			TenantInternetGatewayIDs: []*string{
	// 				to.Ptr("/subscriptions/xxxxxx/resourceGroups/resourcegroupname/providers/Microsoft.ManagedNetworkFabric/internetGateways/example-internetGateway"),
	// 			},
	// 			IPv4AddressSpace: to.Ptr("172.253.0.0/19"),
	// 			IPv6AddressSpace: to.Ptr("::/60"),
	// 			NfcSKU: to.Ptr(armmanagednetworkfabric.NfcSKUStandard),
	// 			ProvisioningState: to.Ptr(armmanagednetworkfabric.ProvisioningStateAccepted),
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
	// 			"keyId": to.Ptr("KeyValue"),
	// 		},
	// 		Location: to.Ptr("eastus"),
	// 		ID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/networkFabricControllers/example-networkController"),
	// 		Name: to.Ptr("example-networkController"),
	// 		Type: to.Ptr("microsoft.managednetworkfabric/networkfabriccontrollers"),
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
