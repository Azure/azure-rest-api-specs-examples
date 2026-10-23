package armcontainerservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v9"
)

// Generated from example definition: 2026-08-02-preview/ManagedClustersGenerateUpgradeInsights_NodeImageLatest.json
func ExampleUpgradeInsightsClient_Generate_generateANodeOSImageUpgradeInsightUsingTheLatestAvailableImage() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerservice.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewUpgradeInsightsClient().Generate(ctx, "rg1", "clustername1", &armcontainerservice.NodeImageGenerateUpgradeInsightsContent{
		UpgradeType:   to.Ptr(armcontainerservice.UpgradeTypeNodeImage),
		AgentPoolName: to.Ptr("nodepool1"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcontainerservice.UpgradeInsightsClientGenerateResponse{
	// 	UpgradeInsight: armcontainerservice.UpgradeInsight{
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ContainerService/managedClusters/clustername1/upgradeInsights/nodeImage-nodepool1-202607.01.0"),
	// 		Name: to.Ptr("nodeImage-nodepool1-202607.01.0"),
	// 		Type: to.Ptr("Microsoft.ContainerService/managedClusters/upgradeInsights"),
	// 		Properties: &armcontainerservice.NodeImageUpgradeInsightProperties{
	// 			ProvisioningState: to.Ptr(armcontainerservice.ResourceProvisioningStateSucceeded),
	// 			UpgradeType: to.Ptr(armcontainerservice.UpgradeTypeNodeImage),
	// 			LastEvaluatedAt: to.Ptr(time.Date(2026, time.July, 7, 10, 0, 0, 0, time.UTC)),
	// 			ExpiresAt: to.Ptr(time.Date(2026, time.July, 21, 10, 0, 0, 0, time.UTC)),
	// 			AgentPoolName: to.Ptr("nodepool1"),
	// 			CurrentNodeImageVersion: to.Ptr("AKSUbuntu-2204gen2-202605.01.0"),
	// 			TargetNodeImageVersion: to.Ptr("AKSUbuntu-2204gen2-202607.01.0"),
	// 			Context: &armcontainerservice.UpgradeContext{
	// 				CurrentUpgradeStrategy: to.Ptr(armcontainerservice.UpgradeInsightStrategyAuto),
	// 			},
	// 			Evaluation: &armcontainerservice.UpgradeInsightEvaluation{
	// 				Status: to.Ptr(armcontainerservice.UpgradeInsightStatusPassed),
	// 				EvaluatedAt: to.Ptr(time.Date(2026, time.July, 7, 10, 0, 0, 0, time.UTC)),
	// 				DetectedIssues: []*armcontainerservice.DetectedUpgradeIssue{
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
