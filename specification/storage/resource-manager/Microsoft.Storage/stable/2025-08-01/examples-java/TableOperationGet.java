
/**
 * Samples for Table Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-01/TableOperationGet.json
     */
    /**
     * Sample code: TableOperationGet.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void tableOperationGet(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getTables().getWithResponse("res3376", "sto328", "table6185",
            com.azure.core.util.Context.NONE);
    }
}
