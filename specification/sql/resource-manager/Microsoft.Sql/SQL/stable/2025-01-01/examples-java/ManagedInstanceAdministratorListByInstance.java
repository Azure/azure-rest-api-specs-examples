
/**
 * Samples for ManagedInstanceAdministrators ListByInstance.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ManagedInstanceAdministratorListByInstance.json
     */
    /**
     * Sample code: List administrators of managed instance.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void listAdministratorsOfManagedInstance(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getManagedInstanceAdministrators().listByInstance("Default-SQL-SouthEastAsia",
            "managedInstance", com.azure.core.util.Context.NONE);
    }
}
