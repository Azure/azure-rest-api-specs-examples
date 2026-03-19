package armalertrulerecommendations_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/alertrulerecommendations/armalertrulerecommendations"
)

// Generated from example definition: 2023-08-01-preview/AlertRuleRecommendations_GetBySubscription_MAC.json
func ExampleClient_NewListByTargetTypePager_listAlertRuleRecommendationsForMonitoringAccountsAtSubscriptionLevel() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armalertrulerecommendations.NewClientFactory("2f00cc51-6809-498f-9ffc-48c42aff570d", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewClient().NewListByTargetTypePager("microsoft.monitor/accounts", nil)
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
		// page = armalertrulerecommendations.ClientListByTargetTypeResponse{
		// 	ListResponse: armalertrulerecommendations.ListResponse{
		// 		Value: []*armalertrulerecommendations.AlertRuleRecommendationResource{
		// 			{
		// 				Name: to.Ptr("NodeRecordingRulesRuleGroup"),
		// 				Type: to.Ptr("Microsoft.AlertsManagement/alertRuleRecommendations"),
		// 				ID: to.Ptr("/subscriptions/2f00cc51-6809-498f-9ffc-48c42aff570d/providers/Microsoft.AlertsManagement/alertRuleRecommendations/NodeRecordingRulesRuleGroup"),
		// 				Properties: &armalertrulerecommendations.AlertRuleRecommendationProperties{
		// 					AlertRuleType: to.Ptr("Microsoft.AlertsManagement/prometheusRuleGroups"),
		// 					Category: to.Ptr("MyCategory1"),
		// 					DisplayInformation: map[string]*string{
		// 						"ruleInfo": to.Ptr("Rule Information for first recording rule."),
		// 					},
		// 					RuleArmTemplate: &armalertrulerecommendations.RuleArmTemplate{
		// 						Schema: to.Ptr("https://schema.management.azure.com/schemas/2019-04-01/deploymentTemplate.json#"),
		// 						ContentVersion: to.Ptr("1.0.0.0"),
		// 						Parameters: map[string]any{
		// 							"actionGroupIds": map[string]any{
		// 								"type": "array",
		// 								"defaultValue": []any{
		// 								},
		// 								"metadata": map[string]any{
		// 									"description": "Insert Action groups ids to attach them to the below alert rules.",
		// 								},
		// 							},
		// 							"alertName": map[string]any{
		// 								"type": "string",
		// 								"defaultValue": "[concat('parameters('alertNamePrefix'), ' - ', parameters('clusterNameForPrometheus'))]",
		// 								"metadata": map[string]any{
		// 									"description": "Name of the alert rule",
		// 								},
		// 								"minLength": 1,
		// 							},
		// 							"alertNamePrefix": map[string]any{
		// 								"type": "string",
		// 								"defaultValue": "NodeRecordingRulesRuleGroup",
		// 								"metadata": map[string]any{
		// 									"description": "prefix of the alert rule name",
		// 								},
		// 								"minLength": 1,
		// 							},
		// 							"clusterNameForPrometheus": map[string]any{
		// 								"type": "string",
		// 							},
		// 							"location": map[string]any{
		// 								"type": "string",
		// 							},
		// 							"targetResourceId": map[string]any{
		// 								"type": "string",
		// 							},
		// 							"targetResourceName": map[string]any{
		// 								"type": "string",
		// 							},
		// 						},
		// 						Resources: []any{
		// 							map[string]any{
		// 								"name": "[parameters('alertName')]",
		// 								"type": "Microsoft.AlertsManagement/prometheusRuleGroups",
		// 								"apiVersion": "2021-07-22-preview",
		// 								"location": "[parameters('location')]",
		// 								"properties": map[string]any{
		// 									"description": "Node Recording Rules RuleGroup",
		// 									"clusterName": "[parameters('clusterNameForPrometheus')]",
		// 									"interval": "PT1M",
		// 									"rules": []any{
		// 										map[string]any{
		// 											"expression": "count without (cpu, mode) (  node_cpu_seconds_total{job=\"node\",mode=\"idle\"})",
		// 											"record": "instance:node_num_cpu:sum",
		// 										},
		// 									},
		// 									"scopes": "[variables('scopes')]",
		// 								},
		// 							},
		// 						},
		// 						Variables: map[string]any{
		// 							"copy": []any{
		// 								map[string]any{
		// 									"name": "actionsForPrometheusRuleGroups",
		// 									"count": "[length(parameters('actionGroupIds'))]",
		// 									"input": map[string]any{
		// 										"actiongroupId": "[parameters('actionGroupIds')[copyIndex('actionsForPrometheusRuleGroups')]]",
		// 									},
		// 								},
		// 							},
		// 							"scopes": "[array(parameters('targetResourceId'))]",
		// 						},
		// 					},
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("KubernetesReccordingRulesRuleGroup"),
		// 				Type: to.Ptr("Microsoft.AlertsManagement/alertRuleRecommendations"),
		// 				ID: to.Ptr("/subscriptions/2f00cc51-6809-498f-9ffc-48c42aff570d/providers/Microsoft.AlertsManagement/alertRuleRecommendations/KubernetesReccordingRulesRuleGroup"),
		// 				Properties: &armalertrulerecommendations.AlertRuleRecommendationProperties{
		// 					AlertRuleType: to.Ptr("Microsoft.AlertsManagement/prometheusRuleGroups"),
		// 					Category: to.Ptr("MyCategory1"),
		// 					DisplayInformation: map[string]*string{
		// 						"ruleInfo": to.Ptr("Rule Information for second recording rule."),
		// 					},
		// 					RuleArmTemplate: &armalertrulerecommendations.RuleArmTemplate{
		// 						Schema: to.Ptr("https://schema.management.azure.com/schemas/2019-04-01/deploymentTemplate.json#"),
		// 						ContentVersion: to.Ptr("1.0.0.0"),
		// 						Parameters: map[string]any{
		// 							"actionGroupIds": map[string]any{
		// 								"type": "array",
		// 								"defaultValue": []any{
		// 								},
		// 								"metadata": map[string]any{
		// 									"description": "Insert Action groups ids to attach them to the below alert rules.",
		// 								},
		// 							},
		// 							"alertName": map[string]any{
		// 								"type": "string",
		// 								"defaultValue": "[concat('parameters('alertNamePrefix'), ' - ', parameters('clusterNameForPrometheus'))]",
		// 								"metadata": map[string]any{
		// 									"description": "Name of the alert rule",
		// 								},
		// 								"minLength": 1,
		// 							},
		// 							"alertNamePrefix": map[string]any{
		// 								"type": "string",
		// 								"defaultValue": "KubernetesReccordingRulesRuleGroup",
		// 								"metadata": map[string]any{
		// 									"description": "prefix of the alert rule name",
		// 								},
		// 								"minLength": 1,
		// 							},
		// 							"clusterNameForPrometheus": map[string]any{
		// 								"type": "string",
		// 							},
		// 							"location": map[string]any{
		// 								"type": "string",
		// 							},
		// 							"targetResourceId": map[string]any{
		// 								"type": "string",
		// 							},
		// 							"targetResourceName": map[string]any{
		// 								"type": "string",
		// 							},
		// 						},
		// 						Resources: []any{
		// 							map[string]any{
		// 								"name": "[parameters('alertName')]",
		// 								"type": "Microsoft.AlertsManagement/prometheusRuleGroups",
		// 								"apiVersion": "2021-07-22-preview",
		// 								"location": "[parameters('location')]",
		// 								"properties": map[string]any{
		// 									"description": "Kubernetes Recording Rules RuleGroup",
		// 									"clusterName": "[parameters('clusterNameForPrometheus')]",
		// 									"interval": "PT1M",
		// 									"rules": []any{
		// 										map[string]any{
		// 											"expression": "sum by (cluster, namespace, pod, container) (  irate(container_cpu_usage_seconds_total{job=\"cadvisor\", image!=\"\"}[5m])) * on (cluster, namespace, pod) group_left(node) topk by (cluster, namespace, pod) (  1, max by(cluster, namespace, pod, node) (kube_pod_info{node!=\"\"}))",
		// 											"record": "node_namespace_pod_container:container_cpu_usage_seconds_total:sum_irate",
		// 										},
		// 									},
		// 									"scopes": "[variables('scopes')]",
		// 								},
		// 							},
		// 						},
		// 						Variables: map[string]any{
		// 							"copy": []any{
		// 								map[string]any{
		// 									"name": "actionsForPrometheusRuleGroups",
		// 									"count": "[length(parameters('actionGroupIds'))]",
		// 									"input": map[string]any{
		// 										"actiongroupId": "[parameters('actionGroupIds')[copyIndex('actionsForPrometheusRuleGroups')]]",
		// 									},
		// 								},
		// 							},
		// 							"scopes": "[array(parameters('targetResourceId'))]",
		// 						},
		// 					},
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("KubernetesAlert-DefaultAlerts"),
		// 				Type: to.Ptr("Microsoft.AlertsManagement/alertRuleRecommendations"),
		// 				ID: to.Ptr("/subscriptions/2f00cc51-6809-498f-9ffc-48c42aff570d/providers/Microsoft.AlertsManagement/alertRuleRecommendations/KubernetesAlert-DefaultAlerts"),
		// 				Properties: &armalertrulerecommendations.AlertRuleRecommendationProperties{
		// 					AlertRuleType: to.Ptr("Microsoft.AlertsManagement/prometheusRuleGroups"),
		// 					Category: to.Ptr("MyCategory2"),
		// 					DisplayInformation: map[string]*string{
		// 						"ruleInfo": to.Ptr("Rule Information for alerting rule."),
		// 					},
		// 					RuleArmTemplate: &armalertrulerecommendations.RuleArmTemplate{
		// 						Schema: to.Ptr("https://schema.management.azure.com/schemas/2019-04-01/deploymentTemplate.json#"),
		// 						ContentVersion: to.Ptr("1.0.0.0"),
		// 						Parameters: map[string]any{
		// 							"actionGroupIds": map[string]any{
		// 								"type": "array",
		// 								"defaultValue": []any{
		// 								},
		// 								"metadata": map[string]any{
		// 									"description": "Insert Action groups ids to attach them to the below alert rules.",
		// 								},
		// 							},
		// 							"alertName": map[string]any{
		// 								"type": "string",
		// 								"defaultValue": "[concat('parameters('alertNamePrefix'), ' - ', parameters('clusterNameForPrometheus'))]",
		// 								"metadata": map[string]any{
		// 									"description": "Name of the alert rule",
		// 								},
		// 								"minLength": 1,
		// 							},
		// 							"alertNamePrefix": map[string]any{
		// 								"type": "string",
		// 								"defaultValue": "KubernetesAlert-DefaultAlerts",
		// 								"metadata": map[string]any{
		// 									"description": "prefix of the alert rule name",
		// 								},
		// 								"minLength": 1,
		// 							},
		// 							"clusterNameForPrometheus": map[string]any{
		// 								"type": "string",
		// 							},
		// 							"location": map[string]any{
		// 								"type": "string",
		// 							},
		// 							"targetResourceId": map[string]any{
		// 								"type": "string",
		// 							},
		// 							"targetResourceName": map[string]any{
		// 								"type": "string",
		// 							},
		// 						},
		// 						Resources: []any{
		// 							map[string]any{
		// 								"name": "[parameters('alertName')]",
		// 								"type": "Microsoft.AlertsManagement/prometheusRuleGroups",
		// 								"apiVersion": "2021-07-22-preview",
		// 								"location": "[parameters('location')]",
		// 								"properties": map[string]any{
		// 									"description": "Kubernetes Alert RuleGroup-DefaultAlerts",
		// 									"clusterName": "[parameters('clusterNameForPrometheus')]",
		// 									"interval": "PT1M",
		// 									"rules": []any{
		// 										map[string]any{
		// 											"Severity": 3,
		// 											"actions": "[variables('actionsForPrometheusRuleGroups')]",
		// 											"alert": "KubePodCrashLooping",
		// 											"expression": "max_over_time(kube_pod_container_status_waiting_reason{reason=\"CrashLoopBackOff\", job=\"kube-state-metrics\"}[5m]) >= 1",
		// 											"for": "PT15M",
		// 											"labels": map[string]any{
		// 												"severity": "warning",
		// 											},
		// 										},
		// 										map[string]any{
		// 											"Severity": 3,
		// 											"actions": "[variables('actionsForPrometheusRuleGroups')]",
		// 											"alert": "KubePodNotReady",
		// 											"expression": "sum by (namespace, pod, cluster) (  max by(namespace, pod, cluster) (    kube_pod_status_phase{job=\"kube-state-metrics\", phase=~\"Pending|Unknown\"}  ) * on(namespace, pod, cluster) group_left(owner_kind) topk by(namespace, pod, cluster) (    1, max by(namespace, pod, owner_kind, cluster) (kube_pod_owner{owner_kind!=\"Job\"})  )) > 0",
		// 											"for": "PT15M",
		// 											"labels": map[string]any{
		// 												"severity": "warning",
		// 											},
		// 										},
		// 									},
		// 									"scopes": "[variables('scopes')]",
		// 								},
		// 							},
		// 						},
		// 						Variables: map[string]any{
		// 							"copy": []any{
		// 								map[string]any{
		// 									"name": "actionsForPrometheusRuleGroups",
		// 									"count": "[length(parameters('actionGroupIds'))]",
		// 									"input": map[string]any{
		// 										"actiongroupId": "[parameters('actionGroupIds')[copyIndex('actionsForPrometheusRuleGroups')]]",
		// 									},
		// 								},
		// 							},
		// 							"scopes": "[array(parameters('targetResourceId'))]",
		// 						},
		// 					},
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
