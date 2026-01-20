package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ccea12be6cd5e64743499d46f7a9c8b60df52db1/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2025-04-01/examples/virtualMachineScaleSetExamples/VirtualMachineScaleSet_ConvertToSinglePlacementGroup_MaximumSet_Gen.json
func ExampleVirtualMachineScaleSetsClient_ConvertToSinglePlacementGroup_virtualMachineScaleSetConvertToSinglePlacementGroupMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewVirtualMachineScaleSetsClient().ConvertToSinglePlacementGroup(ctx, "rgcompute", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", armcompute.VMScaleSetConvertToSinglePlacementGroupInput{
		ActivePlacementGroupID: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaa"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
