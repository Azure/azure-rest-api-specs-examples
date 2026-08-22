
/**
 * Samples for PrivateEndpointConnections Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-12-01/PrivateEndpointConnections_Delete.json
     */
    /**
     * Sample code: PrivateEndpointConnections_Delete.
     * 
     * @param manager Entry point to StorageSyncManager.
     */
    public static void
        privateEndpointConnectionsDelete(com.azure.resourcemanager.storagesync.StorageSyncManager manager) {
        manager.privateEndpointConnections().delete("res6977", "sss2527", "{privateEndpointConnectionName}",
            com.azure.core.util.Context.NONE);
    }
}
