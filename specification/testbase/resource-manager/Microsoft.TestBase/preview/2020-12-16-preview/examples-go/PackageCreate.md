Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Ftestbase%2Farmtestbase%2Fv0.5.0/sdk/resourcemanager/testbase/armtestbase/README.md) on how to add the SDK to your project and authenticate.

```go
package armtestbase_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/testbase/armtestbase"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/testbase/resource-manager/Microsoft.TestBase/preview/2020-12-16-preview/examples/PackageCreate.json
func ExamplePackagesClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armtestbase.NewPackagesClient("476f61a4-952c-422a-b4db-568a828f35df", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreate(ctx,
		"contoso-rg1",
		"contoso-testBaseAccount1",
		"contoso-package2",
		armtestbase.PackageResource{
			Location: to.Ptr("westus"),
			Tags:     map[string]*string{},
			Properties: &armtestbase.PackageProperties{
				ApplicationName: to.Ptr("contoso-package2"),
				BlobPath:        to.Ptr("storageAccountPath/package.zip"),
				FlightingRing:   to.Ptr("Insider Beta Channel"),
				TargetOSList: []*armtestbase.TargetOSInfo{
					{
						OSUpdateType: to.Ptr("Security updates"),
						TargetOSs: []*string{
							to.Ptr("Windows 10 2004"),
							to.Ptr("Windows 10 1903")},
					}},
				Tests: []*armtestbase.Test{
					{
						IsActive: to.Ptr(true),
						TestType: to.Ptr(armtestbase.TestTypeOutOfBoxTest),
						Commands: []*armtestbase.Command{
							{
								Name:              to.Ptr("Install"),
								Action:            to.Ptr(armtestbase.ActionInstall),
								AlwaysRun:         to.Ptr(true),
								ApplyUpdateBefore: to.Ptr(false),
								Content:           to.Ptr("app/scripts/install/job.ps1"),
								ContentType:       to.Ptr(armtestbase.ContentTypePath),
								MaxRunTime:        to.Ptr[int32](1800),
								RestartAfter:      to.Ptr(true),
								RunAsInteractive:  to.Ptr(true),
								RunElevated:       to.Ptr(true),
							},
							{
								Name:              to.Ptr("Launch"),
								Action:            to.Ptr(armtestbase.ActionLaunch),
								AlwaysRun:         to.Ptr(false),
								ApplyUpdateBefore: to.Ptr(true),
								Content:           to.Ptr("app/scripts/launch/job.ps1"),
								ContentType:       to.Ptr(armtestbase.ContentTypePath),
								MaxRunTime:        to.Ptr[int32](1800),
								RestartAfter:      to.Ptr(false),
								RunAsInteractive:  to.Ptr(true),
								RunElevated:       to.Ptr(true),
							},
							{
								Name:              to.Ptr("Close"),
								Action:            to.Ptr(armtestbase.ActionClose),
								AlwaysRun:         to.Ptr(false),
								ApplyUpdateBefore: to.Ptr(false),
								Content:           to.Ptr("app/scripts/close/job.ps1"),
								ContentType:       to.Ptr(armtestbase.ContentTypePath),
								MaxRunTime:        to.Ptr[int32](1800),
								RestartAfter:      to.Ptr(false),
								RunAsInteractive:  to.Ptr(true),
								RunElevated:       to.Ptr(true),
							},
							{
								Name:              to.Ptr("Uninstall"),
								Action:            to.Ptr(armtestbase.ActionUninstall),
								AlwaysRun:         to.Ptr(true),
								ApplyUpdateBefore: to.Ptr(false),
								Content:           to.Ptr("app/scripts/uninstall/job.ps1"),
								ContentType:       to.Ptr(armtestbase.ContentTypePath),
								MaxRunTime:        to.Ptr[int32](1800),
								RestartAfter:      to.Ptr(false),
								RunAsInteractive:  to.Ptr(true),
								RunElevated:       to.Ptr(true),
							}},
					}},
				Version: to.Ptr("1.0.0"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
```
