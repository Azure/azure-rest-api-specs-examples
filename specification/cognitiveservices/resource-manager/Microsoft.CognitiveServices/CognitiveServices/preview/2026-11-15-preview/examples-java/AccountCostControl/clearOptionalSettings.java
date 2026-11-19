
import com.azure.resourcemanager.cognitiveservices.models.AccountCostControl;
import com.azure.resourcemanager.cognitiveservices.models.CostControlPatchProperties;
import com.azure.resourcemanager.cognitiveservices.models.CostControlProviderPropertiesPatch;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for AccountCostControls Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-11-15-preview/AccountCostControl/clearOptionalSettings.json
     */
    /**
     * Sample code: Clear filters and remove named notifications and blocking actions.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void clearFiltersAndRemoveNamedNotificationsAndBlockingActions(
        com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        AccountCostControl resource = manager.accountCostControls().getWithResponse("foundry-resource-group",
            "foundry-account", "production-agents", com.azure.core.util.Context.NONE).getValue();
        resource.update().withProperties(new CostControlPatchProperties().withNotifications(mapOf("usageAlert", null))
            .withProviderProperties(new CostControlProviderPropertiesPatch().withActions(mapOf("stopRequests", null))))
            .withIfMatch("\"00000000-0000-0000-0000-000000000001\"").apply();
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
