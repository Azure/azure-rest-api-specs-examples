
/**
 * Samples for HorizonDbClusters Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/Clusters_Delete.json
     */
    /**
     * Sample code: Delete a HorizonDB cluster.
     * 
     * @param manager Entry point to HorizonDbManager.
     */
    public static void deleteAHorizonDBCluster(com.azure.resourcemanager.horizondb.HorizonDbManager manager) {
        manager.horizonDbClusters().delete("exampleresourcegroup", "examplecluster", com.azure.core.util.Context.NONE);
    }
}
