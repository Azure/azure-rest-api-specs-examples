package armhybridcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridcompute/armhybridcompute/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/f41d0c9332078cb2ef07b749081d94915255ada5/specification/hybridcompute/resource-manager/Microsoft.HybridCompute/stable/2024-07-10/examples/extension/Extensions_Upgrade.json
func ExampleManagementClient_BeginUpgradeExtensions() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewManagementClient().BeginUpgradeExtensions(ctx, "myResourceGroup", "myMachine", armhybridcompute.MachineExtensionUpgrade{
		ExtensionTargets: map[string]*armhybridcompute.ExtensionTargetProperties{
			"Microsoft.Azure.Monitoring": {
				TargetVersion: to.Ptr("2.0"),
			},
			"Microsoft.Compute.CustomScriptExtension": {
				TargetVersion: to.Ptr("1.10"),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
