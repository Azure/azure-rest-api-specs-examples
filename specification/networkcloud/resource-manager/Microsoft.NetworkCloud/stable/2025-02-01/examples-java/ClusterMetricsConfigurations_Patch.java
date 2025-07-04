
import com.azure.resourcemanager.networkcloud.models.ClusterMetricsConfiguration;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for MetricsConfigurations Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/networkcloud/resource-manager/Microsoft.NetworkCloud/stable/2025-02-01/examples/
     * ClusterMetricsConfigurations_Patch.json
     */
    /**
     * Sample code: Patch metrics configuration of cluster.
     * 
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void
        patchMetricsConfigurationOfCluster(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        ClusterMetricsConfiguration resource = manager.metricsConfigurations()
            .getWithResponse("resourceGroupName", "clusterName", "default", com.azure.core.util.Context.NONE)
            .getValue();
        resource.update().withTags(mapOf("key1", "fakeTokenPlaceholder", "key2", "fakeTokenPlaceholder"))
            .withCollectionInterval(15L).withEnabledMetrics(Arrays.asList("metric1", "metric2")).apply();
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
