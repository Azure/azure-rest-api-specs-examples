package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v10"
)

// Generated from example definition: 2025-07-01/NetworkVirtualApplianceSkuGet.json
func ExampleVirtualApplianceSKUsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVirtualApplianceSKUsClient().Get(ctx, "ciscoSdwan", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armnetwork.VirtualApplianceSKUsClientGetResponse{
	// 	VirtualApplianceSKU: armnetwork.VirtualApplianceSKU{
	// 		Name: to.Ptr("ciscoSdwan"),
	// 		Type: to.Ptr("Microsoft.Network/networkVirtualApplianceSkus"),
	// 		Etag: to.Ptr("w/\\00000000-0000-0000-0000-000000000000\\"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/networkVirtualApplianceSkus/ciscoSdwan"),
	// 		Properties: &armnetwork.VirtualApplianceSKUPropertiesFormat{
	// 			AvailableScaleUnits: []*armnetwork.VirtualApplianceSKUInstances{
	// 				{
	// 					InstanceCount: to.Ptr[int32](2),
	// 					ScaleUnit: to.Ptr("1"),
	// 				},
	// 				{
	// 					InstanceCount: to.Ptr[int32](2),
	// 					ScaleUnit: to.Ptr("2"),
	// 				},
	// 			},
	// 			AvailableVersions: []*string{
	// 				to.Ptr("11.12"),
	// 			},
	// 			Vendor: to.Ptr("Cisco"),
	// 		},
	// 	},
	// }
}
