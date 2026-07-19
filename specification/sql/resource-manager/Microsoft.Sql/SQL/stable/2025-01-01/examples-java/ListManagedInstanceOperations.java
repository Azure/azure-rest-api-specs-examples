
/**
 * Samples for ManagedInstanceOperations ListByManagedInstance.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ListManagedInstanceOperations.json
     */
    /**
     * Sample code: List the managed instance management operations.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        listTheManagedInstanceManagementOperations(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getManagedInstanceOperations().listByManagedInstance("sqlcrudtest-7398",
            "sqlcrudtest-4645", com.azure.core.util.Context.NONE);
    }
}
