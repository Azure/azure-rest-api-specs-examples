
import com.azure.core.management.exception.ManagementError;
import com.azure.resourcemanager.connectedcache.models.CacheNodeOldResponse;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for CacheNodesOperations CreateorUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2023-05-01-preview/CacheNodesOperations_CreateorUpdate_MaximumSet_Gen.json
     */
    /**
     * Sample code: CacheNodesOperations_CreateorUpdate.
     * 
     * @param manager Entry point to ConnectedCacheManager.
     */
    public static void
        cacheNodesOperationsCreateorUpdate(com.azure.resourcemanager.connectedcache.ConnectedCacheManager manager) {
        manager.cacheNodesOperations().define("lwrsyhvfpcfrwrim").withRegion("westus")
            .withExistingResourceGroup("rgConnectedCache").withTags(mapOf("key8256", "fakeTokenPlaceholder"))
            .withProperties(new CacheNodeOldResponse().withStatusCode("fakeTokenPlaceholder")
                .withStatusText("bjnsrpzaofjntleoesjwammgbi").withStatusDetails("quuziibkwtgf")
                .withError(new ManagementError()))
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
