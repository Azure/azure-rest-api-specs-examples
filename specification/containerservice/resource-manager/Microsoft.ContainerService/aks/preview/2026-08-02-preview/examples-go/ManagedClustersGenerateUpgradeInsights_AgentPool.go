package armcontainerservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v9"
)

// Generated from example definition: 2026-08-02-preview/ManagedClustersGenerateUpgradeInsights_AgentPool.json
func ExampleUpgradeInsightsClient_Generate_generateASingleAgentPoolKubernetesVersionUpgradeInsight() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerservice.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewUpgradeInsightsClient().Generate(ctx, "rg1", "clustername1", &armcontainerservice.AgentPoolGenerateUpgradeInsightsContent{
		UpgradeType:             to.Ptr(armcontainerservice.UpgradeTypeAgentPool),
		AgentPoolName:           to.Ptr("nodepool1"),
		TargetKubernetesVersion: to.Ptr("1.33.5"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcontainerservice.UpgradeInsightsClientGenerateResponse{
	// 	UpgradeInsight: armcontainerservice.UpgradeInsight{
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ContainerService/managedClusters/clustername1/upgradeInsights/agentPool-nodepool1-1.33.5"),
	// 		Name: to.Ptr("agentPool-nodepool1-1.33.5"),
	// 		Type: to.Ptr("Microsoft.ContainerService/managedClusters/upgradeInsights"),
	// 		Properties: &armcontainerservice.AgentPoolUpgradeInsightProperties{
	// 			ProvisioningState: to.Ptr(armcontainerservice.ResourceProvisioningStateSucceeded),
	// 			UpgradeType: to.Ptr(armcontainerservice.UpgradeTypeAgentPool),
	// 			LastEvaluatedAt: to.Ptr(time.Date(2026, time.July, 7, 10, 0, 0, 0, time.UTC)),
	// 			ExpiresAt: to.Ptr(time.Date(2026, time.July, 21, 10, 0, 0, 0, time.UTC)),
	// 			AgentPoolName: to.Ptr("nodepool1"),
	// 			TargetKubernetesVersion: to.Ptr("1.33.5"),
	// 			CurrentKubernetesVersion: to.Ptr("1.32.7"),
	// 			KubernetesVersionSupportStatus: to.Ptr(armcontainerservice.KubernetesVersionSupportStatusInSupport),
	// 			KubernetesVersionEndOfSupportDate: to.Ptr(time.Date(2026, time.November, 30, 0, 0, 0, 0, time.UTC)),
	// 			Context: &armcontainerservice.UpgradeContext{
	// 				CurrentUpgradeStrategy: to.Ptr(armcontainerservice.UpgradeInsightStrategyManual),
	// 			},
	// 			Evaluation: &armcontainerservice.UpgradeInsightEvaluation{
	// 				Status: to.Ptr(armcontainerservice.UpgradeInsightStatusFailed),
	// 				EvaluatedAt: to.Ptr(time.Date(2026, time.July, 7, 10, 0, 0, 0, time.UTC)),
	// 				DetectedIssues: []*armcontainerservice.DetectedUpgradeIssue{
	// 					{
	// 						IssueType: to.Ptr(armcontainerservice.UpgradeIssueTypeSubnetIPExhausted),
	// 						Severity: to.Ptr(armcontainerservice.UpgradeIssueSeverityBlocking),
	// 						Message: to.Ptr("The subnet does not have sufficient available IP addresses for surge nodes during upgrade."),
	// 						ImpactedResources: []armcontainerservice.UpgradeInsightImpactedResourceClassification{
	// 							&armcontainerservice.ArmImpactedResource{
	// 								Kind: to.Ptr(armcontainerservice.ImpactedResourceKindArm),
	// 								ResourceType: to.Ptr("Microsoft.Network/virtualNetworks/subnets"),
	// 								ResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnet1/subnets/aks-subnet"),
	// 							},
	// 						},
	// 						RecommendedActions: []*armcontainerservice.UpgradeInsightRecommendedAction{
	// 							{
	// 								Title: to.Ptr("Free up or expand subnet address space"),
	// 								Details: to.Ptr("Increase the subnet size or reduce maxSurge so enough IPs are available for surge nodes."),
	// 								DocumentationLink: to.Ptr("https://learn.microsoft.com/azure/aks/upgradeinsights"),
	// 							},
	// 						},
	// 					},
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
