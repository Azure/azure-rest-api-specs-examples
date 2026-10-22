package armcontainerservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v9"
)

// Generated from example definition: 2026-08-02-preview/ManagedClustersListUpgradeInsights.json
func ExampleUpgradeInsightsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerservice.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewUpgradeInsightsClient().NewListPager("rg1", "clustername1", nil)
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
		// page = armcontainerservice.UpgradeInsightsClientListResponse{
		// 	UpgradeInsightListResult: armcontainerservice.UpgradeInsightListResult{
		// 		Value: []*armcontainerservice.UpgradeInsight{
		// 			{
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ContainerService/managedClusters/clustername1/upgradeInsights/controlPlane-1.33.5"),
		// 				Name: to.Ptr("controlPlane-1.33.5"),
		// 				Type: to.Ptr("Microsoft.ContainerService/managedClusters/upgradeInsights"),
		// 				Properties: &armcontainerservice.ControlPlaneUpgradeInsightProperties{
		// 					ProvisioningState: to.Ptr(armcontainerservice.ResourceProvisioningStateSucceeded),
		// 					UpgradeType: to.Ptr(armcontainerservice.UpgradeTypeControlPlane),
		// 					LastEvaluatedAt: to.Ptr(time.Date(2026, time.July, 7, 10, 0, 0, 0, time.UTC)),
		// 					ExpiresAt: to.Ptr(time.Date(2026, time.July, 21, 10, 0, 0, 0, time.UTC)),
		// 					TargetKubernetesVersion: to.Ptr("1.33.5"),
		// 					CurrentKubernetesVersion: to.Ptr("1.32.7"),
		// 					KubernetesVersionSupportStatus: to.Ptr(armcontainerservice.KubernetesVersionSupportStatusInSupport),
		// 					KubernetesVersionEndOfSupportDate: to.Ptr(time.Date(2026, time.November, 30, 0, 0, 0, 0, time.UTC)),
		// 					Context: &armcontainerservice.UpgradeContext{
		// 						CurrentUpgradeStrategy: to.Ptr(armcontainerservice.UpgradeInsightStrategyAuto),
		// 						NextMaintenanceWindow: &armcontainerservice.UpgradeInsightMaintenanceWindow{
		// 							MaintenanceConfigurationName: to.Ptr("aksManagedAutoUpgradeSchedule"),
		// 							Source: to.Ptr(armcontainerservice.MaintenanceWindowSourceClusterAutoUpgradeSchedule),
		// 							StartTime: to.Ptr(time.Date(2026, time.March, 10, 2, 0, 0, 0, time.UTC)),
		// 							EndTime: to.Ptr(time.Date(2026, time.March, 10, 6, 0, 0, 0, time.UTC)),
		// 						},
		// 						LastSuccessfulUpgradeTime: to.Ptr(time.Date(2026, time.January, 15, 8, 30, 0, 0, time.UTC)),
		// 					},
		// 					Evaluation: &armcontainerservice.UpgradeInsightEvaluation{
		// 						Status: to.Ptr(armcontainerservice.UpgradeInsightStatusPassed),
		// 						EvaluatedAt: to.Ptr(time.Date(2026, time.July, 7, 10, 0, 0, 0, time.UTC)),
		// 						DetectedIssues: []*armcontainerservice.DetectedUpgradeIssue{
		// 						},
		// 					},
		// 				},
		// 			},
		// 			{
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ContainerService/managedClusters/clustername1/upgradeInsights/agentPool-nodepool1-1.33.5"),
		// 				Name: to.Ptr("agentPool-nodepool1-1.33.5"),
		// 				Type: to.Ptr("Microsoft.ContainerService/managedClusters/upgradeInsights"),
		// 				Properties: &armcontainerservice.AgentPoolUpgradeInsightProperties{
		// 					ProvisioningState: to.Ptr(armcontainerservice.ResourceProvisioningStateSucceeded),
		// 					UpgradeType: to.Ptr(armcontainerservice.UpgradeTypeAgentPool),
		// 					LastEvaluatedAt: to.Ptr(time.Date(2026, time.July, 7, 10, 0, 0, 0, time.UTC)),
		// 					ExpiresAt: to.Ptr(time.Date(2026, time.July, 21, 10, 0, 0, 0, time.UTC)),
		// 					AgentPoolName: to.Ptr("nodepool1"),
		// 					TargetKubernetesVersion: to.Ptr("1.33.5"),
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
		// 			},
		// 			{
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ContainerService/managedClusters/clustername1/upgradeInsights/nodeImage-nodepool1-202607.01.0"),
		// 				Name: to.Ptr("nodeImage-nodepool1-202607.01.0"),
		// 				Type: to.Ptr("Microsoft.ContainerService/managedClusters/upgradeInsights"),
		// 				Properties: &armcontainerservice.NodeImageUpgradeInsightProperties{
		// 					ProvisioningState: to.Ptr(armcontainerservice.ResourceProvisioningStateSucceeded),
		// 					UpgradeType: to.Ptr(armcontainerservice.UpgradeTypeNodeImage),
		// 					LastEvaluatedAt: to.Ptr(time.Date(2026, time.July, 7, 10, 0, 0, 0, time.UTC)),
		// 					ExpiresAt: to.Ptr(time.Date(2026, time.July, 21, 10, 0, 0, 0, time.UTC)),
		// 					AgentPoolName: to.Ptr("nodepool1"),
		// 					CurrentNodeImageVersion: to.Ptr("AKSUbuntu-2204gen2-202605.01.0"),
		// 					TargetNodeImageVersion: to.Ptr("AKSUbuntu-2204gen2-202607.01.0"),
		// 					Context: &armcontainerservice.UpgradeContext{
		// 						CurrentUpgradeStrategy: to.Ptr(armcontainerservice.UpgradeInsightStrategyAuto),
		// 						NextMaintenanceWindow: &armcontainerservice.UpgradeInsightMaintenanceWindow{
		// 							MaintenanceConfigurationName: to.Ptr("aksManagedNodeOSUpgradeSchedule"),
		// 							Source: to.Ptr(armcontainerservice.MaintenanceWindowSourceNodeOSUpgradeSchedule),
		// 							StartTime: to.Ptr(time.Date(2026, time.March, 5, 2, 0, 0, 0, time.UTC)),
		// 							EndTime: to.Ptr(time.Date(2026, time.March, 5, 6, 0, 0, 0, time.UTC)),
		// 						},
		// 						LastSuccessfulUpgradeTime: to.Ptr(time.Date(2026, time.May, 30, 8, 30, 0, 0, time.UTC)),
		// 					},
		// 					Evaluation: &armcontainerservice.UpgradeInsightEvaluation{
		// 						Status: to.Ptr(armcontainerservice.UpgradeInsightStatusWarning),
		// 						EvaluatedAt: to.Ptr(time.Date(2026, time.July, 7, 10, 0, 0, 0, time.UTC)),
		// 						DetectedIssues: []*armcontainerservice.DetectedUpgradeIssue{
		// 							{
		// 								IssueType: to.Ptr(armcontainerservice.UpgradeIssueTypeNodeImageOutOfSLA),
		// 								Severity: to.Ptr(armcontainerservice.UpgradeIssueSeverityWarning),
		// 								Message: to.Ptr("The node image has not been upgraded within the recommended compliance interval."),
		// 								ImpactedResources: []armcontainerservice.UpgradeInsightImpactedResourceClassification{
		// 									&armcontainerservice.ArmImpactedResource{
		// 										Kind: to.Ptr(armcontainerservice.ImpactedResourceKindArm),
		// 										ResourceType: to.Ptr("Microsoft.ContainerService/managedClusters/agentPools"),
		// 										ResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ContainerService/managedClusters/clustername1/agentPools/nodepool1"),
		// 									},
		// 								},
		// 								RecommendedActions: []*armcontainerservice.UpgradeInsightRecommendedAction{
		// 									{
		// 										Title: to.Ptr("Upgrade the node image"),
		// 										Details: to.Ptr("Trigger a node image upgrade for the agent pool to return to the recommended compliance interval."),
		// 										DocumentationLink: to.Ptr("https://learn.microsoft.com/azure/aks/upgradeinsights"),
		// 									},
		// 								},
		// 							},
		// 						},
		// 					},
		// 				},
		// 			},
		// 			{
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ContainerService/managedClusters/clustername1/upgradeInsights/controlPlaneAndAgentPools-1.33.5"),
		// 				Name: to.Ptr("controlPlaneAndAgentPools-1.33.5"),
		// 				Type: to.Ptr("Microsoft.ContainerService/managedClusters/upgradeInsights"),
		// 				Properties: &armcontainerservice.ControlPlaneAndAgentPoolsUpgradeInsightProperties{
		// 					ProvisioningState: to.Ptr(armcontainerservice.ResourceProvisioningStateSucceeded),
		// 					UpgradeType: to.Ptr(armcontainerservice.UpgradeTypeControlPlaneAndAgentPools),
		// 					LastEvaluatedAt: to.Ptr(time.Date(2026, time.July, 7, 10, 0, 0, 0, time.UTC)),
		// 					ExpiresAt: to.Ptr(time.Date(2026, time.July, 21, 10, 0, 0, 0, time.UTC)),
		// 					TargetKubernetesVersion: to.Ptr("1.33.5"),
		// 					ControlPlane: &armcontainerservice.ScopeEvaluationResult{
		// 						CurrentKubernetesVersion: to.Ptr("1.32.7"),
		// 						KubernetesVersionSupportStatus: to.Ptr(armcontainerservice.KubernetesVersionSupportStatusInSupport),
		// 						KubernetesVersionEndOfSupportDate: to.Ptr(time.Date(2026, time.November, 30, 0, 0, 0, 0, time.UTC)),
		// 						Context: &armcontainerservice.UpgradeContext{
		// 							CurrentUpgradeStrategy: to.Ptr(armcontainerservice.UpgradeInsightStrategyManual),
		// 						},
		// 						Evaluation: &armcontainerservice.UpgradeInsightEvaluation{
		// 							Status: to.Ptr(armcontainerservice.UpgradeInsightStatusFailed),
		// 							EvaluatedAt: to.Ptr(time.Date(2026, time.July, 7, 10, 0, 0, 0, time.UTC)),
		// 							DetectedIssues: []*armcontainerservice.DetectedUpgradeIssue{
		// 								{
		// 									IssueType: to.Ptr(armcontainerservice.UpgradeIssueTypeAPIBreakingChange),
		// 									Severity: to.Ptr(armcontainerservice.UpgradeIssueSeverityBlocking),
		// 									Message: to.Ptr("Deprecated APIs (policy/v1beta1 PodDisruptionBudget) are removed in 1.33.5."),
		// 									ImpactedResources: []armcontainerservice.UpgradeInsightImpactedResourceClassification{
		// 										&armcontainerservice.KubernetesImpactedResource{
		// 											Kind: to.Ptr(armcontainerservice.ImpactedResourceKindKubernetes),
		// 											ResourceType: to.Ptr("PodDisruptionBudget"),
		// 											Namespace: to.Ptr("default"),
		// 											Name: to.Ptr("my-pdb"),
		// 										},
		// 										&armcontainerservice.KubernetesImpactedResource{
		// 											Kind: to.Ptr(armcontainerservice.ImpactedResourceKindKubernetes),
		// 											ResourceType: to.Ptr("PodDisruptionBudget"),
		// 											Namespace: to.Ptr("team-a"),
		// 											Name: to.Ptr("web-pdb"),
		// 										},
		// 									},
		// 									RecommendedActions: []*armcontainerservice.UpgradeInsightRecommendedAction{
		// 										{
		// 											Title: to.Ptr("Migrate to policy/v1 PodDisruptionBudget"),
		// 											Details: to.Ptr("Update the PodDisruptionBudget manifest to apiVersion policy/v1 and re-apply before upgrading."),
		// 											DocumentationLink: to.Ptr("https://learn.microsoft.com/azure/aks/upgradeinsights"),
		// 										},
		// 									},
		// 								},
		// 							},
		// 						},
		// 					},
		// 					AgentPools: []*armcontainerservice.AgentPoolScopeEvaluation{
		// 						{
		// 							AgentPoolName: to.Ptr("nodepool1"),
		// 							CurrentKubernetesVersion: to.Ptr("1.32.7"),
		// 							KubernetesVersionSupportStatus: to.Ptr(armcontainerservice.KubernetesVersionSupportStatusInSupport),
		// 							KubernetesVersionEndOfSupportDate: to.Ptr(time.Date(2026, time.November, 30, 0, 0, 0, 0, time.UTC)),
		// 							Context: &armcontainerservice.UpgradeContext{
		// 								CurrentUpgradeStrategy: to.Ptr(armcontainerservice.UpgradeInsightStrategyManual),
		// 							},
		// 							Evaluation: &armcontainerservice.UpgradeInsightEvaluation{
		// 								Status: to.Ptr(armcontainerservice.UpgradeInsightStatusFailed),
		// 								EvaluatedAt: to.Ptr(time.Date(2026, time.July, 7, 10, 0, 0, 0, time.UTC)),
		// 								DetectedIssues: []*armcontainerservice.DetectedUpgradeIssue{
		// 									{
		// 										IssueType: to.Ptr(armcontainerservice.UpgradeIssueTypeSubnetIPExhausted),
		// 										Severity: to.Ptr(armcontainerservice.UpgradeIssueSeverityBlocking),
		// 										Message: to.Ptr("The subnet does not have sufficient available IP addresses for surge nodes during upgrade."),
		// 										ImpactedResources: []armcontainerservice.UpgradeInsightImpactedResourceClassification{
		// 											&armcontainerservice.ArmImpactedResource{
		// 												Kind: to.Ptr(armcontainerservice.ImpactedResourceKindArm),
		// 												ResourceType: to.Ptr("Microsoft.Network/virtualNetworks/subnets"),
		// 												ResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnet1/subnets/aks-subnet"),
		// 											},
		// 										},
		// 										RecommendedActions: []*armcontainerservice.UpgradeInsightRecommendedAction{
		// 											{
		// 												Title: to.Ptr("Free up or expand subnet address space"),
		// 												Details: to.Ptr("Increase the subnet size or reduce maxSurge so enough IPs are available for surge nodes."),
		// 												DocumentationLink: to.Ptr("https://learn.microsoft.com/azure/aks/upgradeinsights"),
		// 											},
		// 										},
		// 									},
		// 								},
		// 							},
		// 						},
		// 						{
		// 							AgentPoolName: to.Ptr("nodepool2"),
		// 							CurrentKubernetesVersion: to.Ptr("1.32.7"),
		// 							KubernetesVersionSupportStatus: to.Ptr(armcontainerservice.KubernetesVersionSupportStatusInSupport),
		// 							KubernetesVersionEndOfSupportDate: to.Ptr(time.Date(2026, time.November, 30, 0, 0, 0, 0, time.UTC)),
		// 							Context: &armcontainerservice.UpgradeContext{
		// 								CurrentUpgradeStrategy: to.Ptr(armcontainerservice.UpgradeInsightStrategyManual),
		// 							},
		// 							Evaluation: &armcontainerservice.UpgradeInsightEvaluation{
		// 								Status: to.Ptr(armcontainerservice.UpgradeInsightStatusPassed),
		// 								EvaluatedAt: to.Ptr(time.Date(2026, time.July, 7, 10, 0, 0, 0, time.UTC)),
		// 								DetectedIssues: []*armcontainerservice.DetectedUpgradeIssue{
		// 								},
		// 							},
		// 						},
		// 					},
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
