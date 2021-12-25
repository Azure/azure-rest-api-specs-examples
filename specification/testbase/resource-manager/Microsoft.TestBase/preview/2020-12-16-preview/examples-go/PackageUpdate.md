Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Ftestbase%2Farmtestbase%2Fv0.1.0/sdk/resourcemanager/testbase/armtestbase/README.md) on how to add the SDK to your project and authenticate.

```go
package armtestbase_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/testbase/armtestbase"
)

// x-ms-original-file: specification/testbase/resource-manager/Microsoft.TestBase/preview/2020-12-16-preview/examples/PackageUpdate.json
func ExamplePackagesClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armtestbase.NewPackagesClient("<subscription-id>", cred, nil)
	poller, err := client.BeginUpdate(ctx,
		"<resource-group-name>",
		"<test-base-account-name>",
		"<package-name>",
		armtestbase.PackageUpdateParameters{
			Properties: &armtestbase.PackageUpdateParameterProperties{
				BlobPath:      to.StringPtr("<blob-path>"),
				FlightingRing: to.StringPtr("<flighting-ring>"),
				IsEnabled:     to.BoolPtr(false),
				TargetOSList: []*armtestbase.TargetOSInfo{
					{
						OSUpdateType: to.StringPtr("<osupdate-type>"),
						TargetOSs: []*string{
							to.StringPtr("Windows 10 2004"),
							to.StringPtr("Windows 10 1903")},
					}},
				Tests: []*armtestbase.Test{
					{
						IsActive: to.BoolPtr(true),
						TestType: armtestbase.TestTypeOutOfBoxTest.ToPtr(),
						Commands: []*armtestbase.Command{
							{
								Name:              to.StringPtr("<name>"),
								Action:            armtestbase.ActionInstall.ToPtr(),
								AlwaysRun:         to.BoolPtr(true),
								ApplyUpdateBefore: to.BoolPtr(false),
								Content:           to.StringPtr("<content>"),
								ContentType:       armtestbase.ContentTypePath.ToPtr(),
								MaxRunTime:        to.Int32Ptr(1800),
								RestartAfter:      to.BoolPtr(true),
								RunAsInteractive:  to.BoolPtr(true),
								RunElevated:       to.BoolPtr(true),
							},
							{
								Name:              to.StringPtr("<name>"),
								Action:            armtestbase.ActionLaunch.ToPtr(),
								AlwaysRun:         to.BoolPtr(false),
								ApplyUpdateBefore: to.BoolPtr(true),
								Content:           to.StringPtr("<content>"),
								ContentType:       armtestbase.ContentTypePath.ToPtr(),
								MaxRunTime:        to.Int32Ptr(1800),
								RestartAfter:      to.BoolPtr(false),
								RunAsInteractive:  to.BoolPtr(true),
								RunElevated:       to.BoolPtr(true),
							},
							{
								Name:              to.StringPtr("<name>"),
								Action:            armtestbase.ActionClose.ToPtr(),
								AlwaysRun:         to.BoolPtr(false),
								ApplyUpdateBefore: to.BoolPtr(false),
								Content:           to.StringPtr("<content>"),
								ContentType:       armtestbase.ContentTypePath.ToPtr(),
								MaxRunTime:        to.Int32Ptr(1800),
								RestartAfter:      to.BoolPtr(false),
								RunAsInteractive:  to.BoolPtr(true),
								RunElevated:       to.BoolPtr(true),
							},
							{
								Name:              to.StringPtr("<name>"),
								Action:            armtestbase.ActionUninstall.ToPtr(),
								AlwaysRun:         to.BoolPtr(true),
								ApplyUpdateBefore: to.BoolPtr(false),
								Content:           to.StringPtr("<content>"),
								ContentType:       armtestbase.ContentTypePath.ToPtr(),
								MaxRunTime:        to.Int32Ptr(1800),
								RestartAfter:      to.BoolPtr(false),
								RunAsInteractive:  to.BoolPtr(true),
								RunElevated:       to.BoolPtr(true),
							}},
					}},
			},
			Tags: map[string]*string{},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("PackageResource.ID: %s\n", *res.ID)
}
```
