package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ccea12be6cd5e64743499d46f7a9c8b60df52db1/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2025-04-01/examples/virtualMachineExamples/VirtualMachine_Get.json
func ExampleVirtualMachinesClient_Get_getAVirtualMachine() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVirtualMachinesClient().Get(ctx, "myResourceGroup", "myVM", &armcompute.VirtualMachinesClientGetOptions{Expand: to.Ptr(armcompute.InstanceViewTypesUserData)})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.VirtualMachine = armcompute.VirtualMachine{
	// 	Name: to.Ptr("myVM"),
	// 	Type: to.Ptr("Microsoft.Compute/virtualMachines"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/myVM"),
	// 	Location: to.Ptr("West US"),
	// 	Tags: map[string]*string{
	// 		"myTag1": to.Ptr("tagValue1"),
	// 	},
	// 	Etag: to.Ptr("\"1\""),
	// 	ManagedBy: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachineScaleSets/{MyVmss}"),
	// 	Properties: &armcompute.VirtualMachineProperties{
	// 		ApplicationProfile: &armcompute.ApplicationProfile{
	// 			GalleryApplications: []*armcompute.VMGalleryApplication{
	// 				{
	// 					ConfigurationReference: to.Ptr("https://mystorageaccount.blob.core.windows.net/configurations/settings.config"),
	// 					Order: to.Ptr[int32](1),
	// 					PackageReferenceID: to.Ptr("/subscriptions/32c17a9e-aa7b-4ba5-a45b-e324116b6fdb/resourceGroups/myresourceGroupName2/providers/Microsoft.Compute/galleries/myGallery1/applications/MyApplication1/versions/1.0"),
	// 					Tags: to.Ptr("myTag1"),
	// 				},
	// 				{
	// 					PackageReferenceID: to.Ptr("/subscriptions/32c17a9e-aa7b-4ba5-a45b-e324116b6fdg/resourceGroups/myresourceGroupName3/providers/Microsoft.Compute/galleries/myGallery2/applications/MyApplication2/versions/1.1"),
	// 			}},
	// 		},
	// 		AvailabilitySet: &armcompute.SubResource{
	// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/availabilitySets/my-AvailabilitySet"),
	// 		},
	// 		DiagnosticsProfile: &armcompute.DiagnosticsProfile{
	// 			BootDiagnostics: &armcompute.BootDiagnostics{
	// 				Enabled: to.Ptr(true),
	// 				StorageURI: to.Ptr("http://{myStorageAccount}.blob.core.windows.net"),
	// 			},
	// 		},
	// 		ExtensionsTimeBudget: to.Ptr("PT50M"),
	// 		HardwareProfile: &armcompute.HardwareProfile{
	// 			VMSize: to.Ptr(armcompute.VirtualMachineSizeTypesStandardDS3V2),
	// 		},
	// 		NetworkProfile: &armcompute.NetworkProfile{
	// 			NetworkInterfaces: []*armcompute.NetworkInterfaceReference{
	// 				{
	// 					ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/networkInterfaces/{myNIC}"),
	// 			}},
	// 		},
	// 		OSProfile: &armcompute.OSProfile{
	// 			AdminUsername: to.Ptr("admin"),
	// 			ComputerName: to.Ptr("myVM"),
	// 			Secrets: []*armcompute.VaultSecretGroup{
	// 			},
	// 			WindowsConfiguration: &armcompute.WindowsConfiguration{
	// 				EnableAutomaticUpdates: to.Ptr(false),
	// 				ProvisionVMAgent: to.Ptr(true),
	// 			},
	// 		},
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		ProximityPlacementGroup: &armcompute.SubResource{
	// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/proximityPlacementGroups/my-ppg01"),
	// 		},
	// 		StorageProfile: &armcompute.StorageProfile{
	// 			DataDisks: []*armcompute.DataDisk{
	// 				{
	// 					Name: to.Ptr("myDataDisk0"),
	// 					Caching: to.Ptr(armcompute.CachingTypesReadWrite),
	// 					CreateOption: to.Ptr(armcompute.DiskCreateOptionTypesEmpty),
	// 					DiskSizeGB: to.Ptr[int32](30),
	// 					Lun: to.Ptr[int32](0),
	// 					ManagedDisk: &armcompute.ManagedDiskParameters{
	// 						ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/myDataDisk0"),
	// 						StorageAccountType: to.Ptr(armcompute.StorageAccountTypesPremiumLRS),
	// 					},
	// 				},
	// 				{
	// 					Name: to.Ptr("myDataDisk1"),
	// 					Caching: to.Ptr(armcompute.CachingTypesReadWrite),
	// 					CreateOption: to.Ptr(armcompute.DiskCreateOptionTypesAttach),
	// 					DiskSizeGB: to.Ptr[int32](100),
	// 					Lun: to.Ptr[int32](1),
	// 					ManagedDisk: &armcompute.ManagedDiskParameters{
	// 						ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/myDataDisk1"),
	// 						StorageAccountType: to.Ptr(armcompute.StorageAccountTypesPremiumLRS),
	// 					},
	// 			}},
	// 			ImageReference: &armcompute.ImageReference{
	// 				Offer: to.Ptr("WindowsServer"),
	// 				Publisher: to.Ptr("MicrosoftWindowsServer"),
	// 				SKU: to.Ptr("2016-Datacenter"),
	// 				Version: to.Ptr("latest"),
	// 			},
	// 			OSDisk: &armcompute.OSDisk{
	// 				Name: to.Ptr("myOsDisk"),
	// 				Caching: to.Ptr(armcompute.CachingTypesReadWrite),
	// 				CreateOption: to.Ptr(armcompute.DiskCreateOptionTypesFromImage),
	// 				DiskSizeGB: to.Ptr[int32](30),
	// 				ManagedDisk: &armcompute.ManagedDiskParameters{
	// 					ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/myOsDisk"),
	// 					StorageAccountType: to.Ptr(armcompute.StorageAccountTypesPremiumLRS),
	// 				},
	// 				OSType: to.Ptr(armcompute.OperatingSystemTypesWindows),
	// 			},
	// 		},
	// 		TimeCreated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-27T01:02:38.313Z"); return t}()),
	// 		UserData: to.Ptr("RXhhbXBsZSBVc2VyRGF0YQ=="),
	// 		VMID: to.Ptr("0f47b100-583c-48e3-a4c0-aefc2c9bbcc1"),
	// 	},
	// 	Resources: []*armcompute.VirtualMachineExtension{
	// 		{
	// 			Name: to.Ptr("CustomScriptExtension-DSC"),
	// 			Type: to.Ptr("Microsoft.Compute/virtualMachines/extensions"),
	// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/myVM/extensions/CustomScriptExtension-DSC"),
	// 			Location: to.Ptr("west us"),
	// 			Tags: map[string]*string{
	// 				"displayName": to.Ptr("CustomScriptExtension-DSC"),
	// 			},
	// 			Properties: &armcompute.VirtualMachineExtensionProperties{
	// 				Type: to.Ptr("CustomScriptExtension"),
	// 				AutoUpgradeMinorVersion: to.Ptr(true),
	// 				ProvisioningState: to.Ptr("Succeeded"),
	// 				Publisher: to.Ptr("Microsoft.Compute"),
	// 				Settings: map[string]any{
	// 				},
	// 				TypeHandlerVersion: to.Ptr("1.9"),
	// 			},
	// 	}},
	// }
}
