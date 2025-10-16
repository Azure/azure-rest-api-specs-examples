
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/operations_list.json
     */
    /**
     * Sample code: Operations_list.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void operationsList(com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
