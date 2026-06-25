package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v8"
)

// Generated from example definition: 2026-03-01/availabilitySetExamples/AvailabilitySet_CancelMigrationToVirtualMachineScaleSet.json
func ExampleAvailabilitySetsClient_CancelMigrationToVirtualMachineScaleSet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewAvailabilitySetsClient().CancelMigrationToVirtualMachineScaleSet(ctx, "rgcompute", "myAvailabilitySet", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
