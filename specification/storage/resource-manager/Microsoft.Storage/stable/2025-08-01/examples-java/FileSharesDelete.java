
/**
 * Samples for FileShares Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-01/FileSharesDelete.json
     */
    /**
     * Sample code: DeleteShares.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void deleteShares(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getFileShares().deleteWithResponse("res4079", "sto4506", "share9689", null, null,
            com.azure.core.util.Context.NONE);
    }
}
