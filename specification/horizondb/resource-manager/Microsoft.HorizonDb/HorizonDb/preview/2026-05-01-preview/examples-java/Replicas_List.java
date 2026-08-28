
/**
 * Samples for HorizonDbReplicas List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/Replicas_List.json
     */
    /**
     * Sample code: List HorizonDB replicas in a pool.
     * 
     * @param manager Entry point to HorizonDbManager.
     */
    public static void listHorizonDBReplicasInAPool(com.azure.resourcemanager.horizondb.HorizonDbManager manager) {
        manager.horizonDbReplicas().list("exampleresourcegroup", "examplecluster", "examplepool",
            com.azure.core.util.Context.NONE);
    }
}
