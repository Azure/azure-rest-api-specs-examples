import com.azure.core.util.Context;

/** Samples for Extensions Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/hdinsight/resource-manager/Microsoft.HDInsight/stable/2021-06-01/examples/DeleteExtension.json
     */
    /**
     * Sample code: Delete an extension.
     *
     * @param manager Entry point to HDInsightManager.
     */
    public static void deleteAnExtension(com.azure.resourcemanager.hdinsight.HDInsightManager manager) {
        manager.extensions().delete("rg1", "cluster1", "clustermonitoring", Context.NONE);
    }
}
