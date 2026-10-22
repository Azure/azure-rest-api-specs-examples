package armcontainerservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v9"
)

// Generated from example definition: 2026-08-02-preview/ManagedClustersGenerateUpgradeInsights_ControlPlaneAndAgentPools.json
func ExampleUpgradeInsightsClient_Generate_generateACombinedControlPlaneAndAgentPoolsUpgradeInsight() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerservice.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewUpgradeInsightsClient().Generate(ctx, "rg1", "clustername1", &armcontainerservice.ControlPlaneAndAgentPoolsGenerateUpgradeInsightsContent{
		UpgradeType:             to.Ptr(armcontainerservice.UpgradeTypeControlPlaneAndAgentPools),
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
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ContainerService/managedClusters/clustername1/upgradeInsights/controlPlaneAndAgentPools-1.33.5"),
	// 		Name: to.Ptr("controlPlaneAndAgentPools-1.33.5"),
	// 		Type: to.Ptr("Microsoft.ContainerService/managedClusters/upgradeInsights"),
	// 		Properties: &armcontainerservice.ControlPlaneAndAgentPoolsUpgradeInsightProperties{
	// 			ProvisioningState: to.Ptr(armcontainerservice.ResourceProvisioningStateSucceeded),
	// 			UpgradeType: to.Ptr(armcontainerservice.UpgradeTypeControlPlaneAndAgentPools),
	// 			LastEvaluatedAt: to.Ptr(time.Date(2026, time.July, 7, 10, 0, 0, 0, time.UTC)),
	// 			ExpiresAt: to.Ptr(time.Date(2026, time.July, 21, 10, 0, 0, 0, time.UTC)),
	// 			TargetKubernetesVersion: to.Ptr("1.33.5"),
	// 			ControlPlane: &armcontainerservice.ScopeEvaluationResult{
	// 				CurrentKubernetesVersion: to.Ptr("1.32.7"),
	// 				KubernetesVersionSupportStatus: to.Ptr(armcontainerservice.KubernetesVersionSupportStatusInSupport),
	// 				KubernetesVersionEndOfSupportDate: to.Ptr(time.Date(2026, time.November, 30, 0, 0, 0, 0, time.UTC)),
	// 				Context: &armcontainerservice.UpgradeContext{
	// 					CurrentUpgradeStrategy: to.Ptr(armcontainerservice.UpgradeInsightStrategyManual),
	// 				},
	// 				Evaluation: &armcontainerservice.UpgradeInsightEvaluation{
	// 					Status: to.Ptr(armcontainerservice.UpgradeInsightStatusFailed),
	// 					EvaluatedAt: to.Ptr(time.Date(2026, time.July, 7, 10, 0, 0, 0, time.UTC)),
	// 					DetectedIssues: []*armcontainerservice.DetectedUpgradeIssue{
	// 						{
	// 							IssueType: to.Ptr(armcontainerservice.UpgradeIssueTypeAPIBreakingChange),
	// 							Severity: to.Ptr(armcontainerservice.UpgradeIssueSeverityBlocking),
	// 							Message: to.Ptr("Deprecated APIs (policy/v1beta1 PodDisruptionBudget) are removed in 1.33.5."),
	// 							ImpactedResources: []armcontainerservice.UpgradeInsightImpactedResourceClassification{
	// 								&armcontainerservice.KubernetesImpactedResource{
	// 									Kind: to.Ptr(armcontainerservice.ImpactedResourceKindKubernetes),
	// 									ResourceType: to.Ptr("PodDisruptionBudget"),
	// 									Namespace: to.Ptr("default"),
	// 									Name: to.Ptr("my-pdb"),
	// 								},
	// 								&armcontainerservice.KubernetesImpactedResource{
	// 									Kind: to.Ptr(armcontainerservice.ImpactedResourceKindKubernetes),
	// 									ResourceType: to.Ptr("PodDisruptionBudget"),
	// 									Namespace: to.Ptr("team-a"),
	// 									Name: to.Ptr("web-pdb"),
	// 								},
	// 							},
	// 							RecommendedActions: []*armcontainerservice.UpgradeInsightRecommendedAction{
	// 								{
	// 									Title: to.Ptr("Migrate to policy/v1 PodDisruptionBudget"),
	// 									Details: to.Ptr("Update the PodDisruptionBudget manifest to apiVersion policy/v1 and re-apply before upgrading."),
	// 									DocumentationLink: to.Ptr("https://learn.microsoft.com/azure/aks/upgradeinsights"),
	// 								},
	// 							},
	// 						},
	// 					},
	// 				},
	// 			},
	// 			AgentPools: []*armcontainerservice.AgentPoolScopeEvaluation{
	// 				{
	// 					AgentPoolName: to.Ptr("nodepool1"),
	// 					CurrentKubernetesVersion: to.Ptr("1.32.7"),
	// 					KubernetesVersionSupportStatus: to.Ptr(armcontainerservice.KubernetesVersionSupportStatusInSupport),
	// 					KubernetesVersionEndOfSupportDate: to.Ptr(time.Date(2026, time.November, 30, 0, 0, 0, 0, time.UTC)),
	// 					Context: &armcontainerservice.UpgradeContext{
	// 						CurrentUpgradeStrategy: to.Ptr(armcontainerservice.UpgradeInsightStrategyManual),
	// 					},
	// 					Evaluation: &armcontainerservice.UpgradeInsightEvaluation{
	// 						Status: to.Ptr(armcontainerservice.UpgradeInsightStatusFailed),
	// 						EvaluatedAt: to.Ptr(time.Date(2026, time.July, 7, 10, 0, 0, 0, time.UTC)),
	// 						DetectedIssues: []*armcontainerservice.DetectedUpgradeIssue{
	// 							{
	// 								IssueType: to.Ptr(armcontainerservice.UpgradeIssueTypeSubnetIPExhausted),
	// 								Severity: to.Ptr(armcontainerservice.UpgradeIssueSeverityBlocking),
	// 								Message: to.Ptr("The subnet does not have sufficient available IP addresses for surge nodes during upgrade."),
	// 								ImpactedResources: []armcontainerservice.UpgradeInsightImpactedResourceClassification{
	// 									&armcontainerservice.ArmImpactedResource{
	// 										Kind: to.Ptr(armcontainerservice.ImpactedResourceKindArm),
	// 										ResourceType: to.Ptr("Microsoft.Network/virtualNetworks/subnets"),
	// 										ResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnet1/subnets/aks-subnet"),
	// 									},
	// 								},
	// 								RecommendedActions: []*armcontainerservice.UpgradeInsightRecommendedAction{
	// 									{
	// 										Title: to.Ptr("Free up or expand subnet address space"),
	// 										Details: to.Ptr("Increase the subnet size or reduce maxSurge so enough IPs are available for surge nodes."),
	// 										DocumentationLink: to.Ptr("https://learn.microsoft.com/azure/aks/upgradeinsights"),
	// 									},
	// 								},
	// 							},
	// 						},
	// 					},
	// 				},
	// 				{
	// 					AgentPoolName: to.Ptr("nodepool2"),
	// 					CurrentKubernetesVersion: to.Ptr("1.32.7"),
	// 					KubernetesVersionSupportStatus: to.Ptr(armcontainerservice.KubernetesVersionSupportStatusInSupport),
	// 					KubernetesVersionEndOfSupportDate: to.Ptr(time.Date(2026, time.November, 30, 0, 0, 0, 0, time.UTC)),
	// 					Context: &armcontainerservice.UpgradeContext{
	// 						CurrentUpgradeStrategy: to.Ptr(armcontainerservice.UpgradeInsightStrategyManual),
	// 					},
	// 					Evaluation: &armcontainerservice.UpgradeInsightEvaluation{
	// 						Status: to.Ptr(armcontainerservice.UpgradeInsightStatusPassed),
	// 						EvaluatedAt: to.Ptr(time.Date(2026, time.July, 7, 10, 0, 0, 0, time.UTC)),
	// 						DetectedIssues: []*armcontainerservice.DetectedUpgradeIssue{
	// 						},
	// 					},
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
