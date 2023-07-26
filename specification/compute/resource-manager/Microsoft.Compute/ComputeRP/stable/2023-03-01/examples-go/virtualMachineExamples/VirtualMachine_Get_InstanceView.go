package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c53808ba54beef57059371708f1fa6949a11a280/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-03-01/examples/virtualMachineExamples/VirtualMachine_Get_InstanceView.json
func ExampleVirtualMachinesClient_InstanceView_getVirtualMachineInstanceView() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
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
	// res.VirtualMachineInstanceView = armcompute.VirtualMachineInstanceView{
	// 	BootDiagnostics: &armcompute.BootDiagnosticsInstanceView{
	// 		ConsoleScreenshotBlobURI: to.Ptr("https://{myStorageAccount}.blob.core.windows.net/bootdiagnostics-myOsDisk/myOsDisk.screenshot.bmp"),
	// 		SerialConsoleLogBlobURI: to.Ptr("https://{myStorageAccount}.blob.core.windows.net/bootdiagnostics-myOsDisk/myOsDisk.serialconsole.log"),
	// 	},
	// 	ComputerName: to.Ptr("myVM"),
	// 	Disks: []*armcompute.DiskInstanceView{
	// 		{
	// 			Name: to.Ptr("myOsDisk"),
	// 			Statuses: []*armcompute.InstanceViewStatus{
	// 				{
	// 					Code: to.Ptr("ProvisioningState/succeeded"),
	// 					DisplayStatus: to.Ptr("Provisioning succeeded"),
	// 					Level: to.Ptr(armcompute.StatusLevelTypesInfo),
	// 					Time: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-10-14T21:29:47.477089+00:00"); return t}()),
	// 			}},
	// 		},
	// 		{
	// 			Name: to.Ptr("myDataDisk0"),
	// 			Statuses: []*armcompute.InstanceViewStatus{
	// 				{
	// 					Code: to.Ptr("ProvisioningState/succeeded"),
	// 					DisplayStatus: to.Ptr("Provisioning succeeded"),
	// 					Level: to.Ptr(armcompute.StatusLevelTypesInfo),
	// 					Time: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-10-14T21:29:47.461517+00:00"); return t}()),
	// 			}},
	// 	}},
	// 	HyperVGeneration: to.Ptr(armcompute.HyperVGenerationTypeV1),
	// 	OSName: to.Ptr("Windows Server 2016 Datacenter"),
	// 	OSVersion: to.Ptr("Microsoft Windows NT 10.0.14393.0"),
	// 	PatchStatus: &armcompute.VirtualMachinePatchStatus{
	// 		AvailablePatchSummary: &armcompute.AvailablePatchSummary{
	// 			AssessmentActivityID: to.Ptr("68f8b292-dfc2-4646-9781-33cc88631968"),
	// 			CriticalAndSecurityPatchCount: to.Ptr[int32](1),
	// 			LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-04-24T21:02:04.2556154Z"); return t}()),
	// 			OtherPatchCount: to.Ptr[int32](2),
	// 			RebootPending: to.Ptr(true),
	// 			StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-04-24T21:02:04.2556154Z"); return t}()),
	// 			Status: to.Ptr(armcompute.PatchOperationStatusSucceeded),
	// 		},
	// 		ConfigurationStatuses: []*armcompute.InstanceViewStatus{
	// 			{
	// 				Code: to.Ptr("PatchModeConfigurationState/Ready"),
	// 				DisplayStatus: to.Ptr("Status_PatchModeConfigurationState_Ready"),
	// 				Level: to.Ptr(armcompute.StatusLevelTypesInfo),
	// 				Time: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-04-24T21:02:04.2556154Z"); return t}()),
	// 			},
	// 			{
	// 				Code: to.Ptr("AssessmentModeConfigurationState/Pending"),
	// 				DisplayStatus: to.Ptr("Status_AssessmentModeConfigurationState_Pending"),
	// 				Level: to.Ptr(armcompute.StatusLevelTypesInfo),
	// 				Time: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-04-24T21:02:04.2556154Z"); return t}()),
	// 		}},
	// 		LastPatchInstallationSummary: &armcompute.LastPatchInstallationSummary{
	// 			ExcludedPatchCount: to.Ptr[int32](1),
	// 			FailedPatchCount: to.Ptr[int32](1),
	// 			InstallationActivityID: to.Ptr("68f8b292-dfc2-4646-9981-33cc88631968"),
	// 			InstalledPatchCount: to.Ptr[int32](1),
	// 			LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-04-24T21:02:04.2556154Z"); return t}()),
	// 			MaintenanceWindowExceeded: to.Ptr(false),
	// 			NotSelectedPatchCount: to.Ptr[int32](1),
	// 			PendingPatchCount: to.Ptr[int32](1),
	// 			StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-04-24T21:02:04.2556154Z"); return t}()),
	// 			Status: to.Ptr(armcompute.PatchOperationStatusSucceeded),
	// 		},
	// 	},
	// 	PlatformFaultDomain: to.Ptr[int32](1),
	// 	PlatformUpdateDomain: to.Ptr[int32](1),
	// 	Statuses: []*armcompute.InstanceViewStatus{
	// 		{
	// 			Code: to.Ptr("ProvisioningState/succeeded"),
	// 			DisplayStatus: to.Ptr("Provisioning succeeded"),
	// 			Level: to.Ptr(armcompute.StatusLevelTypesInfo),
	// 			Time: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-10-14T21:30:12.8051917+00:00"); return t}()),
	// 		},
	// 		{
	// 			Code: to.Ptr("PowerState/running"),
	// 			DisplayStatus: to.Ptr("VM running"),
	// 			Level: to.Ptr(armcompute.StatusLevelTypesInfo),
	// 	}},
	// 	VMAgent: &armcompute.VirtualMachineAgentInstanceView{
	// 		ExtensionHandlers: []*armcompute.VirtualMachineExtensionHandlerInstanceView{
	// 			{
	// 				Type: to.Ptr("Microsoft.Azure.Security.IaaSAntimalware"),
	// 				Status: &armcompute.InstanceViewStatus{
	// 					Code: to.Ptr("ProvisioningState/succeeded"),
	// 					DisplayStatus: to.Ptr("Ready"),
	// 					Level: to.Ptr(armcompute.StatusLevelTypesInfo),
	// 				},
	// 				TypeHandlerVersion: to.Ptr("1.5.5.9"),
	// 		}},
	// 		Statuses: []*armcompute.InstanceViewStatus{
	// 			{
	// 				Code: to.Ptr("ProvisioningState/succeeded"),
	// 				DisplayStatus: to.Ptr("Ready"),
	// 				Level: to.Ptr(armcompute.StatusLevelTypesInfo),
	// 				Message: to.Ptr("GuestAgent is running and accepting new configurations."),
	// 				Time: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-10-14T23:11:22+00:00"); return t}()),
	// 		}},
	// 		VMAgentVersion: to.Ptr("2.7.41491.949"),
	// 	},
	// 	Extensions: []*armcompute.VirtualMachineExtensionInstanceView{
	// 		{
	// 			Name: to.Ptr("IaaSAntiMalware-ext0"),
	// 			Type: to.Ptr("Microsoft.Azure.Security.IaaSAntimalware"),
	// 			Statuses: []*armcompute.InstanceViewStatus{
	// 				{
	// 					Code: to.Ptr("ProvisioningState/succeeded"),
	// 					DisplayStatus: to.Ptr("Provisioning succeeded"),
	// 					Level: to.Ptr(armcompute.StatusLevelTypesInfo),
	// 					Message: to.Ptr("Microsoft Antimalware enabled"),
	// 			}},
	// 			TypeHandlerVersion: to.Ptr("1.5.5.9"),
	// 	}},
	// }
}
