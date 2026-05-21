
/**
 * Samples for Table Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-01/TableOperationPatch.json
     */
    /**
     * Sample code: TableOperationPatch.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void tableOperationPatch(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getTables().updateWithResponse("res3376", "sto328", "table6185", null,
            com.azure.core.util.Context.NONE);
    }
}
