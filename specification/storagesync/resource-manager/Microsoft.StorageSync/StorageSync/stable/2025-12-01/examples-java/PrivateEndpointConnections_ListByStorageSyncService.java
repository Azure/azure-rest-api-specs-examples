
/**
 * Samples for PrivateEndpointConnections ListByStorageSyncService.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-12-01/PrivateEndpointConnections_ListByStorageSyncService.json
     */
    /**
     * Sample code: PrivateEndpointConnections_ListByStorageSyncService.
     * 
     * @param manager Entry point to StorageSyncManager.
     */
    public static void privateEndpointConnectionsListByStorageSyncService(
        com.azure.resourcemanager.storagesync.StorageSyncManager manager) {
        manager.privateEndpointConnections().listByStorageSyncService("res6977", "sss2527",
            com.azure.core.util.Context.NONE);
    }
}
