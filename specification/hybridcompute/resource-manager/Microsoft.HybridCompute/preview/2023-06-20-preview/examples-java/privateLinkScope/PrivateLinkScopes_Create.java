import java.util.HashMap;
import java.util.Map;

/** Samples for PrivateLinkScopes CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2023-06-20-preview/examples/privateLinkScope/PrivateLinkScopes_Create.json
     */
    /**
     * Sample code: PrivateLinkScopeCreate.
     *
     * @param manager Entry point to HybridComputeManager.
     */
    public static void privateLinkScopeCreate(com.azure.resourcemanager.hybridcompute.HybridComputeManager manager) {
        manager
            .privateLinkScopes()
            .define("my-privatelinkscope")
            .withRegion("westus")
            .withExistingResourceGroup("my-resource-group")
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
