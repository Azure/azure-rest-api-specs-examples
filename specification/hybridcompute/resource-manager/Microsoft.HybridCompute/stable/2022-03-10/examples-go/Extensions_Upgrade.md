Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fhybridcompute%2Farmhybridcompute%2Fv1.0.0/sdk/resourcemanager/hybridcompute/armhybridcompute/README.md) on how to add the SDK to your project and authenticate.

```go
package armhybridcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridcompute/armhybridcompute"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/hybridcompute/resource-manager/Microsoft.HybridCompute/stable/2022-03-10/examples/Extensions_Upgrade.json
func ExampleManagementClient_BeginUpgradeExtensions() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armhybridcompute.NewManagementClient("{subscriptionId}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginUpgradeExtensions(ctx,
		"myResourceGroup",
		"myMachine",
		armhybridcompute.MachineExtensionUpgrade{
			ExtensionTargets: map[string]*armhybridcompute.ExtensionTargetProperties{
				"Microsoft.Azure.Monitoring": {
					TargetVersion: to.Ptr("2.0"),
				},
				"Microsoft.Compute.CustomScriptExtension": {
					TargetVersion: to.Ptr("1.10"),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
```
