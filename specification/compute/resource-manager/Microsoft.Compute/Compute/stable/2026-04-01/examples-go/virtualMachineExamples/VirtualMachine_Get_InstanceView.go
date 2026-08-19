package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v8"
)

// Generated from example definition: 2026-04-01/virtualMachineExamples/VirtualMachine_Get_InstanceView.json
func ExampleVirtualMachinesClient_InstanceView_getVirtualMachineInstanceView() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVirtualMachinesClient().InstanceView(ctx, "myResourceGroup", "myVM", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcompute.VirtualMachinesClientInstanceViewResponse{
	// 	VirtualMachineInstanceView: armcompute.VirtualMachineInstanceView{
	// 		PlatformUpdateDomain: to.Ptr[int32](1),
	// 		PlatformFaultDomain: to.Ptr[int32](1),
	// 		ComputerName: to.Ptr("myVM"),
	// 		OSName: to.Ptr("Windows Server 2016 Datacenter"),
	// 		OSVersion: to.Ptr("Microsoft Windows NT 10.0.14393.0"),
	// 		VMAgent: &armcompute.VirtualMachineAgentInstanceView{
	// 			VMAgentVersion: to.Ptr("2.7.41491.949"),
	// 			Statuses: []*armcompute.InstanceViewStatus{
	// 				{
	// 					Code: to.Ptr("ProvisioningState/succeeded"),
	// 					Level: to.Ptr(armcompute.StatusLevelTypesInfo),
	// 					DisplayStatus: to.Ptr("Ready"),
	// 					Message: to.Ptr("GuestAgent is running and accepting new configurations."),
	// 					Time: to.Ptr(time.Date(2019, time.October, 14, 23, 11, 22, 0, time.UTC)),
	// 				},
	// 			},
	// 			ExtensionHandlers: []*armcompute.VirtualMachineExtensionHandlerInstanceView{
	// 				{
	// 					Type: to.Ptr("Microsoft.Azure.Security.IaaSAntimalware"),
	// 					TypeHandlerVersion: to.Ptr("1.5.5.9"),
	// 					Status: &armcompute.InstanceViewStatus{
	// 						Code: to.Ptr("ProvisioningState/succeeded"),
	// 						Level: to.Ptr(armcompute.StatusLevelTypesInfo),
	// 						DisplayStatus: to.Ptr("Ready"),
	// 					},
	// 				},
	// 			},
	// 		},
	// 		Disks: []*armcompute.DiskInstanceView{
	// 			{
	// 				Name: to.Ptr("myOsDisk"),
	// 				Statuses: []*armcompute.InstanceViewStatus{
	// 					{
	// 						Code: to.Ptr("ProvisioningState/succeeded"),
	// 						Level: to.Ptr(armcompute.StatusLevelTypesInfo),
	// 						DisplayStatus: to.Ptr("Provisioning succeeded"),
	// 						Time: to.Ptr(time.Date(2019, time.October, 14, 21, 29, 47, 477089000, time.UTC)),
	// 					},
	// 				},
	// 			},
	// 			{
	// 				Name: to.Ptr("myDataDisk0"),
	// 				Statuses: []*armcompute.InstanceViewStatus{
	// 					{
	// 						Code: to.Ptr("ProvisioningState/succeeded"),
	// 						Level: to.Ptr(armcompute.StatusLevelTypesInfo),
	// 						DisplayStatus: to.Ptr("Provisioning succeeded"),
	// 						Time: to.Ptr(time.Date(2019, time.October, 14, 21, 29, 47, 461517000, time.UTC)),
	// 					},
	// 				},
	// 			},
	// 		},
	// 		BootDiagnostics: &armcompute.BootDiagnosticsInstanceView{
	// 			ConsoleScreenshotBlobURI: to.Ptr("https://{myStorageAccount}.blob.core.windows.net/bootdiagnostics-myOsDisk/myOsDisk.screenshot.bmp"),
	// 			SerialConsoleLogBlobURI: to.Ptr("https://{myStorageAccount}.blob.core.windows.net/bootdiagnostics-myOsDisk/myOsDisk.serialconsole.log"),
	// 		},
	// 		Extensions: []*armcompute.VirtualMachineExtensionInstanceView{
	// 			{
	// 				Name: to.Ptr("IaaSAntiMalware-ext0"),
	// 				Type: to.Ptr("Microsoft.Azure.Security.IaaSAntimalware"),
	// 				TypeHandlerVersion: to.Ptr("1.5.5.9"),
	// 				Statuses: []*armcompute.InstanceViewStatus{
	// 					{
	// 						Code: to.Ptr("ProvisioningState/succeeded"),
	// 						Level: to.Ptr(armcompute.StatusLevelTypesInfo),
	// 						DisplayStatus: to.Ptr("Provisioning succeeded"),
	// 						Message: to.Ptr("Microsoft Antimalware enabled"),
	// 					},
	// 				},
	// 			},
	// 		},
	// 		HyperVGeneration: to.Ptr(armcompute.HyperVGenerationTypeV1),
	// 		PatchStatus: &armcompute.VirtualMachinePatchStatus{
	// 			AvailablePatchSummary: &armcompute.AvailablePatchSummary{
	// 				Status: to.Ptr(armcompute.PatchOperationStatusSucceeded),
	// 				AssessmentActivityID: to.Ptr("68f8b292-dfc2-4646-9781-33cc88631968"),
	// 				RebootPending: to.Ptr(true),
	// 				CriticalAndSecurityPatchCount: to.Ptr[int32](1),
	// 				OtherPatchCount: to.Ptr[int32](2),
	// 				StartTime: to.Ptr(time.Date(2020, time.April, 24, 21, 2, 4, 255615400, time.UTC)),
	// 				LastModifiedTime: to.Ptr(time.Date(2020, time.April, 24, 21, 2, 4, 255615400, time.UTC)),
	// 			},
	// 			LastPatchInstallationSummary: &armcompute.LastPatchInstallationSummary{
	// 				Status: to.Ptr(armcompute.PatchOperationStatusSucceeded),
	// 				InstallationActivityID: to.Ptr("68f8b292-dfc2-4646-9981-33cc88631968"),
	// 				MaintenanceWindowExceeded: to.Ptr(false),
	// 				NotSelectedPatchCount: to.Ptr[int32](1),
	// 				ExcludedPatchCount: to.Ptr[int32](1),
	// 				PendingPatchCount: to.Ptr[int32](1),
	// 				InstalledPatchCount: to.Ptr[int32](1),
	// 				FailedPatchCount: to.Ptr[int32](1),
	// 				StartTime: to.Ptr(time.Date(2020, time.April, 24, 21, 2, 4, 255615400, time.UTC)),
	// 				LastModifiedTime: to.Ptr(time.Date(2020, time.April, 24, 21, 2, 4, 255615400, time.UTC)),
	// 			},
	// 			ConfigurationStatuses: []*armcompute.InstanceViewStatus{
	// 				{
	// 					Code: to.Ptr("PatchModeConfigurationState/Ready"),
	// 					Level: to.Ptr(armcompute.StatusLevelTypesInfo),
	// 					DisplayStatus: to.Ptr("Status_PatchModeConfigurationState_Ready"),
	// 					Time: to.Ptr(time.Date(2020, time.April, 24, 21, 2, 4, 255615400, time.UTC)),
	// 				},
	// 				{
	// 					Code: to.Ptr("AssessmentModeConfigurationState/Pending"),
	// 					Level: to.Ptr(armcompute.StatusLevelTypesInfo),
	// 					DisplayStatus: to.Ptr("Status_AssessmentModeConfigurationState_Pending"),
	// 					Time: to.Ptr(time.Date(2020, time.April, 24, 21, 2, 4, 255615400, time.UTC)),
	// 				},
	// 			},
	// 		},
	// 		IsVMInStandbyPool: to.Ptr(false),
	// 		CapacityReservationType: to.Ptr(armcompute.CapacityReservationTypeOpen),
	// 		Statuses: []*armcompute.InstanceViewStatus{
	// 			{
	// 				Code: to.Ptr("ProvisioningState/succeeded"),
	// 				Level: to.Ptr(armcompute.StatusLevelTypesInfo),
	// 				DisplayStatus: to.Ptr("Provisioning succeeded"),
	// 				Time: to.Ptr(time.Date(2019, time.October, 14, 21, 30, 12, 805191700, time.UTC)),
	// 			},
	// 			{
	// 				Code: to.Ptr("PowerState/running"),
	// 				Level: to.Ptr(armcompute.StatusLevelTypesInfo),
	// 				DisplayStatus: to.Ptr("VM running"),
	// 			},
	// 		},
	// 	},
	// }
}
