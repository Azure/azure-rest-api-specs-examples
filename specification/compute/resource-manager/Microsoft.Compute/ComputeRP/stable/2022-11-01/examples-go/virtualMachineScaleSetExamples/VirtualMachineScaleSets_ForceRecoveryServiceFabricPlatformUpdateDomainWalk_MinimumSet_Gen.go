package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/19f98c9f526f8db961f172276dd6d6882a86ed86/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-11-01/examples/virtualMachineScaleSetExamples/VirtualMachineScaleSets_ForceRecoveryServiceFabricPlatformUpdateDomainWalk_MinimumSet_Gen.json
func ExampleVirtualMachineScaleSetsClient_ForceRecoveryServiceFabricPlatformUpdateDomainWalk_virtualMachineScaleSetsForceRecoveryServiceFabricPlatformUpdateDomainWalkMinimumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVirtualMachineScaleSetsClient().ForceRecoveryServiceFabricPlatformUpdateDomainWalk(ctx, "rgcompute", "aaaaaaaaaaaa", 9, &armcompute.VirtualMachineScaleSetsClientForceRecoveryServiceFabricPlatformUpdateDomainWalkOptions{Zone: nil,
		PlacementGroupID: nil,
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.RecoveryWalkResponse = armcompute.RecoveryWalkResponse{
	// }
}
