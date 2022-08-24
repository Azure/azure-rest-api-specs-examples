package armconnectedvmware_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/connectedvmware/armconnectedvmware"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/preview/2022-01-10-preview/examples/VirtualMachineInstallPatches.json
func ExampleVirtualMachinesClient_BeginInstallPatches() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armconnectedvmware.NewVirtualMachinesClient("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginInstallPatches(ctx, "myResourceGroupName", "myMachineName", armconnectedvmware.VirtualMachineInstallPatchesParameters{
		MaximumDuration: to.Ptr("PT3H"),
		RebootSetting:   to.Ptr(armconnectedvmware.VMGuestPatchRebootSettingIfRequired),
		WindowsParameters: &armconnectedvmware.WindowsParameters{
			ClassificationsToInclude: []*armconnectedvmware.VMGuestPatchClassificationWindows{
				to.Ptr(armconnectedvmware.VMGuestPatchClassificationWindowsCritical),
				to.Ptr(armconnectedvmware.VMGuestPatchClassificationWindowsSecurity)},
			MaxPatchPublishDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-01-15T02:36:43.0539904+00:00"); return t }()),
		},
	}, nil)
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
