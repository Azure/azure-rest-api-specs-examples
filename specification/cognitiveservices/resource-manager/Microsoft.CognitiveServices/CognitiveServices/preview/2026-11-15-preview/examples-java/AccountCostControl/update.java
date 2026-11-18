
import com.azure.resourcemanager.cognitiveservices.models.AccountCostControl;
import com.azure.resourcemanager.cognitiveservices.models.CostControlActionPatch;
import com.azure.resourcemanager.cognitiveservices.models.CostControlNotificationPatch;
import com.azure.resourcemanager.cognitiveservices.models.CostControlPatchProperties;
import com.azure.resourcemanager.cognitiveservices.models.CostControlProviderPropertiesPatch;
import com.azure.resourcemanager.cognitiveservices.models.CostControlTimePeriodPatch;
import java.time.OffsetDateTime;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for AccountCostControls Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-11-15-preview/AccountCostControl/update.json
     */
    /**
     * Sample code: Merge named budget settings and restore owner-default grouping.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void mergeNamedBudgetSettingsAndRestoreOwnerDefaultGrouping(
        com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        AccountCostControl resource = manager.accountCostControls().getWithResponse("foundry-resource-group",
            "foundry-account", "production-agents", com.azure.core.util.Context.NONE).getValue();
        resource.update()
            .withProperties(new CostControlPatchProperties().withGrouping(Arrays.asList()).withAmount(1500.0D)
                .withTimePeriod(
                    new CostControlTimePeriodPatch().withEndDate(OffsetDateTime.parse("2028-11-15T00:00:00Z")))
                .withNotifications(mapOf("usageAlert", new CostControlNotificationPatch().withEnabled(false)))
                .withProviderProperties(new CostControlProviderPropertiesPatch()
                    .withActions(mapOf("stopRequests", new CostControlActionPatch().withThreshold(90.0D)))))
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
