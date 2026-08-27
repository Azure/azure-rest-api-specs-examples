
/**
 * Samples for HorizonDbClusters Restart.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/Clusters_Restart.json
     */
    /**
     * Sample code: Restart a HorizonDB cluster.
     * 
     * @param manager Entry point to HorizonDbManager.
     */
    public static void restartAHorizonDBCluster(com.azure.resourcemanager.horizondb.HorizonDbManager manager) {
        manager.horizonDbClusters().restart("exampleresourcegroup", "examplecluster", com.azure.core.util.Context.NONE);
    }
}
