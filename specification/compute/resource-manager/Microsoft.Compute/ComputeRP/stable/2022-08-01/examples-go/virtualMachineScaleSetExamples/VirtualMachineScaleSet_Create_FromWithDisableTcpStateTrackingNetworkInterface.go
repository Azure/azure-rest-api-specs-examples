package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-08-01/examples/virtualMachineScaleSetExamples/VirtualMachineScaleSet_Create_FromWithDisableTcpStateTrackingNetworkInterface.json
func ExampleVirtualMachineScaleSetsClient_BeginCreateOrUpdate_createAScaleSetWhereNicConfigHasDisableTcpStateTrackingProperty() {
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
							Name: to.Ptr("{nicConfig1-name}"),
							Properties: &armcompute.VirtualMachineScaleSetNetworkConfigurationProperties{
								DisableTCPStateTracking:     to.Ptr(true),
								EnableAcceleratedNetworking: to.Ptr(true),
								EnableIPForwarding:          to.Ptr(true),
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
						},
						{
							Name: to.Ptr("{nicConfig2-name}"),
							Properties: &armcompute.VirtualMachineScaleSetNetworkConfigurationProperties{
								DisableTCPStateTracking:     to.Ptr(false),
								EnableAcceleratedNetworking: to.Ptr(false),
								EnableIPForwarding:          to.Ptr(false),
								IPConfigurations: []*armcompute.VirtualMachineScaleSetIPConfiguration{
									{
										Name: to.Ptr("{nicConfig2-name}"),
										Properties: &armcompute.VirtualMachineScaleSetIPConfigurationProperties{
											Primary:                 to.Ptr(true),
											PrivateIPAddressVersion: to.Ptr(armcompute.IPVersionIPv4),
											Subnet: &armcompute.APIEntityReference{
												ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/{existing-virtual-network-name}/subnets/{existing-fpga-subnet-name2}"),
											},
										},
									}},
								Primary: to.Ptr(false),
							},
						}},
				},
				OSProfile: &armcompute.VirtualMachineScaleSetOSProfile{
					AdminPassword:      to.Ptr("{your-password}"),
					AdminUsername:      to.Ptr("{your-username}"),
					ComputerNamePrefix: to.Ptr("{vmss-name}"),
				},
				StorageProfile: &armcompute.VirtualMachineScaleSetStorageProfile{
					ImageReference: &armcompute.ImageReference{
						ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/images/{existing-custom-image-name}"),
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
