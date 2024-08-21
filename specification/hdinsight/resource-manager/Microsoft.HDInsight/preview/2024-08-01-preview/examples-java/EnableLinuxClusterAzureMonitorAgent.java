
import com.azure.resourcemanager.hdinsight.models.AzureMonitorRequest;

/**
 * Samples for Extensions EnableAzureMonitorAgent.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hdinsight/resource-manager/Microsoft.HDInsight/preview/2024-08-01-preview/examples/
     * EnableLinuxClusterAzureMonitorAgent.json
     */
    /**
     * Sample code: Enable azure monitoring agent.
     * 
     * @param manager Entry point to HDInsightManager.
     */
    public static void enableAzureMonitoringAgent(com.azure.resourcemanager.hdinsight.HDInsightManager manager) {
        manager.extensions()
            .enableAzureMonitorAgent("rg1", "cluster1", new AzureMonitorRequest()
                .withWorkspaceId("a2090ead-8c9f-4fba-b70e-533e3e003163").withPrimaryKey("fakeTokenPlaceholder"),
                com.azure.core.util.Context.NONE);
    }
}
