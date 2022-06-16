package armsecurityinsight_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsight/armsecurityinsight"
)

// x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2021-09-01-preview/examples/automationRules/CreateAutomationRule.json
func ExampleAutomationRulesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsecurityinsight.NewAutomationRulesClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		"<automation-rule-id>",
		armsecurityinsight.AutomationRule{
			Etag: to.StringPtr("<etag>"),
			Properties: &armsecurityinsight.AutomationRuleProperties{
				Actions: []armsecurityinsight.AutomationRuleActionClassification{
					&armsecurityinsight.AutomationRuleModifyPropertiesAction{
						ActionType: armsecurityinsight.AutomationRuleActionType("ModifyProperties").ToPtr(),
						Order:      to.Int32Ptr(1),
						ActionConfiguration: &armsecurityinsight.AutomationRuleModifyPropertiesActionConfiguration{
							Severity: armsecurityinsight.IncidentSeverity("High").ToPtr(),
						},
					},
					&armsecurityinsight.AutomationRuleRunPlaybookAction{
						ActionType: armsecurityinsight.AutomationRuleActionType("RunPlaybook").ToPtr(),
						Order:      to.Int32Ptr(2),
						ActionConfiguration: &armsecurityinsight.AutomationRuleRunPlaybookActionConfiguration{
							LogicAppResourceID: to.StringPtr("<logic-app-resource-id>"),
							TenantID:           to.StringPtr("<tenant-id>"),
						},
					}},
				DisplayName: to.StringPtr("<display-name>"),
				Order:       to.Int32Ptr(1),
				TriggeringLogic: &armsecurityinsight.AutomationRuleTriggeringLogic{
					Conditions: []armsecurityinsight.AutomationRuleConditionClassification{
						&armsecurityinsight.AutomationRulePropertyValuesCondition{
							ConditionType: armsecurityinsight.AutomationRuleConditionType("Property").ToPtr(),
							ConditionProperties: &armsecurityinsight.AutomationRulePropertyValuesConditionProperties{
								Operator:     armsecurityinsight.AutomationRulePropertyConditionSupportedOperator("Contains").ToPtr(),
								PropertyName: armsecurityinsight.AutomationRulePropertyConditionSupportedProperty("IncidentRelatedAnalyticRuleIds").ToPtr(),
								PropertyValues: []*string{
									to.StringPtr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/alertRules/fab3d2d4-747f-46a7-8ef0-9c0be8112bf7"),
									to.StringPtr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/alertRules/8deb8303-e94d-46ff-96e0-5fd94b33df1a")},
							},
						}},
					IsEnabled:    to.BoolPtr(true),
					TriggersOn:   armsecurityinsight.TriggersOn("Incidents").ToPtr(),
					TriggersWhen: armsecurityinsight.TriggersWhen("Created").ToPtr(),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.AutomationRulesClientCreateOrUpdateResult)
}
