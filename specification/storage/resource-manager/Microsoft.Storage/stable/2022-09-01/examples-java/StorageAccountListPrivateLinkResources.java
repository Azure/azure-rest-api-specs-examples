import com.azure.core.util.Context;

/** Samples for PrivateLinkResources ListByStorageAccount. */
public final class Main {
    /*
     * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2022-09-01/examples/StorageAccountListPrivateLinkResources.json
     */
    /**
     * Sample code: StorageAccountListPrivateLinkResources.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void storageAccountListPrivateLinkResources(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .storageAccounts()
            .manager()
            .serviceClient()
            .getPrivateLinkResources()
            .listByStorageAccountWithResponse("res6977", "sto2527", Context.NONE);
    }
}
