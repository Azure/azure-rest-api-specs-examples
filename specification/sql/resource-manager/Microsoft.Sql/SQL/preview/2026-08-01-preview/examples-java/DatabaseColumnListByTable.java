
/**
 * Samples for DatabaseColumns ListByTable.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-01-preview/DatabaseColumnListByTable.json
     */
    /**
     * Sample code: List database columns.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void listDatabaseColumns(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getDatabaseColumns().listByTable("myRG", "serverName", "myDatabase", "dbo", "table1",
            null, com.azure.core.util.Context.NONE);
    }
}
