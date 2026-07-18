
/**
 * Samples for ManagedInstanceOperations Cancel.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/CancelManagedInstanceOperation.json
     */
    /**
     * Sample code: Cancel the managed instance management operation.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        cancelTheManagedInstanceManagementOperation(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getManagedInstanceOperations().cancelWithResponse("sqlcrudtest-7398",
            "sqlcrudtest-4645", "11111111-1111-1111-1111-111111111111", com.azure.core.util.Context.NONE);
    }
}
