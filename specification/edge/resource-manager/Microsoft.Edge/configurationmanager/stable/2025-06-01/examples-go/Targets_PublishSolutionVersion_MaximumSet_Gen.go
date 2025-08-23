package armworkloadorchestration_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/workloadorchestration/armworkloadorchestration"
)

// Generated from example definition: 2025-06-01/Targets_PublishSolutionVersion_MaximumSet_Gen.json
func ExampleTargetsClient_BeginPublishSolutionVersion() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armworkloadorchestration.NewClientFactory("9D54FE4C-00AF-4836-8F48-B6A9C4E47192", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewTargetsClient().BeginPublishSolutionVersion(ctx, "rgconfigurationmanager", "testname", armworkloadorchestration.SolutionVersionParameter{
		SolutionVersionID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"),
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
	// res = armworkloadorchestration.TargetsClientPublishSolutionVersionResponse{
	// 	SolutionVersion: &armworkloadorchestration.SolutionVersion{
	// 		Properties: &armworkloadorchestration.SolutionVersionProperties{
	// 			SolutionTemplateVersionID: to.Ptr("ykgyybnfeqxcz"),
	// 			Revision: to.Ptr[int32](17),
	// 			TargetDisplayName: to.Ptr("phndhpa"),
	// 			Configuration: to.Ptr("qg"),
	// 			TargetLevelConfiguration: to.Ptr("t"),
	// 			Specification: map[string]any{
	// 			},
	// 			ReviewID: to.Ptr("bcnlj"),
	// 			ExternalValidationID: to.Ptr("yxabw"),
	// 			State: to.Ptr(armworkloadorchestration.StateInReview),
	// 			SolutionInstanceName: to.Ptr("testname"),
	// 			SolutionDependencies: []*armworkloadorchestration.SolutionDependency{
	// 				{
	// 					SolutionVersionID: to.Ptr("jtyhpdivvfbf"),
	// 					SolutionInstanceName: to.Ptr("testname"),
	// 					SolutionTemplateVersionID: to.Ptr("zoagjfizf"),
	// 					TargetID: to.Ptr("qcbujg"),
	// 					Dependencies: []*armworkloadorchestration.SolutionDependency{
	// 					},
	// 				},
	// 			},
	// 			ErrorDetails: &armworkloadorchestration.ErrorDetail{
	// 				Code: to.Ptr("lzmwkzmn"),
	// 				Message: to.Ptr("zjdfzawocrbkgawgampjjulgesr"),
	// 				Target: to.Ptr("rncdbnknqkhjwqzkqqit"),
	// 				Details: []*armworkloadorchestration.ErrorDetail{
	// 				},
	// 				AdditionalInfo: []*armworkloadorchestration.ErrorAdditionalInfo{
	// 					{
	// 						Type: to.Ptr("betjppbcspfewqw"),
	// 						Info: map[string]any{
	// 						},
	// 					},
	// 				},
	// 			},
	// 			ProvisioningState: to.Ptr(armworkloadorchestration.ProvisioningStateSucceeded),
	// 		},
	// 		ExtendedLocation: &armworkloadorchestration.ExtendedLocation{
	// 			Name: to.Ptr("szjrwimeqyiue"),
	// 			Type: to.Ptr(armworkloadorchestration.ExtendedLocationTypeEdgeZone),
	// 		},
	// 		ETag: to.Ptr("btoclmueznpoouhnxqlboavqnsztrz"),
	// 		ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"),
	// 		Name: to.Ptr("lk"),
	// 		Type: to.Ptr("czkgymlhmhlxxpbgcmhstxwtczthq"),
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
