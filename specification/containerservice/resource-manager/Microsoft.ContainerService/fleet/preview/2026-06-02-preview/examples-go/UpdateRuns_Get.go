package armcontainerservicefleet_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservicefleet/armcontainerservicefleet/v3"
)

// Generated from example definition: 2026-06-02-preview/UpdateRuns_Get.json
func ExampleUpdateRunsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerservicefleet.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewUpdateRunsClient().Get(ctx, "rg1", "fleet1", "run1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcontainerservicefleet.UpdateRunsClientGetResponse{
	// 	UpdateRun: armcontainerservicefleet.UpdateRun{
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ContainerService/fleets/myFleet/updateRuns/run1"),
	// 		Name: to.Ptr("run1"),
	// 		Type: to.Ptr("Microsoft.ContainerService/fleets/updateRuns"),
	// 		SystemData: &armcontainerservicefleet.SystemData{
	// 			CreatedBy: to.Ptr("@contoso.com"),
	// 			CreatedByType: to.Ptr(armcontainerservicefleet.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(time.Date(2023, time.March, 1, 1, 10, 8, 395000000, time.UTC)),
	// 			LastModifiedBy: to.Ptr("@contoso.com"),
	// 			LastModifiedByType: to.Ptr(armcontainerservicefleet.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(time.Date(2023, time.March, 1, 1, 10, 8, 395000000, time.UTC)),
	// 		},
	// 		Properties: &armcontainerservicefleet.UpdateRunProperties{
	// 			ProvisioningState: to.Ptr(armcontainerservicefleet.UpdateRunProvisioningStateSucceeded),
	// 			Strategy: &armcontainerservicefleet.UpdateRunStrategy{
	// 				Stages: []*armcontainerservicefleet.UpdateStage{
	// 					{
	// 						Name: to.Ptr("stage1"),
	// 						MaxConcurrency: to.Ptr("10"),
	// 						Groups: []*armcontainerservicefleet.UpdateGroup{
	// 							{
	// 								Name: to.Ptr("group-a"),
	// 								MaxConcurrency: to.Ptr("2"),
	// 								MemberSelector: &armcontainerservicefleet.MemberSelector{
	// 									ByLabel: to.Ptr("tier=frontend"),
	// 								},
	// 								BeforeGates: []*armcontainerservicefleet.GateConfiguration{
	// 									{
	// 										DisplayName: to.Ptr("gate before group-a"),
	// 										Type: to.Ptr(armcontainerservicefleet.GateTypeApproval),
	// 									},
	// 								},
	// 								AfterGates: []*armcontainerservicefleet.GateConfiguration{
	// 									{
	// 										DisplayName: to.Ptr("gate after group-a"),
	// 										Type: to.Ptr(armcontainerservicefleet.GateTypeApproval),
	// 									},
	// 								},
	// 							},
	// 						},
	// 						BeforeGates: []*armcontainerservicefleet.GateConfiguration{
	// 							{
	// 								DisplayName: to.Ptr("gate before stage1"),
	// 								Type: to.Ptr(armcontainerservicefleet.GateTypeApproval),
	// 							},
	// 						},
	// 						AfterGates: []*armcontainerservicefleet.GateConfiguration{
	// 							{
	// 								DisplayName: to.Ptr("gate after stage1"),
	// 								Type: to.Ptr(armcontainerservicefleet.GateTypeApproval),
	// 							},
	// 						},
	// 						AfterStageWaitInSeconds: to.Ptr[int32](3600),
	// 					},
	// 					{
	// 						Name: to.Ptr("stage2"),
	// 						MaxConcurrency: to.Ptr("50%"),
	// 						MemberSelector: &armcontainerservicefleet.MemberSelector{
	// 							ByLabel: to.Ptr("env=production"),
	// 						},
	// 						AfterStageWaitInSeconds: to.Ptr[int32](600),
	// 					},
	// 				},
	// 			},
	// 			ManagedClusterUpdate: &armcontainerservicefleet.ManagedClusterUpdate{
	// 				Upgrade: &armcontainerservicefleet.ManagedClusterUpgradeSpec{
	// 					Type: to.Ptr(armcontainerservicefleet.ManagedClusterUpgradeTypeFull),
	// 					KubernetesVersion: to.Ptr("1.26.1"),
	// 				},
	// 				NodeImageSelection: &armcontainerservicefleet.NodeImageSelection{
	// 					Type: to.Ptr(armcontainerservicefleet.NodeImageSelectionTypeLatest),
	// 				},
	// 			},
	// 			Status: &armcontainerservicefleet.UpdateRunStatus{
	// 				Status: &armcontainerservicefleet.UpdateStatus{
	// 					State: to.Ptr(armcontainerservicefleet.UpdateStatePending),
	// 				},
	// 				Stages: []*armcontainerservicefleet.UpdateStageStatus{
	// 					{
	// 						Status: &armcontainerservicefleet.UpdateStatus{
	// 							State: to.Ptr(armcontainerservicefleet.UpdateStatePending),
	// 						},
	// 						Name: to.Ptr("stage1"),
	// 						MaxConcurrency: to.Ptr[int32](10),
	// 						Groups: []*armcontainerservicefleet.UpdateGroupStatus{
	// 							{
	// 								Status: &armcontainerservicefleet.UpdateStatus{
	// 									State: to.Ptr(armcontainerservicefleet.UpdateStateNotStarted),
	// 								},
	// 								Name: to.Ptr("group-a"),
	// 								MaxConcurrency: to.Ptr[int32](2),
	// 								Members: []*armcontainerservicefleet.MemberUpdateStatus{
	// 									{
	// 										Status: &armcontainerservicefleet.UpdateStatus{
	// 											State: to.Ptr(armcontainerservicefleet.UpdateStateNotStarted),
	// 										},
	// 										Name: to.Ptr("member-one"),
	// 										ClusterResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myClusters/providers/Microsoft.ContainerService/managedClusters/myCluster"),
	// 									},
	// 								},
	// 								BeforeGates: []*armcontainerservicefleet.UpdateRunGateStatus{
	// 									{
	// 										DisplayName: to.Ptr("gate before group-a"),
	// 										Status: &armcontainerservicefleet.UpdateStatus{
	// 											State: to.Ptr(armcontainerservicefleet.UpdateStateNotStarted),
	// 										},
	// 									},
	// 								},
	// 								AfterGates: []*armcontainerservicefleet.UpdateRunGateStatus{
	// 									{
	// 										DisplayName: to.Ptr("gate after group-a"),
	// 										Status: &armcontainerservicefleet.UpdateStatus{
	// 											State: to.Ptr(armcontainerservicefleet.UpdateStateNotStarted),
	// 										},
	// 									},
	// 								},
	// 							},
	// 						},
	// 						BeforeGates: []*armcontainerservicefleet.UpdateRunGateStatus{
	// 							{
	// 								DisplayName: to.Ptr("gate before stage1"),
	// 								GateID: to.Ptr("/subscriptions/subid1/resourcegroups/rg1/providers/Microsoft.ContainerService/fleets/fleet-1/gates/12345678-910a-bcde-f000-000000000000"),
	// 								Status: &armcontainerservicefleet.UpdateStatus{
	// 									State: to.Ptr(armcontainerservicefleet.UpdateStatePending),
	// 								},
	// 							},
	// 						},
	// 						AfterGates: []*armcontainerservicefleet.UpdateRunGateStatus{
	// 							{
	// 								DisplayName: to.Ptr("gate after stage1"),
	// 								Status: &armcontainerservicefleet.UpdateStatus{
	// 									State: to.Ptr(armcontainerservicefleet.UpdateStateNotStarted),
	// 								},
	// 							},
	// 						},
	// 						AfterStageWaitStatus: &armcontainerservicefleet.WaitStatus{
	// 							Status: &armcontainerservicefleet.UpdateStatus{
	// 								State: to.Ptr(armcontainerservicefleet.UpdateStateNotStarted),
	// 							},
	// 							WaitDurationInSeconds: to.Ptr[int32](3600),
	// 						},
	// 					},
	// 					{
	// 						Status: &armcontainerservicefleet.UpdateStatus{
	// 							State: to.Ptr(armcontainerservicefleet.UpdateStateNotStarted),
	// 						},
	// 						Name: to.Ptr("stage2"),
	// 						MaxConcurrency: to.Ptr[int32](50),
	// 						Groups: []*armcontainerservicefleet.UpdateGroupStatus{
	// 						},
	// 						AfterStageWaitStatus: &armcontainerservicefleet.WaitStatus{
	// 							Status: &armcontainerservicefleet.UpdateStatus{
	// 								State: to.Ptr(armcontainerservicefleet.UpdateStateNotStarted),
	// 							},
	// 							WaitDurationInSeconds: to.Ptr[int32](600),
	// 						},
	// 					},
	// 				},
	// 			},
	// 		},
	// 		ETag: to.Ptr("\"EtagValue\""),
	// 	},
	// }
}
