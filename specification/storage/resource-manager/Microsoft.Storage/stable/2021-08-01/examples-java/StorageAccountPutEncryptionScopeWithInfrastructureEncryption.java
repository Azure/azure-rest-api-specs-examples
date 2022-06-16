import com.azure.core.util.Context;
import com.azure.resourcemanager.storage.fluent.models.EncryptionScopeInner;

/** Samples for EncryptionScopes Put. */
public final class Main {
    /*
     * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2021-08-01/examples/StorageAccountPutEncryptionScopeWithInfrastructureEncryption.json
     */
    /**
     * Sample code: StorageAccountPutEncryptionScopeWithInfrastructureEncryption.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void storageAccountPutEncryptionScopeWithInfrastructureEncryption(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .storageAccounts()
            .manager()
            .serviceClient()
            .getEncryptionScopes()
            .putWithResponse(
                "resource-group-name",
                "{storage-account-name}",
                "{encryption-scope-name}",
                new EncryptionScopeInner().withRequireInfrastructureEncryption(true),
                Context.NONE);
    }
}
