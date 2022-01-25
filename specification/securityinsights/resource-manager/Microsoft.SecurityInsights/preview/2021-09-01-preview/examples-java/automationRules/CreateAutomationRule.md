Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-securityinsights_1.0.0-beta.1/sdk/securityinsights/azure-resourcemanager-securityinsights/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.resourcemanager.securityinsights.models.AutomationRuleModifyPropertiesAction;
import com.azure.resourcemanager.securityinsights.models.AutomationRuleModifyPropertiesActionConfiguration;
import com.azure.resourcemanager.securityinsights.models.AutomationRulePropertyConditionSupportedOperator;
import com.azure.resourcemanager.securityinsights.models.AutomationRulePropertyConditionSupportedProperty;
import com.azure.resourcemanager.securityinsights.models.AutomationRulePropertyValuesCondition;
import com.azure.resourcemanager.securityinsights.models.AutomationRulePropertyValuesConditionProperties;
import com.azure.resourcemanager.securityinsights.models.AutomationRuleRunPlaybookAction;
import com.azure.resourcemanager.securityinsights.models.AutomationRuleRunPlaybookActionConfiguration;
import com.azure.resourcemanager.securityinsights.models.AutomationRuleTriggeringLogic;
import com.azure.resourcemanager.securityinsights.models.IncidentSeverity;
import com.azure.resourcemanager.securityinsights.models.TriggersOn;
import com.azure.resourcemanager.securityinsights.models.TriggersWhen;
import java.util.Arrays;

/** Samples for AutomationRules CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2021-09-01-preview/examples/automationRules/CreateAutomationRule.json
     */
    /**
     * Sample code: Creates or updates an automation rule.
     *
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void createsOrUpdatesAnAutomationRule(
        com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager
            .automationRules()
            .define("73e01a99-5cd7-4139-a149-9f2736ff2ab5")
            .withExistingWorkspace("myRg", "myWorkspace")
            .withEtag("\"0300bf09-0000-0000-0000-5c37296e0000\"")
            .withDisplayName("High severity incidents escalation")
            .withOrder(1)
            .withTriggeringLogic(
                new AutomationRuleTriggeringLogic()
                    .withIsEnabled(true)
                    .withTriggersOn(TriggersOn.INCIDENTS)
                    .withTriggersWhen(TriggersWhen.CREATED)
                    .withConditions(
                        Arrays
                            .asList(
                                new AutomationRulePropertyValuesCondition()
                                    .withConditionProperties(
                                        new AutomationRulePropertyValuesConditionProperties()
                                            .withPropertyName(
                                                AutomationRulePropertyConditionSupportedProperty
                                                    .INCIDENT_RELATED_ANALYTIC_RULE_IDS)
                                            .withOperator(AutomationRulePropertyConditionSupportedOperator.CONTAINS)
                                            .withPropertyValues(
                                                Arrays
                                                    .asList(
                                                        "/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/alertRules/fab3d2d4-747f-46a7-8ef0-9c0be8112bf7",
                                                        "/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/alertRules/8deb8303-e94d-46ff-96e0-5fd94b33df1a"))))))
            .withActions(
                Arrays
                    .asList(
                        new AutomationRuleModifyPropertiesAction()
                            .withOrder(1)
                            .withActionConfiguration(
                                new AutomationRuleModifyPropertiesActionConfiguration()
                                    .withSeverity(IncidentSeverity.HIGH)),
                        new AutomationRuleRunPlaybookAction()
                            .withOrder(2)
                            .withActionConfiguration(
                                new AutomationRuleRunPlaybookActionConfiguration()
                                    .withLogicAppResourceId(
                                        "/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.Logic/workflows/IncidentPlaybook")
                                    .withTenantId("ee48efaf-50c6-411b-9345-b2bdc3eb4abc"))))
            .create();
    }
}
```
