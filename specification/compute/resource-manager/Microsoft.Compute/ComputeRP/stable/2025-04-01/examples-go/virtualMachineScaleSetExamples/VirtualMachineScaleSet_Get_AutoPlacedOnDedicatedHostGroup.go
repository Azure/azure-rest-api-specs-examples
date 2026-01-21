package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ccea12be6cd5e64743499d46f7a9c8b60df52db1/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2025-04-01/examples/virtualMachineScaleSetExamples/VirtualMachineScaleSet_Get_AutoPlacedOnDedicatedHostGroup.json
func ExampleVirtualMachineScaleSetsClient_Get_getAVirtualMachineScaleSetPlacedOnADedicatedHostGroupThroughAutomaticPlacement() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVirtualMachineScaleSetsClient().Get(ctx, "myResourceGroup", "myVirtualMachineScaleSet", &armcompute.VirtualMachineScaleSetsClientGetOptions{Expand: nil})
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
	// 	Location: to.Ptr("West US"),
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
	// 		},
	// 	},
	// 	SKU: &armcompute.SKU{
	// 		Name: to.Ptr("Standard_D2s_v3"),
	// 		Capacity: to.Ptr[int64](4),
	// 		Tier: to.Ptr("Standard"),
	// 	},
	// }
}
