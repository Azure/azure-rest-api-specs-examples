
/**
 * Samples for DatabaseColumns ListByDatabase.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ColumnsListByDatabaseMin.json
     */
    /**
     * Sample code: List database columns.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void listDatabaseColumns(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getDatabaseColumns().listByDatabase("myRG", "serverName", "myDatabase", null, null,
            null, null, null, com.azure.core.util.Context.NONE);
    }
}
