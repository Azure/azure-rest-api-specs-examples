package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/compute/resource-manager/Microsoft.Compute/stable/2022-03-01/ComputeRP/examples/virtualMachineScaleSetExamples/VirtualMachineScaleSets_ForceRecoveryServiceFabricPlatformUpdateDomainWalk_MaximumSet_Gen.json
func ExampleVirtualMachineScaleSetsClient_ForceRecoveryServiceFabricPlatformUpdateDomainWalk() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcompute.NewVirtualMachineScaleSetsClient("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.ForceRecoveryServiceFabricPlatformUpdateDomainWalk(ctx,
		"rgcompute",
		"aaaaaaaaaaaaaaaa",
		30,
		&armcompute.VirtualMachineScaleSetsClientForceRecoveryServiceFabricPlatformUpdateDomainWalkOptions{Zone: nil,
			PlacementGroupID: nil,
		})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
