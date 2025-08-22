package armworkloadorchestration_test

import (
	"context"
	"log"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/workloadorchestration/armworkloadorchestration"
)

// Generated from example definition: 2025-06-01/Instances_CreateOrUpdate_MaximumSet_Gen.json
func ExampleInstancesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armworkloadorchestration.NewClientFactory("9D54FE4C-00AF-4836-8F48-B6A9C4E47192", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewInstancesClient().BeginCreateOrUpdate(ctx, "rgconfigurationmanager", "testname", "testname", "testname", armworkloadorchestration.Instance{
		Properties: &armworkloadorchestration.InstanceProperties{
			SolutionVersionID: to.Ptr("acpddbkfclsgxg"),
			TargetID:          to.Ptr("eguutiftuxrsavvckjrv"),
			ActiveState:       to.Ptr(armworkloadorchestration.ActiveStateActive),
			ReconciliationPolicy: &armworkloadorchestration.ReconciliationPolicyProperties{
				State:    to.Ptr(armworkloadorchestration.ReconciliationStateInactive),
				Interval: to.Ptr("szucgzdbydcowvhprhx"),
			},
			SolutionScope: to.Ptr("testname"),
			Status: &armworkloadorchestration.DeploymentStatus{
				LastModified:         to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-06-09T10:12:04.224Z"); return t }()),
				Deployed:             to.Ptr[int32](24),
				ExpectedRunningJobID: to.Ptr[int32](19),
				RunningJobID:         to.Ptr[int32](6),
				Status:               to.Ptr("nnpksn"),
				StatusDetails:        to.Ptr("bslqqnfciczenaltdcmrgg"),
				Generation:           to.Ptr[int32](21),
				TargetStatuses: []*armworkloadorchestration.TargetStatus{
					{
						Name:   to.Ptr("jpbfbxmjvr"),
						Status: to.Ptr("gsgkxfwtyoaepwa"),
						ComponentStatuses: []*armworkloadorchestration.ComponentStatus{
							{
								Name:   to.Ptr("lxzbkoblvaoubknkblwplf"),
								Status: to.Ptr("txtthlvducufbblgtctegtgpzkzgyi"),
							},
						},
					},
				},
			},
		},
		ExtendedLocation: &armworkloadorchestration.ExtendedLocation{
			Name: to.Ptr("szjrwimeqyiue"),
			Type: to.Ptr(armworkloadorchestration.ExtendedLocationTypeEdgeZone),
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
	// res = armworkloadorchestration.InstancesClientCreateOrUpdateResponse{
	// 	Instance: &armworkloadorchestration.Instance{
	// 		Properties: &armworkloadorchestration.InstanceProperties{
	// 			SolutionVersionID: to.Ptr("acpddbkfclsgxg"),
	// 			TargetID: to.Ptr("eguutiftuxrsavvckjrv"),
	// 			ActiveState: to.Ptr(armworkloadorchestration.ActiveStateActive),
	// 			ReconciliationPolicy: &armworkloadorchestration.ReconciliationPolicyProperties{
	// 				State: to.Ptr(armworkloadorchestration.ReconciliationStateInactive),
	// 				Interval: to.Ptr("szucgzdbydcowvhprhx"),
	// 			},
	// 			SolutionScope: to.Ptr("testname"),
	// 			Status: &armworkloadorchestration.DeploymentStatus{
	// 				LastModified: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-06-09T10:12:04.224Z"); return t}()),
	// 				Deployed: to.Ptr[int32](24),
	// 				ExpectedRunningJobID: to.Ptr[int32](19),
	// 				RunningJobID: to.Ptr[int32](6),
	// 				Status: to.Ptr("nnpksn"),
	// 				StatusDetails: to.Ptr("bslqqnfciczenaltdcmrgg"),
	// 				Generation: to.Ptr[int32](21),
	// 				TargetStatuses: []*armworkloadorchestration.TargetStatus{
	// 					{
	// 						Name: to.Ptr("jpbfbxmjvr"),
	// 						Status: to.Ptr("gsgkxfwtyoaepwa"),
	// 						ComponentStatuses: []*armworkloadorchestration.ComponentStatus{
	// 							{
	// 								Name: to.Ptr("lxzbkoblvaoubknkblwplf"),
	// 								Status: to.Ptr("txtthlvducufbblgtctegtgpzkzgyi"),
	// 							},
	// 						},
	// 					},
	// 				},
	// 			},
	// 			DeploymentTimestampEpoch: to.Ptr[int64](3),
	// 			ProvisioningState: to.Ptr(armworkloadorchestration.ProvisioningStateSucceeded),
	// 		},
	// 		ExtendedLocation: &armworkloadorchestration.ExtendedLocation{
	// 			Name: to.Ptr("szjrwimeqyiue"),
	// 			Type: to.Ptr(armworkloadorchestration.ExtendedLocationTypeEdgeZone),
	// 		},
	// 		ETag: to.Ptr("pwgsy"),
	// 		ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"),
	// 		Name: to.Ptr("eklhpbpitmqrsfujmjq"),
	// 		Type: to.Ptr("nmety"),
	// 		SystemData: &armworkloadorchestration.SystemData{
	// 			CreatedBy: to.Ptr("nvjczgdguyvllp"),
	// 			CreatedByType: to.Ptr(armworkloadorchestration.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-06-09T10:11:50.747Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("uzbznzjgvaspvtqhyg"),
	// 			LastModifiedByType: to.Ptr(armworkloadorchestration.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-06-09T10:11:50.747Z"); return t}()),
	// 		},
	// 	},
	// }
}
