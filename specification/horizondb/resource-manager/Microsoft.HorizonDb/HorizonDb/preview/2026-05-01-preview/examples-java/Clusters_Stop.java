
/**
 * Samples for HorizonDbClusters Stop.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/Clusters_Stop.json
     */
    /**
     * Sample code: Stop a HorizonDB cluster.
     * 
     * @param manager Entry point to HorizonDbManager.
     */
    public static void stopAHorizonDBCluster(com.azure.resourcemanager.horizondb.HorizonDbManager manager) {
        manager.horizonDbClusters().stop("exampleresourcegroup", "examplecluster", com.azure.core.util.Context.NONE);
    }
}
