Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsecurity%2Farmsecurity%2Fv0.3.1/sdk/resourcemanager/security/armsecurity/README.md) on how to add the SDK to your project and authenticate.

```go
package armsecurity_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2021-01-01/examples/Alerts/SimulateAlerts_example.json
func ExampleAlertsClient_BeginSimulate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsecurity.NewAlertsClient("<subscription-id>",
		"<asc-location>", cred, nil)
	poller, err := client.BeginSimulate(ctx,
		armsecurity.AlertSimulatorRequestBody{
			Properties: &armsecurity.AlertSimulatorBundlesRequestProperties{
				Kind: armsecurity.Kind("Bundles").ToPtr(),
				Bundles: []*armsecurity.BundleType{
					armsecurity.BundleType("AppServices").ToPtr(),
					armsecurity.BundleType("DNS").ToPtr(),
					armsecurity.BundleType("KeyVaults").ToPtr(),
					armsecurity.BundleType("KubernetesService").ToPtr(),
					armsecurity.BundleType("ResourceManager").ToPtr(),
					armsecurity.BundleType("SqlServers").ToPtr(),
					armsecurity.BundleType("StorageAccounts").ToPtr(),
					armsecurity.BundleType("VirtualMachines").ToPtr()},
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
```
