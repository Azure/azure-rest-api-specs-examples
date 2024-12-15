
/**
 * Samples for Clusters GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/stable/2022-10-01/examples/
     * ClustersGet.json
     */
    /**
     * Sample code: ClustersGet.
     * 
     * @param manager Entry point to LogAnalyticsManager.
     */
    public static void clustersGet(com.azure.resourcemanager.loganalytics.LogAnalyticsManager manager) {
        manager.clusters().getByResourceGroupWithResponse("oiautorest6685", "oiautorest6685",
            com.azure.core.util.Context.NONE);
    }
}
