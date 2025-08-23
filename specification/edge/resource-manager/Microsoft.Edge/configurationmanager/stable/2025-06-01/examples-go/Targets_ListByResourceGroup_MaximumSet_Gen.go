package armworkloadorchestration_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/workloadorchestration/armworkloadorchestration"
)

// Generated from example definition: 2025-06-01/Targets_ListByResourceGroup_MaximumSet_Gen.json
func ExampleTargetsClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armworkloadorchestration.NewClientFactory("9D54FE4C-00AF-4836-8F48-B6A9C4E47192", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewTargetsClient().NewListByResourceGroupPager("rgconfigurationmanager", nil)
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
		// page = armworkloadorchestration.TargetsClientListByResourceGroupResponse{
		// 	TargetListResult: armworkloadorchestration.TargetListResult{
		// 		Value: []*armworkloadorchestration.Target{
		// 			{
		// 				Properties: &armworkloadorchestration.TargetProperties{
		// 					Description: to.Ptr("riabrxtvhlmizyhffdpjeyhvw"),
		// 					DisplayName: to.Ptr("qjlbshhqzfmwxvvynibkoi"),
		// 					TargetSpecification: map[string]any{
		// 					},
		// 					Capabilities: []*string{
		// 						to.Ptr("grjapghdidoao"),
		// 					},
		// 					ContextID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"),
		// 					HierarchyLevel: to.Ptr("octqptfirejhjfavlnfqeiikqx"),
		// 					Status: &armworkloadorchestration.DeploymentStatus{
		// 						LastModified: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-06-09T10:12:04.224Z"); return t}()),
		// 						Deployed: to.Ptr[int32](24),
		// 						ExpectedRunningJobID: to.Ptr[int32](19),
		// 						RunningJobID: to.Ptr[int32](6),
		// 						Status: to.Ptr("nnpksn"),
		// 						StatusDetails: to.Ptr("bslqqnfciczenaltdcmrgg"),
		// 						Generation: to.Ptr[int32](21),
		// 						TargetStatuses: []*armworkloadorchestration.TargetStatus{
		// 							{
		// 								Name: to.Ptr("jpbfbxmjvr"),
		// 								Status: to.Ptr("gsgkxfwtyoaepwa"),
		// 								ComponentStatuses: []*armworkloadorchestration.ComponentStatus{
		// 									{
		// 										Name: to.Ptr("lxzbkoblvaoubknkblwplf"),
		// 										Status: to.Ptr("txtthlvducufbblgtctegtgpzkzgyi"),
		// 									},
		// 								},
		// 							},
		// 						},
		// 					},
		// 					SolutionScope: to.Ptr("testname"),
		// 					State: to.Ptr(armworkloadorchestration.ResourceStateActive),
		// 					ProvisioningState: to.Ptr(armworkloadorchestration.ProvisioningStateSucceeded),
		// 				},
		// 				ETag: to.Ptr("dx"),
		// 				ExtendedLocation: &armworkloadorchestration.ExtendedLocation{
		// 					Name: to.Ptr("szjrwimeqyiue"),
		// 					Type: to.Ptr(armworkloadorchestration.ExtendedLocationTypeEdgeZone),
		// 				},
		// 				Tags: map[string]*string{
		// 					"key612": to.Ptr("vtqzrk"),
		// 				},
		// 				Location: to.Ptr("kckloegmwsjgwtcl"),
		// 				ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"),
		// 				Name: to.Ptr("hcunxaczkbi"),
		// 				Type: to.Ptr("fcgjpcmjvpoxbnslowdrcrd"),
		// 				SystemData: &armworkloadorchestration.SystemData{
		// 					CreatedBy: to.Ptr("nvjczgdguyvllp"),
		// 					CreatedByType: to.Ptr(armworkloadorchestration.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-06-09T10:11:50.747Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("uzbznzjgvaspvtqhyg"),
		// 					LastModifiedByType: to.Ptr(armworkloadorchestration.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-06-09T10:11:50.747Z"); return t}()),
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://microsoft.com/aedz"),
		// 	},
		// }
	}
}
