import com.azure.core.util.Context;

/** Samples for Applications Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/hdinsight/resource-manager/Microsoft.HDInsight/stable/2021-06-01/examples/DeleteApplication.json
     */
    /**
     * Sample code: Delete Application from HDInsight cluster.
     *
     * @param manager Entry point to HDInsightManager.
     */
    public static void deleteApplicationFromHDInsightCluster(
        com.azure.resourcemanager.hdinsight.HDInsightManager manager) {
        manager.applications().delete("rg1", "cluster1", "hue", Context.NONE);
    }
}
