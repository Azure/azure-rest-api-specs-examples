package armalertsmanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/alertsmanagement/armalertsmanagement"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/6d2438481021a94793b07b226df06d5f3c61d51d/specification/alertsmanagement/resource-manager/Microsoft.AlertsManagement/preview/2023-01-01-preview/examples/AlertRuleRecommendations_GetByResource_MAC.json
func ExampleAlertRuleRecommendationsClient_NewListByResourcePager_listAlertRuleRecommendationsForMonitoringAccountsAtResourceLevel() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armalertsmanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAlertRuleRecommendationsClient().NewListByResourcePager("subscriptions/2f00cc51-6809-498f-9ffc-48c42aff570d/resourceGroups/GenevaAlertRP-RunnerResources-eastus/providers/microsoft.monitor/accounts/alertsrp-eastus-pgms", nil)
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
		// page.AlertRuleRecommendationsListResponse = armalertsmanagement.AlertRuleRecommendationsListResponse{
		// 	Value: []*armalertsmanagement.AlertRuleRecommendationResource{
		// 		{
		// 			Name: to.Ptr("NodeRecordingRulesRuleGroup"),
		// 			Type: to.Ptr("Microsoft.AlertsManagement/alertRuleRecommendations"),
		// 			ID: to.Ptr("/subscriptions/2f00cc51-6809-498f-9ffc-48c42aff570d/resourceGroups/GenevaAlertRP-RunnerResources-eastus/providers/microsoft.monitor/accounts/alertsrp-eastus-pgms/providers/Microsoft.AlertsManagement/alertRuleRecommendations/NodeRecordingRulesRuleGroup"),
		// 			Properties: &armalertsmanagement.AlertRuleRecommendationProperties{
		// 				AlertRuleType: to.Ptr("Microsoft.AlertsManagement/prometheusRuleGroups"),
		// 				DisplayInformation: map[string]*string{
		// 					"ruleInfo": to.Ptr("Rule Information for first recording rule."),
		// 				},
		// 				RuleArmTemplate: &armalertsmanagement.RuleArmTemplate{
		// 					Schema: to.Ptr("https://schema.management.azure.com/schemas/2019-04-01/deploymentTemplate.json#"),
		// 					ContentVersion: to.Ptr("1.0.0.0"),
		// 					Parameters: map[string]any{
		// 						"actionGroupIds":map[string]any{
		// 							"type": "array",
		// 							"defaultValue":[]any{
		// 							},
		// 							"metadata":map[string]any{
		// 								"description": "Insert Action groups ids to attach them to the below alert rules.",
		// 							},
		// 						},
		// 						"alertName":map[string]any{
		// 							"type": "string",
		// 							"defaultValue": "[concat('parameters('alertNamePrefix'), ' - ', parameters('clusterNameForPrometheus'))]",
		// 							"metadata":map[string]any{
		// 								"description": "Name of the alert rule",
		// 							},
		// 							"minLength": float64(1),
		// 						},
		// 						"alertNamePrefix":map[string]any{
		// 							"type": "string",
		// 							"defaultValue": "NodeRecordingRulesRuleGroup",
		// 							"metadata":map[string]any{
		// 								"description": "prefix of the alert rule name",
		// 							},
		// 							"minLength": float64(1),
		// 						},
		// 						"clusterNameForPrometheus":map[string]any{
		// 							"type": "string",
		// 						},
		// 						"location":map[string]any{
		// 							"type": "string",
		// 							"defaultValue": "eastus",
		// 						},
		// 						"targetResourceId":map[string]any{
		// 							"type": "string",
		// 							"defaultValue": "/subscriptions/2f00cc51-6809-498f-9ffc-48c42aff570d/resourceGroups/GenevaAlertRP-RunnerResources-eastus/providers/microsoft.monitor/accounts/alertsrp-eastus-pgms",
		// 						},
		// 						"targetResourceName":map[string]any{
		// 							"type": "string",
		// 							"defaultValue": "alertsrp-eastus-pgms",
		// 						},
		// 					},
		// 					Resources: []any{
		// 						map[string]any{
		// 							"name": "[parameters('alertName')]",
		// 							"type": "Microsoft.AlertsManagement/prometheusRuleGroups",
		// 							"apiVersion": "2021-07-22-preview",
		// 							"location": "[parameters('location')]",
		// 							"properties":map[string]any{
		// 								"description": "Node Recording Rules RuleGroup",
		// 								"clusterName": "[parameters('clusterNameForPrometheus')]",
		// 								"interval": "PT1M",
		// 								"rules":[]any{
		// 									map[string]any{
		// 										"expression": "count without (cpu, mode) (  node_cpu_seconds_total{job=\"node\",mode=\"idle\"})",
		// 										"record": "instance:node_num_cpu:sum",
		// 									},
		// 								},
		// 								"scopes": "[variables('scopes')]",
		// 							},
		// 					}},
		// 					Variables: map[string]any{
		// 						"copy":[]any{
		// 							map[string]any{
		// 								"name": "actionsForPrometheusRuleGroups",
		// 								"count": "[length(parameters('actionGroupIds'))]",
		// 								"input":map[string]any{
		// 									"actiongroupId": "[parameters('actionGroupIds')[copyIndex('actionsForPrometheusRuleGroups')]]",
		// 								},
		// 							},
		// 						},
		// 						"scopes": "[array(parameters('targetResourceId'))]",
		// 					},
		// 				},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("KubernetesReccordingRulesRuleGroup"),
		// 			Type: to.Ptr("Microsoft.AlertsManagement/alertRuleRecommendations"),
		// 			ID: to.Ptr("/subscriptions/2f00cc51-6809-498f-9ffc-48c42aff570d/resourceGroups/GenevaAlertRP-RunnerResources-eastus/providers/microsoft.monitor/accounts/alertsrp-eastus-pgms/providers/Microsoft.AlertsManagement/alertRuleRecommendations/KubernetesReccordingRulesRuleGroup"),
		// 			Properties: &armalertsmanagement.AlertRuleRecommendationProperties{
		// 				AlertRuleType: to.Ptr("Microsoft.AlertsManagement/prometheusRuleGroups"),
		// 				DisplayInformation: map[string]*string{
		// 					"ruleInfo": to.Ptr("Rule Information for second recording rule."),
		// 				},
		// 				RuleArmTemplate: &armalertsmanagement.RuleArmTemplate{
		// 					Schema: to.Ptr("https://schema.management.azure.com/schemas/2019-04-01/deploymentTemplate.json#"),
		// 					ContentVersion: to.Ptr("1.0.0.0"),
		// 					Parameters: map[string]any{
		// 						"actionGroupIds":map[string]any{
		// 							"type": "array",
		// 							"defaultValue":[]any{
		// 							},
		// 							"metadata":map[string]any{
		// 								"description": "Insert Action groups ids to attach them to the below alert rules.",
		// 							},
		// 						},
		// 						"alertName":map[string]any{
		// 							"type": "string",
		// 							"defaultValue": "[concat('parameters('alertNamePrefix'), ' - ', parameters('clusterNameForPrometheus'))]",
		// 							"metadata":map[string]any{
		// 								"description": "Name of the alert rule",
		// 							},
		// 							"minLength": float64(1),
		// 						},
		// 						"alertNamePrefix":map[string]any{
		// 							"type": "string",
		// 							"defaultValue": "KubernetesReccordingRulesRuleGroup",
		// 							"metadata":map[string]any{
		// 								"description": "prefix of the alert rule name",
		// 							},
		// 							"minLength": float64(1),
		// 						},
		// 						"clusterNameForPrometheus":map[string]any{
		// 							"type": "string",
		// 						},
		// 						"location":map[string]any{
		// 							"type": "string",
		// 							"defaultValue": "eastus",
		// 						},
		// 						"targetResourceId":map[string]any{
		// 							"type": "string",
		// 							"defaultValue": "/subscriptions/2f00cc51-6809-498f-9ffc-48c42aff570d/resourceGroups/GenevaAlertRP-RunnerResources-eastus/providers/microsoft.monitor/accounts/alertsrp-eastus-pgms",
		// 						},
		// 						"targetResourceName":map[string]any{
		// 							"type": "string",
		// 							"defaultValue": "alertsrp-eastus-pgms",
		// 						},
		// 					},
		// 					Resources: []any{
		// 						map[string]any{
		// 							"name": "[parameters('alertName')]",
		// 							"type": "Microsoft.AlertsManagement/prometheusRuleGroups",
		// 							"apiVersion": "2021-07-22-preview",
		// 							"location": "[parameters('location')]",
		// 							"properties":map[string]any{
		// 								"description": "Kubernetes Recording Rules RuleGroup",
		// 								"clusterName": "[parameters('clusterNameForPrometheus')]",
		// 								"interval": "PT1M",
		// 								"rules":[]any{
		// 									map[string]any{
		// 										"expression": "sum by (cluster, namespace, pod, container) (  irate(container_cpu_usage_seconds_total{job=\"cadvisor\", image!=\"\"}[5m])) * on (cluster, namespace, pod) group_left(node) topk by (cluster, namespace, pod) (  1, max by(cluster, namespace, pod, node) (kube_pod_info{node!=\"\"}))",
		// 										"record": "node_namespace_pod_container:container_cpu_usage_seconds_total:sum_irate",
		// 									},
		// 								},
		// 								"scopes": "[variables('scopes')]",
		// 							},
		// 					}},
		// 					Variables: map[string]any{
		// 						"copy":[]any{
		// 							map[string]any{
		// 								"name": "actionsForPrometheusRuleGroups",
		// 								"count": "[length(parameters('actionGroupIds'))]",
		// 								"input":map[string]any{
		// 									"actiongroupId": "[parameters('actionGroupIds')[copyIndex('actionsForPrometheusRuleGroups')]]",
		// 								},
		// 							},
		// 						},
		// 						"scopes": "[array(parameters('targetResourceId'))]",
		// 					},
		// 				},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("KubernetesAlert-DefaultAlert"),
		// 			Type: to.Ptr("Microsoft.AlertsManagement/alertRuleRecommendations"),
		// 			ID: to.Ptr("/subscriptions/2f00cc51-6809-498f-9ffc-48c42aff570d/resourceGroups/GenevaAlertRP-RunnerResources-eastus/providers/microsoft.monitor/accounts/alertsrp-eastus-pgms/providers/Microsoft.AlertsManagement/alertRuleRecommendations/KubernetesAlert-DefaultAlert"),
		// 			Properties: &armalertsmanagement.AlertRuleRecommendationProperties{
		// 				AlertRuleType: to.Ptr("Microsoft.AlertsManagement/prometheusRuleGroups"),
		// 				DisplayInformation: map[string]*string{
		// 					"ruleInfo": to.Ptr("Rule Information for alerting rule."),
		// 				},
		// 				RuleArmTemplate: &armalertsmanagement.RuleArmTemplate{
		// 					Schema: to.Ptr("https://schema.management.azure.com/schemas/2019-04-01/deploymentTemplate.json#"),
		// 					ContentVersion: to.Ptr("1.0.0.0"),
		// 					Parameters: map[string]any{
		// 						"actionGroupIds":map[string]any{
		// 							"type": "array",
		// 							"defaultValue":[]any{
		// 							},
		// 							"metadata":map[string]any{
		// 								"description": "Insert Action groups ids to attach them to the below alert rules.",
		// 							},
		// 						},
		// 						"alertName":map[string]any{
		// 							"type": "string",
		// 							"defaultValue": "[concat('parameters('alertNamePrefix'), ' - ', parameters('clusterNameForPrometheus'))]",
		// 							"metadata":map[string]any{
		// 								"description": "Name of the alert rule",
		// 							},
		// 							"minLength": float64(1),
		// 						},
		// 						"alertNamePrefix":map[string]any{
		// 							"type": "string",
		// 							"defaultValue": "KubernetesAlert-DefaultAlerts",
		// 							"metadata":map[string]any{
		// 								"description": "prefix of the alert rule name",
		// 							},
		// 							"minLength": float64(1),
		// 						},
		// 						"clusterNameForPrometheus":map[string]any{
		// 							"type": "string",
		// 						},
		// 						"location":map[string]any{
		// 							"type": "string",
		// 							"defaultValue": "eastus",
		// 						},
		// 						"targetResourceId":map[string]any{
		// 							"type": "string",
		// 							"defaultValue": "/subscriptions/2f00cc51-6809-498f-9ffc-48c42aff570d/resourceGroups/GenevaAlertRP-RunnerResources-eastus/providers/microsoft.monitor/accounts/alertsrp-eastus-pgms",
		// 						},
		// 						"targetResourceName":map[string]any{
		// 							"type": "string",
		// 							"defaultValue": "alertsrp-eastus-pgms",
		// 						},
		// 					},
		// 					Resources: []any{
		// 						map[string]any{
		// 							"name": "[parameters('alertName')]",
		// 							"type": "Microsoft.AlertsManagement/prometheusRuleGroups",
		// 							"apiVersion": "2021-07-22-preview",
		// 							"location": "[parameters('location')]",
		// 							"properties":map[string]any{
		// 								"description": "Kubernetes Alert RuleGroup-DefaultAlerts",
		// 								"clusterName": "[parameters('clusterNameForPrometheus')]",
		// 								"interval": "PT1M",
		// 								"rules":[]any{
		// 									map[string]any{
		// 										"Severity": float64(3),
		// 										"actions": "[variables('actionsForPrometheusRuleGroups')]",
		// 										"alert": "KubePodCrashLooping",
		// 										"expression": "max_over_time(kube_pod_container_status_waiting_reason{reason=\"CrashLoopBackOff\", job=\"kube-state-metrics\"}[5m]) >= 1",
		// 										"for": "PT15M",
		// 										"labels":map[string]any{
		// 											"severity": "warning",
		// 										},
		// 									},
		// 									map[string]any{
		// 										"Severity": float64(3),
		// 										"actions": "[variables('actionsForPrometheusRuleGroups')]",
		// 										"alert": "KubePodNotReady",
		// 										"expression": "sum by (namespace, pod, cluster) (  max by(namespace, pod, cluster) (    kube_pod_status_phase{job=\"kube-state-metrics\", phase=~\"Pending|Unknown\"}  ) * on(namespace, pod, cluster) group_left(owner_kind) topk by(namespace, pod, cluster) (    1, max by(namespace, pod, owner_kind, cluster) (kube_pod_owner{owner_kind!=\"Job\"})  )) > 0",
		// 										"for": "PT15M",
		// 										"labels":map[string]any{
		// 											"severity": "warning",
		// 										},
		// 									},
		// 								},
		// 								"scopes": "[variables('scopes')]",
		// 							},
		// 					}},
		// 					Variables: map[string]any{
		// 						"copy":[]any{
		// 							map[string]any{
		// 								"name": "actionsForPrometheusRuleGroups",
		// 								"count": "[length(parameters('actionGroupIds'))]",
		// 								"input":map[string]any{
		// 									"actiongroupId": "[parameters('actionGroupIds')[copyIndex('actionsForPrometheusRuleGroups')]]",
		// 								},
		// 							},
		// 						},
		// 						"scopes": "[array(parameters('targetResourceId'))]",
		// 					},
		// 				},
		// 			},
		// 	}},
		// }
	}
}
