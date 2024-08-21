
/**
 * Samples for PrivateEndpointConnections Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hdinsight/resource-manager/Microsoft.HDInsight/preview/2024-08-01-preview/examples/
     * GetPrivateEndpointConnection.json
     */
    /**
     * Sample code: Get specific private endpoint connection for a specific HDInsight cluster.
     * 
     * @param manager Entry point to HDInsightManager.
     */
    public static void getSpecificPrivateEndpointConnectionForASpecificHDInsightCluster(
        com.azure.resourcemanager.hdinsight.HDInsightManager manager) {
        manager.privateEndpointConnections().getWithResponse("rg1", "cluster1",
            "testprivateep.b3bf5fed-9b12-4560-b7d0-2abe1bba07e2", com.azure.core.util.Context.NONE);
    }
}
