import com.azure.core.util.Context;

/** Samples for EncryptionScopes List. */
public final class Main {
    /*
     * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2022-09-01/examples/StorageAccountEncryptionScopeList.json
     */
    /**
     * Sample code: StorageAccountEncryptionScopeList.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void storageAccountEncryptionScopeList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .storageAccounts()
            .manager()
            .serviceClient()
            .getEncryptionScopes()
            .list("resource-group-name", "{storage-account-name}", null, null, null, Context.NONE);
    }
}
