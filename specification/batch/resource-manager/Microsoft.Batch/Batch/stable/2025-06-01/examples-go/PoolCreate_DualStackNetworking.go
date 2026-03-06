package armbatch_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/batch/armbatch/v4"
)

// Generated from example definition: 2025-06-01/PoolCreate_DualStackNetworking.json
func ExamplePoolClient_Create_createPoolDualStackNetworking() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbatch.NewClientFactory("12345678-1234-1234-1234-123456789012", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPoolClient().Create(ctx, "default-azurebatch", "exampleacc", "dualstackpool", armbatch.Pool{
		Properties: &armbatch.PoolProperties{
			VMSize: to.Ptr("Standard_D4ds_v5"),
			NetworkConfiguration: &armbatch.NetworkConfiguration{
				PublicIPAddressConfiguration: &armbatch.PublicIPAddressConfiguration{
					IPFamilies: []*armbatch.IPFamily{
						to.Ptr(armbatch.IPFamilyIPv4),
						to.Ptr(armbatch.IPFamilyIPv6),
					},
				},
				EndpointConfiguration: &armbatch.PoolEndpointConfiguration{
					InboundNatPools: []*armbatch.InboundNatPool{
						{
							BackendPort:            to.Ptr[int32](22),
							FrontendPortRangeStart: to.Ptr[int32](40000),
							FrontendPortRangeEnd:   to.Ptr[int32](40500),
							Name:                   to.Ptr("sshpool"),
							Protocol:               to.Ptr(armbatch.InboundEndpointProtocolTCP),
							NetworkSecurityGroupRules: []*armbatch.NetworkSecurityGroupRule{
								{
									Access:              to.Ptr(armbatch.NetworkSecurityGroupRuleAccessAllow),
									Priority:            to.Ptr[int32](1000),
									SourceAddressPrefix: to.Ptr("*"),
									SourcePortRanges: []*string{
										to.Ptr("*"),
									},
								},
							},
						},
					},
				},
			},
			DeploymentConfiguration: &armbatch.DeploymentConfiguration{
				VirtualMachineConfiguration: &armbatch.VirtualMachineConfiguration{
					ImageReference: &armbatch.ImageReference{
						Publisher: to.Ptr("Canonical"),
						Offer:     to.Ptr("ubuntu-24_04-lts"),
						SKU:       to.Ptr("server"),
					},
					NodeAgentSKUID: to.Ptr("batch.node.ubuntu 24.04"),
				},
			},
			ScaleSettings: &armbatch.ScaleSettings{
				FixedScale: &armbatch.FixedScaleSettings{
					TargetDedicatedNodes:   to.Ptr[int32](1),
					TargetLowPriorityNodes: to.Ptr[int32](0),
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
	// 		ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789012/resourceGroups/default-azurebatch-japaneast/providers/Microsoft.Batch/batchAccounts/sampleacct/pools/dualstackpool"),
	// 		Name: to.Ptr("dualstackpool"),
	// 		Type: to.Ptr("Microsoft.Batch/batchAccounts/pools"),
	// 		Etag: to.Ptr("W/\"0x8DDC34D4A01A419\""),
	// 		Properties: &armbatch.PoolProperties{
	// 			LastModified: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-07-15T03:11:27.7998105Z"); return t}()),
	// 			CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-07-15T03:11:27.7997967Z"); return t}()),
	// 			ProvisioningState: to.Ptr(armbatch.PoolProvisioningStateSucceeded),
	// 			ProvisioningStateTransitionTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-07-15T03:11:27.7997967Z"); return t}()),
	// 			AllocationState: to.Ptr(armbatch.AllocationStateResizing),
	// 			AllocationStateTransitionTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-07-15T03:11:27.7998136Z"); return t}()),
	// 			VMSize: to.Ptr("Standard_D4ds_v5"),
	// 			InterNodeCommunication: to.Ptr(armbatch.InterNodeCommunicationStateDisabled),
	// 			TaskSlotsPerNode: to.Ptr[int32](1),
	// 			TaskSchedulingPolicy: &armbatch.TaskSchedulingPolicy{
	// 				NodeFillType: to.Ptr(armbatch.ComputeNodeFillTypeSpread),
	// 				JobDefaultOrder: to.Ptr(armbatch.JobDefaultOrderNone),
	// 			},
	// 			DeploymentConfiguration: &armbatch.DeploymentConfiguration{
	// 				VirtualMachineConfiguration: &armbatch.VirtualMachineConfiguration{
	// 					ImageReference: &armbatch.ImageReference{
	// 						Publisher: to.Ptr("Canonical"),
	// 						Offer: to.Ptr("ubuntu-24_04-lts"),
	// 						SKU: to.Ptr("server"),
	// 						Version: to.Ptr("latest"),
	// 					},
	// 					NodeAgentSKUID: to.Ptr("batch.node.ubuntu 24.04"),
	// 				},
	// 			},
	// 			NetworkConfiguration: &armbatch.NetworkConfiguration{
	// 				PublicIPAddressConfiguration: &armbatch.PublicIPAddressConfiguration{
	// 					IPFamilies: []*armbatch.IPFamily{
	// 						to.Ptr(armbatch.IPFamilyIPv4),
	// 						to.Ptr(armbatch.IPFamilyIPv6),
	// 					},
	// 				},
	// 				EndpointConfiguration: &armbatch.PoolEndpointConfiguration{
	// 					InboundNatPools: []*armbatch.InboundNatPool{
	// 						{
	// 							Name: to.Ptr("sshpool"),
	// 							Protocol: to.Ptr(armbatch.InboundEndpointProtocolTCP),
	// 							BackendPort: to.Ptr[int32](22),
	// 							FrontendPortRangeStart: to.Ptr[int32](40000),
	// 							FrontendPortRangeEnd: to.Ptr[int32](40500),
	// 							NetworkSecurityGroupRules: []*armbatch.NetworkSecurityGroupRule{
	// 								{
	// 									Access: to.Ptr(armbatch.NetworkSecurityGroupRuleAccessAllow),
	// 									SourceAddressPrefix: to.Ptr("*"),
	// 									Priority: to.Ptr[int32](1000),
	// 									SourcePortRanges: []*string{
	// 										to.Ptr("*"),
	// 									},
	// 								},
	// 							},
	// 						},
	// 					},
	// 				},
	// 				DynamicVNetAssignmentScope: to.Ptr(armbatch.DynamicVNetAssignmentScopeNone),
	// 				EnableAcceleratedNetworking: to.Ptr(false),
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
	// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-07-15T03:11:27.7994621Z"); return t}()),
	// 			},
	// 			CurrentDedicatedNodes: to.Ptr[int32](0),
	// 			CurrentLowPriorityNodes: to.Ptr[int32](0),
	// 		},
	// 	},
	// }
}
