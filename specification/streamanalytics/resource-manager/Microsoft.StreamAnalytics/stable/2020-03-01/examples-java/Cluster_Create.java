
import com.azure.resourcemanager.streamanalytics.models.ClusterSku;
import com.azure.resourcemanager.streamanalytics.models.ClusterSkuName;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Clusters CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/
     * Cluster_Create.json
     */
    /**
     * Sample code: Create a new cluster.
     * 
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void createANewCluster(com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) {
        manager.clusters().define("An Example Cluster").withRegion("North US").withExistingResourceGroup("sjrg")
            .withTags(mapOf("key", "fakeTokenPlaceholder"))
            .withSku(new ClusterSku().withName(ClusterSkuName.DEFAULT).withCapacity(48)).create();
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
