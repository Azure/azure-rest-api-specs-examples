package armcloudhealth_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cloudhealth/armcloudhealth"
)

// Generated from example definition: 2026-01-01-preview/Entities_Get.json
func ExampleEntitiesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcloudhealth.NewClientFactory("4980D7D5-4E07-47AD-AD34-E76C6BC9F061", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewEntitiesClient().Get(ctx, "rgopenapi", "myHealthModel", "entity1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcloudhealth.EntitiesClientGetResponse{
	// 	Entity: armcloudhealth.Entity{
	// 		Properties: &armcloudhealth.EntityProperties{
	// 			ProvisioningState: to.Ptr(armcloudhealth.HealthModelProvisioningStateSucceeded),
	// 			DisplayName: to.Ptr("Entity 1"),
	// 			CanvasPosition: &armcloudhealth.EntityCoordinates{
	// 				X: to.Ptr[float32](14),
	// 				Y: to.Ptr[float32](13),
	// 			},
	// 			Icon: &armcloudhealth.IconDefinition{
	// 				IconName: to.Ptr("UserFlow"),
	// 			},
	// 			HealthObjective: to.Ptr[float32](62),
	// 			Impact: to.Ptr(armcloudhealth.EntityImpactStandard),
	// 			Tags: map[string]*string{
	// 				"key1376": to.Ptr("sample tag"),
	// 			},
	// 			HealthState: to.Ptr(armcloudhealth.HealthStateHealthy),
	// 			SignalGroups: &armcloudhealth.SignalGroups{
	// 				AzureResource: &armcloudhealth.AzureResourceSignals{
	// 					AuthenticationSetting: to.Ptr("auth123"),
	// 					AzureResourceID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789012/resourceGroups/rg1/providers/Microsoft.Compute/virtualMachines/vm1"),
	// 					AzureResourceKind: to.Ptr("functionapp"),
	// 					Signals: []*armcloudhealth.AzureResourceSignal{
	// 						{
	// 							Name: to.Ptr("uniqueSignalName1"),
	// 							SignalDefinitionName: to.Ptr("sigdef1"),
	// 							SignalKind: to.Ptr(armcloudhealth.SignalKindAzureResourceMetric),
	// 							Status: &armcloudhealth.SignalStatus{
	// 								HealthState: to.Ptr(armcloudhealth.HealthStateHealthy),
	// 								Value: to.Ptr[float64](15),
	// 								ReportedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-11-26T10:00:00Z"); return t}()),
	// 							},
	// 							MetricNamespace: to.Ptr("microsoft.compute/virtualMachines"),
	// 							MetricName: to.Ptr("cpuusage"),
	// 							AggregationType: to.Ptr(armcloudhealth.MetricAggregationTypeNone),
	// 							Dimension: to.Ptr("nodename"),
	// 							DimensionFilter: to.Ptr("node1"),
	// 							DisplayName: to.Ptr("CPU usage"),
	// 							RefreshInterval: to.Ptr(armcloudhealth.RefreshIntervalPT1M),
	// 							TimeGrain: to.Ptr("PT1M"),
	// 							DataUnit: to.Ptr("Count"),
	// 							EvaluationRules: &armcloudhealth.EvaluationRule{
	// 								DegradedRule: &armcloudhealth.ThresholdRuleV2{
	// 									Operator: to.Ptr(armcloudhealth.SignalOperator("LowerThan")),
	// 									Threshold: to.Ptr[float64](10),
	// 								},
	// 								UnhealthyRule: &armcloudhealth.ThresholdRuleV2{
	// 									Operator: to.Ptr(armcloudhealth.SignalOperator("LowerThan")),
	// 									Threshold: to.Ptr[float64](1),
	// 								},
	// 							},
	// 						},
	// 						{
	// 							Name: to.Ptr("uniqueSignalName2"),
	// 							SignalDefinitionName: to.Ptr("sigdef1"),
	// 							SignalKind: to.Ptr(armcloudhealth.SignalKindAzureResourceMetric),
	// 						},
	// 					},
	// 				},
	// 				AzureLogAnalytics: &armcloudhealth.LogAnalyticsSignals{
	// 					AuthenticationSetting: to.Ptr("auth123"),
	// 					LogAnalyticsWorkspaceResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.OperationalInsights/workspaces/myworkspace"),
	// 					Signals: []*armcloudhealth.LogAnalyticsSignal{
	// 						{
	// 							Name: to.Ptr("uniqueSignalName3"),
	// 							SignalKind: to.Ptr(armcloudhealth.SignalKindLogAnalyticsQuery),
	// 							Status: &armcloudhealth.SignalStatus{
	// 								HealthState: to.Ptr(armcloudhealth.HealthStateHealthy),
	// 								Value: to.Ptr[float64](30),
	// 								ReportedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-11-26T10:00:00Z"); return t}()),
	// 							},
	// 							EvaluationRules: &armcloudhealth.EvaluationRule{
	// 								DegradedRule: &armcloudhealth.ThresholdRuleV2{
	// 									Operator: to.Ptr(armcloudhealth.SignalOperatorGreaterThan),
	// 									Threshold: to.Ptr[float64](1),
	// 								},
	// 								UnhealthyRule: &armcloudhealth.ThresholdRuleV2{
	// 									Operator: to.Ptr(armcloudhealth.SignalOperatorGreaterThan),
	// 									Threshold: to.Ptr[float64](5),
	// 								},
	// 							},
	// 							RefreshInterval: to.Ptr(armcloudhealth.RefreshIntervalPT1M),
	// 							QueryText: to.Ptr("print 1"),
	// 							TimeGrain: to.Ptr("PT30M"),
	// 							ValueColumnName: to.Ptr("result"),
	// 							DisplayName: to.Ptr("Test LA signal"),
	// 							DataUnit: to.Ptr("my unit"),
	// 						},
	// 					},
	// 				},
	// 				AzureMonitorWorkspace: &armcloudhealth.AzureMonitorWorkspaceSignals{
	// 					AuthenticationSetting: to.Ptr("auth123"),
	// 					AzureMonitorWorkspaceResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.OperationalInsights/workspaces/myworkspace"),
	// 					Signals: []*armcloudhealth.PrometheusMetricsSignal{
	// 						{
	// 							Name: to.Ptr("pod-cpu-usage"),
	// 							SignalDefinitionName: to.Ptr("PodCpuUsageDefinition"),
	// 							SignalKind: to.Ptr(armcloudhealth.SignalKindPrometheusMetricsQuery),
	// 							Status: &armcloudhealth.SignalStatus{
	// 								HealthState: to.Ptr(armcloudhealth.HealthStateHealthy),
	// 								Value: to.Ptr[float64](45.3),
	// 								ReportedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-12-11T10:30:00Z"); return t}()),
	// 							},
	// 							DisplayName: to.Ptr("Pod CPU Usage"),
	// 							RefreshInterval: to.Ptr(armcloudhealth.RefreshIntervalPT1M),
	// 							DataUnit: to.Ptr("Percent"),
	// 							QueryText: to.Ptr("rate(container_cpu_usage_seconds_total{pod=~\"my-app-.*\"}[5m]) * 100"),
	// 							TimeGrain: to.Ptr("PT5M"),
	// 							EvaluationRules: &armcloudhealth.EvaluationRule{
	// 								DegradedRule: &armcloudhealth.ThresholdRuleV2{
	// 									Operator: to.Ptr(armcloudhealth.SignalOperatorGreaterThan),
	// 									Threshold: to.Ptr[float64](70),
	// 								},
	// 								UnhealthyRule: &armcloudhealth.ThresholdRuleV2{
	// 									Operator: to.Ptr(armcloudhealth.SignalOperatorGreaterThan),
	// 									Threshold: to.Ptr[float64](90),
	// 								},
	// 							},
	// 						},
	// 					},
	// 				},
	// 				Dependencies: &armcloudhealth.DependenciesSignalGroupV2{
	// 					AggregationType: to.Ptr(armcloudhealth.DependenciesAggregationTypeMinHealthy),
	// 					Unit: to.Ptr(armcloudhealth.DependenciesAggregationUnitPercentage),
	// 					DegradedThreshold: to.Ptr[float64](80),
	// 					UnhealthyThreshold: to.Ptr[float64](50),
	// 				},
	// 				External: &armcloudhealth.ExternalSignalGroup{
	// 					Signals: []*armcloudhealth.ExternalSignal{
	// 						{
	// 							Name: to.Ptr("external-signal-1"),
	// 							SignalKind: to.Ptr(armcloudhealth.SignalKindExternalSignal),
	// 							Status: &armcloudhealth.SignalStatus{
	// 								HealthState: to.Ptr(armcloudhealth.HealthStateDegraded),
	// 								Value: to.Ptr[float64](75.5),
	// 								ReportedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-12-15T12:00:00Z"); return t}()),
	// 							},
	// 							EvaluationRules: &armcloudhealth.EvaluationRule{
	// 								DegradedRule: &armcloudhealth.ThresholdRuleV2{
	// 									Operator: to.Ptr(armcloudhealth.SignalOperatorGreaterThan),
	// 									Threshold: to.Ptr[float64](70),
	// 								},
	// 								UnhealthyRule: &armcloudhealth.ThresholdRuleV2{
	// 									Operator: to.Ptr(armcloudhealth.SignalOperatorGreaterThan),
	// 									Threshold: to.Ptr[float64](90),
	// 								},
	// 							},
	// 						},
	// 					},
	// 				},
	// 			},
	// 			DiscoveredBy: to.Ptr("discoveryRule1"),
	// 			Alerts: &armcloudhealth.EntityAlerts{
	// 				Unhealthy: &armcloudhealth.AlertConfiguration{
	// 					Severity: to.Ptr(armcloudhealth.AlertSeveritySev1),
	// 					Description: to.Ptr("Alert description"),
	// 					ActionGroupIDs: []*string{
	// 						to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.Insights/actionGroups/myactiongroup"),
	// 					},
	// 				},
	// 				Degraded: &armcloudhealth.AlertConfiguration{
	// 					Severity: to.Ptr(armcloudhealth.AlertSeveritySev4),
	// 					Description: to.Ptr("Alert description"),
	// 					ActionGroupIDs: []*string{
	// 						to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.Insights/actionGroups/myactiongroup"),
	// 					},
	// 				},
	// 			},
	// 		},
	// 		ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789012/resourceGroups/myResourceGroup/providers/Microsoft.CloudHealth/healthmodels/myHealthModel/entities/entity1"),
	// 		Name: to.Ptr("entity1"),
	// 		Type: to.Ptr("Microsoft.CloudHealth/healthmodels/entities"),
	// 		SystemData: &armcloudhealth.SystemData{
	// 			CreatedBy: to.Ptr("cbhzxxlvkmufetjjjwtk"),
	// 			CreatedByType: to.Ptr(armcloudhealth.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-09-18T14:04:09.327Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("arz"),
	// 			LastModifiedByType: to.Ptr(armcloudhealth.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-09-18T14:04:09.328Z"); return t}()),
	// 		},
	// 	},
	// }
}
