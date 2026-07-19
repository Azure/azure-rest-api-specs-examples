
/**
 * Samples for RestorableDroppedManagedDatabases ListByInstance.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/RestorableDroppedManagedDatabaseListByManagedInstance.json
     */
    /**
     * Sample code: List restorable dropped databases by managed instances.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        listRestorableDroppedDatabasesByManagedInstances(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getRestorableDroppedManagedDatabases().listByInstance("Test1", "managedInstance",
            com.azure.core.util.Context.NONE);
    }
}
