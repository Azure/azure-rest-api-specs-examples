
/**
 * Samples for ManagedDatabaseColumns ListByDatabase.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ManagedColumnsListByDatabaseMin.json
     */
    /**
     * Sample code: List managed database columns.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void listManagedDatabaseColumns(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getManagedDatabaseColumns().listByDatabase("myRG", "serverName", "myDatabase", null,
            null, null, null, null, com.azure.core.util.Context.NONE);
    }
}
