import com.azure.core.util.Context;

/** Samples for Clusters GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/hdinsight/resource-manager/Microsoft.HDInsight/stable/2021-06-01/examples/GetLinuxHadoopCluster.json
     */
    /**
     * Sample code: Get Hadoop on Linux cluster.
     *
     * @param manager Entry point to HDInsightManager.
     */
    public static void getHadoopOnLinuxCluster(com.azure.resourcemanager.hdinsight.HDInsightManager manager) {
        manager.clusters().getByResourceGroupWithResponse("rg1", "cluster1", Context.NONE);
    }
}
