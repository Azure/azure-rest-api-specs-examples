
/**
 * Samples for PrivateEndpointConnections List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01/StorageAccountListPrivateEndpointConnections.json
     */
    /**
     * Sample code: StorageAccountListPrivateEndpointConnections.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void
        storageAccountListPrivateEndpointConnections(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getPrivateEndpointConnections().list("res6977", "sto2527",
            com.azure.core.util.Context.NONE);
    }
}
