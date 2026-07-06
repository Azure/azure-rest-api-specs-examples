
import com.azure.resourcemanager.storage.fluent.models.PrivateEndpointConnectionInner;
import com.azure.resourcemanager.storage.models.PrivateEndpointServiceConnectionStatus;
import com.azure.resourcemanager.storage.models.PrivateLinkServiceConnectionState;

/**
 * Samples for PrivateEndpointConnections Put.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01/StorageAccountPutPrivateEndpointConnection.json
     */
    /**
     * Sample code: StorageAccountPutPrivateEndpointConnection.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void
        storageAccountPutPrivateEndpointConnection(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getPrivateEndpointConnections().putWithResponse("res7687", "sto9699",
            "{privateEndpointConnectionName}",
            new PrivateEndpointConnectionInner().withPrivateLinkServiceConnectionState(
                new PrivateLinkServiceConnectionState().withStatus(PrivateEndpointServiceConnectionStatus.APPROVED)
                    .withDescription("Auto-Approved")),
            com.azure.core.util.Context.NONE);
    }
}
