
/**
 * Samples for Extensions DisableAzureMonitor.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hdinsight/resource-manager/Microsoft.HDInsight/preview/2024-08-01-preview/examples/
     * DisableLinuxClusterAzureMonitor.json
     */
    /**
     * Sample code: Disable azure monitor.
     * 
     * @param manager Entry point to HDInsightManager.
     */
    public static void disableAzureMonitor(com.azure.resourcemanager.hdinsight.HDInsightManager manager) {
        manager.extensions().disableAzureMonitor("rg1", "cluster1", com.azure.core.util.Context.NONE);
    }
}
