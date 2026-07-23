
import com.azure.resourcemanager.containerregistry.fluent.models.CacheRuleInner;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for CacheRules Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/CacheRuleCreate.json
     */
    /**
     * Sample code: CacheRuleCreate.
     * 
     * @param manager Entry point to ContainerRegistryManager.
     */
    public static void cacheRuleCreate(com.azure.resourcemanager.containerregistry.ContainerRegistryManager manager) {
        manager.serviceClient().getCacheRules().create("myResourceGroup", "myRegistry", "myCacheRule",
            new CacheRuleInner().withCredentialSetResourceId("fakeTokenPlaceholder")
                .withSourceRepository("docker.io/library/hello-world")
                .withTargetRepository("cached-docker-hub/hello-world"),
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
