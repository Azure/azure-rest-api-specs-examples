package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v8"
)

// Generated from example definition: 2026-03-01/virtualMachineScaleSetExamples/VirtualMachineScaleSet_ListSkus_MaximumSet_Gen.json
func ExampleVirtualMachineScaleSetsClient_NewListSKUsPager_virtualMachineScaleSetListSkusMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewVirtualMachineScaleSetsClient().NewListSKUsPager("rgcompute", "aaaaaa", nil)
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
		// page = armcompute.VirtualMachineScaleSetsClientListSKUsResponse{
		// 	VirtualMachineScaleSetListSKUsResult: armcompute.VirtualMachineScaleSetListSKUsResult{
		// 		Value: []*armcompute.VirtualMachineScaleSetSKU{
		// 			{
		// 				ResourceType: to.Ptr("aaaaaaaaaaaaaaaaa"),
		// 				SKU: &armcompute.SKU{
		// 					Name: to.Ptr("DSv3-Type1"),
		// 					Tier: to.Ptr("aaa"),
		// 					Capacity: to.Ptr[int64](7),
		// 				},
		// 				Capacity: &armcompute.VirtualMachineScaleSetSKUCapacity{
		// 					Minimum: to.Ptr[int64](22),
		// 					Maximum: to.Ptr[int64](27),
		// 					DefaultCapacity: to.Ptr[int64](20),
		// 					ScaleType: to.Ptr(armcompute.VirtualMachineScaleSetSKUScaleTypeAutomatic),
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("a://example.com/aaaaaaaaaaaaaa"),
		// 	},
		// }
	}
}
