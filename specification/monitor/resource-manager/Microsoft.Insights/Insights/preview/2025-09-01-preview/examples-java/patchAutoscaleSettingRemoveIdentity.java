
import com.azure.resourcemanager.monitor.models.AutoscaleSettingResourcePatch;
import com.azure.resourcemanager.monitor.models.ManagedServiceIdentity;
import com.azure.resourcemanager.monitor.models.ManagedServiceIdentityType;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for AutoscaleSettings Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01-preview/patchAutoscaleSettingRemoveIdentity.json
     */
    /**
     * Sample code: Remove the managed identity from an autoscale setting.
     * 
     * @param manager Entry point to MonitorManager.
     */
    public static void
        removeTheManagedIdentityFromAnAutoscaleSetting(com.azure.resourcemanager.monitor.MonitorManager manager) {
        manager.serviceClient().getAutoscaleSettings().updateWithResponse("TestingMetricsScaleSet", "MySetting",
            new AutoscaleSettingResourcePatch()
                .withIdentity(new ManagedServiceIdentity().withType(ManagedServiceIdentityType.NONE)),
            com.azure.core.util.Context.NONE);
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
