import com.azure.core.util.Context;

/** Samples for PrivateEndpointConnections Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2021-04-01/examples/StorageAccountGetPrivateEndpointConnection.json
     */
    /**
     * Sample code: StorageAccountGetPrivateEndpointConnection.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void storageAccountGetPrivateEndpointConnection(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .storageAccounts()
            .manager()
            .serviceClient()
            .getPrivateEndpointConnections()
            .getWithResponse("res6977", "sto2527", "{privateEndpointConnectionName}", Context.NONE);
    }
}
