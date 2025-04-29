package armcontainerservicefleet_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservicefleet/armcontainerservicefleet/v2"
)

// Generated from example definition: 2025-03-01/UpdateRuns_ListByFleet_MaximumSet_Gen.json
func ExampleUpdateRunsClient_NewListByFleetPager_listsTheUpdateRunResourcesByFleetGeneratedByMaximumSetRule() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerservicefleet.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewUpdateRunsClient().NewListByFleetPager("rgfleets", "fleet1", nil)
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
		// page = armcontainerservicefleet.UpdateRunsClientListByFleetResponse{
		// 	UpdateRunListResult: armcontainerservicefleet.UpdateRunListResult{
		// 		Value: []*armcontainerservicefleet.UpdateRun{
		// 			{
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ContainerService/fleets/myFleet/updateRuns/run1"),
		// 				Name: to.Ptr("run1"),
		// 				Type: to.Ptr("Microsoft.ContainerService/fleets/updateRuns"),
		// 				SystemData: &armcontainerservicefleet.SystemData{
		// 					CreatedBy: to.Ptr("@contoso.com"),
		// 					CreatedByType: to.Ptr(armcontainerservicefleet.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-03-01T01:10:08.395Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("@contoso.com"),
		// 					LastModifiedByType: to.Ptr(armcontainerservicefleet.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-03-01T01:10:08.395Z"); return t}()),
		// 				},
		// 				Properties: &armcontainerservicefleet.UpdateRunProperties{
		// 					ProvisioningState: to.Ptr(armcontainerservicefleet.UpdateRunProvisioningStateSucceeded),
		// 					Strategy: &armcontainerservicefleet.UpdateRunStrategy{
		// 						Stages: []*armcontainerservicefleet.UpdateStage{
		// 							{
		// 								Name: to.Ptr("stage1"),
		// 								Groups: []*armcontainerservicefleet.UpdateGroup{
		// 									{
		// 										Name: to.Ptr("group-a"),
		// 									},
		// 								},
		// 								AfterStageWaitInSeconds: to.Ptr[int32](3600),
		// 							},
		// 						},
		// 					},
		// 					ManagedClusterUpdate: &armcontainerservicefleet.ManagedClusterUpdate{
		// 						Upgrade: &armcontainerservicefleet.ManagedClusterUpgradeSpec{
		// 							Type: to.Ptr(armcontainerservicefleet.ManagedClusterUpgradeTypeFull),
		// 							KubernetesVersion: to.Ptr("1.26.1"),
		// 						},
		// 						NodeImageSelection: &armcontainerservicefleet.NodeImageSelection{
		// 							Type: to.Ptr(armcontainerservicefleet.NodeImageSelectionTypeLatest),
		// 							CustomNodeImageVersions: []*armcontainerservicefleet.NodeImageVersion{
		// 								{
		// 									Version: to.Ptr("wkcmcqhecdbsoskkny"),
		// 								},
		// 							},
		// 						},
		// 					},
		// 					Status: &armcontainerservicefleet.UpdateRunStatus{
		// 						Status: &armcontainerservicefleet.UpdateStatus{
		// 							State: to.Ptr(armcontainerservicefleet.UpdateStateNotStarted),
		// 							StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-02-14T23:29:32.763Z"); return t}()),
		// 							CompletedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-02-14T23:29:32.763Z"); return t}()),
		// 							Error: &armcontainerservicefleet.ErrorDetail{
		// 								Code: to.Ptr("ezwocfahsfmbddlqgloysjkthkn"),
		// 								Message: to.Ptr("udtnrxlgadzqlogclb"),
		// 								Target: to.Ptr("l"),
		// 								Details: []*armcontainerservicefleet.ErrorDetail{
		// 								},
		// 								AdditionalInfo: []*armcontainerservicefleet.ErrorAdditionalInfo{
		// 									{
		// 										Type: to.Ptr("fzgprz"),
		// 										Info: &armcontainerservicefleet.ErrorAdditionalInfoInfo{
		// 										},
		// 									},
		// 								},
		// 							},
		// 						},
		// 						Stages: []*armcontainerservicefleet.UpdateStageStatus{
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
		// 												Info: &armcontainerservicefleet.ErrorAdditionalInfoInfo{
		// 												},
		// 											},
		// 										},
		// 									},
		// 								},
		// 								Name: to.Ptr("stage1"),
		// 								Groups: []*armcontainerservicefleet.UpdateGroupStatus{
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
		// 														Info: &armcontainerservicefleet.ErrorAdditionalInfoInfo{
		// 														},
		// 													},
		// 												},
		// 											},
		// 										},
		// 										Name: to.Ptr("group-a"),
		// 										Members: []*armcontainerservicefleet.MemberUpdateStatus{
		// 											{
		// 												Status: &armcontainerservicefleet.UpdateStatus{
		// 													State: to.Ptr(armcontainerservicefleet.UpdateStateNotStarted),
		// 													StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-02-14T23:29:32.767Z"); return t}()),
		// 													CompletedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-02-14T23:29:32.767Z"); return t}()),
		// 													Error: &armcontainerservicefleet.ErrorDetail{
		// 														Code: to.Ptr("ezwocfahsfmbddlqgloysjkthkn"),
		// 														Message: to.Ptr("udtnrxlgadzqlogclb"),
		// 														Target: to.Ptr("l"),
		// 														Details: []*armcontainerservicefleet.ErrorDetail{
		// 														},
		// 														AdditionalInfo: []*armcontainerservicefleet.ErrorAdditionalInfo{
		// 															{
		// 																Type: to.Ptr("fzgprz"),
		// 																Info: &armcontainerservicefleet.ErrorAdditionalInfoInfo{
		// 																},
		// 															},
		// 														},
		// 													},
		// 												},
		// 												Name: to.Ptr("member-one"),
		// 												ClusterResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myClusters/providers/Microsoft.ContainerService/managedClusters/myCluster"),
		// 												OperationID: to.Ptr("islvvdetacuskjzmkcxc"),
		// 												Message: to.Ptr("xrvhotarzemcgeen"),
		// 											},
		// 										},
		// 									},
		// 								},
		// 								AfterStageWaitStatus: &armcontainerservicefleet.WaitStatus{
		// 									Status: &armcontainerservicefleet.UpdateStatus{
		// 										State: to.Ptr(armcontainerservicefleet.UpdateStateNotStarted),
		// 										StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-02-14T23:29:32.767Z"); return t}()),
		// 										CompletedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-02-14T23:29:32.767Z"); return t}()),
		// 										Error: &armcontainerservicefleet.ErrorDetail{
		// 											Code: to.Ptr("ezwocfahsfmbddlqgloysjkthkn"),
		// 											Message: to.Ptr("udtnrxlgadzqlogclb"),
		// 											Target: to.Ptr("l"),
		// 											Details: []*armcontainerservicefleet.ErrorDetail{
		// 											},
		// 											AdditionalInfo: []*armcontainerservicefleet.ErrorAdditionalInfo{
		// 												{
		// 													Type: to.Ptr("fzgprz"),
		// 													Info: &armcontainerservicefleet.ErrorAdditionalInfoInfo{
		// 													},
		// 												},
		// 											},
		// 										},
		// 									},
		// 									WaitDurationInSeconds: to.Ptr[int32](3600),
		// 								},
		// 							},
		// 						},
		// 						NodeImageSelection: &armcontainerservicefleet.NodeImageSelectionStatus{
		// 							SelectedNodeImageVersions: []*armcontainerservicefleet.NodeImageVersion{
		// 								{
		// 									Version: to.Ptr("wkcmcqhecdbsoskkny"),
		// 								},
		// 							},
		// 						},
		// 					},
		// 					UpdateStrategyID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rgfleets/providers/Microsoft.ContainerService/fleets/fleet1/updateStrategies/strategy1"),
		// 					AutoUpgradeProfileID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rgfleets/providers/Microsoft.ContainerService/fleets/fleet1/autoUpgradeProfiles/aup1"),
		// 				},
		// 				ETag: to.Ptr("\"EtagValue\""),
		// 			},
		// 		},
		// 		NextLink: to.Ptr("http://nextlink.contoso.com"),
		// 	},
		// }
	}
}
