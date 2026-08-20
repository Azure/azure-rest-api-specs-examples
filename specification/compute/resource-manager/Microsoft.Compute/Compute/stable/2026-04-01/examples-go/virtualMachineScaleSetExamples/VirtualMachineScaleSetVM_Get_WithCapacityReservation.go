package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v8"
)

// Generated from example definition: 2026-04-01/virtualMachineScaleSetExamples/VirtualMachineScaleSetVM_Get_WithCapacityReservation.json
func ExampleVirtualMachineScaleSetVMsClient_Get_getVMScaleSetVMWithCapacityReservation() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVirtualMachineScaleSetVMsClient().Get(ctx, "myResourceGroup", "{vmss-name}", "0", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcompute.VirtualMachineScaleSetVMsClientGetResponse{
	// 	VirtualMachineScaleSetVM: armcompute.VirtualMachineScaleSetVM{
	// 		Name: to.Ptr("{vmss-vm-name}"),
	// 		ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachineScaleSets/{vmss-name}/virtualMachines/0"),
	// 		Type: to.Ptr("Microsoft.Compute/virtualMachines"),
	// 		Location: to.Ptr("westus"),
	// 		Tags: map[string]*string{
	// 			"myTag1": to.Ptr("tagValue1"),
	// 		},
	// 		Etag: to.Ptr("\"1\""),
	// 		Properties: &armcompute.VirtualMachineScaleSetVMProperties{
	// 			LatestModelApplied: to.Ptr(true),
	// 			ModelDefinitionApplied: to.Ptr("VirtualMachineScaleSet"),
	// 			VMID: to.Ptr("42af9fdf-b906-4ad7-9905-8316209ff619"),
	// 			HardwareProfile: &armcompute.HardwareProfile{
	// 			},
	// 			StorageProfile: &armcompute.StorageProfile{
	// 				ImageReference: &armcompute.ImageReference{
	// 					Publisher: to.Ptr("MicrosoftWindowsServer"),
	// 					Offer: to.Ptr("WindowsServer"),
	// 					SKU: to.Ptr("2012-R2-Datacenter"),
	// 					Version: to.Ptr("4.127.20180315"),
	// 					ExactVersion: to.Ptr("4.127.20180315"),
	// 				},
	// 				OSDisk: &armcompute.OSDisk{
	// 					OSType: to.Ptr(armcompute.OperatingSystemTypesWindows),
	// 					Name: to.Ptr("vmss3176_vmss3176_0_OsDisk_1_6d72b805e50e4de6830303c5055077fc"),
	// 					CreateOption: to.Ptr(armcompute.DiskCreateOptionTypesFromImage),
	// 					Caching: to.Ptr(armcompute.CachingTypesNone),
	// 					ManagedDisk: &armcompute.ManagedDiskParameters{
	// 						StorageAccountType: to.Ptr(armcompute.StorageAccountTypesStandardLRS),
	// 						ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/vmss3176_vmss3176_0_OsDisk_1_6d72b805e50e4de6830303c5055077fc"),
	// 					},
	// 					DiskSizeGB: to.Ptr[int32](127),
	// 				},
	// 				DataDisks: []*armcompute.DataDisk{
	// 				},
	// 			},
	// 			OSProfile: &armcompute.OSProfile{
	// 				ComputerName: to.Ptr("test000000"),
	// 				AdminUsername: to.Ptr("Foo12"),
	// 				WindowsConfiguration: &armcompute.WindowsConfiguration{
	// 					ProvisionVMAgent: to.Ptr(true),
	// 					EnableAutomaticUpdates: to.Ptr(true),
	// 				},
	// 				Secrets: []*armcompute.VaultSecretGroup{
	// 				},
	// 				AllowExtensionOperations: to.Ptr(true),
	// 				RequireGuestProvisionSignal: to.Ptr(true),
	// 			},
	// 			CapacityReservation: &armcompute.CapacityReservationProfile{
	// 				CapacityReservationGroup: &armcompute.SubResource{
	// 					ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/capacityReservationGroups/{crg-name}"),
	// 				},
	// 			},
	// 			NetworkProfile: &armcompute.NetworkProfile{
	// 				NetworkInterfaces: []*armcompute.NetworkInterfaceReference{
	// 					{
	// 						ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachineScaleSets/{vmss-name}/virtualMachines/0/networkInterfaces/vmsstestnetconfig5415"),
	// 					},
	// 				},
	// 			},
	// 			ProvisioningState: to.Ptr("Succeeded"),
	// 		},
	// 	},
	// }
}
