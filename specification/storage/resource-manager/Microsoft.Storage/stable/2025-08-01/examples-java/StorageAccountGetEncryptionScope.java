
/**
 * Samples for EncryptionScopes Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-01/StorageAccountGetEncryptionScope.json
     */
    /**
     * Sample code: StorageAccountGetEncryptionScope.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void storageAccountGetEncryptionScope(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getEncryptionScopes().getWithResponse("resource-group-name", "accountname",
            "{encryption-scope-name}", com.azure.core.util.Context.NONE);
    }
}
