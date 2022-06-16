import com.azure.core.util.Context;

/** Samples for StorageAccounts ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2021-09-01/examples/StorageAccountListByResourceGroup.json
     */
    /**
     * Sample code: StorageAccountListByResourceGroup.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void storageAccountListByResourceGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .storageAccounts()
            .manager()
            .serviceClient()
            .getStorageAccounts()
            .listByResourceGroup("res6117", Context.NONE);
    }
}
