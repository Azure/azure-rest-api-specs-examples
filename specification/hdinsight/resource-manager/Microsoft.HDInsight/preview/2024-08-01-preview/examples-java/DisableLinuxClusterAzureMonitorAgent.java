
/**
 * Samples for Extensions DisableAzureMonitorAgent.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hdinsight/resource-manager/Microsoft.HDInsight/preview/2024-08-01-preview/examples/
     * DisableLinuxClusterAzureMonitorAgent.json
     */
    /**
     * Sample code: Disable azure monitor agent.
     * 
     * @param manager Entry point to HDInsightManager.
     */
    public static void disableAzureMonitorAgent(com.azure.resourcemanager.hdinsight.HDInsightManager manager) {
        manager.extensions().disableAzureMonitorAgent("rg1", "cluster1", com.azure.core.util.Context.NONE);
    }
}
