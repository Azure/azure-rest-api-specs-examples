
/**
 * Samples for HorizonDbReplicas Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/Replicas_Delete.json
     */
    /**
     * Sample code: Delete a HorizonDB replica.
     * 
     * @param manager Entry point to HorizonDbManager.
     */
    public static void deleteAHorizonDBReplica(com.azure.resourcemanager.horizondb.HorizonDbManager manager) {
        manager.horizonDbReplicas().delete("exampleresourcegroup", "examplecluster", "examplepool", "examplereplica",
            com.azure.core.util.Context.NONE);
    }
}
