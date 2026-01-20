package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ccea12be6cd5e64743499d46f7a9c8b60df52db1/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2025-04-01/examples/virtualMachineScaleSetExamples/VirtualMachineScaleSet_Create_WithPriorityMixPolicy.json
func ExampleVirtualMachineScaleSetsClient_BeginCreateOrUpdate_createAScaleSetWithPriorityMixPolicy() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualMachineScaleSetsClient().BeginCreateOrUpdate(ctx, "myResourceGroup", "{vmss-name}", armcompute.VirtualMachineScaleSet{
		Location: to.Ptr("westus"),
		Properties: &armcompute.VirtualMachineScaleSetProperties{
			OrchestrationMode:        to.Ptr(armcompute.OrchestrationModeFlexible),
			PlatformFaultDomainCount: to.Ptr[int32](1),
			PriorityMixPolicy: &armcompute.PriorityMixPolicy{
				BaseRegularPriorityCount:           to.Ptr[int32](10),
				RegularPriorityPercentageAboveBase: to.Ptr[int32](50),
			},
			VirtualMachineProfile: &armcompute.VirtualMachineScaleSetVMProfile{
				NetworkProfile: &armcompute.VirtualMachineScaleSetNetworkProfile{
					NetworkAPIVersion: to.Ptr(armcompute.NetworkAPIVersionTwoThousandTwenty1101),
					NetworkInterfaceConfigurations: []*armcompute.VirtualMachineScaleSetNetworkConfiguration{
						{
							Name: to.Ptr("{vmss-name}"),
							Properties: &armcompute.VirtualMachineScaleSetNetworkConfigurationProperties{
								EnableAcceleratedNetworking: to.Ptr(false),
								EnableIPForwarding:          to.Ptr(true),
								IPConfigurations: []*armcompute.VirtualMachineScaleSetIPConfiguration{
									{
										Name: to.Ptr("{vmss-name}"),
										Properties: &armcompute.VirtualMachineScaleSetIPConfigurationProperties{
											ApplicationGatewayBackendAddressPools: []*armcompute.SubResource{},
											LoadBalancerBackendAddressPools:       []*armcompute.SubResource{},
											Primary:                               to.Ptr(true),
											PublicIPAddressConfiguration: &armcompute.VirtualMachineScaleSetPublicIPAddressConfiguration{
												Name: to.Ptr("{vmss-name}"),
												Properties: &armcompute.VirtualMachineScaleSetPublicIPAddressConfigurationProperties{
													IdleTimeoutInMinutes: to.Ptr[int32](15),
												},
											},
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
				},
				Priority: to.Ptr(armcompute.VirtualMachinePriorityTypesSpot),
				StorageProfile: &armcompute.VirtualMachineScaleSetStorageProfile{
					ImageReference: &armcompute.ImageReference{
						Offer:     to.Ptr("0001-com-ubuntu-server-focal"),
						Publisher: to.Ptr("Canonical"),
						SKU:       to.Ptr("20_04-lts-gen2"),
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
			Name:     to.Ptr("Standard_A8m_v2"),
			Capacity: to.Ptr[int64](2),
			Tier:     to.Ptr("Standard"),
		},
	}, &armcompute.VirtualMachineScaleSetsClientBeginCreateOrUpdateOptions{IfMatch: nil,
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
	// res.VirtualMachineScaleSet = armcompute.VirtualMachineScaleSet{
	// 	Name: to.Ptr("{vmss-name}"),
	// 	Type: to.Ptr("Microsoft.Compute/virtualMachineScaleSets"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachineScaleSets/{vmss-name}"),
	// 	Location: to.Ptr("westus"),
	// 	Properties: &armcompute.VirtualMachineScaleSetProperties{
	// 		OrchestrationMode: to.Ptr(armcompute.OrchestrationModeFlexible),
	// 		PlatformFaultDomainCount: to.Ptr[int32](1),
	// 		PriorityMixPolicy: &armcompute.PriorityMixPolicy{
	// 			BaseRegularPriorityCount: to.Ptr[int32](10),
	// 			RegularPriorityPercentageAboveBase: to.Ptr[int32](50),
	// 		},
	// 		VirtualMachineProfile: &armcompute.VirtualMachineScaleSetVMProfile{
	// 			NetworkProfile: &armcompute.VirtualMachineScaleSetNetworkProfile{
	// 				NetworkAPIVersion: to.Ptr(armcompute.NetworkAPIVersionTwoThousandTwenty1101),
	// 				NetworkInterfaceConfigurations: []*armcompute.VirtualMachineScaleSetNetworkConfiguration{
	// 					{
	// 						Name: to.Ptr("{vmss-name}"),
	// 						Properties: &armcompute.VirtualMachineScaleSetNetworkConfigurationProperties{
	// 							EnableAcceleratedNetworking: to.Ptr(false),
	// 							EnableIPForwarding: to.Ptr(true),
	// 							IPConfigurations: []*armcompute.VirtualMachineScaleSetIPConfiguration{
	// 								{
	// 									Name: to.Ptr("{vmss-name}"),
	// 									Properties: &armcompute.VirtualMachineScaleSetIPConfigurationProperties{
	// 										ApplicationGatewayBackendAddressPools: []*armcompute.SubResource{
	// 										},
	// 										LoadBalancerBackendAddressPools: []*armcompute.SubResource{
	// 										},
	// 										Primary: to.Ptr(true),
	// 										PublicIPAddressConfiguration: &armcompute.VirtualMachineScaleSetPublicIPAddressConfiguration{
	// 											Name: to.Ptr("{vmss-name}"),
	// 											Properties: &armcompute.VirtualMachineScaleSetPublicIPAddressConfigurationProperties{
	// 												IdleTimeoutInMinutes: to.Ptr[int32](15),
	// 											},
	// 										},
	// 										Subnet: &armcompute.APIEntityReference{
	// 											ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/{existing-virtual-network-name}/subnets/{existing-subnet-name}"),
	// 										},
	// 									},
	// 							}},
	// 							Primary: to.Ptr(true),
	// 						},
	// 				}},
	// 			},
	// 			OSProfile: &armcompute.VirtualMachineScaleSetOSProfile{
	// 				AdminUsername: to.Ptr("{your-username}"),
	// 				ComputerNamePrefix: to.Ptr("{vmss-name}"),
	// 			},
	// 			Priority: to.Ptr(armcompute.VirtualMachinePriorityTypesSpot),
	// 			StorageProfile: &armcompute.VirtualMachineScaleSetStorageProfile{
	// 				ImageReference: &armcompute.ImageReference{
	// 					Offer: to.Ptr("0001-com-ubuntu-server-focal"),
	// 					Publisher: to.Ptr("Canonical"),
	// 					SKU: to.Ptr("20_04-lts-gen2"),
	// 					Version: to.Ptr("latest"),
	// 				},
	// 				OSDisk: &armcompute.VirtualMachineScaleSetOSDisk{
	// 					Caching: to.Ptr(armcompute.CachingTypesReadWrite),
	// 					CreateOption: to.Ptr(armcompute.DiskCreateOptionTypesFromImage),
	// 					ManagedDisk: &armcompute.VirtualMachineScaleSetManagedDiskParameters{
	// 						StorageAccountType: to.Ptr(armcompute.StorageAccountTypesStandardLRS),
	// 					},
	// 				},
	// 			},
	// 		},
	// 	},
	// 	SKU: &armcompute.SKU{
	// 		Name: to.Ptr("Standard_A8m_v2"),
	// 		Capacity: to.Ptr[int64](2),
	// 		Tier: to.Ptr("Standard"),
	// 	},
	// }
}
