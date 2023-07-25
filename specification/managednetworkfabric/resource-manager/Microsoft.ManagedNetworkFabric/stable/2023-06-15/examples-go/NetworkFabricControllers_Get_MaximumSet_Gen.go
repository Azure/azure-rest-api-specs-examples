package armmanagednetworkfabric_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managednetworkfabric/armmanagednetworkfabric"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/925ba149e17454ce91ecd3f9f4134effb2f97844/specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/NetworkFabricControllers_Get_MaximumSet_Gen.json
func ExampleNetworkFabricControllersClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagednetworkfabric.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewNetworkFabricControllersClient().Get(ctx, "example-rg", "example-networkController", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.NetworkFabricController = armmanagednetworkfabric.NetworkFabricController{
	// 	Name: to.Ptr("example-networkController"),
	// 	Type: to.Ptr("microsoft.managednetworkfabric/networkfabriccontrollers"),
	// 	ID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/networkFabricControllers/example-networkController"),
	// 	SystemData: &armmanagednetworkfabric.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-05-29T05:17:40.665Z"); return t}()),
	// 		CreatedBy: to.Ptr("user@mail.com"),
	// 		CreatedByType: to.Ptr(armmanagednetworkfabric.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-05-29T05:17:40.665Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("user@mail.com"),
	// 		LastModifiedByType: to.Ptr(armmanagednetworkfabric.CreatedByTypeUser),
	// 	},
	// 	Location: to.Ptr("eastus"),
	// 	Properties: &armmanagednetworkfabric.NetworkFabricControllerProperties{
	// 		Annotation: to.Ptr("annotation"),
	// 		InfrastructureExpressRouteConnections: []*armmanagednetworkfabric.ExpressRouteConnectionInformation{
	// 			{
	// 				ExpressRouteCircuitID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.Network/expressRouteCircuits/expressRouteCircuitName"),
	// 		}},
	// 		WorkloadExpressRouteConnections: []*armmanagednetworkfabric.ExpressRouteConnectionInformation{
	// 			{
	// 				ExpressRouteCircuitID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.Network/expressRouteCircuits/expressRouteCircuitName"),
	// 		}},
	// 		InfrastructureServices: &armmanagednetworkfabric.ControllerServices{
	// 			IPv4AddressSpaces: []*string{
	// 				to.Ptr("172.253.0.0/19")},
	// 				IPv6AddressSpaces: []*string{
	// 				},
	// 			},
	// 			IPv4AddressSpace: to.Ptr("172.253.0.0/19"),
	// 			IPv6AddressSpace: to.Ptr("::/60"),
	// 			IsWorkloadManagementNetworkEnabled: to.Ptr(armmanagednetworkfabric.IsWorkloadManagementNetworkEnabledTrue),
	// 			ManagedResourceGroupConfiguration: &armmanagednetworkfabric.ManagedResourceGroupConfiguration{
	// 				Name: to.Ptr("managedResourceGroupName"),
	// 				Location: to.Ptr("eastus"),
	// 			},
	// 			NetworkFabricIDs: []*string{
	// 				to.Ptr("/subscriptions/xxxxxx/resourceGroups/resourcegroupname/providers/Microsoft.ManagedNetworkFabric/networkFabrics/eaxmple-fabric")},
	// 				NfcSKU: to.Ptr(armmanagednetworkfabric.NfcSKUStandard),
	// 				ProvisioningState: to.Ptr(armmanagednetworkfabric.ProvisioningStateSucceeded),
	// 				TenantInternetGatewayIDs: []*string{
	// 					to.Ptr("/subscriptions/xxxxxx/resourceGroups/resourcegroupname/providers/Microsoft.ManagedNetworkFabric/internetGateways/example-internetGateway")},
	// 					WorkloadManagementNetwork: to.Ptr(true),
	// 					WorkloadServices: &armmanagednetworkfabric.ControllerServices{
	// 						IPv4AddressSpaces: []*string{
	// 							to.Ptr("172.253.28.0/22")},
	// 							IPv6AddressSpaces: []*string{
	// 							},
	// 						},
	// 					},
	// 				}
}
