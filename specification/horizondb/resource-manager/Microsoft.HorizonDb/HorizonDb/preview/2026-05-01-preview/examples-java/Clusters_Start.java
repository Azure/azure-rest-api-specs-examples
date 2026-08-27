
/**
 * Samples for HorizonDbClusters Start.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/Clusters_Start.json
     */
    /**
     * Sample code: Start a HorizonDB cluster.
     * 
     * @param manager Entry point to HorizonDbManager.
     */
    public static void startAHorizonDBCluster(com.azure.resourcemanager.horizondb.HorizonDbManager manager) {
        manager.horizonDbClusters().start("exampleresourcegroup", "examplecluster", com.azure.core.util.Context.NONE);
    }
}
