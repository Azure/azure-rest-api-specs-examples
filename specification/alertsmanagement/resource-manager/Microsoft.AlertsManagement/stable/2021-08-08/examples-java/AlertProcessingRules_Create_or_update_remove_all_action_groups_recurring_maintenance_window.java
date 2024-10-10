
import com.azure.resourcemanager.alertsmanagement.models.AlertProcessingRuleProperties;
import com.azure.resourcemanager.alertsmanagement.models.Condition;
import com.azure.resourcemanager.alertsmanagement.models.DaysOfWeek;
import com.azure.resourcemanager.alertsmanagement.models.Field;
import com.azure.resourcemanager.alertsmanagement.models.Operator;
import com.azure.resourcemanager.alertsmanagement.models.RemoveAllActionGroups;
import com.azure.resourcemanager.alertsmanagement.models.Schedule;
import com.azure.resourcemanager.alertsmanagement.models.WeeklyRecurrence;
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
     * AlertProcessingRules_Create_or_update_remove_all_action_groups_recurring_maintenance_window.json
     */
    /**
     * Sample code: Create or update a rule that removes all action groups from all alerts on any VM in two resource
     * groups during a recurring maintenance window (2200-0400 every Sat and Sun, India Standard Time).
     * 
     * @param manager Entry point to AlertsManagementManager.
     */
    public static void
        createOrUpdateARuleThatRemovesAllActionGroupsFromAllAlertsOnAnyVMInTwoResourceGroupsDuringARecurringMaintenanceWindow22000400EverySatAndSunIndiaStandardTime(
            com.azure.resourcemanager.alertsmanagement.AlertsManagementManager manager) {
        manager.alertProcessingRules().define("RemoveActionGroupsRecurringMaintenance").withRegion("Global")
            .withExistingResourceGroup("alertscorrelationrg").withTags(mapOf())
            .withProperties(new AlertProcessingRuleProperties()
                .withScopes(Arrays.asList("/subscriptions/subId1/resourceGroups/RGId1",
                    "/subscriptions/subId1/resourceGroups/RGId2"))
                .withConditions(Arrays.asList(new Condition().withField(Field.TARGET_RESOURCE_TYPE)
                    .withOperator(Operator.EQUALS).withValues(Arrays.asList("microsoft.compute/virtualmachines"))))
                .withSchedule(new Schedule().withTimeZone("India Standard Time").withRecurrences(
                    Arrays.asList(new WeeklyRecurrence().withStartTime("22:00:00").withEndTime("04:00:00")
                        .withDaysOfWeek(Arrays.asList(DaysOfWeek.SATURDAY, DaysOfWeek.SUNDAY)))))
                .withActions(Arrays.asList(new RemoveAllActionGroups()))
                .withDescription(
                    "Remove all ActionGroups from all Vitual machine Alerts during the recurring maintenance")
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
