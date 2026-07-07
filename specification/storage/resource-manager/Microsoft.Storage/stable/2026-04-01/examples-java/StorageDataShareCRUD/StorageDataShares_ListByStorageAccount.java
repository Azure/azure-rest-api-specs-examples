
/**
 * Samples for DataShares ListByStorageAccount.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01/StorageDataShareCRUD/StorageDataShares_ListByStorageAccount.json
     */
    /**
     * Sample code: ListDataSharesByStorageAccount.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void listDataSharesByStorageAccount(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getDataShares().listByStorageAccount("testrg", "teststorageaccount",
            com.azure.core.util.Context.NONE);
    }
}
