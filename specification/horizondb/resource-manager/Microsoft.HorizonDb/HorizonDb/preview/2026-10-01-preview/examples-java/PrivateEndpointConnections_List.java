
/**
 * Samples for PrivateEndpointConnections List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/PrivateEndpointConnections_List.json
     */
    /**
     * Sample code: List private endpoint connections on a HorizonDB cluster.
     * 
     * @param manager Entry point to HorizonDbManager.
     */
    public static void listPrivateEndpointConnectionsOnAHorizonDBCluster(
        com.azure.resourcemanager.horizondb.HorizonDbManager manager) {
        manager.privateEndpointConnections().list("example-resource-group", "example-cluster",
            com.azure.core.util.Context.NONE);
    }
}
