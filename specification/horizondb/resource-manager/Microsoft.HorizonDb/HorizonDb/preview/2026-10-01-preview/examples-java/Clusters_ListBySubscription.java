
/**
 * Samples for Clusters List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/Clusters_ListBySubscription.json
     */
    /**
     * Sample code: List HorizonDB clusters in a subscription.
     * 
     * @param manager Entry point to HorizonDbManager.
     */
    public static void
        listHorizonDBClustersInASubscription(com.azure.resourcemanager.horizondb.HorizonDbManager manager) {
        manager.clusters().list(com.azure.core.util.Context.NONE);
    }
}
