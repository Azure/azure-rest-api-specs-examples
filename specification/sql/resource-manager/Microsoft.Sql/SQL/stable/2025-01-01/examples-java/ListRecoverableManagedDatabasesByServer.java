
/**
 * Samples for RecoverableManagedDatabases ListByInstance.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ListRecoverableManagedDatabasesByServer.json
     */
    /**
     * Sample code: List recoverable databases by managed instances.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        listRecoverableDatabasesByManagedInstances(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getRecoverableManagedDatabases().listByInstance("Test1", "managedInstance",
            com.azure.core.util.Context.NONE);
    }
}
