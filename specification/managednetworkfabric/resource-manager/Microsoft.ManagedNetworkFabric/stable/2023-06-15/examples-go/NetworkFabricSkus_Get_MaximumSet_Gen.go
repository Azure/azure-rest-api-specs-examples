package armmanagednetworkfabric_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managednetworkfabric/armmanagednetworkfabric"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/925ba149e17454ce91ecd3f9f4134effb2f97844/specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/NetworkFabricSkus_Get_MaximumSet_Gen.json
func ExampleNetworkFabricSKUsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagednetworkfabric.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewNetworkFabricSKUsClient().Get(ctx, "example-fabricsku", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.NetworkFabricSKU = armmanagednetworkfabric.NetworkFabricSKU{
	// 	Name: to.Ptr("example-fabricsku"),
	// 	Type: to.Ptr("microsoft.managednetworkfabric/networkFabricSkus"),
	// 	ID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/providers/Microsoft.ManagedNetworkFabric/networkFabricSkus/example-fabricsku"),
	// 	SystemData: &armmanagednetworkfabric.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-11T16:44:43.644Z"); return t}()),
	// 		CreatedBy: to.Ptr("email@address.com"),
	// 		CreatedByType: to.Ptr(armmanagednetworkfabric.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-11T16:44:43.644Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("User@email.com"),
	// 		LastModifiedByType: to.Ptr(armmanagednetworkfabric.CreatedByTypeUser),
	// 	},
	// 	Properties: &armmanagednetworkfabric.NetworkFabricSKUProperties{
	// 		Type: to.Ptr(armmanagednetworkfabric.FabricSKUTypeSingleRack),
	// 		MaxComputeRacks: to.Ptr[int32](4),
	// 		MaximumServerCount: to.Ptr[int32](9),
	// 		ProvisioningState: to.Ptr(armmanagednetworkfabric.ProvisioningStateSucceeded),
	// 		SupportedVersions: []*string{
	// 			to.Ptr("1.0.0")},
	// 			Details: to.Ptr("https://azure/fabricskuDetails"),
	// 		},
	// 	}
}
