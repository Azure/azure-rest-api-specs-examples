package armconnectedvmware_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/connectedvmware/armconnectedvmware"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7473936304533e6716fc4563401bf265dda4cb64/specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/preview/2022-01-10-preview/examples/VirtualMachineInstallPatches.json
func ExampleVirtualMachinesClient_BeginInstallPatches() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armconnectedvmware.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualMachinesClient().BeginInstallPatches(ctx, "myResourceGroupName", "myMachineName", armconnectedvmware.VirtualMachineInstallPatchesParameters{
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
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.VirtualMachineInstallPatchesResult = armconnectedvmware.VirtualMachineInstallPatchesResult{
	// 	ExcludedPatchCount: to.Ptr[int32](0),
	// 	FailedPatchCount: to.Ptr[int32](0),
	// 	InstallationActivityID: to.Ptr("68f8b292-dfc2-4646-9781-33cc88631968"),
	// 	InstalledPatchCount: to.Ptr[int32](3),
	// 	LastModifiedDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-15T02:16:06.9740000Z"); return t}()),
	// 	MaintenanceWindowExceeded: to.Ptr(false),
	// 	NotSelectedPatchCount: to.Ptr[int32](0),
	// 	OSType: to.Ptr(armconnectedvmware.OsTypeUMWindows),
	// 	PendingPatchCount: to.Ptr[int32](2),
	// 	RebootStatus: to.Ptr(armconnectedvmware.VMGuestPatchRebootStatusCompleted),
	// 	StartDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-01-15T02:15:06.9740000Z"); return t}()),
	// 	StartedBy: to.Ptr(armconnectedvmware.PatchOperationStartedByUser),
	// 	Status: to.Ptr(armconnectedvmware.PatchOperationStatusSucceeded),
	// }
}
