
import com.azure.resourcemanager.alertsmanagement.models.AlertProcessingRule;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for AlertProcessingRules Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/alertsmanagement/resource-manager/Microsoft.AlertsManagement/stable/2021-08-08/examples/
     * AlertProcessingRules_Patch.json
     */
    /**
     * Sample code: PatchAlertProcessingRule.
     * 
     * @param manager Entry point to AlertsManagementManager.
     */
    public static void
        patchAlertProcessingRule(com.azure.resourcemanager.alertsmanagement.AlertsManagementManager manager) {
        AlertProcessingRule resource
            = manager.alertProcessingRules().getByResourceGroupWithResponse("alertscorrelationrg", "WeeklySuppression",
                com.azure.core.util.Context.NONE).getValue();
        resource.update().withTags(mapOf("key1", "fakeTokenPlaceholder", "key2", "fakeTokenPlaceholder"))
            .withEnabled(false).apply();
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
