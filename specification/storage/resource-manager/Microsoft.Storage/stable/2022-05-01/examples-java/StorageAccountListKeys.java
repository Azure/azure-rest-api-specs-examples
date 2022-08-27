import com.azure.core.util.Context;

/** Samples for StorageAccounts ListKeys. */
public final class Main {
    /*
     * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2022-05-01/examples/StorageAccountListKeys.json
     */
    /**
     * Sample code: StorageAccountListKeys.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void storageAccountListKeys(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .storageAccounts()
            .manager()
            .serviceClient()
            .getStorageAccounts()
            .listKeysWithResponse("res418", "sto2220", null, Context.NONE);
    }
}
