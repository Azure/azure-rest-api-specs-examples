package armcontainerservicefleet_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservicefleet/armcontainerservicefleet"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/3066a973f4baf2e2bf072a013b585a820bb10146/specification/containerservice/resource-manager/Microsoft.ContainerService/fleet/preview/2023-08-15-preview/examples/UpdateRuns_ListByFleet.json
func ExampleUpdateRunsClient_NewListByFleetPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerservicefleet.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewUpdateRunsClient().NewListByFleetPager("rg1", "fleet1", nil)
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
		// page.UpdateRunListResult = armcontainerservicefleet.UpdateRunListResult{
		// 	Value: []*armcontainerservicefleet.UpdateRun{
		// 		{
		// 			Name: to.Ptr("run1"),
		// 			Type: to.Ptr("Microsoft.ContainerService/fleets/updateRuns"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ContainerService/fleets/myFleet/updateRuns/run1"),
		// 			SystemData: &armcontainerservicefleet.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-03-01T01:10:08.395Z"); return t}()),
		// 				CreatedBy: to.Ptr("@contoso.com"),
		// 				CreatedByType: to.Ptr(armcontainerservicefleet.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-03-01T01:10:08.395Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("@contoso.com"),
		// 				LastModifiedByType: to.Ptr(armcontainerservicefleet.CreatedByTypeUser),
		// 			},
		// 			ETag: to.Ptr("\"EtagValue\""),
		// 			Properties: &armcontainerservicefleet.UpdateRunProperties{
		// 				ManagedClusterUpdate: &armcontainerservicefleet.ManagedClusterUpdate{
		// 					NodeImageSelection: &armcontainerservicefleet.NodeImageSelection{
		// 						Type: to.Ptr(armcontainerservicefleet.NodeImageSelectionTypeLatest),
		// 					},
		// 					Upgrade: &armcontainerservicefleet.ManagedClusterUpgradeSpec{
		// 						Type: to.Ptr(armcontainerservicefleet.ManagedClusterUpgradeTypeFull),
		// 						KubernetesVersion: to.Ptr("1.26.1"),
		// 					},
		// 				},
		// 				ProvisioningState: to.Ptr(armcontainerservicefleet.UpdateRunProvisioningStateSucceeded),
		// 				Status: &armcontainerservicefleet.UpdateRunStatus{
		// 					Stages: []*armcontainerservicefleet.UpdateStageStatus{
		// 						{
		// 							Name: to.Ptr("stage1"),
		// 							AfterStageWaitStatus: &armcontainerservicefleet.WaitStatus{
		// 								Status: &armcontainerservicefleet.UpdateStatus{
		// 									State: to.Ptr(armcontainerservicefleet.UpdateStateNotStarted),
		// 								},
		// 								WaitDurationInSeconds: to.Ptr[int32](3600),
		// 							},
		// 							Groups: []*armcontainerservicefleet.UpdateGroupStatus{
		// 								{
		// 									Name: to.Ptr("group-a"),
		// 									Members: []*armcontainerservicefleet.MemberUpdateStatus{
		// 										{
		// 											Name: to.Ptr("member-one"),
		// 											ClusterResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myClusters/providers/Microsoft.ContainerService/managedClusters/myCluster"),
		// 											Status: &armcontainerservicefleet.UpdateStatus{
		// 												State: to.Ptr(armcontainerservicefleet.UpdateStateNotStarted),
		// 											},
		// 									}},
		// 									Status: &armcontainerservicefleet.UpdateStatus{
		// 										State: to.Ptr(armcontainerservicefleet.UpdateStateNotStarted),
		// 									},
		// 							}},
		// 							Status: &armcontainerservicefleet.UpdateStatus{
		// 								State: to.Ptr(armcontainerservicefleet.UpdateStateNotStarted),
		// 							},
		// 					}},
		// 					Status: &armcontainerservicefleet.UpdateStatus{
		// 						State: to.Ptr(armcontainerservicefleet.UpdateStateNotStarted),
		// 					},
		// 				},
		// 				Strategy: &armcontainerservicefleet.UpdateRunStrategy{
		// 					Stages: []*armcontainerservicefleet.UpdateStage{
		// 						{
		// 							Name: to.Ptr("stage1"),
		// 							AfterStageWaitInSeconds: to.Ptr[int32](3600),
		// 							Groups: []*armcontainerservicefleet.UpdateGroup{
		// 								{
		// 									Name: to.Ptr("group-a"),
		// 							}},
		// 					}},
		// 				},
		// 			},
		// 	}},
		// }
	}
}
