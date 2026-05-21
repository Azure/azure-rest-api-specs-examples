
/**
 * Samples for EncryptionScopes List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-01/StorageAccountEncryptionScopeList.json
     */
    /**
     * Sample code: StorageAccountEncryptionScopeList.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void storageAccountEncryptionScopeList(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getEncryptionScopes().list("resource-group-name", "accountname", null, null, null,
            com.azure.core.util.Context.NONE);
    }
}
