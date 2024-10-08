
import com.azure.resourcemanager.hdinsight.models.ClusterMonitoringRequest;

/**
 * Samples for Extensions EnableMonitoring.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hdinsight/resource-manager/Microsoft.HDInsight/preview/2024-08-01-preview/examples/
     * EnableLinuxClusterMonitoring.json
     */
    /**
     * Sample code: Enable cluster monitoring.
     * 
     * @param manager Entry point to HDInsightManager.
     */
    public static void enableClusterMonitoring(com.azure.resourcemanager.hdinsight.HDInsightManager manager) {
        manager.extensions()
            .enableMonitoring(
                "rg1", "cluster1", new ClusterMonitoringRequest()
                    .withWorkspaceId("a2090ead-8c9f-4fba-b70e-533e3e003163").withPrimaryKey("fakeTokenPlaceholder"),
                com.azure.core.util.Context.NONE);
    }
}
