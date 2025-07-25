
import com.azure.resourcemanager.networkcloud.models.L2Network;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for L2Networks Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/networkcloud/resource-manager/Microsoft.NetworkCloud/stable/2025-02-01/examples/L2Networks_Patch.
     * json
     */
    /**
     * Sample code: Patch L2 network.
     * 
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void patchL2Network(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        L2Network resource = manager.l2Networks()
            .getByResourceGroupWithResponse("resourceGroupName", "l2NetworkName", com.azure.core.util.Context.NONE)
            .getValue();
        resource.update().withTags(mapOf("key1", "fakeTokenPlaceholder", "key2", "fakeTokenPlaceholder")).apply();
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
