
/**
 * Samples for Table Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-01/TableOperationPut.json
     */
    /**
     * Sample code: TableOperationPut.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void tableOperationPut(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getTables().createWithResponse("res3376", "sto328", "table6185", null,
            com.azure.core.util.Context.NONE);
    }
}
