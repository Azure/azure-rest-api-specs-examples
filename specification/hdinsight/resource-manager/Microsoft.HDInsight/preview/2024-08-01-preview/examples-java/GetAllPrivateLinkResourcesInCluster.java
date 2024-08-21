
/**
 * Samples for PrivateLinkResources ListByCluster.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hdinsight/resource-manager/Microsoft.HDInsight/preview/2024-08-01-preview/examples/
     * GetAllPrivateLinkResourcesInCluster.json
     */
    /**
     * Sample code: Get all private link resources in a specific HDInsight cluster.
     * 
     * @param manager Entry point to HDInsightManager.
     */
    public static void getAllPrivateLinkResourcesInASpecificHDInsightCluster(
        com.azure.resourcemanager.hdinsight.HDInsightManager manager) {
        manager.privateLinkResources().listByClusterWithResponse("rg1", "cluster1", com.azure.core.util.Context.NONE);
    }
}
