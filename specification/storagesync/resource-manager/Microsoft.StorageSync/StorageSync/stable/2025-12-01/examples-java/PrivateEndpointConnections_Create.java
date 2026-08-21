
import com.azure.resourcemanager.storagesync.models.PrivateEndpointServiceConnectionStatus;
import com.azure.resourcemanager.storagesync.models.PrivateLinkServiceConnectionState;

/**
 * Samples for PrivateEndpointConnections Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-12-01/PrivateEndpointConnections_Create.json
     */
    /**
     * Sample code: PrivateEndpointConnections_Create.
     * 
     * @param manager Entry point to StorageSyncManager.
     */
    public static void
        privateEndpointConnectionsCreate(com.azure.resourcemanager.storagesync.StorageSyncManager manager) {
        manager.privateEndpointConnections().define("{privateEndpointConnectionName}")
            .withExistingStorageSyncService("res7687", "sss2527")
            .withPrivateLinkServiceConnectionState(new PrivateLinkServiceConnectionState()
                .withStatus(PrivateEndpointServiceConnectionStatus.APPROVED).withDescription("Auto-Approved"))
            .create();
    }
}
