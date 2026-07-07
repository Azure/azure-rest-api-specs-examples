
/**
 * Samples for Table Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01/TableOperationDelete.json
     */
    /**
     * Sample code: TableOperationDelete.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void tableOperationDelete(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getTables().deleteWithResponse("res3376", "sto328", "table6185",
            com.azure.core.util.Context.NONE);
    }
}
