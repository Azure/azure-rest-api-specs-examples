
/**
 * Samples for DataShares Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-01/StorageDataShareCRUD/StorageDataShares_Get.json
     */
    /**
     * Sample code: GetDataShare.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void getDataShare(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getDataShares().getWithResponse("testrg", "teststorageaccount", "testdatashare",
            com.azure.core.util.Context.NONE);
    }
}
