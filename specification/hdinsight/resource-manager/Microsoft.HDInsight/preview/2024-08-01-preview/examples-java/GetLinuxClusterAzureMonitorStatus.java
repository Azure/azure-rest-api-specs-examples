
/**
 * Samples for Extensions GetAzureMonitorStatus.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hdinsight/resource-manager/Microsoft.HDInsight/preview/2024-08-01-preview/examples/
     * GetLinuxClusterAzureMonitorStatus.json
     */
    /**
     * Sample code: Get azure monitor status.
     * 
     * @param manager Entry point to HDInsightManager.
     */
    public static void getAzureMonitorStatus(com.azure.resourcemanager.hdinsight.HDInsightManager manager) {
        manager.extensions().getAzureMonitorStatusWithResponse("rg1", "cluster1", com.azure.core.util.Context.NONE);
    }
}
