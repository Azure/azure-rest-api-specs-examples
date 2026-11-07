
import com.azure.resourcemanager.monitor.fluent.models.AzureMonitorPrivateLinkScopeInner;
import com.azure.resourcemanager.monitor.models.AccessMode;
import com.azure.resourcemanager.monitor.models.AccessModeSettings;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for PrivateLinkScopes CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2023-06-01-preview/PrivateLinkScopesCreate.json
     */
    /**
     * Sample code: PrivateLinkScopeCreate.
     * 
     * @param manager Entry point to MonitorManager.
     */
    public static void privateLinkScopeCreate(com.azure.resourcemanager.monitor.MonitorManager manager) {
        manager.serviceClient().getPrivateLinkScopes().createOrUpdateWithResponse("my-resource-group",
            "my-privatelinkscope",
            new AzureMonitorPrivateLinkScopeInner().withLocation("Global")
                .withAccessModeSettings(new AccessModeSettings().withQueryAccessMode(AccessMode.OPEN)
                    .withIngestionAccessMode(AccessMode.OPEN).withExclusions(Arrays.asList())),
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
