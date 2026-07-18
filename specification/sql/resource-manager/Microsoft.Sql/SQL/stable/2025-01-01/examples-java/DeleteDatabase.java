
/**
 * Samples for Databases Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/DeleteDatabase.json
     */
    /**
     * Sample code: Deletes a database.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void deletesADatabase(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getDatabases().delete("Default-SQL-SouthEastAsia", "testsvr", "testdb",
            com.azure.core.util.Context.NONE);
    }
}
