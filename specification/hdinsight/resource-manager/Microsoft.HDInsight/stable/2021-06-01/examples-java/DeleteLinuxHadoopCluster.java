import com.azure.core.util.Context;

/** Samples for Clusters Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/hdinsight/resource-manager/Microsoft.HDInsight/stable/2021-06-01/examples/DeleteLinuxHadoopCluster.json
     */
    /**
     * Sample code: Delete Hadoop on Linux cluster.
     *
     * @param manager Entry point to HDInsightManager.
     */
    public static void deleteHadoopOnLinuxCluster(com.azure.resourcemanager.hdinsight.HDInsightManager manager) {
        manager.clusters().delete("rg1", "cluster1", Context.NONE);
    }
}
