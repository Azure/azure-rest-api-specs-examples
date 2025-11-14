package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7033e85e1f80ef5cd9ca664b538ed193a8fd815b/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2025-04-01/examples/availabilitySetExamples/AvailabilitySet_CancelMigrationToVirtualMachineScaleSet.json
func ExampleAvailabilitySetsClient_CancelMigrationToVirtualMachineScaleSet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewAvailabilitySetsClient().CancelMigrationToVirtualMachineScaleSet(ctx, "rgcompute", "myAvailabilitySet", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
