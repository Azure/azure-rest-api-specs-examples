
/**
 * Samples for HorizonDbPrivateEndpointConnections List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/PrivateEndpointConnections_List.json
     */
    /**
     * Sample code: List all private endpoint connections on a cluster.
     * 
     * @param manager Entry point to HorizonDbManager.
     */
    public static void
        listAllPrivateEndpointConnectionsOnACluster(com.azure.resourcemanager.horizondb.HorizonDbManager manager) {
        manager.horizonDbPrivateEndpointConnections().list("exampleresourcegroup", "examplecluster",
            com.azure.core.util.Context.NONE);
    }
}
