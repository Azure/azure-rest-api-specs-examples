import com.azure.core.util.Context;
import com.azure.resourcemanager.storage.fluent.models.EncryptionScopeInner;
import com.azure.resourcemanager.storage.models.EncryptionScopeKeyVaultProperties;
import com.azure.resourcemanager.storage.models.EncryptionScopeSource;

/** Samples for EncryptionScopes Patch. */
public final class Main {
    /*
     * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2022-09-01/examples/StorageAccountPatchEncryptionScope.json
     */
    /**
     * Sample code: StorageAccountPatchEncryptionScope.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void storageAccountPatchEncryptionScope(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .storageAccounts()
            .manager()
            .serviceClient()
            .getEncryptionScopes()
            .patchWithResponse(
                "resource-group-name",
                "{storage-account-name}",
                "{encryption-scope-name}",
                new EncryptionScopeInner()
                    .withSource(EncryptionScopeSource.MICROSOFT_KEY_VAULT)
                    .withKeyVaultProperties(
                        new EncryptionScopeKeyVaultProperties()
                            .withKeyUri("https://testvault.vault.core.windows.net/keys/key1/863425f1358359c")),
                Context.NONE);
    }
}
