
import com.azure.resourcemanager.alertsmanagement.models.AlertProcessingRuleProperties;
import com.azure.resourcemanager.alertsmanagement.models.DailyRecurrence;
import com.azure.resourcemanager.alertsmanagement.models.DaysOfWeek;
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
     * AlertProcessingRules_Create_or_update_remove_all_action_groups_outside_business_hours.json
     */
    /**
     * Sample code: Create or update a rule that removes all action groups outside business hours (Mon-Fri 09:00-17:00,
     * Eastern Standard Time).
     * 
     * @param manager Entry point to AlertsManagementManager.
     */
    public static void
        createOrUpdateARuleThatRemovesAllActionGroupsOutsideBusinessHoursMonFri09001700EasternStandardTime(
            com.azure.resourcemanager.alertsmanagement.AlertsManagementManager manager) {
        manager.alertProcessingRules().define("RemoveActionGroupsOutsideBusinessHours").withRegion("Global")
            .withExistingResourceGroup("alertscorrelationrg").withTags(mapOf())
            .withProperties(new AlertProcessingRuleProperties().withScopes(Arrays.asList("/subscriptions/subId1"))
                .withSchedule(new Schedule().withTimeZone("Eastern Standard Time").withRecurrences(
                    Arrays.asList(new DailyRecurrence().withStartTime("17:00:00").withEndTime("09:00:00"),
                        new WeeklyRecurrence().withDaysOfWeek(Arrays.asList(DaysOfWeek.SATURDAY, DaysOfWeek.SUNDAY)))))
                .withActions(Arrays.asList(new RemoveAllActionGroups()))
                .withDescription("Remove all ActionGroups outside business hours").withEnabled(true))
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
