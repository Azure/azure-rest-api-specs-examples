
/**
 * Samples for Clusters Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/Clusters_Delete.json
     */
    /**
     * Sample code: Delete a HorizonDB cluster.
     * 
     * @param manager Entry point to HorizonDbManager.
     */
    public static void deleteAHorizonDBCluster(com.azure.resourcemanager.horizondb.HorizonDbManager manager) {
        manager.clusters().delete("example-resource-group", "example-cluster", com.azure.core.util.Context.NONE);
    }
}
