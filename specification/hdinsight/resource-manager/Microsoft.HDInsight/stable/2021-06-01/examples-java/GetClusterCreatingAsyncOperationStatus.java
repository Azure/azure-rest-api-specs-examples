import com.azure.core.util.Context;

/** Samples for Clusters GetAzureAsyncOperationStatus. */
public final class Main {
    /*
     * x-ms-original-file: specification/hdinsight/resource-manager/Microsoft.HDInsight/stable/2021-06-01/examples/GetClusterCreatingAsyncOperationStatus.json
     */
    /**
     * Sample code: Get Async Operation Status of Creating Cluster.
     *
     * @param manager Entry point to HDInsightManager.
     */
    public static void getAsyncOperationStatusOfCreatingCluster(
        com.azure.resourcemanager.hdinsight.HDInsightManager manager) {
        manager
            .clusters()
            .getAzureAsyncOperationStatusWithResponse(
                "rg1", "cluster1", "CF938302-6B4D-44A0-A6D2-C0D67E847AEC", Context.NONE);
    }
}
