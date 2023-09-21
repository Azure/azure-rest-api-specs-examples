package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c78b5d8bd3aff2d82a5f034d9164b1a9ac030e09/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-07-01/examples/virtualMachineScaleSetExamples/VirtualMachineScaleSet_ListSkus_MaximumSet_Gen.json
func ExampleVirtualMachineScaleSetsClient_NewListSKUsPager_virtualMachineScaleSetListSkusMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
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
		// page.VirtualMachineScaleSetListSKUsResult = armcompute.VirtualMachineScaleSetListSKUsResult{
		// 	Value: []*armcompute.VirtualMachineScaleSetSKU{
		// 		{
		// 			Capacity: &armcompute.VirtualMachineScaleSetSKUCapacity{
		// 				DefaultCapacity: to.Ptr[int64](20),
		// 				Maximum: to.Ptr[int64](27),
		// 				Minimum: to.Ptr[int64](22),
		// 				ScaleType: to.Ptr(armcompute.VirtualMachineScaleSetSKUScaleTypeAutomatic),
		// 			},
		// 			ResourceType: to.Ptr("aaaaaaaaaaaaaaaaa"),
		// 			SKU: &armcompute.SKU{
		// 				Name: to.Ptr("DSv3-Type1"),
		// 				Capacity: to.Ptr[int64](7),
		// 				Tier: to.Ptr("aaa"),
		// 			},
		// 	}},
		// }
	}
}
