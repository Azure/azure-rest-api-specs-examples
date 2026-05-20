
/**
 * Samples for DataShares Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-01/StorageDataShareCRUD/StorageDataShares_Delete.json
     */
    /**
     * Sample code: DeleteDataShare.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void deleteDataShare(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getDataShares().delete("testrg", "teststorageaccount", "testdatashare",
            com.azure.core.util.Context.NONE);
    }
}
