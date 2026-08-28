
/**
 * Samples for HorizonDbClusters List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/Clusters_ListBySubscription.json
     */
    /**
     * Sample code: List HorizonDB clusters by subscription.
     * 
     * @param manager Entry point to HorizonDbManager.
     */
    public static void
        listHorizonDBClustersBySubscription(com.azure.resourcemanager.horizondb.HorizonDbManager manager) {
        manager.horizonDbClusters().list(com.azure.core.util.Context.NONE);
    }
}
