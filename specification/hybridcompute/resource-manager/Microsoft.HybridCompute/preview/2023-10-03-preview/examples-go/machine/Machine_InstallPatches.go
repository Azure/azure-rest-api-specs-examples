package armhybridcompute_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridcompute/armhybridcompute/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/71a0c7adf2a6e169ab9a33c7cf36bb93db083e86/specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2023-10-03-preview/examples/machine/Machine_InstallPatches.json
func ExampleMachinesClient_BeginInstallPatches() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewMachinesClient().BeginInstallPatches(ctx, "myResourceGroupName", "myMachineName", armhybridcompute.MachineInstallPatchesParameters{
		MaximumDuration: to.Ptr("PT4H"),
		RebootSetting:   to.Ptr(armhybridcompute.VMGuestPatchRebootSettingIfRequired),
		WindowsParameters: &armhybridcompute.WindowsParameters{
			ClassificationsToInclude: []*armhybridcompute.VMGuestPatchClassificationWindows{
				to.Ptr(armhybridcompute.VMGuestPatchClassificationWindowsCritical),
				to.Ptr(armhybridcompute.VMGuestPatchClassificationWindowsSecurity)},
			MaxPatchPublishDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-19T02:36:43.053Z"); return t }()),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.MachineInstallPatchesResult = armhybridcompute.MachineInstallPatchesResult{
	// 	ExcludedPatchCount: to.Ptr[int32](0),
	// 	FailedPatchCount: to.Ptr[int32](0),
	// 	InstallationActivityID: to.Ptr("68f8b292-dfc2-4646-9781-33cc88631968"),
	// 	InstalledPatchCount: to.Ptr[int32](3),
	// 	LastModifiedDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-22T02:16:06.974Z"); return t}()),
	// 	MaintenanceWindowExceeded: to.Ptr(false),
	// 	NotSelectedPatchCount: to.Ptr[int32](0),
	// 	OSType: to.Ptr(armhybridcompute.OsTypeWindows),
	// 	PendingPatchCount: to.Ptr[int32](2),
	// 	RebootStatus: to.Ptr(armhybridcompute.VMGuestPatchRebootStatusCompleted),
	// 	StartDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-22T02:15:06.974Z"); return t}()),
	// 	StartedBy: to.Ptr(armhybridcompute.PatchOperationStartedByUser),
	// 	Status: to.Ptr(armhybridcompute.PatchOperationStatusSucceeded),
	// }
}
