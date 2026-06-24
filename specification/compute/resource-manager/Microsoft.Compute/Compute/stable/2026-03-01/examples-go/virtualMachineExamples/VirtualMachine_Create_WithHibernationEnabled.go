package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v8"
)

// Generated from example definition: 2026-03-01/virtualMachineExamples/VirtualMachine_Create_WithHibernationEnabled.json
func ExampleVirtualMachinesClient_BeginCreateOrUpdate_createAVMWithHibernationEnabled() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualMachinesClient().BeginCreateOrUpdate(ctx, "myResourceGroup", "{vm-name}", armcompute.VirtualMachine{
		Location: to.Ptr("eastus2euap"),
		Properties: &armcompute.VirtualMachineProperties{
			HardwareProfile: &armcompute.HardwareProfile{
				VMSize: to.Ptr(armcompute.VirtualMachineSizeTypesStandardD2SV3),
			},
			AdditionalCapabilities: &armcompute.AdditionalCapabilities{
				HibernationEnabled: to.Ptr(true),
			},
			StorageProfile: &armcompute.StorageProfile{
				ImageReference: &armcompute.ImageReference{
					Publisher: to.Ptr("MicrosoftWindowsServer"),
					Offer:     to.Ptr("WindowsServer"),
					SKU:       to.Ptr("2019-Datacenter"),
					Version:   to.Ptr("latest"),
				},
				OSDisk: &armcompute.OSDisk{
					Caching: to.Ptr(armcompute.CachingTypesReadWrite),
					ManagedDisk: &armcompute.ManagedDiskParameters{
						StorageAccountType: to.Ptr(armcompute.StorageAccountTypesStandardLRS),
					},
					Name:         to.Ptr("vmOSdisk"),
					CreateOption: to.Ptr(armcompute.DiskCreateOptionTypesFromImage),
				},
			},
			NetworkProfile: &armcompute.NetworkProfile{
				NetworkInterfaces: []*armcompute.NetworkInterfaceReference{
					{
						ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/networkInterfaces/{existing-nic-name}"),
						Properties: &armcompute.NetworkInterfaceReferenceProperties{
							Primary: to.Ptr(true),
						},
					},
				},
			},
			OSProfile: &armcompute.OSProfile{
				AdminUsername: to.Ptr("{your-username}"),
				ComputerName:  to.Ptr("{vm-name}"),
				AdminPassword: to.Ptr("{your-password}"),
			},
			DiagnosticsProfile: &armcompute.DiagnosticsProfile{
				BootDiagnostics: &armcompute.BootDiagnostics{
					StorageURI: to.Ptr("http://{existing-storage-account-name}.blob.core.windows.net"),
					Enabled:    to.Ptr(true),
				},
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
	// res = armcompute.VirtualMachinesClientCreateOrUpdateResponse{
	// 	VirtualMachine: armcompute.VirtualMachine{
	// 		ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/{vm-name}"),
	// 		Type: to.Ptr("Microsoft.Compute/virtualMachines"),
	// 		Properties: &armcompute.VirtualMachineProperties{
	// 			OSProfile: &armcompute.OSProfile{
	// 				AdminUsername: to.Ptr("{your-username}"),
	// 				Secrets: []*armcompute.VaultSecretGroup{
	// 				},
	// 				ComputerName: to.Ptr("{vm-name}"),
	// 				WindowsConfiguration: &armcompute.WindowsConfiguration{
	// 					ProvisionVMAgent: to.Ptr(true),
	// 					EnableAutomaticUpdates: to.Ptr(true),
	// 				},
	// 			},
	// 			NetworkProfile: &armcompute.NetworkProfile{
	// 				NetworkInterfaces: []*armcompute.NetworkInterfaceReference{
	// 					{
	// 						ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/networkInterfaces/nsgExistingNic"),
	// 						Properties: &armcompute.NetworkInterfaceReferenceProperties{
	// 							Primary: to.Ptr(true),
	// 						},
	// 					},
	// 				},
	// 			},
	// 			StorageProfile: &armcompute.StorageProfile{
	// 				ImageReference: &armcompute.ImageReference{
	// 					Publisher: to.Ptr("MicrosoftWindowsServer"),
	// 					Offer: to.Ptr("WindowsServer"),
	// 					SKU: to.Ptr("2019-Datacenter"),
	// 					Version: to.Ptr("latest"),
	// 				},
	// 				OSDisk: &armcompute.OSDisk{
	// 					OSType: to.Ptr(armcompute.OperatingSystemTypesWindows),
	// 					Caching: to.Ptr(armcompute.CachingTypesReadWrite),
	// 					CreateOption: to.Ptr(armcompute.DiskCreateOptionTypesFromImage),
	// 					Name: to.Ptr("vmOSdisk"),
	// 					ManagedDisk: &armcompute.ManagedDiskParameters{
	// 						StorageAccountType: to.Ptr(armcompute.StorageAccountTypesStandardLRS),
	// 					},
	// 				},
	// 				DataDisks: []*armcompute.DataDisk{
	// 				},
	// 			},
	// 			DiagnosticsProfile: &armcompute.DiagnosticsProfile{
	// 				BootDiagnostics: &armcompute.BootDiagnostics{
	// 					StorageURI: to.Ptr("http://nsgdiagnostic.blob.core.windows.net"),
	// 					Enabled: to.Ptr(true),
	// 				},
	// 			},
	// 			VMID: to.Ptr("676420ba-7a24-4bfe-80bd-9c841ee184fa"),
	// 			HardwareProfile: &armcompute.HardwareProfile{
	// 				VMSize: to.Ptr(armcompute.VirtualMachineSizeTypesStandardD2SV3),
	// 			},
	// 			AdditionalCapabilities: &armcompute.AdditionalCapabilities{
	// 				HibernationEnabled: to.Ptr(true),
	// 			},
	// 			SecurityProfile: &armcompute.SecurityProfile{
	// 				SecurityType: to.Ptr(armcompute.SecurityTypesStandard),
	// 			},
	// 			ProvisioningState: to.Ptr("Updating"),
	// 		},
	// 		Name: to.Ptr("{vm-name}"),
	// 		Location: to.Ptr("eastus2euap"),
	// 	},
	// }
}
