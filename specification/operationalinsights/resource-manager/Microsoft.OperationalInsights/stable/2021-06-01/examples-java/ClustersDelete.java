/** Samples for Clusters Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/stable/2021-06-01/examples/ClustersDelete.json
     */
    /**
     * Sample code: ClustersDelete.
     *
     * @param manager Entry point to LogAnalyticsManager.
     */
    public static void clustersDelete(com.azure.resourcemanager.loganalytics.LogAnalyticsManager manager) {
        manager.clusters().delete("oiautorest6685", "oiautorest6685", com.azure.core.util.Context.NONE);
    }
}
