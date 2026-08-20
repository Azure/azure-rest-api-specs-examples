package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v8"
)

// Generated from example definition: 2026-04-01/virtualMachineScaleSetExamples/VirtualMachineScaleSet_Create_WithSpotPlusPriorityFlex.json
func ExampleVirtualMachineScaleSetsClient_BeginCreateOrUpdate_createAScaleSetWithSpotPlusPriorityFlexible() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualMachineScaleSetsClient().BeginCreateOrUpdate(ctx, "myResourceGroup", "{vmss-name}", armcompute.VirtualMachineScaleSet{
		SKU: &armcompute.SKU{
			Capacity: to.Ptr[int64](10),
			Name:     to.Ptr("Mix"),
		},
		Location: to.Ptr("westus"),
		Properties: &armcompute.VirtualMachineScaleSetProperties{
			OrchestrationMode:        to.Ptr(armcompute.OrchestrationModeFlexible),
			PlatformFaultDomainCount: to.Ptr[int32](1),
			SinglePlacementGroup:     to.Ptr(false),
			VirtualMachineProfile: &armcompute.VirtualMachineScaleSetVMProfile{
				StorageProfile: &armcompute.VirtualMachineScaleSetStorageProfile{
					ImageReference: &armcompute.ImageReference{
						Publisher: to.Ptr("Canonical"),
						Offer:     to.Ptr("0001-com-ubuntu-server-focal"),
						SKU:       to.Ptr("20_04-lts-gen2"),
						Version:   to.Ptr("latest"),
					},
					OSDisk: &armcompute.VirtualMachineScaleSetOSDisk{
						CreateOption: to.Ptr(armcompute.DiskCreateOptionTypesFromImage),
						Caching:      to.Ptr(armcompute.CachingTypesReadWrite),
						ManagedDisk: &armcompute.VirtualMachineScaleSetManagedDiskParameters{
							StorageAccountType: to.Ptr(armcompute.StorageAccountTypesStandardLRS),
						},
					},
				},
				OSProfile: &armcompute.VirtualMachineScaleSetOSProfile{
					ComputerNamePrefix: to.Ptr("{vmss-name}"),
					AdminUsername:      to.Ptr("{your-username}"),
				},
				NetworkProfile: &armcompute.VirtualMachineScaleSetNetworkProfile{
					NetworkInterfaceConfigurations: []*armcompute.VirtualMachineScaleSetNetworkConfiguration{
						{
							Name: to.Ptr("{vmss-name}"),
							Properties: &armcompute.VirtualMachineScaleSetNetworkConfigurationProperties{
								Primary:                     to.Ptr(true),
								EnableIPForwarding:          to.Ptr(true),
								EnableAcceleratedNetworking: to.Ptr(false),
								IPConfigurations: []*armcompute.VirtualMachineScaleSetIPConfiguration{
									{
										Name: to.Ptr("{vmss-name}"),
										Properties: &armcompute.VirtualMachineScaleSetIPConfigurationProperties{
											Subnet: &armcompute.APIEntityReference{
												ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/{existing-virtual-network-name}/subnets/{existing-subnet-name}"),
											},
											Primary:                               to.Ptr(true),
											ApplicationGatewayBackendAddressPools: []*armcompute.SubResource{},
											LoadBalancerBackendAddressPools:       []*armcompute.SubResource{},
											PublicIPAddressConfiguration: &armcompute.VirtualMachineScaleSetPublicIPAddressConfiguration{
												Name: to.Ptr("{vmss-name}"),
												Properties: &armcompute.VirtualMachineScaleSetPublicIPAddressConfigurationProperties{
													IdleTimeoutInMinutes: to.Ptr[int32](15),
												},
											},
										},
									},
								},
							},
						},
					},
					NetworkAPIVersion: to.Ptr(armcompute.NetworkAPIVersionTwoThousandTwenty1101),
				},
				Priority:       to.Ptr(armcompute.VirtualMachinePriorityTypesSpotPlus),
				EvictionPolicy: to.Ptr(armcompute.VirtualMachineEvictionPolicyTypesDeallocate),
				BillingProfile: &armcompute.BillingProfile{
					MaxPrice: to.Ptr[float64](-1),
				},
			},
			PriorityMixPolicy: &armcompute.PriorityMixPolicy{
				BaseRegularPriorityCount:           to.Ptr[int32](4),
				RegularPriorityPercentageAboveBase: to.Ptr[int32](50),
			},
			SKUProfile: &armcompute.SKUProfile{
				VMSizes: []*armcompute.SKUProfileVMSize{
					{
						Name: to.Ptr("Standard_D8s_v5"),
					},
					{
						Name: to.Ptr("Standard_E16s_v5"),
					},
					{
						Name: to.Ptr("Standard_D2s_v5"),
					},
				},
				AllocationStrategy: to.Ptr(armcompute.AllocationStrategyCapacityOptimized),
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
	// res = armcompute.VirtualMachineScaleSetsClientCreateOrUpdateResponse{
	// 	VirtualMachineScaleSet: armcompute.VirtualMachineScaleSet{
	// 		SKU: &armcompute.SKU{
	// 			Capacity: to.Ptr[int64](10),
	// 			Name: to.Ptr("Mix"),
	// 		},
	// 		Name: to.Ptr("{vmss-name}"),
	// 		Properties: &armcompute.VirtualMachineScaleSetProperties{
	// 			OrchestrationMode: to.Ptr(armcompute.OrchestrationModeFlexible),
	// 			PlatformFaultDomainCount: to.Ptr[int32](1),
	// 			SinglePlacementGroup: to.Ptr(false),
	// 			UniqueID: to.Ptr("c2d3e4f5-a6b7-8901-cdef-012345678901"),
	// 			VirtualMachineProfile: &armcompute.VirtualMachineScaleSetVMProfile{
	// 				StorageProfile: &armcompute.VirtualMachineScaleSetStorageProfile{
	// 					ImageReference: &armcompute.ImageReference{
	// 						Publisher: to.Ptr("Canonical"),
	// 						Offer: to.Ptr("0001-com-ubuntu-server-focal"),
	// 						SKU: to.Ptr("20_04-lts-gen2"),
	// 						Version: to.Ptr("latest"),
	// 					},
	// 					OSDisk: &armcompute.VirtualMachineScaleSetOSDisk{
	// 						CreateOption: to.Ptr(armcompute.DiskCreateOptionTypesFromImage),
	// 						Caching: to.Ptr(armcompute.CachingTypesReadWrite),
	// 						ManagedDisk: &armcompute.VirtualMachineScaleSetManagedDiskParameters{
	// 							StorageAccountType: to.Ptr(armcompute.StorageAccountTypesStandardLRS),
	// 						},
	// 					},
	// 				},
	// 				OSProfile: &armcompute.VirtualMachineScaleSetOSProfile{
	// 					ComputerNamePrefix: to.Ptr("{vmss-name}"),
	// 					AdminUsername: to.Ptr("{your-username}"),
	// 				},
	// 				NetworkProfile: &armcompute.VirtualMachineScaleSetNetworkProfile{
	// 					NetworkInterfaceConfigurations: []*armcompute.VirtualMachineScaleSetNetworkConfiguration{
	// 						{
	// 							Name: to.Ptr("{vmss-name}"),
	// 							Properties: &armcompute.VirtualMachineScaleSetNetworkConfigurationProperties{
	// 								Primary: to.Ptr(true),
	// 								EnableIPForwarding: to.Ptr(true),
	// 								EnableAcceleratedNetworking: to.Ptr(false),
	// 								IPConfigurations: []*armcompute.VirtualMachineScaleSetIPConfiguration{
	// 									{
	// 										Name: to.Ptr("{vmss-name}"),
	// 										Properties: &armcompute.VirtualMachineScaleSetIPConfigurationProperties{
	// 											Subnet: &armcompute.APIEntityReference{
	// 												ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/{existing-virtual-network-name}/subnets/{existing-subnet-name}"),
	// 											},
	// 											Primary: to.Ptr(true),
	// 											ApplicationGatewayBackendAddressPools: []*armcompute.SubResource{
	// 											},
	// 											LoadBalancerBackendAddressPools: []*armcompute.SubResource{
	// 											},
	// 											PublicIPAddressConfiguration: &armcompute.VirtualMachineScaleSetPublicIPAddressConfiguration{
	// 												Name: to.Ptr("{vmss-name}"),
	// 												Properties: &armcompute.VirtualMachineScaleSetPublicIPAddressConfigurationProperties{
	// 													IdleTimeoutInMinutes: to.Ptr[int32](15),
	// 												},
	// 											},
	// 										},
	// 									},
	// 								},
	// 							},
	// 						},
	// 					},
	// 					NetworkAPIVersion: to.Ptr(armcompute.NetworkAPIVersionTwoThousandTwenty1101),
	// 				},
	// 				Priority: to.Ptr(armcompute.VirtualMachinePriorityTypesSpotPlus),
	// 				EvictionPolicy: to.Ptr(armcompute.VirtualMachineEvictionPolicyTypesDeallocate),
	// 				BillingProfile: &armcompute.BillingProfile{
	// 					MaxPrice: to.Ptr[float64](-1),
	// 				},
	// 			},
	// 			PriorityMixPolicy: &armcompute.PriorityMixPolicy{
	// 				BaseRegularPriorityCount: to.Ptr[int32](4),
	// 				RegularPriorityPercentageAboveBase: to.Ptr[int32](50),
	// 			},
	// 			SKUProfile: &armcompute.SKUProfile{
	// 				VMSizes: []*armcompute.SKUProfileVMSize{
	// 					{
	// 						Name: to.Ptr("Standard_D8s_v5"),
	// 					},
	// 					{
	// 						Name: to.Ptr("Standard_E16s_v5"),
	// 					},
	// 					{
	// 						Name: to.Ptr("Standard_D2s_v5"),
	// 					},
	// 				},
	// 				AllocationStrategy: to.Ptr(armcompute.AllocationStrategyCapacityOptimized),
	// 			},
	// 			ProvisioningState: to.Ptr("Creating"),
	// 		},
	// 		Location: to.Ptr("westus"),
	// 		Type: to.Ptr("Microsoft.Compute/virtualMachineScaleSets"),
	// 		ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachineScaleSets/{vmss-name}"),
	// 	},
	// }
}
