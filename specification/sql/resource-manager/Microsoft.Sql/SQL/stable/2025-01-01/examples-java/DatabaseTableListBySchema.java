
/**
 * Samples for DatabaseTables ListBySchema.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/DatabaseTableListBySchema.json
     */
    /**
     * Sample code: List database tables.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void listDatabaseTables(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getDatabaseTables().listBySchema("myRG", "serverName", "myDatabase", "dbo", null,
            com.azure.core.util.Context.NONE);
    }
}
