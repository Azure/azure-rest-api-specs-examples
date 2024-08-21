
/**
 * Samples for Extensions DisableMonitoring.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hdinsight/resource-manager/Microsoft.HDInsight/preview/2024-08-01-preview/examples/
     * DisableLinuxClusterMonitoring.json
     */
    /**
     * Sample code: Disable cluster monitoring.
     * 
     * @param manager Entry point to HDInsightManager.
     */
    public static void disableClusterMonitoring(com.azure.resourcemanager.hdinsight.HDInsightManager manager) {
        manager.extensions().disableMonitoring("rg1", "cluster1", com.azure.core.util.Context.NONE);
    }
}
