
/**
 * Samples for PrivateEndpointConnections Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/PrivateEndpointConnections_Delete.json
     */
    /**
     * Sample code: Delete a private endpoint connection.
     * 
     * @param manager Entry point to HorizonDbManager.
     */
    public static void deleteAPrivateEndpointConnection(com.azure.resourcemanager.horizondb.HorizonDbManager manager) {
        manager.privateEndpointConnections().delete("example-resource-group", "example-cluster",
            "example-private-endpoint-connection", com.azure.core.util.Context.NONE);
    }
}
