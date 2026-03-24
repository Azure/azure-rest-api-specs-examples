
import com.azure.resourcemanager.containerregistry.fluent.models.ReplicationInner;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Replications Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01-preview/ReplicationCreate.json
     */
    /**
     * Sample code: ReplicationCreate.
     * 
     * @param manager Entry point to ContainerRegistryManager.
     */
    public static void replicationCreate(com.azure.resourcemanager.containerregistry.ContainerRegistryManager manager) {
        manager.serviceClient().getReplications().create("myResourceGroup", "myRegistry", "myReplication",
            new ReplicationInner().withLocation("eastus").withTags(mapOf("key", "fakeTokenPlaceholder")),
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
