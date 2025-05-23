
import com.azure.resourcemanager.connectedvmware.models.HostModel;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Hosts Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/stable/2023-10-01/examples/
     * UpdateHost.json
     */
    /**
     * Sample code: UpdateHost.
     * 
     * @param manager Entry point to ConnectedVMwareManager.
     */
    public static void updateHost(com.azure.resourcemanager.connectedvmware.ConnectedVMwareManager manager) {
        HostModel resource = manager.hosts()
            .getByResourceGroupWithResponse("testrg", "HRHost", com.azure.core.util.Context.NONE).getValue();
        resource.update().withTags(mapOf("tag1", "value1", "tag2", "value2")).apply();
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
