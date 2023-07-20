/** Samples for Extensions DisableAzureMonitor. */
public final class Main {
    /*
     * x-ms-original-file: specification/hdinsight/resource-manager/Microsoft.HDInsight/preview/2023-04-15-preview/examples/DisableLinuxClusterAzureMonitor.json
     */
    /**
     * Sample code: Enable cluster monitoring.
     *
     * @param manager Entry point to HDInsightManager.
     */
    public static void enableClusterMonitoring(com.azure.resourcemanager.hdinsight.HDInsightManager manager) {
        manager.extensions().disableAzureMonitor("rg1", "cluster1", com.azure.core.util.Context.NONE);
    }
}
