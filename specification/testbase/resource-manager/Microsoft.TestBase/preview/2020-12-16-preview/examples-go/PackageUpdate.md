Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Ftestbase%2Farmtestbase%2Fv0.4.0/sdk/resourcemanager/testbase/armtestbase/README.md) on how to add the SDK to your project and authenticate.

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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/testbase/resource-manager/Microsoft.TestBase/preview/2020-12-16-preview/examples/PackageUpdate.json
func ExamplePackagesClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armtestbase.NewPackagesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginUpdate(ctx,
		"<resource-group-name>",
		"<test-base-account-name>",
		"<package-name>",
		armtestbase.PackageUpdateParameters{
			Properties: &armtestbase.PackageUpdateParameterProperties{
				BlobPath:      to.Ptr("<blob-path>"),
				FlightingRing: to.Ptr("<flighting-ring>"),
				IsEnabled:     to.Ptr(false),
				TargetOSList: []*armtestbase.TargetOSInfo{
					{
						OSUpdateType: to.Ptr("<osupdate-type>"),
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
								Name:              to.Ptr("<name>"),
								Action:            to.Ptr(armtestbase.ActionInstall),
								AlwaysRun:         to.Ptr(true),
								ApplyUpdateBefore: to.Ptr(false),
								Content:           to.Ptr("<content>"),
								ContentType:       to.Ptr(armtestbase.ContentTypePath),
								MaxRunTime:        to.Ptr[int32](1800),
								RestartAfter:      to.Ptr(true),
								RunAsInteractive:  to.Ptr(true),
								RunElevated:       to.Ptr(true),
							},
							{
								Name:              to.Ptr("<name>"),
								Action:            to.Ptr(armtestbase.ActionLaunch),
								AlwaysRun:         to.Ptr(false),
								ApplyUpdateBefore: to.Ptr(true),
								Content:           to.Ptr("<content>"),
								ContentType:       to.Ptr(armtestbase.ContentTypePath),
								MaxRunTime:        to.Ptr[int32](1800),
								RestartAfter:      to.Ptr(false),
								RunAsInteractive:  to.Ptr(true),
								RunElevated:       to.Ptr(true),
							},
							{
								Name:              to.Ptr("<name>"),
								Action:            to.Ptr(armtestbase.ActionClose),
								AlwaysRun:         to.Ptr(false),
								ApplyUpdateBefore: to.Ptr(false),
								Content:           to.Ptr("<content>"),
								ContentType:       to.Ptr(armtestbase.ContentTypePath),
								MaxRunTime:        to.Ptr[int32](1800),
								RestartAfter:      to.Ptr(false),
								RunAsInteractive:  to.Ptr(true),
								RunElevated:       to.Ptr(true),
							},
							{
								Name:              to.Ptr("<name>"),
								Action:            to.Ptr(armtestbase.ActionUninstall),
								AlwaysRun:         to.Ptr(true),
								ApplyUpdateBefore: to.Ptr(false),
								Content:           to.Ptr("<content>"),
								ContentType:       to.Ptr(armtestbase.ContentTypePath),
								MaxRunTime:        to.Ptr[int32](1800),
								RestartAfter:      to.Ptr(false),
								RunAsInteractive:  to.Ptr(true),
								RunElevated:       to.Ptr(true),
							}},
					}},
			},
			Tags: map[string]*string{},
		},
		&armtestbase.PackagesClientBeginUpdateOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
