
import com.azure.resourcemanager.alertsmanagement.models.AddActionGroups;
import com.azure.resourcemanager.alertsmanagement.models.AlertProcessingRuleProperties;
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
     * AlertProcessingRules_Create_or_update_add_action_group_all_alerts_in_subscription.json
     */
    /**
     * Sample code: Create or update a rule that adds an action group to all alerts in a subscription.
     * 
     * @param manager Entry point to AlertsManagementManager.
     */
    public static void createOrUpdateARuleThatAddsAnActionGroupToAllAlertsInASubscription(
        com.azure.resourcemanager.alertsmanagement.AlertsManagementManager manager) {
        manager.alertProcessingRules().define("AddActionGroupToSubscription").withRegion("Global")
            .withExistingResourceGroup("alertscorrelationrg").withTags(mapOf())
            .withProperties(new AlertProcessingRuleProperties().withScopes(Arrays.asList("/subscriptions/subId1"))
                .withActions(Arrays.asList(new AddActionGroups().withActionGroupIds(Arrays.asList(
                    "/subscriptions/subId1/resourcegroups/RGId1/providers/microsoft.insights/actiongroups/ActionGroup1"))))
                .withDescription("Add ActionGroup1 to all alerts in the subscription").withEnabled(true))
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
