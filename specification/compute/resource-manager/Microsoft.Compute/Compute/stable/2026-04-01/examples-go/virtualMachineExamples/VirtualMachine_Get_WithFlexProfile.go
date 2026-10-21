package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v9"
)

// Generated from example definition: 2026-04-01/virtualMachineExamples/VirtualMachine_Get_WithFlexProfile.json
func ExampleVirtualMachinesClient_Get_getAVirtualMachineWithAFlexProfile() {
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
	// 		ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/myVM"),
	// 		Name: to.Ptr("myVM"),
	// 		Type: to.Ptr("Microsoft.Compute/virtualMachines"),
	// 		Location: to.Ptr("westus"),
	// 		Tags: map[string]*string{
	// 			"department": to.Ptr("engineering"),
	// 			"project": to.Ptr("alpha"),
	// 		},
	// 		Properties: &armcompute.VirtualMachineProperties{
	// 			VMID: to.Ptr("a149cd25-409f-4e6c-82c1-7c9e34d8e8e8"),
	// 			HardwareProfile: &armcompute.HardwareProfile{
	// 				VMSize: to.Ptr(armcompute.VirtualMachineSizeTypes("Standard_D4as_v5")),
	// 			},
	// 			FlexProfile: &armcompute.FlexProfile{
	// 				PreferredVMSize: to.Ptr("Standard_D4s_v5"),
	// 			},
	// 			StorageProfile: &armcompute.StorageProfile{
	// 				ImageReference: &armcompute.ImageReference{
	// 					Publisher: to.Ptr("Canonical"),
	// 					Offer: to.Ptr("0001-com-ubuntu-server-focal"),
	// 					SKU: to.Ptr("20_04-lts-gen2"),
	// 					Version: to.Ptr("20.04.202406010"),
	// 				},
	// 				OSDisk: &armcompute.OSDisk{
	// 					OSType: to.Ptr(armcompute.OperatingSystemTypesLinux),
	// 					Name: to.Ptr("myVM_OsDisk_1"),
	// 					CreateOption: to.Ptr(armcompute.DiskCreateOptionTypesFromImage),
	// 					Caching: to.Ptr(armcompute.CachingTypesReadWrite),
	// 					ManagedDisk: &armcompute.ManagedDiskParameters{
	// 						ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/myVM_OsDisk_1"),
	// 						StorageAccountType: to.Ptr(armcompute.StorageAccountTypesPremiumLRS),
	// 					},
	// 					DeleteOption: to.Ptr(armcompute.DiskDeleteOptionTypesDelete),
	// 				},
	// 			},
	// 			OSProfile: &armcompute.OSProfile{
	// 				ComputerName: to.Ptr("myVM"),
	// 				AdminUsername: to.Ptr("azureuser"),
	// 				LinuxConfiguration: &armcompute.LinuxConfiguration{
	// 					DisablePasswordAuthentication: to.Ptr(true),
	// 					SSH: &armcompute.SSHConfiguration{
	// 						PublicKeys: []*armcompute.SSHPublicKey{
	// 							{
	// 								Path: to.Ptr("/home/azureuser/.ssh/authorized_keys"),
	// 								KeyData: to.Ptr("ssh-rsa AAAA...{example}"),
	// 							},
	// 						},
	// 					},
	// 					ProvisionVMAgent: to.Ptr(true),
	// 				},
	// 				AllowExtensionOperations: to.Ptr(true),
	// 			},
	// 			NetworkProfile: &armcompute.NetworkProfile{
	// 				NetworkInterfaces: []*armcompute.NetworkInterfaceReference{
	// 					{
	// 						ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/networkInterfaces/myNic"),
	// 						Properties: &armcompute.NetworkInterfaceReferenceProperties{
	// 							Primary: to.Ptr(true),
	// 							DeleteOption: to.Ptr(armcompute.DeleteOptionsDetach),
	// 						},
	// 					},
	// 				},
	// 			},
	// 			ProvisioningState: to.Ptr("Succeeded"),
	// 			TimeCreated: to.Ptr(time.Date(2026, time.April, 1, 12, 34, 56, 123456700, time.UTC)),
	// 		},
	// 	},
	// }
}
