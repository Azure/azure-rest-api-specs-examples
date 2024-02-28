package armalertsmanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/alertsmanagement/armalertsmanagement"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/6d2438481021a94793b07b226df06d5f3c61d51d/specification/alertsmanagement/resource-manager/Microsoft.AlertsManagement/preview/2023-01-01-preview/examples/AlertRuleRecommendations_GetBySubscription_VM.json
func ExampleAlertRuleRecommendationsClient_NewListByTargetTypePager_listAlertRuleRecommendationsForVirtualMachinesAtSubscriptionLevel() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armalertsmanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAlertRuleRecommendationsClient().NewListByTargetTypePager("microsoft.compute/virtualmachines", nil)
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
		// 			Name: to.Ptr("Percentage CPU"),
		// 			Type: to.Ptr("Microsoft.AlertsManagement/alertRuleRecommendations"),
		// 			ID: to.Ptr("/subscriptions/2f00cc51-6809-498f-9ffc-48c42aff570d/providers/Microsoft.AlertsManagement/alertRuleRecommendations/Percentage CPU"),
		// 			Properties: &armalertsmanagement.AlertRuleRecommendationProperties{
		// 				AlertRuleType: to.Ptr("Microsoft.Insights/metricAlerts"),
		// 				DisplayInformation: map[string]*string{
		// 					"displayUnits": to.Ptr("Percentage"),
		// 					"infoBallonLink": to.Ptr("Rule1 InfoBalloon Link"),
		// 					"infoBallonMessage": to.Ptr("Rule1 InfoBalloon Message"),
		// 					"metricDisplayName": to.Ptr("Percentage CPU"),
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
		// 							"defaultValue": "[concat('parameters('alertNamePrefix'), ' - ', parameters('targetResourceName'))]",
		// 							"metadata":map[string]any{
		// 								"description": "Name of the alert rule",
		// 							},
		// 							"minLength": float64(1),
		// 						},
		// 						"alertNamePrefix":map[string]any{
		// 							"type": "string",
		// 							"defaultValue": "Percentage CPU",
		// 							"metadata":map[string]any{
		// 								"description": "prefix of the alert rule name",
		// 							},
		// 							"minLength": float64(1),
		// 						},
		// 						"alertSeverity":map[string]any{
		// 							"type": "int",
		// 							"allowedValues":[]any{
		// 								float64(0),
		// 								float64(1),
		// 								float64(2),
		// 								float64(3),
		// 								float64(4),
		// 							},
		// 							"defaultValue": float64(3),
		// 							"metadata":map[string]any{
		// 								"description": "Severity of alert {0,1,2,3,4}",
		// 							},
		// 						},
		// 						"targetResourceId":map[string]any{
		// 							"type": "string",
		// 						},
		// 						"targetResourceName":map[string]any{
		// 							"type": "string",
		// 						},
		// 						"threshold":map[string]any{
		// 							"type": "int",
		// 							"defaultValue": float64(80),
		// 							"metadata":map[string]any{
		// 								"description": "The threshold value at which the alert is activated.",
		// 							},
		// 						},
		// 					},
		// 					Resources: []any{
		// 						map[string]any{
		// 							"name": "[parameters('alertName')]",
		// 							"type": "Microsoft.Insights/metricAlerts",
		// 							"apiVersion": "2018-03-01",
		// 							"location": "global",
		// 							"properties":map[string]any{
		// 								"description": "Percentage CPU is greater than 80 %",
		// 								"actions": "[variables('actionsForMetricAlerts')]",
		// 								"criteria":map[string]any{
		// 									"allOf":[]any{
		// 										map[string]any{
		// 											"name": "Metric1",
		// 											"metricName": "Percentage CPU",
		// 											"operator": "GreaterThan",
		// 											"threshold": "[parameters('threshold')]",
		// 											"timeAggregation": "Average",
		// 										},
		// 									},
		// 									"odata.type": "Microsoft.Azure.Monitor.MultipleResourceMultipleMetricCriteria",
		// 								},
		// 								"enabled": true,
		// 								"evaluationFrequency": "PT5M",
		// 								"scopes": "[variables('scopes')]",
		// 								"severity": "[parameters('alertSeverity')]",
		// 								"windowSize": "PT5M",
		// 							},
		// 					}},
		// 					Variables: map[string]any{
		// 						"copy":[]any{
		// 							map[string]any{
		// 								"name": "actionsForMetricAlerts",
		// 								"count": "[length(parameters('actionGroupIds'))]",
		// 								"input":map[string]any{
		// 									"actiongroupId": "[parameters('actionGroupIds')[copyIndex('actionsForMetricAlerts')]]",
		// 								},
		// 							},
		// 						},
		// 						"scopes": "[array(parameters('targetResourceId'))]",
		// 					},
		// 				},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Available Memory Bytes"),
		// 			Type: to.Ptr("Microsoft.AlertsManagement/alertRuleRecommendations"),
		// 			ID: to.Ptr("/subscriptions/2f00cc51-6809-498f-9ffc-48c42aff570d/providers/Microsoft.AlertsManagement/alertRuleRecommendations/Available Memory Bytes"),
		// 			Properties: &armalertsmanagement.AlertRuleRecommendationProperties{
		// 				AlertRuleType: to.Ptr("Microsoft.Insights/metricAlerts"),
		// 				DisplayInformation: map[string]*string{
		// 					"displayUnits": to.Ptr("Gigabytes"),
		// 					"infoBallonLink": to.Ptr("Rule2 InfoBalloon Link"),
		// 					"infoBallonMessage": to.Ptr("Rule2 InfoBalloon Message"),
		// 					"metricDisplayName": to.Ptr("Available Memory Bytes"),
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
		// 							"defaultValue": "[concat('parameters('alertNamePrefix'), ' - ', parameters('targetResourceName'))]",
		// 							"metadata":map[string]any{
		// 								"description": "Name of the alert rule",
		// 							},
		// 							"minLength": float64(1),
		// 						},
		// 						"alertNamePrefix":map[string]any{
		// 							"type": "string",
		// 							"defaultValue": "Available Memory Bytes",
		// 							"metadata":map[string]any{
		// 								"description": "prefix of the alert rule name",
		// 							},
		// 							"minLength": float64(1),
		// 						},
		// 						"alertSeverity":map[string]any{
		// 							"type": "int",
		// 							"allowedValues":[]any{
		// 								float64(0),
		// 								float64(1),
		// 								float64(2),
		// 								float64(3),
		// 								float64(4),
		// 							},
		// 							"defaultValue": float64(3),
		// 							"metadata":map[string]any{
		// 								"description": "Severity of alert {0,1,2,3,4}",
		// 							},
		// 						},
		// 						"targetResourceId":map[string]any{
		// 							"type": "string",
		// 						},
		// 						"targetResourceName":map[string]any{
		// 							"type": "string",
		// 						},
		// 						"threshold":map[string]any{
		// 							"type": "int",
		// 							"defaultValue": float64(1000000000),
		// 							"metadata":map[string]any{
		// 								"description": "The threshold value at which the alert is activated.",
		// 							},
		// 						},
		// 					},
		// 					Resources: []any{
		// 						map[string]any{
		// 							"name": "[parameters('alertName')]",
		// 							"type": "Microsoft.Insights/metricAlerts",
		// 							"apiVersion": "2018-03-01",
		// 							"location": "global",
		// 							"properties":map[string]any{
		// 								"description": "Available Memory Bytes is less than 1 GB",
		// 								"actions": "[variables('actionsForMetricAlerts')]",
		// 								"criteria":map[string]any{
		// 									"allOf":[]any{
		// 										map[string]any{
		// 											"name": "Metric1",
		// 											"metricName": "Available Memory Bytes",
		// 											"operator": "LessThan",
		// 											"threshold": "[parameters('threshold')]",
		// 											"timeAggregation": "Average",
		// 										},
		// 									},
		// 									"odata.type": "Microsoft.Azure.Monitor.MultipleResourceMultipleMetricCriteria",
		// 								},
		// 								"enabled": true,
		// 								"evaluationFrequency": "PT5M",
		// 								"scopes": "[variables('scopes')]",
		// 								"severity": "[parameters('alertSeverity')]",
		// 								"windowSize": "PT5M",
		// 							},
		// 					}},
		// 					Variables: map[string]any{
		// 						"copy":[]any{
		// 							map[string]any{
		// 								"name": "actionsForMetricAlerts",
		// 								"count": "[length(parameters('actionGroupIds'))]",
		// 								"input":map[string]any{
		// 									"actiongroupId": "[parameters('actionGroupIds')[copyIndex('actionsForMetricAlerts')]]",
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
