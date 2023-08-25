/** Samples for EncryptionScopes List. */
public final class Main {
    /*
     * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2023-01-01/examples/StorageAccountEncryptionScopeList.json
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
            .list("resource-group-name", "accountname", null, null, null, com.azure.core.util.Context.NONE);
    }
}
