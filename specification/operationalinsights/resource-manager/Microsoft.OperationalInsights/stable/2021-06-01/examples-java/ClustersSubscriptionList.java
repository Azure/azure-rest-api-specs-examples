/** Samples for Clusters List. */
public final class Main {
    /*
     * x-ms-original-file: specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/stable/2021-06-01/examples/ClustersSubscriptionList.json
     */
    /**
     * Sample code: ClustersSubscriptionList.
     *
     * @param manager Entry point to LogAnalyticsManager.
     */
    public static void clustersSubscriptionList(com.azure.resourcemanager.loganalytics.LogAnalyticsManager manager) {
        manager.clusters().list(com.azure.core.util.Context.NONE);
    }
}
