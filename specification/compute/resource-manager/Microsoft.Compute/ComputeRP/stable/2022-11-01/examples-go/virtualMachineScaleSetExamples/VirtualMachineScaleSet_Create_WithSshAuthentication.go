package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-11-01/examples/virtualMachineScaleSetExamples/VirtualMachineScaleSet_Create_WithSshAuthentication.json
func ExampleVirtualMachineScaleSetsClient_BeginCreateOrUpdate_createAScaleSetWithSshAuthentication() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcompute.NewVirtualMachineScaleSetsClient("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx, "myResourceGroup", "{vmss-name}", armcompute.VirtualMachineScaleSet{
		Location: to.Ptr("westus"),
		Properties: &armcompute.VirtualMachineScaleSetProperties{
			Overprovision: to.Ptr(true),
			UpgradePolicy: &armcompute.UpgradePolicy{
				Mode: to.Ptr(armcompute.UpgradeModeManual),
			},
			VirtualMachineProfile: &armcompute.VirtualMachineScaleSetVMProfile{
				NetworkProfile: &armcompute.VirtualMachineScaleSetNetworkProfile{
					NetworkInterfaceConfigurations: []*armcompute.VirtualMachineScaleSetNetworkConfiguration{
						{
							Name: to.Ptr("{vmss-name}"),
							Properties: &armcompute.VirtualMachineScaleSetNetworkConfigurationProperties{
								EnableIPForwarding: to.Ptr(true),
								IPConfigurations: []*armcompute.VirtualMachineScaleSetIPConfiguration{
									{
										Name: to.Ptr("{vmss-name}"),
										Properties: &armcompute.VirtualMachineScaleSetIPConfigurationProperties{
											Subnet: &armcompute.APIEntityReference{
												ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/{existing-virtual-network-name}/subnets/{existing-subnet-name}"),
											},
										},
									}},
								Primary: to.Ptr(true),
							},
						}},
				},
				OSProfile: &armcompute.VirtualMachineScaleSetOSProfile{
					AdminUsername:      to.Ptr("{your-username}"),
					ComputerNamePrefix: to.Ptr("{vmss-name}"),
					LinuxConfiguration: &armcompute.LinuxConfiguration{
						DisablePasswordAuthentication: to.Ptr(true),
						SSH: &armcompute.SSHConfiguration{
							PublicKeys: []*armcompute.SSHPublicKey{
								{
									Path:    to.Ptr("/home/{your-username}/.ssh/authorized_keys"),
									KeyData: to.Ptr("ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCeClRAk2ipUs/l5voIsDC5q9RI+YSRd1Bvd/O+axgY4WiBzG+4FwJWZm/mLLe5DoOdHQwmU2FrKXZSW4w2sYE70KeWnrFViCOX5MTVvJgPE8ClugNl8RWth/tU849DvM9sT7vFgfVSHcAS2yDRyDlueii+8nF2ym8XWAPltFVCyLHRsyBp5YPqK8JFYIa1eybKsY3hEAxRCA+/7bq8et+Gj3coOsuRmrehav7rE6N12Pb80I6ofa6SM5XNYq4Xk0iYNx7R3kdz0Jj9XgZYWjAHjJmT0gTRoOnt6upOuxK7xI/ykWrllgpXrCPu3Ymz+c+ujaqcxDopnAl2lmf69/J1"),
								}},
						},
					},
				},
				StorageProfile: &armcompute.VirtualMachineScaleSetStorageProfile{
					ImageReference: &armcompute.ImageReference{
						Offer:     to.Ptr("WindowsServer"),
						Publisher: to.Ptr("MicrosoftWindowsServer"),
						SKU:       to.Ptr("2016-Datacenter"),
						Version:   to.Ptr("latest"),
					},
					OSDisk: &armcompute.VirtualMachineScaleSetOSDisk{
						Caching:      to.Ptr(armcompute.CachingTypesReadWrite),
						CreateOption: to.Ptr(armcompute.DiskCreateOptionTypesFromImage),
						ManagedDisk: &armcompute.VirtualMachineScaleSetManagedDiskParameters{
							StorageAccountType: to.Ptr(armcompute.StorageAccountTypesStandardLRS),
						},
					},
				},
			},
		},
		SKU: &armcompute.SKU{
			Name:     to.Ptr("Standard_D1_v2"),
			Capacity: to.Ptr[int64](3),
			Tier:     to.Ptr("Standard"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
