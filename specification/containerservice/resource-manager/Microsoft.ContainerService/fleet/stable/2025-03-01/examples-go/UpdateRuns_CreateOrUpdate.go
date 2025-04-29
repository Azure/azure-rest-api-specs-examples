package armcontainerservicefleet_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservicefleet/armcontainerservicefleet/v2"
)

// Generated from example definition: 2025-03-01/UpdateRuns_CreateOrUpdate.json
func ExampleUpdateRunsClient_BeginCreateOrUpdate_createAnUpdateRun() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerservicefleet.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewUpdateRunsClient().BeginCreateOrUpdate(ctx, "rg1", "fleet1", "run1", armcontainerservicefleet.UpdateRun{
		Properties: &armcontainerservicefleet.UpdateRunProperties{
			UpdateStrategyID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ContainerService/fleets/myFleet/updateStrategies/strategy1"),
			Strategy: &armcontainerservicefleet.UpdateRunStrategy{
				Stages: []*armcontainerservicefleet.UpdateStage{
					{
						Name: to.Ptr("stage1"),
						Groups: []*armcontainerservicefleet.UpdateGroup{
							{
								Name: to.Ptr("group-a"),
							},
						},
						AfterStageWaitInSeconds: to.Ptr[int32](3600),
					},
				},
			},
			ManagedClusterUpdate: &armcontainerservicefleet.ManagedClusterUpdate{
				Upgrade: &armcontainerservicefleet.ManagedClusterUpgradeSpec{
					Type:              to.Ptr(armcontainerservicefleet.ManagedClusterUpgradeTypeFull),
					KubernetesVersion: to.Ptr("1.26.1"),
				},
				NodeImageSelection: &armcontainerservicefleet.NodeImageSelection{
					Type: to.Ptr(armcontainerservicefleet.NodeImageSelectionTypeLatest),
				},
			},
		},
	}, nil)
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
	// res = armcontainerservicefleet.UpdateRunsClientCreateOrUpdateResponse{
	// 	UpdateRun: &armcontainerservicefleet.UpdateRun{
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ContainerService/fleets/myFleet/updateRuns/run1"),
	// 		Name: to.Ptr("run1"),
	// 		Type: to.Ptr("Microsoft.ContainerService/fleets/updateRuns"),
	// 		SystemData: &armcontainerservicefleet.SystemData{
	// 			CreatedBy: to.Ptr("@contoso.com"),
	// 			CreatedByType: to.Ptr(armcontainerservicefleet.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-03-01T01:10:08.395Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("@contoso.com"),
	// 			LastModifiedByType: to.Ptr(armcontainerservicefleet.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-03-01T01:10:08.395Z"); return t}()),
	// 		},
	// 		Properties: &armcontainerservicefleet.UpdateRunProperties{
	// 			ProvisioningState: to.Ptr(armcontainerservicefleet.UpdateRunProvisioningStateSucceeded),
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
	// 				},
	// 			},
	// 			Status: &armcontainerservicefleet.UpdateRunStatus{
	// 				Status: &armcontainerservicefleet.UpdateStatus{
	// 					State: to.Ptr(armcontainerservicefleet.UpdateStateNotStarted),
	// 				},
	// 				Stages: []*armcontainerservicefleet.UpdateStageStatus{
	// 					{
	// 						Status: &armcontainerservicefleet.UpdateStatus{
	// 							State: to.Ptr(armcontainerservicefleet.UpdateStateNotStarted),
	// 						},
	// 						Name: to.Ptr("stage1"),
	// 						Groups: []*armcontainerservicefleet.UpdateGroupStatus{
	// 							{
	// 								Status: &armcontainerservicefleet.UpdateStatus{
	// 									State: to.Ptr(armcontainerservicefleet.UpdateStateNotStarted),
	// 								},
	// 								Name: to.Ptr("group-a"),
	// 								Members: []*armcontainerservicefleet.MemberUpdateStatus{
	// 									{
	// 										Status: &armcontainerservicefleet.UpdateStatus{
	// 											State: to.Ptr(armcontainerservicefleet.UpdateStateNotStarted),
	// 										},
	// 										Name: to.Ptr("member-one"),
	// 										ClusterResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myClusters/providers/Microsoft.ContainerService/managedClusters/myCluster"),
	// 									},
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
	// 				},
	// 			},
	// 		},
	// 		ETag: to.Ptr("\"EtagValue\""),
	// 	},
	// }
}
