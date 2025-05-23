
import com.azure.resourcemanager.storage.fluent.models.PrivateEndpointConnectionInner;
import com.azure.resourcemanager.storage.models.PrivateEndpointServiceConnectionStatus;
import com.azure.resourcemanager.storage.models.PrivateLinkServiceConnectionState;

/**
 * Samples for PrivateEndpointConnections Put.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2024-01-01/examples/
     * StorageAccountPutPrivateEndpointConnection.json
     */
    /**
     * Sample code: StorageAccountPutPrivateEndpointConnection.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        storageAccountPutPrivateEndpointConnection(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.storageAccounts().manager().serviceClient().getPrivateEndpointConnections().putWithResponse("res7687",
            "sto9699", "{privateEndpointConnectionName}",
            new PrivateEndpointConnectionInner().withPrivateLinkServiceConnectionState(
                new PrivateLinkServiceConnectionState().withStatus(PrivateEndpointServiceConnectionStatus.APPROVED)
                    .withDescription("Auto-Approved")),
            com.azure.core.util.Context.NONE);
    }
}
