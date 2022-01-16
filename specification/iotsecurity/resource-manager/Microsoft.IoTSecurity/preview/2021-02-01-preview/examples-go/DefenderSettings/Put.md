Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fiotsecurity%2Farmiotsecurity%2Fv0.2.0/sdk/resourcemanager/iotsecurity/armiotsecurity/README.md) on how to add the SDK to your project and authenticate.

```go
package armiotsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iotsecurity/armiotsecurity"
)

// x-ms-original-file: specification/iotsecurity/resource-manager/Microsoft.IoTSecurity/preview/2021-02-01-preview/examples/DefenderSettings/Put.json
func ExampleDefenderSettingsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armiotsecurity.NewDefenderSettingsClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		armiotsecurity.DefenderSettingsModel{
			Properties: &armiotsecurity.DefenderSettingsProperties{
				DeviceQuota: to.Int32Ptr(2000),
				MdeIntegration: &armiotsecurity.DefenderSettingsPropertiesMdeIntegration{
					Status: armiotsecurity.MdeIntegration("Enabled").ToPtr(),
				},
				OnboardingKind: armiotsecurity.OnboardingKind("Default").ToPtr(),
				SentinelWorkspaceResourceIDs: []*string{
					to.StringPtr("/subscriptions/c4930e90-cd72-4aa5-93e9-2d081d129569/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace1")},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.DefenderSettingsClientCreateOrUpdateResult)
}
```
