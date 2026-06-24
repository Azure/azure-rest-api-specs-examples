package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v8"
)

// Generated from example definition: 2026-03-01/virtualMachineScaleSetExamples/VirtualMachineScaleSetVM_SimulateEviction.json
func ExampleVirtualMachineScaleSetVMsClient_SimulateEviction() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewVirtualMachineScaleSetVMsClient().SimulateEviction(ctx, "ResourceGroup", "VmScaleSetName", "InstanceId", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
