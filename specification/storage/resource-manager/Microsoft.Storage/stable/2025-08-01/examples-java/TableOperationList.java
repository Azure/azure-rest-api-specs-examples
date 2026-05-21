
/**
 * Samples for Table List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-01/TableOperationList.json
     */
    /**
     * Sample code: TableOperationList.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void tableOperationList(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getTables().list("res9290", "sto328", com.azure.core.util.Context.NONE);
    }
}
