
import com.azure.resourcemanager.hdinsight.models.AzureMonitorRequest;

/**
 * Samples for Extensions EnableAzureMonitor.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hdinsight/resource-manager/Microsoft.HDInsight/preview/2024-08-01-preview/examples/
     * EnableLinuxClusterAzureMonitor.json
     */
    /**
     * Sample code: Enable azure monitor.
     * 
     * @param manager Entry point to HDInsightManager.
     */
    public static void enableAzureMonitor(com.azure.resourcemanager.hdinsight.HDInsightManager manager) {
        manager.extensions()
            .enableAzureMonitor("rg1", "cluster1", new AzureMonitorRequest()
                .withWorkspaceId("a2090ead-8c9f-4fba-b70e-533e3e003163").withPrimaryKey("fakeTokenPlaceholder"),
                com.azure.core.util.Context.NONE);
    }
}
