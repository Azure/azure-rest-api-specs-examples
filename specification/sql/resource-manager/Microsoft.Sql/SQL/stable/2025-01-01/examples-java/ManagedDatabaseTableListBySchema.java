
/**
 * Samples for ManagedDatabaseTables ListBySchema.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ManagedDatabaseTableListBySchema.json
     */
    /**
     * Sample code: List managed database tables.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void listManagedDatabaseTables(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getManagedDatabaseTables().listBySchema("myRG", "myManagedInstanceName", "myDatabase",
            "dbo", null, com.azure.core.util.Context.NONE);
    }
}
