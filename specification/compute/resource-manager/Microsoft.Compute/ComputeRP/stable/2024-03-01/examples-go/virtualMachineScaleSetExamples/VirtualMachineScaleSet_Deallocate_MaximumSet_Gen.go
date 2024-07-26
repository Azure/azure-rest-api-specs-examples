package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v6"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/81a4ee5a83ae38620c0e1404793caffe005d26e4/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-03-01/examples/virtualMachineScaleSetExamples/VirtualMachineScaleSet_Deallocate_MaximumSet_Gen.json
func ExampleVirtualMachineScaleSetsClient_BeginDeallocate_virtualMachineScaleSetDeallocateMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualMachineScaleSetsClient().BeginDeallocate(ctx, "rgcompute", "aaaaaaaaaaaaaaaaaaaaaaaaaaaa", &armcompute.VirtualMachineScaleSetsClientBeginDeallocateOptions{Hibernate: to.Ptr(true),
		VMInstanceIDs: &armcompute.VirtualMachineScaleSetVMInstanceIDs{
			InstanceIDs: []*string{
				to.Ptr("aaaaaaaaaaaaaaaaa")},
		},
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
