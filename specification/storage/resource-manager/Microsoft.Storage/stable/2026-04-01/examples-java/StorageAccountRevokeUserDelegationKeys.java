
/**
 * Samples for StorageAccounts RevokeUserDelegationKeys.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01/StorageAccountRevokeUserDelegationKeys.json
     */
    /**
     * Sample code: StorageAccountRevokeUserDelegationKeys.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void
        storageAccountRevokeUserDelegationKeys(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getStorageAccounts().revokeUserDelegationKeysWithResponse("res4167", "sto3539",
            com.azure.core.util.Context.NONE);
    }
}
