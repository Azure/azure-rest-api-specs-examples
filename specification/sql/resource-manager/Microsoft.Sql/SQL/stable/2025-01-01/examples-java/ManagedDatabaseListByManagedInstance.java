
/**
 * Samples for ManagedDatabases ListByInstance.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ManagedDatabaseListByManagedInstance.json
     */
    /**
     * Sample code: List databases by managed instances.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void listDatabasesByManagedInstances(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getManagedDatabases().listByInstance("Test1", "managedInstance",
            com.azure.core.util.Context.NONE);
    }
}
