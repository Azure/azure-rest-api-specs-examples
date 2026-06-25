package armcompute_test

import (
	"context"
	"log"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v8"
)

// Generated from example definition: 2026-03-01/virtualMachineExamples/VirtualMachine_InstallPatches.json
func ExampleVirtualMachinesClient_BeginInstallPatches() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualMachinesClient().BeginInstallPatches(ctx, "myResourceGroupName", "myVMName", armcompute.VirtualMachineInstallPatchesParameters{
		MaximumDuration: to.Ptr("PT4H"),
		RebootSetting:   to.Ptr(armcompute.VMGuestPatchRebootSettingIfRequired),
		WindowsParameters: &armcompute.WindowsParameters{
			ClassificationsToInclude: []*armcompute.VMGuestPatchClassificationWindows{
				to.Ptr(armcompute.VMGuestPatchClassificationWindowsCritical),
				to.Ptr(armcompute.VMGuestPatchClassificationWindowsSecurity),
			},
			MaxPatchPublishDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-19T02:36:43.0539904+00:00"); return t }()),
			PatchNameMasksToInclude: []*string{
				to.Ptr("*SQL*"),
			},
			PatchNameMasksToExclude: []*string{
				to.Ptr("*Windows*"),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcompute.VirtualMachinesClientInstallPatchesResponse{
	// 	VirtualMachineInstallPatchesResult: armcompute.VirtualMachineInstallPatchesResult{
	// 		Status: to.Ptr(armcompute.PatchOperationStatusSucceeded),
	// 		InstallationActivityID: to.Ptr("68f8b292-dfc2-4646-9781-33cc88631968"),
	// 		RebootStatus: to.Ptr(armcompute.VMGuestPatchRebootStatusCompleted),
	// 		MaintenanceWindowExceeded: to.Ptr(false),
	// 		ExcludedPatchCount: to.Ptr[int32](0),
	// 		NotSelectedPatchCount: to.Ptr[int32](0),
	// 		PendingPatchCount: to.Ptr[int32](2),
	// 		InstalledPatchCount: to.Ptr[int32](3),
	// 		FailedPatchCount: to.Ptr[int32](0),
	// 		StartDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-04-24T21:02:04.2556154Z"); return t}()),
	// 		Patches: []*armcompute.PatchInstallationDetail{
	// 			{
	// 				PatchID: to.Ptr("35428702-5784-4ba4-a6e0-5222258b5411"),
	// 				Name: to.Ptr("Definition Update for Windows Defender Antivirus - KB2267602 (Definition 1.279.1373.0)"),
	// 				Version: to.Ptr(""),
	// 				KbID: to.Ptr("2267602"),
	// 				Classifications: []*string{
	// 					to.Ptr("Definition Updates"),
	// 				},
	// 				InstallationState: to.Ptr(armcompute.PatchInstallationStateInstalled),
	// 			},
	// 			{
	// 				PatchID: to.Ptr("39f9cdd1-795c-4d0e-8c0a-73ab3f31746d"),
	// 				Name: to.Ptr("Windows Malicious Software Removal Tool x64 - October 2018 (KB890830)"),
	// 				Version: to.Ptr(""),
	// 				KbID: to.Ptr("890830"),
	// 				Classifications: []*string{
	// 					to.Ptr("Update Rollups"),
	// 				},
	// 				InstallationState: to.Ptr(armcompute.PatchInstallationStatePending),
	// 			},
	// 		},
	// 	},
	// }
}
