
/**
 * Samples for PrivateEndpointConnections Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-01/PrivateEndpointConnections_Delete.json
     */
    /**
     * Sample code: PrivateEndpointConnections_Delete.
     * 
     * @param manager Entry point to FileSharesManager.
     */
    public static void
        privateEndpointConnectionsDelete(com.azure.resourcemanager.fileshares.FileSharesManager manager) {
        manager.privateEndpointConnections().delete("rgfileshares", "fileshare", "privateEndpointConnection1",
            com.azure.core.util.Context.NONE);
    }
}
