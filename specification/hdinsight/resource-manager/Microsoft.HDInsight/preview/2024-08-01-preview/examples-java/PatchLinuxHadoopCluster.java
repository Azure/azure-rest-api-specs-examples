
import com.azure.resourcemanager.hdinsight.models.Cluster;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Clusters Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hdinsight/resource-manager/Microsoft.HDInsight/preview/2024-08-01-preview/examples/
     * PatchLinuxHadoopCluster.json
     */
    /**
     * Sample code: Patch HDInsight Linux clusters.
     * 
     * @param manager Entry point to HDInsightManager.
     */
    public static void patchHDInsightLinuxClusters(com.azure.resourcemanager.hdinsight.HDInsightManager manager) {
        Cluster resource = manager.clusters()
            .getByResourceGroupWithResponse("rg1", "cluster1", com.azure.core.util.Context.NONE).getValue();
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
