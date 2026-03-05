package armbatch_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/batch/armbatch/v4"
)

// Generated from example definition: 2025-06-01/PoolCreate_ConfidentialDiskEncryptionSet_ForUserSubscriptionAccounts.json
func ExamplePoolClient_Create_createPoolConfidentialDiskEncryptionSetForUserSubscriptionAccounts() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbatch.NewClientFactory("12345678-1234-1234-1234-123456789012", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPoolClient().Create(ctx, "default-azurebatch-japaneast", "sampleacct", "testpool", armbatch.Pool{
		Properties: &armbatch.PoolProperties{
			VMSize: to.Ptr("Standard_DC2as_v5"),
			TaskSchedulingPolicy: &armbatch.TaskSchedulingPolicy{
				NodeFillType: to.Ptr(armbatch.ComputeNodeFillTypePack),
			},
			DeploymentConfiguration: &armbatch.DeploymentConfiguration{
				VirtualMachineConfiguration: &armbatch.VirtualMachineConfiguration{
					ImageReference: &armbatch.ImageReference{
						Publisher: to.Ptr("MicrosoftWindowsServer"),
						Offer:     to.Ptr("WindowsServer"),
						SKU:       to.Ptr("2019-datacenter-core-g2"),
						Version:   to.Ptr("latest"),
					},
					NodeAgentSKUID: to.Ptr("batch.node.windows amd64"),
					SecurityProfile: &armbatch.SecurityProfile{
						SecurityType:     to.Ptr(armbatch.SecurityTypesConfidentialVM),
						EncryptionAtHost: to.Ptr(false),
						UefiSettings: &armbatch.UefiSettings{
							VTpmEnabled:       to.Ptr(true),
							SecureBootEnabled: to.Ptr(true),
						},
					},
					OSDisk: &armbatch.OSDisk{
						ManagedDisk: &armbatch.ManagedDisk{
							StorageAccountType: to.Ptr(armbatch.StorageAccountTypeStandardLRS),
							DiskEncryptionSet: &armbatch.DiskEncryptionSetParameters{
								ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789012/resourceGroups/default-azurebatch-japaneast/providers/Microsoft.Compute/diskEncryptionSets/DiskEncryptionSetId"),
							},
							SecurityProfile: &armbatch.VMDiskSecurityProfile{
								SecurityEncryptionType: to.Ptr(armbatch.SecurityEncryptionTypesDiskWithVMGuestState),
								DiskEncryptionSet: &armbatch.DiskEncryptionSetParameters{
									ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789012/resourceGroups/default-azurebatch-japaneast/providers/Microsoft.Compute/diskEncryptionSets/ConfidentialDiskEncryptionSetId"),
								},
							},
						},
					},
					DataDisks: []*armbatch.DataDisk{
						{
							Lun:        to.Ptr[int32](0),
							DiskSizeGB: to.Ptr[int32](1024),
							ManagedDisk: &armbatch.ManagedDisk{
								StorageAccountType: to.Ptr(armbatch.StorageAccountTypeStandardLRS),
								DiskEncryptionSet: &armbatch.DiskEncryptionSetParameters{
									ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789012/resourceGroups/default-azurebatch-japaneast/providers/Microsoft.Compute/diskEncryptionSets/DiskEncryptionSetId"),
								},
							},
						},
					},
				},
			},
			ScaleSettings: &armbatch.ScaleSettings{
				FixedScale: &armbatch.FixedScaleSettings{
					TargetDedicatedNodes: to.Ptr[int32](1),
					ResizeTimeout:        to.Ptr("PT15M"),
				},
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armbatch.PoolClientCreateResponse{
	// 	ETag: "W/\"0x8D4EDFEBFADF4AB\""
	// 	Pool: &armbatch.Pool{
	// 		ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789012/resourceGroups/default-azurebatch-japaneast/providers/Microsoft.Batch/batchAccounts/sampleacct/pools/testpool"),
	// 		Name: to.Ptr("testpool"),
	// 		Type: to.Ptr("Microsoft.Batch/batchAccounts/pools"),
	// 		Etag: to.Ptr("W/\"0x8D4EDFEBFADF4AB\""),
	// 		Properties: &armbatch.PoolProperties{
	// 			LastModified: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-06-30T07:30:49.7522647Z"); return t}()),
	// 			CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-06-30T07:30:49.7522634Z"); return t}()),
	// 			ProvisioningState: to.Ptr(armbatch.PoolProvisioningStateSucceeded),
	// 			ProvisioningStateTransitionTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-06-30T07:30:49.7522634Z"); return t}()),
	// 			AllocationState: to.Ptr(armbatch.AllocationStateResizing),
	// 			AllocationStateTransitionTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-06-30T07:30:49.7522647Z"); return t}()),
	// 			VMSize: to.Ptr("STANDARD_DC2as_v5"),
	// 			InterNodeCommunication: to.Ptr(armbatch.InterNodeCommunicationStateDisabled),
	// 			TaskSlotsPerNode: to.Ptr[int32](1),
	// 			TaskSchedulingPolicy: &armbatch.TaskSchedulingPolicy{
	// 				NodeFillType: to.Ptr(armbatch.ComputeNodeFillTypePack),
	// 				JobDefaultOrder: to.Ptr(armbatch.JobDefaultOrderNone),
	// 			},
	// 			DeploymentConfiguration: &armbatch.DeploymentConfiguration{
	// 				VirtualMachineConfiguration: &armbatch.VirtualMachineConfiguration{
	// 					ImageReference: &armbatch.ImageReference{
	// 						Publisher: to.Ptr("MicrosoftWindowsServer"),
	// 						Offer: to.Ptr("WindowsServer"),
	// 						SKU: to.Ptr("2019-datacenter-core-g2"),
	// 						Version: to.Ptr("latest"),
	// 					},
	// 					NodeAgentSKUID: to.Ptr("batch.node.windows amd64"),
	// 					DataDisks: []*armbatch.DataDisk{
	// 						{
	// 							Lun: to.Ptr[int32](0),
	// 							Caching: to.Ptr(armbatch.CachingTypeNone),
	// 							DiskSizeGB: to.Ptr[int32](1024),
	// 							ManagedDisk: &armbatch.ManagedDisk{
	// 								StorageAccountType: to.Ptr(armbatch.StorageAccountTypeStandardLRS),
	// 								DiskEncryptionSet: &armbatch.DiskEncryptionSetParameters{
	// 									ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789012/resourceGroups/default-azurebatch-japaneast/providers/Microsoft.Compute/diskEncryptionSets/DiskEncryptionSetId"),
	// 								},
	// 							},
	// 						},
	// 					},
	// 					OSDisk: &armbatch.OSDisk{
	// 						Caching: to.Ptr(armbatch.CachingTypeNone),
	// 						ManagedDisk: &armbatch.ManagedDisk{
	// 							StorageAccountType: to.Ptr(armbatch.StorageAccountTypeStandardLRS),
	// 							SecurityProfile: &armbatch.VMDiskSecurityProfile{
	// 								SecurityEncryptionType: to.Ptr(armbatch.SecurityEncryptionTypesDiskWithVMGuestState),
	// 								DiskEncryptionSet: &armbatch.DiskEncryptionSetParameters{
	// 									ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789012/resourceGroups/default-azurebatch-japaneast/providers/Microsoft.Compute/diskEncryptionSets/ConfidentialDiskEncryptionSetId"),
	// 								},
	// 							},
	// 							DiskEncryptionSet: &armbatch.DiskEncryptionSetParameters{
	// 								ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789012/resourceGroups/default-azurebatch-japaneast/providers/Microsoft.Compute/diskEncryptionSets/DiskEncryptionSetId"),
	// 							},
	// 						},
	// 					},
	// 					SecurityProfile: &armbatch.SecurityProfile{
	// 						SecurityType: to.Ptr(armbatch.SecurityTypesConfidentialVM),
	// 						EncryptionAtHost: to.Ptr(false),
	// 						UefiSettings: &armbatch.UefiSettings{
	// 							SecureBootEnabled: to.Ptr(true),
	// 							VTpmEnabled: to.Ptr(true),
	// 						},
	// 					},
	// 				},
	// 			},
	// 			ScaleSettings: &armbatch.ScaleSettings{
	// 				FixedScale: &armbatch.FixedScaleSettings{
	// 					TargetDedicatedNodes: to.Ptr[int32](1),
	// 					TargetLowPriorityNodes: to.Ptr[int32](0),
	// 					ResizeTimeout: to.Ptr("PT15M"),
	// 				},
	// 			},
	// 			ResizeOperationStatus: &armbatch.ResizeOperationStatus{
	// 				TargetDedicatedNodes: to.Ptr[int32](1),
	// 				NodeDeallocationOption: to.Ptr(armbatch.ComputeNodeDeallocationOptionRequeue),
	// 				ResizeTimeout: to.Ptr("PT15M"),
	// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-06-30T07:30:49.7522324Z"); return t}()),
	// 			},
	// 			CurrentDedicatedNodes: to.Ptr[int32](0),
	// 			CurrentLowPriorityNodes: to.Ptr[int32](0),
	// 		},
	// 	},
	// }
}
