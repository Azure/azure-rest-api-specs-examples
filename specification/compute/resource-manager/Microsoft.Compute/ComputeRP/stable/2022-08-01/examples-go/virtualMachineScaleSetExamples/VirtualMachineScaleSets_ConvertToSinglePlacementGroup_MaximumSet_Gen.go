package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-08-01/examples/virtualMachineScaleSetExamples/VirtualMachineScaleSets_ConvertToSinglePlacementGroup_MaximumSet_Gen.json
func ExampleVirtualMachineScaleSetsClient_ConvertToSinglePlacementGroup_virtualMachineScaleSetsConvertToSinglePlacementGroupMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcompute.NewVirtualMachineScaleSetsClient("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.ConvertToSinglePlacementGroup(ctx, "rgcompute", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", armcompute.VMScaleSetConvertToSinglePlacementGroupInput{
		ActivePlacementGroupID: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaa"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
