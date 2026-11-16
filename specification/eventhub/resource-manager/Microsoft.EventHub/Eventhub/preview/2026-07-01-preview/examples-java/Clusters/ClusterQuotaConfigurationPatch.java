
import com.azure.resourcemanager.eventhubs.fluent.models.ClusterQuotaConfigurationPropertiesInner;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Configuration Patch.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/Clusters/ClusterQuotaConfigurationPatch.json
     */
    /**
     * Sample code: ClustersQuotasConfigurationPatch.
     * 
     * @param manager Entry point to EventHubsManager.
     */
    public static void clustersQuotasConfigurationPatch(com.azure.resourcemanager.eventhubs.EventHubsManager manager) {
        manager.serviceClient().getConfigurations().patchWithResponse("ArunMonocle", "testCluster",
            new ClusterQuotaConfigurationPropertiesInner()
                .withSettings(mapOf("eventhub-per-namespace-quota", "20", "namespaces-per-cluster-quota", "200")),
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
