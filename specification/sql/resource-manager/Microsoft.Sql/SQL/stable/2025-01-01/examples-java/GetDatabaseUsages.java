
/**
 * Samples for DatabaseUsages ListByDatabase.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/GetDatabaseUsages.json
     */
    /**
     * Sample code: Gets database usages.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void getsDatabaseUsages(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getDatabaseUsages().listByDatabase("Default-SQL-SouthEastAsia", "testsvr", "testdb",
            com.azure.core.util.Context.NONE);
    }
}
