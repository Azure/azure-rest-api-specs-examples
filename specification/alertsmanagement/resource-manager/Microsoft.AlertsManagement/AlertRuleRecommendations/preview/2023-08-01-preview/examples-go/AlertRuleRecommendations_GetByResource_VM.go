package armalertrulerecommendations_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/alertrulerecommendations/armalertrulerecommendations"
)

// Generated from example definition: 2023-08-01-preview/AlertRuleRecommendations_GetByResource_VM.json
func ExampleClient_NewListByResourcePager_listAlertRuleRecommendationsForVirtualMachinesAtResourceLevel() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armalertrulerecommendations.NewClientFactory("<subscriptionID>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewClient().NewListByResourcePager("subscriptions/2f00cc51-6809-498f-9ffc-48c42aff570d/resourcegroups/test/providers/Microsoft.Compute/virtualMachines/testMachineCanBeSafelyDeleted", nil)
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
		// page = armalertrulerecommendations.ClientListByResourceResponse{
		// 	ListResponse: armalertrulerecommendations.ListResponse{
		// 		Value: []*armalertrulerecommendations.AlertRuleRecommendationResource{
		// 			{
		// 				Name: to.Ptr("Percentage CPU"),
		// 				Type: to.Ptr("Microsoft.AlertsManagement/alertRuleRecommendations"),
		// 				ID: to.Ptr("/subscriptions/2f00cc51-6809-498f-9ffc-48c42aff570d/resourcegroups/test/providers/Microsoft.Compute/virtualMachines/testMachineCanBeSafelyDeleted/providers/Microsoft.AlertsManagement/alertRuleRecommendations/Percentage CPU"),
		// 				Properties: &armalertrulerecommendations.AlertRuleRecommendationProperties{
		// 					AlertRuleType: to.Ptr("Microsoft.Insights/metricAlerts"),
		// 					DisplayInformation: map[string]*string{
		// 						"displayUnits": to.Ptr("Percentage"),
		// 						"infoBallonLink": to.Ptr("Rule1 InfoBalloon Link"),
		// 						"infoBallonMessage": to.Ptr("Rule1 InfoBalloon Message"),
		// 						"metricDisplayName": to.Ptr("Percentage CPU"),
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
		// 								"defaultValue": "[concat('parameters('alertNamePrefix'), ' - ', parameters('targetResourceName'))]",
		// 								"metadata": map[string]any{
		// 									"description": "Name of the alert rule",
		// 								},
		// 								"minLength": 1,
		// 							},
		// 							"alertNamePrefix": map[string]any{
		// 								"type": "string",
		// 								"defaultValue": "Percentage CPU",
		// 								"metadata": map[string]any{
		// 									"description": "prefix of the alert rule name",
		// 								},
		// 								"minLength": 1,
		// 							},
		// 							"alertSeverity": map[string]any{
		// 								"type": "int",
		// 								"allowedValues": []any{
		// 									0,
		// 									1,
		// 									2,
		// 									3,
		// 									4,
		// 								},
		// 								"defaultValue": 3,
		// 								"metadata": map[string]any{
		// 									"description": "Severity of alert {0,1,2,3,4}",
		// 								},
		// 							},
		// 							"targetResourceId": map[string]any{
		// 								"type": "string",
		// 								"defaultValue": "/subscriptions/2f00cc51-6809-498f-9ffc-48c42aff570d/resourcegroups/test/providers/Microsoft.Compute/virtualMachines/testMachineCanBeSafelyDeleted",
		// 							},
		// 							"targetResourceName": map[string]any{
		// 								"type": "string",
		// 								"defaultValue": "testmachinecanbesafelydeleted",
		// 							},
		// 							"threshold": map[string]any{
		// 								"type": "int",
		// 								"defaultValue": 80,
		// 								"metadata": map[string]any{
		// 									"description": "The threshold value at which the alert is activated.",
		// 								},
		// 							},
		// 						},
		// 						Resources: []any{
		// 							map[string]any{
		// 								"name": "[parameters('alertName')]",
		// 								"type": "Microsoft.Insights/metricAlerts",
		// 								"apiVersion": "2018-03-01",
		// 								"location": "global",
		// 								"properties": map[string]any{
		// 									"description": "Percentage CPU is greater than 80 %",
		// 									"actions": "[variables('actionsForMetricAlerts')]",
		// 									"criteria": map[string]any{
		// 										"allOf": []any{
		// 											map[string]any{
		// 												"name": "Metric1",
		// 												"metricName": "Percentage CPU",
		// 												"operator": "GreaterThan",
		// 												"threshold": "[parameters('threshold')]",
		// 												"timeAggregation": "Average",
		// 											},
		// 										},
		// 										"odata.type": "Microsoft.Azure.Monitor.MultipleResourceMultipleMetricCriteria",
		// 									},
		// 									"enabled": true,
		// 									"evaluationFrequency": "PT5M",
		// 									"scopes": "[variables('scopes')]",
		// 									"severity": "[parameters('alertSeverity')]",
		// 									"windowSize": "PT5M",
		// 								},
		// 							},
		// 						},
		// 						Variables: map[string]any{
		// 							"copy": []any{
		// 								map[string]any{
		// 									"name": "actionsForMetricAlerts",
		// 									"count": "[length(parameters('actionGroupIds'))]",
		// 									"input": map[string]any{
		// 										"actiongroupId": "[parameters('actionGroupIds')[copyIndex('actionsForMetricAlerts')]]",
		// 									},
		// 								},
		// 							},
		// 							"scopes": "[array(parameters('targetResourceId'))]",
		// 						},
		// 					},
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Available Memory Bytes"),
		// 				Type: to.Ptr("Microsoft.AlertsManagement/alertRuleRecommendations"),
		// 				ID: to.Ptr("/subscriptions/2f00cc51-6809-498f-9ffc-48c42aff570d/resourcegroups/test/providers/Microsoft.Compute/virtualMachines/testMachineCanBeSafelyDeleted/providers/Microsoft.AlertsManagement/alertRuleRecommendations/Available Memory Bytes"),
		// 				Properties: &armalertrulerecommendations.AlertRuleRecommendationProperties{
		// 					AlertRuleType: to.Ptr("Microsoft.Insights/metricAlerts"),
		// 					DisplayInformation: map[string]*string{
		// 						"displayUnits": to.Ptr("Gigabytes"),
		// 						"infoBallonLink": to.Ptr("Rule2 InfoBalloon Link"),
		// 						"infoBallonMessage": to.Ptr("Rule2 InfoBalloon Message"),
		// 						"metricDisplayName": to.Ptr("Available Memory Bytes"),
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
		// 								"defaultValue": "[concat('parameters('alertNamePrefix'), ' - ', parameters('targetResourceName'))]",
		// 								"metadata": map[string]any{
		// 									"description": "Name of the alert rule",
		// 								},
		// 								"minLength": 1,
		// 							},
		// 							"alertNamePrefix": map[string]any{
		// 								"type": "string",
		// 								"defaultValue": "Available Memory Bytes",
		// 								"metadata": map[string]any{
		// 									"description": "prefix of the alert rule name",
		// 								},
		// 								"minLength": 1,
		// 							},
		// 							"alertSeverity": map[string]any{
		// 								"type": "int",
		// 								"allowedValues": []any{
		// 									0,
		// 									1,
		// 									2,
		// 									3,
		// 									4,
		// 								},
		// 								"defaultValue": 3,
		// 								"metadata": map[string]any{
		// 									"description": "Severity of alert {0,1,2,3,4}",
		// 								},
		// 							},
		// 							"targetResourceId": map[string]any{
		// 								"type": "string",
		// 								"defaultValue": "/subscriptions/2f00cc51-6809-498f-9ffc-48c42aff570d/resourcegroups/test/providers/Microsoft.Compute/virtualMachines/testMachineCanBeSafelyDeleted",
		// 							},
		// 							"targetResourceName": map[string]any{
		// 								"type": "string",
		// 								"defaultValue": "testmachinecanbesafelydeleted",
		// 							},
		// 							"threshold": map[string]any{
		// 								"type": "int",
		// 								"defaultValue": 1000000000,
		// 								"metadata": map[string]any{
		// 									"description": "The threshold value at which the alert is activated.",
		// 								},
		// 							},
		// 						},
		// 						Resources: []any{
		// 							map[string]any{
		// 								"name": "[parameters('alertName')]",
		// 								"type": "Microsoft.Insights/metricAlerts",
		// 								"apiVersion": "2018-03-01",
		// 								"location": "global",
		// 								"properties": map[string]any{
		// 									"description": "Available Memory Bytes is less than 1 GB",
		// 									"actions": "[variables('actionsForMetricAlerts')]",
		// 									"criteria": map[string]any{
		// 										"allOf": []any{
		// 											map[string]any{
		// 												"name": "Metric1",
		// 												"metricName": "Available Memory Bytes",
		// 												"operator": "LessThan",
		// 												"threshold": "[parameters('threshold')]",
		// 												"timeAggregation": "Average",
		// 											},
		// 										},
		// 										"odata.type": "Microsoft.Azure.Monitor.MultipleResourceMultipleMetricCriteria",
		// 									},
		// 									"enabled": true,
		// 									"evaluationFrequency": "PT5M",
		// 									"scopes": "[variables('scopes')]",
		// 									"severity": "[parameters('alertSeverity')]",
		// 									"windowSize": "PT5M",
		// 								},
		// 							},
		// 						},
		// 						Variables: map[string]any{
		// 							"copy": []any{
		// 								map[string]any{
		// 									"name": "actionsForMetricAlerts",
		// 									"count": "[length(parameters('actionGroupIds'))]",
		// 									"input": map[string]any{
		// 										"actiongroupId": "[parameters('actionGroupIds')[copyIndex('actionsForMetricAlerts')]]",
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
