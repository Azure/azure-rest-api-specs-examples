package armcomputeschedule_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/computeschedule/armcomputeschedule/v2"
)

// Generated from example definition: 2026-04-15-preview/ScheduledActions_VirtualMachinesExecuteCreateFlex_MaximumSet_Gen.json
func ExampleScheduledActionsClient_VirtualMachinesExecuteCreateFlex_scheduledActionsVirtualMachinesExecuteCreateFlexMaximumSet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcomputeschedule.NewClientFactory("732116BD-AF31-4E74-9283-B387C44B4A44", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewScheduledActionsClient().VirtualMachinesExecuteCreateFlex(ctx, "eastus2", armcomputeschedule.ExecuteCreateFlexRequest{
		ResourceConfigParameters: &armcomputeschedule.ResourceProvisionFlexPayload{
			VirtualMachineBaseProfile: &armcomputeschedule.BulkVMConfiguration{
				Name:              to.Ptr("baseFlexVmConfig"),
				ComputeAPIVersion: to.Ptr("2024-07-01"),
				ResourceGroupName: to.Ptr("myResourceGroup"),
				Zones: []*string{
					to.Ptr("1"),
				},
				Identity: &armcomputeschedule.VirtualMachineIdentity{
					Type: to.Ptr(armcomputeschedule.ResourceIdentityTypeSystemAssigned),
				},
				Tags: map[string]*string{
					"environment": to.Ptr("production"),
					"department":  to.Ptr("engineering"),
				},
				Properties: &armcomputeschedule.BulkActionVMProperties{
					StorageProfile: &armcomputeschedule.StorageProfile{
						ImageReference: &armcomputeschedule.ImageReference{
							Publisher: to.Ptr("Canonical"),
							Offer:     to.Ptr("0001-com-ubuntu-server-jammy"),
							SKU:       to.Ptr("22_04-lts-gen2"),
							Version:   to.Ptr("latest"),
						},
						OSDisk: &armcomputeschedule.OSDisk{
							OSType:       to.Ptr(armcomputeschedule.OperatingSystemTypesLinux),
							Name:         to.Ptr("myOsDisk"),
							Caching:      to.Ptr(armcomputeschedule.CachingTypesReadWrite),
							CreateOption: to.Ptr(armcomputeschedule.DiskCreateOptionTypesFromImage),
							DiskSizeGB:   to.Ptr[int32](128),
							ManagedDisk: &armcomputeschedule.ManagedDiskParameters{
								StorageAccountType: to.Ptr(armcomputeschedule.StorageAccountTypesPremiumLRS),
							},
							DeleteOption: to.Ptr(armcomputeschedule.DiskDeleteOptionTypesDelete),
						},
						DataDisks: []*armcomputeschedule.DataDisk{
							{
								Lun:          to.Ptr[int32](0),
								Name:         to.Ptr("myDataDisk-0"),
								Caching:      to.Ptr(armcomputeschedule.CachingTypesReadOnly),
								CreateOption: to.Ptr(armcomputeschedule.DiskCreateOptionTypesEmpty),
								DiskSizeGB:   to.Ptr[int32](256),
								ManagedDisk: &armcomputeschedule.ManagedDiskParameters{
									StorageAccountType: to.Ptr(armcomputeschedule.StorageAccountTypesPremiumLRS),
								},
								DeleteOption: to.Ptr(armcomputeschedule.DiskDeleteOptionTypesDelete),
							},
						},
						DiskControllerType: to.Ptr(armcomputeschedule.DiskControllerTypesSCSI),
					},
					AdditionalCapabilities: &armcomputeschedule.AdditionalCapabilities{
						UltraSSDEnabled:    to.Ptr(false),
						HibernationEnabled: to.Ptr(false),
					},
					OSProfile: &armcomputeschedule.OSProfile{
						ComputerName:  to.Ptr("myFlexVM"),
						AdminUsername: to.Ptr("azureuser"),
						LinuxConfiguration: &armcomputeschedule.LinuxConfiguration{
							DisablePasswordAuthentication: to.Ptr(true),
							SSH: &armcomputeschedule.SSHConfiguration{
								PublicKeys: []*armcomputeschedule.SSHPublicKey{
									{
										Path:    to.Ptr("/home/azureuser/.ssh/authorized_keys"),
										KeyData: to.Ptr("ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQ..."),
									},
								},
							},
							ProvisionVMAgent: to.Ptr(true),
							PatchSettings: &armcomputeschedule.LinuxPatchSettings{
								PatchMode:      to.Ptr(armcomputeschedule.LinuxVMGuestPatchModeAutomaticByPlatform),
								AssessmentMode: to.Ptr(armcomputeschedule.LinuxPatchAssessmentModeAutomaticByPlatform),
							},
						},
						AllowExtensionOperations: to.Ptr(true),
					},
					NetworkProfile: &armcomputeschedule.NetworkProfile{
						NetworkInterfaces: []*armcomputeschedule.NetworkInterfaceReference{
							{
								Properties: &armcomputeschedule.NetworkInterfaceReferenceProperties{
									Primary:      to.Ptr(true),
									DeleteOption: to.Ptr(armcomputeschedule.DeleteOptionsDelete),
								},
								ID: to.Ptr("/subscriptions/732116BD-AF31-4E74-9283-B387C44B4A44/resourceGroups/myResourceGroup/providers/Microsoft.Network/networkInterfaces/myNic"),
							},
						},
					},
					SecurityProfile: &armcomputeschedule.SecurityProfile{
						UefiSettings: &armcomputeschedule.UefiSettings{
							SecureBootEnabled: to.Ptr(true),
							VTpmEnabled:       to.Ptr(true),
						},
						SecurityType: to.Ptr(armcomputeschedule.SecurityTypesTrustedLaunch),
					},
					DiagnosticsProfile: &armcomputeschedule.DiagnosticsProfile{
						BootDiagnostics: &armcomputeschedule.BootDiagnostics{
							Enabled: to.Ptr(true),
						},
					},
				},
				VMExtensions: []*armcomputeschedule.BulkActionVMExtension{
					{
						Name: to.Ptr("AzureMonitorLinuxAgent"),
						Properties: &armcomputeschedule.BulkActionVMExtensionProperties{
							Publisher:               to.Ptr("Microsoft.Azure.Monitor"),
							Type:                    to.Ptr("AzureMonitorLinuxAgent"),
							TypeHandlerVersion:      to.Ptr("1.0"),
							AutoUpgradeMinorVersion: to.Ptr(true),
							EnableAutomaticUpgrade:  to.Ptr(true),
							Settings:                map[string]any{},
							SuppressFailures:        to.Ptr(false),
						},
					},
				},
			},
			VirtualMachineOverrides: []*armcomputeschedule.BulkVMConfiguration{
				{
					Name:              to.Ptr("overrideFlexVmConfig-0"),
					ComputeAPIVersion: to.Ptr("2024-07-01"),
					Zones: []*string{
						to.Ptr("2"),
					},
					Tags: map[string]*string{
						"environment": to.Ptr("production"),
						"department":  to.Ptr("engineering"),
						"role":        to.Ptr("web-server"),
					},
					Properties: &armcomputeschedule.BulkActionVMProperties{
						StorageProfile: &armcomputeschedule.StorageProfile{
							OSDisk: &armcomputeschedule.OSDisk{
								OSType:       to.Ptr(armcomputeschedule.OperatingSystemTypesLinux),
								Name:         to.Ptr("overrideOsDisk"),
								Caching:      to.Ptr(armcomputeschedule.CachingTypesReadWrite),
								CreateOption: to.Ptr(armcomputeschedule.DiskCreateOptionTypesFromImage),
								DiskSizeGB:   to.Ptr[int32](256),
								ManagedDisk: &armcomputeschedule.ManagedDiskParameters{
									StorageAccountType: to.Ptr(armcomputeschedule.StorageAccountTypesPremiumLRS),
								},
								DeleteOption: to.Ptr(armcomputeschedule.DiskDeleteOptionTypesDelete),
							},
						},
						NetworkProfile: &armcomputeschedule.NetworkProfile{
							NetworkInterfaces: []*armcomputeschedule.NetworkInterfaceReference{
								{
									Properties: &armcomputeschedule.NetworkInterfaceReferenceProperties{
										Primary:      to.Ptr(true),
										DeleteOption: to.Ptr(armcomputeschedule.DeleteOptionsDelete),
									},
									ID: to.Ptr("/subscriptions/732116BD-AF31-4E74-9283-B387C44B4A44/resourceGroups/myResourceGroup/providers/Microsoft.Network/networkInterfaces/myNic-override"),
								},
							},
						},
					},
				},
			},
			ResourceCount:  to.Ptr[int32](24),
			ResourcePrefix: to.Ptr("myFlexVm"),
			FlexProperties: &armcomputeschedule.FlexProperties{
				VMSizeProfiles: []*armcomputeschedule.VMSizeProfile{
					{
						Name: to.Ptr("Standard_D2s_v3"),
						Rank: to.Ptr[int32](24),
					},
					{
						Name: to.Ptr("Standard_D2s_v3"),
						Rank: to.Ptr[int32](24),
					},
				},
				OSType: to.Ptr(armcomputeschedule.OsTypeWindows),
				PriorityProfile: &armcomputeschedule.PriorityProfile{
					Type:               to.Ptr(armcomputeschedule.PriorityTypeRegular),
					AllocationStrategy: to.Ptr(armcomputeschedule.AllocationStrategyLowestPrice),
				},
				ZoneAllocationPolicy: &armcomputeschedule.ZoneAllocationPolicy{
					DistributionStrategy: to.Ptr(armcomputeschedule.DistributionStrategyBestEffortSingleZone),
					ZonePreferences: []*armcomputeschedule.ZonePreference{
						{
							Zone: to.Ptr("1"),
							Rank: to.Ptr[int32](21),
						},
					},
				},
			},
		},
		ExecutionParameters: &armcomputeschedule.ExecutionParameters{
			OptimizationPreference: to.Ptr(armcomputeschedule.OptimizationPreferenceCost),
			RetryPolicy: &armcomputeschedule.RetryPolicy{
				RetryCount:           to.Ptr[int32](3),
				RetryWindowInMinutes: to.Ptr[int32](30),
				OnFailureAction:      to.Ptr(armcomputeschedule.ResourceOperationTypeUnknown),
			},
		},
		Correlationid: to.Ptr("a1b2c3d4-e5f6-7890-abcd-ef1234567890"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcomputeschedule.ScheduledActionsClientVirtualMachinesExecuteCreateFlexResponse{
	// 	CreateFlexResourceOperationResponse: armcomputeschedule.CreateFlexResourceOperationResponse{
	// 		Description: to.Ptr("Flex create operation completed successfully"),
	// 		Type: to.Ptr("VirtualMachine"),
	// 		Location: to.Ptr("eastus2"),
	// 		Results: []*armcomputeschedule.ResourceOperation{
	// 			{
	// 				ResourceID: to.Ptr("/subscriptions/11111111-1111-1111-1111-111111111111/resourceGroups/rgcomputeschedule/providers/Microsoft.Compute/virtualMachines/vm1"),
	// 				ErrorCode: to.Ptr("Success"),
	// 				ErrorDetails: to.Ptr(""),
	// 				Operation: &armcomputeschedule.ResourceOperationDetails{
	// 					OperationID: to.Ptr("9b51e9d1-7f85-4c2b-a1c6-8dfd6916fee4"),
	// 					ResourceID: to.Ptr("/subscriptions/11111111-1111-1111-1111-111111111111/resourceGroups/rgcomputeschedule/providers/Microsoft.Compute/virtualMachines/vm1"),
	// 					OpType: to.Ptr(armcomputeschedule.ResourceOperationTypeUnknown),
	// 					SubscriptionID: to.Ptr("732116BD-AF31-4E74-9283-B387C44B4A44"),
	// 					Deadline: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-03-12T02:39:42.468Z"); return t}()),
	// 					DeadlineType: to.Ptr(armcomputeschedule.DeadlineTypeUnknown),
	// 					State: to.Ptr(armcomputeschedule.OperationStateUnknown),
	// 					Timezone: to.Ptr("UTC"),
	// 					TimeZone: to.Ptr("UTC"),
	// 					ResourceOperationError: &armcomputeschedule.ResourceOperationError{
	// 						ErrorCode: to.Ptr(""),
	// 						ErrorDetails: to.Ptr(""),
	// 					},
	// 					FallbackOperationInfo: &armcomputeschedule.FallbackOperationInfo{
	// 						LastOpType: to.Ptr(armcomputeschedule.ResourceOperationTypeUnknown),
	// 						Status: to.Ptr("Succeeded"),
	// 						Error: &armcomputeschedule.ResourceOperationError{
	// 							ErrorCode: to.Ptr(""),
	// 							ErrorDetails: to.Ptr(""),
	// 						},
	// 					},
	// 					CompletedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-03-12T02:39:42.468Z"); return t}()),
	// 					RetryPolicy: &armcomputeschedule.RetryPolicy{
	// 						RetryCount: to.Ptr[int32](3),
	// 						RetryWindowInMinutes: to.Ptr[int32](30),
	// 						OnFailureAction: to.Ptr(armcomputeschedule.ResourceOperationTypeUnknown),
	// 					},
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
