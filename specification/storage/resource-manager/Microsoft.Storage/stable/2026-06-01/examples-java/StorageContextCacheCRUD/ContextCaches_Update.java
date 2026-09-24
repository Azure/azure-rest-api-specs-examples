
import com.azure.resourcemanager.storage.models.ContextCachePropertiesUpdate;
import com.azure.resourcemanager.storage.models.ContextCacheUpdate;
import com.azure.resourcemanager.storage.models.Identity;
import com.azure.resourcemanager.storage.models.IdentityType;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for ContextCaches Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-01/StorageContextCacheCRUD/ContextCaches_Update.json
     */
    /**
     * Sample code: Update a Context Cache tags.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void updateAContextCacheTags(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getContextCaches().update("testrg", "testaccount",
            new ContextCacheUpdate().withTags(mapOf("environment", "production", "team", "context-cache"))
                .withIdentity(new Identity().withType(IdentityType.SYSTEM_ASSIGNED)).withProperties(
                    new ContextCachePropertiesUpdate().withDescription("Updated Prompt Service account description")),
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
