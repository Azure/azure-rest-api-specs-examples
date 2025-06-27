
import com.azure.resourcemanager.servicefabricmanagedclusters.models.NodeType;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for NodeTypes Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-03-01-preview/NodeTypePatchOperation_example.json
     */
    /**
     * Sample code: Patch a node type.
     * 
     * @param manager Entry point to ServiceFabricManagedClustersManager.
     */
    public static void patchANodeType(
        com.azure.resourcemanager.servicefabricmanagedclusters.ServiceFabricManagedClustersManager manager) {
        NodeType resource = manager.nodeTypes()
            .getWithResponse("resRg", "myCluster", "BE", com.azure.core.util.Context.NONE).getValue();
        resource.update().withTags(mapOf("a", "b")).apply();
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
