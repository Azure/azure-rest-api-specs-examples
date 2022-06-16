import com.azure.core.util.Context;

/** Samples for EncryptionScopes Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2021-08-01/examples/StorageAccountGetEncryptionScope.json
     */
    /**
     * Sample code: StorageAccountGetEncryptionScope.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void storageAccountGetEncryptionScope(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .storageAccounts()
            .manager()
            .serviceClient()
            .getEncryptionScopes()
            .getWithResponse("resource-group-name", "{storage-account-name}", "{encryption-scope-name}", Context.NONE);
    }
}
