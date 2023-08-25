import com.azure.resourcemanager.storage.fluent.models.EncryptionScopeInner;

/** Samples for EncryptionScopes Put. */
public final class Main {
    /*
     * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2023-01-01/examples/StorageAccountPutEncryptionScope.json
     */
    /**
     * Sample code: StorageAccountPutEncryptionScope.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void storageAccountPutEncryptionScope(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .storageAccounts()
            .manager()
            .serviceClient()
            .getEncryptionScopes()
            .putWithResponse(
                "resource-group-name",
                "accountname",
                "{encryption-scope-name}",
                new EncryptionScopeInner(),
                com.azure.core.util.Context.NONE);
    }
}
