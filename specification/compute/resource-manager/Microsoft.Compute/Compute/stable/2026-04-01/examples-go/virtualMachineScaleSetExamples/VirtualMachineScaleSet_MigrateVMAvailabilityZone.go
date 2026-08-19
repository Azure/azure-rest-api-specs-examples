package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v8"
)

// Generated from example definition: 2026-04-01/virtualMachineScaleSetExamples/VirtualMachineScaleSet_MigrateVMAvailabilityZone.json
func ExampleVirtualMachineScaleSetsClient_BeginMigrateVMAvailabilityZone() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualMachineScaleSetsClient().BeginMigrateVMAvailabilityZone(ctx, "myResourceGroup", "{vmss-name}", armcompute.MigrateVMAvailabilityZoneInput{
		InstanceIDs: []*string{
			to.Ptr("0"),
			to.Ptr("1"),
			to.Ptr("2"),
		},
		TargetZone: to.Ptr("2"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
}
