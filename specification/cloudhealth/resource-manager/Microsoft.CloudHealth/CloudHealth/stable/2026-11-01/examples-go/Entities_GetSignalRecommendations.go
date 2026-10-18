package armcloudhealth_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cloudhealth/armcloudhealth"
)

// Generated from example definition: 2026-11-01/Entities_GetSignalRecommendations.json
func ExampleEntitiesClient_GetSignalRecommendations() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcloudhealth.NewClientFactory("abcdef12-3456-7890-abcd-ef1234567890", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewEntitiesClient().GetSignalRecommendations(ctx, "online-store-rg", "online-store", "online-store-aks", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcloudhealth.EntitiesClientGetSignalRecommendationsResponse{
	// 	GetSignalRecommendationsResponse: armcloudhealth.GetSignalRecommendationsResponse{
	// 		RecommendedSignals: []*armcloudhealth.SignalConfiguration{
	// 			{
	// 				SignalID: to.Ptr("2c2d2ac1-c9bc-4bb7-997d-56af23c81f41"),
	// 				DisplayName: to.Ptr("AKS node CPU usage"),
	// 				Description: to.Ptr("Monitors average CPU usage across nodes in the cluster."),
	// 				ApplicableResourceTypes: []*string{
	// 					to.Ptr("Microsoft.ContainerService/managedClusters"),
	// 				},
	// 				RefreshInterval: to.Ptr(armcloudhealth.RefreshIntervalPT5M),
	// 				DataUnit: to.Ptr("Percent"),
	// 				Configuration: &armcloudhealth.AzureResourceMetricRecommendationConfiguration{
	// 					SignalKind: to.Ptr(armcloudhealth.SignalRecommendationKindAzureResourceMetric),
	// 					MetricNamespace: to.Ptr("Microsoft.ContainerService/managedClusters"),
	// 					MetricName: to.Ptr("node_cpu_usage_percentage"),
	// 					AggregationType: to.Ptr(armcloudhealth.MetricAggregationTypeAverage),
	// 					TimeGrain: to.Ptr("PT5M"),
	// 					DimensionFilter: to.Ptr("nodepool eq 'systempool'"),
	// 				},
	// 				EvaluationRules: &armcloudhealth.EvaluationRule{
	// 					DegradedRule: &armcloudhealth.ThresholdRuleV2{
	// 						Operator: to.Ptr(armcloudhealth.SignalOperatorGreaterThan),
	// 						Threshold: to.Ptr[float64](75),
	// 					},
	// 					UnhealthyRule: &armcloudhealth.ThresholdRuleV2{
	// 						Operator: to.Ptr(armcloudhealth.SignalOperatorGreaterThan),
	// 						Threshold: to.Ptr[float64](90),
	// 					},
	// 				},
	// 			},
	// 			{
	// 				SignalID: to.Ptr("bf7e57d0-1b92-4fe0-8dc8-f48d5ec89f14"),
	// 				DisplayName: to.Ptr("Restarting Kubernetes containers"),
	// 				Description: to.Ptr("Detects containers that restart repeatedly within the evaluation window."),
	// 				ApplicableResourceTypes: []*string{
	// 					to.Ptr("Microsoft.ContainerService/managedClusters"),
	// 				},
	// 				RefreshInterval: to.Ptr(armcloudhealth.RefreshIntervalPT5M),
	// 				DataUnit: to.Ptr("Count"),
	// 				Configuration: &armcloudhealth.PrometheusMetricsRecommendationConfiguration{
	// 					SignalKind: to.Ptr(armcloudhealth.SignalRecommendationKindPrometheusMetricsQuery),
	// 					QueryText: to.Ptr("sum(increase(kube_pod_container_status_restarts_total{cluster=\"{{entity.name}}\"}[5m]))"),
	// 					TimeGrain: to.Ptr("PT5M"),
	// 					RequiredMetrics: []*string{
	// 						to.Ptr("kube_pod_container_status_restarts_total"),
	// 					},
	// 					RequiredScrapeTargets: []*string{
	// 						to.Ptr("kube-state-metrics"),
	// 					},
	// 				},
	// 				EvaluationRules: &armcloudhealth.EvaluationRule{
	// 					DegradedRule: &armcloudhealth.ThresholdRuleV2{
	// 						Operator: to.Ptr(armcloudhealth.SignalOperatorGreaterThan),
	// 						Threshold: to.Ptr[float64](3),
	// 					},
	// 					UnhealthyRule: &armcloudhealth.ThresholdRuleV2{
	// 						Operator: to.Ptr(armcloudhealth.SignalOperatorGreaterThan),
	// 						Threshold: to.Ptr[float64](10),
	// 					},
	// 				},
	// 			},
	// 		},
	// 		RecommendedConfigurations: []*armcloudhealth.SignalConfiguration{
	// 			{
	// 				SignalID: to.Ptr("f482d4ea-3c8f-44e6-9a36-a24b6279291e"),
	// 				DisplayName: to.Ptr("AKS control-plane errors"),
	// 				Description: to.Ptr("Counts error-level control-plane audit events for the cluster."),
	// 				ApplicableResourceTypes: []*string{
	// 					to.Ptr("Microsoft.ContainerService/managedClusters"),
	// 				},
	// 				RefreshInterval: to.Ptr(armcloudhealth.RefreshIntervalPT5M),
	// 				DataUnit: to.Ptr("Count"),
	// 				Configuration: &armcloudhealth.LogAnalyticsQueryRecommendationConfiguration{
	// 					SignalKind: to.Ptr(armcloudhealth.SignalRecommendationKindLogAnalyticsQuery),
	// 					QueryText: to.Ptr("AKSAuditAdmin | where _ResourceId =~ '{{entity.azureResourceId}}' | where toint(ResponseStatus.code) >= 400 | summarize value = count() by bin(TimeGenerated, 5m)"),
	// 					TimeGrain: to.Ptr("PT5M"),
	// 					ValueColumnName: to.Ptr("value"),
	// 					RequiredTables: []*string{
	// 						to.Ptr("AKSAuditAdmin"),
	// 					},
	// 					RequiredDiagnosticSettingCategories: []*string{
	// 						to.Ptr("kube-audit-admin"),
	// 					},
	// 				},
	// 				EvaluationRules: &armcloudhealth.EvaluationRule{
	// 					DegradedRule: &armcloudhealth.ThresholdRuleV2{
	// 						Operator: to.Ptr(armcloudhealth.SignalOperatorGreaterThan),
	// 						Threshold: to.Ptr[float64](0),
	// 					},
	// 					UnhealthyRule: &armcloudhealth.ThresholdRuleV2{
	// 						Operator: to.Ptr(armcloudhealth.SignalOperatorGreaterThan),
	// 						Threshold: to.Ptr[float64](10),
	// 					},
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
