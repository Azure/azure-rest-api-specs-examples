
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/mysql/resource-manager/Microsoft.DBforMySQL/ServiceOperations/stable/2023-12-30/examples/
     * OperationsList.json
     */
    /**
     * Sample code: OperationList.
     * 
     * @param manager Entry point to MySqlManager.
     */
    public static void operationList(com.azure.resourcemanager.mysqlflexibleserver.MySqlManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
