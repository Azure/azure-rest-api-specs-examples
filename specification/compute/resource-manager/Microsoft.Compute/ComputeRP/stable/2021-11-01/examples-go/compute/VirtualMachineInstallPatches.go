package armcompute_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/VirtualMachineInstallPatches.json
func ExampleVirtualMachinesClient_BeginInstallPatches() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armcompute.NewVirtualMachinesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginInstallPatches(ctx,
		"<resource-group-name>",
		"<vm-name>",
		armcompute.VirtualMachineInstallPatchesParameters{
			MaximumDuration: to.Ptr("<maximum-duration>"),
			RebootSetting:   to.Ptr(armcompute.VMGuestPatchRebootSettingIfRequired),
			WindowsParameters: &armcompute.WindowsParameters{
				ClassificationsToInclude: []*armcompute.VMGuestPatchClassificationWindows{
					to.Ptr(armcompute.VMGuestPatchClassificationWindowsCritical),
					to.Ptr(armcompute.VMGuestPatchClassificationWindowsSecurity)},
				MaxPatchPublishDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-19T02:36:43.0539904+00:00"); return t }()),
			},
		},
		&armcompute.VirtualMachinesClientBeginInstallPatchesOptions{ResumeToken: ""})
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
