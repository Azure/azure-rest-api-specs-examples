
/**
 * Samples for DeletedAccounts List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01/DeletedAccountList.json
     */
    /**
     * Sample code: DeletedAccountList.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void deletedAccountList(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getDeletedAccounts().list(com.azure.core.util.Context.NONE);
    }
}
