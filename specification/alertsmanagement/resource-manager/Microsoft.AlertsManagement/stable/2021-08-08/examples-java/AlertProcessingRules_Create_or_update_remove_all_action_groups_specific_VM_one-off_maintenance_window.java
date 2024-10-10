
import com.azure.resourcemanager.alertsmanagement.models.AlertProcessingRuleProperties;
import com.azure.resourcemanager.alertsmanagement.models.RemoveAllActionGroups;
import com.azure.resourcemanager.alertsmanagement.models.Schedule;
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
     * AlertProcessingRules_Create_or_update_remove_all_action_groups_specific_VM_one-off_maintenance_window.json
     */
    /**
     * Sample code: Create or update a rule that removes all action groups from alerts on a specific VM during a one-off
     * maintenance window (1800-2000 at a specific date, Pacific Standard Time).
     * 
     * @param manager Entry point to AlertsManagementManager.
     */
    public static void
        createOrUpdateARuleThatRemovesAllActionGroupsFromAlertsOnASpecificVMDuringAOneOffMaintenanceWindow18002000AtASpecificDatePacificStandardTime(
            com.azure.resourcemanager.alertsmanagement.AlertsManagementManager manager) {
        manager.alertProcessingRules().define("RemoveActionGroupsMaintenanceWindow").withRegion("Global")
            .withExistingResourceGroup("alertscorrelationrg").withTags(mapOf())
            .withProperties(new AlertProcessingRuleProperties()
                .withScopes(Arrays.asList(
                    "/subscriptions/subId1/resourceGroups/RGId1/providers/Microsoft.Compute/virtualMachines/VMName"))
                .withSchedule(new Schedule().withEffectiveFrom("2021-04-15T18:00:00")
                    .withEffectiveUntil("2021-04-15T20:00:00").withTimeZone("Pacific Standard Time"))
                .withActions(Arrays.asList(new RemoveAllActionGroups()))
                .withDescription("Removes all ActionGroups from all Alerts on VMName during the maintenance window")
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
