
/**
 * Samples for PrivateEndpointConnections Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2023-05-01/examples/
     * StorageAccountDeletePrivateEndpointConnection.json
     */
    /**
     * Sample code: StorageAccountDeletePrivateEndpointConnection.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        storageAccountDeletePrivateEndpointConnection(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.storageAccounts().manager().serviceClient().getPrivateEndpointConnections().deleteWithResponse("res6977",
            "sto2527", "{privateEndpointConnectionName}", com.azure.core.util.Context.NONE);
    }
}
