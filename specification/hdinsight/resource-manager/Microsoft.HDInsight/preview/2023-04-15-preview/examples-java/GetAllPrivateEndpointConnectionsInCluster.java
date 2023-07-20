/** Samples for PrivateEndpointConnections ListByCluster. */
public final class Main {
    /*
     * x-ms-original-file: specification/hdinsight/resource-manager/Microsoft.HDInsight/preview/2023-04-15-preview/examples/GetAllPrivateEndpointConnectionsInCluster.json
     */
    /**
     * Sample code: Get all private endpoint connections for a specific HDInsight cluster.
     *
     * @param manager Entry point to HDInsightManager.
     */
    public static void getAllPrivateEndpointConnectionsForASpecificHDInsightCluster(
        com.azure.resourcemanager.hdinsight.HDInsightManager manager) {
        manager.privateEndpointConnections().listByCluster("rg1", "cluster1", com.azure.core.util.Context.NONE);
    }
}
