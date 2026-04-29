package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v8"
)

// Generated from example definition: 2025-11-01/virtualMachineExamples/VirtualMachine_MigrateToVirtualMachineScaleSet.json
func ExampleVirtualMachinesClient_BeginMigrateToVMScaleSet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualMachinesClient().BeginMigrateToVMScaleSet(ctx, "myResourceGroup", "myVMName", &armcompute.VirtualMachinesClientBeginMigrateToVMScaleSetOptions{
		Parameters: &armcompute.MigrateVMToVirtualMachineScaleSetInput{
			TargetFaultDomain: to.Ptr[int32](0),
			TargetVMSize:      to.Ptr("Standard_D1_v2"),
		}})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
