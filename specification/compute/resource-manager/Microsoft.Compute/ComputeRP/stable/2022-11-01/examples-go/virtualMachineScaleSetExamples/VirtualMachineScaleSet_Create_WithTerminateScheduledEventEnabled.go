package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-11-01/examples/virtualMachineScaleSetExamples/VirtualMachineScaleSet_Create_WithTerminateScheduledEventEnabled.json
func ExampleVirtualMachineScaleSetsClient_BeginCreateOrUpdate_createAScaleSetWithTerminateScheduledEventsEnabled() {
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
					AdminPassword:      to.Ptr("{your-password}"),
					AdminUsername:      to.Ptr("{your-username}"),
					ComputerNamePrefix: to.Ptr("{vmss-name}"),
				},
				ScheduledEventsProfile: &armcompute.ScheduledEventsProfile{
					TerminateNotificationProfile: &armcompute.TerminateNotificationProfile{
						Enable:           to.Ptr(true),
						NotBeforeTimeout: to.Ptr("PT5M"),
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
