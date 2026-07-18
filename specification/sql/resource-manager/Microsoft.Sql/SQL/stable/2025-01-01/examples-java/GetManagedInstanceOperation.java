
/**
 * Samples for ManagedInstanceOperations Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/GetManagedInstanceOperation.json
     */
    /**
     * Sample code: Gets the managed instance management operation.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        getsTheManagedInstanceManagementOperation(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getManagedInstanceOperations().getWithResponse("sqlcrudtest-7398", "sqlcrudtest-4645",
            "00000000-1111-2222-3333-444444444444", com.azure.core.util.Context.NONE);
    }
}
