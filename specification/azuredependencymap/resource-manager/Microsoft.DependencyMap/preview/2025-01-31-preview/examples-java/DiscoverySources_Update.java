
import com.azure.resourcemanager.dependencymap.models.DiscoverySourceResource;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for DiscoverySources Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-31-preview/DiscoverySources_Update.json
     */
    /**
     * Sample code: DiscoverySources_Update - generated by [MaximumSet] rule.
     * 
     * @param manager Entry point to DependencyMapManager.
     */
    public static void discoverySourcesUpdateGeneratedByMaximumSetRule(
        com.azure.resourcemanager.dependencymap.DependencyMapManager manager) {
        DiscoverySourceResource resource = manager.discoverySources()
            .getWithResponse("rgdependencyMap", "mapsTest1", "sourceTest1", com.azure.core.util.Context.NONE)
            .getValue();
        resource.update().withTags(mapOf()).apply();
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
