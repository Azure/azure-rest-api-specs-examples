import com.azure.core.util.Context;

/** Samples for Extensions GetMonitoringStatus. */
public final class Main {
    /*
     * x-ms-original-file: specification/hdinsight/resource-manager/Microsoft.HDInsight/stable/2021-06-01/examples/GetLinuxClusterMonitoringStatus.json
     */
    /**
     * Sample code: Enable cluster monitoring.
     *
     * @param manager Entry point to HDInsightManager.
     */
    public static void enableClusterMonitoring(com.azure.resourcemanager.hdinsight.HDInsightManager manager) {
        manager.extensions().getMonitoringStatusWithResponse("rg1", "cluster1", Context.NONE);
    }
}
