package armhybridcompute_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridcompute/armhybridcompute"
)

// x-ms-original-file: specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2021-06-10-preview/examples/Extensions_Upgrade.json
func ExampleManagementClient_BeginUpgradeExtensions() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armhybridcompute.NewManagementClient("<subscription-id>", cred, nil)
	poller, err := client.BeginUpgradeExtensions(ctx,
		"<resource-group-name>",
		"<machine-name>",
		armhybridcompute.MachineExtensionUpgrade{
			ExtensionTargets: map[string]*armhybridcompute.ExtensionTargetProperties{
				"Microsoft.Azure.Monitoring": {
					TargetVersion: to.StringPtr("<target-version>"),
				},
				"Microsoft.Compute.CustomScriptExtension": {
					TargetVersion: to.StringPtr("<target-version>"),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}
