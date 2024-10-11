
import com.azure.resourcemanager.alertsmanagement.models.AlertProcessingRuleProperties;
import com.azure.resourcemanager.alertsmanagement.models.Condition;
import com.azure.resourcemanager.alertsmanagement.models.Field;
import com.azure.resourcemanager.alertsmanagement.models.Operator;
import com.azure.resourcemanager.alertsmanagement.models.RemoveAllActionGroups;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for AlertProcessingRules CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/alertsmanagement/resource-manager/Microsoft.AlertsManagement/stable/2021-08-08/examples/
     * AlertProcessingRules_Create_or_update_remove_all_action_groups_from_specific_alert_rule.json
     */
    /**
     * Sample code: Create or update a rule that removes all action groups from all alerts in a subscription coming from
     * a specific alert rule.
     * 
     * @param manager Entry point to AlertsManagementManager.
     */
    public static void
        createOrUpdateARuleThatRemovesAllActionGroupsFromAllAlertsInASubscriptionComingFromASpecificAlertRule(
            com.azure.resourcemanager.alertsmanagement.AlertsManagementManager manager) {
        manager.alertProcessingRules().define("RemoveActionGroupsSpecificAlertRule").withRegion("Global")
            .withExistingResourceGroup("alertscorrelationrg").withTags(mapOf())
            .withProperties(new AlertProcessingRuleProperties().withScopes(Arrays.asList("/subscriptions/subId1"))
                .withConditions(Arrays.asList(new Condition().withField(Field.ALERT_RULE_ID)
                    .withOperator(Operator.EQUALS).withValues(Arrays.asList(
                        "/subscriptions/suubId1/resourceGroups/Rgid2/providers/microsoft.insights/activityLogAlerts/RuleName"))))
                .withActions(Arrays.asList(new RemoveAllActionGroups()))
                .withDescription("Removes all ActionGroups from all Alerts that fire on above AlertRule")
                .withEnabled(true))
            .create();
    }

    // Use "Map.of" if available
    @SuppressWarnings("unchecked")
    private static <T> Map<String, T> mapOf(Object... inputs) {
        Map<String, T> map = new HashMap<>();
        for (int i = 0; i < inputs.length; i += 2) {
            String key = (String) inputs[i];
            T value = (T) inputs[i + 1];
            map.put(key, value);
        }
        return map;
    }
}
