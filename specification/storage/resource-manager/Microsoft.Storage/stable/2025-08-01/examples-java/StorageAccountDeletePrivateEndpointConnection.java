
/**
 * Samples for PrivateEndpointConnections Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-01/StorageAccountDeletePrivateEndpointConnection.json
     */
    /**
     * Sample code: StorageAccountDeletePrivateEndpointConnection.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void
        storageAccountDeletePrivateEndpointConnection(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getPrivateEndpointConnections().deleteWithResponse("res6977", "sto2527",
            "{privateEndpointConnectionName}", com.azure.core.util.Context.NONE);
    }
}
