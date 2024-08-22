
/**
 * Samples for PrivateLinkResources Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hdinsight/resource-manager/Microsoft.HDInsight/preview/2024-08-01-preview/examples/
     * GetPrivateLinkResource.json
     */
    /**
     * Sample code: Get specific private link resource in a specific HDInsight cluster.
     * 
     * @param manager Entry point to HDInsightManager.
     */
    public static void getSpecificPrivateLinkResourceInASpecificHDInsightCluster(
        com.azure.resourcemanager.hdinsight.HDInsightManager manager) {
        manager.privateLinkResources().getWithResponse("rg1", "cluster1", "gateway", com.azure.core.util.Context.NONE);
    }
}
