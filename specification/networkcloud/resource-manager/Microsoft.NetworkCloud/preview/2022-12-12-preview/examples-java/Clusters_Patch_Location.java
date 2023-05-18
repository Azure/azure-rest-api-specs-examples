import com.azure.resourcemanager.networkcloud.models.Cluster;
import java.util.HashMap;
import java.util.Map;

/** Samples for Clusters Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/networkcloud/resource-manager/Microsoft.NetworkCloud/preview/2022-12-12-preview/examples/Clusters_Patch_Location.json
     */
    /**
     * Sample code: Patch cluster location.
     *
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void patchClusterLocation(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        Cluster resource =
            manager
                .clusters()
                .getByResourceGroupWithResponse("resourceGroupName", "clusterName", com.azure.core.util.Context.NONE)
                .getValue();
        resource
            .update()
            .withTags(mapOf("key1", "myvalue1", "key2", "myvalue2"))
            .withClusterLocation("Foo Street, 3rd Floor, row 9")
            .apply();
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
