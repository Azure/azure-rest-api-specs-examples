
/**
 * Samples for Clusters Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2020-03-01-preview/examples/
     * Cluster_Delete.json
     */
    /**
     * Sample code: Delete a cluster.
     * 
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void deleteACluster(com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) {
        manager.clusters().delete("sjrg", "testcluster", com.azure.core.util.Context.NONE);
    }
}
