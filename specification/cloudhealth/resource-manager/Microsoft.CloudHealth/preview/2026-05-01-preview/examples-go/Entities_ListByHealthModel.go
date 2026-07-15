package armcloudhealth_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cloudhealth/armcloudhealth"
)

// Generated from example definition: 2026-05-01-preview/Entities_ListByHealthModel.json
func ExampleEntitiesClient_NewListByHealthModelPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcloudhealth.NewClientFactory("abcdef12-3456-7890-abcd-ef1234567890", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewEntitiesClient().NewListByHealthModelPager("online-store-rg", "online-store", nil)
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
		// page = armcloudhealth.EntitiesClientListByHealthModelResponse{
		// 	EntityListResult: armcloudhealth.EntityListResult{
		// 		Value: []*armcloudhealth.Entity{
		// 			{
		// 				Properties: &armcloudhealth.EntityProperties{
		// 					DisplayName: to.Ptr("Online Store"),
		// 					CanvasPosition: &armcloudhealth.EntityCoordinates{
		// 						X: to.Ptr[float32](120),
		// 						Y: to.Ptr[float32](40),
		// 					},
		// 					Icon: &armcloudhealth.IconDefinition{
		// 						IconName: to.Ptr("Globe"),
		// 					},
		// 					HealthObjective: to.Ptr[float32](99.9),
		// 					Impact: to.Ptr(armcloudhealth.EntityImpactStandard),
		// 					Tags: map[string]*string{
		// 						"environment": to.Ptr("production"),
		// 						"team": to.Ptr("online-store"),
		// 					},
		// 					SignalGroups: &armcloudhealth.SignalGroups{
		// 						Dependencies: &armcloudhealth.DependenciesSignalGroupV2{
		// 							AggregationType: to.Ptr(armcloudhealth.DependenciesAggregationTypeWorstOf),
		// 						},
		// 					},
		// 					ProvisioningState: to.Ptr(armcloudhealth.HealthModelProvisioningStateSucceeded),
		// 					HealthState: to.Ptr(armcloudhealth.HealthStateHealthy),
		// 				},
		// 				ID: to.Ptr("/subscriptions/abcdef12-3456-7890-abcd-ef1234567890/resourceGroups/online-store-rg/providers/Microsoft.CloudHealth/healthmodels/online-store/entities/storefront"),
		// 				Name: to.Ptr("storefront"),
		// 				Type: to.Ptr("Microsoft.CloudHealth/healthmodels/entities"),
		// 				SystemData: &armcloudhealth.SystemData{
		// 					CreatedBy: to.Ptr("admin@contoso.com"),
		// 					CreatedByType: to.Ptr(armcloudhealth.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-05-04T08:15:00.000Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("admin@contoso.com"),
		// 					LastModifiedByType: to.Ptr(armcloudhealth.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-05-04T09:30:00.000Z"); return t}()),
		// 				},
		// 			},
		// 			{
		// 				Properties: &armcloudhealth.EntityProperties{
		// 					DisplayName: to.Ptr("Web Frontend"),
		// 					CanvasPosition: &armcloudhealth.EntityCoordinates{
		// 						X: to.Ptr[float32](240),
		// 						Y: to.Ptr[float32](120),
		// 					},
		// 					Icon: &armcloudhealth.IconDefinition{
		// 						IconName: to.Ptr("AppService"),
		// 					},
		// 					HealthObjective: to.Ptr[float32](99.9),
		// 					Impact: to.Ptr(armcloudhealth.EntityImpactStandard),
		// 					Tags: map[string]*string{
		// 						"environment": to.Ptr("production"),
		// 						"team": to.Ptr("online-store"),
		// 					},
		// 					SignalGroups: &armcloudhealth.SignalGroups{
		// 						AzureResource: &armcloudhealth.AzureResourceSignals{
		// 							AuthenticationSetting: to.Ptr("default-auth"),
		// 							AzureResourceID: to.Ptr("/subscriptions/abcdef12-3456-7890-abcd-ef1234567890/resourceGroups/online-store-rg/providers/Microsoft.Web/sites/online-store-web"),
		// 							AzureResourceKind: to.Ptr("app"),
		// 							Signals: []*armcloudhealth.AzureResourceSignal{
		// 								{
		// 									SignalKind: to.Ptr(armcloudhealth.SignalKindAzureResourceMetric),
		// 									Name: to.Ptr("http-5xx"),
		// 									Status: &armcloudhealth.SignalStatus{
		// 										HealthState: to.Ptr(armcloudhealth.HealthStateHealthy),
		// 										Value: to.Ptr[float64](0),
		// 										ReportedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-05-04T09:30:00.000Z"); return t}()),
		// 									},
		// 									DisplayName: to.Ptr("HTTP 5xx errors"),
		// 									RefreshInterval: to.Ptr(armcloudhealth.RefreshIntervalPT1M),
		// 									DataUnit: to.Ptr("Count"),
		// 									MetricNamespace: to.Ptr("Microsoft.Web/sites"),
		// 									MetricName: to.Ptr("Http5xx"),
		// 									TimeGrain: to.Ptr("PT5M"),
		// 									AggregationType: to.Ptr(armcloudhealth.MetricAggregationTypeTotal),
		// 									EvaluationRules: &armcloudhealth.EvaluationRule{
		// 										DegradedRule: &armcloudhealth.ThresholdRuleV2{
		// 											Operator: to.Ptr(armcloudhealth.SignalOperatorGreaterThan),
		// 											Threshold: to.Ptr[float64](5),
		// 										},
		// 										UnhealthyRule: &armcloudhealth.ThresholdRuleV2{
		// 											Operator: to.Ptr(armcloudhealth.SignalOperatorGreaterThan),
		// 											Threshold: to.Ptr[float64](20),
		// 										},
		// 									},
		// 								},
		// 								{
		// 									SignalKind: to.Ptr(armcloudhealth.SignalKindAzureResourceMetric),
		// 									Name: to.Ptr("response-time"),
		// 									Status: &armcloudhealth.SignalStatus{
		// 										HealthState: to.Ptr(armcloudhealth.HealthStateHealthy),
		// 										Value: to.Ptr[float64](0.42),
		// 										ReportedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-05-04T09:30:00.000Z"); return t}()),
		// 									},
		// 									DisplayName: to.Ptr("Average response time"),
		// 									RefreshInterval: to.Ptr(armcloudhealth.RefreshIntervalPT1M),
		// 									DataUnit: to.Ptr("Seconds"),
		// 									MetricNamespace: to.Ptr("Microsoft.Web/sites"),
		// 									MetricName: to.Ptr("AverageResponseTime"),
		// 									TimeGrain: to.Ptr("PT5M"),
		// 									AggregationType: to.Ptr(armcloudhealth.MetricAggregationTypeAverage),
		// 									EvaluationRules: &armcloudhealth.EvaluationRule{
		// 										DegradedRule: &armcloudhealth.ThresholdRuleV2{
		// 											Operator: to.Ptr(armcloudhealth.SignalOperatorGreaterThan),
		// 											Threshold: to.Ptr[float64](1),
		// 										},
		// 										UnhealthyRule: &armcloudhealth.ThresholdRuleV2{
		// 											Operator: to.Ptr(armcloudhealth.SignalOperatorGreaterThan),
		// 											Threshold: to.Ptr[float64](3),
		// 										},
		// 									},
		// 								},
		// 							},
		// 							ResourceHealth: &armcloudhealth.AzureResourceHealthSignal{
		// 								Enabled: to.Ptr(armcloudhealth.ResourceHealthAvailabilityStateSignalBehaviorEnabled),
		// 								SignalName: to.Ptr("resourcehealth-availabilitystate"),
		// 								Status: &armcloudhealth.AzureResourceHealthSignalStatus{
		// 									HealthState: to.Ptr(armcloudhealth.HealthStateHealthy),
		// 									ReportedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-05-04T09:30:00.000Z"); return t}()),
		// 									AvailabilityState: to.Ptr(armcloudhealth.ResourceHealthAvailabilityStateAvailable),
		// 									Summary: to.Ptr("The App Service is available."),
		// 								},
		// 							},
		// 						},
		// 						AzureLogAnalytics: &armcloudhealth.LogAnalyticsSignals{
		// 							AuthenticationSetting: to.Ptr("default-auth"),
		// 							LogAnalyticsWorkspaceResourceID: to.Ptr("/subscriptions/abcdef12-3456-7890-abcd-ef1234567890/resourceGroups/online-store-rg/providers/Microsoft.OperationalInsights/workspaces/online-store-law"),
		// 							Signals: []*armcloudhealth.LogAnalyticsSignal{
		// 								{
		// 									SignalKind: to.Ptr(armcloudhealth.SignalKindLogAnalyticsQuery),
		// 									Name: to.Ptr("failed-request-rate"),
		// 									Status: &armcloudhealth.SignalStatus{
		// 										HealthState: to.Ptr(armcloudhealth.HealthStateDegraded),
		// 										Value: to.Ptr[float64](2.3),
		// 										ReportedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-05-04T09:30:00.000Z"); return t}()),
		// 									},
		// 									DisplayName: to.Ptr("Failed request rate"),
		// 									RefreshInterval: to.Ptr(armcloudhealth.RefreshIntervalPT5M),
		// 									DataUnit: to.Ptr("Percent"),
		// 									QueryText: to.Ptr("requests | where timestamp > ago(5m) | summarize failureRate = 100.0 * countif(success == false) / count()"),
		// 									TimeGrain: to.Ptr("PT5M"),
		// 									ValueColumnName: to.Ptr("failureRate"),
		// 									EvaluationRules: &armcloudhealth.EvaluationRule{
		// 										DegradedRule: &armcloudhealth.ThresholdRuleV2{
		// 											Operator: to.Ptr(armcloudhealth.SignalOperatorGreaterThan),
		// 											Threshold: to.Ptr[float64](1),
		// 										},
		// 										UnhealthyRule: &armcloudhealth.ThresholdRuleV2{
		// 											Operator: to.Ptr(armcloudhealth.SignalOperatorGreaterThan),
		// 											Threshold: to.Ptr[float64](5),
		// 										},
		// 									},
		// 								},
		// 							},
		// 						},
		// 						Dependencies: &armcloudhealth.DependenciesSignalGroupV2{
		// 							AggregationType: to.Ptr(armcloudhealth.DependenciesAggregationTypeWorstOf),
		// 						},
		// 					},
		// 					Alerts: &armcloudhealth.EntityAlerts{
		// 						Unhealthy: &armcloudhealth.AlertConfiguration{
		// 							Severity: to.Ptr(armcloudhealth.AlertSeveritySev1),
		// 							Description: to.Ptr("Web frontend is unhealthy."),
		// 							ActionGroupIDs: []*string{
		// 								to.Ptr("/subscriptions/abcdef12-3456-7890-abcd-ef1234567890/resourceGroups/online-store-rg/providers/Microsoft.Insights/actionGroups/online-store-oncall"),
		// 							},
		// 						},
		// 						Degraded: &armcloudhealth.AlertConfiguration{
		// 							Severity: to.Ptr(armcloudhealth.AlertSeveritySev3),
		// 							Description: to.Ptr("Web frontend is degraded."),
		// 							ActionGroupIDs: []*string{
		// 								to.Ptr("/subscriptions/abcdef12-3456-7890-abcd-ef1234567890/resourceGroups/online-store-rg/providers/Microsoft.Insights/actionGroups/online-store-oncall"),
		// 							},
		// 						},
		// 					},
		// 					ProvisioningState: to.Ptr(armcloudhealth.HealthModelProvisioningStateSucceeded),
		// 					HealthState: to.Ptr(armcloudhealth.HealthStateDegraded),
		// 				},
		// 				ID: to.Ptr("/subscriptions/abcdef12-3456-7890-abcd-ef1234567890/resourceGroups/online-store-rg/providers/Microsoft.CloudHealth/healthmodels/online-store/entities/web-frontend"),
		// 				Name: to.Ptr("web-frontend"),
		// 				Type: to.Ptr("Microsoft.CloudHealth/healthmodels/entities"),
		// 				SystemData: &armcloudhealth.SystemData{
		// 					CreatedBy: to.Ptr("admin@contoso.com"),
		// 					CreatedByType: to.Ptr(armcloudhealth.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-05-04T08:15:00.000Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("admin@contoso.com"),
		// 					LastModifiedByType: to.Ptr(armcloudhealth.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-05-04T09:30:00.000Z"); return t}()),
		// 				},
		// 			},
		// 			{
		// 				Properties: &armcloudhealth.EntityProperties{
		// 					DisplayName: to.Ptr("Orders API"),
		// 					CanvasPosition: &armcloudhealth.EntityCoordinates{
		// 						X: to.Ptr[float32](360),
		// 						Y: to.Ptr[float32](240),
		// 					},
		// 					Icon: &armcloudhealth.IconDefinition{
		// 						IconName: to.Ptr("Kubernetes"),
		// 					},
		// 					HealthObjective: to.Ptr[float32](99.9),
		// 					Impact: to.Ptr(armcloudhealth.EntityImpactStandard),
		// 					Tags: map[string]*string{
		// 						"environment": to.Ptr("production"),
		// 						"team": to.Ptr("online-store"),
		// 					},
		// 					SignalGroups: &armcloudhealth.SignalGroups{
		// 						AzureResource: &armcloudhealth.AzureResourceSignals{
		// 							AuthenticationSetting: to.Ptr("default-auth"),
		// 							AzureResourceID: to.Ptr("/subscriptions/abcdef12-3456-7890-abcd-ef1234567890/resourceGroups/online-store-rg/providers/Microsoft.ContainerService/managedClusters/online-store-aks"),
		// 							AzureResourceKind: to.Ptr("managedClusters"),
		// 							Signals: []*armcloudhealth.AzureResourceSignal{
		// 								{
		// 									SignalKind: to.Ptr(armcloudhealth.SignalKindAzureResourceMetric),
		// 									Name: to.Ptr("node-cpu"),
		// 									Status: &armcloudhealth.SignalStatus{
		// 										HealthState: to.Ptr(armcloudhealth.HealthStateHealthy),
		// 										Value: to.Ptr[float64](41.2),
		// 										ReportedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-05-04T09:30:00.000Z"); return t}()),
		// 									},
		// 									DisplayName: to.Ptr("Node CPU utilization"),
		// 									RefreshInterval: to.Ptr(armcloudhealth.RefreshIntervalPT1M),
		// 									DataUnit: to.Ptr("Percent"),
		// 									MetricNamespace: to.Ptr("Microsoft.ContainerService/managedClusters"),
		// 									MetricName: to.Ptr("node_cpu_usage_percentage"),
		// 									TimeGrain: to.Ptr("PT5M"),
		// 									AggregationType: to.Ptr(armcloudhealth.MetricAggregationTypeAverage),
		// 									EvaluationRules: &armcloudhealth.EvaluationRule{
		// 										DegradedRule: &armcloudhealth.ThresholdRuleV2{
		// 											Operator: to.Ptr(armcloudhealth.SignalOperatorGreaterThan),
		// 											Threshold: to.Ptr[float64](70),
		// 										},
		// 										UnhealthyRule: &armcloudhealth.ThresholdRuleV2{
		// 											Operator: to.Ptr(armcloudhealth.SignalOperatorGreaterThan),
		// 											Threshold: to.Ptr[float64](90),
		// 										},
		// 									},
		// 								},
		// 							},
		// 							ResourceHealth: &armcloudhealth.AzureResourceHealthSignal{
		// 								Enabled: to.Ptr(armcloudhealth.ResourceHealthAvailabilityStateSignalBehaviorEnabled),
		// 								SignalName: to.Ptr("resourcehealth-availabilitystate"),
		// 								Status: &armcloudhealth.AzureResourceHealthSignalStatus{
		// 									HealthState: to.Ptr(armcloudhealth.HealthStateHealthy),
		// 									ReportedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-05-04T09:30:00.000Z"); return t}()),
		// 									AvailabilityState: to.Ptr(armcloudhealth.ResourceHealthAvailabilityStateAvailable),
		// 									Summary: to.Ptr("The managed cluster is available."),
		// 								},
		// 							},
		// 						},
		// 						AzureMonitorWorkspace: &armcloudhealth.AzureMonitorWorkspaceSignals{
		// 							AuthenticationSetting: to.Ptr("default-auth"),
		// 							AzureMonitorWorkspaceResourceID: to.Ptr("/subscriptions/abcdef12-3456-7890-abcd-ef1234567890/resourceGroups/online-store-rg/providers/Microsoft.Monitor/accounts/online-store-amw"),
		// 							Signals: []*armcloudhealth.PrometheusMetricsSignal{
		// 								{
		// 									SignalKind: to.Ptr(armcloudhealth.SignalKindPrometheusMetricsQuery),
		// 									Name: to.Ptr("error-rate"),
		// 									Status: &armcloudhealth.SignalStatus{
		// 										HealthState: to.Ptr(armcloudhealth.HealthStateHealthy),
		// 										Value: to.Ptr[float64](0.4),
		// 										ReportedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-05-04T09:30:00.000Z"); return t}()),
		// 									},
		// 									DisplayName: to.Ptr("HTTP 5xx error rate"),
		// 									RefreshInterval: to.Ptr(armcloudhealth.RefreshIntervalPT1M),
		// 									DataUnit: to.Ptr("Percent"),
		// 									QueryText: to.Ptr("sum(rate(http_requests_total{job=\"orders-api\", code=~\"5..\"}[5m])) / sum(rate(http_requests_total{job=\"orders-api\"}[5m])) * 100"),
		// 									TimeGrain: to.Ptr("PT5M"),
		// 									EvaluationRules: &armcloudhealth.EvaluationRule{
		// 										DegradedRule: &armcloudhealth.ThresholdRuleV2{
		// 											Operator: to.Ptr(armcloudhealth.SignalOperatorGreaterThan),
		// 											Threshold: to.Ptr[float64](1),
		// 										},
		// 										UnhealthyRule: &armcloudhealth.ThresholdRuleV2{
		// 											Operator: to.Ptr(armcloudhealth.SignalOperatorGreaterThan),
		// 											Threshold: to.Ptr[float64](5),
		// 										},
		// 									},
		// 								},
		// 								{
		// 									SignalKind: to.Ptr(armcloudhealth.SignalKindPrometheusMetricsQuery),
		// 									Name: to.Ptr("p95-latency"),
		// 									Status: &armcloudhealth.SignalStatus{
		// 										HealthState: to.Ptr(armcloudhealth.HealthStateHealthy),
		// 										Value: to.Ptr[float64](180),
		// 										ReportedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-05-04T09:30:00.000Z"); return t}()),
		// 									},
		// 									DisplayName: to.Ptr("p95 request latency"),
		// 									RefreshInterval: to.Ptr(armcloudhealth.RefreshIntervalPT1M),
		// 									DataUnit: to.Ptr("MilliSeconds"),
		// 									QueryText: to.Ptr("histogram_quantile(0.95, sum by (le) (rate(http_request_duration_seconds_bucket{job=\"orders-api\"}[5m]))) * 1000"),
		// 									TimeGrain: to.Ptr("PT5M"),
		// 									EvaluationRules: &armcloudhealth.EvaluationRule{
		// 										DegradedRule: &armcloudhealth.ThresholdRuleV2{
		// 											Operator: to.Ptr(armcloudhealth.SignalOperatorGreaterThan),
		// 											Threshold: to.Ptr[float64](300),
		// 										},
		// 										UnhealthyRule: &armcloudhealth.ThresholdRuleV2{
		// 											Operator: to.Ptr(armcloudhealth.SignalOperatorGreaterThan),
		// 											Threshold: to.Ptr[float64](800),
		// 										},
		// 									},
		// 								},
		// 								{
		// 									SignalKind: to.Ptr(armcloudhealth.SignalKindPrometheusMetricsQuery),
		// 									Name: to.Ptr("pod-cpu"),
		// 									SignalDefinitionName: to.Ptr("pod-cpu-usage"),
		// 									Status: &armcloudhealth.SignalStatus{
		// 										HealthState: to.Ptr(armcloudhealth.HealthStateHealthy),
		// 										Value: to.Ptr[float64](45.3),
		// 										ReportedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-05-04T09:30:00.000Z"); return t}()),
		// 									},
		// 									DisplayName: to.Ptr("Pod CPU utilization"),
		// 									RefreshInterval: to.Ptr(armcloudhealth.RefreshIntervalPT1M),
		// 									DataUnit: to.Ptr("Percent"),
		// 									QueryText: to.Ptr("sum(rate(container_cpu_usage_seconds_total{namespace=\"online-store\", pod=~\"orders-api-.*\"}[5m])) * 100"),
		// 									TimeGrain: to.Ptr("PT5M"),
		// 									EvaluationRules: &armcloudhealth.EvaluationRule{
		// 										DegradedRule: &armcloudhealth.ThresholdRuleV2{
		// 											Operator: to.Ptr(armcloudhealth.SignalOperatorGreaterThan),
		// 											Threshold: to.Ptr[float64](70),
		// 										},
		// 										UnhealthyRule: &armcloudhealth.ThresholdRuleV2{
		// 											Operator: to.Ptr(armcloudhealth.SignalOperatorGreaterThan),
		// 											Threshold: to.Ptr[float64](90),
		// 										},
		// 									},
		// 								},
		// 							},
		// 						},
		// 						AzureLogAnalytics: &armcloudhealth.LogAnalyticsSignals{
		// 							AuthenticationSetting: to.Ptr("default-auth"),
		// 							LogAnalyticsWorkspaceResourceID: to.Ptr("/subscriptions/abcdef12-3456-7890-abcd-ef1234567890/resourceGroups/online-store-rg/providers/Microsoft.OperationalInsights/workspaces/online-store-law"),
		// 							Signals: []*armcloudhealth.LogAnalyticsSignal{
		// 								{
		// 									SignalKind: to.Ptr(armcloudhealth.SignalKindLogAnalyticsQuery),
		// 									Name: to.Ptr("unhealthy-pods"),
		// 									Status: &armcloudhealth.SignalStatus{
		// 										HealthState: to.Ptr(armcloudhealth.HealthStateHealthy),
		// 										Value: to.Ptr[float64](0),
		// 										ReportedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-05-04T09:30:00.000Z"); return t}()),
		// 									},
		// 									DisplayName: to.Ptr("Unhealthy pods"),
		// 									RefreshInterval: to.Ptr(armcloudhealth.RefreshIntervalPT5M),
		// 									DataUnit: to.Ptr("Count"),
		// 									QueryText: to.Ptr("KubePodInventory | where TimeGenerated > ago(5m) | where Namespace == 'online-store' | where PodStatus != 'Running' | summarize unhealthyPods = dcount(Name)"),
		// 									TimeGrain: to.Ptr("PT5M"),
		// 									ValueColumnName: to.Ptr("unhealthyPods"),
		// 									EvaluationRules: &armcloudhealth.EvaluationRule{
		// 										DegradedRule: &armcloudhealth.ThresholdRuleV2{
		// 											Operator: to.Ptr(armcloudhealth.SignalOperatorGreaterThan),
		// 											Threshold: to.Ptr[float64](0),
		// 										},
		// 										UnhealthyRule: &armcloudhealth.ThresholdRuleV2{
		// 											Operator: to.Ptr(armcloudhealth.SignalOperatorGreaterThan),
		// 											Threshold: to.Ptr[float64](2),
		// 										},
		// 									},
		// 								},
		// 							},
		// 						},
		// 						Dependencies: &armcloudhealth.DependenciesSignalGroupV2{
		// 							AggregationType: to.Ptr(armcloudhealth.DependenciesAggregationTypeMinHealthy),
		// 							Unit: to.Ptr(armcloudhealth.DependenciesAggregationUnitPercentage),
		// 							DegradedThreshold: to.Ptr[float64](100),
		// 							UnhealthyThreshold: to.Ptr[float64](50),
		// 							IgnoreUnknown: to.Ptr(true),
		// 						},
		// 					},
		// 					Alerts: &armcloudhealth.EntityAlerts{
		// 						Unhealthy: &armcloudhealth.AlertConfiguration{
		// 							Severity: to.Ptr(armcloudhealth.AlertSeveritySev1),
		// 							Description: to.Ptr("Orders API is unhealthy."),
		// 							ActionGroupIDs: []*string{
		// 								to.Ptr("/subscriptions/abcdef12-3456-7890-abcd-ef1234567890/resourceGroups/online-store-rg/providers/Microsoft.Insights/actionGroups/online-store-oncall"),
		// 							},
		// 						},
		// 						Degraded: &armcloudhealth.AlertConfiguration{
		// 							Severity: to.Ptr(armcloudhealth.AlertSeveritySev3),
		// 							Description: to.Ptr("Orders API is degraded."),
		// 							ActionGroupIDs: []*string{
		// 								to.Ptr("/subscriptions/abcdef12-3456-7890-abcd-ef1234567890/resourceGroups/online-store-rg/providers/Microsoft.Insights/actionGroups/online-store-oncall"),
		// 							},
		// 						},
		// 					},
		// 					ProvisioningState: to.Ptr(armcloudhealth.HealthModelProvisioningStateSucceeded),
		// 					HealthState: to.Ptr(armcloudhealth.HealthStateHealthy),
		// 				},
		// 				ID: to.Ptr("/subscriptions/abcdef12-3456-7890-abcd-ef1234567890/resourceGroups/online-store-rg/providers/Microsoft.CloudHealth/healthmodels/online-store/entities/orders-api"),
		// 				Name: to.Ptr("orders-api"),
		// 				Type: to.Ptr("Microsoft.CloudHealth/healthmodels/entities"),
		// 				SystemData: &armcloudhealth.SystemData{
		// 					CreatedBy: to.Ptr("admin@contoso.com"),
		// 					CreatedByType: to.Ptr(armcloudhealth.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-05-04T08:15:00.000Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("admin@contoso.com"),
		// 					LastModifiedByType: to.Ptr(armcloudhealth.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-05-04T09:30:00.000Z"); return t}()),
		// 				},
		// 			},
		// 			{
		// 				Properties: &armcloudhealth.EntityProperties{
		// 					DisplayName: to.Ptr("Orders Database"),
		// 					CanvasPosition: &armcloudhealth.EntityCoordinates{
		// 						X: to.Ptr[float32](480),
		// 						Y: to.Ptr[float32](360),
		// 					},
		// 					Icon: &armcloudhealth.IconDefinition{
		// 						IconName: to.Ptr("SQLDatabase"),
		// 					},
		// 					HealthObjective: to.Ptr[float32](99.5),
		// 					Impact: to.Ptr(armcloudhealth.EntityImpactStandard),
		// 					Tags: map[string]*string{
		// 						"environment": to.Ptr("production"),
		// 						"team": to.Ptr("online-store"),
		// 					},
		// 					SignalGroups: &armcloudhealth.SignalGroups{
		// 						AzureResource: &armcloudhealth.AzureResourceSignals{
		// 							AuthenticationSetting: to.Ptr("default-auth"),
		// 							AzureResourceID: to.Ptr("/subscriptions/abcdef12-3456-7890-abcd-ef1234567890/resourceGroups/online-store-rg/providers/Microsoft.Sql/servers/online-store-sql/databases/orders"),
		// 							Signals: []*armcloudhealth.AzureResourceSignal{
		// 								{
		// 									SignalKind: to.Ptr(armcloudhealth.SignalKindAzureResourceMetric),
		// 									Name: to.Ptr("sql-cpu"),
		// 									SignalDefinitionName: to.Ptr("sql-cpu-percent"),
		// 									Status: &armcloudhealth.SignalStatus{
		// 										HealthState: to.Ptr(armcloudhealth.HealthStateHealthy),
		// 										Value: to.Ptr[float64](38.5),
		// 										ReportedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-05-04T09:30:00.000Z"); return t}()),
		// 									},
		// 									DisplayName: to.Ptr("CPU utilization"),
		// 									RefreshInterval: to.Ptr(armcloudhealth.RefreshIntervalPT1M),
		// 									DataUnit: to.Ptr("Percent"),
		// 									MetricNamespace: to.Ptr("Microsoft.Sql/servers/databases"),
		// 									MetricName: to.Ptr("cpu_percent"),
		// 									TimeGrain: to.Ptr("PT5M"),
		// 									AggregationType: to.Ptr(armcloudhealth.MetricAggregationTypeAverage),
		// 									EvaluationRules: &armcloudhealth.EvaluationRule{
		// 										DegradedRule: &armcloudhealth.ThresholdRuleV2{
		// 											Operator: to.Ptr(armcloudhealth.SignalOperatorGreaterThan),
		// 											Threshold: to.Ptr[float64](70),
		// 										},
		// 										UnhealthyRule: &armcloudhealth.ThresholdRuleV2{
		// 											Operator: to.Ptr(armcloudhealth.SignalOperatorGreaterThan),
		// 											Threshold: to.Ptr[float64](90),
		// 										},
		// 									},
		// 								},
		// 								{
		// 									SignalKind: to.Ptr(armcloudhealth.SignalKindAzureResourceMetric),
		// 									Name: to.Ptr("sql-dtu"),
		// 									Status: &armcloudhealth.SignalStatus{
		// 										HealthState: to.Ptr(armcloudhealth.HealthStateHealthy),
		// 										Value: to.Ptr[float64](52.1),
		// 										ReportedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-05-04T09:30:00.000Z"); return t}()),
		// 									},
		// 									DisplayName: to.Ptr("DTU consumption"),
		// 									RefreshInterval: to.Ptr(armcloudhealth.RefreshIntervalPT1M),
		// 									DataUnit: to.Ptr("Percent"),
		// 									MetricNamespace: to.Ptr("Microsoft.Sql/servers/databases"),
		// 									MetricName: to.Ptr("dtu_consumption_percent"),
		// 									TimeGrain: to.Ptr("PT5M"),
		// 									AggregationType: to.Ptr(armcloudhealth.MetricAggregationTypeAverage),
		// 									EvaluationRules: &armcloudhealth.EvaluationRule{
		// 										DegradedRule: &armcloudhealth.ThresholdRuleV2{
		// 											Operator: to.Ptr(armcloudhealth.SignalOperatorGreaterThan),
		// 											Threshold: to.Ptr[float64](75),
		// 										},
		// 										UnhealthyRule: &armcloudhealth.ThresholdRuleV2{
		// 											Operator: to.Ptr(armcloudhealth.SignalOperatorGreaterThan),
		// 											Threshold: to.Ptr[float64](90),
		// 										},
		// 									},
		// 								},
		// 							},
		// 							ResourceHealth: &armcloudhealth.AzureResourceHealthSignal{
		// 								Enabled: to.Ptr(armcloudhealth.ResourceHealthAvailabilityStateSignalBehaviorEnabled),
		// 								SignalName: to.Ptr("resourcehealth-availabilitystate"),
		// 								Status: &armcloudhealth.AzureResourceHealthSignalStatus{
		// 									HealthState: to.Ptr(armcloudhealth.HealthStateHealthy),
		// 									ReportedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-05-04T09:30:00.000Z"); return t}()),
		// 									AvailabilityState: to.Ptr(armcloudhealth.ResourceHealthAvailabilityStateAvailable),
		// 									Summary: to.Ptr("The database is available."),
		// 								},
		// 							},
		// 						},
		// 						Dependencies: &armcloudhealth.DependenciesSignalGroupV2{
		// 							AggregationType: to.Ptr(armcloudhealth.DependenciesAggregationTypeWorstOf),
		// 						},
		// 					},
		// 					ProvisioningState: to.Ptr(armcloudhealth.HealthModelProvisioningStateSucceeded),
		// 					HealthState: to.Ptr(armcloudhealth.HealthStateHealthy),
		// 				},
		// 				ID: to.Ptr("/subscriptions/abcdef12-3456-7890-abcd-ef1234567890/resourceGroups/online-store-rg/providers/Microsoft.CloudHealth/healthmodels/online-store/entities/orders-db"),
		// 				Name: to.Ptr("orders-db"),
		// 				Type: to.Ptr("Microsoft.CloudHealth/healthmodels/entities"),
		// 				SystemData: &armcloudhealth.SystemData{
		// 					CreatedBy: to.Ptr("admin@contoso.com"),
		// 					CreatedByType: to.Ptr(armcloudhealth.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-05-04T08:15:00.000Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("admin@contoso.com"),
		// 					LastModifiedByType: to.Ptr(armcloudhealth.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-05-04T09:30:00.000Z"); return t}()),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
