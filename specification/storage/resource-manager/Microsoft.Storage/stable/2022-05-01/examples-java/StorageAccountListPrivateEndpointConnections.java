import com.azure.core.util.Context;

/** Samples for PrivateEndpointConnections List. */
public final class Main {
    /*
     * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2022-05-01/examples/StorageAccountListPrivateEndpointConnections.json
     */
    /**
     * Sample code: StorageAccountListPrivateEndpointConnections.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void storageAccountListPrivateEndpointConnections(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .storageAccounts()
            .manager()
            .serviceClient()
            .getPrivateEndpointConnections()
            .list("res6977", "sto2527", Context.NONE);
    }
}
