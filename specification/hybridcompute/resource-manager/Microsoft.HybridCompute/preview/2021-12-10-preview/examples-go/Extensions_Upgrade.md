```go
package armhybridcompute_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridcompute/armhybridcompute"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2021-12-10-preview/examples/Extensions_Upgrade.json
func ExampleManagementClient_BeginUpgradeExtensions() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armhybridcompute.NewManagementClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginUpgradeExtensions(ctx,
		"<resource-group-name>",
		"<machine-name>",
		armhybridcompute.MachineExtensionUpgrade{
			ExtensionTargets: map[string]*armhybridcompute.ExtensionTargetProperties{
				"Microsoft.Azure.Monitoring": {
					TargetVersion: to.Ptr("<target-version>"),
				},
				"Microsoft.Compute.CustomScriptExtension": {
					TargetVersion: to.Ptr("<target-version>"),
				},
			},
		},
		&armhybridcompute.ManagementClientBeginUpgradeExtensionsOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fhybridcompute%2Farmhybridcompute%2Fv0.4.0/sdk/resourcemanager/hybridcompute/armhybridcompute/README.md) on how to add the SDK to your project and authenticate.
