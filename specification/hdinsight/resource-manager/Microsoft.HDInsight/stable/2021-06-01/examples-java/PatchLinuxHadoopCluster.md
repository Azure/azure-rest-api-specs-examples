```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.hdinsight.models.Cluster;
import java.util.HashMap;
import java.util.Map;

/** Samples for Clusters Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/hdinsight/resource-manager/Microsoft.HDInsight/stable/2021-06-01/examples/PatchLinuxHadoopCluster.json
     */
    /**
     * Sample code: Patch HDInsight Linux clusters.
     *
     * @param manager Entry point to HDInsightManager.
     */
    public static void patchHDInsightLinuxClusters(com.azure.resourcemanager.hdinsight.HDInsightManager manager) {
        Cluster resource =
            manager.clusters().getByResourceGroupWithResponse("rg1", "cluster1", Context.NONE).getValue();
        resource.update().withTags(mapOf("key1", "val1", "key2", "val2")).apply();
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
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-hdinsight_1.0.0-beta.5/sdk/hdinsight/azure-resourcemanager-hdinsight/README.md) on how to add the SDK to your project and authenticate.
