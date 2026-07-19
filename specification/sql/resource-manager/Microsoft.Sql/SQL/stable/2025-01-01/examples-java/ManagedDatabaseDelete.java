
/**
 * Samples for ManagedDatabases Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ManagedDatabaseDelete.json
     */
    /**
     * Sample code: Delete managed database.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void deleteManagedDatabase(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getManagedDatabases().delete("Default-SQL-SouthEastAsia", "managedInstance", "testdb",
            com.azure.core.util.Context.NONE);
    }
}
