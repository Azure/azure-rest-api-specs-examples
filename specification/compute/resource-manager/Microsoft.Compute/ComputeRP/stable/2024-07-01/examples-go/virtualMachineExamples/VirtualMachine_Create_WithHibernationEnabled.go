package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v6"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/069a65e8a6d1a6c0c58d9a9d97610b7103b6e8a5/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-07-01/examples/virtualMachineExamples/VirtualMachine_Create_WithHibernationEnabled.json
func ExampleVirtualMachinesClient_BeginCreateOrUpdate_createAVmWithHibernationEnabled() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualMachinesClient().BeginCreateOrUpdate(ctx, "myResourceGroup", "{vm-name}", armcompute.VirtualMachine{
		Location: to.Ptr("eastus2euap"),
		Properties: &armcompute.VirtualMachineProperties{
			AdditionalCapabilities: &armcompute.AdditionalCapabilities{
				HibernationEnabled: to.Ptr(true),
			},
			DiagnosticsProfile: &armcompute.DiagnosticsProfile{
				BootDiagnostics: &armcompute.BootDiagnostics{
					Enabled:    to.Ptr(true),
					StorageURI: to.Ptr("http://{existing-storage-account-name}.blob.core.windows.net"),
				},
			},
			HardwareProfile: &armcompute.HardwareProfile{
				VMSize: to.Ptr(armcompute.VirtualMachineSizeTypesStandardD2SV3),
			},
			NetworkProfile: &armcompute.NetworkProfile{
				NetworkInterfaces: []*armcompute.NetworkInterfaceReference{
					{
						ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/networkInterfaces/{existing-nic-name}"),
						Properties: &armcompute.NetworkInterfaceReferenceProperties{
							Primary: to.Ptr(true),
						},
					}},
			},
			OSProfile: &armcompute.OSProfile{
				AdminPassword: to.Ptr("{your-password}"),
				AdminUsername: to.Ptr("{your-username}"),
				ComputerName:  to.Ptr("{vm-name}"),
			},
			StorageProfile: &armcompute.StorageProfile{
				ImageReference: &armcompute.ImageReference{
					Offer:     to.Ptr("WindowsServer"),
					Publisher: to.Ptr("MicrosoftWindowsServer"),
					SKU:       to.Ptr("2019-Datacenter"),
					Version:   to.Ptr("latest"),
				},
				OSDisk: &armcompute.OSDisk{
					Name:         to.Ptr("vmOSdisk"),
					Caching:      to.Ptr(armcompute.CachingTypesReadWrite),
					CreateOption: to.Ptr(armcompute.DiskCreateOptionTypesFromImage),
					ManagedDisk: &armcompute.ManagedDiskParameters{
						StorageAccountType: to.Ptr(armcompute.StorageAccountTypesStandardLRS),
					},
				},
			},
		},
	}, &armcompute.VirtualMachinesClientBeginCreateOrUpdateOptions{IfMatch: nil,
		IfNoneMatch: nil,
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.VirtualMachine = armcompute.VirtualMachine{
	// 	Name: to.Ptr("{vm-name}"),
	// 	Type: to.Ptr("Microsoft.Compute/virtualMachines"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/{vm-name}"),
	// 	Location: to.Ptr("eastus2euap"),
	// 	Properties: &armcompute.VirtualMachineProperties{
	// 		AdditionalCapabilities: &armcompute.AdditionalCapabilities{
	// 			HibernationEnabled: to.Ptr(true),
	// 		},
	// 		DiagnosticsProfile: &armcompute.DiagnosticsProfile{
	// 			BootDiagnostics: &armcompute.BootDiagnostics{
	// 				Enabled: to.Ptr(true),
	// 				StorageURI: to.Ptr("http://nsgdiagnostic.blob.core.windows.net"),
	// 			},
	// 		},
	// 		HardwareProfile: &armcompute.HardwareProfile{
	// 			VMSize: to.Ptr(armcompute.VirtualMachineSizeTypesStandardD2SV3),
	// 		},
	// 		NetworkProfile: &armcompute.NetworkProfile{
	// 			NetworkInterfaces: []*armcompute.NetworkInterfaceReference{
	// 				{
	// 					ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/networkInterfaces/nsgExistingNic"),
	// 					Properties: &armcompute.NetworkInterfaceReferenceProperties{
	// 						Primary: to.Ptr(true),
	// 					},
	// 			}},
	// 		},
	// 		OSProfile: &armcompute.OSProfile{
	// 			AdminUsername: to.Ptr("{your-username}"),
	// 			ComputerName: to.Ptr("{vm-name}"),
	// 			Secrets: []*armcompute.VaultSecretGroup{
	// 			},
	// 			WindowsConfiguration: &armcompute.WindowsConfiguration{
	// 				EnableAutomaticUpdates: to.Ptr(true),
	// 				ProvisionVMAgent: to.Ptr(true),
	// 			},
	// 		},
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		StorageProfile: &armcompute.StorageProfile{
	// 			DataDisks: []*armcompute.DataDisk{
	// 			},
	// 			ImageReference: &armcompute.ImageReference{
	// 				Offer: to.Ptr("WindowsServer"),
	// 				Publisher: to.Ptr("MicrosoftWindowsServer"),
	// 				SKU: to.Ptr("2019-Datacenter"),
	// 				Version: to.Ptr("latest"),
	// 			},
	// 			OSDisk: &armcompute.OSDisk{
	// 				Name: to.Ptr("vmOSdisk"),
	// 				Caching: to.Ptr(armcompute.CachingTypesReadWrite),
	// 				CreateOption: to.Ptr(armcompute.DiskCreateOptionTypesFromImage),
	// 				ManagedDisk: &armcompute.ManagedDiskParameters{
	// 					StorageAccountType: to.Ptr(armcompute.StorageAccountTypesStandardLRS),
	// 				},
	// 				OSType: to.Ptr(armcompute.OperatingSystemTypesWindows),
	// 			},
	// 		},
	// 		VMID: to.Ptr("676420ba-7a24-4bfe-80bd-9c841ee184fa"),
	// 	},
	// }
}
