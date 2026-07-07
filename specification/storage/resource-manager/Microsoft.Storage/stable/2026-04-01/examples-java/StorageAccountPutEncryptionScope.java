
import com.azure.resourcemanager.storage.fluent.models.EncryptionScopeInner;

/**
 * Samples for EncryptionScopes Put.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01/StorageAccountPutEncryptionScope.json
     */
    /**
     * Sample code: StorageAccountPutEncryptionScope.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void storageAccountPutEncryptionScope(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getEncryptionScopes().putWithResponse("resource-group-name", "accountname",
            "{encryption-scope-name}", new EncryptionScopeInner(), com.azure.core.util.Context.NONE);
    }
}
