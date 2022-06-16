import com.azure.core.util.Context;

/** Samples for PrivateLinkResources Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/hdinsight/resource-manager/Microsoft.HDInsight/stable/2021-06-01/examples/GetPrivateLinkResource.json
     */
    /**
     * Sample code: Get specific private link resource in a specific HDInsight cluster.
     *
     * @param manager Entry point to HDInsightManager.
     */
    public static void getSpecificPrivateLinkResourceInASpecificHDInsightCluster(
        com.azure.resourcemanager.hdinsight.HDInsightManager manager) {
        manager.privateLinkResources().getWithResponse("rg1", "cluster1", "gateway", Context.NONE);
    }
}
