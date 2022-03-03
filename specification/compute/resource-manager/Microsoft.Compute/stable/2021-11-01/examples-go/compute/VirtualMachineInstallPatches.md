Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fcompute%2Farmcompute%2Fv0.4.0/sdk/resourcemanager/compute/armcompute/README.md) on how to add the SDK to your project and authenticate.

```go
package armcompute_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute"
)

// x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/VirtualMachineInstallPatches.json
func ExampleVirtualMachinesClient_BeginInstallPatches() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcompute.NewVirtualMachinesClient("<subscription-id>", cred, nil)
	poller, err := client.BeginInstallPatches(ctx,
		"<resource-group-name>",
		"<vm-name>",
		armcompute.VirtualMachineInstallPatchesParameters{
			MaximumDuration: to.StringPtr("<maximum-duration>"),
			RebootSetting:   armcompute.VMGuestPatchRebootSetting("IfRequired").ToPtr(),
			WindowsParameters: &armcompute.WindowsParameters{
				ClassificationsToInclude: []*armcompute.VMGuestPatchClassificationWindows{
					armcompute.VMGuestPatchClassificationWindows("Critical").ToPtr(),
					armcompute.VMGuestPatchClassificationWindows("Security").ToPtr()},
				MaxPatchPublishDate: to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-19T02:36:43.0539904+00:00"); return t }()),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.VirtualMachinesClientInstallPatchesResult)
}
```
