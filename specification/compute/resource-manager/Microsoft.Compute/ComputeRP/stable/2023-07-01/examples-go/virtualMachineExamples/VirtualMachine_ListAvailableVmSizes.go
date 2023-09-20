package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c78b5d8bd3aff2d82a5f034d9164b1a9ac030e09/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-07-01/examples/virtualMachineExamples/VirtualMachine_ListAvailableVmSizes.json
func ExampleVirtualMachinesClient_NewListAvailableSizesPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewVirtualMachinesClient().NewListAvailableSizesPager("myResourceGroup", "myVmName", nil)
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
		// page.VirtualMachineSizeListResult = armcompute.VirtualMachineSizeListResult{
		// 	Value: []*armcompute.VirtualMachineSize{
		// 		{
		// 			Name: to.Ptr("Standard_A1_V2"),
		// 			MaxDataDiskCount: to.Ptr[int32](2),
		// 			MemoryInMB: to.Ptr[int32](2048),
		// 			NumberOfCores: to.Ptr[int32](1),
		// 			OSDiskSizeInMB: to.Ptr[int32](1047552),
		// 			ResourceDiskSizeInMB: to.Ptr[int32](10240),
		// 		},
		// 		{
		// 			Name: to.Ptr("Standard_A2_V2"),
		// 			MaxDataDiskCount: to.Ptr[int32](4),
		// 			MemoryInMB: to.Ptr[int32](4096),
		// 			NumberOfCores: to.Ptr[int32](2),
		// 			OSDiskSizeInMB: to.Ptr[int32](1047552),
		// 			ResourceDiskSizeInMB: to.Ptr[int32](20480),
		// 	}},
		// }
	}
}
