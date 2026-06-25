package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v8"
)

// Generated from example definition: 2026-03-01/virtualMachineExamples/VirtualMachine_Get_WithVMSizeProperties.json
func ExampleVirtualMachinesClient_Get_getAVirtualMachineWithVMSizeProperties() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVirtualMachinesClient().Get(ctx, "myResourceGroup", "myVM", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcompute.VirtualMachinesClientGetResponse{
	// 	VirtualMachine: armcompute.VirtualMachine{
	// 		Name: to.Ptr("myVM"),
	// 		ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/myVM"),
	// 		Type: to.Ptr("Microsoft.Compute/virtualMachines"),
	// 		Location: to.Ptr("West US"),
	// 		Tags: map[string]*string{
	// 			"myTag1": to.Ptr("tagValue1"),
	// 		},
	// 		Properties: &armcompute.VirtualMachineProperties{
	// 			VMID: to.Ptr("0f47b100-583c-48e3-a4c0-aefc2c9bbcc1"),
	// 			AvailabilitySet: &armcompute.SubResource{
	// 				ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/availabilitySets/my-AvailabilitySet"),
	// 			},
	// 			HardwareProfile: &armcompute.HardwareProfile{
	// 				VMSize: to.Ptr(armcompute.VirtualMachineSizeTypesStandardDS3V2),
	// 				VMSizeProperties: &armcompute.VMSizeProperties{
	// 					VCPUsAvailable: to.Ptr[int32](1),
	// 					VCPUsPerCore: to.Ptr[int32](1),
	// 				},
	// 			},
	// 			StorageProfile: &armcompute.StorageProfile{
	// 				ImageReference: &armcompute.ImageReference{
	// 					Publisher: to.Ptr("MicrosoftWindowsServer"),
	// 					Offer: to.Ptr("WindowsServer"),
	// 					SKU: to.Ptr("2016-Datacenter"),
	// 					Version: to.Ptr("latest"),
	// 				},
	// 				OSDisk: &armcompute.OSDisk{
	// 					OSType: to.Ptr(armcompute.OperatingSystemTypesWindows),
	// 					Name: to.Ptr("myOsDisk"),
	// 					CreateOption: to.Ptr(armcompute.DiskCreateOptionTypesFromImage),
	// 					Caching: to.Ptr(armcompute.CachingTypesReadWrite),
	// 					ManagedDisk: &armcompute.ManagedDiskParameters{
	// 						StorageAccountType: to.Ptr(armcompute.StorageAccountTypesPremiumLRS),
	// 						ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/myOsDisk"),
	// 					},
	// 					DiskSizeGB: to.Ptr[int32](30),
	// 				},
	// 				DataDisks: []*armcompute.DataDisk{
	// 					{
	// 						Lun: to.Ptr[int32](0),
	// 						Name: to.Ptr("myDataDisk0"),
	// 						CreateOption: to.Ptr(armcompute.DiskCreateOptionTypesEmpty),
	// 						Caching: to.Ptr(armcompute.CachingTypesReadWrite),
	// 						ManagedDisk: &armcompute.ManagedDiskParameters{
	// 							StorageAccountType: to.Ptr(armcompute.StorageAccountTypesPremiumLRS),
	// 							ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/myDataDisk0"),
	// 						},
	// 						DiskSizeGB: to.Ptr[int32](30),
	// 					},
	// 					{
	// 						Lun: to.Ptr[int32](1),
	// 						Name: to.Ptr("myDataDisk1"),
	// 						CreateOption: to.Ptr(armcompute.DiskCreateOptionTypesAttach),
	// 						Caching: to.Ptr(armcompute.CachingTypesReadWrite),
	// 						ManagedDisk: &armcompute.ManagedDiskParameters{
	// 							StorageAccountType: to.Ptr(armcompute.StorageAccountTypesPremiumLRS),
	// 							ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/myDataDisk1"),
	// 						},
	// 						DiskSizeGB: to.Ptr[int32](100),
	// 					},
	// 				},
	// 			},
	// 			ApplicationProfile: &armcompute.ApplicationProfile{
	// 				GalleryApplications: []*armcompute.VMGalleryApplication{
	// 					{
	// 						Tags: to.Ptr("myTag1"),
	// 						Order: to.Ptr[int32](1),
	// 						PackageReferenceID: to.Ptr("/subscriptions/32c17a9e-aa7b-4ba5-a45b-e324116b6fdb/resourceGroups/myresourceGroupName2/providers/Microsoft.Compute/galleries/myGallery1/applications/MyApplication1/versions/1.0"),
	// 						ConfigurationReference: to.Ptr("https://mystorageaccount.blob.core.windows.net/configurations/settings.config"),
	// 					},
	// 					{
	// 						PackageReferenceID: to.Ptr("/subscriptions/32c17a9e-aa7b-4ba5-a45b-e324116b6fdg/resourceGroups/myresourceGroupName3/providers/Microsoft.Compute/galleries/myGallery2/applications/MyApplication2/versions/1.1"),
	// 					},
	// 				},
	// 			},
	// 			UserData: to.Ptr("RXhhbXBsZSBVc2VyRGF0YQ=="),
	// 			OSProfile: &armcompute.OSProfile{
	// 				ComputerName: to.Ptr("myVM"),
	// 				AdminUsername: to.Ptr("admin"),
	// 				WindowsConfiguration: &armcompute.WindowsConfiguration{
	// 					ProvisionVMAgent: to.Ptr(true),
	// 					EnableAutomaticUpdates: to.Ptr(false),
	// 				},
	// 				Secrets: []*armcompute.VaultSecretGroup{
	// 				},
	// 			},
	// 			SecurityProfile: &armcompute.SecurityProfile{
	// 				SecurityType: to.Ptr(armcompute.SecurityTypesStandard),
	// 			},
	// 			NetworkProfile: &armcompute.NetworkProfile{
	// 				NetworkInterfaces: []*armcompute.NetworkInterfaceReference{
	// 					{
	// 						ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/networkInterfaces/{myNIC}"),
	// 					},
	// 				},
	// 			},
	// 			DiagnosticsProfile: &armcompute.DiagnosticsProfile{
	// 				BootDiagnostics: &armcompute.BootDiagnostics{
	// 					Enabled: to.Ptr(true),
	// 					StorageURI: to.Ptr("http://{myStorageAccount}.blob.core.windows.net"),
	// 				},
	// 			},
	// 			ExtensionsTimeBudget: to.Ptr("PT50M"),
	// 			ProvisioningState: to.Ptr("Succeeded"),
	// 		},
	// 		Resources: []*armcompute.VirtualMachineExtension{
	// 			{
	// 				Name: to.Ptr("CustomScriptExtension-DSC"),
	// 				ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/myVM/extensions/CustomScriptExtension-DSC"),
	// 				Type: to.Ptr("Microsoft.Compute/virtualMachines/extensions"),
	// 				Location: to.Ptr("west us"),
	// 				Tags: map[string]*string{
	// 					"displayName": to.Ptr("CustomScriptExtension-DSC"),
	// 				},
	// 				Properties: &armcompute.VirtualMachineExtensionProperties{
	// 					AutoUpgradeMinorVersion: to.Ptr(true),
	// 					ProvisioningState: to.Ptr("Succeeded"),
	// 					Publisher: to.Ptr("Microsoft.Compute"),
	// 					Type: to.Ptr("CustomScriptExtension"),
	// 					TypeHandlerVersion: to.Ptr("1.9"),
	// 					Settings: map[string]any{
	// 					},
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
