
/**
 * Samples for OperationProgress Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/mysql/resource-manager/Microsoft.DBforMySQL/ServiceOperations/stable/2023-12-30/examples/
     * OperationProgress_Get_ImportFromStorage.json
     */
    /**
     * Sample code: OperationProgress_Get ImportFromStorage.
     * 
     * @param manager Entry point to MySqlManager.
     */
    public static void
        operationProgressGetImportFromStorage(com.azure.resourcemanager.mysqlflexibleserver.MySqlManager manager) {
        manager.operationProgress().getWithResponse("westus", "00000000-0000-0000-0000-000000000000",
            com.azure.core.util.Context.NONE);
    }
}
