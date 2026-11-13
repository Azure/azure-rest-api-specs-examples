
import com.azure.resourcemanager.containerregistry.models.CacheRuleUpdateParameters;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for CacheRules Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-15-preview/CacheRuleUpdate.json
     */
    /**
     * Sample code: CacheRuleUpdate.
     * 
     * @param manager Entry point to ContainerRegistryManager.
     */
    public static void cacheRuleUpdate(com.azure.resourcemanager.containerregistry.ContainerRegistryManager manager) {
        manager.serviceClient().getCacheRules().update("myResourceGroup", "myRegistry", "myCacheRule",
            new CacheRuleUpdateParameters().withCredentialSetResourceId("fakeTokenPlaceholder"),
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
