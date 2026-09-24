
import com.azure.resourcemanager.storage.fluent.models.ContextCacheInner;
import com.azure.resourcemanager.storage.models.ContextCacheAccountKind;
import com.azure.resourcemanager.storage.models.ContextCacheProperties;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for ContextCaches CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-01/StorageContextCacheCRUD/ContextCaches_CreateOrUpdate.json
     */
    /**
     * Sample code: Create a Context Cache.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void createAContextCache(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getContextCaches().createOrUpdate("testrg", "testaccount",
            new ContextCacheInner().withLocation("eastus").withTags(mapOf("environment", "test"))
                .withProperties(new ContextCacheProperties().withAccountKind(ContextCacheAccountKind.REGIONAL)
                    .withDescription("Test Azure Context Cache account")),
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
