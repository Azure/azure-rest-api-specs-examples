package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v8"
)

// Generated from example definition: 2026-03-01/virtualMachineScaleSetExamples/VirtualMachineScaleSet_Reimage_MaximumSet_Gen.json
func ExampleVirtualMachineScaleSetsClient_BeginReimage_virtualMachineScaleSetReimageMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualMachineScaleSetsClient().BeginReimage(ctx, "rgcompute", "aaaaaaaaaaaaaaaaaaaaaaaaaa", &armcompute.VirtualMachineScaleSetsClientBeginReimageOptions{
		VMScaleSetReimageInput: &armcompute.VirtualMachineScaleSetReimageParameters{
			InstanceIDs: []*string{
				to.Ptr("aaaaaaaaaa"),
			},
			ForceUpdateOSDiskForEphemeral: to.Ptr(true),
			TempDisk:                      to.Ptr(true),
		}})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcompute.VirtualMachineScaleSetsClientReimageResponse{
	// }
}
