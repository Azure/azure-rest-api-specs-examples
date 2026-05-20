
/**
 * Samples for FileShares List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-01/DeletedFileSharesList.json
     */
    /**
     * Sample code: ListDeletedShares.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void listDeletedShares(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getFileShares().list("res9290", "sto1590", null, null, "deleted",
            com.azure.core.util.Context.NONE);
    }
}
