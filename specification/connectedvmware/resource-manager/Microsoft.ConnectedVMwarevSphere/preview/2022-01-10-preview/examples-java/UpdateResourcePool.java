import com.azure.core.util.Context;
import com.azure.resourcemanager.connectedvmware.models.ResourcePool;
import java.util.HashMap;
import java.util.Map;

/** Samples for ResourcePools Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/preview/2022-01-10-preview/examples/UpdateResourcePool.json
     */
    /**
     * Sample code: UpdateResourcePool.
     *
     * @param manager Entry point to ConnectedVMwareManager.
     */
    public static void updateResourcePool(com.azure.resourcemanager.connectedvmware.ConnectedVMwareManager manager) {
        ResourcePool resource =
            manager.resourcePools().getByResourceGroupWithResponse("testrg", "HRPool", Context.NONE).getValue();
        resource.update().withTags(mapOf("tag1", "value1", "tag2", "value2")).apply();
    }

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
