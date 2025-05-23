
import com.azure.resourcemanager.synapse.models.BigDataPoolResourceInfo;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for BigDataPools Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/UpdateBigDataPool.
     * json
     */
    /**
     * Sample code: Update a Big Data pool.
     * 
     * @param manager Entry point to SynapseManager.
     */
    public static void updateABigDataPool(com.azure.resourcemanager.synapse.SynapseManager manager) {
        BigDataPoolResourceInfo resource = manager.bigDataPools().getWithResponse("ExampleResourceGroup",
            "ExampleWorkspace", "ExamplePool", com.azure.core.util.Context.NONE).getValue();
        resource.update().withTags(mapOf("key", "fakeTokenPlaceholder")).apply();
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
