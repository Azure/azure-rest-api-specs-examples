
/**
 * Samples for Endpoints List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/Endpoints_List.json
     */
    /**
     * Sample code: List HorizonDB endpoints in a cluster.
     * 
     * @param manager Entry point to HorizonDbManager.
     */
    public static void listHorizonDBEndpointsInACluster(com.azure.resourcemanager.horizondb.HorizonDbManager manager) {
        manager.endpoints().list("example-resource-group", "example-cluster", com.azure.core.util.Context.NONE);
    }
}
