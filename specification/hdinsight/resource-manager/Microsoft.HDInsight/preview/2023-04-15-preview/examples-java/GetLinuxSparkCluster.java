/** Samples for Clusters GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/hdinsight/resource-manager/Microsoft.HDInsight/preview/2023-04-15-preview/examples/GetLinuxSparkCluster.json
     */
    /**
     * Sample code: Get Spark on Linux cluster.
     *
     * @param manager Entry point to HDInsightManager.
     */
    public static void getSparkOnLinuxCluster(com.azure.resourcemanager.hdinsight.HDInsightManager manager) {
        manager.clusters().getByResourceGroupWithResponse("rg1", "cluster1", com.azure.core.util.Context.NONE);
    }
}
