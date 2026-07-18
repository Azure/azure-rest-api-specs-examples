
/**
 * Samples for DatabaseOperations ListByDatabase.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ListDatabaseOperations.json
     */
    /**
     * Sample code: List the database management operations.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void listTheDatabaseManagementOperations(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getDatabaseOperations().listByDatabase("Default-SQL-SouthEastAsia", "testsvr", "testdb",
            com.azure.core.util.Context.NONE);
    }
}
