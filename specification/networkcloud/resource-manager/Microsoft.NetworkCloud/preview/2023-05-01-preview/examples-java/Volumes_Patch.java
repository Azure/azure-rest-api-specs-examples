import com.azure.resourcemanager.networkcloud.models.Volume;
import java.util.HashMap;
import java.util.Map;

/** Samples for Volumes Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/networkcloud/resource-manager/Microsoft.NetworkCloud/preview/2023-05-01-preview/examples/Volumes_Patch.json
     */
    /**
     * Sample code: Patch volume.
     *
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void patchVolume(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        Volume resource =
            manager
                .volumes()
                .getByResourceGroupWithResponse("resourceGroupName", "volumeName", com.azure.core.util.Context.NONE)
                .getValue();
        resource.update().withTags(mapOf("key1", "myvalue1", "key2", "myvalue2")).apply();
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
