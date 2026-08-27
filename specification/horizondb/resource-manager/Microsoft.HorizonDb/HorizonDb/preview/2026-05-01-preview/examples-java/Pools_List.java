
/**
 * Samples for HorizonDbPools List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/Pools_List.json
     */
    /**
     * Sample code: List HorizonDB pools in a cluster.
     * 
     * @param manager Entry point to HorizonDbManager.
     */
    public static void listHorizonDBPoolsInACluster(com.azure.resourcemanager.horizondb.HorizonDbManager manager) {
        manager.horizonDbPools().list("exampleresourcegroup", "examplecluster", com.azure.core.util.Context.NONE);
    }
}
