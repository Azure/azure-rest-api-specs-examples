package armcontainerservicefleet_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservicefleet/armcontainerservicefleet/v3"
)

// Generated from example definition: 2025-04-01-preview/UpdateRuns_Skip_MaximumSet_Gen.json
func ExampleUpdateRunsClient_BeginSkip_skipsOneOrMoreMemberGroupStageAfterStageWaitSOfAnUpdateRunGeneratedByMaximumSetRule() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerservicefleet.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewUpdateRunsClient().BeginSkip(ctx, "rgfleets", "fleet1", "fleet1", armcontainerservicefleet.SkipProperties{
		Targets: []*armcontainerservicefleet.SkipTarget{
			{
				Type: to.Ptr(armcontainerservicefleet.TargetTypeMember),
				Name: to.Ptr("member-one"),
			},
			{
				Type: to.Ptr(armcontainerservicefleet.TargetTypeAfterStageWait),
				Name: to.Ptr("stage1"),
			},
		},
	}, &armcontainerservicefleet.UpdateRunsClientBeginSkipOptions{
		IfMatch: to.Ptr("rncfubdzrhcihvpqflbsjvoau")})
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
	// res = armcontainerservicefleet.UpdateRunsClientSkipResponse{
	// 	UpdateRun: &armcontainerservicefleet.UpdateRun{
	// 		Properties: &armcontainerservicefleet.UpdateRunProperties{
	// 			UpdateStrategyID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ContainerService/fleets/myFleet/updateStrategies/strategy1"),
	// 			Strategy: &armcontainerservicefleet.UpdateRunStrategy{
	// 				Stages: []*armcontainerservicefleet.UpdateStage{
	// 					{
	// 						Name: to.Ptr("stage1"),
	// 						Groups: []*armcontainerservicefleet.UpdateGroup{
	// 							{
	// 								Name: to.Ptr("group-a"),
	// 							},
	// 						},
	// 						AfterStageWaitInSeconds: to.Ptr[int32](3600),
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
	// 					CustomNodeImageVersions: []*armcontainerservicefleet.NodeImageVersion{
	// 						{
	// 							Version: to.Ptr("wkcmcqhecdbsoskkny"),
	// 						},
	// 					},
	// 				},
	// 			},
	// 			ProvisioningState: to.Ptr(armcontainerservicefleet.UpdateRunProvisioningStateSucceeded),
	// 			Status: &armcontainerservicefleet.UpdateRunStatus{
	// 				Status: &armcontainerservicefleet.UpdateStatus{
	// 					StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-02-14T23:29:32.768Z"); return t}()),
	// 					CompletedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-02-14T23:29:32.768Z"); return t}()),
	// 					State: to.Ptr(armcontainerservicefleet.UpdateStateNotStarted),
	// 					Error: &armcontainerservicefleet.ErrorDetail{
	// 						Code: to.Ptr("ezwocfahsfmbddlqgloysjkthkn"),
	// 						Message: to.Ptr("udtnrxlgadzqlogclb"),
	// 						Target: to.Ptr("l"),
	// 						Details: []*armcontainerservicefleet.ErrorDetail{
	// 						},
	// 						AdditionalInfo: []*armcontainerservicefleet.ErrorAdditionalInfo{
	// 							{
	// 								Type: to.Ptr("fzgprz"),
	// 								Info: map[string]any{
	// 								},
	// 							},
	// 						},
	// 					},
	// 				},
	// 				Stages: []*armcontainerservicefleet.UpdateStageStatus{
	// 					{
	// 						Status: &armcontainerservicefleet.UpdateStatus{
	// 							State: to.Ptr(armcontainerservicefleet.UpdateStateNotStarted),
	// 							StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-02-14T23:29:32.767Z"); return t}()),
	// 							CompletedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-02-14T23:29:32.767Z"); return t}()),
	// 							Error: &armcontainerservicefleet.ErrorDetail{
	// 								Code: to.Ptr("ezwocfahsfmbddlqgloysjkthkn"),
	// 								Message: to.Ptr("udtnrxlgadzqlogclb"),
	// 								Target: to.Ptr("l"),
	// 								Details: []*armcontainerservicefleet.ErrorDetail{
	// 								},
	// 								AdditionalInfo: []*armcontainerservicefleet.ErrorAdditionalInfo{
	// 									{
	// 										Type: to.Ptr("fzgprz"),
	// 										Info: map[string]any{
	// 										},
	// 									},
	// 								},
	// 							},
	// 						},
	// 						Name: to.Ptr("stage1"),
	// 						Groups: []*armcontainerservicefleet.UpdateGroupStatus{
	// 							{
	// 								Status: &armcontainerservicefleet.UpdateStatus{
	// 									State: to.Ptr(armcontainerservicefleet.UpdateStateNotStarted),
	// 									StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-02-14T23:29:32.767Z"); return t}()),
	// 									CompletedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-02-14T23:29:32.767Z"); return t}()),
	// 									Error: &armcontainerservicefleet.ErrorDetail{
	// 										Code: to.Ptr("ezwocfahsfmbddlqgloysjkthkn"),
	// 										Message: to.Ptr("udtnrxlgadzqlogclb"),
	// 										Target: to.Ptr("l"),
	// 										Details: []*armcontainerservicefleet.ErrorDetail{
	// 										},
	// 										AdditionalInfo: []*armcontainerservicefleet.ErrorAdditionalInfo{
	// 											{
	// 												Type: to.Ptr("fzgprz"),
	// 												Info: map[string]any{
	// 												},
	// 											},
	// 										},
	// 									},
	// 								},
	// 								Name: to.Ptr("group-a"),
	// 								Members: []*armcontainerservicefleet.MemberUpdateStatus{
	// 									{
	// 										Status: &armcontainerservicefleet.UpdateStatus{
	// 											State: to.Ptr(armcontainerservicefleet.UpdateStateNotStarted),
	// 											StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-02-14T23:29:32.767Z"); return t}()),
	// 											CompletedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-02-14T23:29:32.767Z"); return t}()),
	// 											Error: &armcontainerservicefleet.ErrorDetail{
	// 												Code: to.Ptr("ezwocfahsfmbddlqgloysjkthkn"),
	// 												Message: to.Ptr("udtnrxlgadzqlogclb"),
	// 												Target: to.Ptr("l"),
	// 												Details: []*armcontainerservicefleet.ErrorDetail{
	// 												},
	// 												AdditionalInfo: []*armcontainerservicefleet.ErrorAdditionalInfo{
	// 													{
	// 														Type: to.Ptr("fzgprz"),
	// 														Info: map[string]any{
	// 														},
	// 													},
	// 												},
	// 											},
	// 										},
	// 										Name: to.Ptr("member-one"),
	// 										ClusterResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myClusters/providers/Microsoft.ContainerService/managedClusters/myCluster"),
	// 										OperationID: to.Ptr("islvvdetacuskjzmkcxc"),
	// 										Message: to.Ptr("xrvhotarzemcgeen"),
	// 									},
	// 								},
	// 							},
	// 						},
	// 						AfterStageWaitStatus: &armcontainerservicefleet.WaitStatus{
	// 							Status: &armcontainerservicefleet.UpdateStatus{
	// 								State: to.Ptr(armcontainerservicefleet.UpdateStateNotStarted),
	// 								StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-02-14T23:29:32.767Z"); return t}()),
	// 								CompletedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-02-14T23:29:32.767Z"); return t}()),
	// 								Error: &armcontainerservicefleet.ErrorDetail{
	// 									Code: to.Ptr("ezwocfahsfmbddlqgloysjkthkn"),
	// 									Message: to.Ptr("udtnrxlgadzqlogclb"),
	// 									Target: to.Ptr("l"),
	// 									Details: []*armcontainerservicefleet.ErrorDetail{
	// 									},
	// 									AdditionalInfo: []*armcontainerservicefleet.ErrorAdditionalInfo{
	// 										{
	// 											Type: to.Ptr("fzgprz"),
	// 											Info: map[string]any{
	// 											},
	// 										},
	// 									},
	// 								},
	// 							},
	// 							WaitDurationInSeconds: to.Ptr[int32](3600),
	// 						},
	// 					},
	// 				},
	// 				NodeImageSelection: &armcontainerservicefleet.NodeImageSelectionStatus{
	// 					SelectedNodeImageVersions: []*armcontainerservicefleet.NodeImageVersion{
	// 						{
	// 							Version: to.Ptr("wkcmcqhecdbsoskkny"),
	// 						},
	// 					},
	// 				},
	// 			},
	// 			AutoUpgradeProfileID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rgfleets/providers/Microsoft.ContainerService/fleets/fleet1/autoUpgradeProfiles/aup1"),
	// 		},
	// 		ETag: to.Ptr("\"EtagValue\""),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ContainerService/fleets/myFleet/updateRuns/run1"),
	// 		Name: to.Ptr("run1"),
	// 		Type: to.Ptr("Microsoft.ContainerService/fleets/updateRuns"),
	// 		SystemData: &armcontainerservicefleet.SystemData{
	// 			CreatedBy: to.Ptr("someUser"),
	// 			CreatedByType: to.Ptr(armcontainerservicefleet.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-23T05:40:40.657Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("someOtherUser"),
	// 			LastModifiedByType: to.Ptr(armcontainerservicefleet.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-23T05:40:40.657Z"); return t}()),
	// 		},
	// 	},
	// }
}
