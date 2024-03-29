package armmanagednetworkfabric_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managednetworkfabric/armmanagednetworkfabric"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d03c1964cb76ffd6884d10a1871bbe779a2f68ef/specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/preview/2023-02-01-preview/examples/NetworkFabricSkus_Get_MaximumSet_Gen.json
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
	res, err := clientFactory.NewNetworkFabricSKUsClient().Get(ctx, "networkFabricSkuName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.NetworkFabricSKU = armmanagednetworkfabric.NetworkFabricSKU{
	// 	Name: to.Ptr("networkFabricSkuName"),
	// 	Type: to.Ptr("microsoft.managednetworkfabric/networkfabricsSku"),
	// 	ID: to.Ptr("/subscriptions/xxxxxx/resourceGroups/resourcegroupname/providers/Microsoft.ManagedNetworkFabric/networkFabrics/networkFabricSku/networkFabricSkuName"),
	// 	SystemData: &armmanagednetworkfabric.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-02-18T12:02:09.236Z"); return t}()),
	// 		CreatedBy: to.Ptr("User@email.com"),
	// 		CreatedByType: to.Ptr(armmanagednetworkfabric.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-02-18T12:02:09.236Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("Userid"),
	// 		LastModifiedByType: to.Ptr(armmanagednetworkfabric.CreatedByTypeApplication),
	// 	},
	// 	Properties: &armmanagednetworkfabric.NetworkFabricSKUProperties{
	// 		Type: to.Ptr("MultiRack"),
	// 		DetailsURI: to.Ptr("https://baseurl/images/skus/networkFabricSkuName"),
	// 		MaxComputeRacks: to.Ptr[int32](8),
	// 		MaxSupportedVer: to.Ptr("0.1.2"),
	// 		MinSupportedVer: to.Ptr("0.1.1"),
	// 		ProvisioningState: to.Ptr(armmanagednetworkfabric.ProvisioningStateSucceeded),
	// 	},
	// }
}
