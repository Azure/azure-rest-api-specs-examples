package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ccea12be6cd5e64743499d46f7a9c8b60df52db1/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2025-04-01/examples/virtualMachineScaleSetExamples/VirtualMachineScaleSet_Get_WithDiskControllerType.json
func ExampleVirtualMachineScaleSetsClient_Get_getVmScaleSetVmWithDiskControllerType() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVirtualMachineScaleSetsClient().Get(ctx, "myResourceGroup", "myVirtualMachineScaleSet", &armcompute.VirtualMachineScaleSetsClientGetOptions{Expand: to.Ptr(armcompute.ExpandTypesForGetVMScaleSetsUserData)})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.VirtualMachineScaleSet = armcompute.VirtualMachineScaleSet{
	// 	Name: to.Ptr("myVirtualMachineScaleSet"),
	// 	Type: to.Ptr("Microsoft.Compute/virtualMachineScaleSets"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachineScaleSets/myVirtualMachineScaleSet"),
	// 	Location: to.Ptr("westus"),
	// 	Tags: map[string]*string{
	// 		"myTag1": to.Ptr("tagValue1"),
	// 	},
	// 	Properties: &armcompute.VirtualMachineScaleSetProperties{
	// 		DoNotRunExtensionsOnOverprovisionedVMs: to.Ptr(false),
	// 		HostGroup: &armcompute.SubResource{
	// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/hostGroups/myHostGroup"),
	// 		},
	// 		Overprovision: to.Ptr(false),
	// 		PlatformFaultDomainCount: to.Ptr[int32](1),
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		SinglePlacementGroup: to.Ptr(false),
	// 		UpgradePolicy: &armcompute.UpgradePolicy{
	// 			AutomaticOSUpgradePolicy: &armcompute.AutomaticOSUpgradePolicy{
	// 				EnableAutomaticOSUpgrade: to.Ptr(false),
	// 			},
	// 			Mode: to.Ptr(armcompute.UpgradeModeAutomatic),
	// 		},
	// 		VirtualMachineProfile: &armcompute.VirtualMachineScaleSetVMProfile{
	// 			ApplicationProfile: &armcompute.ApplicationProfile{
	// 				GalleryApplications: []*armcompute.VMGalleryApplication{
	// 					{
	// 						ConfigurationReference: to.Ptr("https://mystorageaccount.blob.core.windows.net/configurations/settings.config"),
	// 						Order: to.Ptr[int32](1),
	// 						PackageReferenceID: to.Ptr("/subscriptions/32c17a9e-aa7b-4ba5-a45b-e324116b6fdb/resourceGroups/myresourceGroupName2/providers/Microsoft.Compute/galleries/myGallery1/applications/MyApplication1/versions/1.0"),
	// 						Tags: to.Ptr("myTag1"),
	// 					},
	// 					{
	// 						PackageReferenceID: to.Ptr("/subscriptions/32c17a9e-aa7b-4ba5-a45b-e324116b6fdg/resourceGroups/myresourceGroupName3/providers/Microsoft.Compute/galleries/myGallery2/applications/MyApplication2/versions/1.1"),
	// 				}},
	// 			},
	// 			NetworkProfile: &armcompute.VirtualMachineScaleSetNetworkProfile{
	// 				NetworkInterfaceConfigurations: []*armcompute.VirtualMachineScaleSetNetworkConfiguration{
	// 					{
	// 						Name: to.Ptr("myNic"),
	// 						Properties: &armcompute.VirtualMachineScaleSetNetworkConfigurationProperties{
	// 							IPConfigurations: []*armcompute.VirtualMachineScaleSetIPConfiguration{
	// 								{
	// 									Name: to.Ptr("myIPConfig"),
	// 									Properties: &armcompute.VirtualMachineScaleSetIPConfigurationProperties{
	// 										Primary: to.Ptr(true),
	// 										Subnet: &armcompute.APIEntityReference{
	// 											ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/myVNet/subnets/mySubnet"),
	// 										},
	// 									},
	// 							}},
	// 							NetworkSecurityGroup: &armcompute.SubResource{
	// 								ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/networkSecurityGroups/myNetworkSecurityGroup"),
	// 							},
	// 							Primary: to.Ptr(true),
	// 						},
	// 				}},
	// 			},
	// 			OSProfile: &armcompute.VirtualMachineScaleSetOSProfile{
	// 				AdminUsername: to.Ptr("admin"),
	// 				ComputerNamePrefix: to.Ptr("myVirtualMachineScaleSet"),
	// 				LinuxConfiguration: &armcompute.LinuxConfiguration{
	// 					DisablePasswordAuthentication: to.Ptr(false),
	// 				},
	// 			},
	// 			StorageProfile: &armcompute.VirtualMachineScaleSetStorageProfile{
	// 				DataDisks: []*armcompute.VirtualMachineScaleSetDataDisk{
	// 				},
	// 				DiskControllerType: to.Ptr("NVMe"),
	// 				ImageReference: &armcompute.ImageReference{
	// 					Offer: to.Ptr("databricks"),
	// 					Publisher: to.Ptr("azuredatabricks"),
	// 					SKU: to.Ptr("databricksworker"),
	// 					Version: to.Ptr("3.15.2"),
	// 				},
	// 				OSDisk: &armcompute.VirtualMachineScaleSetOSDisk{
	// 					Caching: to.Ptr(armcompute.CachingTypesReadWrite),
	// 					CreateOption: to.Ptr(armcompute.DiskCreateOptionTypesFromImage),
	// 					DiskSizeGB: to.Ptr[int32](30),
	// 					ManagedDisk: &armcompute.VirtualMachineScaleSetManagedDiskParameters{
	// 						StorageAccountType: to.Ptr(armcompute.StorageAccountTypesPremiumLRS),
	// 					},
	// 				},
	// 			},
	// 			UserData: to.Ptr("RXhhbXBsZSBVc2VyRGF0YQ=="),
	// 		},
	// 	},
	// 	SKU: &armcompute.SKU{
	// 		Name: to.Ptr("Standard_D2s_v3"),
	// 		Capacity: to.Ptr[int64](4),
	// 		Tier: to.Ptr("Standard"),
	// 	},
	// }
}
