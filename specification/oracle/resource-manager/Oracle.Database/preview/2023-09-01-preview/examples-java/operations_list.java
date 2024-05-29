
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/oracle/resource-manager/Oracle.Database/preview/2023-09-01-preview/examples/operations_list.json
     */
    /**
     * Sample code: List Operations.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void listOperations(com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
