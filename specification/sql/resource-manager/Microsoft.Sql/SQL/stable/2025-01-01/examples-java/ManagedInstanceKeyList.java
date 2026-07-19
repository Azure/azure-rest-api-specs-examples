
/**
 * Samples for ManagedInstanceKeys ListByInstance.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ManagedInstanceKeyList.json
     */
    /**
     * Sample code: List the keys for a managed instance.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void listTheKeysForAManagedInstance(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getManagedInstanceKeys().listByInstance("sqlcrudtest-7398", "sqlcrudtest-4645", null,
            com.azure.core.util.Context.NONE);
    }
}
