
/**
 * Samples for Extensions GetAzureMonitorAgentStatus.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hdinsight/resource-manager/Microsoft.HDInsight/preview/2024-08-01-preview/examples/
     * GetLinuxClusterAzureMonitorAgentStatus.json
     */
    /**
     * Sample code: Get azure monitor agent status.
     * 
     * @param manager Entry point to HDInsightManager.
     */
    public static void getAzureMonitorAgentStatus(com.azure.resourcemanager.hdinsight.HDInsightManager manager) {
        manager.extensions().getAzureMonitorAgentStatusWithResponse("rg1", "cluster1",
            com.azure.core.util.Context.NONE);
    }
}
