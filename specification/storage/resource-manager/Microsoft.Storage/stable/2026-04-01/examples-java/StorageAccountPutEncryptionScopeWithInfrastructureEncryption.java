
import com.azure.resourcemanager.storage.fluent.models.EncryptionScopeInner;

/**
 * Samples for EncryptionScopes Put.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01/StorageAccountPutEncryptionScopeWithInfrastructureEncryption.json
     */
    /**
     * Sample code: StorageAccountPutEncryptionScopeWithInfrastructureEncryption.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void storageAccountPutEncryptionScopeWithInfrastructureEncryption(
        com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getEncryptionScopes().putWithResponse("resource-group-name", "accountname",
            "{encryption-scope-name}", new EncryptionScopeInner().withRequireInfrastructureEncryption(true),
            com.azure.core.util.Context.NONE);
    }
}
