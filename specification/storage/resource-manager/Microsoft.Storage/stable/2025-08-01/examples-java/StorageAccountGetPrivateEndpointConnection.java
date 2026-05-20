
/**
 * Samples for PrivateEndpointConnections Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-01/StorageAccountGetPrivateEndpointConnection.json
     */
    /**
     * Sample code: StorageAccountGetPrivateEndpointConnection.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void
        storageAccountGetPrivateEndpointConnection(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getPrivateEndpointConnections().getWithResponse("res6977", "sto2527",
            "{privateEndpointConnectionName}", com.azure.core.util.Context.NONE);
    }
}
