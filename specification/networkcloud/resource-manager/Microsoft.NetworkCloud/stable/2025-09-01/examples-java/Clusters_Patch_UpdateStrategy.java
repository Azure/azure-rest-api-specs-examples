
import com.azure.resourcemanager.networkcloud.models.Cluster;
import com.azure.resourcemanager.networkcloud.models.ClusterUpdateStrategy;
import com.azure.resourcemanager.networkcloud.models.ClusterUpdateStrategyType;
import com.azure.resourcemanager.networkcloud.models.ValidationThresholdType;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Clusters Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/networkcloud/resource-manager/Microsoft.NetworkCloud/stable/2025-09-01/examples/
     * Clusters_Patch_UpdateStrategy.json
     */
    /**
     * Sample code: Patch update strategy.
     * 
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void patchUpdateStrategy(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        Cluster resource = manager.clusters()
            .getByResourceGroupWithResponse("resourceGroupName", "clusterName", com.azure.core.util.Context.NONE)
            .getValue();
        resource.update().withTags(mapOf("key1", "fakeTokenPlaceholder", "key2", "fakeTokenPlaceholder"))
            .withUpdateStrategy(
                new ClusterUpdateStrategy().withMaxUnavailable(4L).withStrategyType(ClusterUpdateStrategyType.RACK)
                    .withThresholdType(ValidationThresholdType.COUNT_SUCCESS).withThresholdValue(4L)
                    .withWaitTimeMinutes(10L))
            .apply();
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
