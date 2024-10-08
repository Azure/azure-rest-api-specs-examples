
/**
 * Samples for Extensions GetMonitoringStatus.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hdinsight/resource-manager/Microsoft.HDInsight/preview/2024-08-01-preview/examples/
     * GetLinuxClusterMonitoringStatus.json
     */
    /**
     * Sample code: Get cluster monitoring status.
     * 
     * @param manager Entry point to HDInsightManager.
     */
    public static void getClusterMonitoringStatus(com.azure.resourcemanager.hdinsight.HDInsightManager manager) {
        manager.extensions().getMonitoringStatusWithResponse("rg1", "cluster1", com.azure.core.util.Context.NONE);
    }
}
