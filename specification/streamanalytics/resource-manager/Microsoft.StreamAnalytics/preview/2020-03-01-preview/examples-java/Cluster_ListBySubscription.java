
/**
 * Samples for Clusters List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2020-03-01-preview/examples/
     * Cluster_ListBySubscription.json
     */
    /**
     * Sample code: List the clusters in a subscription.
     * 
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void
        listTheClustersInASubscription(com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) {
        manager.clusters().list(com.azure.core.util.Context.NONE);
    }
}
