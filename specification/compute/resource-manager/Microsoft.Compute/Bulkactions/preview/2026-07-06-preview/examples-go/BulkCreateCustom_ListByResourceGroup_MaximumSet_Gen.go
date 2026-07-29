package armbulkactions_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armbulkactions"
)

// Generated from example definition: 2026-07-06-preview/BulkCreateCustom_ListByResourceGroup_MaximumSet_Gen.json
func ExampleBulkCreateCustomClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbulkactions.NewClientFactory("1FBA3C66-5C9C-4391-B72F-9F52735FC9F2", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewBulkCreateCustomClient().NewListByResourceGroupPager("rgBulkactions", "eastus", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page = armbulkactions.BulkCreateCustomClientListByResourceGroupResponse{
		// 	BulkCreateCustomListResult: armbulkactions.BulkCreateCustomListResult{
		// 		Value: []*armbulkactions.LocationBasedBulkCreateCustom{
		// 			{
		// 				Properties: &armbulkactions.BulkCreateCustomProperties{
		// 					CreatedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-07-08T18:20:00Z"); return t}()),
		// 					ProvisioningState: to.Ptr(armbulkactions.ProvisioningStateSucceeded),
		// 					Capacity: to.Ptr[int32](10),
		// 					CapacityType: to.Ptr(armbulkactions.CapacityTypeVM),
		// 					PriorityProfile: &armbulkactions.BulkCreateCustomPriorityProfile{
		// 						Type: to.Ptr(armbulkactions.PriorityTypeSpot),
		// 						MaxPricePerVM: to.Ptr[float32](0.2),
		// 						EvictionPolicy: to.Ptr(armbulkactions.EvictionPolicyDelete),
		// 						AllocationStrategy: to.Ptr(armbulkactions.BulkCreateCustomAllocationStrategyLowestPrice),
		// 					},
		// 					VMSizesProfile: []*armbulkactions.BulkCreateCustomVMSizeProfile{
		// 						{
		// 							Name: to.Ptr("Standard_D2s_v5"),
		// 							Rank: to.Ptr[int32](1),
		// 						},
		// 						{
		// 							Name: to.Ptr("Standard_D4s_v5"),
		// 							Rank: to.Ptr[int32](2),
		// 						},
		// 					},
		// 					ComputeProfile: &armbulkactions.ComputeProfile{
		// 						VirtualMachineProfile: &armbulkactions.BulkactionVMProperties{
		// 							StorageProfile: &armbulkactions.StorageProfile{
		// 								ImageReference: &armbulkactions.ImageReference{
		// 									Publisher: to.Ptr("Canonical"),
		// 									Offer: to.Ptr("0001-com-ubuntu-server-jammy"),
		// 									SKU: to.Ptr("22_04-lts-gen2"),
		// 									Version: to.Ptr("latest"),
		// 								},
		// 								OSDisk: &armbulkactions.OSDisk{
		// 									OSType: to.Ptr(armbulkactions.OperatingSystemTypesLinux),
		// 									Caching: to.Ptr(armbulkactions.CachingTypesReadWrite),
		// 									CreateOption: to.Ptr(armbulkactions.DiskCreateOptionTypesFromImage),
		// 									ManagedDisk: &armbulkactions.ManagedDiskParametersContent{
		// 										StorageAccountType: to.Ptr(armbulkactions.StorageAccountTypesPremiumLRS),
		// 									},
		// 									DeleteOption: to.Ptr(armbulkactions.DiskDeleteOptionTypesDelete),
		// 								},
		// 							},
		// 							OSProfile: &armbulkactions.OSProfile{
		// 								ComputerName: to.Ptr("bulkvm"),
		// 								AdminUsername: to.Ptr("azureuser"),
		// 								LinuxConfiguration: &armbulkactions.LinuxConfiguration{
		// 									DisablePasswordAuthentication: to.Ptr(true),
		// 									SSH: &armbulkactions.SSHConfiguration{
		// 										PublicKeys: []*armbulkactions.SSHPublicKey{
		// 											{
		// 												Path: to.Ptr("/home/azureuser/.ssh/authorized_keys"),
		// 												KeyData: to.Ptr("ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDExampleKeyReplaceWithYourOwn azureuser@bulklaunch"),
		// 											},
		// 										},
		// 									},
		// 								},
		// 							},
		// 							NetworkProfile: &armbulkactions.NetworkProfile{
		// 								NetworkAPIVersion: to.Ptr(armbulkactions.NetworkAPIVersion20201101),
		// 								NetworkInterfaceConfigurations: []*armbulkactions.VirtualMachineNetworkInterfaceConfiguration{
		// 									{
		// 										Name: to.Ptr("bulkvm-nic"),
		// 										Properties: &armbulkactions.VirtualMachineNetworkInterfaceConfigurationProperties{
		// 											Primary: to.Ptr(true),
		// 											DeleteOption: to.Ptr(armbulkactions.DeleteOptionsDelete),
		// 											IPConfigurations: []*armbulkactions.VirtualMachineNetworkInterfaceIPConfiguration{
		// 												{
		// 													Name: to.Ptr("bulkvm-ipconfig"),
		// 													Properties: &armbulkactions.VirtualMachineNetworkInterfaceIPConfigurationProperties{
		// 														Primary: to.Ptr(true),
		// 														Subnet: &armbulkactions.SubResource{
		// 															ID: to.Ptr("/subscriptions/1FBA3C66-5C9C-4391-B72F-9F52735FC9F2/resourceGroups/rgBulkactions/providers/Microsoft.Network/virtualNetworks/bulkvnet/subnets/default"),
		// 														},
		// 													},
		// 												},
		// 											},
		// 										},
		// 									},
		// 								},
		// 							},
		// 						},
		// 						ComputeAPIVersion: to.Ptr("2024-11-01"),
		// 					},
		// 					ZoneAllocationPolicy: &armbulkactions.BulkCreateCustomZoneAllocationPolicy{
		// 						DistributionStrategy: to.Ptr(armbulkactions.BulkCreateCustomDistributionStrategyBestEffortBalanced),
		// 						ZonePreferences: []*armbulkactions.ZonePreference{
		// 							{
		// 								Zone: to.Ptr("1"),
		// 								Rank: to.Ptr[int32](1),
		// 							},
		// 							{
		// 								Zone: to.Ptr("2"),
		// 								Rank: to.Ptr[int32](2),
		// 							},
		// 						},
		// 					},
		// 					ExecutionParameters: &armbulkactions.ExecutionParameters{
		// 						RetryPolicy: &armbulkactions.RetryPolicy{
		// 							RetryWindowInMinutes: to.Ptr[int32](30),
		// 							OnFailureAction: to.Ptr(armbulkactions.ResourceOperationTypeDelete),
		// 						},
		// 					},
		// 				},
		// 				Zones: []*string{
		// 					to.Ptr("1"),
		// 					to.Ptr("2"),
		// 				},
		// 				Tags: map[string]*string{
		// 					"workload": to.Ptr("batch-render"),
		// 					"env": to.Ptr("prod"),
		// 				},
		// 				Identity: &armbulkactions.ManagedServiceIdentity{
		// 					Type: to.Ptr(armbulkactions.ManagedServiceIdentityTypeSystemAssigned),
		// 					PrincipalID: to.Ptr("aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee"),
		// 					TenantID: to.Ptr("11111111-2222-3333-4444-555555555555"),
		// 				},
		// 				ID: to.Ptr("/subscriptions/1FBA3C66-5C9C-4391-B72F-9F52735FC9F2/resourceGroups/rgBulkactions/providers/Microsoft.Compute/locations/eastus/bulkCreateCustom/85c374f7-9857-4fd7-9267-81019219c362"),
		// 				Name: to.Ptr("85c374f7-9857-4fd7-9267-81019219c362"),
		// 				Type: to.Ptr("Microsoft.Compute/locations/bulkCreateCustom"),
		// 				SystemData: &armbulkactions.SystemData{
		// 					CreatedBy: to.Ptr("user@contoso.com"),
		// 					CreatedByType: to.Ptr(armbulkactions.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-07-08T18:20:00Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("user@contoso.com"),
		// 					LastModifiedByType: to.Ptr(armbulkactions.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-07-08T18:25:00Z"); return t}()),
		// 				},
		// 			},
		// 			{
		// 				Properties: &armbulkactions.BulkCreateCustomProperties{
		// 					CreatedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-07-08T18:20:00Z"); return t}()),
		// 					ProvisioningState: to.Ptr(armbulkactions.ProvisioningStateCreating),
		// 					Capacity: to.Ptr[int32](10),
		// 					CapacityType: to.Ptr(armbulkactions.CapacityTypeVM),
		// 					PriorityProfile: &armbulkactions.BulkCreateCustomPriorityProfile{
		// 						Type: to.Ptr(armbulkactions.PriorityTypeSpot),
		// 						MaxPricePerVM: to.Ptr[float32](0.2),
		// 						EvictionPolicy: to.Ptr(armbulkactions.EvictionPolicyDelete),
		// 						AllocationStrategy: to.Ptr(armbulkactions.BulkCreateCustomAllocationStrategyLowestPrice),
		// 					},
		// 					VMSizesProfile: []*armbulkactions.BulkCreateCustomVMSizeProfile{
		// 						{
		// 							Name: to.Ptr("Standard_D2s_v5"),
		// 							Rank: to.Ptr[int32](1),
		// 						},
		// 						{
		// 							Name: to.Ptr("Standard_D4s_v5"),
		// 							Rank: to.Ptr[int32](2),
		// 						},
		// 					},
		// 					ComputeProfile: &armbulkactions.ComputeProfile{
		// 						VirtualMachineProfile: &armbulkactions.BulkactionVMProperties{
		// 							StorageProfile: &armbulkactions.StorageProfile{
		// 								ImageReference: &armbulkactions.ImageReference{
		// 									Publisher: to.Ptr("Canonical"),
		// 									Offer: to.Ptr("0001-com-ubuntu-server-jammy"),
		// 									SKU: to.Ptr("22_04-lts-gen2"),
		// 									Version: to.Ptr("latest"),
		// 								},
		// 								OSDisk: &armbulkactions.OSDisk{
		// 									OSType: to.Ptr(armbulkactions.OperatingSystemTypesLinux),
		// 									Caching: to.Ptr(armbulkactions.CachingTypesReadWrite),
		// 									CreateOption: to.Ptr(armbulkactions.DiskCreateOptionTypesFromImage),
		// 									ManagedDisk: &armbulkactions.ManagedDiskParametersContent{
		// 										StorageAccountType: to.Ptr(armbulkactions.StorageAccountTypesPremiumLRS),
		// 									},
		// 									DeleteOption: to.Ptr(armbulkactions.DiskDeleteOptionTypesDelete),
		// 								},
		// 							},
		// 							OSProfile: &armbulkactions.OSProfile{
		// 								ComputerName: to.Ptr("bulkvm"),
		// 								AdminUsername: to.Ptr("azureuser"),
		// 								LinuxConfiguration: &armbulkactions.LinuxConfiguration{
		// 									DisablePasswordAuthentication: to.Ptr(true),
		// 									SSH: &armbulkactions.SSHConfiguration{
		// 										PublicKeys: []*armbulkactions.SSHPublicKey{
		// 											{
		// 												Path: to.Ptr("/home/azureuser/.ssh/authorized_keys"),
		// 												KeyData: to.Ptr("ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDExampleKeyReplaceWithYourOwn azureuser@bulklaunch"),
		// 											},
		// 										},
		// 									},
		// 								},
		// 							},
		// 							NetworkProfile: &armbulkactions.NetworkProfile{
		// 								NetworkAPIVersion: to.Ptr(armbulkactions.NetworkAPIVersion20201101),
		// 								NetworkInterfaceConfigurations: []*armbulkactions.VirtualMachineNetworkInterfaceConfiguration{
		// 									{
		// 										Name: to.Ptr("bulkvm-nic"),
		// 										Properties: &armbulkactions.VirtualMachineNetworkInterfaceConfigurationProperties{
		// 											Primary: to.Ptr(true),
		// 											DeleteOption: to.Ptr(armbulkactions.DeleteOptionsDelete),
		// 											IPConfigurations: []*armbulkactions.VirtualMachineNetworkInterfaceIPConfiguration{
		// 												{
		// 													Name: to.Ptr("bulkvm-ipconfig"),
		// 													Properties: &armbulkactions.VirtualMachineNetworkInterfaceIPConfigurationProperties{
		// 														Primary: to.Ptr(true),
		// 														Subnet: &armbulkactions.SubResource{
		// 															ID: to.Ptr("/subscriptions/1FBA3C66-5C9C-4391-B72F-9F52735FC9F2/resourceGroups/rgBulkactions/providers/Microsoft.Network/virtualNetworks/bulkvnet/subnets/default"),
		// 														},
		// 													},
		// 												},
		// 											},
		// 										},
		// 									},
		// 								},
		// 							},
		// 						},
		// 						ComputeAPIVersion: to.Ptr("2024-11-01"),
		// 					},
		// 					ZoneAllocationPolicy: &armbulkactions.BulkCreateCustomZoneAllocationPolicy{
		// 						DistributionStrategy: to.Ptr(armbulkactions.BulkCreateCustomDistributionStrategyBestEffortBalanced),
		// 						ZonePreferences: []*armbulkactions.ZonePreference{
		// 							{
		// 								Zone: to.Ptr("1"),
		// 								Rank: to.Ptr[int32](1),
		// 							},
		// 							{
		// 								Zone: to.Ptr("2"),
		// 								Rank: to.Ptr[int32](2),
		// 							},
		// 						},
		// 					},
		// 					ExecutionParameters: &armbulkactions.ExecutionParameters{
		// 						RetryPolicy: &armbulkactions.RetryPolicy{
		// 							RetryWindowInMinutes: to.Ptr[int32](30),
		// 							OnFailureAction: to.Ptr(armbulkactions.ResourceOperationTypeDelete),
		// 						},
		// 					},
		// 				},
		// 				Zones: []*string{
		// 					to.Ptr("1"),
		// 					to.Ptr("2"),
		// 				},
		// 				Tags: map[string]*string{
		// 					"workload": to.Ptr("batch-render"),
		// 					"env": to.Ptr("prod"),
		// 				},
		// 				Identity: &armbulkactions.ManagedServiceIdentity{
		// 					Type: to.Ptr(armbulkactions.ManagedServiceIdentityTypeSystemAssigned),
		// 					PrincipalID: to.Ptr("aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee"),
		// 					TenantID: to.Ptr("11111111-2222-3333-4444-555555555555"),
		// 				},
		// 				ID: to.Ptr("/subscriptions/1FBA3C66-5C9C-4391-B72F-9F52735FC9F2/resourceGroups/rgBulkactions/providers/Microsoft.Compute/locations/eastus/bulkCreateCustom/b1d0f6a2-3c4e-4f5a-8b7c-9d0e1f2a3b4c"),
		// 				Name: to.Ptr("b1d0f6a2-3c4e-4f5a-8b7c-9d0e1f2a3b4c"),
		// 				Type: to.Ptr("Microsoft.Compute/locations/bulkCreateCustom"),
		// 				SystemData: &armbulkactions.SystemData{
		// 					CreatedBy: to.Ptr("user@contoso.com"),
		// 					CreatedByType: to.Ptr(armbulkactions.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-07-08T18:20:00Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("user@contoso.com"),
		// 					LastModifiedByType: to.Ptr(armbulkactions.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-07-08T18:25:00Z"); return t}()),
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://management.azure.com/subscriptions/1FBA3C66-5C9C-4391-B72F-9F52735FC9F2/resourceGroups/rgBulkactions/providers/Microsoft.Compute/locations/eastus/bulkCreateCustom?api-version=2026-07-06-preview&$skiptoken=page2"),
		// 	},
		// }
	}
}
