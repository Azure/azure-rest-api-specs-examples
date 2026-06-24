package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v8"
)

// Generated from example definition: 2026-03-01/virtualMachineScaleSetExamples/VirtualMachineScaleSet_Get_WithDiskControllerType.json
func ExampleVirtualMachineScaleSetsClient_Get_getVMScaleSetVMWithDiskControllerType() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVirtualMachineScaleSetsClient().Get(ctx, "myResourceGroup", "myVirtualMachineScaleSet", &armcompute.VirtualMachineScaleSetsClientGetOptions{
		Expand: to.Ptr(armcompute.ExpandTypesForGetVMScaleSetsUserData)})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcompute.VirtualMachineScaleSetsClientGetResponse{
	// 	VirtualMachineScaleSet: armcompute.VirtualMachineScaleSet{
	// 		Name: to.Ptr("myVirtualMachineScaleSet"),
	// 		ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachineScaleSets/myVirtualMachineScaleSet"),
	// 		Type: to.Ptr("Microsoft.Compute/virtualMachineScaleSets"),
	// 		Location: to.Ptr("westus"),
	// 		Tags: map[string]*string{
	// 			"myTag1": to.Ptr("tagValue1"),
	// 		},
	// 		SKU: &armcompute.SKU{
	// 			Name: to.Ptr("Standard_D2s_v3"),
	// 			Tier: to.Ptr("Standard"),
	// 			Capacity: to.Ptr[int64](4),
	// 		},
	// 		Properties: &armcompute.VirtualMachineScaleSetProperties{
	// 			SinglePlacementGroup: to.Ptr(false),
	// 			UpgradePolicy: &armcompute.UpgradePolicy{
	// 				Mode: to.Ptr(armcompute.UpgradeModeAutomatic),
	// 				AutomaticOSUpgradePolicy: &armcompute.AutomaticOSUpgradePolicy{
	// 					EnableAutomaticOSUpgrade: to.Ptr(false),
	// 				},
	// 			},
	// 			VirtualMachineProfile: &armcompute.VirtualMachineScaleSetVMProfile{
	// 				StorageProfile: &armcompute.VirtualMachineScaleSetStorageProfile{
	// 					OSDisk: &armcompute.VirtualMachineScaleSetOSDisk{
	// 						CreateOption: to.Ptr(armcompute.DiskCreateOptionTypesFromImage),
	// 						Caching: to.Ptr(armcompute.CachingTypesReadWrite),
	// 						ManagedDisk: &armcompute.VirtualMachineScaleSetManagedDiskParameters{
	// 							StorageAccountType: to.Ptr(armcompute.StorageAccountTypesPremiumLRS),
	// 						},
	// 						DiskSizeGB: to.Ptr[int32](30),
	// 					},
	// 					ImageReference: &armcompute.ImageReference{
	// 						Publisher: to.Ptr("azuredatabricks"),
	// 						Offer: to.Ptr("databricks"),
	// 						SKU: to.Ptr("databricksworker"),
	// 						Version: to.Ptr("3.15.2"),
	// 					},
	// 					DataDisks: []*armcompute.VirtualMachineScaleSetDataDisk{
	// 					},
	// 					DiskControllerType: to.Ptr("NVMe"),
	// 				},
	// 				ApplicationProfile: &armcompute.ApplicationProfile{
	// 					GalleryApplications: []*armcompute.VMGalleryApplication{
	// 						{
	// 							Tags: to.Ptr("myTag1"),
	// 							Order: to.Ptr[int32](1),
	// 							PackageReferenceID: to.Ptr("/subscriptions/32c17a9e-aa7b-4ba5-a45b-e324116b6fdb/resourceGroups/myresourceGroupName2/providers/Microsoft.Compute/galleries/myGallery1/applications/MyApplication1/versions/1.0"),
	// 							ConfigurationReference: to.Ptr("https://mystorageaccount.blob.core.windows.net/configurations/settings.config"),
	// 						},
	// 						{
	// 							PackageReferenceID: to.Ptr("/subscriptions/32c17a9e-aa7b-4ba5-a45b-e324116b6fdg/resourceGroups/myresourceGroupName3/providers/Microsoft.Compute/galleries/myGallery2/applications/MyApplication2/versions/1.1"),
	// 						},
	// 					},
	// 				},
	// 				UserData: to.Ptr("RXhhbXBsZSBVc2VyRGF0YQ=="),
	// 				OSProfile: &armcompute.VirtualMachineScaleSetOSProfile{
	// 					ComputerNamePrefix: to.Ptr("myVirtualMachineScaleSet"),
	// 					AdminUsername: to.Ptr("admin"),
	// 					LinuxConfiguration: &armcompute.LinuxConfiguration{
	// 						DisablePasswordAuthentication: to.Ptr(false),
	// 					},
	// 				},
	// 				NetworkProfile: &armcompute.VirtualMachineScaleSetNetworkProfile{
	// 					NetworkInterfaceConfigurations: []*armcompute.VirtualMachineScaleSetNetworkConfiguration{
	// 						{
	// 							Name: to.Ptr("myNic"),
	// 							Properties: &armcompute.VirtualMachineScaleSetNetworkConfigurationProperties{
	// 								Primary: to.Ptr(true),
	// 								IPConfigurations: []*armcompute.VirtualMachineScaleSetIPConfiguration{
	// 									{
	// 										Name: to.Ptr("myIPConfig"),
	// 										Properties: &armcompute.VirtualMachineScaleSetIPConfigurationProperties{
	// 											Primary: to.Ptr(true),
	// 											Subnet: &armcompute.APIEntityReference{
	// 												ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/myVNet/subnets/mySubnet"),
	// 											},
	// 										},
	// 									},
	// 								},
	// 								NetworkSecurityGroup: &armcompute.SubResource{
	// 									ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/networkSecurityGroups/myNetworkSecurityGroup"),
	// 								},
	// 							},
	// 						},
	// 					},
	// 				},
	// 				SecurityProfile: &armcompute.SecurityProfile{
	// 					SecurityType: to.Ptr(armcompute.SecurityTypesStandard),
	// 				},
	// 			},
	// 			ProvisioningState: to.Ptr("succeeded"),
	// 			Overprovision: to.Ptr(false),
	// 			DoNotRunExtensionsOnOverprovisionedVMs: to.Ptr(false),
	// 			PlatformFaultDomainCount: to.Ptr[int32](1),
	// 			HostGroup: &armcompute.SubResource{
	// 				ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/hostGroups/myHostGroup"),
	// 			},
	// 		},
	// 	},
	// }
}
